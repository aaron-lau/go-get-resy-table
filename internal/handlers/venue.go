// internal/handlers/venue.go
package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    
    "github.com/spf13/viper"
    httputils "github.com/aaron-lau/go-get-resy-table/pkg/http"
    "github.com/aaron-lau/go-get-resy-table/internal/resy"
)

func GetVenueConfig(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Get venue_id from query parameters
    venueID := r.URL.Query().Get("venue_id")
    if venueID == "" {
        http.Error(w, "venue_id is required", http.StatusBadRequest)
        return
    }

    log.Printf("🏪 Fetching config for venue ID: %s", venueID)

    // Create request object
    reqObj := &httputils.Req{
        QueryParams: map[string]string{
            "venue_id": venueID,
        },
    }

    // Make request to Resy API
    resp, _, err := httputils.Get("https://api.resy.com/2/config", reqObj)
    if err != nil {
        log.Printf("❌ Venue config request failed: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	
    // Parse the full response
	var venueConfig resy.VenueConfigResponse
    if err := json.Unmarshal(resp, &venueConfig); err != nil {
        log.Printf("❌ Failed to parse venue config: %v", err)
        log.Printf("Raw response: %s", string(resp))
        http.Error(w, "Failed to parse venue configuration", http.StatusInternalServerError)
        return
    }

    // Extract the specific fields we want
    simplified := resy.SimplifiedVenueConfig{
        VenueName:      venueConfig.Venue.Name,
        LeadTimeInDays: venueConfig.LeadTimeInDays,
        MaxPartySize:   venueConfig.Venue.MaxPartySize,
        MinPartySize:   venueConfig.Venue.MinPartySize,
        PhoneNumber:    venueConfig.Venue.Contact.PhoneNumber,
    }

	if viper.GetBool("DEBUG") {
		log.Printf("📍 Venue: %s", simplified.VenueName)
		log.Printf("⏰ Lead Time: %d days", simplified.LeadTimeInDays)
		log.Printf("👥 Party Size: %d-%d people", simplified.MinPartySize, simplified.MaxPartySize)
		log.Printf("📞 Phone: %s", simplified.PhoneNumber)
	}

    // Return the simplified response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(simplified)
}