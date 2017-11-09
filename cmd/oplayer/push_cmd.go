package oplayer

import (
	"os"

	"github.com/setekhid/ketos/pkg/registry"
	"github.com/spf13/cobra"
)

var (
	PushLayer = &cobra.Command{
		Use:   "push-layer",
		Short: "push-layer [--repository|-r image-name] directory ignores...",
		Args:  cobra.MinimumNArgs(1),

		RunE: pushLayer,
	}
)

func init() {

	flags := PushLayer.Flags()
	flags.StringP("repository", "r", "",
		"repository name with or without tag")
}

func pushLayer(cmd *cobra.Command, args []string) error {

	repoName, err := cmd.Flags().GetString("repository")
	if err != nil {
		return err
	}

	if len(repoName) <= 0 {
		repoName, err = DefaultImageRepo()
		if err != nil {
			return err
		}
	}

	layerDirectory := args[1]
	ignoreDirectories := args[1:]

	repo, _, err := registry.DockerImage(repoName).Connect()
	if err != nil {
		return err
	}

	digest, err := repo.PostLayerDirectory(layerDirectory, ignoreDirectories...)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write([]byte(digest.String()))
	return err
}
