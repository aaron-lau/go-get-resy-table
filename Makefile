# Makefile

# Variables
BINARY_NAME=reservation-bot
GO=go

.PHONY: run test build clean curl-test

# Run the application
run:
	$(GO) run cmd/main.go

# Run all tests
test:
	$(GO) test ./... -v

# Build the application
build:
	$(GO) build -o bin/$(BINARY_NAME) cmd/main.go

# Clean build files
clean:
	rm -rf bin/

# Test with curl
curl-test:
	curl -X POST http://localhost:8080/book \
		-H "Content-Type: application/json" \
		-H "X-Resy-Auth-Token: ${RESY_AUTH_KEY}" \
		-d '{"restaurant_name": "Test Restaurant", "date": "2024-02-14", "time": "19:00", "party_size": 2}'

# Run with environment variables
run-with-env:
	RESY_API_KEY=${RESY_API_KEY} RESY_AUTH_KEY=${RESY_AUTH_KEY} $(GO) run cmd/main.go

# Test specific packages
test-resy:
	$(GO) test ./internal/resy/... -v

test-handlers:
	$(GO) test ./internal/handlers/... -v

# Run with debug output
run-debug:
	RESY_API_KEY=${RESY_API_KEY} RESY_AUTH_KEY=${RESY_AUTH_KEY} DEBUG=true $(GO) run cmd/main.go