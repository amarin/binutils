PROJECT_NAME := "binutils"
PKG := "github.com/amarin/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage coverhtml lint tidy

all: build

## Lint the files
lint:
	@golint -set_exit_status ${PKG_LIST}

## Run unittests
test:
	@go test -short ${PKG_LIST}

## Tidy go dependencies
tidy:
	@go mod tidy

## Get the dependencies
dep: tidy
	@go get -v -d ./...

## Run data race detector
race: dep
	@go test -race -short ${PKG_LIST}

## Run memory sanitizer
msan: dep
	@go test -msan -short ${PKG_LIST}

## Remove previous build
clean:
	@rm -f $(PROJECT_NAME)

## Build the binary file
build: dep
	@go build -i -v $(PKG)

## Display this help screen
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
