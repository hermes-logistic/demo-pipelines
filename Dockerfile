# Use the official Golang image version 1.21 as the base image
FROM golang:1.21 AS bin-stage

# Set the default shell for the subsequent commands
SHELL ["/bin/bash", "-c"]

# Create a directory for the Go project
RUN mkdir -p /go/src/github.com/tventura-hermes/go-api

# Set the working directory inside the container
WORKDIR /go/src/github.com/tventura-hermes/go-api

# Copy everything from the current directory to the working directory in the container
COPY . .

# Initialize Go modules for the project
RUN go mod init go-api

# Install Gin framework for building the API
RUN go get -u github.com/gin-gonic/gin

# Install necessary packages for testing
RUN go get github.com/stretchr/testify/assert

# Install package for validating structures
RUN go get -u github.com/go-playground/validator/v10

# Install necessary packages for the database (GORM and Postgres driver)
RUN go get -u gorm.io/gorm ; go get -u gorm.io/driver/postgres@v1.5.4

# Install MongoDB driver
RUN go get -u go.mongodb.org/mongo-driver/mongo

# Install UUID package
RUN go get -u github.com/google/uuid

# Install packages necessary for documentation (Swagger)
RUN go get -u github.com/swaggo/swag/cmd/swag ; go install github.com/swaggo/swag/cmd/swag@latest ; go get -u github.com/swaggo/files ; go get -u github.com/swaggo/gin-swagger ; swag init

RUN go mod tidy

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api

# Create a new stage from the debian:11 image for the release
FROM alpine:latest AS release-stage

# Set the working directory inside the container
WORKDIR /

# Copy the built executable from the previous stage to the new stage
COPY --from=bin-stage /go-api /go-api

# Expose port 8080 to the outside world
EXPOSE 8080:8080

# Set the command to run when the container starts
ENTRYPOINT ["/go-api"]
