package polymorphic

// Implements checks, if object implements T interface.
// Returns boolean.
func Implements[T any](object interface{}) bool {
	_, ok := any(object).(T)
	return ok
}
