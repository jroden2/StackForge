package bundle

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/jroden2/stackforge/pkg/domain"
)

func LoadBundles(dir string) ([]domain.Bundle, error) {
	var bundles []domain.Bundle
	if files, err := filepath.Glob(filepath.Join(dir, "*.json")); err != nil {
		return bundles, err
	} else {
		for _, file := range files {
			if data, err := os.ReadFile(file); err != nil {
				return bundles, err
			} else {
				var bundle domain.Bundle
				if err := json.Unmarshal(data, &bundle); err != nil {
					return bundles, err
				}
				bundles = append(bundles, bundle)
			}
		}
	}
}
