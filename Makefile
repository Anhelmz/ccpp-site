# CCPP Full-Stack Application Makefile

.PHONY: help install-backend install-frontend install dev-backend dev-frontend dev build clean test-backend test-frontend up stop

FRONTEND_PID_FILE := $(CURDIR)/.frontend_dev.pid
FRONTEND_LOG := $(CURDIR)/frontend-dev.log

# Default target
help:
	@echo "CCPP Full-Stack Application"
	@echo "Available commands:"
	@echo "  install-backend    - Install Go backend dependencies"
	@echo "  install-frontend   - Install Node.js frontend dependencies"
	@echo "  install           - Install all dependencies"
	@echo "  dev-backend       - Start backend development server"
	@echo "  dev-frontend      - Start frontend development server"
	@echo "  dev               - Start both backend and frontend in development mode"
	@echo "  build             - Build both backend and frontend for production"
	@echo "  clean             - Clean build artifacts and dependencies"
	@echo "  test-backend      - Run backend tests"
	@echo "  test-frontend     - Run frontend tests"
	@echo "  up                - Start all Docker services"
	@echo "  stop              - Stop all Docker services"

# Install dependencies
install-backend:
	@echo "Installing Go backend dependencies..."
	cd backend && go mod tidy && go mod download

install-frontend:
	@echo "Installing Node.js frontend dependencies..."
	cd frontend && npm install

install: install-backend install-frontend
	@echo "All dependencies installed successfully!"

# Development servers
dev-backend:
	@echo "Starting backend development server..."
	cd backend && go run cmd/server/main.go

dev-frontend:
	@echo "Starting frontend development server..."
	cd frontend && npm run dev

dev: install
	@echo "Starting development servers..."
	@echo "Backend will be available at http://localhost:8080"
	@echo "Frontend will be available at http://localhost:3000"
	@echo "Press Ctrl+C to stop all servers"
	@trap 'kill %1 %2' INT; \
	$(MAKE) dev-backend & \
	$(MAKE) dev-frontend & \
	wait

# Build for production
build-backend:
	@echo "Building backend for production..."
	cd backend && go build -o ../bin/server cmd/server/main.go

build-frontend:
	@echo "Building frontend for production..."
	cd frontend && npm run build

build: install build-backend build-frontend
	@echo "Production build completed!"

# Testing
test-backend:
	@echo "Running backend tests..."
	cd backend && go test ./...

test-frontend:
	@echo "Running frontend tests..."
	cd frontend && npm run test

# Clean up
clean:
	@echo "Cleaning up build artifacts and dependencies..."
	rm -rf bin/
	cd frontend && rm -rf node_modules/ dist/
	@echo "Clean completed!"

# Database operations
db-setup:
	@echo "Setting up PostgreSQL database..."
	@echo "Starting PostgreSQL with Docker..."
	docker-compose up -d postgres
	@echo "PostgreSQL is now running on localhost:5432"
	@echo "Database name: ccpp_db, User: postgres, Password: password"

db-stop:
	@echo "Stopping PostgreSQL database..."
	docker-compose down postgres

db-reset:
	@echo "Resetting PostgreSQL database..."
	docker-compose down postgres
	docker-compose up -d postgres
	@echo "Database reset completed!"

db-logs:
	@echo "Showing PostgreSQL logs..."
	docker-compose logs -f postgres

# Docker Compose lifecycle
up:
	@echo "Starting all Docker services..."
	docker-compose up -d
	@echo "Installing frontend dependencies..."
	cd frontend && npm install
	@echo "Starting frontend development server..."
	@if [ -f $(FRONTEND_PID_FILE) ] && kill -0 $$(cat $(FRONTEND_PID_FILE)) 2>/dev/null; then \
		echo "Frontend dev server already running with PID $$(cat $(FRONTEND_PID_FILE))"; \
	else \
		cd frontend && nohup npm run dev > $(FRONTEND_LOG) 2>&1 & echo $$! > $(FRONTEND_PID_FILE); \
		echo "Frontend dev server started with PID $$(cat $(FRONTEND_PID_FILE)). Logs: $(FRONTEND_LOG)"; \
	fi

stop:
	@echo "Stopping frontend development server..."
	@if [ -f $(FRONTEND_PID_FILE) ]; then \
		PID=$$(cat $(FRONTEND_PID_FILE)); \
		if kill -0 $$PID 2>/dev/null; then \
			kill $$PID && echo "Killed frontend dev server (PID $$PID)"; \
		else \
			echo "Frontend dev server not running (stale PID $$PID)"; \
		fi; \
		rm -f $(FRONTEND_PID_FILE); \
	else \
		echo "No frontend dev server PID file found"; \
	fi
	@echo "Stopping all Docker services..."
	docker-compose down

# Quick start (install and run in development)
start: install dev

# Production run (build and run)
run: build
	@echo "Starting production servers..."
	@echo "Backend: ./bin/server"
	@echo "Frontend: Serve the frontend/dist directory with a web server"
