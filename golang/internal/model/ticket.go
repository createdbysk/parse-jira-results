package model

type Ticket map[string]string

func NewTicket(input map[string]string) (Ticket, error) {
	return Ticket(input), nil
}
