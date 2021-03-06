package rootpath

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

// OverlayFS presents an object to parse path to the right layer path
type OverlayFS struct {
	top    string
	lowers []string
}

// NewOverlayFS generate a new OverlayFS object
func NewOverlayFS(layers ...string) (*OverlayFS, error) {

	absLayers := []string{}
	for _, layer := range layers {

		abs, err := filepath.Abs(layer)
		if err != nil {
			return nil, errors.Wrap(err, "calc abs path")
		}

		absLayers = append(absLayers, abs)
	}

	return NewSimpleOverlayFS(absLayers...), nil
}

func NewSimpleOverlayFS(layers ...string) *OverlayFS {
	return &OverlayFS{
		top:    layers[len(layers)-1],
		lowers: layers[:len(layers)-1],
	}
}

// NewOverlayFSFromEnv generate OverlayFS instance from environment variable
// KETOS_ROOTPATH_LAYERS, default value is "/:/_ketos"
func NewOverlayFSFromEnv() (*OverlayFS, error) {

	layers := os.Getenv("KETOS_ROOTPATH_LAYERS")
	if len(layers) <= 0 {
		return NewDefaultOverlayFS(), nil
	}

	return NewOverlayFS(filepath.SplitList(layers)...)
}

func NewDefaultOverlayFS() *OverlayFS {
	return NewSimpleOverlayFS(
		string(filepath.Separator),
		filepath.Join(string(filepath.Separator), "_ketos"),
	)
}

// Expand expand the path to the right layer path
func (r *OverlayFS) ExpandPath(path string) (string, error) {

	path, err := r.cleanPath(path)
	if err != nil {
		return "", err
	}

	// check top layer
	topLayerPath := filepath.Join(r.top, path)
	if _, err := os.Stat(topLayerPath); err == nil || !os.IsNotExist(err) {
		return topLayerPath, err
	}

	// check lower layers
	for i := len(r.lowers) - 1; i >= 0; i-- {

		lowerLayer := r.lowers[i]
		lowerLayerPath := filepath.Join(lowerLayer, path)

		if _, err = os.Stat(lowerLayerPath); err == nil || !os.IsNotExist(err) {
			return lowerLayerPath, err
		}
	}

	return topLayerPath, nil
}

// IsTopLayer check if the file path is in top layer
func (r *OverlayFS) IsTopLayer(path string) bool {

	path = filepath.Clean(path)
	return filepath.HasPrefix(path, r.top)
}

// CopyForWriting copy the lower layer file to top layer for writting
func (r *OverlayFS) CopyForWriting(path string) (string, error) {

	path, err := r.cleanPath(path)
	if err != nil {
		return "", err
	}

	lowerFilePath, err := r.ExpandPath(path)
	if err != nil {
		return "", err
	}

	// exists on top layer
	if r.IsTopLayer(lowerFilePath) {
		return lowerFilePath, nil
	}

	// open top layer file
	topFilePath := filepath.Join(r.top, path)
	topFile, err := os.Create(topFilePath)
	if err != nil {
		return "", errors.Wrap(err, "create top layer file")
	}
	defer topFile.Close()

	// open lower layer file
	lowerFile, err := os.Open(lowerFilePath)
	if err != nil {
		return "", errors.Wrap(err, "open lower layer file")
	}
	defer lowerFile.Close()

	_, err = io.Copy(topFile, lowerFile)
	if err != nil {
		return "", errors.Wrap(err, "copy lower level file to container")
	}

	return topFilePath, nil
}

// ShrinkPath shrink the expanded path to the fake rootfs path
func (r *OverlayFS) ShrinkPath(expandedPath string) (string, error) {

	expandedPath, err := r.cleanPath(expandedPath)
	if err != nil {
		return "", err
	}

	if filepath.HasPrefix(expandedPath, r.top+string(filepath.Separator)) {
		return expandedPath[len(r.top):], nil
	}

	// check lower layers
	for i := len(r.lowers) - 1; i >= 0; i-- {

		lower := r.lowers[i]
		if filepath.HasPrefix(expandedPath, lower+string(filepath.Separator)) {
			return expandedPath[len(lower):], nil
		}
	}

	return "", errors.New("unknown expanded path")
}

func (r *OverlayFS) prepareOnTopLayer(path string) (string, error) {

	topPath := filepath.Join(r.top, path)
	topDir := filepath.Dir(topPath)

	// FIXME file permission should use the lower layer's
	err := os.MkdirAll(topDir, os.ModePerm)
	if err != nil {
		return "", errors.Wrap(err, "prepare top layer")
	}

	return topPath, nil
}

// WipeFile wipe a regular file on top layer
func (r *OverlayFS) WipeFile(path string) error {

	path, err := r.cleanPath(path)
	if err != nil {
		return err
	}

	topPath, err := r.prepareOnTopLayer(path)
	if err != nil {
		return err
	}

	topDir, fileName := filepath.Split(topPath)

	file, err := os.Create(filepath.Join(topDir, ".wh."+fileName))
	if err != nil {
		return errors.Wrap(err, "wipe file")
	}
	defer file.Close()

	return nil
}

// WipeFolder wipe a folder on top layer
func (r *OverlayFS) WipeFolder(path string) error {

	path, err := r.cleanPath(path)
	if err != nil {
		return err
	}

	topPath, err := r.prepareOnTopLayer(path)
	if err != nil {
		return err
	}

	topDir, folderName := filepath.Split(topPath)

	err = os.Mkdir(filepath.Join(topDir, ".wh."+folderName), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "wipe folder")
	}

	return nil
}

func (r *OverlayFS) cleanPath(path string) (string, error) {

	path = filepath.Clean(path)
	if !filepath.HasPrefix(path, "/") {
		return "", errors.New("can't find location")
	}
	return path, nil
}
