package model

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/db"

type OwnedTicketsDB struct {
	TicketId           int64   `json:"ticket_id"`
	EventDescription   string  `json:"event_description"`
	StripeReceiptEmail *string `json:"stripe_receipt_email"`
}

func GetAllOwnedTickets() ([]OwnedTicketsDB, db.DatabaseError) {
	query := "SELECT * FROM owned_tickets ORDER BY ticket_id DESC"
	rows, err := db.Query[OwnedTicketsDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetOwnedTicketsByEmail(email string) ([]OwnedTicketsDB, db.DatabaseError) {
	query := "SELECT * FROM owned_tickets WHERE stripe_receipt_email = $1 ORDER BY ticket_id DESC"
	rows, err := db.Query[OwnedTicketsDB](query, email)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetOwnedTicketById(ticketId int64) (*OwnedTicketsDB, db.DatabaseError) {
	query := "SELECT * FROM owned_tickets WHERE ticket_id = $1"
	rows, err := db.Query[OwnedTicketsDB](query, ticketId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetOwnedTicketById", "Owned ticket not found", "owned-ticket-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetOwnedTicketsByEventDescription(eventDescription string) ([]OwnedTicketsDB, db.DatabaseError) {
	query := "SELECT * FROM owned_tickets WHERE event_description ILIKE '%' || $1 || '%' ORDER BY ticket_id DESC"
	rows, err := db.Query[OwnedTicketsDB](query, eventDescription)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetPaidOwnedTickets() ([]OwnedTicketsDB, db.DatabaseError) {
	query := "SELECT * FROM owned_tickets WHERE stripe_receipt_email IS NOT NULL ORDER BY ticket_id DESC"
	rows, err := db.Query[OwnedTicketsDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetFreeOwnedTickets() ([]OwnedTicketsDB, db.DatabaseError) {
	query := "SELECT * FROM owned_tickets WHERE stripe_receipt_email IS NULL ORDER BY ticket_id DESC"
	rows, err := db.Query[OwnedTicketsDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}
