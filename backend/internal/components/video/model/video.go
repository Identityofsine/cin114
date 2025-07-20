package model

type Video struct {
	VideoId     int64    `json:"videoId"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Boxart      *Boxart  `json:"boxart,omitempty"`
	Weight      int      `json:"weight"`
	Credits     []Credit `json:"credits,omitempty"`
	Links       []Link   `json:"links,omitempty"`
	VideoType   string   `json:"videoType"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}
