package registry

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
)

// TarLayer tar a docker image layer
func TarLayer(out io.Writer, root string, fileList []string) error {

	fileList = SortFiles4Layer(fileList)

	gzipOut := gzip.NewWriter(out)
	defer gzipOut.Close()
	tarOut := tar.NewWriter(gzipOut)
	defer tarOut.Close()

	for _, file := range fileList {

		err := func() error {

			path, err := filepath.Rel(root, file)
			if err != nil {
				return errors.Wrap(err, "calc rel path of root")
			}

			finfo, err := os.Stat(file)
			if err != nil {
				return errors.Wrap(err, "get file stat")
			}

			tarHdr, err := tar.FileInfoHeader(finfo, file)
			if err != nil {
				return errors.Wrap(err, "parsing tar header")
			}
			tarHdr.Name = path

			err = tarOut.WriteHeader(tarHdr)
			if err != nil {
				return errors.Wrap(err, "write tar file header")
			}

			if finfo.IsDir() {
				return nil
			}

			f, err := os.Open(file)
			if err != nil {
				return errors.Wrap(err, "open layer file")
			}
			defer f.Close()

			_, err = io.Copy(tarOut, f)
			if err != nil {
				return errors.Wrap(err, "tar layer file")
			}

			return nil
		}()
		if err != nil {
			return err
		}
	}

	return nil
}

// UntarLayer untar a layer to tar.Reader
func UntarLayer(in io.Reader) (*tar.Reader, error) {

	gzipIn, err := gzip.NewReader(in)
	if err != nil {
		return nil, errors.Wrap(err, "un-gzip")
	}

	return tar.NewReader(gzipIn), nil
}

// TarLayerAndDigest tar a new layer and return its digest
func TarLayerAndDigest(
	out io.Writer, root string, fileList []string) (digest.Digest, error) {

	digester := digest.Canonical.Digester()
	err := TarLayer(io.MultiWriter(out, digester.Hash()), root, fileList)

	return digester.Digest(), err
}

// SortFiles4Layer sort file list for layer
func SortFiles4Layer(fileList []string) []string {

	sorted := make([]string, len(fileList))
	copy(sorted, fileList)
	sort.Strings(sorted)

	return sorted
}

// TarLayerDirectory tar the whole directory to a layer
func TarLayerDirectory(
	out io.Writer, root string, ignores ...string) (digest.Digest, error) {

	// calculate absolute root path
	var err error
	root, err = filepath.Abs(root)
	if err != nil {
		return "", errors.Wrap(err, "calc absolute path of root")
	}

	// calculate relative path of ignores to root
	for i := 0; i < len(ignores); i++ {

		if !filepath.IsAbs(ignores[i]) {
			ignores[i], err = filepath.Abs(ignores[i])
			if err != nil {
				return "", errors.Wrap(err, "calc aboluste path of ignores")
			}
		}

		ignores[i], err = filepath.Rel(root, ignores[i])
		if err != nil {
			return "", errors.Wrap(err, "calc ignored related path to root")
		}
	}
	sort.Strings(ignores)

	// walk all files
	fileList := []string{}
	err = filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return errors.Wrap(err, "calc related layer path")
			}

			position := sort.SearchStrings(ignores, relPath)
			if position < len(ignores) && ignores[position] == relPath {

				if info.IsDir() {
					return filepath.SkipDir
				}

				return nil
			}

			fileList = append(fileList, path)
			return nil
		},
	)
	if err != nil {
		return "", err
	}

	return TarLayerAndDigest(out, root, fileList)
}

// UntarLayerDirectory untar a layer to directory
func UntarLayerDirectory(in io.Reader, root string) error {

	reader, err := UntarLayer(in)
	if err != nil {
		return err
	}

	for {

		header, err := reader.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "untar file header")
		}

		err = func() error {

			// FIXME using fileInfo to init physical file
			fileInfo := header.FileInfo()
			filePath := filepath.Join(root, header.Name)

			if fileInfo.IsDir() {
				err = os.MkdirAll(filePath, os.ModePerm)
				if err != nil {
					return errors.Wrap(err, "make layer directory")
				}
				return nil
			}

			file, err := os.Create(filePath)
			if err != nil {
				return errors.Wrap(err, "create layer file")
			}
			defer file.Close()

			_, err = io.Copy(file, reader)
			if err != nil {
				return errors.Wrap(err, "write layer file")
			}

			return nil
		}()
		if err != nil {
			return err
		}
	}
}
