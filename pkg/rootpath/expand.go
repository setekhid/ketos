package rootpath

import (
	"io"
	"os"
	"path/filepath"
)

/*
 *  Working Directory
 *  |
 *  +- .ketos
 *  |  |
 *  |  +- layers (each layers, folder may named with digest number)
 *  |  |
 *  |  +- tags (each tag manifest file)
 *  |
 *  +- asset_file.txt
 */

// ExpandPath expand the path to fake rootfs
func ExpandPath(path string, ro bool) (string, error) {

	// check the top layer contains
	topLayer := KetosChrootDir + string(filepath.Separator) + path
	_, err := os.Stat(topLayer)
	if err == nil || !os.IsNotExist(err) {
		return topLayer, err
	}

	// seeking lower layer path
	lowerLayer := path
	if KetosChrootToImg {
		lowerLayer, err = expandOverlayPath(path)
		if err != nil {
			return "", err
		}
		if len(lowerLayer) <= 0 {
			return topLayer, nil
		}
	}

	// readonly
	if ro {
		return lowerLayer, nil
	}

	// write to top layer directory, or doesn't exist
	info, err := os.Stat(lowerLayer)
	if err != nil {

		if os.IsNotExist(err) {
			return topLayer, nil
		}

		return "", err
	}
	if info.IsDir() {
		return topLayer, os.MkdirAll(topLayer, os.ModePerm)
	}

	// copy regular file
	lowerLayerFile, err := os.Open(lowerLayer)
	if err != nil {
		return "", err
	}
	defer lowerLayerFile.Close()

	topLayerFile, err := os.Create(topLayer)
	if err != nil {
		return "", err
	}
	defer topLayerFile.Close()

	_, err = io.Copy(topLayerFile, lowerLayerFile)
	return topLayer, err
}
