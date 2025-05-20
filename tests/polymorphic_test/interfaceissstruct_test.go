package polymorphic

import (
	"testing"

	"github.com/dywoq/polymorphic"
) 

type testInterface interface {
	value() string
}

type s struct{} // implements testInterface
type d struct{} // doesn't implement testInterface

func (s) value() string { return "s" }

func TestInterfaceIsStruct(t *testing.T) {
	tests := []struct {
		name     string
		got      bool
		expected bool
	}{
		{
			name:     "s implements testInterface",
			got:      polymorphic.InterfaceIsStruct[s, testInterface](s{}),
			expected: true,
		},
		{
			name:     "d doesn't implement testInterface",
			got:      polymorphic.InterfaceIsStruct[s, testInterface](d{}),
			expected: false,
		},
	}

	for _, tt := range tests {
		if tt.got != tt.expected {
			t.Errorf("InterfaceIsStruct[T any]() got %t, expected %t (name: %s)", tt.got, tt.expected, tt.name)
		}
	}
}
