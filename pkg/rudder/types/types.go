package types

// Resource represents a generic resource.
type Resource[T any] struct {
	ID    string
	Value T
}
