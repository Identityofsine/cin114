package model

type UserWhitelist struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
