package model

import "time"

type Ticket struct {
	Name      string    `json: "name"`
	Priority  string    `json: "priority"`
	Type      string    `json: "type"`
	StartDate time.Time `json: "startDate"`
	EndDate   time.Time `json: "endDate"`
}
