package domain

type InstallLogic struct {
	Method     string   `json:"method"`
	Identifier string   `json:"identifier"`
	URL        string   `json:"url"`
	Options    []string `json:"options"`
	PipeTo     string   `json:"pipeTo"`
	// cURL
	UpgradeScript   string `json:"upgradeScript,omitempty"`
	UninstallScript string `json:"uninstallScript,omitempty"`
}
