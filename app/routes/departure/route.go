package departure

import (
	"travelinaja/app/controllers/departure"

	"github.com/gofiber/fiber/v2"
)

func DepartureRoutes(app fiber.Router) {
	app.Post("/create", departure.CreateDeparture)
	app.Get("/get", departure.GetDepartures)
	app.Get("/get/:id", departure.GetDepartureByID)
	app.Put("/update/:id", departure.UpdateDeparture)
	app.Delete("/delete/:id", departure.DeleteDeparture)
}
