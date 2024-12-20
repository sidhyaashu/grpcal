# Stage 1: Build the Go binary using golang:alpine image for smaller size
FROM golang:1.23-alpine AS builder

# Install necessary build dependencies
RUN apk add --no-cache build-base

WORKDIR /app

# Copy the Go modules and install dependencies first for caching purposes
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire source code
COPY . ./

# Build the Go application
RUN go build -o bin/server server/main.go

# Stage 2: Create a minimal final image
FROM debian:bookworm-slim

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/server .

# Expose the application port
EXPOSE 8080

# Run the server binary
CMD ["./server"]
