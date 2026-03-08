package managers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jroden2/stackforge/pkg/domain"
)

var homebrewBootstrap = domain.Package{
	ID:   "homebrew",
	Name: "Homebrew",
	InstallLogic: domain.InstallLogic{
		Identifier: "brew",
		URL:        "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh",
	},
}

var nodeBootstrap = domain.Package{
	ID:   "nodejs",
	Name: "NVM + Node.js LTS",
	InstallLogic: domain.InstallLogic{
		Identifier: "node",
		URL:        "https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.4/install.sh",
	},
}

func bootstrapInstall(pkg domain.Package) error {
	fmt.Printf("%s is not installed, attempting to install via curl\n", pkg.Name)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("curl -fsSL %s | bash", pkg.InstallLogic.URL))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("bootstrap: failed to install %s: %w", pkg.Name, err)
	}
	return nil
}
