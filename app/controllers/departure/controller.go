package controllers

import (
	"travelinaja/app/models"
	services "travelinaja/app/services/departure"
	"travelinaja/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DepartureController struct {
	DepartureService services.DepartureService
}

func NewDepartureController(service services.DepartureService) *DepartureController {
	return &DepartureController{
		DepartureService: service,
	}
}

func (controller *DepartureController) CreateDeparture(c *fiber.Ctx) error {
	departure := new(models.Departure)

	if err := c.BodyParser(departure); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErrors := utils.Validate(departure)
	if len(validationErrors.Errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": validationErrors.Errors,
			},
		})
	}

	if err := controller.DepartureService.CreateDeparture(departure); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created departure",
	})
}

func (controller *DepartureController) GetDepartures(c *fiber.Ctx) error {
	departures, err := controller.DepartureService.GetDepartures()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if len(departures) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Departures not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetched all departures",
		"data":    departures,
	})
}

func (controller *DepartureController) GetDepartureByID(c *fiber.Ctx) error {
	id := c.Params("id")
	departureID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	departure, err := controller.DepartureService.GetDepartureByID(departureID)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Departure not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetched departure",
		"data":    departure,
	})
}

func (controller *DepartureController) UpdateDeparture(c *fiber.Ctx) error {
	departure := new(models.Departure)

	if err := c.BodyParser(departure); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErrors := utils.Validate(departure)
	if len(validationErrors.Errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": validationErrors.Errors,
			},
		})
	}

	id := c.Params("id")
	departureID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	if err := controller.DepartureService.UpdateDeparture(departureID, departure); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Departure not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully updated departure",
	})
}

func (controller *DepartureController) DeleteDeparture(c *fiber.Ctx) error {
	id := c.Params("id")
	departureID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	if err := controller.DepartureService.DeleteDeparture(departureID); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Departure not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted departure",
	})
}
