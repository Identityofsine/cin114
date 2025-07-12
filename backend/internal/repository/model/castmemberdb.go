package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

type CastMemberDB struct {
	CastMemberId int64     `json:"cast_member_id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetAllCastMembers() ([]CastMemberDB, db.DatabaseError) {
	query := "SELECT * FROM cast_members ORDER BY name ASC"
	rows, err := db.Query[CastMemberDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetCastMemberById(castMemberId int64) (*CastMemberDB, db.DatabaseError) {
	query := "SELECT * FROM cast_members WHERE cast_member_id = $1"
	rows, err := db.Query[CastMemberDB](query, castMemberId)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetCastMemberById", "Cast member not found", "cast-member-not-found", 404)
	}
	return &(*rows)[0], nil
}

func GetCastMemberByName(name string) (*CastMemberDB, db.DatabaseError) {
	query := "SELECT * FROM cast_members WHERE name = $1"
	rows, err := db.Query[CastMemberDB](query, name)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetCastMemberByName", "Cast member not found", "cast-member-not-found", 404)
	}
	return &(*rows)[0], nil
}

func CreateCastMember(castMember *CastMemberDB) db.DatabaseError {
	query := "INSERT INTO cast_members (name) VALUES ($1) RETURNING cast_member_id, created_at, updated_at"
	rows, err := db.Query[CastMemberDB](query, castMember.Name)
	if err != nil {
		return err
	}
	if len(*rows) > 0 {
		castMember.CastMemberId = (*rows)[0].CastMemberId
		castMember.CreatedAt = (*rows)[0].CreatedAt
		castMember.UpdatedAt = (*rows)[0].UpdatedAt
	}
	return nil
}

func UpdateCastMember(castMember *CastMemberDB) db.DatabaseError {
	query := "UPDATE cast_members SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE cast_member_id = $2"
	_, err := db.Query[CastMemberDB](query, castMember.Name, castMember.CastMemberId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCastMember(castMemberId int64) db.DatabaseError {
	query := "DELETE FROM cast_members WHERE cast_member_id = $1"
	_, err := db.Query[CastMemberDB](query, castMemberId)
	if err != nil {
		return err
	}
	return nil
}

func CastMemberExists(castMemberId int64) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM cast_members WHERE cast_member_id = $1)"
	rows, err := db.Query[bool](query, castMemberId)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}

func CastMemberExistsByName(name string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM cast_members WHERE name = $1)"
	rows, err := db.Query[bool](query, name)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
