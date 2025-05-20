package polymorphic

import (
	"reflect"
	"runtime"
)

// Interface is a structure that holds data about an interface.
// It uses reflect and runtime package to retrieve information from the interface.
type Interface struct {
	Name            string
	Type            reflect.Type
	Kind            reflect.Kind
	NumberOfMethods int
	MethodNames     []string
	PackagePath     string
	File            string
	Line            int
}

// FillInterface populates the fields of the provided Interface struct with information
// derived from the generic type T.
//
// If T is not an interface, or the function failed to get the current line and file,
// this function will panic.
//
// If you created an instance of Interface that way:
//
//	i := Interface{}
//
// Then you can fill the instance with FillInterface:
//
//	FillInterface[SomeInterface](&i)
func FillInterface[T any](i *Interface) {
	tType := reflect.TypeOf(new(T)).Elem()

	if tType.Kind() != reflect.Interface {
		panic("polymorphic: T is not an interface")
	}

	numMethods := tType.NumMethod()
	methodNames := make([]string, numMethods)
	for i := range numMethods {
		methodNames[i] = tType.Method(i).Name
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("didn't get the line and the file of the code")
	}

	i.Name = tType.Name()
	i.Type = tType
	i.Kind = tType.Kind()
	i.NumberOfMethods = numMethods
	i.MethodNames = methodNames
	i.PackagePath = tType.PkgPath()
	i.File = file
	i.Line = line
}

// NewInterface creates new instance of Interface structure, fills it with FillInterface
// and returns it.
func NewInterface[T any]() Interface {
	instance := Interface{}
	FillInterface[T](&instance)
	return instance
}
