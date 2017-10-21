package rootpath

import (
	"encoding/json"
	manifestv2 "github.com/docker/distribution/manifest/schema2"
	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/ketos/metadata"
	"os"
	"path/filepath"
)

func expandOverlayPath(path string) string {

	ketosFolder, err := metadata.KetosFolder()
	if err != nil {
		panic(errors.Wrap(err, "seek ketos folder"))
	}

	manifestFileName := filepath.Join(ketosFolder, "tags", KetosChrootImgTag+".manifest")
	manifestFile, err := os.Open(manifestFileName)
	if err != nil {
		panic(errors.Wrap(err, "open manifest file"))
	}
	defer manifestFile.Close()

	manifest := &manifestv2.DeserializedManifest{}
	err = json.NewDecoder(manifestFile).Decode(manifest)
	if err != nil {
		panic(errors.Wrap(err, "decode manifest"))
	}

	for i := len(manifest.Layers) - 1; i >= 0; i-- {

		layer := manifest.Layers[i]
		digest := layer.Digest.Encoded()
		layerPath := filepath.Join(ketosFolder, "layers", digest)
		layerFilePath := filepath.Join(layerPath, path)

		_, err = os.Stat(layerFilePath)
		if err == nil {
			return layerFilePath
		}
		if !os.IsNotExist(err) {
			panic(errors.Wrap(err, "checking layer file"))
		}
	}

	return filepath.Join(filepath.Dir(ketosFolder), path)
}
