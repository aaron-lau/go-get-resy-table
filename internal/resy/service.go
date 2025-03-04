// internal/resy/service.go
package resy

type Service struct {
	client *Client
}

func NewService(client *Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) BookReservation(req *ReservationRequest) (*ReservationResponse, error) {
	// Add any business logic here before making the API call
	if err := s.validateRequest(req); err != nil {
		return &ReservationResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Call the Resy API through the client
	return s.client.BookReservation(req)
}

func (s *Service) validateRequest(req *ReservationRequest) error {
	if req.RestaurantName == "" {
		return ErrMissingRestaurantName
	}
	if req.Date == "" {
		return ErrMissingDate
	}
	if req.Time == "" {
		return ErrMissingTime
	}
	if req.PartySize <= 0 {
		return ErrInvalidPartySize
	}
	return nil
}
