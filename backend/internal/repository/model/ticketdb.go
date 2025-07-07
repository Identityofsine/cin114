package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type TicketDB struct {
	TicketId        int64     `json:"ticket_id"`
	EventId         int64     `json:"event_id"`
	StripePaymentId *string   `json:"stripe_payment_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func GetAllTickets() ([]TicketDB, db.DatabaseError) {
	query := "SELECT * FROM tickets ORDER BY created_at DESC"
	rows, err := db.Query[TicketDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetTicketById(ticketId int64) (*TicketDB, db.DatabaseError) {
	query := "SELECT * FROM tickets WHERE ticket_id = $1"
	rows, err := db.Query[TicketDB](query, ticketId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetTicketById", "Ticket not found", "ticket-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetTicketsByEventId(eventId int64) ([]TicketDB, db.DatabaseError) {
	query := "SELECT * FROM tickets WHERE event_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query[TicketDB](query, eventId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetTicketsByStripePaymentId(stripePaymentId string) ([]TicketDB, db.DatabaseError) {
	query := "SELECT * FROM tickets WHERE stripe_payment_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query[TicketDB](query, stripePaymentId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateTicket(ticket *TicketDB) db.DatabaseError {
	query := "INSERT INTO tickets (event_id, stripe_payment_id) VALUES ($1, $2) RETURNING ticket_id, created_at, updated_at"
	rows, err := db.Query[TicketDB](query, ticket.EventId, ticket.StripePaymentId)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		ticket.TicketId = (*rows)[0].TicketId
		ticket.CreatedAt = (*rows)[0].CreatedAt
		ticket.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateTicketStripePayment(ticketId int64, stripePaymentId *string) db.DatabaseError {
	query := "UPDATE tickets SET stripe_payment_id = $1, updated_at = CURRENT_TIMESTAMP WHERE ticket_id = $2"
	_, err := db.Query[TicketDB](query, stripePaymentId, ticketId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTicket(ticketId int64) db.DatabaseError {
	query := "DELETE FROM tickets WHERE ticket_id = $1"
	_, err := db.Query[TicketDB](query, ticketId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTicketsByEventId(eventId int64) db.DatabaseError {
	query := "DELETE FROM tickets WHERE event_id = $1"
	_, err := db.Query[TicketDB](query, eventId)
	if err != nil {
		return err
	}
	return nil
}

func TicketExists(ticketId int64) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM tickets WHERE ticket_id = $1)"
	rows, err := db.Query[bool](query, ticketId)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}

func GetTicketCountByEventId(eventId int64) (int, db.DatabaseError) {
	query := "SELECT COUNT(*) FROM tickets WHERE event_id = $1"
	rows, err := db.Query[int](query, eventId)
	if err != nil {
		return 0, err
	}
	if len(*rows) == 0 {
		return 0, nil
	}
	return (*rows)[0], nil
}
