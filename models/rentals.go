package models

type Rental struct {
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
