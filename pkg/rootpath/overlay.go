package rootpath

import (
	"encoding/json"
	manifestv2 "github.com/docker/distribution/manifest/schema2"
	"github.com/setekhid/ketos/pkg/metadata"
	"os"
	"path/filepath"
)

func expandOverlayPath(path string) (string, error) {

	ketosFolder, err := metadata.KetosFolder()
	if err != nil {
		return "", err
	}

	manifestFileName := filepath.Join(ketosFolder,
		"tags", KetosChrootImgTag+".manifest")
	manifestFile, err := os.Open(manifestFileName)
	if err != nil {
		return "", err
	}
	defer manifestFile.Close()

	manifest := &manifestv2.Manifest{}
	err = json.NewDecoder(manifestFile).Decode(manifest)
	if err != nil {
		return "", err
	}

	for i := len(manifest.Layers) - 1; i >= 0; i-- {

		layer := manifest.Layers[i]
		digest := layer.Digest.Hex()
		layerPath := filepath.Join(ketosFolder, "layers", digest)
		layerFilePath := filepath.Join(layerPath, path)

		_, err = os.Stat(layerFilePath)
		if err == nil {
			return layerFilePath, nil
		}
		if !os.IsNotExist(err) {
			return "", err
		}
	}

	return "", nil
}
