# Use an official Golang runtime as a parent image
FROM golang:1.22-alpine AS build

# Set the working directory in the container to /app
WORKDIR /app

# Add a .dockerignore file
# COPY ../../.dockerignore ./

# Copy the go.mod and go.sum files from the root directory
COPY ../../go.mod ../../go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the shop-svc directory contents into the container at /app/shop-svc
COPY ../../shop-svc /app/shop-svc

# Set the working directory to the shop-svc directory
WORKDIR /app/shop-svc

# Install any needed dependencies
RUN apk update && apk add --no-cache ca-certificates

# Build the Go app
RUN go build -o /app/shop-svc ./cmd/main.go

# Create a new non-root user with an ID of 1001 and add the user to the wheel group
RUN addgroup -g 1001 appuser && \
    adduser -u 1001 -G wheel -s /bin/false appuser

# Expose the port on which the app runs
EXPOSE 50052

# Set the default command to run the app
CMD ["/app/shop-svc"]
