package pull

import (
	"fmt"
	"strings"

	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "pull",
		Short: "pull [tag]",
		Args:  cobra.MaximumNArgs(1),

		RunE: pullMain,
	}
)

func init() {
}

func pullMain(cmd *cobra.Command, args []string) error {

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

	manifest, err := repo.GetManifest(tag)
	if err != nil {
		return err
	}

	// sync layers
	for _, layer := range manifest.FSLayers {
		layerPath := meta.LayerPath(layer.BlobSum)
		err = repo.GetLayer2Directory(layer.BlobSum, layerPath)
		if err != nil {
			return err
		}
	}

	// sync manifest
	return meta.PutManifest(tag, manifest)
}

func pullMain0(cmd *cobra.Command, args []string) error {
	name, tag, err := validRef(args[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fetch manifest, then get every layer
	err = pullV2("library/"+name, tag)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("%s download successfull\n", args[0])
	return nil
}

func validRef(ref string) (string, string, error) {
	idx := strings.LastIndex(ref, "/")
	if idx != -1 {
		ref = string([]byte(ref)[idx+1:])
	}
	repo := strings.Split(ref, ":")
	if len(repo) > 2 {
		return "", "", fmt.Errorf("image format error, should be \"registry/lib/name:tag\"")
	}

	if len(repo) == 2 {
		return repo[0], repo[1], nil
	}

	return repo[0], "", nil
}
