GO_TEST_COVERAGE_MODE     ?= count
GO_TEST_COVERAGE_FILENAME ?= cover.out
PACKAGES                  ?= go list ./... | grep -v vendor
SHELL                     ?= /bin/bash -euo pipefail

.PHONY: format
format:
	@(goimports --ungroup -w ./internal/ ./*.go)

.PHOMY: deps
deps:
	@(go mod tidy)
	@(go mod vendor)
	@(go mod verify)
