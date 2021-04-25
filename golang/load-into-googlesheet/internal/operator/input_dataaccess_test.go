package operator

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

type mockNewReaderConnection struct {
	c                           Connection
	reader                      io.Reader
	previousNewReaderConnection newReaderConnectionType
}

// NewMockNewReaderConnection mocks operator.NewReaderConnection.
// It returns the MockFunction interface.
// Caller must call mockFunction.Unpatch(). The idomatic way to do this
//
//		mockFunction := NewMockReaderConnection(connection)
// 		defer mockFunction.Unpatch()
func NewMockNewReaderConnection(connection Connection) MockFunction {
	c := &mockNewReaderConnection{
		c:                           connection,
		previousNewReaderConnection: newReaderConnection,
	}
	// Set the global variable
	newReaderConnection = c.newReaderConnection

	return c
}

func (c *MockNewReaderConnection) newReaderConnection(reader io.Reader) Connection {
	c.reader = reader
	return c.Connection
}

func (c *mockNewReaderConnection) Unpatch() {
	// Restore the global variable
	newReaderConnection = c.previousNewReaderConnection
}

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
