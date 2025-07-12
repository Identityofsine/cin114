package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type VideoLinksDB struct {
	VideoLinkId int64     `json:"video_link_id"`
	VideoId     int64     `json:"video_id"`
	LinkType    string    `json:"link_type"`
	LinkUrl     string    `json:"link_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Extended struct for joins with video info
type VideoLinksWithDetailsDB struct {
	VideoLinksDB
	VideoTitle string `json:"video_title"`
	VideoType  string `json:"video_type"`
}

func GetLinksByVideoId(videoId int64) ([]VideoLinksDB, db.DatabaseError) {
	query := "SELECT * FROM video_links WHERE video_id = $1 ORDER BY link_type ASC"
	rows, err := db.Query[VideoLinksDB](query, videoId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetLinksByVideoIdWithDetails(videoId int64) ([]VideoLinksWithDetailsDB, db.DatabaseError) {
	query := `SELECT 
		vl.*, v.title as video_title, v.video_type
		FROM video_links vl
		JOIN videos v ON vl.video_id = v.video_id
		WHERE vl.video_id = $1 
		ORDER BY vl.link_type ASC`

	rows, err := db.Query[VideoLinksWithDetailsDB](query, videoId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetLinksByType(linkType string) ([]VideoLinksDB, db.DatabaseError) {
	query := "SELECT * FROM video_links WHERE link_type = $1 ORDER BY created_at DESC"
	rows, err := db.Query[VideoLinksDB](query, linkType)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetLinksByTypeWithDetails(linkType string) ([]VideoLinksWithDetailsDB, db.DatabaseError) {
	query := `SELECT 
		vl.*, v.title as video_title, v.video_type
		FROM video_links vl
		JOIN videos v ON vl.video_id = v.video_id
		WHERE vl.link_type = $1 
		ORDER BY vl.created_at DESC`

	rows, err := db.Query[VideoLinksWithDetailsDB](query, linkType)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateVideoLink(link *VideoLinksDB) db.DatabaseError {
	query := "INSERT INTO video_links (video_id, link_type, link_url) VALUES ($1, $2, $3) RETURNING video_link_id, created_at, updated_at"
	rows, err := db.Query[VideoLinksDB](query, link.VideoId, link.LinkType, link.LinkUrl)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		link.VideoLinkId = (*rows)[0].VideoLinkId
		link.CreatedAt = (*rows)[0].CreatedAt
		link.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateVideoLink(link *VideoLinksDB) db.DatabaseError {
	query := "UPDATE video_links SET link_type = $1, link_url = $2, updated_at = CURRENT_TIMESTAMP WHERE video_link_id = $3"
	_, err := db.Query[VideoLinksDB](query, link.LinkType, link.LinkUrl, link.VideoLinkId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoLink(videoLinkId int64) db.DatabaseError {
	query := "DELETE FROM video_links WHERE video_link_id = $1"
	_, err := db.Query[VideoLinksDB](query, videoLinkId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLinksByVideoId(videoId int64) db.DatabaseError {
	query := "DELETE FROM video_links WHERE video_id = $1"
	_, err := db.Query[VideoLinksDB](query, videoId)
	if err != nil {
		return err
	}
	return nil
}

func VideoLinkExists(videoLinkId int64) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM video_links WHERE video_link_id = $1)"
	rows, err := db.Query[bool](query, videoLinkId)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
