.PHONY: help build run test migrate docs dev dev-local run-test run-prod \
        test-connection test-connection-dev test-connection-prod \
        migrate-test migrate-prod\
        env-check env-check-test env-check-prod env-setup-test env-setup-prod \
        setup db-reset

help:
	@echo "Available commands:"
	@echo ""
	@echo "🏗️  Build & Run:"
	@echo "  build           - Build the application"
	@echo "  run             - Run the application"
	@echo "  dev             - Run in development mode (alias for dev-local)"
	@echo ""
	@echo "🚀 Environment-specific startup:"
	@echo "  dev-local       - Start in LOCAL development environment"
	@echo "  run-test        - Start in TEST environment"
	@echo "  run-prod        - Start in PRODUCTION environment"
	@echo ""
	@echo "🧪 Testing:"
	@echo "  test            - Run all tests"
	@echo "  test-connection - Test database and Redis connections (test env)"
	@echo "  test-connection-dev - Test connections (development env)"
	@echo "  test-connection-prod - Test connections (production env)"
	@echo ""
	@echo "🗄️  Database:"
	@echo "  migrate         - Run database migrations (development)"
	@echo "  migrate-test    - Run database migrations (test env)"
	@echo "  migrate-prod    - Run database migrations (production env)"
	@echo "  db-reset        - Reset database (migrate)"
	@echo ""
	@echo "🔧 Environment:"
	@echo "  env-check       - Check required environment variables"
	@echo "  env-setup-test  - Setup .env.test file"
	@echo "  env-setup-prod  - Setup .env.production file"
	@echo "🔧 Utilities:"
	@echo "  setup           - Full development setup"


# Build the application
build:
	go build -o bin/server ./cmd/server

# Run the application
run: build
	./bin/server

# Development mode (alias for dev-local)
dev: dev-local

# Environment-specific startup commands
dev-local:
	@echo "🚀 Starting in LOCAL development environment..."
	ENVIRONMENT=development go run ./cmd/server

run-test: env-check-test
	@echo "🧪 Starting in TEST environment..."
	ENVIRONMENT=test go run ./cmd/server

run-prod: env-check-prod
	@echo "🏭 Starting in PRODUCTION environment..."
	ENVIRONMENT=production go run ./cmd/server

# Test database and Redis connections
test-connection-dev:
	@echo "🔌 Testing database and Redis connections (development)..."
	ENVIRONMENT=development go test ./test/connection/... -v

test-connection: env-check-test
	@echo "🔌 Testing database and Redis connections (test)..."
	ENVIRONMENT=test go test ./test/connection/... -v

test-connection-prod: env-check-prod
	@echo "🔌 Testing database and Redis connections (production)..."
	ENVIRONMENT=production go test ./test/connection/... -v

# Run all tests
test:
	go test ./... -v

# Run database migrations
migrate:
	@echo "🗄️ Running database migrations (development)..."
	ENVIRONMENT=development go run ./cmd/migrate

migrate-test: env-check-test
	@echo "🗄️ Running database migrations (test)..."
	ENVIRONMENT=test go run ./cmd/migrate

migrate-prod: env-check-prod
	@echo "🗄️ Running database migrations (production)..."
	ENVIRONMENT=production go run ./cmd/migrate


# Generate API documentation
docs:
	swag init -g ./cmd/server/main.go -o ./docs

# Database commands
db-reset: migrate

# Full development setup
setup: test-connection-dev migrate
	@echo "Development environment setup complete!"

# Environment variable checks
env-check:
	@echo "🔧 Checking environment variables..."
	@echo "Current environment: $${ENVIRONMENT:-development}"

env-check-test:
	@echo "🔧 Checking test environment variables..."
	@if [ -z "$${DB_USER}" ]; then echo "❌ DB_USER is not set"; exit 1; fi
	@if [ -z "$${DB_PASSWORD}" ]; then echo "❌ DB_PASSWORD is not set"; exit 1; fi
	@if [ -z "$${JWT_SECRET}" ]; then echo "❌ JWT_SECRET is not set"; exit 1; fi
	@echo "✅ All required test environment variables are set"

env-check-prod:
	@echo "🔧 Checking production environment variables..."
	@if [ -z "$${DB_USER}" ]; then echo "❌ DB_USER is not set"; exit 1; fi
	@if [ -z "$${DB_PASSWORD}" ]; then echo "❌ DB_PASSWORD is not set"; exit 1; fi
	@if [ -z "$${JWT_SECRET}" ]; then echo "❌ JWT_SECRET is not set"; exit 1; fi
	@if [ -z "$${REDIS_PASSWORD}" ]; then echo "❌ REDIS_PASSWORD is not set"; exit 1; fi
	@echo "✅ All required production environment variables are set"

# Helper commands for environment setup
env-setup-test:
	@echo "🔧 Setting up test environment..."
	@echo "Copy .env.test.example to .env.test and edit it:"
	@echo "  cp .env.test.example .env.test"
	@echo "  # Edit .env.test with your values"
	@echo "  source .env.test"
	@echo "  make test-connection"

env-setup-prod:
	@echo "🔧 Setting up production environment..."
	@echo "Copy .env.production.example to .env.production and edit it:"
	@echo "  cp .env.production.example .env.production"
	@echo "  # Edit .env.production with your values"
	@echo "  source .env.production"
	@echo "  make test-connection-prod"
