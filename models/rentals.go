package models

import "time"

type Rental struct {
	ID              int
	User            *User
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

type RentalSearchQuery struct {
	PriceMin *int
	PriceMax *int
	Limit    *int
	Offset   *int
	IDs      []int
	Near     []float32
	Sort     *string
}
