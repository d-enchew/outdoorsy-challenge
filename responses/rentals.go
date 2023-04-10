package responses

type Rental struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Type            string      `json:"type"`
	VehicleMake     string      `json:"make"`
	VehicleModel    string      `json:"model"`
	VehicleYear     int         `json:"year"`
	VehicleLength   float32     `json:"length"`
	Sleeps          int         `json:"sleeps"`
	Location        Location    `json:"location"`
	PricePerDay     PricePerDay `json:"price"`
	PrimaryImageUrl string      `json:"primary_image_url"`
	User            User        `json:"user"`
}

type PricePerDay struct {
	Day int `json:"day"`
}

type Location struct {
	HomeCity    string  `json:"city"`
	HomeState   string  `json:"state"`
	HomeZip     string  `json:"zip"`
	HomeCountry string  `json:"country"`
	Lat         float32 `json:"lat"`
	Lng         float32 `json:"lng"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
