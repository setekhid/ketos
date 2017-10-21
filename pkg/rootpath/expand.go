package rootpath

import (
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

func ExpandPath(path string) string {

	// Simple combine
	if !KetosChrootWD {
		return KetosChrootRoot + string(filepath.Separator) + path
	}

	// combine with overlay

	return ""
}
