// internal/resy/models.go
package resy

import "errors"

// Error definitions
var (
	ErrMissingRestaurantName = errors.New("restaurant name is required")
	ErrMissingDate           = errors.New("date is required")
	ErrMissingTime           = errors.New("time is required")
	ErrInvalidPartySize      = errors.New("party size must be greater than 0")
)

type ReservationRequest struct {
	RestaurantName string `json:"restaurant_name"`
	Date           string `json:"date"`
	Time           string `json:"time"`
	PartySize      int    `json:"party_size"`
}

type ReservationResponse struct {
	Success       bool   `json:"success"`
	ReservationID string `json:"reservation_id,omitempty"`
	Error         string `json:"error,omitempty"`
}

type VenueConfigRequest struct {
    VenueID string `json:"venue_id"`
}

type VenueConfigResponse struct {
    Meta struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
    } `json:"meta"`
    Data struct {
        Name        string `json:"name"`
        ID          int    `json:"id"`
        URL         string `json:"url"`
        Config      map[string]interface{} `json:"config"`
        // Add more fields as needed
    } `json:"data"`
}