# Use an official Go runtime as a parent image
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o bin/task-manager ./Delivery/main.go

# Use a minimal base image to run the Go app
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/task-manager .

# Expose port 8081 to the outside world
EXPOSE 8081

# Run the binary program produced by `go build`
CMD ["./task-manager"]
