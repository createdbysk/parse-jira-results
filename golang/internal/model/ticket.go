package model

import "time"

type Ticket struct {
	Name      string    `json: "name"`
	StartDate time.Time `json: "startDate"`
	EndDate   time.Time `json: "endDate"`
	Priority  string    `json: "priority"`
	Type      string    `json: "type"`
}
