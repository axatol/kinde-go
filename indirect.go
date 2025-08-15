package kinde

func Indirect[T any](value T) *T {
	return &value
}
