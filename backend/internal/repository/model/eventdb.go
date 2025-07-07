package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type EventDB struct {
	EventId          int64      `json:"event_id"`
	Description      string     `json:"description"`
	ShortDescription *string    `json:"short_description"`
	ExpirationDate   *time.Time `json:"expiration_date"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

func GetAllEvents() ([]EventDB, db.DatabaseError) {
	query := "SELECT * FROM events ORDER BY created_at DESC"
	rows, err := db.Query[EventDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetEventById(eventId int64) (*EventDB, db.DatabaseError) {
	query := "SELECT * FROM events WHERE event_id = $1"
	rows, err := db.Query[EventDB](query, eventId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetEventById", "Event not found", "event-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetActiveEvents() ([]EventDB, db.DatabaseError) {
	query := "SELECT * FROM events WHERE expiration_date IS NULL OR expiration_date > NOW() ORDER BY created_at DESC"
	rows, err := db.Query[EventDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateEvent(event *EventDB) db.DatabaseError {
	query := "INSERT INTO events (description, short_description, expiration_date) VALUES ($1, $2, $3) RETURNING event_id, created_at, updated_at"
	rows, err := db.Query[EventDB](query, event.Description, event.ShortDescription, event.ExpirationDate)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		event.EventId = (*rows)[0].EventId
		event.CreatedAt = (*rows)[0].CreatedAt
		event.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateEvent(event *EventDB) db.DatabaseError {
	query := "UPDATE events SET description = $1, short_description = $2, expiration_date = $3, updated_at = CURRENT_TIMESTAMP WHERE event_id = $4"
	_, err := db.Query[EventDB](query, event.Description, event.ShortDescription, event.ExpirationDate, event.EventId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(eventId int64) db.DatabaseError {
	query := "DELETE FROM events WHERE event_id = $1"
	_, err := db.Query[EventDB](query, eventId)
	if err != nil {
		return err
	}
	return nil
}

func EventExists(eventId int64) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM events WHERE event_id = $1)"
	rows, err := db.Query[bool](query, eventId)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
