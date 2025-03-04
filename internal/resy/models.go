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
    LeadTimeInDays int `json:"lead_time_in_days"`
    Venue struct {
        Config struct {
            AllowBypassPaymentMethod int `json:"allow_bypass_payment_method"`
            AllowMultipleResys      int `json:"allow_multiple_resys"`
            EnableInvite           int `json:"enable_invite"`
            EnableResypay          int `json:"enable_resypay"`
            HospitalityIncluded    int `json:"hospitality_included"`
        } `json:"config"`
        Contact struct {
            PhoneNumber string `json:"phone_number"`
        } `json:"contact"`
        MaxPartySize int    `json:"max_party_size"`
        MinPartySize int    `json:"min_party_size"`
        Name         string `json:"name"`
    } `json:"venue"`
}

type SimplifiedVenueConfig struct {
    VenueName      string `json:"venue_name"`
    LeadTimeInDays int    `json:"lead_time_in_days"`
    MaxPartySize   int    `json:"max_party_size"`
    MinPartySize   int    `json:"min_party_size"`
    PhoneNumber    string `json:"phone_number"`
}