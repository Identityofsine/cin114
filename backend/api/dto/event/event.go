package event

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

type EventLocation struct {
	EventLocationId     int64     `json:"event_location_id"`
	EventId             int64     `json:"event_id"`
	LocationName        string    `json:"location_name"`
	LocationDescription *string   `json:"location_description"`
	LocationAddress     *string   `json:"location_address"`
	Longitude           float64   `json:"longitude"`
	Latitude            float64   `json:"latitude"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type EventImage struct {
	EventImageId int64     `json:"event_image_id"`
	EventId      int64     `json:"event_id"`
	ImageUrl     string    `json:"image_url"`
	ImageType    string    `json:"image_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EventImageLk struct {
	EventImageLk          string    `json:"event_image_lk"`
	EventImageDescription string    `json:"event_image_description"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type Event struct {
	EventId          int64           `json:"event_id"`
	VideoId          int64           `json:"video_id,omitempty"`
	Description      string          `json:"description"`
	ShortDescription *string         `json:"short_description"`
	ExpirationDate   *time.Time      `json:"expiration_date"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	Locations        []EventLocation `json:"locations,omitempty"`
	Images           []EventImage    `json:"images,omitempty"`
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

type CreateEventLocationRequest struct {
	EventId             int64   `json:"event_id" binding:"required"`
	LocationName        string  `json:"location_name" binding:"required"`
	LocationDescription *string `json:"location_description"`
	Longitude           float64 `json:"longitude" binding:"required"`
	Latitude            float64 `json:"latitude" binding:"required"`
}

type UpdateEventLocationRequest struct {
	LocationName        string  `json:"location_name" binding:"required"`
	LocationDescription *string `json:"location_description"`
	Longitude           float64 `json:"longitude" binding:"required"`
	Latitude            float64 `json:"latitude" binding:"required"`
}

type CreateEventImageRequest struct {
	EventId   int64  `json:"event_id" binding:"required"`
	ImageUrl  string `json:"image_url" binding:"required"`
	ImageType string `json:"image_type" binding:"required"`
}

type UpdateEventImageRequest struct {
	ImageUrl  string `json:"image_url" binding:"required"`
	ImageType string `json:"image_type" binding:"required"`
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

// Event mapping functions
func Map(object EventDB) Event {
	return Event{
		EventId:          object.EventId,
		VideoId:          object.VideoId,
		Description:      object.Description,
		ShortDescription: object.ShortDescription,
		ExpirationDate:   object.ExpirationDate,
		CreatedAt:        object.CreatedAt,
		UpdatedAt:        object.UpdatedAt,
		Locations:        []EventLocation{},
		Images:           []EventImage{},
	}
}

func MapWithChildren(object EventDB, locations []EventLocationDB, images []EventImageDB) Event {
	event := Map(object)
	event.Locations = MapLocations(locations)
	event.Images = MapImages(images)
	return event
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

// Event Location mapping functions
func MapLocation(object EventLocationDB) EventLocation {
	return EventLocation{
		EventLocationId:     object.EventLocationId,
		EventId:             object.EventId,
		LocationName:        object.LocationName,
		LocationDescription: object.LocationDescription,
		LocationAddress:     object.LocationAddress,
		Longitude:           object.Longitude,
		Latitude:            object.Latitude,
		CreatedAt:           object.CreatedAt,
		UpdatedAt:           object.UpdatedAt,
	}
}

func MapLocations(objects []EventLocationDB) []EventLocation {
	locations := make([]EventLocation, len(objects))
	for i, obj := range objects {
		locations[i] = MapLocation(obj)
	}
	return locations
}

func ReverseMapLocation(object EventLocation) EventLocationDB {
	return EventLocationDB{
		EventLocationId:     object.EventLocationId,
		EventId:             object.EventId,
		LocationName:        object.LocationName,
		LocationDescription: object.LocationDescription,
		LocationAddress:     object.LocationAddress,
		Longitude:           object.Longitude,
		Latitude:            object.Latitude,
		CreatedAt:           object.CreatedAt,
		UpdatedAt:           object.UpdatedAt,
	}
}

func MapCreateLocationRequest(object CreateEventLocationRequest) EventLocationDB {
	return EventLocationDB{
		EventId:             object.EventId,
		LocationName:        object.LocationName,
		LocationDescription: object.LocationDescription,
		Longitude:           object.Longitude,
		Latitude:            object.Latitude,
	}
}

func MapUpdateLocationRequest(object UpdateEventLocationRequest) EventLocationDB {
	return EventLocationDB{
		LocationName:        object.LocationName,
		LocationDescription: object.LocationDescription,
		Longitude:           object.Longitude,
		Latitude:            object.Latitude,
	}
}

// Event Image mapping functions
func MapImage(object EventImageDB) EventImage {
	return EventImage{
		EventImageId: object.EventImageId,
		EventId:      object.EventId,
		ImageUrl:     object.ImageUrl,
		ImageType:    object.ImageType,
		CreatedAt:    object.CreatedAt,
		UpdatedAt:    object.UpdatedAt,
	}
}

func MapImages(objects []EventImageDB) []EventImage {
	images := make([]EventImage, len(objects))
	for i, obj := range objects {
		images[i] = MapImage(obj)
	}
	return images
}

func ReverseMapImage(object EventImage) EventImageDB {
	return EventImageDB{
		EventImageId: object.EventImageId,
		EventId:      object.EventId,
		ImageUrl:     object.ImageUrl,
		ImageType:    object.ImageType,
		CreatedAt:    object.CreatedAt,
		UpdatedAt:    object.UpdatedAt,
	}
}

func MapCreateImageRequest(object CreateEventImageRequest) EventImageDB {
	return EventImageDB{
		EventId:   object.EventId,
		ImageUrl:  object.ImageUrl,
		ImageType: object.ImageType,
	}
}

func MapUpdateImageRequest(object UpdateEventImageRequest) EventImageDB {
	return EventImageDB{
		ImageUrl:  object.ImageUrl,
		ImageType: object.ImageType,
	}
}

// Event Image Lookup mapping functions
func MapImageLk(object EventImageLkDB) EventImageLk {
	return EventImageLk{
		EventImageLk:          object.EventImageLk,
		EventImageDescription: object.EventImageDescription,
		CreatedAt:             object.CreatedAt,
		UpdatedAt:             object.UpdatedAt,
	}
}

func MapImageLks(objects []EventImageLkDB) []EventImageLk {
	imageLks := make([]EventImageLk, len(objects))
	for i, obj := range objects {
		imageLks[i] = MapImageLk(obj)
	}
	return imageLks
}
