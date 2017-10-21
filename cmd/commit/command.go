package commit

import (
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command {
		Use: "commit",
		Short: "commit current working tree to a layer of image",
		RunE: commit_main,
	}
)

func init() {

	// TODO flags
}

func commit_main(cmd *cobra.Command, args []string) error {

	// TODO
	return nil
}
