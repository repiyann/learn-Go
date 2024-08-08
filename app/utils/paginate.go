package utils

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := c.QueryInt("page", 1)
		limit := c.QueryInt("limit", 10)

		offset := (page - 1) * limit
		return db.Order("created_at asc").Offset(offset).Limit(limit)
	}
}
