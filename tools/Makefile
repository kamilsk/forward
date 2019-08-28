BINPATH     = $(shell dirname $(PWD))/bin
GO111MODULE = on
GOFLAGS     = -mod=
SHELL       = /bin/bash -euo pipefail

export PATH := $(BINPATH):$(PATH)

.DEFAULT_GOAL = deps

.PHONY: env
env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GO111MODULE: $(GO111MODULE)"
	@echo "GOFLAGS:     $(GOFLAGS)"
	@echo "PATH:        $(PATH)"
	@echo "SHELL:       $(SHELL)"


.PHONY: brew
brew:
	@echo nothing to do

.PHONY: deps
deps:
	@go get -m -u github.com/kamilsk/egg
	@go build -o $(BINPATH)/egg github.com/kamilsk/egg
	@go mod tidy && go mod vendor && go mod verify
	@GOFLAGS='-mod=vendor' egg --build
	@mv bin/* $(BINPATH) && rm -rf bin

.PHONY: update
update:
	@go get -u -tags=tools
