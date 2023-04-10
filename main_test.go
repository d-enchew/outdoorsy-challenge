package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"outdoorsy/responses"
	"testing"
)

func TestMinMaxPrice(t *testing.T) {
	minPrice, maxPrice := 5000, 9000
	resp, err := http.Get(fmt.Sprintf("http://localhost:1212/rentals?price_min=%d&price_max=%d", minPrice, maxPrice))
	if err != nil {
		t.Error("Error on executing request")
	}
	rentals := []responses.Rental{}
	err = json.NewDecoder(resp.Body).Decode(&rentals)
	if err != nil {
		t.Error("Error on decoding response")
	}
	for _, rental := range rentals {
		if rental.PricePerDay.Day < minPrice || rental.PricePerDay.Day > maxPrice {
			t.Error("Endpoint returned a wrong value")
		}
	}
}

func TestLimit(t *testing.T) {
	limit := 5
	resp, err := http.Get(fmt.Sprintf("http://localhost:1212/rentals?limit=%d", limit))
	if err != nil {
		t.Error("Error on executing request")
	}
	rentals := []responses.Rental{}
	err = json.NewDecoder(resp.Body).Decode(&rentals)
	if err != nil {
		t.Error("Error on decoding response")
	}

	if len(rentals) > limit {
		// handle case of not having enough results (less than limit results returned)
		t.Error("Returned more result than needed")
	}
}

func TestGetByIDs(t *testing.T) {
	ids := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
	}
	resp, err := http.Get(fmt.Sprintf("http://localhost:1212/rentals?ids=%s", "1,2,3"))
	if err != nil {
		t.Error("Error on executing request")
	}
	rentals := []responses.Rental{}
	err = json.NewDecoder(resp.Body).Decode(&rentals)
	if err != nil {
		t.Error("Error on decoding response")
	}

	for _, rental := range rentals {
		if _, ok := ids[rental.ID]; !ok {
			t.Error("Received unwanted rental")
		}
	}
}

func TestSortPrice(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:1212/rentals?sort=price"))
	if err != nil {
		t.Error("Error on executing request")
	}
	rentals := []responses.Rental{}
	err = json.NewDecoder(resp.Body).Decode(&rentals)
	if err != nil {
		t.Error("Error on decoding response")
	}

	currentPrice := 0
	for _, rental := range rentals {
		if currentPrice > rental.PricePerDay.Day {
			t.Errorf("Incorrect ordering returned")
		}
		currentPrice = rental.PricePerDay.Day
	}

}

func TestMultipleCriteria(t *testing.T) {
	minPrice, maxPrice := 5000, 15000
	limit := 6
	resp, err := http.Get(fmt.Sprintf("http://localhost:1212/rentals?sort=price&price_min=%d&price_max=%d&limit=%d", minPrice, maxPrice, limit))
	if err != nil {
		t.Error("Error on executing request")
	}
	rentals := []responses.Rental{}
	err = json.NewDecoder(resp.Body).Decode(&rentals)
	if err != nil {
		t.Error("Error on decoding response")
	}

	if len(rentals) > limit {
		t.Error("Returned more result than needed")
	}

	currentPrice := 0
	for _, rental := range rentals {
		if rental.PricePerDay.Day < minPrice || rental.PricePerDay.Day > maxPrice {
			t.Error("Endpoint returned a wrong value")
		}
		if currentPrice > rental.PricePerDay.Day {
			t.Errorf("Incorrect ordering returned")
		}
		currentPrice = rental.PricePerDay.Day
	}
}
