package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/tickets/service"
)

// GetTicketsByEventId handles GET /tickets/event/:eventId
func GetTicketsByEventId(c *gin.Context) {
	eventId := c.Param("eventId")

	tickets, err := service.GetTicketsByEventIdService(eventId)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
		"count":   len(tickets),
	})
}

// GetAllOwnedTickets handles GET /tickets
func GetAllOwnedTickets(c *gin.Context) {
	tickets, err := service.GetAllOwnedTicketsService()
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
		"count":   len(tickets),
	})
}

// GetOwnedTicketsByEmail handles GET /tickets/email/:email
func GetOwnedTicketsByEmail(c *gin.Context) {
	email := c.Param("email")

	tickets, err := service.GetOwnedTicketsByEmailService(email)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tickets": tickets,
		"count":   len(tickets),
	})
}

// GetOwnedTicketById handles GET /tickets/:ticketId
func GetOwnedTicketById(c *gin.Context) {
	ticketId := c.Param("ticketId")

	ticket, err := service.GetOwnedTicketByIdService(ticketId)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ticket": ticket,
	})
}

// CheckEventHasTickets handles GET /tickets/event/:eventId/exists
func CheckEventHasTickets(c *gin.Context) {
	eventId := c.Param("eventId")

	tickets, err := service.GetTicketsByEventIdService(eventId)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"has_tickets":  len(tickets) > 0,
		"ticket_count": len(tickets),
	})
}
