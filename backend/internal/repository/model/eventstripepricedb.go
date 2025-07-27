package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type EventStripePriceDB struct {
	Id            int64     `json:"id"`
	EventId       int64     `json:"event_id"`
	StripePriceId string    `json:"stripe_price_id"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// EventStripePriceWithDetailsDB includes event and price details
type EventStripePriceWithDetailsDB struct {
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

func GetAllEventStripePrices() ([]EventStripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM event_stripe_price ORDER BY created_at DESC"
	rows, err := db.Query[EventStripePriceDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetEventStripePriceById(id int64) (*EventStripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM event_stripe_price WHERE id = $1"
	rows, err := db.Query[EventStripePriceDB](query, id)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetEventStripePriceById", "Event stripe price not found", "event-stripe-price-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetEventStripePricesByEventId(eventId int64) ([]EventStripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM event_stripe_price WHERE event_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query[EventStripePriceDB](query, eventId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetActiveEventStripePriceByEventId(eventId int64) (*EventStripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM event_stripe_price WHERE event_id = $1 AND is_active = true"
	rows, err := db.Query[EventStripePriceDB](query, eventId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetActiveEventStripePriceByEventId", "Active event stripe price not found", "active-event-stripe-price-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetEventStripePricesByStripePriceId(stripePriceId string) ([]EventStripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM event_stripe_price WHERE stripe_price_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query[EventStripePriceDB](query, stripePriceId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetActiveEventStripePrices() ([]EventStripePriceDB, db.DatabaseError) {
	query := "SELECT * FROM event_stripe_price WHERE is_active = true ORDER BY created_at DESC"
	rows, err := db.Query[EventStripePriceDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetEventStripePricesWithDetails() ([]EventStripePriceWithDetailsDB, db.DatabaseError) {
	query := `
		SELECT 
			esp.id,
			esp.event_id,
			e.description AS event_description,
			e.short_description AS event_short_description,
			e.expiration_date AS event_expiration_date,
			esp.stripe_price_id,
			sp.active AS price_active,
			sp.currency AS price_currency,
			sp.type AS price_type,
			sp.unit_amount AS price_unit_amount,
			esp.is_active,
			esp.created_at,
			esp.updated_at
		FROM event_stripe_price esp
		JOIN events e ON esp.event_id = e.event_id
		JOIN stripe_prices sp ON esp.stripe_price_id = sp.id
		ORDER BY esp.created_at DESC
	`
	rows, err := db.Query[EventStripePriceWithDetailsDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateEventStripePrice(eventStripePrice *EventStripePriceDB) db.DatabaseError {
	query := `INSERT INTO event_stripe_price (
		event_id, stripe_price_id, is_active
	) VALUES (
		$1, $2, $3
	) RETURNING id, created_at, updated_at`

	type CreateResult struct {
		Id        int64     `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	rows, err := db.Query[CreateResult](query,
		eventStripePrice.EventId, eventStripePrice.StripePriceId, eventStripePrice.IsActive,
	)

	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		eventStripePrice.Id = (*rows)[0].Id
		eventStripePrice.CreatedAt = (*rows)[0].CreatedAt
		eventStripePrice.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateEventStripePriceActive(id int64, isActive bool) db.DatabaseError {
	query := "UPDATE event_stripe_price SET is_active = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2"
	_, err := db.Query[EventStripePriceDB](query, isActive, id)
	if err != nil {
		return err
	}
	return nil
}

func DeactivateAllEventStripePricesForEvent(eventId int64) db.DatabaseError {
	query := "UPDATE event_stripe_price SET is_active = false, updated_at = CURRENT_TIMESTAMP WHERE event_id = $1"
	_, err := db.Query[EventStripePriceDB](query, eventId)
	if err != nil {
		return err
	}
	return nil
}

func SetActiveEventStripePriceForEvent(eventId int64, stripePriceId string) db.DatabaseError {
	// First deactivate all existing prices for this event
	err := DeactivateAllEventStripePricesForEvent(eventId)
	if err != nil {
		return err
	}

	// Check if this event-price combination already exists
	query := "SELECT id FROM event_stripe_price WHERE event_id = $1 AND stripe_price_id = $2"
	rows, err := db.Query[EventStripePriceDB](query, eventId, stripePriceId)
	if err != nil {
		return err
	}

	if len(*rows) > 0 {
		// Update existing record to active
		return UpdateEventStripePriceActive((*rows)[0].Id, true)
	} else {
		// Create new active record
		newRecord := &EventStripePriceDB{
			EventId:       eventId,
			StripePriceId: stripePriceId,
			IsActive:      true,
		}
		return CreateEventStripePrice(newRecord)
	}
}

func DeleteEventStripePrice(id int64) db.DatabaseError {
	query := "DELETE FROM event_stripe_price WHERE id = $1"
	_, err := db.Query[EventStripePriceDB](query, id)
	if err != nil {
		return err
	}
	return nil
}

func EventStripePriceExists(eventId int64, stripePriceId string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM event_stripe_price WHERE event_id = $1 AND stripe_price_id = $2)"
	rows, err := db.Query[bool](query, eventId, stripePriceId)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
