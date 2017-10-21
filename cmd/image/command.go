package image

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	defaultDir = ".ketos/tags"
)

var (
	Command = &cobra.Command{
		Use:   "image",
		Short: "show image in dir",
		RunE:  imageMain,
	}
)

func init() {
	flags := Command.Flags()
	flags.StringP("workDir", "d", "", "show images in work dir")
}

func imageMain(cmd *cobra.Command, args []string) error {
	dir, _ := cmd.Flags().GetString("workDir")
	if dir == "" {
		dir = defaultDir
	}

	err := showImages(dir)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
