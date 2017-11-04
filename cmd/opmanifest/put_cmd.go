package opmanifest

import (
	"encoding/json"
	"os"

	manifestV1 "github.com/docker/distribution/manifest/schema1"
	"github.com/pkg/errors"
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/spf13/cobra"
)

var (
	PutManifest = &cobra.Command{
		Use:   "put-manifest",
		Short: "put-manifest [image-name]:tag",
		Args:  cobra.ExactArgs(1),

		RunE: putManifest,
	}
)

func init() {
}

func putManifest(cmd *cobra.Command, args []string) error {

	image, err := FillImage(args[0])
	if err != nil {
		return err
	}

	repo, tag, err := registry.DockerImage(image).Connect()
	if err != nil {
		return err
	}

	manifest := &manifestV1.Manifest{}
	err = json.NewDecoder(os.Stdin).Decode(manifest)
	if err != nil {
		return errors.Wrap(err, "unmarshal manifest")
	}

	return repo.PutManifest(tag, manifest)
}
