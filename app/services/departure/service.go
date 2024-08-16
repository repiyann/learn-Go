package services

import (
	"travelinaja/app/models"
	repositories "travelinaja/app/repositories/departure"

	"github.com/google/uuid"
)

type DepartureService interface {
	CreateDeparture(departure *models.Departure) error
	GetDepartures() ([]models.Departure, error)
	GetDepartureByID(id uuid.UUID) (*models.Departure, error)
	UpdateDeparture(id uuid.UUID, departure *models.Departure) error
	DeleteDeparture(id uuid.UUID) error
}

type departureService struct {
	DepartureRepository repositories.DepartureRepository
}

func NewDepartureService(repo repositories.DepartureRepository) DepartureService {
	return &departureService{
		DepartureRepository: repo,
	}
}

func (s *departureService) CreateDeparture(departure *models.Departure) error {
	departure.DepartID = uuid.New()

	return s.DepartureRepository.CreateDeparture(departure)
}

func (s *departureService) GetDepartures() ([]models.Departure, error) {
	return s.DepartureRepository.GetDepartures()
}

func (s *departureService) GetDepartureByID(id uuid.UUID) (*models.Departure, error) {
	return s.DepartureRepository.GetDepartureByID(id)
}

func (s *departureService) UpdateDeparture(id uuid.UUID, departure *models.Departure) error {
	return s.DepartureRepository.UpdateDeparture(id, departure)
}

func (s *departureService) DeleteDeparture(id uuid.UUID) error {
	return s.DepartureRepository.DeleteDeparture(id)
}
