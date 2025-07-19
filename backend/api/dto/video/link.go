package video

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

// MapLink maps VideoLinksDB to Link model
func MapLink(object VideoLinksDB) Link {
	return Link{
		VideoLinkId: object.VideoLinkId,
		VideoId:     object.VideoId,
		LinkType:    object.LinkType,
		LinkUrl:     object.LinkUrl,
		CreatedAt:   object.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   object.UpdatedAt.Format(time.RFC3339),
	}
}

// MapAllLinks maps a slice of VideoLinksDB to Link models
func MapAllLinks(objects []VideoLinksDB) []Link {
	links := make([]Link, len(objects))
	for i, object := range objects {
		links[i] = MapLink(object)
	}
	return links
}

// MapLinkWithDetails maps VideoLinksWithDetailsDB to Link model
func MapLinkWithDetails(object VideoLinksWithDetailsDB) Link {
	return Link{
		VideoLinkId: object.VideoLinkId,
		VideoId:     object.VideoId,
		LinkType:    object.LinkType,
		LinkUrl:     object.LinkUrl,
		CreatedAt:   object.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   object.UpdatedAt.Format(time.RFC3339),
	}
}

// MapAllLinksWithDetails maps a slice of VideoLinksWithDetailsDB to Link models
func MapAllLinksWithDetails(objects []VideoLinksWithDetailsDB) []Link {
	links := make([]Link, len(objects))
	for i, object := range objects {
		links[i] = MapLinkWithDetails(object)
	}
	return links
}

// ReverseMapLink maps Link model to VideoLinksDB
func ReverseMapLink(object Link) VideoLinksDB {
	createdAt, _ := time.Parse(time.RFC3339, object.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, object.UpdatedAt)

	return VideoLinksDB{
		VideoLinkId: object.VideoLinkId,
		VideoId:     object.VideoId,
		LinkType:    object.LinkType,
		LinkUrl:     object.LinkUrl,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
