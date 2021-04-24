package config

import (
	"os"
	"reflect"
	"testing"
)

func TestInputContext(t *testing.T) {
	// GIVEN
	expected := &InputContext{
		Reader: os.Stdin,
	}

	// WHEN
	actual := NewInputContext()

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestInputContext: expected %v, actual %v",
			expected,
			actual,
		)
	}
}
