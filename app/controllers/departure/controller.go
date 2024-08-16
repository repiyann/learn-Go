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
			"message": fiber.Map{
				"errors": "Invalid input",
			},
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
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created departure",
	})
}

func (controller *DepartureController) GetDepartures(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	departures, totalCount, err := controller.DepartureService.GetDepartures(limit, offset)

	if err != nil {
		return err
	}

	if len(departures) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": "Departures not found",
			},
		})
	}

	totalPages := (totalCount + limit - 1) / limit

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Successfully fetched all departures",
		"data":        departures,
		"currentPage": page,
		"totalPages":  totalPages,
		"limit":       limit,
	})
}

func (controller *DepartureController) GetDepartureByID(c *fiber.Ctx) error {
	id := c.Params("id")
	departureID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": "Invalid ID format",
			},
		})
	}

	departure, err := controller.DepartureService.GetDepartureByID(departureID)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fiber.Map{
					"errors": "Departure not found",
				},
			})
		}

		return err
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
			"message": fiber.Map{
				"errors": "Invalid input",
			},
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
			"message": fiber.Map{
				"errors": "Invalid ID format",
			},
		})
	}

	if err := controller.DepartureService.UpdateDeparture(departureID, departure); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fiber.Map{
					"errors": "Departure not found",
				},
			})
		}

		return err
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
			"message": fiber.Map{
				"errors": "Invalid ID format",
			},
		})
	}

	if err := controller.DepartureService.DeleteDeparture(departureID); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fiber.Map{
					"errors": "Departure not found",
				},
			})
		}

		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted departure",
	})
}
