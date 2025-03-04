
# Go Get Resy Table 🍽️

A Go application that helps automate restaurant reservations on Resy.

## Features

- Make restaurant reservations through Resy's API
- Local development and testing support
- GCP Cloud Run deployment ready
- Configurable through environment variables
- Detailed logging for debugging

## Prerequisites

- Go 1.24 or later
- Make
- Docker (for container builds)
- Resy API credentials
- GCP account (for deployment)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/aaron-lau/go-get-resy-table.git
cd go-get-resy-table
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file:
```env
RESY_API_KEY=your_api_key_here
RESY_AUTH_KEY=your_auth_key_here
PORT=8080
DEBUG=true
```

## Usage

### Local Development

Or with debug logging:
```bash
DEBUG=true make run
```

### Testing

Run all tests:
```bash
make test
```

Test the API endpoint:
```bash
make curl-test
```

### Example Request

```bash
curl -X POST http://localhost:8080/book \
  -H "Content-Type: application/json" \
  -H "X-Resy-Auth-Token: ${RESY_AUTH_KEY}" \
  -d '{
    "restaurant_name": "Test Restaurant",
    "date": "2024-02-14",
    "time": "19:00",
    "party_size": 2
  }'
```

## Project Structure

```
go-get-resy-table/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   └── reservation.go
│   └── resy/
│       ├── client.go
│       ├── models.go
│       └── service.go
├── terraform/
│   ├── main.tf
│   ├── variables.tf
│   └── outputs.tf
├── Dockerfile
├── Makefile
└── README.md
```

## Available Make Commands

- `make run`: Run the application
- `make test`: Run all tests
- `make build`: Build the application
- `make clean`: Clean build artifacts
- `make curl-test`: Test the API endpoint
- `make run-debug`: Run with debug logging

## Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| RESY_API_KEY | Resy API Key | Yes | - |
| RESY_AUTH_KEY | Resy Auth Token | Yes | - |
| PORT | Server Port | No | 8080 |
| DEBUG | Enable Debug Logging | No | false |

## Deployment

### Build Docker Image

```bash
docker build -t go-get-resy-table .
```

### Deploy to GCP Cloud Run

```bash
cd terraform
terraform init
terraform apply
```

## TODO

- [ ] Implement actual Resy API integration
- [ ] Add authentication middleware
- [ ] Add rate limiting
- [ ] Implement retry logic
- [ ] Add metrics and monitoring
- [ ] Add more comprehensive tests
- [ ] Add CI/CD pipeline

## Acknowledgments

- Thanks to Resy for their API
- Built with Go and ❤️
