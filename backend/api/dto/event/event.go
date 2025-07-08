package event

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

type Event struct {
	EventId          int64      `json:"event_id"`
	Description      string     `json:"description"`
	ShortDescription *string    `json:"short_description"`
	ExpirationDate   *time.Time `json:"expiration_date"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type CreateEventRequest struct {
	Description      string     `json:"description" binding:"required"`
	ShortDescription *string    `json:"short_description"`
	ExpirationDate   *time.Time `json:"expiration_date"`
}

type UpdateEventRequest struct {
	Description      string     `json:"description" binding:"required"`
	ShortDescription *string    `json:"short_description"`
	ExpirationDate   *time.Time `json:"expiration_date"`
}

type CreateCheckoutRequest struct {
	Quantity   int64  `json:"quantity" binding:"required"`
	SuccessURL string `json:"success_url" `
	CancelURL  string `json:"cancel_url" `
}

type CheckoutResponse struct {
	CheckoutURL string `json:"checkout_url"`
	SessionId   string `json:"session_id"`
}

func Map(object EventDB) Event {
	return Event{
		EventId:          object.EventId,
		Description:      object.Description,
		ShortDescription: object.ShortDescription,
		ExpirationDate:   object.ExpirationDate,
		CreatedAt:        object.CreatedAt,
		UpdatedAt:        object.UpdatedAt,
	}
}

func ReverseMap(object Event) EventDB {
	return EventDB{
		EventId:          object.EventId,
		Description:      object.Description,
		ShortDescription: object.ShortDescription,
		ExpirationDate:   object.ExpirationDate,
		CreatedAt:        object.CreatedAt,
		UpdatedAt:        object.UpdatedAt,
	}
}

func MapCreateRequest(object CreateEventRequest) EventDB {
	return EventDB{
		Description:      object.Description,
		ShortDescription: object.ShortDescription,
		ExpirationDate:   object.ExpirationDate,
	}
}

func MapUpdateRequest(object UpdateEventRequest) EventDB {
	return EventDB{
		Description:      object.Description,
		ShortDescription: object.ShortDescription,
		ExpirationDate:   object.ExpirationDate,
	}
}
