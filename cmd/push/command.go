package push

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	defaultRegistry = "127.0.0.1:5000"
	defaultDir      = ".ketos/tags"
	defaultlLayer   = ".ketos/layers"
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

	name, tag, err := parseRef(args[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = push(name, tag)
	if err != nil {
		return err
	}
	fmt.Printf("image %s push successful\n", args[0])

	return nil
}

func parseRef(ref string) (string, string, error) {
	repo := strings.Split(ref, ":")
	if len(repo) > 2 {
		return "", "", fmt.Errorf("%s format error", ref)
	}

	if len(repo) == 2 {
		return repo[0], repo[1], nil
	}

	return repo[0], "latest", nil
}
