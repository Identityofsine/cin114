package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/eventstripprice"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/eventstripprice/service"
)

// GET /api/v1/event-stripe-prices
func GetAllEventStripePrices(c *gin.Context) {
	eventStripePrices, err := service.GetAllEventStripePricesService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventStripePrices)
}

// GET /api/v1/event-stripe-prices/:id
func GetEventStripePriceById(c *gin.Context) {
	id := c.Param("id")
	eventStripePrice, err := service.GetEventStripePriceByIdService(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventStripePrice)
}

// GET /api/v1/events/:eventId/stripe-prices
func GetEventStripePricesByEventId(c *gin.Context) {
	eventId := c.Param("eventId")
	eventStripePrices, err := service.GetEventStripePricesByEventIdService(eventId)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventStripePrices)
}

// GET /api/v1/events/:eventId/stripe-prices/active
func GetActiveEventStripePriceByEventId(c *gin.Context) {
	eventId := c.Param("eventId")
	eventStripePrice, err := service.GetActiveEventStripePriceByEventIdService(eventId)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventStripePrice)
}

// GET /api/v1/stripe-prices/:stripePriceId/events
func GetEventStripePricesByStripePriceId(c *gin.Context) {
	stripePriceId := c.Param("stripePriceId")
	eventStripePrices, err := service.GetEventStripePricesByStripePriceIdService(stripePriceId)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventStripePrices)
}

// GET /api/v1/event-stripe-prices/active
func GetActiveEventStripePrices(c *gin.Context) {
	eventStripePrices, err := service.GetActiveEventStripePricesService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventStripePrices)
}

// GET /api/v1/event-stripe-prices/details
func GetEventStripePricesWithDetails(c *gin.Context) {
	eventStripePrices, err := service.GetEventStripePricesWithDetailsService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventStripePrices)
}

// POST /api/v1/event-stripe-prices
func CreateEventStripePrice(c *gin.Context) {
	var request dto.CreateEventStripePriceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventStripePrice, dbErr := service.CreateEventStripePriceService(request)
	if dbErr != nil {
		c.JSON(dbErr.Code, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, eventStripePrice)
}

// PATCH /api/v1/event-stripe-prices/:id/active
func UpdateEventStripePriceActive(c *gin.Context) {
	id := c.Param("id")
	var request dto.UpdateEventStripePriceActiveRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dbErr := service.UpdateEventStripePriceActiveService(id, request); dbErr != nil {
		c.JSON(dbErr.Code, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event stripe price active status updated successfully"})
}

// PUT /api/v1/events/:eventId/stripe-prices/:stripePriceId/activate
func SetActiveEventStripePriceForEvent(c *gin.Context) {
	eventId := c.Param("eventId")
	stripePriceId := c.Param("stripePriceId")

	if dbErr := service.SetActiveEventStripePriceForEventService(eventId, stripePriceId); dbErr != nil {
		c.JSON(dbErr.Code, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event stripe price activated successfully"})
}

// DELETE /api/v1/event-stripe-prices/:id
func DeleteEventStripePrice(c *gin.Context) {
	id := c.Param("id")

	if dbErr := service.DeleteEventStripePriceService(id); dbErr != nil {
		c.JSON(dbErr.Code, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event stripe price deleted successfully"})
}
