# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build main.go -o myapp

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
