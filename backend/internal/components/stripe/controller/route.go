package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

// This directory contains the routes responsible for handling the requests
// of the stripe component of this web application

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	// Stripe webhook routes
	router.POST("/stripe/webhook", HandleStripeWebhook)

	return router
}

var (
	StripeRoute = route{}
)
