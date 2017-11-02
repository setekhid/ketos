package push

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/spf13/cobra"
)

var (
	defaultRegistry = "127.0.0.1:5000"
	defaultTagDir   = ""
	defaultlLayer   = ""
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

	ketosFolder, err := metadata.KetosFolder()
	if err != nil {
		fmt.Println(err)
		return err
	}

	name, tag, err := parseRef(args[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	defaultTagDir = filepath.Join(ketosFolder, "tags")
	defaultlLayer = filepath.Join(ketosFolder, "layers")

	//name = filepath.Join(defaultRegistry, name)
	err = push(name, tag)
	if err != nil {
		return err
	}
	fmt.Printf("image %s push successful\n", args[0])

	return nil
}

func parseRef(ref string) (string, string, error) {

	sepInd := strings.LastIndex(ref, ":")
	if sepInd < 0 {
		return "", "", fmt.Errorf("%s format error", ref)
	}

	return ref[:sepInd], ref[sepInd+1:], nil
}
