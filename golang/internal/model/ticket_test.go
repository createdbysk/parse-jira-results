package model

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestUnmarshalJSON(t *testing.T) {
	// GIVEN
	serialized, _ := json.Marshal(map[string]string{
		"name":      "TICKET-1",
		"startDate": "2021-03-01T00:00:00Z",
		"endDate":   "2021-03-02T00:00:00Z",
		"priority":  "3-Low",
		"type":      "Task",
	})
	startDate, _ := time.Parse("2006-01-02", "2021-03-01")
	endDate, _ := time.Parse("2006-01-02", "2021-03-02")
	expected := Ticket{
		Name:      "TICKET-1",
		StartDate: startDate,
		EndDate:   endDate,
		Priority:  "3-Low",
		Type:      "Task",
	}

	// WHEN
	var actual Ticket
	err := json.Unmarshal(serialized, &actual)

	if err != nil {
		t.Errorf("Ticket.UnmarshalJSON %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"Ticket.UnmarshalJSON expected: %v actual: %v",
			expected, actual,
		)
	}
}
