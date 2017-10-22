package push

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	manifestV2 "github.com/docker/distribution/manifest/schema2"
	"github.com/setekhid/ketos/client"
)

func push(name, tag string) error {
	hub, err := client.NewRegitry("http://127.0.0.1/", "", "")
	if err != nil {
		return err
	}

	manifest, err := marshalManifest(defaultTagDir, tag)
	if err != nil {
		return err
	}

	for _, l := range manifest.Layers {
		digest := l.Digest
		exist, err := hub.HasLayer(name, digest)
		if err != nil {
			return err
		}
		if exist {
			continue
		}

		layer := l.Digest.Encoded()
		hasL := getLayer(layer)
		if !hasL {
			fmt.Printf("layer %s not exist in %s\n", layer, defaultlLayer)
			return nil
		}

		gz := filepath.Join(defaultlLayer, layer+".tar.gz")
		contents, err := os.Open(gz)
		if err != nil {
			return err
		}

		err = hub.UploadLayer(name, digest, contents)
		if err != nil {
			return err
		}

		fmt.Printf("layer %s push successful\n", layer)
	}

	return nil
}

func getLayer(layer string) bool {
	subdir, err := ioutil.ReadDir(defaultlLayer)
	if err != nil {
		fmt.Println(err)
		return false
	}

	for _, d := range subdir {
		if d.Name() == layer {
			return true
		}
	}

	return false
}

func marshalManifest(tagDir, tag string) (*manifestV2.DeserializedManifest, error) {
	path := filepath.Join(tagDir, tag+".manifest")
	fm, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	manifest := &manifestV2.DeserializedManifest{}

	err = json.NewDecoder(fm).Decode(manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}
