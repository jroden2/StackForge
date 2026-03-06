package managers

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/jroden2/stackforge/pkg/domain"
	"github.com/jroden2/stackforge/pkg/installer"
)

type BrewManager struct{}

func init() {
	installer.RegisterManager(BrewManager{})
}

func (mgr BrewManager) Name() string {
	return "homebrew"
}

func (mgr BrewManager) Install(pkg domain.Package) error {
	if mgr.IsInstalled(pkg) {
		fmt.Println(installer.ErrPackageAlreadyInstalled)
		return fmt.Errorf("%w", installer.ErrPackageAlreadyInstalled)
	}
	cmd := mgr.buildManageCmd("install", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Installing", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("installing %s failed: %w", pkg.InstallLogic.Identifier, err)
	}
	return nil
}

func (mgr BrewManager) Uninstall(pkg domain.Package) error {
	if !mgr.IsInstalled(pkg) {
		fmt.Println(installer.ErrPackageNotInstalled)
		return fmt.Errorf("%w", installer.ErrPackageNotInstalled)
	}
	cmd := mgr.buildManageCmd("uninstall", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Uninstalling", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("uninstalling %s failed: %w", pkg.InstallLogic.Identifier, err)
	}
	return nil
}

func (mgr BrewManager) Upgrade(pkg domain.Package) error {
	if !mgr.IsInstalled(pkg) {
		fmt.Println(installer.ErrPackageNotInstalled)
		return mgr.Install(pkg)
	}
	cmd := mgr.buildManageCmd("uninstall", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Upgrading", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("upgrading %s failed: %w", pkg.InstallLogic.Identifier, err)
	}
	return nil
}

// utils
func (mgr BrewManager) buildManageCmd(action string, pkg domain.Package) *exec.Cmd {
	if pkg.InstallLogic.Method == "cask" {
		return exec.Command("brew", action, "--cask", pkg.InstallLogic.Identifier)
	}
	return exec.Command("brew", action, pkg.InstallLogic.Identifier)
}

func (mgr BrewManager) IsInstalled(pkg domain.Package) bool {
	cmd := exec.Command("brew", "list", "--version", pkg.InstallLogic.Identifier)
	return cmd.Run() == nil
}

func RunInstall(mgr installer.Manager, pkg domain.Package, forced bool) error {
	if forced {
		return mgr.Upgrade(pkg)
	}
	if mgr.IsInstalled(pkg) {
		fmt.Println(installer.ErrPackageAlreadyInstalled)
		return nil
	}
	return mgr.Install(pkg)
}
func RunUninstall(mgr installer.Manager, pkg domain.Package) error {
	if err := mgr.Uninstall(pkg); err != nil {
		if errors.Is(err, installer.ErrPackageNotInstalled) {
			fmt.Println("Skipping", pkg.Name)
			return nil
		}
		return err
	}
	return nil
}
