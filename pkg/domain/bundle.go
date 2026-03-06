package domain

type Bundle struct {
	BundleID    string    `json:"bundleId"`
	Title       string    `json:"title"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	Packages    []Package `json:"packages"`
}
