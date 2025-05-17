# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Copy dependencies first (for better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the codebase
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose the port specified in the environment
EXPOSE 8080

# Run the binary
CMD ["./main"]
