package service

import (
	"strconv"

	dto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/ticket"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

// GetTicketsByEventIdService returns all owned tickets for a specific event
func GetTicketsByEventIdService(eventId string) ([]dto.OwnedTicket, db.DatabaseError) {
	eventIdInt, parseErr := strconv.ParseInt(eventId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetTicketsByEventIdService", "Invalid event ID format", "invalid-event-id", 400)
	}

	ownedTickets, err := model.GetOwnedTicketsByEventId(eventIdInt)
	if err != nil {
		return nil, err
	}

	result := dto.MapOwnedTickets(ownedTickets)
	return result, nil
}

// GetAllOwnedTicketsService returns all owned tickets
func GetAllOwnedTicketsService() ([]dto.OwnedTicket, db.DatabaseError) {
	ownedTickets, err := model.GetAllOwnedTickets()
	if err != nil {
		return nil, err
	}

	result := dto.MapOwnedTickets(ownedTickets)
	return result, nil
}

// GetOwnedTicketsByEmailService returns all owned tickets for a specific email
func GetOwnedTicketsByEmailService(email string) ([]dto.OwnedTicket, db.DatabaseError) {
	ownedTickets, err := model.GetOwnedTicketsByEmail(email)
	if err != nil {
		return nil, err
	}

	result := dto.MapOwnedTickets(ownedTickets)
	return result, nil
}

// GetOwnedTicketByIdService returns a specific owned ticket by ID
func GetOwnedTicketByIdService(ticketId string) (*dto.OwnedTicket, db.DatabaseError) {
	ticketIdInt, parseErr := strconv.ParseInt(ticketId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetOwnedTicketByIdService", "Invalid ticket ID format", "invalid-ticket-id", 400)
	}

	ownedTicket, err := model.GetOwnedTicketById(ticketIdInt)
	if err != nil {
		return nil, err
	}

	result := dto.MapOwnedTicket(*ownedTicket)
	return &result, nil
}
