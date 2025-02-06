.PHONY: default
SHELL := /bin/bash # Use bash syntax
APP_EXECUTABLE="out/transaction-service"
GOPATH=$(shell go env GOPATH)

export GOPATH
default: setup lint test build

setup: --cp-config ## Install all the dependencies
	# Installing go tools for coverage and linting
	go install github.com/axw/gocov/gocov@latest
	go install github.com/t-yuki/gocover-cobertura@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin

--cp-config:
	cp application.yml{.sample,}

lint:  ## Run all the linters
	golangci-lint run

build: ## Build the application
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE) ./cmd/server

run: build ## Run the application
	./$(APP_EXECUTABLE) server

test: ## Run all the tests
	mkdir -p coverage/
	GO111MODULE=on go clean -testcache && go test -race ./... -covermode=atomic -coverprofile=coverage/coverage.out

test.cover: test ## Generate coverage report
	GO111MODULE=on gocov convert coverage/coverage.out | gocov report 2>&1 | tee coverage/coverage.txt

test.report: test.cover ## Generate coverage report in cobertura format
	GO111MODULE=on go tool cover -html coverage/coverage.out -o coverage/coverage.html
	GO111MODULE=on gocover-cobertura < coverage/coverage.out > coverage/coverage.xml

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_.\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'