package video

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
)

// MapCredit maps VideoCreditsDB to Credit model
func MapCredit(object VideoCreditsDB) Credit {
	return Credit{
		CreditId:    object.VideoCreditId,
		VideoId:     object.VideoId,
		CastMember:  nil, // Cast member not included in basic mapping
		CreditRole:  object.CreditRole,
		CreditOrder: object.CreditOrder,
		CreatedAt:   object.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   object.UpdatedAt.Format(time.RFC3339),
	}
}

// MapAllCredits maps a slice of VideoCreditsDB to Credit models
func MapAllCredits(objects []VideoCreditsDB) []Credit {
	credits := make([]Credit, len(objects))
	for i, object := range objects {
		credits[i] = MapCredit(object)
	}
	return credits
}

// MapCreditWithCast maps VideoCreditsWithDetailsDB to Credit model with Cast member
func MapCreditWithCast(object VideoCreditsWithDetailsDB) Credit {
	cast := &Cast{
		CastMemberId: object.CastMemberId,
		Name:         object.CastMemberName,
		CreatedAt:    object.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    object.UpdatedAt.Format(time.RFC3339),
	}

	return Credit{
		CreditId:    object.VideoCreditId,
		VideoId:     object.VideoId,
		CastMember:  cast,
		CreditRole:  object.CreditRole,
		CreditOrder: object.CreditOrder,
		CreatedAt:   object.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   object.UpdatedAt.Format(time.RFC3339),
	}
}

// MapAllCreditsWithCast maps a slice of VideoCreditsWithDetailsDB to Credit models with Cast members
func MapAllCreditsWithCast(objects []VideoCreditsWithDetailsDB) []Credit {
	credits := make([]Credit, len(objects))
	for i, object := range objects {
		credits[i] = MapCreditWithCast(object)
	}
	return credits
}

// ReverseMapCredit maps Credit model to VideoCreditsDB
func ReverseMapCredit(object Credit) VideoCreditsDB {
	createdAt, _ := time.Parse(time.RFC3339, object.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, object.UpdatedAt)

	castMemberId := int64(0)
	if object.CastMember != nil {
		castMemberId = object.CastMember.CastMemberId
	}

	return VideoCreditsDB{
		VideoCreditId: object.CreditId,
		VideoId:       object.VideoId,
		CastMemberId:  castMemberId,
		CreditRole:    object.CreditRole,
		CreditOrder:   object.CreditOrder,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}
