package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

// This directory contains the routes responsible for handling the requests
// of the tickets component of this web application

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	// Ticket management routes
	router.GET("/tickets", GetAllOwnedTickets)
	router.GET("/tickets/:ticketId", GetOwnedTicketById)
	router.GET("/tickets/email/:email", GetOwnedTicketsByEmail)
	router.GET("/tickets/event/:eventId", GetTicketsByEventId)
	router.GET("/tickets/event/:eventId/exists", CheckEventHasTickets)

	return router
}

var (
	TicketRoute = route{}
)
