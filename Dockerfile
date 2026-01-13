# Build Stage
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Install git if needed for dependencies (optional, but good practice)
RUN apk add --no-cache git

# Copy dependency files first to leverage caching
COPY go.mod ./
# COPY go.sum ./ # Uncomment if you have a go.sum

RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
# -ldflags="-w -s" reduces binary size by stripping debug symbols
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /mi-app main.go

# Final Stage
FROM alpine:latest

# Install CA certificates for HTTPS calls
RUN apk --no-cache add ca-certificates

# creates a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser

# Copy the binary from builder
COPY --from=builder /mi-app .

# Change ownership to non-root user
RUN chown appuser:appgroup /home/appuser/mi-app

# Switch to non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["./mi-app"]