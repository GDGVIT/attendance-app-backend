# Start from golang base image
FROM golang:1.20-alpine as builder

# Install git.
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy only necessary files
COPY go.mod go.sum ./
COPY . .

# Build the Go application
RUN go build -o app

# Set the DEBUG environment variable to False in production
ENV DEBUG=False

# Set the entry point for your application
CMD ["./app"]
