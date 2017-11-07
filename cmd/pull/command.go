package pull

import (
	"os"

	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/setekhid/ketos/pkg/registry"
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
		packFile := meta.PackFile(layer.BlobSum)

		// cache pack
		err := func() error {

			pack, err := os.Create(packFile)
			if err != nil {
				return err
			}
			defer pack.Close()

			err = repo.GetLayer(layer.BlobSum, pack)
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			return err
		}

		// untar layer
		err = func() error {

			pack, err := os.Open(packFile)
			if err != nil {
				return err
			}
			defer pack.Close()

			err = registry.UntarLayerDirectory(pack, layerPath)
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			return err
		}
	}

	// sync manifest
	return meta.PutManifest(tag, manifest)
}
