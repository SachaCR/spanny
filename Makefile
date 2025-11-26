.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the service
	go run ./cmd/spanny

.PHONY: build
build: ## Build the service
	go build ./cmd/spanny

.PHONY: install
install: ## Install Spanny in your $GOPATH/bin
	go install ./cmd/spanny

.PHONY: gotest
gotest: ## Test the service with default test runner
	GO_ENV=test go test -v ./...

.PHONY: test
test: ## Test the service with gotestsum
	GO_ENV=test gotestsum --format dots

.PHONY: coverage
coverage: ## Evaluate test coverage
	GO_ENV=test go test -v -coverprofile coverage.out ./... && go tool cover -html=coverage.out
