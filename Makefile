GO111MODULES=on
BIN = $(CURDIR)/bin

APP = ./cmd/spork

.PHONY: build
## build: cleans and builds project
build: server client

.PHONY: mkdir_bin
## mkdir_bin: makes bin directory
mkdir_bin:
	@mkdir -p bin/

.PHONY: run
## run: runs project
run:
	go run -race ${APP} $@

.PHONY: clean
## clean: cleans build directory
clean:
	@echo "Cleaning"
	@go clean

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: server
server: clean mkdir_bin
	@echo "Building..."
	@go build -o ${BIN}/ ./cmd/server

.PHONY: client
client: clean mkdir_bin
	@echo "Building..."
	@go build -o ${BIN}/ ./cmd/client
