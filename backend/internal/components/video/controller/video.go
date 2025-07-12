package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/service"
)

// /video/:videoId
func GetVideoById(c *gin.Context) {
	// Get the video ID from the request context
	videoId := c.Param("videoId")
	if videoId == "" {
		c.JSON(400, gin.H{"error": "videoId is required"})
		return
	}

	// Call the service to get the video details
	videoDetails, err := service.GetVideoByIdService(videoId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve video details"})
		return
	}

	c.JSON(200, videoDetails)
}
