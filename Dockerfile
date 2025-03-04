# Dockerfile
FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-get-resy-table ./cmd/main.go

EXPOSE 8080

CMD ["/go-get-resy-table"]