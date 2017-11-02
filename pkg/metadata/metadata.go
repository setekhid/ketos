package metadata

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	manifestV1 "github.com/docker/distribution/manifest/schema1"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Metadatas struct {
	folders MetaFolders
}

func NewMetadata(path string, image string) (*Metadatas, error) {

	path, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.Wrap(err, "calc absolute path of ketos folder")
	}

	data := &Metadatas{MetaFolders(path)}
	err = data.init(image)

	return data, err
}

func ConnMetadata(path string) (*Metadatas, error) {

	path, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.Wrap(err, "calc absolute path of ketos folder")
	}

	data := &Metadatas{MetaFolders(path)}
	inited, err := data.hasInited()
	if err != nil {
		return nil, err
	}
	if !inited {
		return nil, errors.New("ketos metadata folder hasn't been initialized")
	}

	return data, err
}

func (d *Metadatas) init(image string) error {

	inited, err := d.hasInited()
	if err != nil {
		return err
	}
	if inited {
		return nil
	}

	err = d.folders.InitFolders()
	if err != nil {
		return err
	}

	err = d.folders.InitConfig(image)
	if err != nil {
		return err
	}

	return nil
}

func (d *Metadatas) hasInited() (bool, error) {

	_, err := os.Stat(d.folders.Config())
	if err != nil && !os.IsNotExist(err) {
		return false, errors.Wrap(err, "check ketos config file")
	}

	return err == nil, nil
}

type KetosConfig struct {
	InitImageName string `yaml:"init_image_name"`

	Repository struct {
		Name     string `yaml:"name"`
		Registry string `yaml:"registry"`
	}
}

func (d *Metadatas) GetConfig() (*KetosConfig, error) {

	content, err := ioutil.ReadFile(d.folders.Config())
	if err != nil {
		return nil, errors.Wrap(err, "reading ketos config")
	}

	config := &KetosConfig{}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal ketos config")
	}

	return config, nil
}

func (d *Metadatas) ListTags() ([]string, error) {

	infos, err := ioutil.ReadDir(d.folders.Manifests())
	if err != nil {
		return nil, errors.Wrap(err, "reading manifest directory")
	}

	tags := []string{}
	for _, info := range infos {

		name := info.Name()
		if strings.HasSuffix(name, ".json") {
			tag := name[:len(name)-len(".json")]
			tags = append(tags, tag)
		}
	}

	return tags, nil
}

func (d *Metadatas) GetManifest(tag string) (*manifestV1.Manifest, error) {

	content, err := ioutil.ReadFile(d.folders.Manifest(tag))
	if err != nil {
		return nil, errors.Wrap(err, "reading manifest file")
	}

	manifest := &manifestV1.Manifest{}
	err = json.Unmarshal(content, manifest)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal manifest file")
	}

	return manifest, nil
}

func (d *Metadatas) PutManifest(
	tag string, manifest *manifestV1.Manifest) error {

	content, err := json.Marshal(manifest)
	if err != nil {
		return errors.Wrap(err, "marshal manifest json")
	}

	err = ioutil.WriteFile(d.folders.Manifest(tag), content, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "write down manifest")
	}

	return nil
}

func (d *Metadatas) LayerPath(digest digest.Digest) string {
	return d.folders.Layer(digest)
}

func (d *Metadatas) ContainerPath() string {
	return d.folders.Container()
}

func (d *Metadatas) MetaFolderPath() string {
	return string(d.folders)
}
