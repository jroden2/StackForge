package installer

import "github.com/jroden2/stackforge/pkg/domain"

type Manager interface {
	Name() string
	Install(pkg domain.Package) error
	IsInstalled(pkg domain.Package) bool
	Upgrade(pkg domain.Package) error
	Uninstall(pkg domain.Package) error
}
