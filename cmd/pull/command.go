package pull

import (
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command {
		Use: "pull",
		Short: "pull down image from registry",
		RunE: pull_main,
	}
)

func init() {

	// TODO flags
}

func pull_main(cmd *cobra.Command, args []string) error {

	// TODO
	return nil
}
