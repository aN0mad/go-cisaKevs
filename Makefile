# Change these variables as necessary.
MAIN_PACKAGE_PATH=cmd/kevcli/main.go
BINARY_NAME=cisaKevs
BIN_DIR=./bin

## help: print this help message
.PHONY: help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## build: build the unix version
.PHONY: build
build:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -ldflags="-s -w" -o=${BIN_DIR}/${BINARY_NAME}.elf ${MAIN_PACKAGE_PATH}

## buildwin: build the windows version
.PHONY: build
buildwin:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o=${BIN_DIR}/${BINARY_NAME}.exe ${MAIN_PACKAGE_PATH}

## vars: print the variables used in this Makefile
.PHONY: vars
vars:
	@echo "MAIN_PACKAGE_PATH: ${MAIN_PACKAGE_PATH}"
	@echo "BINARY_NAME: ${BINARY_NAME}"
	@echo "BIN_DIR: ${BIN_DIR}"

## all: build all applications for unix and windows
.PHONY: all
all:
	make build
	make buildwin

## clean: clean the repository
.PHONY: clean
clean:
	rm -rf ${BIN_DIR}/*
	go clean -cache
	go mod tidy -v
	mv ./logs /tmp/${BINARY_NAME}-logs-$(shell date +%Y%m%d-%H%M%S) || true