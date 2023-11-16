# Use the official Go image as a parent image
FROM golang:latest

# Set GOPATH to an empty string within the container
ARG GOPATH=

# Copy your Go application source code into the container
COPY ./ ./

# Build the Go application
RUN go build -o main .

# Expose a port (if your application listens on a specific port)
EXPOSE 8080

# Command to run your Go application
CMD ["./main"]