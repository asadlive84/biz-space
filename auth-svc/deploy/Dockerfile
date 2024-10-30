# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum from the project root
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire auth-svc directory
COPY ../auth-svc .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ../../auth-svc

# Final stage
FROM scratch
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]