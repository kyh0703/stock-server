package types

type Controller interface {
	Path() string
	Routes()
}
