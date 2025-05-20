package polymorphic_test

import (
	"testing"

	"github.com/dywoq/polymorphic"
)

type testInterface interface {
	value() string
}

type testA struct{} // implements testInterface
type testB struct{} // doesn't implement testInterface

func (ta testA) value() string {
	return "HI!"
}

func TestImplements(t *testing.T) {
	a := testA{}
	b := testB{}
	tests := []struct {
		name     string
		got      bool
		expected bool
	}{
		{
			name:     "testA implements testInterface",
			got:      polymorphic.Implements[testInterface](a),
			expected: true,
		},
		{
			name:     "testB doesn't implement testInterface",
			got:      polymorphic.Implements[testInterface](b),
			expected: false,
		},
	}

	for _, tt := range tests {
		if tt.got != tt.expected {
			t.Errorf("Implements[T any]() got %t, expected %t (name: %s)", tt.got, tt.expected, tt.name)
		}
	}
}
