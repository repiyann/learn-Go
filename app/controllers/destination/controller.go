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
			"message": fiber.Map{
				"errors": "Invalid input",
			},
		})
	}

	validationErrors := utils.Validate(destination)
	if len(validationErrors.Errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": validationErrors.Errors,
			},
		})
	}

	if err := controller.DestinationService.CreateDestination(destination); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created destination",
	})
}

func (controller *DestinationController) GetDestinations(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	destinations, totalCount, err := controller.DestinationService.GetDestinations(limit, offset)

	if err != nil {
		return err
	}

	if len(destinations) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": "Destinations not found",
			},
		})
	}

	totalPages := (totalCount + limit - 1) / limit

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Successfully fetched all destinations",
		"data":        destinations,
		"currentPage": page,
		"totalPages":  totalPages,
		"limit":       limit,
	})
}

func (controller *DestinationController) GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	destiID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": "Invalid ID format",
			},
		})
	}

	destination, err := controller.DestinationService.GetDestinationByID(destiID)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fiber.Map{
					"errors": "Destination not found",
				},
			})
		}

		return err
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
			"message": fiber.Map{
				"errors": "Invalid input",
			},
		})
	}

	validationErrors := utils.Validate(destination)
	if len(validationErrors.Errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": validationErrors.Errors,
			},
		})
	}

	id := c.Params("id")
	destiID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fiber.Map{
				"errors": "Invalid ID format",
			},
		})
	}

	if err := controller.DestinationService.UpdateDestination(destiID, destination); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fiber.Map{
					"errors": "Destination not found",
				},
			})
		}

		return err
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
			"message": fiber.Map{
				"errors": "Invalid ID format",
			},
		})
	}

	if err := controller.DestinationService.DeleteDestination(destiID); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fiber.Map{
					"errors": "Destination not found",
				},
			})
		}

		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted destination",
	})
}
