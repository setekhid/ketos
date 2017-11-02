package metadata

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
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
