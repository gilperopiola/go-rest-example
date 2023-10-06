# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory to the container's workspace
COPY . .

# Build the Go app from the /cmd directory
RUN go build -o main ./cmd/

# Expose port 8040 to the outside world
EXPOSE 8040

# Command to run the executable
CMD ["./main"]
