package model

import "github.com/identityofsine/fofx-go-gin-api-template/pkg/db"

type UserEmailWhitelistDb struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// logic
func CreateUserEmailWhitelist(email, description string) db.DatabaseError {
	query := "INSERT INTO user_email_whitelist (email, description) VALUES ($1, $2)"

	_, err := db.Query[UserEmailWhitelistDb](query, email, description)
	if err != nil {
		return err
	}
	return nil
}

func GetUserEmailWhitelistByEmail(email string) (*UserEmailWhitelistDb, db.DatabaseError) {
	query := "SELECT * FROM user_email_whitelist WHERE email = $1"
	rows, err := db.Query[UserEmailWhitelistDb](query, email)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, db.NewDatabaseError("GetUserEmailWhitelistByEmail", "Email not found in whitelist", "email-not-found", 404)
	}
	return &(*rows)[0], nil
}
