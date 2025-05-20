package polymorphic

import "reflect"

// InterfaceIsStruct checks if an arbitrary 'object' can implement interface I,
// and if its underlying concrete type is the specific struct type S.
//
// Type parameter S must be a struct. The function will panic if S is not a struct when invoked.
// Type parameter I must be an interface. The function will panic if I is not an interface when invoked.
//
// The function returns true if:
// 1. 'object' is a valid (non-nil) value.
// 2. The concrete type of 'object' implements interface I.
// 3. The concrete type of 'object' is precisely the struct type S.
//
// It returns false otherwise.
func InterfaceIsStruct[S any, I any](object any) bool {
	var sType S
	structReflectionType := reflect.TypeOf(sType)

	if structReflectionType.Kind() == reflect.Ptr {
		structReflectionType = structReflectionType.Elem()
	}
	if structReflectionType.Kind() != reflect.Struct {
		panic("polymorphic: Type parameter S must be a struct type (or a pointer to a struct)")
	}

	var iType I
	interfaceReflectionType := reflect.TypeOf(&iType).Elem()
	if interfaceReflectionType.Kind() != reflect.Interface {
		panic("polymorphic: Type parameter I must be an interface type")
	}

	objectValue := reflect.ValueOf(object)
	if !objectValue.IsValid() {
		return false
	}

	objectConcreteType := objectValue.Type()
	if objectConcreteType.Kind() == reflect.Ptr {
		objectConcreteType = objectConcreteType.Elem()
	}

	if !objectConcreteType.Implements(interfaceReflectionType) {
		return false
	}

	return objectConcreteType == structReflectionType
}
