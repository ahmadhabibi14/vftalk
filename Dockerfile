FROM golang:1.21-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build main.go

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["./main"]