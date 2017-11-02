package catmanifest

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "cat-manifest",
		Short: "cat-manifest [image-name]:tag",
		Args:  cobra.ExactArgs(1),

		RunE: catManifest,
	}
)

func init() {
}

func catManifest(cmd *cobra.Command, args []string) error {

	image := args[0]
	if strings.HasPrefix(image, ":") {

		tag := image

		folder, err := metadata.KetosFolder()
		if err != nil {
			return err
		}

		meta, err := metadata.ConnMetadata(folder)
		if err != nil {
			return err
		}

		conf, err := meta.GetConfig()
		if err != nil {
			return err
		}

		image = conf.Repository.Registry + "/" + conf.Repository.Name + tag
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
