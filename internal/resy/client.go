// internal/resy/client.go
package resy

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    
    httputils "github.com/aaron-lau/go-get-resy-table/pkg/http"
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
    reqBody, err := json.Marshal(req)
    if err != nil {
        log.Printf("❌ Error marshaling request: %v", err)
        return nil, fmt.Errorf("failed to marshal request: %w", err)
    }

    log.Printf("📤 Sending request to Resy API...")
	log.Printf("%s", reqBody)
	log.Printf("%s", httputils.Req{})
    resp, statusCode, err := httputils.Get("https://api.resy.com/2/user", &httputils.Req{})
    
    if err != nil {
        log.Printf("❌ Error making reservation request: %v", err)
        return nil, err
    }

    // Log the raw response for debugging
    log.Printf("📥 Received response (status %d): %s", statusCode, string(resp))

    if statusCode != 200 {
        return &ReservationResponse{
            Success: false,
            Error:   fmt.Sprintf("unexpected status code: %d - %s", statusCode, string(resp)),
        }, nil
    }

    var response ReservationResponse
    if err := json.Unmarshal(resp, &response); err != nil {
        log.Printf("❌ Failed to unmarshal response: %v\nResponse body: %s", err, string(resp))
        return nil, fmt.Errorf("failed to unmarshal response: %w", err)
    }

    return &response, nil
}

func (c *Client) debugLog(format string, v ...interface{}) {
    if c.debug {
        log.Printf(format, v...)
    }
}