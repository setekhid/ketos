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

		_, err := os.Stat(filepath.Join(path, ".ketos"))
		if err == nil {
			return filepath.Join(path, ".ketos"), nil
		}

		if !os.IsNotExist(err) {
			return "", errors.Wrap(err, "open ketos metadata")
		}

		if path == "/" {
			return "", errors.New("didn't find metadata")
		}

		path = filepath.Dir(path)
	}
}

func KetosFolder() (string, error) {
	return SeekKetosFolder("./")
}
