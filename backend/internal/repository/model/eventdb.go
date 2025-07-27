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
	VideoId          int64      `json:"video_id"`
	Cap              int64      `json:"capacity,omitempty"` // Optional field for event capacity
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

func GetEventByIdWithChildren(eventId int64) (*EventDB, []EventLocationDB, []EventImageDB, db.DatabaseError) {
	// Get the main event
	event, err := GetEventById(eventId)
	if err != nil {
		return nil, nil, nil, err
	}

	// Get locations for this event
	locations, err := GetLocationsByEventId(eventId)
	if err != nil {
		return event, nil, nil, err
	}

	// Get images for this event
	images, err := GetImagesByEventId(eventId)
	if err != nil {
		return event, locations, nil, err
	}

	return event, locations, images, nil
}

func GetAllEventsWithChildren() ([]EventDB, map[int64][]EventLocationDB, map[int64][]EventImageDB, db.DatabaseError) {
	// Get all events
	events, err := GetAllEvents()
	if err != nil {
		return nil, nil, nil, err
	}

	locationsMap := make(map[int64][]EventLocationDB)
	imagesMap := make(map[int64][]EventImageDB)

	// Get locations and images for each event
	for _, event := range events {
		locations, err := GetLocationsByEventId(event.EventId)
		if err != nil {
			return events, locationsMap, imagesMap, err
		}
		locationsMap[event.EventId] = locations

		images, err := GetImagesByEventId(event.EventId)
		if err != nil {
			return events, locationsMap, imagesMap, err
		}
		imagesMap[event.EventId] = images
	}

	return events, locationsMap, imagesMap, nil
}

func GetActiveEventsWithChildren() ([]EventDB, map[int64][]EventLocationDB, map[int64][]EventImageDB, db.DatabaseError) {
	// Get active events
	events, err := GetActiveEvents()
	if err != nil {
		return nil, nil, nil, err
	}

	locationsMap := make(map[int64][]EventLocationDB)
	imagesMap := make(map[int64][]EventImageDB)

	// Get locations and images for each event
	for _, event := range events {
		locations, err := GetLocationsByEventId(event.EventId)
		if err != nil {
			return events, locationsMap, imagesMap, err
		}
		locationsMap[event.EventId] = locations

		images, err := GetImagesByEventId(event.EventId)
		if err != nil {
			return events, locationsMap, imagesMap, err
		}
		imagesMap[event.EventId] = images
	}

	return events, locationsMap, imagesMap, nil
}
