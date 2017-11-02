package oplayer

import (
	"github.com/opencontainers/go-digest"
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/spf13/cobra"
)

var (
	PullLayer = &cobra.Command{
		Use:   "pull-layer",
		Short: "pull-layer [--repository|-r image-name] layer-digest directory",
		Args:  cobra.ExactArgs(2),

		RunE: pullLayer,
	}
)

func init() {

	flags := PullLayer.Flags()
	flags.StringP("repository", "r", "",
		"repository name with or without tag")
}

func pullLayer(cmd *cobra.Command, args []string) error {

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

	digest, err := digest.Parse(args[0])
	if err != nil {
		return err
	}

	layerDirectory := args[1]

	repo, _, err := registry.DockerImage(repoName).Connect()
	if err != nil {
		return err
	}

	return repo.GetLayer2Directory(digest, layerDirectory)
}
