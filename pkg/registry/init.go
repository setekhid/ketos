package registry

func init() {

	ConfigRegistry(DefaultRegistry, false, "ktm7bvr4.mirror.aliyuncs.com")
	ConfigRegistry(QuayRegistry, false, "")
	ConfigRegistry(GoogleRegistry, false, "")
	ConfigRegistry(LocalhostRegistry, true, "")
	ConfigRegistry(LocalComposeRegistry, true, "")
}
