package model

import (
	"reflect"
	"testing"
)

func TestNewTicket(t *testing.T) {
	// GIVEN
	input := map[string]string{
		"name":        "TICKET-1",
		"priority":    "3-Low",
		"type":        "Task",
		"status":      "Open",
		"resolution":  "Unresolved",
		"sprint":      "Sprint",
		"createdDate": "2021-02-28T00:00:00Z",
		"startDate":   "2021-03-01T00:00:00Z",
		"endDate":     "2021-03-02T00:00:00Z",
	}

	expected := TicketFixture()

	// WHEN
	actual, err := NewTicket(input)

	// THEN
	if err != nil {
		t.Errorf("NewTicket errored: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"NewTicket() expected: %v actual: %v",
			expected, actual,
		)
	}
}
