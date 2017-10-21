package chroot

import (
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command {
		Use: "chroot",
		Aliases: []string{"chr"},
		Short: "change to image base as root fs",
		RunE: chroot_main,
	}
)

func init() {

	// TODO flags
}

func chroot_main(cmd *cobra.Command, args []string) error {

	// TODO
	return nil
}
