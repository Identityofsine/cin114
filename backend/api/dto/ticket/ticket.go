package ticket

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

type OwnedTicket struct {
	TicketId           int64     `json:"ticket_id"`
	EventId            int64     `json:"event_id"`
	EventDescription   string    `json:"event_description"`
	StripeReceiptEmail *string   `json:"stripe_receipt_email"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// Mapping functions
func MapOwnedTicket(object OwnedTicketsDB) OwnedTicket {
	return OwnedTicket{
		TicketId:           object.TicketId,
		EventId:            object.EventId,
		EventDescription:   object.EventDescription,
		StripeReceiptEmail: object.StripeReceiptEmail,
		CreatedAt:          object.CreatedAt,
		UpdatedAt:          object.UpdatedAt,
	}
}

func MapOwnedTickets(objects []OwnedTicketsDB) []OwnedTicket {
	result := make([]OwnedTicket, len(objects))
	for i, object := range objects {
		result[i] = MapOwnedTicket(object)
	}
	return result
}
