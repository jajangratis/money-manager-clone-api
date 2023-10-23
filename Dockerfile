# Use the official Go image as the base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code and go.mod/go.sum files to the container
COPY . .

# Install all Go packages listed in go.mod
RUN go get -u ./...

# Build your Go application
RUN go build -o main .

# Expose the port your application will run on
EXPOSE 8000

# Run your application
CMD ["./main"]