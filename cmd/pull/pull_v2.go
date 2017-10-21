package pull

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/docker/distribution"
	"github.com/setekhid/ketos/client"
)

var ()

func pullV2(name, tag string) error {
	// fetch manifest
	hub, err := client.NewRegitry("", "", "")
	if err != nil {
		return err
	}

	manifest, err := hub.ManifestV2(name, tag)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("m.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
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

	config, err := hub.DownloadLayer(name, manifest.Config.Digest)
	if err != nil {
		return err
	}

	fmt.Printf("config %+v\n", config)
	cf, err := os.OpenFile(manifest.Config.Digest.Encoded()+".json", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer cf.Close()
	buf1 := bytes.Buffer{}
	buf1.ReadFrom(config)
	cf.Write(buf1.Bytes())

	// fetch layers
	layers := manifest.Layers
	for _, l := range layers {
		fmt.Printf("download layer %v\n", l.Digest.Encoded())

		contents, err := hub.DownloadLayer(name, l.Digest)
		if err != nil {
			return err
		}
		defer contents.Close()

		digest := l.Digest.Encoded()
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
