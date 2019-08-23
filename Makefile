SHELL       = /bin/bash -euo pipefail
PKGS        = $(shell go list ./... | grep -v vendor)
GO111MODULE = on
GOFLAGS     = -mod=vendor
TIMEOUT     = 1s
BIN         = $(shell basename $(shell pwd))


.DEFAULT_GOAL = test-with-coverage


.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: update
update:
	@go get -mod= -u


.PHONY: format
format:
	@goimports -local $(dirname $(go list -m)) -ungroup -w .

.PHONY: generate
generate:
	@go generate $(PKGS)

.PHONY: refresh
refresh: generate format


.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PKGS)

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PKGS) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PKGS)

.PHONY: test-smoke
test-smoke:
	@echo not implemented yet


.PHONY: build
build:
	@go build -o bin/$(BIN) .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh

.PHONY: install
install:
	@go build -o $(GOPATH)/bin/$(BIN) .
