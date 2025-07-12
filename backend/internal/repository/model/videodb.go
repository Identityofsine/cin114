package model

import (
	"database/sql"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type VideoDB struct {
	VideoId            int64          `json:"video_id"`
	Title              string         `json:"title"`
	Description        *string        `json:"description"`
	Weight             int            `json:"weight"`
	UseBoxartAsPreview bool           `json:"use_boxart_as_preview"`
	BoxartTitle        *string        `json:"boxart_title"`
	BoxartCaption      *string        `json:"boxart_caption"`
	BoxartImg          *string        `json:"boxart_img"`
	BoxartVideo        *string        `json:"boxart_video"`
	Url                string         `json:"url"`
	Date               *string        `json:"date"`
	Img                *string        `json:"img"`
	StyleJson          sql.NullString `json:"style_json"`
	VideoType          string         `json:"video_type"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	Previewable        bool           `json:"previewable"`
}

func GetAllVideos() ([]VideoDB, db.DatabaseError) {
	query := "SELECT * FROM videos ORDER BY weight ASC, created_at DESC"
	rows, err := db.Query[VideoDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetVideoById(videoId int64) (*VideoDB, db.DatabaseError) {
	query := "SELECT * FROM videos WHERE video_id = $1"
	rows, err := db.Query[VideoDB](query, videoId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetVideoById", "Video not found", "video-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetVideosByType(videoType string) ([]VideoDB, db.DatabaseError) {
	query := "SELECT * FROM videos WHERE video_type = $1 ORDER BY weight ASC, created_at DESC"
	rows, err := db.Query[VideoDB](query, videoType)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetVideosByUrl(url string) (*VideoDB, db.DatabaseError) {
	query := "SELECT * FROM videos WHERE url = $1"
	rows, err := db.Query[VideoDB](query, url)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetVideosByUrl", "Video not found", "video-not-found", 404)
	}
	return &(*rows)[0], nil
}

func CreateVideo(video *VideoDB) db.DatabaseError {
	query := `INSERT INTO videos (
		title, description, weight, use_boxart_as_preview, boxart_title, 
		boxart_caption, boxart_img, boxart_video, url, date, img, 
		style_json, video_type
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
	RETURNING video_id, created_at, updated_at`

	rows, err := db.Query[VideoDB](query,
		video.Title, video.Description, video.Weight, video.UseBoxartAsPreview,
		video.BoxartTitle, video.BoxartCaption, video.BoxartImg, video.BoxartVideo,
		video.Url, video.Date, video.Img, video.StyleJson, video.VideoType)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		video.VideoId = (*rows)[0].VideoId
		video.CreatedAt = (*rows)[0].CreatedAt
		video.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateVideo(video *VideoDB) db.DatabaseError {
	query := `UPDATE videos SET 
		title = $1, description = $2, weight = $3, use_boxart_as_preview = $4,
		boxart_title = $5, boxart_caption = $6, boxart_img = $7, boxart_video = $8,
		url = $9, date = $10, img = $11, style_json = $12, video_type = $13,
		updated_at = CURRENT_TIMESTAMP 
		WHERE video_id = $14`

	_, err := db.Query[VideoDB](query,
		video.Title, video.Description, video.Weight, video.UseBoxartAsPreview,
		video.BoxartTitle, video.BoxartCaption, video.BoxartImg, video.BoxartVideo,
		video.Url, video.Date, video.Img, video.StyleJson, video.VideoType, video.VideoId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideo(videoId int64) db.DatabaseError {
	query := "DELETE FROM videos WHERE video_id = $1"
	_, err := db.Query[VideoDB](query, videoId)
	if err != nil {
		return err
	}
	return nil
}

func VideoExists(videoId int64) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM videos WHERE video_id = $1)"
	rows, err := db.Query[bool](query, videoId)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}

func GetVideoByUrl(url string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM videos WHERE url = $1)"
	rows, err := db.Query[bool](query, url)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
