// pkg/http/utils.go
package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// Constants for configuration
const (
	DefaultTimeout = 3 * time.Second
	UserAgent     = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	ResyOrigin    = "https://resy.com"
)

// ResyHeaders contains the standard headers for Resy API requests
type ResyHeaders struct {
	APIKey    string
	AuthToken string
}

// Req represents an HTTP request configuration
type Req struct {
	QueryParams map[string]string
	Body        []byte
	Headers     map[string]string // Additional custom headers
	Timeout     time.Duration     // Optional custom timeout
}

// getResyHeaders retrieves authentication headers from configuration
func getResyHeaders() (*ResyHeaders, error) {
	apiKey := viper.GetString("RESY_API_KEY")
	authToken := viper.GetString("RESY_AUTH_KEY")

	if apiKey == "" || authToken == "" {
		return nil, fmt.Errorf("missing required authentication credentials: API_KEY=%v, AUTH_KEY=%v", 
			apiKey != "", authToken != "")
	}

	if viper.GetBool("DEBUG") {
		log.Printf("🔑 Auth Headers - API Key: %s, Auth Token: %s", maskString(apiKey), maskString(authToken))
	}

	return &ResyHeaders{
		APIKey:    apiKey,
		AuthToken: authToken,
	}, nil
}

// maskString masks a string for secure logging
func maskString(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:4] + "****" + s[len(s)-4:]
}

// setDefaultHeaders sets the standard headers for Resy API requests
func setDefaultHeaders(req *http.Request, headers *ResyHeaders) {
	req.Header.Set("user-agent", UserAgent)
	req.Header.Set("origin", ResyOrigin)
	req.Header.Set("referrer", ResyOrigin)
	req.Header.Set("x-origin", ResyOrigin)
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("authorization", fmt.Sprintf(`ResyAPI api_key="%s"`, headers.APIKey))
	req.Header.Set("x-resy-auth-token", headers.AuthToken)
	req.Header.Set("x-resy-universal-auth", headers.AuthToken)
}

// createRequest creates an HTTP request with all necessary headers
func createRequest(method, url string, p *Req) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(p.Body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	headers, err := getResyHeaders()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth headers: %w", err)
	}

	setDefaultHeaders(req, headers)

	// Add custom headers if provided
	for key, value := range p.Headers {
		req.Header.Set(key, value)
	}

	// Add query parameters if provided
	if p.QueryParams != nil {
		q := req.URL.Query()
		for key, val := range p.QueryParams {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	return req, nil
}

// makeRequest executes an HTTP request with logging and error handling
func makeRequest(req *http.Request, timeout time.Duration) ([]byte, int, error) {
	client := &http.Client{Timeout: timeout}

	if viper.GetBool("DEBUG") {
		log.Printf("📤 Making %s request to %s", req.Method, req.URL)
		log.Printf("Headers: %v", req.Header)
		if req.Body != nil {
			log.Printf("Body: %s", req.Body)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response: %w", err)
	}

	if viper.GetBool("DEBUG") {
		log.Printf("📥 Response Status: %d", resp.StatusCode)
		log.Printf("Response Body: %s", string(body))
	}

	return body, resp.StatusCode, nil
}

// template creates an HTTP request handler with the specified method and content type
func template(method string, contentType string) func(string, *Req) ([]byte, int, error) {
	return func(url string, p *Req) ([]byte, int, error) {
		if p == nil {
			p = &Req{}
		}

		req, err := createRequest(method, url, p)
		if err != nil {
			return nil, 0, err
		}

		if contentType != "" {
			req.Header.Set("content-type", contentType)
		}

		timeout := DefaultTimeout
		if p.Timeout > 0 {
			timeout = p.Timeout
		}

		return makeRequest(req, timeout)
	}
}

// HTTP request methods
func PostJSON(url string, p *Req) ([]byte, int, error) {
	return template(http.MethodPost, "application/json")(url, p)
}

func PostForm(url string, p *Req) ([]byte, int, error) {
	return template(http.MethodPost, "application/x-www-form-urlencoded")(url, p)
}

func Get(url string, p *Req) ([]byte, int, error) {
	return template(http.MethodGet, "")(url, p)
}