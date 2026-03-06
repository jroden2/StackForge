package bundle

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/jroden2/stackforge/pkg/domain"
)

func LoadBundles(dir string) ([]domain.Bundle, error) {
	var bundles []domain.Bundle

	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		var bundle domain.Bundle
		if err := json.Unmarshal(data, &bundle); err != nil {
			return nil, err
		}

		bundles = append(bundles, bundle)
	}

	return bundles, nil
}
