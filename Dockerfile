# Backend Dockerfile
FROM golang:1.19-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install gcc dan musl-dev untuk build go-sqlite3
RUN apk add --no-cache gcc musl-dev

# Copy source code
COPY . .
# Copy database folder dari root project ke dalam image
COPY ../database ./database

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Set environment variables
ENV GIN_MODE=release

# Expose port 8081
EXPOSE 8081

# Command to run the executable
CMD ["./main"]