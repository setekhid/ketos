package push

import (
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/spf13/cobra"
	"os"
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

		packFile := meta.PackFile(layer.BlobSum)
		has, err := repo.HasLayer(layer.BlobSum)
		if err != nil {
			return err
		}

		if !has {

			err = func() error {
				pack, err := os.Open(packFile)
				if err != nil {
					return err
				}
				defer pack.Close()

				return repo.PutLayer(layer.BlobSum, pack)
			}()
			if err != nil {
				return err
			}
		}
	}

	// sync manifest
	return repo.PutManifest(tag, manifest)
}
