package installer

import "github.com/jroden2/stackforge/pkg/domain"

type Manager interface {
	Name() string
	Install(pkg domain.Package) error
}
