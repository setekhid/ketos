package chroot

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var (
	Command = &cobra.Command{
		Use:     "chroot",
		Aliases: []string{"chr"},
		Short:   "change to image base as root fs",
		RunE:    chrootMain,
	}
)

func init() {

	flags := Command.Flags()
	flags.StringP("engine", "e", "ld.so",
		"specify the chroot implementation engine name")
	flags.StringP("image-tag", "t", "latest",
		"the image repository tag, when overlay")
}

func chrootMain(cmd *cobra.Command, args []string) error {

	engineName, err := cmd.Flags().GetString("engine")
	if err != nil {
		return errors.Wrap(err, "parsing flag --engine")
	}
	imageTag, err := cmd.Flags().GetString("image-tag")
	if err != nil {
		return errors.Wrap(err, "parsing flag --image-tag")
	}

	userCommand := args

	root, err := SeekKetosRoot("./")
	if err != nil {
		return errors.Wrap(err, "seeking ketos root")
	}

	engine, err := NewEngineByName(engineName)
	if err != nil {
		return err
	}
	engine.Run(root, imageTag, userCommand, os.Stdin, os.Stdout, os.Stderr)

	return nil
}
