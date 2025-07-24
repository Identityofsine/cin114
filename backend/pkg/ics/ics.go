package ics

import (
	"fmt"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/api/dto/event"
)

const (
	ics = "BEGIN:VCALENDAR\nVERSION:2.0\nCALSCALE:GREGORIAN\nPRODID:-//CIN114//%s//EN\nBEGIN:VEVENT\nUID:%d\nDTSTAMP:%s\nDTSTART:%s\nDTEND:%s\nSUMMARY:%s\nDESCRIPTION:%s\nLOCATION:%s\nEND:VEVENT\nEND:VCALENDAR"
)

func GenerateICSForEvent(
	event event.Event,
) ([]byte, error) {

	if event.ShortDescription == nil || event.EventId == 0 || event.CreatedAt.IsZero() || event.ExpirationDate.IsZero() {
		return nil, fmt.Errorf("invalid event data")
	}

	location := "No location specified"
	if event.Locations != nil && len(event.Locations) > 0 {
		location = *event.Locations[0].LocationAddress // Assuming the first location is the primary one
	}

	// Load Eastern Time zone to convert database times to UTC
	easternTZ, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, fmt.Errorf("failed to load timezone: %v", err)
	}

	// The database time is stored as Eastern Time, so we need to tell Go that
	// and then convert it to UTC for the ICS file
	eventStartEastern := time.Date(
		event.ExpirationDate.Year(),
		event.ExpirationDate.Month(),
		event.ExpirationDate.Day(),
		event.ExpirationDate.Hour(),
		event.ExpirationDate.Minute(),
		event.ExpirationDate.Second(),
		event.ExpirationDate.Nanosecond(),
		easternTZ,
	)
	eventEndEastern := eventStartEastern.Add(time.Hour * 2) // Assuming event lasts 2 hours

	// Convert to UTC for universal calendar compatibility
	eventStartUTC := eventStartEastern.UTC()
	eventEndUTC := eventEndEastern.UTC()

	fileContent := fmt.Sprintf(
		ics,
		*event.ShortDescription,
		event.EventId,
		event.CreatedAt.UTC().Format("20060102T150405Z"), // DTSTAMP in UTC
		eventStartUTC.Format("20060102T150405Z"),         // Event start in UTC
		eventEndUTC.Format("20060102T150405Z"),           // Event end in UTC
		*event.ShortDescription,
		event.Description,
		location,
	)

	return []byte(fileContent), nil
}
