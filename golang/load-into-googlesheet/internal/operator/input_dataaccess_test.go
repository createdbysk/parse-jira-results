package operator

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"local.dev/sheetsLoader/internal/config"
)

func readerFixture(data string) io.Reader {
	return strings.NewReader(data)
}

func mockInputContextFixture(reader io.Reader) *config.InputContext {
	return &config.InputContext{Reader: reader}
}

func TestReadFromReader(t *testing.T) {
	// GIVEN
	delimitedText := "This|is|the|header\n" +
		"Fields|in|the|record\n"

	r := readerFixture(delimitedText)
	inputContext := mockInputContextFixture(r)
	connection := NewReaderConnection(inputContext)

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
