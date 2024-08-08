package models

import (
	"time"

	"github.com/google/uuid"
)

type Destination struct {
	DestiID      uuid.UUID `json:"destiID" gorm:"type:char(36);primaryKey"`
	DestiCity    string    `json:"destiCity" validate:"required"`
	DestiAddress string    `json:"destiAddress" validate:"required"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
