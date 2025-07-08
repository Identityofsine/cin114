package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type EventImageLkDB struct {
	EventImageLk          string    `json:"event_image_lk"`
	EventImageDescription string    `json:"event_image_description"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type EventImageDB struct {
	EventImageId int64     `json:"event_image_id"`
	EventId      int64     `json:"event_id"`
	ImageUrl     string    `json:"image_url"`
	ImageType    string    `json:"image_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Event Image Lookup functions
func GetAllEventImageLks() ([]EventImageLkDB, db.DatabaseError) {
	query := "SELECT * FROM event_image_lks ORDER BY event_image_lk"
	rows, err := db.Query[EventImageLkDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetEventImageLkById(imageType string) (*EventImageLkDB, db.DatabaseError) {
	query := "SELECT * FROM event_image_lks WHERE event_image_lk = $1"
	rows, err := db.Query[EventImageLkDB](query, imageType)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetEventImageLkById", "Image type not found", "image-type-not-found", 404)
	}
	return &(*rows)[0], nil
}

// Event Image functions
func GetImagesByEventId(eventId int64) ([]EventImageDB, db.DatabaseError) {
	query := "SELECT * FROM event_images WHERE event_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query[EventImageDB](query, eventId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetEventImageById(imageId int64) (*EventImageDB, db.DatabaseError) {
	query := "SELECT * FROM event_images WHERE event_image_id = $1"
	rows, err := db.Query[EventImageDB](query, imageId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetEventImageById", "Event image not found", "event-image-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetImagesByEventIdAndType(eventId int64, imageType string) ([]EventImageDB, db.DatabaseError) {
	query := "SELECT * FROM event_images WHERE event_id = $1 AND image_type = $2 ORDER BY created_at DESC"
	rows, err := db.Query[EventImageDB](query, eventId, imageType)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateEventImage(image *EventImageDB) db.DatabaseError {
	query := "INSERT INTO event_images (event_id, image_url, image_type) VALUES ($1, $2, $3) RETURNING event_image_id, created_at, updated_at"
	rows, err := db.Query[EventImageDB](query, image.EventId, image.ImageUrl, image.ImageType)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		image.EventImageId = (*rows)[0].EventImageId
		image.CreatedAt = (*rows)[0].CreatedAt
		image.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateEventImage(image *EventImageDB) db.DatabaseError {
	query := "UPDATE event_images SET image_url = $1, image_type = $2, updated_at = CURRENT_TIMESTAMP WHERE event_image_id = $3"
	_, err := db.Query[EventImageDB](query, image.ImageUrl, image.ImageType, image.EventImageId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEventImage(imageId int64) db.DatabaseError {
	query := "DELETE FROM event_images WHERE event_image_id = $1"
	_, err := db.Query[EventImageDB](query, imageId)
	if err != nil {
		return err
	}
	return nil
}
