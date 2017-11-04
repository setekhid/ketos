package opmanifest

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/spf13/cobra"
)

var (
	CatManifest = &cobra.Command{
		Use:   "cat-manifest",
		Short: "cat-manifest [image-name]:tag",
		Args:  cobra.ExactArgs(1),

		RunE: catManifest,
	}
)

func init() {
}

func catManifest(cmd *cobra.Command, args []string) error {

	image, err := FillImage(args[0])
	if err != nil {
		return err
	}

	repo, tag, err := registry.DockerImage(image).Connect()
	if err != nil {
		return err
	}

	manifest, err := repo.GetManifest(tag)
	if err != nil {
		return err
	}

	content, err := json.Marshal(manifest)
	if err != nil {
		return errors.Wrap(err, "marshal manifest")
	}

	_, err = os.Stdout.Write(content)
	if err != nil {
		return errors.Wrap(err, "output manifest to stdout")
	}

	return nil
}
