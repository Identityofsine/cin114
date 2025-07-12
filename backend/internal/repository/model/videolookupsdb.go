package model

import (
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

// Video Type Lookup
type VideoTypeLkDB struct {
	VideoTypeLk          string    `json:"video_type_lk"`
	VideoTypeDescription string    `json:"video_type_description"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// Credit Role Lookup
type CreditRoleLkDB struct {
	CreditRoleLk          string    `json:"credit_role_lk"`
	CreditRoleDescription string    `json:"credit_role_description"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// Link Type Lookup
type LinkTypeLkDB struct {
	LinkTypeLk          string    `json:"link_type_lk"`
	LinkTypeDescription string    `json:"link_type_description"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// Video Type Lookup Functions
func GetAllVideoTypes() ([]VideoTypeLkDB, db.DatabaseError) {
	query := "SELECT * FROM video_type_lks ORDER BY video_type_lk ASC"
	rows, err := db.Query[VideoTypeLkDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetVideoTypeByKey(videoTypeLk string) (*VideoTypeLkDB, db.DatabaseError) {
	query := "SELECT * FROM video_type_lks WHERE video_type_lk = $1"
	rows, err := db.Query[VideoTypeLkDB](query, videoTypeLk)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetVideoTypeByKey", "Video type not found", "video-type-not-found", 404)
	}
	return &(*rows)[0], nil
}

func CreateVideoType(videoType *VideoTypeLkDB) db.DatabaseError {
	query := "INSERT INTO video_type_lks (video_type_lk, video_type_description) VALUES ($1, $2)"
	_, err := db.Query[VideoTypeLkDB](query, videoType.VideoTypeLk, videoType.VideoTypeDescription)
	if err != nil {
		return err
	}
	return nil
}

func UpdateVideoType(videoType *VideoTypeLkDB) db.DatabaseError {
	query := "UPDATE video_type_lks SET video_type_description = $1, updated_at = CURRENT_TIMESTAMP WHERE video_type_lk = $2"
	_, err := db.Query[VideoTypeLkDB](query, videoType.VideoTypeDescription, videoType.VideoTypeLk)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoType(videoTypeLk string) db.DatabaseError {
	query := "DELETE FROM video_type_lks WHERE video_type_lk = $1"
	_, err := db.Query[VideoTypeLkDB](query, videoTypeLk)
	if err != nil {
		return err
	}
	return nil
}

func VideoTypeExists(videoTypeLk string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM video_type_lks WHERE video_type_lk = $1)"
	rows, err := db.Query[bool](query, videoTypeLk)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}

// Credit Role Lookup Functions
func GetAllCreditRoles() ([]CreditRoleLkDB, db.DatabaseError) {
	query := "SELECT * FROM credit_role_lks ORDER BY credit_role_lk ASC"
	rows, err := db.Query[CreditRoleLkDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetCreditRoleByKey(creditRoleLk string) (*CreditRoleLkDB, db.DatabaseError) {
	query := "SELECT * FROM credit_role_lks WHERE credit_role_lk = $1"
	rows, err := db.Query[CreditRoleLkDB](query, creditRoleLk)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetCreditRoleByKey", "Credit role not found", "credit-role-not-found", 404)
	}
	return &(*rows)[0], nil
}

func CreateCreditRole(creditRole *CreditRoleLkDB) db.DatabaseError {
	query := "INSERT INTO credit_role_lks (credit_role_lk, credit_role_description) VALUES ($1, $2)"
	_, err := db.Query[CreditRoleLkDB](query, creditRole.CreditRoleLk, creditRole.CreditRoleDescription)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCreditRole(creditRole *CreditRoleLkDB) db.DatabaseError {
	query := "UPDATE credit_role_lks SET credit_role_description = $1, updated_at = CURRENT_TIMESTAMP WHERE credit_role_lk = $2"
	_, err := db.Query[CreditRoleLkDB](query, creditRole.CreditRoleDescription, creditRole.CreditRoleLk)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCreditRole(creditRoleLk string) db.DatabaseError {
	query := "DELETE FROM credit_role_lks WHERE credit_role_lk = $1"
	_, err := db.Query[CreditRoleLkDB](query, creditRoleLk)
	if err != nil {
		return err
	}
	return nil
}

func CreditRoleExists(creditRoleLk string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM credit_role_lks WHERE credit_role_lk = $1)"
	rows, err := db.Query[bool](query, creditRoleLk)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}

// Link Type Lookup Functions
func GetAllLinkTypes() ([]LinkTypeLkDB, db.DatabaseError) {
	query := "SELECT * FROM link_type_lks ORDER BY link_type_lk ASC"
	rows, err := db.Query[LinkTypeLkDB](query)
	if err != nil {
		return nil, err
	}
	return *rows, nil
}

func GetLinkTypeByKey(linkTypeLk string) (*LinkTypeLkDB, db.DatabaseError) {
	query := "SELECT * FROM link_type_lks WHERE link_type_lk = $1"
	rows, err := db.Query[LinkTypeLkDB](query, linkTypeLk)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetLinkTypeByKey", "Link type not found", "link-type-not-found", 404)
	}
	return &(*rows)[0], nil
}

func CreateLinkType(linkType *LinkTypeLkDB) db.DatabaseError {
	query := "INSERT INTO link_type_lks (link_type_lk, link_type_description) VALUES ($1, $2)"
	_, err := db.Query[LinkTypeLkDB](query, linkType.LinkTypeLk, linkType.LinkTypeDescription)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLinkType(linkType *LinkTypeLkDB) db.DatabaseError {
	query := "UPDATE link_type_lks SET link_type_description = $1, updated_at = CURRENT_TIMESTAMP WHERE link_type_lk = $2"
	_, err := db.Query[LinkTypeLkDB](query, linkType.LinkTypeDescription, linkType.LinkTypeLk)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLinkType(linkTypeLk string) db.DatabaseError {
	query := "DELETE FROM link_type_lks WHERE link_type_lk = $1"
	_, err := db.Query[LinkTypeLkDB](query, linkTypeLk)
	if err != nil {
		return err
	}
	return nil
}

func LinkTypeExists(linkTypeLk string) (bool, db.DatabaseError) {
	query := "SELECT EXISTS(SELECT 1 FROM link_type_lks WHERE link_type_lk = $1)"
	rows, err := db.Query[bool](query, linkTypeLk)
	if err != nil {
		return false, err
	}
	if len(*rows) == 0 {
		return false, nil
	}
	return (*rows)[0], nil
}
