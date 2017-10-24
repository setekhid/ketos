package metadata

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func SeekKetosFolder(path string) (string, error) {

	path, err := filepath.Abs(path)
	if err != nil {
		return "", errors.Wrap(err, "abs path")
	}

	for {

		mayKetos := filepath.Join(path, ".ketos")
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

func KetosFolder() (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "get current working directory")
	}

	return SeekKetosFolder(wd)
}
