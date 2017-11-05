package push

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/spf13/cobra"
)

var (
	defaultRegistry = "127.0.0.1:5000"
	defaultTagDir   = ""
	defaultlLayer   = ""
)

var (
	Command = &cobra.Command{
		Use:   "push",
		Short: "push [tag]",
		Args:  cobra.MaximumNArgs(1),

		RunE: pushMain,
	}
)

func pushMain(cmd *cobra.Command, args []string) error {

	tag := "latest"
	if len(args) > 0 {
		tag = args[0]
	}

	// metadata folder
	meta, err := metadata.CurrentMetadatas()
	if err != nil {
		return err
	}

	// repository
	repo, err := meta.ConnectRegistry()
	if err != nil {
		return err
	}

	manifest, err := meta.GetManifest(tag)
	if err != nil {
		return err
	}

	// sync layers
	for _, layer := range manifest.FSLayers {

		layerPath := meta.LayerPath(layer.BlobSum)
		has, err := repo.HasLayer(layer.BlobSum)
		if err != nil {
			return err
		}

		if !has {

			digest, err := repo.PostLayerDirectory(layerPath)
			if err != nil {
				return err
			}

			if digest.String() != layer.BlobSum.String() {
				return errors.New("post a layer with wrong digest")
			}
		}
	}

	// sync manifest
	return repo.PutManifest(tag, manifest)
}

func pushMain0(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		fmt.Println("image name:tag should be provide")
		return nil
	}

	ketosFolder, err := metadata.KetosFolder()
	if err != nil {
		fmt.Println(err)
		return err
	}

	name, tag, err := parseRef(args[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	defaultTagDir = filepath.Join(ketosFolder, "tags")
	defaultlLayer = filepath.Join(ketosFolder, "layers")

	//name = filepath.Join(defaultRegistry, name)
	err = push(name, tag)
	if err != nil {
		return err
	}
	fmt.Printf("image %s push successful\n", args[0])

	return nil
}

func parseRef(ref string) (string, string, error) {

	sepInd := strings.LastIndex(ref, ":")
	if sepInd < 0 {
		return "", "", fmt.Errorf("%s format error", ref)
	}

	return ref[:sepInd], ref[sepInd+1:], nil
}
