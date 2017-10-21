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

func ExpandPath(path string, ro bool) string {

	// check the top layer contains
	topLayer := KetosChrootRoot + string(filepath.Separator) + path
	_, err := os.Stat(topLayer)
	if err == nil || !os.IsNotExist(err) {
		return topLayer
	}

	// seeking lower layer path
	lowerLayer := path
	if !KetosChrootWD {
		lowerLayer = expandRootfsPath(path)
	} else {
		lowerLayer = expandOverlayPath(path)
	}

	// readonly
	if ro {
		return lowerLayer
	}

	// may write

	// write to top layer directory, or doesn't exist
	info, err := os.Stat(lowerLayer)
	if err != nil || info.IsDir() {
		return topLayer
	}

	// copy regular file
	lowerLayerFile, err := os.Open(lowerLayer)
	if err != nil {
		return topLayer
	}
	defer lowerLayerFile.Close()
	topLayerFile, err := os.Create(topLayer)
	if err != nil {
		return topLayer
	}
	defer topLayerFile.Close()

	io.Copy(topLayerFile, lowerLayerFile)
	return topLayer
}

func expandRootfsPath(path string) string {
	return path
}
