package video

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

// MapBoxart maps VideoDB to Boxart model
func MapBoxart(object VideoDB) Boxart {
	return Boxart{
		VideoId:            object.VideoId,
		UseBoxartAsPreview: object.UseBoxartAsPreview,
		BoxartTitle:        object.BoxartTitle,
		BoxartCaption:      object.BoxartCaption,
		BoxartImg:          object.BoxartImg,
		BoxartVideo:        object.BoxartVideo,
	}
}

// ReverseMapBoxart maps Boxart model to VideoDB (partial mapping)
func ReverseMapBoxart(object Boxart) VideoDB {
	return VideoDB{
		VideoId:            object.VideoId,
		UseBoxartAsPreview: object.UseBoxartAsPreview,
		BoxartTitle:        object.BoxartTitle,
		BoxartCaption:      object.BoxartCaption,
		BoxartImg:          object.BoxartImg,
		BoxartVideo:        object.BoxartVideo,
	}
}
