package tag

import (
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "tag",
		Short: "tag src-tag dest-tag",
		Args:  cobra.ExactArgs(2),

		RunE: tagMain,
	}
)

func init() {
}

func tagMain(cmd *cobra.Command, args []string) error {

	srcTag, destTag := args[0], args[1]

	meta, err := metadata.CurrentMetadatas()
	if err != nil {
		return err
	}

	manifest, err := meta.GetManifest(srcTag)
	manifest.Tag = destTag
	return meta.PutManifest(destTag, manifest)
}
