package installer

var registry = map[string]Manager{}

func RegisterManager(mgr Manager) {
	registry[mgr.Name()] = mgr
}

func GetManager(name string) (Manager, bool) {
	mgr, ok := registry[name]
	return mgr, ok
}
