package polymorphic

import "reflect"

// Implements checks, if object implements T interface.
// It panics if T is not an interface.
// Returns boolean.
func Implements[T any](object interface{}) bool {
	if tType := reflect.TypeOf(new(T)).Elem(); tType.Kind() != reflect.Interface {
		panic("T is not an interface")
	}
	_, ok := any(object).(T)
	return ok
}
