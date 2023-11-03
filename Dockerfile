# Dockerfile for the Go application
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-matching-service .

CMD ["./order-matching-service"]
