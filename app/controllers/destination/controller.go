package controllers

import (
	"travelinaja/app/models"
	services "travelinaja/app/services/destination"
	"travelinaja/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DestinationController struct {
	DestinationService services.DestinationService
}

func NewDestinationController(service services.DestinationService) *DestinationController {
	return &DestinationController{
		DestinationService: service,
	}
}

func (controller *DestinationController) CreateDestination(c *fiber.Ctx) error {
	destination := new(models.Destination)

	if err := c.BodyParser(destination); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.Validate(destination); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := controller.DestinationService.CreateDestination(destination); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created destination",
	})
}

func (controller *DestinationController) GetDestinations(c *fiber.Ctx) error {
	destinations, err := controller.DestinationService.GetDestinations()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if len(destinations) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Destinations not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetched all destinations",
		"data":    destinations,
	})
}

func (controller *DestinationController) GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	destiID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	destination, err := controller.DestinationService.GetDestinationByID(destiID)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Destination not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetched destination",
		"data":    destination,
	})
}

func (controller *DestinationController) UpdateDestination(c *fiber.Ctx) error {
	destination := new(models.Destination)

	if err := c.BodyParser(destination); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.Validate(destination); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := c.Params("id")
	destiID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	if err := controller.DestinationService.UpdateDestination(destiID, destination); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Destination not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully updated destination",
	})
}

func (controller *DestinationController) DeleteDestination(c *fiber.Ctx) error {
	id := c.Params("id")
	destiID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	if err := controller.DestinationService.DeleteDestination(destiID); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Destination not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted destination",
	})
}
