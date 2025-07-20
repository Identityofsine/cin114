package model

type Link struct {
	VideoLinkId int64  `json:"videoLinkId"`
	VideoId     int64  `json:"videoId"`
	LinkType    string `json:"linkType"`
	LinkUrl     string `json:"linkUrl"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
