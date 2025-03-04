// internal/handlers/reservation.go
package handlers

import (
	"encoding/json"
	"net/http"
    "fmt"
	"log"

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
        log.Printf("❌ Error decoding request: %v", err)
        http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
        return
    }

    log.Printf("📝 Received reservation request for %s", req.RestaurantName)
    
    resp, err := h.service.BookReservation(&req)
    if err != nil {
        log.Printf("❌ Error booking reservation: %v", err)
        http.Error(w, fmt.Sprintf("Error booking reservation: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        log.Printf("❌ Error encoding response: %v", err)
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }
}
