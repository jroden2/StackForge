package managers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jroden2/stackforge/pkg/domain"
	"github.com/jroden2/stackforge/pkg/installer"
)

type NpmManager struct{}

func init() {
	installer.RegisterManager(NpmManager{})
}

func (mgr NpmManager) Name() string {
	return "npm"
}

func (mgr NpmManager) IsInstalled(pkg domain.Package) bool {
	cmd := exec.Command("npm", "list", "-g", "--depth=0", pkg.InstallLogic.Identifier)
	return cmd.Run() == nil
}

func (mgr NpmManager) Install(pkg domain.Package) error {
	if mgr.IsInstalled(pkg) {
		return fmt.Errorf("%w", installer.ErrPackageAlreadyInstalled)
	}
	cmd := exec.Command("npm", "install", "-g", pkg.InstallLogic.Identifier)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Installing", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("installing %s failed: %w", pkg.Name, err)
	}
	return nil
}

func (mgr NpmManager) Upgrade(pkg domain.Package) error {
	if !mgr.IsInstalled(pkg) {
		return mgr.Install(pkg)
	}
	cmd := exec.Command("npm", "update", "-g", pkg.InstallLogic.Identifier)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Upgrading", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("upgrading %s failed: %w", pkg.Name, err)
	}
	return nil
}

func (mgr NpmManager) Uninstall(pkg domain.Package) error {
	if !mgr.IsInstalled(pkg) {
		return fmt.Errorf("%w", installer.ErrPackageNotInstalled)
	}
	cmd := exec.Command("npm", "uninstall", "-g", pkg.InstallLogic.Identifier)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Uninstalling", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("uninstalling %s failed: %w", pkg.Name, err)
	}
	return nil
}

func (mgr NpmManager) checkDependencies() error {
	// Todo: call the dep check - install if needed
	return nil
}
