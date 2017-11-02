package oplayer

import (
	"os"

	"github.com/opencontainers/go-digest"
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/spf13/cobra"
)

var (
	HasLayer = &cobra.Command{
		Use:   "has-layer",
		Short: "has-layer [--repository|-r image-name] layer-digest",
		Args:  cobra.ExactArgs(1),

		RunE: hasLayer,
	}
)

func init() {

	flags := HasLayer.Flags()
	flags.StringP("repository", "r", "",
		"set the repo image name with or without tag")
}

func hasLayer(cmd *cobra.Command, args []string) error {

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

	repo, _, err := registry.DockerImage(repoName).Connect()
	if err != nil {
		return err
	}

	has, err := repo.HasLayer(digest)
	if err != nil {
		return err
	}

	if !has {
		os.Exit(1)
	}

	return nil
}
