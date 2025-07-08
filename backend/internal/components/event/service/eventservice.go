package service

import (
	"strconv"

	dto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/event"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/payment"
)

// GetAllEventsService returns all events
func GetAllEventsService() ([]dto.Event, db.DatabaseError) {
	events, err := model.GetAllEvents()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Event, len(events))
	for i, event := range events {
		result[i] = dto.Map(event)
	}

	return result, nil
}

// GetEventByIdService returns a specific event by ID
func GetEventByIdService(id string) (*dto.Event, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetEventById", "Invalid ID format", "invalid-id", 400)
	}

	event, err := model.GetEventById(idInt)
	if err != nil {
		return nil, err
	}

	result := dto.Map(*event)
	return &result, nil
}

// GetActiveEventsService returns all active events (not expired)
func GetActiveEventsService() ([]dto.Event, db.DatabaseError) {
	events, err := model.GetActiveEvents()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Event, len(events))
	for i, event := range events {
		result[i] = dto.Map(event)
	}

	return result, nil
}

// CreateEventCheckoutService creates a checkout session for an event
func CreateEventCheckoutService(eventId string, request dto.CreateCheckoutRequest) (*dto.CheckoutResponse, db.DatabaseError) {
	eventIdInt, parseErr := strconv.ParseInt(eventId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("CreateEventCheckout", "Invalid event ID format", "invalid-event-id", 400)
	}

	// Create the checkout form for the payment service
	checkoutForm := payment.CheckoutSessionForm{
		EventId:    eventIdInt,
		Quantity:   request.Quantity,
		SuccessURL: request.SuccessURL,
		CancelURL:  request.CancelURL,
	}

	// Create the checkout session using the payment service
	session, err := payment.CreateCheckoutSessionForEvent(checkoutForm)
	if err != nil {
		return nil, err
	}

	// Return the response with checkout URL and session ID
	response := &dto.CheckoutResponse{
		CheckoutURL: session.URL,
		SessionId:   session.ID,
	}

	return response, nil
}

// EventExistsService checks if an event exists
func EventExistsService(id string) (bool, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return false, db.NewDatabaseError("EventExists", "Invalid ID format", "invalid-id", 400)
	}

	return model.EventExists(idInt)
}
