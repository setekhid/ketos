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
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "show xcb version",
		RunE:    versionMain,
	}
)

func versionMain(cmd *cobra.Command, args []string) error {
	fmt.Printf("Version: %10s\nCommit: %10s\n", Version, Commit)

	return nil
}
