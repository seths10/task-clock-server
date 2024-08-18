include .env

all: build

build:
	@echo "Building..."
	@go build -o ${BINARY} cmd/api/main.go

run:
	@echo "Running..."
	@go run cmd/api/main.go