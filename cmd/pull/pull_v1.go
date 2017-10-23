package pull

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/docker/distribution"
	client "github.com/setekhid/ketos/pkg/registry"
)

func pullV1(name, tag string) error {
	// fetch manifest
	hub, err := client.NewRegitry("", "", "")
	if err != nil {
		return err
	}

	manifest, err := hub.Manifest(name, tag)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("m1.json", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(manifest)
	if err != nil {
		fmt.Println("encode json ", err)
		return err
	}

	fmt.Printf("%+v\n", manifest)

	// fetch layers
	layers := manifest.FSLayers
	for _, l := range layers {
		fmt.Printf("download layer %v\n", l.BlobSum.Encoded())

		contents, err := hub.DownloadLayer(name, l.BlobSum)
		if err != nil {
			return err
		}
		defer contents.Close()

		digest := l.BlobSum.Encoded()
		tmpDir, err := ioutil.TempDir(".", digest)
		if err != nil {
			return err
		}
		gz := filepath.Join(tmpDir, digest)
		fd, err := os.OpenFile(gz+".tar.gz", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer fd.Close()

		buf := bytes.Buffer{}
		buf.ReadFrom(contents)
		_, err = fd.Write(buf.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}
