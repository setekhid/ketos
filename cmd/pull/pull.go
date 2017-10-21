package pull

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/docker/distribution"
	"github.com/setekhid/ketos/client"
)

func pull(name, tag string) error {
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

	maniDig, err := hub.ManifestDigest(name, tag)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", manifest)
	fmt.Printf("%+v\n", maniDig)

	// fetch layers
	layers := manifest.Layers
	for _, l := range layers {
		fmt.Printf("layer %v\n", l.Digest.Encoded())

		contents, err := hub.DownloadLayer(name, l.Digest)
		if err != nil {
			return err
		}
		defer contents.Close()

		fd, err := os.OpenFile(l.Digest.Encoded()+".tar", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer fd.Close()
		buf := bytes.Buffer{}
		buf.ReadFrom(contents)
		tw := gzip.NewWriter(fd)
		tw.Write(buf.Bytes())
		//	fmt.Println(buf.String())
	}

	return nil
}
