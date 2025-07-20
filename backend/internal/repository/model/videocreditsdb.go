package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type VideoCreditsDB struct {
	VideoCreditId int64     `json:"video_credit_id"`
	VideoId       int64     `json:"video_id"`
	CastMemberId  int64     `json:"cast_member_id"`
	CreditRole    string    `json:"credit_role"`
	CreditOrder   int       `json:"credit_order"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Extended struct for joins with cast member and video info
type VideoCreditsWithDetailsDB struct {
	VideoCreditsDB
	CastMemberName string `json:"cast_member_name"`
	VideoTitle     string `json:"video_title"`
	VideoType      string `json:"video_type"`
}

func GetCreditsByVideoId(videoId int64) ([]VideoCreditsDB, db.DatabaseError) {
	query := "SELECT * FROM video_credits WHERE video_id = $1 ORDER BY credit_order ASC, credit_role ASC"
	rows, err := db.Query[VideoCreditsDB](query, videoId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetCreditsByCastMemberId(castMemberId int64) ([]VideoCreditsDB, db.DatabaseError) {
	query := "SELECT * FROM video_credits WHERE cast_member_id = $1 ORDER BY credit_order ASC, credit_role ASC"
	rows, err := db.Query[VideoCreditsDB](query, castMemberId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetCreditsByVideoIdWithDetails(videoId int64) ([]VideoCreditsWithDetailsDB, db.DatabaseError) {
	query := `SELECT 
		vc.*, cm.name as cast_member_name, v.title as video_title, v.video_type
		FROM video_credits vc
		JOIN cast_members cm ON vc.cast_member_id = cm.cast_member_id
		JOIN videos v ON vc.video_id = v.video_id
		WHERE vc.video_id = $1 
		ORDER BY vc.credit_order ASC, vc.credit_role ASC`

	rows, err := db.Query[VideoCreditsWithDetailsDB](query, videoId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetCreditsByCastMemberIdWithDetails(castMemberId int64) ([]VideoCreditsWithDetailsDB, db.DatabaseError) {
	query := `SELECT 
		vc.*, cm.name as cast_member_name, v.title as video_title, v.video_type
		FROM video_credits vc
		JOIN cast_members cm ON vc.cast_member_id = cm.cast_member_id
		JOIN videos v ON vc.video_id = v.video_id
		WHERE vc.cast_member_id = $1 
		ORDER BY vc.credit_order ASC, vc.credit_role ASC`

	rows, err := db.Query[VideoCreditsWithDetailsDB](query, castMemberId)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetCreditsByRole(creditRole string) ([]VideoCreditsDB, db.DatabaseError) {
	query := "SELECT * FROM video_credits WHERE credit_role = $1 ORDER BY credit_order ASC"
	rows, err := db.Query[VideoCreditsDB](query, creditRole)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func CreateVideoCredit(credit *VideoCreditsDB) db.DatabaseError {
	query := "INSERT INTO video_credits (video_id, cast_member_id, credit_role, credit_order) VALUES ($1, $2, $3, $4) RETURNING video_credit_id, created_at, updated_at"
	rows, err := db.Query[VideoCreditsDB](query, credit.VideoId, credit.CastMemberId, credit.CreditRole, credit.CreditOrder)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		credit.VideoCreditId = (*rows)[0].VideoCreditId
		credit.CreatedAt = (*rows)[0].CreatedAt
		credit.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateVideoCredit(credit *VideoCreditsDB) db.DatabaseError {
	query := "UPDATE video_credits SET credit_role = $1, credit_order = $2, updated_at = CURRENT_TIMESTAMP WHERE video_credit_id = $3"
	_, err := db.Query[VideoCreditsDB](query, credit.CreditRole, credit.CreditOrder, credit.VideoCreditId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoCredit(videoCreditId int64) db.DatabaseError {
	query := "DELETE FROM video_credits WHERE video_credit_id = $1"
	_, err := db.Query[VideoCreditsDB](query, videoCreditId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCreditsByVideoId(videoId int64) db.DatabaseError {
	query := "DELETE FROM video_credits WHERE video_id = $1"
	_, err := db.Query[VideoCreditsDB](query, videoId)
	if err != nil {
		return err
	}
	return nil
}

func VideoCreditExists(videoCreditId int64) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM video_credits WHERE video_credit_id = $1)"
	rows, err := db.Query[bool](query, videoCreditId)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
