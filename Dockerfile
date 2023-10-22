# Use the official Golang image as a base image for the builder stage
FROM golang:1.19-alpine AS builder

# Install gcc and other necessary build tools
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory to the container's workspace
COPY . .

# Run unit tests. IDK why this logic works, but it works
RUN if [ "$fast" = "true" ]; then \
    go test ./...; \
fi

# Build the Go app from the /cmd directory
RUN go build -o go-rest-example ./cmd/

# Start a new stage with a lightweight base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage to the current stage
COPY --from=builder /app/go-rest-example /app/go-rest-example
COPY --from=builder /app/.env /app

# Expose port 8040 to the outside world
EXPOSE 8040

# Command to run the executable
CMD ["./go-rest-example"]
