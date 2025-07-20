package video

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

// Map maps VideoDB to Video model
func Map(object VideoDB) Video {
	return Video{
		VideoId:     object.VideoId,
		Title:       object.Title,
		Description: object.Description,
		Weight:      object.Weight,
		VideoType:   object.VideoType,
		CreatedAt:   object.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   object.UpdatedAt.Format(time.RFC3339),
	}
}

// MapAll maps a slice of VideoDB to Video models
func MapAll(objects []VideoDB) []Video {
	videos := make([]Video, len(objects))
	for i, object := range objects {
		videos[i] = Map(object)
	}
	return videos
}

// ReverseMap maps Video model to VideoDB
func ReverseMap(object Video) VideoDB {
	createdAt, _ := time.Parse(time.RFC3339, object.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, object.UpdatedAt)

	return VideoDB{
		VideoId:     object.VideoId,
		Title:       object.Title,
		Description: object.Description,
		Weight:      object.Weight,
		VideoType:   object.VideoType,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
