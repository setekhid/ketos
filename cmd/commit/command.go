package commit

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	manifestV1 "github.com/docker/distribution/manifest/schema1"
	imageV1 "github.com/docker/docker/image"
	"github.com/opencontainers/go-digest"
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "commit",
		Short: "commit [--tag|-t tag-name]",
		Args:  cobra.NoArgs,

		RunE: commitMain,
	}
)

func init() {

	flags := Command.Flags()
	flags.StringP("tag", "t", "", "commit to a tag")
}

type blackHole struct{}

func (_ blackHole) Write(p []byte) (int, error) { return len(p), nil }
func (_ blackHole) Read(p []byte) (int, error)  { return 0, nil }
func (_ blackHole) Close() error                { return nil }

func commitMain(cmd *cobra.Command, args []string) error {

	tag, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	meta, err := metadata.CurrentMetadatas()
	if err != nil {
		return err
	}

	containerPath := meta.ContainerPath()
	metaFolderPath := meta.MetaFolderPath()

	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	tmpPack := filepath.Join(tmpDir, "container-pack.tar.gz")
	digest := digest.Digest("unknow")

	// tar container pack
	err = func() error {
		containerPack, err := os.Create(tmpPack)
		if err != nil {
			return err
		}
		defer containerPack.Close()

		digest, err = registry.TarLayerDirectory(containerPack,
			containerPath, metaFolderPath)

		return err
	}()
	if err != nil {
		return err
	}

	// save container pack
	packFile := meta.PackFile(digest)
	err = os.Rename(tmpPack, packFile)
	if err != nil {
		return err
	}
	pack, err := os.Open(packFile)
	if err != nil {
		return err
	}
	defer pack.Close()

	layerPath := meta.LayerPath(digest)
	err = registry.UntarLayerDirectory(pack, layerPath)
	if err != nil {
		return err
	}

	// update manifest
	manifest, err := meta.GetManifest(tag)
	if err != nil {
		return err
	}

	manifest.FSLayers = append(
		[]manifestV1.FSLayer{
			{BlobSum: digest},
		},
		manifest.FSLayers...,
	)

	shit := manifest.History[0]
	imageShit := &imageV1.V1Image{}
	err = json.Unmarshal([]byte(shit.V1Compatibility), imageShit)
	if err != nil {
		return err
	}

	imageJson, err := json.Marshal(&imageV1.V1Image{
		ID:      digest.Hex(),
		Parent:  imageShit.ID,
		Created: time.Now(),
	})
	if err != nil {
		return err
	}
	manifest.History = append(
		[]manifestV1.History{
			{V1Compatibility: string(imageJson)},
		},
		manifest.History...,
	)

	return meta.PutManifest(tag, manifest)
}
