BIN         = $(shell basename $(shell go list -m))
BINPATH     = $(PWD)/bin
COMMIT      = $(shell git rev-parse --verify HEAD)
DATE        = $(shell date +%Y-%m-%dT%T%Z)
GO111MODULE = on
GOFLAGS     = -mod=vendor
MODULE      = $(shell go list -m)
PACKAGES    = $(shell go list ./...)
PATHS       = $(shell go list ./... | sed -e "s|$(shell go list -m)/\{0,1\}||g")
SHELL       = /bin/bash -euo pipefail
TIMEOUT     = 1s

export PATH := $(BINPATH):$(PATH)

.DEFAULT_GOAL = test-with-coverage

.PHONY: env
env:
	@echo "BIN:         $(BIN)"
	@echo "BINPATH:     $(BINPATH)"
	@echo "COMMIT:      $(COMMIT)"
	@echo "DATE:        $(DATE)"
	@echo "GO111MODULE: $(GO111MODULE)"
	@echo "GOFLAGS:     $(GOFLAGS)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATH:        $(PATH)"
	@echo "PATHS:       $(PATHS)"
	@echo "SHELL:       $(SHELL)"
	@echo "TIMEOUT:     $(TIMEOUT)"


.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: format
format:
	@goimports -local $(dir $(shell go list -m)) -ungroup -w $(PATHS)

.PHONY: generate
generate:
	@go generate $(PACKAGES)

.PHONY: update
update:
	@go get -mod= -u

.PHONY: refresh
refresh: update deps generate format test-with-coverage


.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PACKAGES) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-smoke
test-smoke:
	@echo nothing to do


.PHONY: build
build:
	@go build -o bin/$(BIN) -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)" .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh
