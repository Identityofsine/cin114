package video

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

// MapVideoWithDetails maps VideoDB with related data to complete Video model
func MapVideoWithDetails(videoDB VideoDB, boxart *Boxart, credits []Credit, links []Link) Video {
	video := Map(videoDB)

	if boxart != nil {
		video.Boxart = boxart
	}

	if len(credits) > 0 {
		video.Credits = credits
	}

	if len(links) > 0 {
		video.Links = links
	}

	return video
}

// MapVideoWithBoxart maps VideoDB with boxart data to Video model
func MapVideoWithBoxart(videoDB VideoDB, boxartDB *VideoDB) Video {
	video := Map(videoDB)

	if boxartDB != nil {
		boxart := MapBoxart(*boxartDB)
		video.Boxart = &boxart
	}

	return video
}

// MapVideoWithCredits maps VideoDB with credits data to Video model
func MapVideoWithCredits(videoDB VideoDB, creditsDB []VideoCreditsWithDetailsDB) Video {
	video := Map(videoDB)

	if len(creditsDB) > 0 {
		credits := MapAllCreditsWithCast(creditsDB)
		video.Credits = credits
	}

	return video
}

// MapVideoWithLinks maps VideoDB with links data to Video model
func MapVideoWithLinks(videoDB VideoDB, linksDB []VideoLinksWithDetailsDB) Video {
	video := Map(videoDB)

	if len(linksDB) > 0 {
		links := MapAllLinksWithDetails(linksDB)
		video.Links = links
	}

	return video
}

// MapAllVideosWithDetails maps multiple VideoDB with related data to Video models
func MapAllVideosWithDetails(videosDB []VideoDB, boxarts map[int64]*Boxart, credits map[int64][]Credit, links map[int64][]Link) []Video {
	videos := make([]Video, len(videosDB))
	for i, videoDB := range videosDB {
		boxart := boxarts[videoDB.VideoId]
		videoCredits := credits[videoDB.VideoId]
		videoLinks := links[videoDB.VideoId]
		videos[i] = MapVideoWithDetails(videoDB, boxart, videoCredits, videoLinks)
	}
	return videos
}
