.DEFAULT_GOAL = install

SHELL = /bin/bash -euo pipefail

GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   =
GOPROXY     = direct
GOTAGS      = -tags integration,tools
MODULE      = $(shell go list -m)
PACKAGES    = $(shell go list $(GOTAGS) ./...)
PATHS       = $(shell go list $(GOTAGS) ./... | sed -e "s|$(shell go list -m)/\{0,1\}||g")
TIMEOUT     = 1s
VENDOR      = $(dir $(MODULE))

.PHONY: go-env
go-env:
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(strip $(shell go env GOFLAGS))"
	@echo "GOPRIVATE:   $(strip $(shell go env GOPRIVATE))"
	@echo "GOPROXY:     $(strip $(shell go env GOPROXY))"
	@echo "GOTAGS:      $(GOTAGS)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATHS:       $(strip $(PATHS))"
	@echo "TIMEOUT:     $(TIMEOUT)"
	@echo "VENDOR:      $(VENDOR)"

BINPATH = $(shell dirname $(PWD))/bin

export PATH := $(BINPATH):$(PATH)

.PHONY: tools-env
tools-env:
	@echo "BINPATH:     $(BINPATH)"

.PHONY: brew
brew:
	@brew bundle

.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: update
update:
	@go get $(GOTAGS) -mod= -u

.PHONY: build
build:
	@egg --build
	@mv bin/* $(BINPATH) && rm -rf bin

.PHONY: egg
egg:
	@go get -d -mod= -u github.com/kamilsk/egg
	@go build -mod= -o $(BINPATH)/egg github.com/kamilsk/egg


.PHONY: env
env: go-env tools-env

.PHONY: install
install: egg build
