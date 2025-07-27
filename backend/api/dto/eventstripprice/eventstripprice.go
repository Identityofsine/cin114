package eventstripprice

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

type EventStripePrice struct {
	Id            int64     `json:"id"`
	EventId       int64     `json:"event_id"`
	StripePriceId string    `json:"stripe_price_id"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type EventStripePriceWithDetails struct {
	Id                    int64      `json:"id"`
	EventId               int64      `json:"event_id"`
	EventDescription      string     `json:"event_description"`
	EventShortDescription *string    `json:"event_short_description"`
	EventExpirationDate   *time.Time `json:"event_expiration_date"`
	StripePriceId         string     `json:"stripe_price_id"`
	PriceActive           bool       `json:"price_active"`
	PriceCurrency         string     `json:"price_currency"`
	PriceType             string     `json:"price_type"`
	PriceUnitAmount       *int64     `json:"price_unit_amount"`
	IsActive              bool       `json:"is_active"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

type CreateEventStripePriceRequest struct {
	EventId int64 `json:"event_id" binding:"required"`
	Price   int64 `json:"price" binding:"required"`
}

type UpdateEventStripePriceActiveRequest struct {
	IsActive bool `json:"is_active" binding:"required"`
}

func Map(object EventStripePriceDB) EventStripePrice {
	return EventStripePrice{
		Id:            object.Id,
		EventId:       object.EventId,
		StripePriceId: object.StripePriceId,
		IsActive:      object.IsActive,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}
}

func MapWithDetails(object EventStripePriceWithDetailsDB) EventStripePriceWithDetails {
	return EventStripePriceWithDetails{
		Id:                    object.Id,
		EventId:               object.EventId,
		EventDescription:      object.EventDescription,
		EventShortDescription: object.EventShortDescription,
		EventExpirationDate:   object.EventExpirationDate,
		StripePriceId:         object.StripePriceId,
		PriceActive:           object.PriceActive,
		PriceCurrency:         object.PriceCurrency,
		PriceType:             object.PriceType,
		PriceUnitAmount:       object.PriceUnitAmount,
		IsActive:              object.IsActive,
		CreatedAt:             object.CreatedAt,
		UpdatedAt:             object.UpdatedAt,
	}
}

func ReverseMap(object EventStripePrice) EventStripePriceDB {
	return EventStripePriceDB{
		Id:            object.Id,
		EventId:       object.EventId,
		StripePriceId: object.StripePriceId,
		IsActive:      object.IsActive,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}
}

func MapCreateRequest(object CreateEventStripePriceRequest) EventStripePriceDB {
	return EventStripePriceDB{
		EventId:       object.EventId,
		StripePriceId: "",   // This will need to be set after creating the stripe price
		IsActive:      true, // Default to active when creating
	}
}
