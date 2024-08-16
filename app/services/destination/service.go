package services

import (
	"travelinaja/app/models"
	repositories "travelinaja/app/repositories/destination"

	"github.com/google/uuid"
)

type DestinationService interface {
	CreateDestination(destination *models.Destination) error
	GetDestinations() ([]models.Destination, error)
	GetDestinationByID(id uuid.UUID) (*models.Destination, error)
	UpdateDestination(id uuid.UUID, destination *models.Destination) error
	DeleteDestination(id uuid.UUID) error
}

type destinationService struct {
	DestinationRepository repositories.DestinationRepository
}

func NewDestinationService(repo repositories.DestinationRepository) DestinationService {
	return &destinationService{
		DestinationRepository: repo,
	}
}

func (s *destinationService) CreateDestination(destination *models.Destination) error {
	destination.DestiID = uuid.New()

	return s.DestinationRepository.CreateDestination(destination)
}

func (s *destinationService) GetDestinations() ([]models.Destination, error) {
	return s.DestinationRepository.GetDestinations()
}

func (s *destinationService) GetDestinationByID(id uuid.UUID) (*models.Destination, error) {
	return s.DestinationRepository.GetDestinationByID(id)
}

func (s *destinationService) UpdateDestination(id uuid.UUID, destination *models.Destination) error {
	return s.DestinationRepository.UpdateDestination(id, destination)
}

func (s *destinationService) DeleteDestination(id uuid.UUID) error {
	return s.DestinationRepository.DeleteDestination(id)
}
