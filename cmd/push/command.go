package push

import (
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command {
		Use: "push",
		Short: "push current image to registry",
		RunE: push_main,
	}
)

func init() {

	// TODO flags
}

func push_main(cmd *cobra.Command, args []string) error {

	// TODO
	return nil
}
