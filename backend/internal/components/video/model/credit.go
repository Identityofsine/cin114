package model

type Credit struct {
	CreditId    int64  `json:"creditId"`
	VideoId     int64  `json:"videoId"`
	CastMember  *Cast  `json:"castMember,omitempty"`
	CreditRole  string `json:"creditRole"`
	CreditOrder int    `json:"creditOrder"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
