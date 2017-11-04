package oplayer

import (
	"github.com/setekhid/ketos/pkg/metadata"
)

func DefaultImageRepo() (string, error) {

	folder, err := metadata.KetosFolder()
	if err != nil {
		return "", err
	}

	meta, err := metadata.ConnMetadata(folder)
	if err != nil {
		return "", err
	}

	conf, err := meta.GetConfig()
	if err != nil {
		return "", err
	}

	return conf.Repository.Registry + "/" + conf.Repository.Name, nil
}
