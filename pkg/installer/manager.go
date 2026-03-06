package installer

type Manager interface {
	Name() string
	Install() error
}
