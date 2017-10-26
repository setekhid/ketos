package main

import (
	"archive/tar"
	"os"
)

func main() {

	file, err := os.OpenFile("./taste.tar",
		os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tarw := tar.NewWriter(file)
	defer tarw.Close()

	asset1_content := []byte("this is asset1")
	asset1 := tar.Header{
		Name: "assets/asset1.txt",
		Size: int64(len(asset1_content)),
		Mode: int64(os.ModePerm),
	}

	err = tarw.WriteHeader(&asset1)
	if err != nil {
		panic(err)
	}
	_, err = tarw.Write(asset1_content)
	if err != nil {
		panic(err)
	}

	asset2_content := []byte("this is an etc file")
	asset2 := tar.Header{
		Name: "etc/config.yaml",
		Size: int64(len(asset2_content)),
		Mode: int64(os.ModePerm),
	}

	err = tarw.WriteHeader(&asset2)
	if err != nil {
		panic(err)
	}
	_, err = tarw.Write(asset2_content)
	if err != nil {
		panic(err)
	}

	asset_dir := tar.Header{
		Name:     "abc/bbc",
		Typeflag: tar.TypeDir,
		Mode:     int64(os.ModePerm),
	}
	err = tarw.WriteHeader(&asset_dir)
	if err != nil {
		panic(err)
	}
}
