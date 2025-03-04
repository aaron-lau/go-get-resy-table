// internal/handlers/venue.go
package handlers

import (
    "log"
    "net/http"
    
    "github.com/spf13/viper"
    httputils "github.com/aaron-lau/go-get-resy-table/pkg/http"
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
    resp, statusCode, err := httputils.Get("https://api.resy.com/2/config", reqObj)
    if err != nil {
        log.Printf("❌ Venue config request failed: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	if viper.GetBool("DEBUG") {
    	log.Printf("📥 Received response (status %d): %s", statusCode, string(resp))
	}

    // Pass through the response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    w.Write(resp)
}