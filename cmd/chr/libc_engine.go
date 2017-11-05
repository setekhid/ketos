package main

import (
	"github.com/setekhid/ketos/pkg/metadata"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func newLibcChrootExecutor(env []string) (ChrootExecutor, error) {

	libcHookPath := os.Getenv("KETOS_LIBC_HOOKER")
	if len(libcHookPath) <= 0 {
		libcHookPath = "/usr/local/lib/libketos-hookroot.so"
	}

	executor := func(repoPath, tagName string, userCommand []string) error {

		meta, err := metadata.GetMetadatas(repoPath)
		if err != nil {
			return err
		}
		manifest, err := meta.GetManifest(tagName)
		if err != nil {
			return err
		}

		rootfsLayers := []string{}
		for _, layer := range manifest.FSLayers {
			layerPath := meta.LayerPath(layer.BlobSum)
			rootfsLayers = append(rootfsLayers, layerPath)
		}
		rootfsLayers = append(rootfsLayers, meta.ContainerPath())

		exe := exec.Command(userCommand[0], userCommand[1:]...)
		exe.Env = append(exe.Env,
			"LD_PRELOAD="+libcHookPath,
			"KETOS_ROOTPATH_LAYERS="+
				strings.Join(rootfsLayers, string(filepath.ListSeparator)))
		exe.Stdin = os.Stdin
		exe.Stdout = os.Stdout
		exe.Stderr = os.Stderr

		return exe.Run()
	}

	return ExecutorFunc(executor), nil
}

func init() { AddExecutor("libc", newLibcChrootExecutor) }
