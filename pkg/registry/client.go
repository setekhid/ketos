package registry

import (
	"io"
	"os"
	"strings"

	manifestV1 "github.com/docker/distribution/manifest/schema1"
	"github.com/docker/libtrust"
	hreg "github.com/heroku/docker-registry-client/registry"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
)

var (
	DefaultRegistry   = "registry-1.docker.io"
	QuayRegistry      = "quay.io"
	GoogleRegistry    = "gcr.io"
	LocalhostRegistry = "localhost:5000"

	// Deprecated
	defaultRegistry = "https://registry-1.docker.io/"
	defaultName     = ""
	defaultPasswd   = ""
)

// Deprecated
func NewRegitry(registry, name, passwd string) (*hreg.Registry, error) {
	if registry == "" {
		registry = defaultRegistry
	}

	if name == "" {
		name = defaultName
	}

	if passwd == "" {
		passwd = defaultPasswd
	}

	hub, err := hreg.New(registry, name, passwd)
	if err != nil {
		return nil, err
	}

	return hub, err
}

// ImageName represents a docker image name
type DockerImage string

// IsFromHub check if this image should be from docker hub
func (img DockerImage) IsFromHub() bool {
	return strings.Count(string(img), "/") < 2
}

// Split splits the image whole name to registry url, repository name and image
// tag
func (img DockerImage) Split() (string, string, string) {

	registry, img := img.SplitRegistry()
	img, tag := img.SplitTag()
	repository, _ := img.ToRepository()

	return registry, repository, tag
}

// Connect connects to repository, return repo and tag from DockerImage
func (img DockerImage) Connect() (*Repository, string, error) {

	registry, repository, tag := img.Split()
	repo, err := ConnectRepository(registry, repository)
	return repo, tag, err
}

// ToRepository check if image name doesn't contain tag field, caller's
// reposibility to make sure registry field already been splitted
func (img DockerImage) ToRepository() (string, bool) {
	return string(img), strings.Count(string(img), ":") <= 0
}

// SplitRegistry return the registry's url which holding this image, and the
// rest of image name
func (img DockerImage) SplitRegistry() (string, DockerImage) {

	if img.IsFromHub() {

		registry, imageName := DefaultRegistry, string(img)
		if strings.Count(imageName, "/") < 1 {
			imageName = "library/" + imageName
		}

		return registry, DockerImage(imageName)
	}

	name := string(img)
	ind := strings.Index(name, "/")

	return name[:ind], DockerImage(name[ind+1:])
}

// SplitTag return the rest of image name and image tag
func (img DockerImage) SplitTag() (DockerImage, string) {

	name := string(img)

	ind := strings.LastIndex(name, ":")
	if ind < 0 {
		return img, "latest"
	}

	return DockerImage(name[:ind]), name[ind+1:]
}

// Repository represents a docker registry repository
type Repository struct {
	conn *hreg.Registry
	repo string
}

// ConnectRepository get a connection to a registry's repository
func ConnectRepository(registry, repository string) (*Repository, error) {

	conf := GetRegistryConfig(registry)

	proto := "https"
	if conf.Insecure {
		proto = "http"
	}
	url := proto + "://" + registry

	conn, err := NewRegitry(url, "", "")
	if err != nil {
		return nil, err
	}

	return &Repository{conn, repository}, nil
}

// GetManifest gets manifest of a tag of image
func (r *Repository) GetManifest(tag string) (*manifestV1.Manifest, error) {

	manifest, err := r.conn.Manifest(r.repo, tag)
	if err != nil {
		return nil, err
	}

	return &manifest.Manifest, nil
}

// PutManifest puts manifest to a tag of image
func (r *Repository) PutManifest(
	tag string, manifest *manifestV1.Manifest) error {

	key, err := libtrust.GenerateRSA2048PrivateKey()
	if err != nil {
		return errors.Wrap(err, "generate rsa2048 private key")
	}

	signed, err := manifestV1.Sign(manifest, key)
	if err != nil {
		return errors.Wrap(err, "sign manifest")
	}

	return r.conn.PutManifest(r.repo, tag, signed)
}

// GetLayer gets the layer of image
func (r *Repository) GetLayer(digest digest.Digest, content io.Writer) error {

	rc, err := r.conn.DownloadLayer(r.repo, digest)
	if err != nil {
		return errors.Wrap(err, "download layer")
	}
	defer rc.Close()

	_, err = io.Copy(content, rc)
	if err != nil {
		return errors.Wrap(err, "output layer content")
	}

	return nil
}

// GetLayer2Directory fetching the layer and untar it
func (r *Repository) GetLayer2Directory(
	digest digest.Digest, root string) error {

	pipeR, pipeW, err := os.Pipe()
	if err != nil {
		return errors.Wrap(err, "open pipe for caching layer")
	}
	defer pipeR.Close()

	err = func() error {
		defer pipeW.Close()
		return r.GetLayer(digest, pipeW)
	}()
	if err != nil {
		return err
	}

	return UntarLayerDirectory(pipeR, root)
}

// PutLayer puts the layer of image
func (r *Repository) PutLayer(digest digest.Digest, content io.Reader) error {

	return r.conn.UploadLayer(r.repo, digest, content)
}

// PutLayer4Directory put layer from directory
func (r *Repository) PostLayerDirectory(
	root string, ignores ...string) (digest.Digest, error) {

	pipeR, pipeW, err := os.Pipe()
	if err != nil {
		return "", errors.Wrap(err, "open pipe for caching layer")
	}
	defer pipeR.Close()

	digest, err := func() (digest.Digest, error) {
		defer pipeW.Close()
		return TarLayerDirectory(pipeW, root, ignores...)
	}()
	if err != nil {
		return "", err
	}

	return digest, r.PutLayer(digest, pipeR)
}
