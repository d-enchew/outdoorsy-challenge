package repositories

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"outdoorsy/models"
	"strings"
	"time"
)

// Repositorer interface defines the functions, used to manage a repository. An implementation of the interface
// should be able to handle all operations, despite the repository type - DB (Mongo, Postgres, MySQL), File, etc.
type Repositorer interface {
	Get(ctx context.Context, ID int) (*models.Rental, error)
	List(ctx context.Context, searchCriteria models.RentalSearchQuery) ([]*models.Rental, error)
}

// repository represents an implementation of the Repositorer interface, executing functions on a Postgres database
type rentalRepository struct {
	db *gorm.DB
}

// InitializeRepository creates a postgres managing repository object.
func InitializeRepository(connectionString string) (Repositorer, error) {
	db, err := gorm.Open(postgres.Open(connectionString))

	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &rentalRepository{db: db}, nil
}

// Get executes the lookup in a Postgres databases
func (rr *rentalRepository) Get(ctx context.Context, ID int) (*models.Rental, error) {
	rental := &rental{}

	err := rr.db.WithContext(ctx).
		Table("rentals").
		Preload("User").
		Find(rental, "id = ?", ID).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rental.toModel(), nil
}

// List executes the lookup for multiple elements in a Postgres database
func (rr *rentalRepository) List(ctx context.Context, searchCriteria models.RentalSearchQuery) ([]*models.Rental, error) {
	const distanceThreshold = 50
	query := rr.db.WithContext(ctx).
		Table("rentals").
		Preload("User")

	if searchCriteria.Limit != nil {
		query = query.Limit(*searchCriteria.Limit)
	}
	if searchCriteria.Offset != nil {
		query = query.Offset(*searchCriteria.Offset)
	}
	if searchCriteria.IDs != nil && len(searchCriteria.IDs) > 0 {
		query = query.Where("id in ?", searchCriteria.IDs)
	}

	if searchCriteria.Sort != nil {
		query = query.Order(getSortValue(*searchCriteria.Sort))
	}

	if searchCriteria.PriceMin != nil {
		query = query.Where("price_per_day >= ?", searchCriteria.PriceMin)
	}

	if searchCriteria.PriceMax != nil {
		query = query.Where("price_per_day <= ?", searchCriteria.PriceMax)
	}

	if searchCriteria.Near != nil && len(searchCriteria.Near) > 0 {
		query = query.Where("lat > ? and lat < ? and lng > ? and lng < ?", searchCriteria.Near[0]-distanceThreshold, searchCriteria.Near[0]+distanceThreshold, searchCriteria.Near[1]-distanceThreshold, searchCriteria.Near[1]+distanceThreshold)
	}
	rentals := []rental{}
	err := query.Find(&rentals).Error
	if err != nil {
		return nil, err
	}
	result := []*models.Rental{}
	for _, r := range rentals {
		result = append(result, r.toModel())
	}
	return result, err

}

// getSortValue handle all cases whose names differ from column names in the db
func getSortValue(sortCriteria string) string {
	sortCriteria = strings.ToLower(sortCriteria)
	switch sortCriteria {
	case "price":
		return "price_per_day"

	case "make", "model", "year", "length":
		return "vehicle_" + sortCriteria
	}
	return sortCriteria
}

type rental struct {
	ID              int
	UserID          int
	User            User `gorm:"foreignKey:ID;references:UserID"`
	Name            string
	Type            string
	Description     string
	Sleeps          int
	PricePerDay     int64
	HomeCity        string
	HomeState       string
	HomeZip         string
	HomeCountry     string
	VehicleMake     string
	VehicleModel    string
	VehicleYear     int
	VehicleLength   float32
	Created         time.Time
	Updated         time.Time
	Lat             float32
	Lng             float32
	PrimaryImageUrl string
}

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// toModel converts the DTO object into a model
func (user User) toModel() *models.User {
	return &models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

// toModel converts the DTO object into a model
func (rental rental) toModel() *models.Rental {
	return &models.Rental{
		ID:              rental.ID,
		User:            rental.User.toModel(),
		Name:            rental.Name,
		Type:            rental.Type,
		Description:     rental.Description,
		Sleeps:          rental.Sleeps,
		PricePerDay:     rental.PricePerDay,
		HomeCity:        rental.HomeCity,
		HomeState:       rental.HomeState,
		HomeZip:         rental.HomeZip,
		HomeCountry:     rental.HomeCountry,
		VehicleMake:     rental.VehicleMake,
		VehicleModel:    rental.VehicleModel,
		VehicleYear:     rental.VehicleYear,
		VehicleLength:   rental.VehicleLength,
		Created:         rental.Created,
		Updated:         rental.Updated,
		Lat:             rental.Lat,
		Lng:             rental.Lng,
		PrimaryImageUrl: rental.PrimaryImageUrl,
	}
}
