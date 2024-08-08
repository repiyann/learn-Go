package destination

import (
	"travelinaja/app/controllers/destination"

	"github.com/gofiber/fiber/v2"
)

func DestinationRoutes(app fiber.Router) {
	app.Post("/create", destination.CreateDestination)
	app.Get("/get", destination.GetDestinations)
	app.Get("/get/:id", destination.GetDestinationByID)
	app.Put("/update/:id", destination.UpdateDestination)
	app.Delete("/delete/:id", destination.DeleteDestination)
}
