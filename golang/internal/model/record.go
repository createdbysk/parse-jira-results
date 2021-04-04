package model

import (
	"errors"
	"fmt"
)

func NewRecord(ticket Ticket, fields []string) ([]string, error) {
	result := make([]string, len(fields))

	for i, field := range fields {
		if value, exists := ticket[field]; exists {
			result[i] = value
		} else {
			return nil, errors.New(fmt.Sprintf("Field %s not present in ticket", field))
		}
	}

	return result, nil
}
