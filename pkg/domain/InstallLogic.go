package domain

type InstallLogic struct {
	Method     string   `json:"method"`
	Identifier string   `json:"identifier"`
	URL        string   `json:"url"`
	Options    []string `json:"options"`
	PipeTo     string   `json:"pipeTo"`
}
