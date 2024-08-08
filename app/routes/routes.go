package routes

import (
	"travelinaja/app/routes/departure"
	"travelinaja/app/routes/destination"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	version := api.Group("/v1")

	departureGroup := version.Group("/departure")
	destinationGroup := version.Group("/destination")

	departure.DepartureRoutes(departureGroup)
	destination.DestinationRoutes(destinationGroup)
}
