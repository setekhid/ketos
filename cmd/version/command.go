package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Commit  string
)

var (
	Command = &cobra.Command{
		Use:   "version",
		Short: "show xcb version",
		Args:  cobra.NoArgs,

		RunE: versionMain,
	}
)

func versionMain(cmd *cobra.Command, args []string) error {

	_, err := fmt.Printf("Version: %10s\nCommit: %10s\n", Version, Commit)
	return err
}
