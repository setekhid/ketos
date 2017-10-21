package image

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	manifestV2 "github.com/docker/distribution/manifest/schema2"
)

func showImages(dir string) error {
	subdir, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	fmt.Printf("%-20s%-20s%-20s\n", "IMAGE ID", "IMAGE NAME", "IMAGE TAG")
	for _, tag := range subdir {
		manifest := filepath.Join(dir, tag.Name(), "manifest.json")
		if _, err := os.Stat(manifest); err != nil && os.IsNotExist(err) {
			fmt.Printf("%s manifest not exist\n", manifest)
			continue
		}

		fd, err := os.OpenFile(manifest, os.O_RDONLY, 0644)
		if err != nil {
			fmt.Printf("open file %s %s\n", manifest, err)
			continue
		}
		defer fd.Close()
		mj := &manifestV2.DeserializedManifest{}
		err = json.NewDecoder(fd).Decode(mj)
		if err != nil {
			fmt.Errorf("unmarshal json %s", err)
		}

		id := mj.Config.Digest.String()
		fmt.Printf("%-20s%-20s%-20s\n", formatID(id), "IMAGE NAME", tag.Name())
	}

	return nil
}

func formatID(digest string) string {
	prefix := "sha256:"
	return string([]byte(digest)[len(prefix) : len(prefix)+12])
}
