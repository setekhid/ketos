package push

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "push",
		Short: "push current image to registry",
		RunE:  pushMain,
	}
)

func pushMain(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		fmt.Println("image name:tag should be provide")
		return nil
	}

	return nil
}

func parseRef(ref string) {

}
