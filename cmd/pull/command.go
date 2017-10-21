package pull

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "pull",
		Short: "pull down image from registry, only support pull from docker.io",
		RunE:  pullMain,
	}
)

func init() {

	// flags for pull command

}

func pullMain(cmd *cobra.Command, args []string) error {
	name, tag, err := validRef(args[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fetch manifest, then get every layer
	err = pullV1("library/"+name, tag)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("%s download successfull\n", args[0])
	return nil
}

func validRef(ref string) (string, string, error) {
	idx := strings.LastIndex(ref, "/")
	if idx != -1 {
		ref = string([]byte(ref)[idx+1:])
	}
	repo := strings.Split(ref, ":")
	if len(repo) > 2 {
		return "", "", fmt.Errorf("image format error, should be \"registry/lib/name:tag\"")
	}

	if len(repo) == 2 {
		return repo[0], repo[1], nil
	}

	return repo[0], "", nil
}
