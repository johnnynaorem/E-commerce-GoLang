# Use a proper base image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies required for Go builds
RUN apk add --no-cache git build-base

# Copy all files
COPY . .

COPY service.json /app/service-account.json
ENV GOOGLE_APPLICATION_CREDENTIALS="/app/service-account.json"

# Download dependencies
RUN go mod tidy && go mod vendor

# Expose port
EXPOSE 8080

# Use correct CMD syntax with forward slashes
CMD [ "go", "run", "./cmd/main.go" ]
