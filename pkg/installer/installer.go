package installer

import (
	"errors"
	"fmt"

	"github.com/jroden2/stackforge/pkg/domain"
)

func InstallBundle(bundle domain.Bundle) error {
	fmt.Println("Installing bundle ...")
	for _, pkg := range bundle.Packages {
		if manager, ok := GetManager(pkg.Manager); !ok {
			return errors.New(fmt.Sprintf("Package %s manager not found", pkg.Name))
		} else {
			err := manager.Install(pkg)
			if err != nil {
				fmt.Printf("Package %s install failed: %s\n", pkg.Name, err)
				return err
			}
		}
	}
	return nil
}
