package pull

import (
	"fmt"
	"strings"

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

	// flags for pull command

}

func pull_main(cmd *cobra.Command, args []string) error {
	name, tag, err := validRef(args[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fetch manifest, then get every layer
	return nil
}

func validRef(ref string) (string, string, error) {
	repo := strings.Split(ref, ":")
	if len(repo) > 2 {
		return "", "", fmt.Errorf("image format error, should be \"name:tag\"")
	}

	if len(repo) == 2 {
		return repo[0], repo[1], nil
	}

	return repo[0], "", nil
}
