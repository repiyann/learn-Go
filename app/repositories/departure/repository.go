package repositories

import (
	"travelinaja/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartureRepository interface {
	CreateDeparture(departure *models.Departure) error
	GetDepartures() ([]models.Departure, error)
	GetDepartureByID(id uuid.UUID) (*models.Departure, error)
	UpdateDeparture(id uuid.UUID, departure *models.Departure) error
	DeleteDeparture(id uuid.UUID) error
}

type departureRepository struct {
	DB *gorm.DB
}

func NewDepartureRepository(db *gorm.DB) DepartureRepository {
	return &departureRepository{
		DB: db,
	}
}

func (r *departureRepository) CreateDeparture(departure *models.Departure) error {
	return r.DB.Create(departure).Error
}

func (r *departureRepository) GetDepartures() ([]models.Departure, error) {
	var departures []models.Departure
	err := r.DB.Find(&departures).Error

	return departures, err
}

func (r *departureRepository) GetDepartureByID(id uuid.UUID) (*models.Departure, error) {
	var departure models.Departure
	err := r.DB.First(&departure, "depart_id = ?", id).Error

	return &departure, err
}

func (r *departureRepository) UpdateDeparture(id uuid.UUID, departure *models.Departure) error {
	result := r.DB.Model(&models.Departure{}).Where("depart_id = ?", id).Updates(departure)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *departureRepository) DeleteDeparture(id uuid.UUID) error {
	result := r.DB.Where("depart_id = ?", id).Delete(&models.Departure{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
