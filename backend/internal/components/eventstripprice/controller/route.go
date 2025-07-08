package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

// This directory contains the routes responsible for handling the requests
// of the event stripe price component of this web application

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	// Event stripe price management routes
	router.GET("/event-stripe-prices", GetAllEventStripePrices)
	router.GET("/event-stripe-prices/:id", GetEventStripePriceById)
	router.GET("/event-stripe-prices/active", GetActiveEventStripePrices)
	router.GET("/event-stripe-prices/details", GetEventStripePricesWithDetails)
	router.POST("/event-stripe-prices", CreateEventStripePrice)
	router.PATCH("/event-stripe-prices/:id/active", UpdateEventStripePriceActive)
	router.DELETE("/event-stripe-prices/:id", DeleteEventStripePrice)

	// Event-specific price routes
	router.GET("/event-stripe-prices/event/:eventId", GetEventStripePricesByEventId)
	router.GET("/event-stripe-prices/event/:eventId/active", GetActiveEventStripePriceByEventId)
	router.PUT("/event-stripe-prices/event/:eventId/stripe-price/:stripePriceId/activate", SetActiveEventStripePriceForEvent)

	// Stripe price-specific event routes
	router.GET("/stripe-prices/:stripePriceId/events", GetEventStripePricesByStripePriceId)

	return router
}

var (
	EventStripePriceRoute = route{}
)
