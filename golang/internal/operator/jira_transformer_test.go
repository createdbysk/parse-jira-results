package operator

import (
	"reflect"
	"testing"
)

type testDataStructure struct {
	field1 string
	field2 string
}

type field1Extractor struct {
}

func (e *field1Extractor) Extract(from interface{}, to map[string]interface{}) {
	to["field1"] = from.(testDataStructure).field1
}

type field2Extractor struct {
}

func (e *field2Extractor) Extract(from interface{}, to map[string]interface{}) {
	to["field2"] = from.(testDataStructure).field2
}

func TestExtractFieldsTransformer(t *testing.T) {
	// GIVEN
	ds := testDataStructure{"field1", "field2"}

	eft := NewExtractFieldsTransformer(
		&field1Extractor{},
		&field2Extractor{},
	)

	expected := map[string]interface{}{
		"field1": "field1",
		"field2": "field2",
	}

	// WHEN
	actual := make(map[string]interface{})
	eft.Transform(ds, actual)

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestExtractFieldsTransformer: expected %v, actual %v",
			expected,
			actual,
		)
	}
}
