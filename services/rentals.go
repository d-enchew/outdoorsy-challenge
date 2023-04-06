package services

import "outdoorsy/models"

type Servicer interface {
	Get(ID int) (*models.Rental, error)
	List(searchCriteria models.RentalSearchQuery) ([]*models.Rental, error)
}

type rentalService struct {
}

func InitializeService() Servicer {
	return &rentalService{}
}

func (rs *rentalService) Get(ID int) (*models.Rental, error) {

	return nil, nil
}
func (rs *rentalService) List(searchCriteria models.RentalSearchQuery) ([]*models.Rental, error) {

	return nil, nil
}
