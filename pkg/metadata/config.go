package metadata

type KetosConfig struct {
	InitImageName string `yaml:"init_image_name"`

	Repository struct {
		Name     string `yaml:"name"`
		Registry string `yaml:"registry"`
	}
}
