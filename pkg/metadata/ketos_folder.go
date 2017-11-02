package metadata

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/registry"
)

const KetosMetaFolder = ".ketos"

// SeekKetosFolder seek .ketos from path to root
func SeekKetosFolder(path string) (string, error) {

	path, err := filepath.Abs(path)
	if err != nil {
		return "", errors.Wrap(err, "abs path")
	}

	for {

		mayKetos := filepath.Join(path, KetosMetaFolder)
		_, err := os.Stat(mayKetos)
		if err == nil {
			return mayKetos, nil
		}

		if !os.IsNotExist(err) {
			return "", errors.Wrap(err, "open ketos metadata")
		}

		parent := filepath.Dir(path)
		if path == parent {
			return "", errors.New("didn't find metadata")
		}
		path = parent
	}
}

// KetosFolder seek .ketos from current working directory to root
func KetosFolder() (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "get current working directory")
	}

	return SeekKetosFolder(wd)
}

// MetaFolders represents ketos metadata folders
type MetaFolders string

// InitFS initialize file system
func (m MetaFolders) InitFolders() error {

	err := os.MkdirAll(filepath.Dir(m.Config()), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "mkdir directory of config")
	}

	err = os.MkdirAll(m.Container(), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "mkdir container directory")
	}

	err = os.MkdirAll(m.Layers(), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "mkdir layers")
	}

	err = os.MkdirAll(m.Manifests(), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "mkdir manifests")
	}

	return nil
}

// InitConfig initialize config file for ketos repo
func (m MetaFolders) InitConfig(image string) error {

	config := &KetosConfig{}
	config.InitImageName = image

	config.Repository.Registry, config.Repository.Name, _ =
		registry.DockerImage(image).Split()

	content, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "marshal ketos config")
	}

	err = ioutil.WriteFile(m.Config(), content, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "write down config file")
	}

	return nil
}

// Config get the config file path
func (m MetaFolders) Config() string {
	return filepath.Join(string(m), "config.yaml")
}

// Container return the top working layer
func (m MetaFolders) Container() string {
	return filepath.Dir(string(m))
}

// Layers return the layers folder path
func (m MetaFolders) Layers() string {
	return filepath.Join(string(m), "layers")
}

// Layer return the specified layer directory
func (m MetaFolders) Layer(digest digest.Digest) string {
	return filepath.Join(m.Layers(), digest.Hex())
}

// MetaLayer return the json file of config layer
func (m MetaFolders) MetaLayer(digest digest.Digest) string {
	return filepath.Join(m.Layers(), digest.Hex()+".json")
}

// Manifests return the manifests folder path
func (m MetaFolders) Manifests() string {
	return filepath.Join(string(m), "manifests")
}

// Manifest get the specific manifest
func (m MetaFolders) Manifest(tag string) string {
	return filepath.Join(m.Manifests(), tag+".json")
}
