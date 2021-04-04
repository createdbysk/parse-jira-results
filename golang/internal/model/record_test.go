package model

import (
	"reflect"
	"testing"
)

func TestNewRecord(t *testing.T) {
	// GIVEN
	fields := []string{
		"name",
		"type",
		"priority",
		"createdDate",
		"startDate",
		"endDate",
	}
	ticket := TicketFixture()

	expected := RecordFixture()

	// WHEN
	actual, err := NewRecord(ticket, fields)

	// THEN
	if err != nil {
		t.Errorf("NewRecord errored: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"NewRecord() expected: %v actual: %v",
			expected, actual,
		)
	}
}
