package installer

import "fmt"

var registry = map[string]Manager{}

func RegisterManager(mgr Manager) {
	registry[mgr.Name()] = mgr
}

func GetManager(name string) (Manager, bool) {
	mgr, ok := registry[name]
	fmt.Println(name)
	return mgr, ok
}
