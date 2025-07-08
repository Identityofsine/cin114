package service

import (
	"strconv"

	dto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/event"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/payment"
)

// GetAllEventsService returns all events with their locations and images
func GetAllEventsService() ([]dto.Event, db.DatabaseError) {
	events, locationsMap, imagesMap, err := model.GetAllEventsWithChildren()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Event, len(events))
	for i, event := range events {
		locations := locationsMap[event.EventId]
		images := imagesMap[event.EventId]
		result[i] = dto.MapWithChildren(event, locations, images)
	}

	return result, nil
}

// GetEventByIdService returns a specific event by ID with its locations and images
func GetEventByIdService(id string) (*dto.Event, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetEventById", "Invalid ID format", "invalid-id", 400)
	}

	event, locations, images, err := model.GetEventByIdWithChildren(idInt)
	if err != nil {
		return nil, err
	}

	result := dto.MapWithChildren(*event, locations, images)
	return &result, nil
}

// GetActiveEventsService returns all active events (not expired) with their locations and images
func GetActiveEventsService() ([]dto.Event, db.DatabaseError) {
	events, locationsMap, imagesMap, err := model.GetActiveEventsWithChildren()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Event, len(events))
	for i, event := range events {
		locations := locationsMap[event.EventId]
		images := imagesMap[event.EventId]
		result[i] = dto.MapWithChildren(event, locations, images)
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
