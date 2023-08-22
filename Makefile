# Env Variables
# =============================================================================================
ROOT_DIR=$(PWD)
BINARY_PATH=$(ROOT_DIR)/bin/myapp

# set it to "open -a" on MacOS
OPEN=

# keep it empty to run go commands locally
RUN=docker compose run --rm app

# Rules
# =============================================================================================
.DEFAULT: usage

.PHONY: usage
usage:
	@echo '+---------------------------------------------------------------------------------------+'
	@echo '| Event Stream "make" Usage                                                                  |'
	@echo '+---------------------------------------------------------------------------------------+'
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "|- \033[33m%-15s\033[0m -> %s\n", $$1, $$2}'

.PHONY: build
build: ## Builds the local dev images
	@docker compose build app

.PHONY: up
up: ## Prepares local dev environment
	@docker compose up -d --force-recreate app

.PHONY: test
test: ## Run tests
	@$(RUN) go test -race ./...

.PHONY: gencoverage
gencoverage: ## Launches the tests with coverage
	@echo "Generating coverage ..."
	@$(RUN) go test -coverpkg=./... -coverprofile coverage/coverage.out ./...

.PHONY: coverage-html
coverage-html: gencoverage ## Launches the test suite with coverage and open html report
	@$(RUN) go tool cover -html=coverage/coverage.out -o coverage/index.html
	@$(OPEN) firefox coverage/index.html

.PHONY: coverage-text
coverage-text: gencoverage ## Launches the test suite with coverage and s report
	@$(RUN) go tool cover -func coverage/coverage.out

.PHONY: build-prod-image
build-prod-image: ## Builds the production docker image
	@docker compose build app-prod

.PHONY: run-prod-image
run-prod-image: build-prod-image ## Runs the production image locally to troubleshoot it. For local development prefer the dev image ("app" service) and use `make up` instead
	@docker compose up -d --force-recreate app-prod
