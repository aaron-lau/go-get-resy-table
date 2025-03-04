// internal/handlers/reservation.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aaron-lau/go-get-resy-table/internal/resy"
)

type ReservationHandler struct {
	service *resy.Service
}

func NewReservationHandler(service *resy.Service) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (h *ReservationHandler) BookReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req resy.ReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.service.BookReservation(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
