package model

type Boxart struct {
	VideoId            int64   `json:"videoId"`
	UseBoxartAsPreview bool    `json:"useBoxartAsPreview"`
	BoxartTitle        *string `json:"boxartTitle"`
	BoxartCaption      *string `json:"boxartCaption"`
	BoxartImg          *string `json:"boxartImg"`
	BoxartVideo        *string `json:"boxartVideo"`
}
