package pull

import (
	"errors"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "pull",
		Short: "pull down image from registry",
		RunE:  pull_main,
	}
)

func init() {

	flags := Command.Flags()
	flags.StringP("demo-flag", "D", "demo", "demo flag")

	// TODO flags
}

func pull_main(cmd *cobra.Command, args []string) error {

	demo_value, err := cmd.Flags().GetString("demo-flag")
	if err != nil {
		return errors.New("didn't find demo-flag")
	}
	_ = demo_value

	// TODO
	return nil
}
