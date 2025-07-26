package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/event"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/event/service"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

// GET /api/v1/events
func GetAllEvents(c *gin.Context) {
	events, err := service.GetAllEventsService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// GET /api/v1/events/:id
func GetEventById(c *gin.Context) {
	id := c.Param("id")
	event, err := service.GetEventByIdService(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

// GET /api/v1/events/:id/ics
func GetEventICS(c *gin.Context) {
	id := c.Param("id")
	icsData, err := service.GetEventICS(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	if icsData == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.Header("Content-Type", "text/calendar")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=event_%s.ics", id))
	c.Data(http.StatusOK, "text/calendar", icsData)
}

// GET /api/v1/events/active
func GetActiveEvents(c *gin.Context) {
	events, err := service.GetActiveEventsService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// POST /api/v1/events/:id/checkout
func CreateEventCheckout(c *gin.Context) {
	id := c.Param("id")
	var request dto.CreateCheckoutRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	successURL := config.GetStripeSettings().StripeRedirectURL + "/thank-you"
	cancelURL := config.GetStripeSettings().StripeRedirectURL

	storedlogs.LogInfo(
		fmt.Sprintf(
			"Creating checkout session for event %s with quantity %d, success URL: %s, cancel URL: %s",
			id, request.Quantity, successURL, cancelURL,
		),
	)

	request.SuccessURL = successURL
	request.CancelURL = cancelURL

	checkout, dbErr := service.CreateEventCheckoutService(id, request)
	if dbErr != nil {
		c.JSON(dbErr.Code, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, checkout)
}

// GET /api/v1/events/:id/exists
func CheckEventExists(c *gin.Context) {
	id := c.Param("id")

	exists, dbErr := service.EventExistsService(id)
	if dbErr != nil {
		c.JSON(dbErr.Code, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
