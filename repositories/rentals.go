package repositories

import "outdoorsy/models"

type Repositorer interface {
	Get(ID int) (*models.Rental, error)
	List(searchCriteria models.RentalSearchQuery) ([]*models.Rental, error)
}

type rentalRepository struct {
}

func InitializeRepository() Repositorer {
	return &rentalRepository{}
}

func (rr *rentalRepository) Get(ID int) (*models.Rental, error) {

	return nil, nil
}
func (rr *rentalRepository) List(searchCriteria models.RentalSearchQuery) ([]*models.Rental, error) {

	return nil, nil
}
