package registry

var (
	registryConfigs       = map[string]RegistryConfig{}
	defaultRegistryConfig = RegistryConfig{
		Insecure: true,
	}
)

// RegistryConfig represents config information for a docker registry
type RegistryConfig struct {
	Name     string
	Insecure bool
}

// ConfigRegistry config a docker registry
func ConfigRegistry(name string, insecure bool) {
	registryConfigs[name] = RegistryConfig{
		Name:     name,
		Insecure: insecure,
	}
}

// GetRegistryConfig get the config of a docker registry, and return default if
// doesn't exist
func GetRegistryConfig(name string) RegistryConfig {

	config, ok := registryConfigs[name]
	if ok {
		return config
	}

	config = defaultRegistryConfig
	config.Name = name

	return config
}
