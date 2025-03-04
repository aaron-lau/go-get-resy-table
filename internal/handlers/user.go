// internal/handlers/user.go
package handlers

import (
    "log"
    "net/http"
    
    "github.com/spf13/viper"
    httputils "github.com/aaron-lau/go-get-resy-table/pkg/http"
)

func TestUserAuth(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    log.Printf("🔑 Testing Resy authentication...")
    
    // Create request with debug info
    reqObj := &httputils.Req{
        QueryParams: map[string]string{},
    }

    if viper.GetBool("DEBUG") {
    	log.Printf("📝 Request object: %+v", reqObj)
	}
    
    resp, statusCode, err := httputils.Get("https://api.resy.com/2/user", reqObj)
    
    if err != nil {
        log.Printf("❌ Auth test failed: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	if viper.GetBool("DEBUG") {
    	log.Printf("📥 Received response (status %d): %s", statusCode, string(resp))
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    w.Write(resp)
}