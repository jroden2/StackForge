package domain

type Package struct {
	ID           string       "json:id"
	Name         string       "json:name"
	Manager      string       "json:manager"
	InstallLogic InstallLogic "json:installLogic"
}
