.PHONY: dev prod build clean help hot

# Default target
help:
	@echo "Available commands:"
	@echo "  make dev     - Start local development server"
	@echo "  make hot     - Start with hot reload (recommended for development)"
	@echo "  make prod    - Start production server with Docker"
	@echo "  make build   - Build production Docker image"
	@echo "  make clean   - Clean up containers and images"
	@echo "  make deps    - Install Go dependencies"
	@echo "  make air     - Install Air for hot reload"

# Local development (requires Go 1.24+ installed)
dev:
	@echo "Starting local development server..."
	go mod download
	go run main.go

# Hot reload development (recommended)
hot:
	@echo "Starting hot reload development server..."
	@echo "Server will automatically restart when files change"
	@echo "Watching: .go, .html, .css files"
	@echo "Press Ctrl+C to stop"
	@if ! command -v air &> /dev/null; then \
		echo "Installing Air for hot reload..."; \
		go install github.com/air-verse/air@latest; \
	fi
	$(shell go env GOPATH)/bin/air

# Production deployment
prod:
	@echo "Starting production server..."
	docker compose up -d web

# Build production image
build:
	@echo "Building production Docker image..."
	docker compose build web

# Clean up
clean:
	@echo "Cleaning up containers and images..."
	docker compose down
	docker rmi zoelabsde-web:latest 2>/dev/null || true
	rm -rf tmp/

# Install dependencies
deps:
	@echo "Installing Go dependencies..."
	go mod download
	go mod tidy

# Install Air for hot reload
air:
	@echo "Installing Air for hot reload..."
	go install github.com/air-verse/air@latest
