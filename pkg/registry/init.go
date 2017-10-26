package registry

func init() {

	ConfigRegistry(DefaultRegistry, false)
	ConfigRegistry(QuayRegistry, false)
	ConfigRegistry(GoogleRegistry, false)
	ConfigRegistry(LocalhostRegistry, true)
}
