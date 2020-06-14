GO111MODULES=on
BIN = $(CURDIR)/bin

APP = ./cmd/spork

.PHONY: build
## build: cleans and builds project
build: clean
	@echo "Building..."
	@mkdir -p bin/
	@go build -o ${BIN}/ ${APP}

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
