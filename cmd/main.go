// cmd/main.go
package main

import (
	"log"
	"net/http"

	"github.com/aaron-lau/go-get-resy-table/internal/config"
	"github.com/aaron-lau/go-get-resy-table/internal/handlers"
	"github.com/aaron-lau/go-get-resy-table/internal/resy"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: .env file not found")
    }

    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    resyClient := resy.NewClient(cfg.ResyAPIKey, cfg.ResyAuthKey, cfg.Debug)
    resyService := resy.NewService(resyClient)
	reservationHandler := handlers.NewReservationHandler(resyService)

	http.HandleFunc("/book", reservationHandler.BookReservation)
    http.HandleFunc("/test-auth", handlers.TestUserAuth)
    http.HandleFunc("/venue/config", handlers.GetVenueConfig)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
