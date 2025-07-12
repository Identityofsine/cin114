package video

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

// MapCast maps CastMemberDB to Cast model
func MapCast(object CastMemberDB) Cast {
	return Cast{
		CastMemberId: object.CastMemberId,
		Name:         object.Name,
		CreatedAt:    object.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    object.UpdatedAt.Format(time.RFC3339),
	}
}

// MapAllCast maps a slice of CastMemberDB to Cast models
func MapAllCast(objects []CastMemberDB) []Cast {
	casts := make([]Cast, len(objects))
	for i, object := range objects {
		casts[i] = MapCast(object)
	}
	return casts
}

// ReverseMapCast maps Cast model to CastMemberDB
func ReverseMapCast(object Cast) CastMemberDB {
	createdAt, _ := time.Parse(time.RFC3339, object.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, object.UpdatedAt)

	return CastMemberDB{
		CastMemberId: object.CastMemberId,
		Name:         object.Name,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
