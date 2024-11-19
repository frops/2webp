# Use a lightweight Go image based on Alpine
FROM golang:1.23-alpine AS builder

# Install required libraries for libwebp and build tools
RUN apk add --no-cache libwebp-dev build-base

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the application
RUN go build -o /2webp ./cmd/2webp

# Create a smaller runtime image
FROM alpine:latest

# Install runtime libraries for libwebp
RUN apk add --no-cache libwebp

# Copy the compiled binary from the builder
COPY --from=builder /2webp /usr/local/bin/2webp

# Default command to show help
ENTRYPOINT ["2webp"]