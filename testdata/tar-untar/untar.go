package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./taste.tar")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tarr := tar.NewReader(file)

	for {

		hdr, err := tarr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("tared file: %v\n", hdr.Name)
		_, err = io.Copy(os.Stdout, tarr)
		if err != nil {
			panic(err)
		}

		fmt.Println()
	}
}
