package commit

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/docker/distribution"
	manifestV1 "github.com/docker/distribution/manifest/schema1"
	manifestv2 "github.com/docker/distribution/manifest/schema2"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/setekhid/ketos/pkg/registry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	Command = &cobra.Command{
		Use:   "commit",
		Short: "commit [--tag|-t tag-name]",
		Args:  cobra.NoArgs,

		RunE: commitMain,
	}
)

func init() {

	flags := Command.Flags()
	flags.StringP("tag", "t", "", "commit to a tag")
}

type blackHole struct{}

func (_ blackHole) Write(p []byte) (int, error) { return len(p), nil }
func (_ blackHole) Read(p []byte) (int, error)  { return 0, nil }
func (_ blackHole) Close() error                { return nil }

func commitMain(cmd *cobra.Command, args []string) error {

	tag, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	meta, err := metadata.CurrentMetadatas()
	if err != nil {
		return err
	}

	containerPath := meta.ContainerPath()
	metaFolderPath := meta.MetaFolderPath()

	// calculate layer digest
	digest, err := registry.TarLayerDirectory(blackHole{},
		containerPath, metaFolderPath)
	if err != nil {
		return err
	}

	layerPath := meta.LayerPath(digest)
	ignorePath, err := filepath.Rel(containerPath, metaFolderPath)
	if err != nil {
		return errors.Wrap(err, "calc metadata related path")
	}

	// archive container files to layer
	err = filepath.Walk(
		containerPath,
		func(path string, info os.FileInfo, err error) error {

			relPath, err := filepath.Rel(containerPath, path)
			if err != nil {
				return errors.Wrap(err, "calc related path")
			}

			baseName := filepath.Base(path)
			newPath := filepath.Join(layerPath, baseName)

			if relPath == "." || relPath == ignorePath {
				return nil
			}

			err = os.Rename(path, newPath)
			if err != nil {
				return errors.Wrapf(err, "move %v to %v", path, newPath)
			}

			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		},
	)

	// update manifest
	manifest, err := meta.GetManifest(tag)
	if err != nil {
		return err
	}

	manifest.FSLayers = append(manifest.FSLayers,
		manifestV1.FSLayer{
			BlobSum: digest,
		})

	return meta.PutManifest(tag, manifest)
}

func commitMain0(cmd *cobra.Command, args []string) error {

	commitTag, err := cmd.Flags().GetString("tag")
	if err != nil {
		return errors.Wrap(err, "parsing tag")
	}

	ketosFolder, err := metadata.KetosFolder()
	if err != nil {
		return errors.Wrap(err, "seek ketos folder")
	}

	buff := &bytes.Buffer{}
	err = gztarKetosContainer(filepath.Dir(ketosFolder), buff)
	if err != nil {
		return errors.Wrap(err, "gzip container layer")
	}

	contentSize := buff.Len()
	digestNumber := digest.FromBytes(buff.Bytes())
	layerTar, err := os.Create(filepath.Join(ketosFolder, "layers", digestNumber.Encoded()+".tar.gz"))
	if err != nil {
		return errors.Wrap(err, "create layer.tar.gz")
	}
	defer layerTar.Close()
	_, err = io.Copy(layerTar, buff)
	if err != nil {
		return errors.Wrap(err, "write down layer.tar.gz")
	}

	archiveSrc := filepath.Dir(ketosFolder)
	archiveDest := filepath.Join(ketosFolder, "layers", digestNumber.Encoded())
	err = os.MkdirAll(archiveDest, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "make layer directory")
	}
	err = archiveKetosContainer(archiveSrc, archiveDest)
	if err != nil {
		return errors.Wrap(err, "archive ketos container")
	}

	err = updateImageTag(ketosFolder, commitTag, digestNumber, int64(contentSize))
	if err != nil {
		return errors.Wrap(err, "update image tag")
	}

	return nil
}

func updateImageTag(ketosDir string, imageTag string, dige digest.Digest, size int64) error {

	manifestFile := filepath.Join(ketosDir, "tags", imageTag+".manifest")

	manifestContent, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return errors.Wrap(err, "read old manifest")
	}

	manifest := &manifestv2.Manifest{}
	err = json.Unmarshal(manifestContent, manifest)
	if err != nil {
		return errors.Wrap(err, "unmarshal old manifest")
	}
	log.Println(string(manifestContent))

	manifest.Layers = append(manifest.Layers, distribution.Descriptor{
		Digest:    dige,
		MediaType: "application/vnd.docker.image.rootfs.diff.tar.gzip",
		Size:      size,
	})

	manifestContent, err = json.Marshal(manifest)
	if err != nil {
		return errors.Wrap(err, "marshal new manifest")
	}
	log.Println(string(manifestContent))

	err = os.Remove(manifestFile)
	if err != nil {
		return errors.Wrap(err, "remove old manifest file")
	}
	err = ioutil.WriteFile(manifestFile, manifestContent, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "write down new manifest file")
	}

	return nil
}

func archiveKetosContainer(root string, dest string) error {
	return filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if path == root {
				return nil
			}

			if strings.Index(path, ".ketos") >= 0 {
				return filepath.SkipDir
			}

			err = os.Rename(path, filepath.Join(dest, filepath.Base(path)))
			if err != nil {
				return errors.Wrap(err, "rename file to archive")
			}

			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		},
	)
}

func gztarKetosContainer(root string, out io.Writer) error {

	gzipOut := gzip.NewWriter(out)
	defer gzipOut.Close()
	tarOut := tar.NewWriter(gzipOut)
	defer tarOut.Close()

	return filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if path == root {
				return nil
			}

			if strings.Index(path, ".ketos") >= 0 {
				return filepath.SkipDir
			}

			if info.IsDir() {
				return nil
			}

			tarHdr, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return errors.Wrap(err, "parsing tar header")
			}

			err = tarOut.WriteHeader(tarHdr)
			if err != nil {
				return errors.Wrap(err, "write tar file header")
			}

			file, err := os.Open(path)
			if err != nil {
				return errors.Wrap(err, "open layer file")
			}
			defer file.Close()

			_, err = io.Copy(tarOut, file)
			if err != nil {
				return errors.Wrap(err, "tar layer file")
			}

			return nil
		},
	)
}
