package model

func TicketFixture() Ticket {
	return Ticket{
		"name":        "TICKET-1",
		"priority":    "3-Low",
		"type":        "Task",
		"status":      "Done",
		"resolution":  "Resolved",
		"sprint":      "Sprint",
		"createdDate": "2021-02-28T00:00:00Z",
		"startDate":   "2021-03-01T00:00:00Z",
		"endDate":     "2021-03-02T00:00:00Z",
	}
}

func RecordFixture() []string {
	ticket := TicketFixture()
	return []string{
		ticket["name"],
		ticket["type"],
		ticket["priority"],
		ticket["createdDate"],
		ticket["startDate"],
		ticket["endDate"],
	}
}
