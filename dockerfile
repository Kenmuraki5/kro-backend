# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's working directory
COPY ./cmd/myapp .

copy go.mod go.sum .

# Copy the 'internal' folder to the container
COPY ./internal ./internal

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8060

# Command to run the executable
CMD ["./main"]
