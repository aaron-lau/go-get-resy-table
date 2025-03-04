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
	curl -v -X POST http://localhost:8080/book \
		-H "Content-Type: application/json" \
		-d '{}'

test-auth:
	curl -v http://localhost:8080/test-auth | jq '.'

test-auth-direct:
	curl -v https://api.resy.com/2/user \
		-H "Authorization: ResyAPI api_key=\"${RESY_API_KEY}\"" \
		-H "x-resy-auth-token: ${RESY_AUTH_KEY}" \
		-H "x-resy-universal-auth: ${RESY_AUTH_KEY}" | jq '.'


test-venue-config:
	curl -v "http://localhost:8080/venue/config?venue_id=49338" | jq '.'

test-venue-config-direct:
	curl -v "https://api.resy.com/2/config?venue_id=49338" \
		-H "Authorization: ResyAPI api_key=\"${RESY_API_KEY}\"" \
		-H "x-resy-auth-token: ${RESY_AUTH_KEY}" \
		-H "x-resy-universal-auth: ${RESY_AUTH_KEY}" | jq '.'

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