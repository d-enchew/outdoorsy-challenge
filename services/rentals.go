package services

import (
	"context"
	"github.com/sirupsen/logrus"
	"outdoorsy/models"
	"outdoorsy/repositories"
)

// Servicer defines the list of functions, which a service implementation must include
type Servicer interface {
	Get(ctx context.Context, ID int) (*models.Rental, error)
	List(ctx context.Context, searchCriteria models.RentalSearchQuery) ([]*models.Rental, error)
}

// rentalService implementation of the Servicer interface
type rentalService struct {
	repository repositories.Repositorer
	logger     *logrus.Logger
}

// InitializeService Creates a new instance of a service
func InitializeService(repository repositories.Repositorer, logger *logrus.Logger) Servicer {
	return &rentalService{
		repository: repository,
		logger:     logger,
	}
}

// Get Handles the search in the repository for a single rental
func (rs *rentalService) Get(ctx context.Context, ID int) (*models.Rental, error) {
	rental, err := rs.repository.Get(ctx, ID)
	if err != nil {
		rs.logger.Error(err)
		return nil, err
	}
	return rental, nil
}

// List Handles the lookup of multiple elements from the repository, based on a search criteria
func (rs *rentalService) List(ctx context.Context, searchCriteria models.RentalSearchQuery) ([]*models.Rental, error) {
	rentals, err := rs.repository.List(ctx, searchCriteria)
	if err != nil {
		rs.logger.Error(err)
		return nil, err
	}
	return rentals, nil
}
