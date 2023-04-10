package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"outdoorsy/models"
	"outdoorsy/responses"
	"outdoorsy/services"
	"strconv"
	"strings"
)

// HTTPHandler defines the methods which should be implemented by any implementation, in order to run this API.
// Using the interface, it is possible to easily switch between different modules for http routing and handling - mux, gin, etc
type HTTPHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

// httpHandler Implementation of the HTTPHandler interface
type httpHandler struct {
	service services.Servicer
}

// InitializeHandler creates a new handler to handle incoming HTTP requests, using gorilla mux
func InitializeHandler(service services.Servicer) HTTPHandler {
	return &httpHandler{service: service}
}

// Get Handles the endpoint, related to retrieving a single rental's information
func (rh *httpHandler) Get(w http.ResponseWriter, r *http.Request) {
	idQuery, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "Rental ID parameter is missing", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		http.Error(w, "Invalid ID passed", http.StatusBadRequest)
		return
	}
	rental, err := rh.service.Get(r.Context(), id)
	if err != nil {
		http.Error(w, "Error while getting the rental", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(convertRentalToResponse(rental))
	if err != nil {
		http.Error(w, "Error while encoding the result", http.StatusInternalServerError)
		return
	}
}

// List Handles the list rentals endpoint, allowing optional parameters for filtering, paging and sorting
func (rh *httpHandler) List(w http.ResponseWriter, r *http.Request) {
	searchQuery := models.RentalSearchQuery{}

	priceMinParam := r.URL.Query().Get("price_min")
	if priceMinParam != "" {
		priceMin, err := strconv.Atoi(priceMinParam)
		if err != nil {
			http.Error(w, "Invalid value passed for price_min", http.StatusBadRequest)
			return
		}
		searchQuery.PriceMin = &priceMin
	}
	priceMaxParam := r.URL.Query().Get("price_max")
	if priceMaxParam != "" {
		priceMax, err := strconv.Atoi(priceMaxParam)
		if err != nil {
			http.Error(w, "Invalid value passed for price_max", http.StatusBadRequest)
			return
		}
		searchQuery.PriceMax = &priceMax
	}
	limitParam := r.URL.Query().Get("limit")
	if limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			http.Error(w, "Invalid value passed for limit", http.StatusBadRequest)
			return
		}
		searchQuery.Limit = &limit
	}
	offsetParam := r.URL.Query().Get("offset")
	if offsetParam != "" {
		offset, err := strconv.Atoi(offsetParam)
		if err != nil {
			http.Error(w, "Invalid value passed for offset", http.StatusBadRequest)
			return
		}
		searchQuery.Offset = &offset
	}
	idsParam := r.URL.Query().Get("ids")
	if idsParam != "" {
		ids := strings.Split(idsParam, ",")
		for _, value := range ids {
			id, err := strconv.Atoi(value)
			if err != nil {
				http.Error(w, "Invalid value for ID passed", http.StatusBadRequest)
				return
			}
			searchQuery.IDs = append(searchQuery.IDs, id)
		}
	}
	nearParam := r.URL.Query().Get("near")
	if nearParam != "" {
		nearValues := strings.Split(nearParam, ",")
		for _, value := range nearValues {
			result, err := strconv.ParseFloat(value, 32)
			if err != nil {
				http.Error(w, "Invalid value passed for lon or lat", http.StatusBadRequest)
				return
			}
			searchQuery.Near = append(searchQuery.Near, float32(result))
		}
	}
	sort := r.URL.Query().Get("sort")
	if sort != "" {
		if isValid := validateSortCriteria(sort); !isValid {
			http.Error(w, "Incorrect sort criteria passed", http.StatusBadRequest)
			return
		}
		searchQuery.Sort = &sort
	}

	rentals, err := rh.service.List(r.Context(), searchQuery)
	if err != nil {
		http.Error(w, "Error while searching for rentals", http.StatusInternalServerError)
		return
	}
	response := make([]responses.Rental, len(rentals))
	for i, rental := range rentals {
		response[i] = convertRentalToResponse(rental)
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error while encoding the result", http.StatusInternalServerError)
		return
	}

}

// validateSortCriteria Validates if the sort criteria is a valid value. Could be extended in order to run more ordering criteria
func validateSortCriteria(sortCriteria string) bool {
	isValid := false
	sortCriteria = strings.ToLower(sortCriteria)
	switch sortCriteria {
	case "id",
		"name",
		"description",
		"type",
		"make",
		"model",
		"year",
		"sleeps",
		"price":
		isValid = true

	}
	return isValid
}

// convertRentalToResponse Converts a model to a response structure
func convertRentalToResponse(rental *models.Rental) responses.Rental {
	return responses.Rental{
		ID:            rental.ID,
		Name:          rental.Name,
		Description:   rental.Description,
		Type:          rental.Type,
		VehicleMake:   rental.VehicleMake,
		VehicleModel:  rental.VehicleModel,
		VehicleYear:   rental.VehicleYear,
		VehicleLength: rental.VehicleLength,
		Sleeps:        rental.Sleeps,
		Location: responses.Location{
			HomeCity:    rental.HomeCity,
			HomeState:   rental.HomeState,
			HomeZip:     rental.HomeZip,
			HomeCountry: rental.HomeCountry,
			Lat:         rental.Lat,
			Lng:         rental.Lng,
		},
		PricePerDay: responses.PricePerDay{
			Day: int(rental.PricePerDay),
		},
		PrimaryImageUrl: rental.PrimaryImageUrl,
		User: responses.User{
			ID:        rental.User.ID,
			FirstName: rental.User.FirstName,
			LastName:  rental.User.LastName,
		},
	}
}
