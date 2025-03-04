// internal/resy/service_test.go
package resy

import (
	"testing"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *ReservationRequest
		wantErr error
	}{
		{
			name: "valid request",
			req: &ReservationRequest{
				RestaurantName: "Test Restaurant",
				Date:           "2024-01-20",
				Time:           "19:00",
				PartySize:      2,
			},
			wantErr: nil,
		},
		{
			name: "missing restaurant",
			req: &ReservationRequest{
				Date:      "2024-01-20",
				Time:      "19:00",
				PartySize: 2,
			},
			wantErr: ErrMissingRestaurantName,
		},
		// Add more test cases
	}

	service := NewService(nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.validateRequest(tt.req)
			if err != tt.wantErr {
				t.Errorf("validateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
