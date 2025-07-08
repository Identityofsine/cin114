package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type EventLocationDB struct {
	EventLocationId     int64     `json:"event_location_id"`
	EventId             int64     `json:"event_id"`
	LocationName        string    `json:"location_name"`
	LocationDescription *string   `json:"location_description"`
	Longitude           float64   `json:"longitude"`
	Latitude            float64   `json:"latitude"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func GetLocationsByEventId(eventId int64) ([]EventLocationDB, db.DatabaseError) {
	query := "SELECT * FROM event_locations WHERE event_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query[EventLocationDB](query, eventId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetLocationById(locationId int64) (*EventLocationDB, db.DatabaseError) {
	query := "SELECT * FROM event_locations WHERE event_location_id = $1"
	rows, err := db.Query[EventLocationDB](query, locationId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetLocationById", "Location not found", "location-not-found", 404)
	}
	return &(*rows)[0], nil
}

func CreateEventLocation(location *EventLocationDB) db.DatabaseError {
	query := "INSERT INTO event_locations (event_id, location_name, location_description, longitude, latitude) VALUES ($1, $2, $3, $4, $5) RETURNING event_location_id, created_at, updated_at"
	rows, err := db.Query[EventLocationDB](query, location.EventId, location.LocationName, location.LocationDescription, location.Longitude, location.Latitude)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		location.EventLocationId = (*rows)[0].EventLocationId
		location.CreatedAt = (*rows)[0].CreatedAt
		location.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateEventLocation(location *EventLocationDB) db.DatabaseError {
	query := "UPDATE event_locations SET location_name = $1, location_description = $2, longitude = $3, latitude = $4, updated_at = CURRENT_TIMESTAMP WHERE event_location_id = $5"
	_, err := db.Query[EventLocationDB](query, location.LocationName, location.LocationDescription, location.Longitude, location.Latitude, location.EventLocationId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEventLocation(locationId int64) db.DatabaseError {
	query := "DELETE FROM event_locations WHERE event_location_id = $1"
	_, err := db.Query[EventLocationDB](query, locationId)
	if err != nil {
		return err
	}
	return nil
}
