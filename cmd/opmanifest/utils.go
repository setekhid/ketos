package opmanifest

import (
	"strings"

	"github.com/setekhid/ketos/pkg/metadata"
)

func FillImage(mayOnlyTag string) (string, error) {

	if strings.HasPrefix(mayOnlyTag, ":") {

		tag := mayOnlyTag

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

		return conf.Repository.Registry + "/" + conf.Repository.Name + tag, nil
	}

	return mayOnlyTag, nil
}
