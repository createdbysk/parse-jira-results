package model

import "time"

type Ticket struct {
	Name       string    `json: "name"`
	Priority   string    `json: "priority"`
	Type       string    `json: "type"`
	Status     string    `json: "status"`
	Resolution string    `json: "resolution"`
	Sprint     string    `json: "sprint"`
	StartDate  time.Time `json: "startDate"`
	EndDate    time.Time `json: "endDate"`
}
