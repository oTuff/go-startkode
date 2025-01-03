# Read .env variables for use with commands.
include .env
.EXPORT_ALL_VARIABLES:

# Declare all targets as phony to avoid conflicts with file names
.PHONY:*

# Show available targets
default:
	@echo -e "\033[1;32mAvailable targets:\033[0m"
	@grep "^[^ ]*:#" Makefile | awk -F':' '{printf "%-20s \033[1;34m%s\033[0m\n", $$1, $$2}'

install-tools:# Install dev tools
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/go-gremlins/gremlins/cmd/gremlins@latest

run:# Run with live reload
	air --build.cmd "go build -o tmp/" #api" --build.bin "./bin/api"

build:# Build application
	go build -o bin/

docs:# Generate swagger documentation
	swag init

sqlc:# Generate slqc boilerplate
	sqlc generate

migrate-new:# Create new sql migration
	@read -p "Enter migration name: " NAME; \
	goose create -dir db/migrations $$NAME sql

migrate-up:# Migrate the db
	@goose -dir db/migrations postgres "$(DBSTRING)" up

migrate-down:# Roll back the db
	@goose -dir db/migrations postgres "$(DBSTRING)" down

test:# Run tests
	go test ./...

test-verbose:# Run tests with verbose flag
	go test ./... -v

test-cover:# Show test coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

test-cover-html:# Show test coverage in the browser
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-mutation:# Run mutation testing with gremlins
	gremlins r
