// internal/resy/client.go
package resy

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

type Client struct {
    apiKey     string
    authKey    string
    httpClient *http.Client
    debug      bool
}

func NewClient(apiKey, authKey string, debug bool) *Client {
    return &Client{
        apiKey:     apiKey,
        authKey:    authKey,
        httpClient: &http.Client{},
        debug:      debug,
    }
}

func (c *Client) BookReservation(req *ReservationRequest) (*ReservationResponse, error) {
    startTime := time.Now()
    log.Printf("🔄 Starting reservation request for %s", req.RestaurantName)

    reqBody, err := json.Marshal(req)
    if err != nil {
        log.Printf("❌ Error marshaling request: %v", err)
        return nil, fmt.Errorf("failed to marshal request: %w", err)
    }

    headers := map[string]string{
        "Authorization":     fmt.Sprintf("ResyAPI api_key=\"%s\"", c.apiKey),
        "x-resy-auth-token": c.authKey,
        "Content-Type":      "application/json",
    }

    // Pretty print the request details
    log.Printf("📝 Request Details:")
    log.Printf("  Restaurant: %s", req.RestaurantName)
    log.Printf("  Date: %s", req.Date)
    log.Printf("  Time: %s", req.Time)
    log.Printf("  Party Size: %d", req.PartySize)
    log.Printf("  Request Body: %s", string(reqBody))
    log.Printf("  Headers: %v", headers)

    // TODO: Make actual API call
    // For now, return mock response with simulated delay
    time.Sleep(500 * time.Millisecond) // Simulate API call

    response := &ReservationResponse{
        Success:       true,
        ReservationID: fmt.Sprintf("mock-%s-%s-%s", req.RestaurantName, req.Date, time.Now().Format("150405")),
    }

    duration := time.Since(startTime)
    log.Printf("✅ Reservation completed in %v", duration)
    log.Printf("📎 Reservation ID: %s", response.ReservationID)

    return response, nil
}

func (c *Client) debugLog(format string, v ...interface{}) {
    if c.debug {
        log.Printf(format, v...)
    }
}