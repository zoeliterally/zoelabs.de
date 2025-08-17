# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy all source files
COPY . .

# Download and tidy dependencies
RUN go mod download && go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy static files
COPY --from=builder /app/static ./static

# Copy templates
COPY --from=builder /app/templates ./templates

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
