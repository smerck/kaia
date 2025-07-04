.PHONY: api web web-install web-build web-dev test lint lint-go lint-web all

# Run the Go API server
api:
	@echo "Starting Go API server..."
	ANTHROPIC_API_KEY=$$ANTHROPIC_API_KEY go run ./cmd/kaia-api

# Install frontend dependencies
web-install:
	cd web && npm install

# Build the frontend for production
web-build:
	cd web && npm run build

# Run the frontend dev server
web-dev:
	cd web && npm run dev

# Run all Go tests
test:
	go test ./...

# Run Go linting
lint-go:
	@echo "Running Go linter..."
	golangci-lint run

# Run frontend linting
lint-web:
	@echo "Running frontend linter..."
	cd web && npm run lint

# Run all linting
lint: lint-go lint-web

# Run both API and web (dev mode, requires two terminals)
all: api web-dev
	@echo "Open two terminals and run 'make api' and 'make web-dev'"
