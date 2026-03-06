package managers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jroden2/stackforge/pkg/domain"
	"github.com/jroden2/stackforge/pkg/installer"
)

type CurlManager struct{}

func init() {
	installer.RegisterManager(CurlManager{})
}

func (mgr CurlManager) Name() string {
	return "curl"
}

func (mgr CurlManager) IsInstalled(pkg domain.Package) bool {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("command -v %s", pkg.InstallLogic.Identifier))
	return cmd.Run() == nil
}

func (mgr CurlManager) Install(pkg domain.Package) error {
	if mgr.IsInstalled(pkg) {
		return fmt.Errorf("%w", installer.ErrPackageAlreadyInstalled)
	}
	cmd := exec.Command("bash", "-c", fmt.Sprintf("curl -fsSL %s | bash", pkg.InstallLogic.Identifier))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Installing", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("installing %s failed: %w", pkg.Name, err)
	}
	return nil
}

func (mgr CurlManager) Upgrade(pkg domain.Package) error {
	if !mgr.IsInstalled(pkg) {
		return mgr.Install(pkg)
	}
	if pkg.InstallLogic.UpgradeScript == "" {
		return fmt.Errorf("%w for %s", installer.ErrUpgradeNotSupported, pkg.Name)
	}
	cmd := exec.Command("bash", "-c", fmt.Sprintf("curl -fsSL %s | bash", pkg.InstallLogic.UpgradeScript))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Upgrading", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("upgrading %s failed: %w", pkg.Name, err)
	}
	return nil
}

func (mgr CurlManager) Uninstall(pkg domain.Package) error {
	if !mgr.IsInstalled(pkg) {
		return fmt.Errorf("%w", installer.ErrPackageNotInstalled)
	}
	if pkg.InstallLogic.UninstallScript == "" {
		return fmt.Errorf("%w for %s", installer.ErrUninstallNotSupported, pkg.Name)
	}
	cmd := exec.Command("bash", "-c", fmt.Sprintf("curl -fsSL %s | bash", pkg.InstallLogic.UninstallScript))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("Uninstalling", pkg.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("uninstalling %s failed: %w", pkg.Name, err)
	}
	return nil
}
