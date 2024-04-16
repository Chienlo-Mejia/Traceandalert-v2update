package location

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

func GetUserLocation() (Location, error) {
	apiKey := "AIzaSyBOgoIW9c_8vgl0i7mUGoltXND2VzSQHFg"
	url := fmt.Sprintf("https://www.googleapis.com/geolocation/v1/geolocate?key=%s", apiKey)

	// Create an empty JSON request body
	requestBody := []byte(`{}`)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return Location{}, fmt.Errorf("failed to get location data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var location Location
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return Location{}, fmt.Errorf("failed to decode location data: %v", err)
	}

	return location, nil
}
