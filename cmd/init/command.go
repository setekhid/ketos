package initk

import (
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	Command = &cobra.Command{
		Use:   "init",
		Short: "init [--image|-I setekhid/scratch:latest] .",
		Args:  cobra.ExactArgs(1),

		RunE: initMain,
	}
)

func init() {

	flags := Command.Flags()
	flags.StringP("image", "I", "setekhid/scratch:latest",
		"initialize image")
}

func initMain(cmd *cobra.Command, args []string) error {

	initImageName, err := cmd.Flags().GetString("image")
	if err != nil {
		return err
	}
	workingDir := args[0]

	meta, err := metadata.NewMetadata(
		filepath.Join(workingDir, metadata.KetosMetaFolder),
		initImageName)
	if err != nil {
		return err
	}

	// TODO pull down init image
	_ = meta

	return nil
}
