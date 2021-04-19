package operator

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadFromReader(t *testing.T) {
	// GIVEN
	delimitedText := "This|is|the|header\n" +
		"Fields|in|the|record\n"

	r := strings.NewReader(delimitedText)
	connection := NewReaderConnection(r)

	expected := map[string]interface{}{
		"delimitedText": delimitedText,
		"hasNext":       false,
	}

	// WHEN
	var data string
	input := NewDelimitedTextInput()
	iterator, err := input.Read(connection)

	if err != nil {
		t.Fatal(err)
	}

	hasNext := iterator.Next(&data)

	actual := map[string]interface{}{
		"delimitedText": data,
		"hasNext":       hasNext,
	}

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestReadFromReader: expected %v, actual %v",
			expected,
			actual,
		)
	}
}
