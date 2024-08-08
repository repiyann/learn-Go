package destination

import (
	"math"
	"travelinaja/app/database"
	"travelinaja/app/models"
	"travelinaja/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateDestination(c *fiber.Ctx) error {
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

	destination.DestiID = uuid.New()
	result := database.DBConn.Create(&destination)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created destination",
	})
}

func GetDestinations(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	var totalDestinations int64
	if err := database.DBConn.Model(&models.Destination{}).Count(&totalDestinations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	totalPages := int(math.Ceil(float64(totalDestinations) / float64(limit)))

	destinations := []models.Destination{}
	result := database.DBConn.Offset(offset).Limit(limit).Order("created_at asc").Find(&destinations)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	if len(destinations) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Destinations not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Successfully fetched all destinations",
		"data":       destinations,
		"page":       page,
		"limit":      limit,
		"totalPages": totalPages,
	})
}

func GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	destiID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	var destination models.Destination
	result := database.DBConn.First(&destination, destiID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Destination not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetched destination",
		"data":    destination,
	})
}

func UpdateDestination(c *fiber.Ctx) error {
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

	result := database.DBConn.Model(&models.Destination{}).Where("desti_id = ?", destiID).Updates(destination)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Destination not found",
		})
	}
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully updated destination",
	})
}

func DeleteDestination(c *fiber.Ctx) error {
	id := c.Params("id")
	destiID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}

	result := database.DBConn.Where("desti_id", destiID).Delete(&models.Destination{})

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Destination not found",
		})
	}
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted destination",
	})
}
