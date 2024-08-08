package models

import (
	"time"

	"github.com/google/uuid"
)

type Departure struct {
	DepartID      uuid.UUID `json:"departID" gorm:"type:char(36);primaryKey"`
	DepartCity    string    `json:"departCity" validate:"required"`
	DepartAddress string    `json:"departAddress" validate:"required"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
