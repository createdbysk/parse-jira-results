package operator

import (
	"bytes"
	"testing"
)

func TestCSVFileLoader(t *testing.T) {
	// GIVEN
	writer := bytes.NewBufferString("")
	delimiter := '?'
	record := []string{
		"Hello",
		"World!",
	}
	expected := "Hello?World!\n"

	// WHEN
	var loader Loader
	var actual string
	var err error
	if loader, err = NewCSVFileLoader(writer, delimiter); err == nil {
		if err = loader.Load(record); err == nil {
			loader.Flush()
			actual = writer.String()
		}
	}

	// THEN
	if err != nil {
		t.Errorf("NewCSVFileLoader() failed: Error = %v", err)
	}
	if loader == nil {
		t.Error("NewCSVFileLoader() returned nil loader.")
	}
	if actual != expected {
		t.Errorf(
			"NewCSVFileLoader() did not prouce the expected result."+
				"Expected %v, Actual %v",
			expected, actual,
		)
	}
}
