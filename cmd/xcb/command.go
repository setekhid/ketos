package main

import (
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command {
		Use: "xcb",
		Aliases: []string{"xcbuild"},
		Short: "cross container builder",
		RunE: xcb_main,
	}
)

func init() {

	// TODO
}

func xcb_main(cmd *cobra.Command, args []string) error {

	// TODO
	return nil
}
