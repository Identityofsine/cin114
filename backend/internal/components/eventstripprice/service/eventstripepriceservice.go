package service

import (
	"strconv"

	eventDto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/event"
	dto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/eventstripprice"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/payment"
)

// GetAllEventStripePricesService returns all event stripe price associations
func GetAllEventStripePricesService() ([]dto.EventStripePrice, db.DatabaseError) {
	eventStripePrices, err := model.GetAllEventStripePrices()
	if err != nil {
		return nil, err
	}

	result := make([]dto.EventStripePrice, len(eventStripePrices))
	for i, esp := range eventStripePrices {
		result[i] = dto.Map(esp)
	}

	return result, nil
}

// GetEventStripePriceByIdService returns a specific event stripe price association by ID
func GetEventStripePriceByIdService(id string) (*dto.EventStripePrice, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetEventStripePriceById", "Invalid ID format", "invalid-id", 400)
	}

	eventStripePrice, err := model.GetEventStripePriceById(idInt)
	if err != nil {
		return nil, err
	}

	result := dto.Map(*eventStripePrice)
	return &result, nil
}

// GetEventStripePricesByEventIdService returns all price associations for a specific event
func GetEventStripePricesByEventIdService(eventId string) ([]dto.EventStripePrice, db.DatabaseError) {
	eventIdInt, parseErr := strconv.ParseInt(eventId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetEventStripePricesByEventId", "Invalid event ID format", "invalid-event-id", 400)
	}

	eventStripePrices, err := model.GetEventStripePricesByEventId(eventIdInt)
	if err != nil {
		return nil, err
	}

	result := make([]dto.EventStripePrice, len(eventStripePrices))
	for i, esp := range eventStripePrices {
		result[i] = dto.Map(esp)
	}

	return result, nil
}

// GetActiveEventStripePriceByEventIdService returns the active price for a specific event
func GetActiveEventStripePriceByEventIdService(eventId string) (*dto.EventStripePrice, db.DatabaseError) {
	eventIdInt, parseErr := strconv.ParseInt(eventId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetActiveEventStripePriceByEventId", "Invalid event ID format", "invalid-event-id", 400)
	}

	eventStripePrice, err := model.GetActiveEventStripePriceByEventId(eventIdInt)
	if err != nil {
		return nil, err
	}

	result := dto.Map(*eventStripePrice)
	return &result, nil
}

// GetEventStripePricesByStripePriceIdService returns all event associations for a specific Stripe price
func GetEventStripePricesByStripePriceIdService(stripePriceId string) ([]dto.EventStripePrice, db.DatabaseError) {
	eventStripePrices, err := model.GetEventStripePricesByStripePriceId(stripePriceId)
	if err != nil {
		return nil, err
	}

	result := make([]dto.EventStripePrice, len(eventStripePrices))
	for i, esp := range eventStripePrices {
		result[i] = dto.Map(esp)
	}

	return result, nil
}

// GetActiveEventStripePricesService returns all active event stripe price associations
func GetActiveEventStripePricesService() ([]dto.EventStripePrice, db.DatabaseError) {
	eventStripePrices, err := model.GetActiveEventStripePrices()
	if err != nil {
		return nil, err
	}

	result := make([]dto.EventStripePrice, len(eventStripePrices))
	for i, esp := range eventStripePrices {
		result[i] = dto.Map(esp)
	}

	return result, nil
}

// GetEventStripePricesWithDetailsService returns event stripe prices with joined event and price details
func GetEventStripePricesWithDetailsService() ([]dto.EventStripePriceWithDetails, db.DatabaseError) {
	eventStripePrices, err := model.GetEventStripePricesWithDetails()
	if err != nil {
		return nil, err
	}

	result := make([]dto.EventStripePriceWithDetails, len(eventStripePrices))
	for i, esp := range eventStripePrices {
		result[i] = dto.MapWithDetails(esp)
	}

	return result, nil
}

// CreateEventStripePriceService creates a new event stripe price association
func CreateEventStripePriceService(request dto.CreateEventStripePriceRequest) (*dto.EventStripePrice, db.DatabaseError) {

	eventdb, err := model.GetEventById(request.EventId)
	if err != nil {
		return nil, err
	}

	event := eventDto.Map(*eventdb)

	obj, err := payment.CreateStripePriceForScreening(
		payment.EventStripePriceForm{
			Event: event,
			Price: request.Price,
		})
	if err != nil {
		return nil, err
	}

	eventStripePrice := dto.Map(*obj)

	return &eventStripePrice, nil
}

// UpdateEventStripePriceActiveService updates the active status of an event stripe price
func UpdateEventStripePriceActiveService(id string, request dto.UpdateEventStripePriceActiveRequest) db.DatabaseError {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return db.NewDatabaseError("UpdateEventStripePriceActive", "Invalid ID format", "invalid-id", 400)
	}

	return model.UpdateEventStripePriceActive(idInt, request.IsActive)
}

// SetActiveEventStripePriceForEventService sets a specific price as active for an event (deactivates others)
func SetActiveEventStripePriceForEventService(eventId string, stripePriceId string) db.DatabaseError {
	eventIdInt, parseErr := strconv.ParseInt(eventId, 10, 64)
	if parseErr != nil {
		return db.NewDatabaseError("SetActiveEventStripePriceForEvent", "Invalid event ID format", "invalid-event-id", 400)
	}

	return model.SetActiveEventStripePriceForEvent(eventIdInt, stripePriceId)
}

// DeleteEventStripePriceService deletes an event stripe price association
func DeleteEventStripePriceService(id string) db.DatabaseError {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return db.NewDatabaseError("DeleteEventStripePrice", "Invalid ID format", "invalid-id", 400)
	}

	return model.DeleteEventStripePrice(idInt)
}
