package installer

import (
	"errors"

	"github.com/jroden2/stackforge/pkg/domain"
)

var (
	ErrPackageNotInstalled     = errors.New("package not installed")
	ErrPackageAlreadyInstalled = errors.New("package already installed")
	ErrUninstallNotSupported   = errors.New("uninstall not supported for manager")
)

type Manager interface {
	Name() string
	Install(pkg domain.Package) error
	IsInstalled(pkg domain.Package) bool
	Upgrade(pkg domain.Package) error
	Uninstall(pkg domain.Package) error
}
