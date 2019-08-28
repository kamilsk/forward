BIN         = $(shell basename $(shell go list -m))
BINPATH     = $(shell dirname $(PWD))/bin
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

.DEFAULT_GOAL = deps

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


.PHONY: brew
brew:
	@echo nothing to do

.PHONY: deps
deps:
	@go get -m -mod= github.com/kamilsk/egg
	@go build -mod= -o $(BINPATH)/egg github.com/kamilsk/egg
	@go mod tidy && go mod vendor && go mod verify
	@GOFLAGS='-mod=' egg --build
	@mv bin/* $(BINPATH) && rm -rf bin

.PHONY: update
update:
	@go get -mod= -u
