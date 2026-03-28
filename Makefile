.PHONY: all help build run test vet cover lint fmt tidy clean docker-build check vuln install-lint

BINARY := ideal
CMD := ./cmd/$(BINARY)
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w \
	-X 'github.com/stryker/ideal/internal/version.Version=$(VERSION)' \
	-X 'github.com/stryker/ideal/internal/version.Commit=$(COMMIT)' \
	-X 'github.com/stryker/ideal/internal/version.BuildDate=$(BUILD_DATE)'
# Pin matches .github/workflows/ci.yml (golangci-lint-action).
GOLANGCI_LINT_VERSION ?= v2.11.4

all: test build ## Default: test + build

help: ## List targets
	@grep -E '^[a-zA-Z0-9_.-]+:.*?##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-16s %s\n", $$1, $$2}'

check: vet test lint build ## Full local gate (needs golangci-lint)

build: ## Write bin/ideal with embedded git version
	mkdir -p bin
	go build -trimpath -ldflags "$(LDFLAGS)" -o bin/$(BINARY) $(CMD)

run: ## go run ./cmd/ideal
	go run $(CMD)

test: ## Tests with race + shuffle
	go test -count=1 -race -shuffle=on ./...

vet: ## go vet
	go vet ./...

cover: ## HTML coverage report (coverage.html)
	go test -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "wrote coverage.html"

lint: ## golangci-lint (install: brew or make install-lint)
	golangci-lint run ./...

fmt: ## gofmt/goimports via golangci-lint fmt
	golangci-lint fmt ./...

install-lint: ## Install pinned golangci-lint into GOPATH/bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b "$$(go env GOPATH)/bin" $(GOLANGCI_LINT_VERSION)

tidy: ## go mod tidy
	go mod tidy

vuln: ## govulncheck (same pin as CI)
	go run golang.org/x/vuln/cmd/govulncheck@v1.1.4 ./...

clean: ## Remove bin/ and coverage artifacts
	rm -rf bin/ coverage.out coverage.html

docker-build: ## docker build with VERSION / COMMIT / BUILD_DATE build-args
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		-t $(BINARY):local .
