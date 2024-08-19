include .env

all: build

build:
	@echo "Building..."
	@go build -o ${BINARY} ./cmd/api

run:
	@echo "Running..."
	@go run cmd/api/main.go