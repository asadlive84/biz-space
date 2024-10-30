FROM golang:1.22-alpine AS builder

WORKDIR /build


# Copy the api-gateway directory
COPY api-gateway ./api-gateway
# Copy go.mod and go.sum from the project root
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download



# Build the application
WORKDIR /build/api-gateway
RUN go build -o api-gateway ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/api-gateway/api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]