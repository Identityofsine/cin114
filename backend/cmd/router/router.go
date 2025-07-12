package router

import (
	"github.com/gin-gonic/gin"
	authController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/auth/controller"
	eventController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/event/controller"
	eventStripePriceController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/eventstripprice/controller"
	healthController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/health/controller"
	registerController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/register/controller"
	logsController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/storedlogs/controller"
	stripeController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/stripe/controller"
	videoController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/video/controller"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set up the router
	setupRoutes(r)

	return r
}

func setupRoutes(engine *gin.Engine) {
	// Set up the routes for the application
	api := engine.Group("/api/v1")

	// Top Level Middleware
	api.Use(middlewares.UseCors().Middleware)

	//inject your routes here:
	//login
	authController.AuthRoute.UseRouter(api)
	registerController.RegisterRoute.UseRouter(api)

	eventController.EventRoute.UseRouter(api)
	healthController.HealthRoute.UseRouter(api)
	stripeController.StripeRoute.UseRouter(api)
	videoController.VideoRoute.UseRouter(api)

	api.Use(middlewares.UseAuthenticationEnforcementMiddleware().Middleware)

	logsController.LogsRoute.UseRouter(api)
	eventStripePriceController.EventStripePriceRoute.UseRouter(api)

}
