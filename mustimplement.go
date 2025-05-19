package polymorphic

// MustImplement checks if object implements T interface and panics if it doesn't.
// Useful for assertions at initialization time.
func MustImplement[T any](object interface{}, message string) {
	if !Implements[T](object) {
		panic(message)
	}
}
