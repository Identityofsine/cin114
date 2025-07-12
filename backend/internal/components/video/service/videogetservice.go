package service

import (
	"strconv"

	videoDto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/video"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
)

// GetAllVideosService returns all videos with basic information
func GetAllVideosService() ([]Video, db.DatabaseError) {
	videosDB, err := GetAllVideos()
	if err != nil {
		return nil, err
	}

	return videoDto.MapAll(videosDB), nil
}

// GetVideoByIdService returns a specific video by ID with basic information
func GetVideoByIdService(id string) (*Video, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetVideoById", "Invalid ID format", "invalid-id", 400)
	}

	videoDB, err := GetVideoById(idInt)
	if err != nil {
		return nil, err
	}

	video := videoDto.Map(*videoDB)
	return &video, nil
}

// GetVideosByTypeService returns all videos of a specific type
func GetVideosByTypeService(videoType string) ([]Video, db.DatabaseError) {
	videosDB, err := GetVideosByType(videoType)
	if err != nil {
		return nil, err
	}

	return videoDto.MapAll(videosDB), nil
}

// GetVideoByUrlService returns a specific video by URL
func GetVideoByUrlService(url string) (*Video, db.DatabaseError) {
	videoDB, err := GetVideosByUrl(url)
	if err != nil {
		return nil, err
	}

	video := videoDto.Map(*videoDB)
	return &video, nil
}

// GetVideoWithBoxartService returns a video with its boxart information
func GetVideoWithBoxartService(id string) (*Video, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetVideoWithBoxart", "Invalid ID format", "invalid-id", 400)
	}

	videoDB, err := GetVideoById(idInt)
	if err != nil {
		return nil, err
	}

	video := videoDto.MapVideoWithBoxart(*videoDB, videoDB)
	return &video, nil
}

// GetVideoWithCreditsService returns a video with its credits and cast information
func GetVideoWithCreditsService(id string) (*Video, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetVideoWithCredits", "Invalid ID format", "invalid-id", 400)
	}

	videoDB, err := GetVideoById(idInt)
	if err != nil {
		return nil, err
	}

	creditsDB, err := GetCreditsByVideoIdWithDetails(idInt)
	if err != nil {
		return nil, err
	}

	video := videoDto.MapVideoWithCredits(*videoDB, creditsDB)
	return &video, nil
}

// GetVideoWithLinksService returns a video with its links information
func GetVideoWithLinksService(id string) (*Video, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetVideoWithLinks", "Invalid ID format", "invalid-id", 400)
	}

	videoDB, err := GetVideoById(idInt)
	if err != nil {
		return nil, err
	}

	linksDB, err := GetLinksByVideoIdWithDetails(idInt)
	if err != nil {
		return nil, err
	}

	video := videoDto.MapVideoWithLinks(*videoDB, linksDB)
	return &video, nil
}

// GetVideoWithAllDetailsService returns a video with all related information (boxart, credits, links)
func GetVideoWithAllDetailsService(id string) (*Video, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetVideoWithAllDetails", "Invalid ID format", "invalid-id", 400)
	}

	videoDB, err := GetVideoById(idInt)
	if err != nil {
		return nil, err
	}

	// Get boxart
	boxart := videoDto.MapBoxart(*videoDB)

	// Get credits with cast details
	creditsDB, err := GetCreditsByVideoIdWithDetails(idInt)
	if err != nil {
		return nil, err
	}
	credits := videoDto.MapAllCreditsWithCast(creditsDB)

	// Get links
	linksDB, err := GetLinksByVideoIdWithDetails(idInt)
	if err != nil {
		return nil, err
	}
	links := videoDto.MapAllLinksWithDetails(linksDB)

	video := videoDto.MapVideoWithDetails(*videoDB, &boxart, credits, links)
	return &video, nil
}

// GetAllVideosWithAllDetailsService returns all videos with all related information
func GetAllVideosWithAllDetailsService() ([]Video, db.DatabaseError) {
	videosDB, err := GetAllVideos()
	if err != nil {
		return nil, err
	}

	// Create maps to store related data
	boxarts := make(map[int64]*Boxart)
	credits := make(map[int64][]Credit)
	links := make(map[int64][]Link)

	// Process each video to get related data
	for _, videoDB := range videosDB {
		// Get boxart
		boxart := videoDto.MapBoxart(videoDB)
		boxarts[videoDB.VideoId] = &boxart

		// Get credits
		creditsDB, err := GetCreditsByVideoIdWithDetails(videoDB.VideoId)
		if err == nil {
			credits[videoDB.VideoId] = videoDto.MapAllCreditsWithCast(creditsDB)
		}

		// Get links
		linksDB, err := GetLinksByVideoIdWithDetails(videoDB.VideoId)
		if err == nil {
			links[videoDB.VideoId] = videoDto.MapAllLinksWithDetails(linksDB)
		}
	}

	return videoDto.MapAllVideosWithDetails(videosDB, boxarts, credits, links), nil
}

// GetCastMembersService returns all cast members
func GetCastMembersService() ([]Cast, db.DatabaseError) {
	castMembersDB, err := GetAllCastMembers()
	if err != nil {
		return nil, err
	}

	return videoDto.MapAllCast(castMembersDB), nil
}

// GetCastMemberByIdService returns a specific cast member by ID
func GetCastMemberByIdService(id string) (*Cast, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetCastMemberById", "Invalid ID format", "invalid-id", 400)
	}

	castMemberDB, err := GetCastMemberById(idInt)
	if err != nil {
		return nil, err
	}

	cast := videoDto.MapCast(*castMemberDB)
	return &cast, nil
}

// GetCastMemberByNameService returns a specific cast member by name
func GetCastMemberByNameService(name string) (*Cast, db.DatabaseError) {
	castMemberDB, err := GetCastMemberByName(name)
	if err != nil {
		return nil, err
	}

	cast := videoDto.MapCast(*castMemberDB)
	return &cast, nil
}

// GetCreditsByVideoIdService returns all credits for a specific video
func GetCreditsByVideoIdService(videoId string) ([]Credit, db.DatabaseError) {
	videoIdInt, parseErr := strconv.ParseInt(videoId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetCreditsByVideoId", "Invalid video ID format", "invalid-video-id", 400)
	}

	creditsDB, err := GetCreditsByVideoIdWithDetails(videoIdInt)
	if err != nil {
		return nil, err
	}

	return videoDto.MapAllCreditsWithCast(creditsDB), nil
}

// GetCreditsByCastMemberIdService returns all credits for a specific cast member
func GetCreditsByCastMemberIdService(castMemberId string) ([]Credit, db.DatabaseError) {
	castMemberIdInt, parseErr := strconv.ParseInt(castMemberId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetCreditsByCastMemberId", "Invalid cast member ID format", "invalid-cast-member-id", 400)
	}

	creditsDB, err := GetCreditsByCastMemberIdWithDetails(castMemberIdInt)
	if err != nil {
		return nil, err
	}

	return videoDto.MapAllCreditsWithCast(creditsDB), nil
}

// GetLinksByVideoIdService returns all links for a specific video
func GetLinksByVideoIdService(videoId string) ([]Link, db.DatabaseError) {
	videoIdInt, parseErr := strconv.ParseInt(videoId, 10, 64)
	if parseErr != nil {
		return nil, db.NewDatabaseError("GetLinksByVideoId", "Invalid video ID format", "invalid-video-id", 400)
	}

	linksDB, err := GetLinksByVideoIdWithDetails(videoIdInt)
	if err != nil {
		return nil, err
	}

	return videoDto.MapAllLinksWithDetails(linksDB), nil
}

// GetLinksByTypeService returns all links of a specific type
func GetLinksByTypeService(linkType string) ([]Link, db.DatabaseError) {
	linksDB, err := GetLinksByTypeWithDetails(linkType)
	if err != nil {
		return nil, err
	}

	return videoDto.MapAllLinksWithDetails(linksDB), nil
}

// VideoExistsService checks if a video exists
func VideoExistsService(id string) (bool, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return false, db.NewDatabaseError("VideoExists", "Invalid ID format", "invalid-id", 400)
	}

	return VideoExists(idInt)
}

// CastMemberExistsService checks if a cast member exists
func CastMemberExistsService(id string) (bool, db.DatabaseError) {
	idInt, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		return false, db.NewDatabaseError("CastMemberExists", "Invalid ID format", "invalid-id", 400)
	}

	return CastMemberExists(idInt)
}
