package repositories

import (
	"travelinaja/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DestinationRepository interface {
	CreateDestination(destination *models.Destination) error
	GetDestinations() ([]models.Destination, error)
	GetDestinationByID(id uuid.UUID) (*models.Destination, error)
	UpdateDestination(id uuid.UUID, destination *models.Destination) error
	DeleteDestination(id uuid.UUID) error
}

type destinationRepository struct {
	DB *gorm.DB
}

func NewDestinationRepository(db *gorm.DB) DestinationRepository {
	return &destinationRepository{
		DB: db,
	}
}

func (r *destinationRepository) CreateDestination(destination *models.Destination) error {
	return r.DB.Create(destination).Error
}

func (r *destinationRepository) GetDestinations() ([]models.Destination, error) {
	var destination []models.Destination
	err := r.DB.Find(&destination).Error

	return destination, err
}

func (r *destinationRepository) GetDestinationByID(id uuid.UUID) (*models.Destination, error) {
	var destination models.Destination
	err := r.DB.First(&destination, "desti_id = ?", id).Error

	return &destination, err
}

func (r *destinationRepository) UpdateDestination(id uuid.UUID, destination *models.Destination) error {
	result := r.DB.Model(&models.Destination{}).Where("desti_id", id).Updates(destination)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (r *destinationRepository) DeleteDestination(id uuid.UUID) error {
	result := r.DB.Where("desti_id = ?", id).Delete(&models.Destination{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
