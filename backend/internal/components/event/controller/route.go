package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

// This directory contains the routes responsible for handling the requests
// of the event component of this web application

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	// Event management routes
	router.GET("/events", GetAllEvents)
	router.GET("/events/active", GetActiveEvents)
	router.GET("/events/:id", GetEventById)
	router.GET("/events/:id/exists", CheckEventExists)
	router.POST("/events/:id/checkout", CreateEventCheckout)

	return router
}

var (
	EventRoute = route{}
)
