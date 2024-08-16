package destination

import (
	controllers "travelinaja/app/controllers/destination"
	"travelinaja/app/database"
	repositories "travelinaja/app/repositories/destination"
	services "travelinaja/app/services/destination"

	"github.com/gofiber/fiber/v2"
)

func DestinationRoutes(app fiber.Router) {
	db := database.DBConn

	destinationRepo := repositories.NewDestinationRepository(db)
	destinationService := services.NewDestinationService(destinationRepo)
	destinationController := controllers.NewDestinationController(destinationService)

	app.Post("/create", destinationController.CreateDestination)
	app.Get("/get", destinationController.GetDestinations)
	app.Get("/get/:id", destinationController.GetDestinationByID)
	app.Put("/update/:id", destinationController.UpdateDestination)
	app.Delete("/delete/:id", destinationController.DeleteDestination)
}
