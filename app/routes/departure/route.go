package departure

import (
	controllers "travelinaja/app/controllers/departure"
	"travelinaja/app/database"
	repositories "travelinaja/app/repositories/departure"
	services "travelinaja/app/services/departure"

	"github.com/gofiber/fiber/v2"
)

func DepartureRoutes(app fiber.Router) {
	db := database.DBConn

	departureRepo := repositories.NewDepartureRepository(db)
	departureService := services.NewDepartureService(departureRepo)
	departureController := controllers.NewDepartureController(departureService)

	app.Post("/create", departureController.CreateDeparture)
	app.Get("/get", departureController.GetDepartures)
	app.Get("/get/:id", departureController.GetDepartureByID)
	app.Put("/update/:id", departureController.UpdateDeparture)
	app.Delete("/delete/:id", departureController.DeleteDeparture)
}
