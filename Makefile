SHELL := /bin/bash -euo pipefail


.PHONY: deps
deps:
	@(go mod tidy && go mod vendor && go mod verify)

.PHONY: update
update:
	@(go get -u)


.PHONY: format
format:
	@(goimports -ungroup -w .)

.PHONY: generate
generate:
	@(go generate ./...)

.PHONY: refresh
refresh: generate format


.PHONY: test
test:                         #| Runs tests with race.
	@(go test -race -timeout 1s ./...)

.PHONY: test-check
test-check:                   #| Fast runs tests to check their compilation errors.
	@(go test -run=^hack ./...)

.PHONY: test-with-coverage
test-with-coverage:           #| Runs tests with coverage.
	@(go test -cover -timeout 1s  ./...)

.PHONY: test-with-coverage-formatted
test-with-coverage-formatted: #| Runs tests with coverage and formats the result.
	@(go test -cover -timeout 1s  ./... | column -t | sort -r)

.PHONY: test-with-coverage-profile
test-with-coverage-profile:   #| Runs tests with coverage and collects the result.
	@(go test -covermode count -coverprofile cover.out -timeout 1s ./...)

.PHONY: test-example
test-example:                 #| Runs example tests with coverage and collects the result.
	@(go test -covermode count -coverprofile -run=Example -timeout 1s -v example.out ./...)


.PHONY: sync
sync:
	@(git stash && git pull --rebase && git stash pop || true)

.PHONY: upgrade
upgrade: sync update deps refresh test-with-coverage-formatted


.PHONY: install
install:
	@(go build -mod vendor -o forward . && chmod +x forward && echo "$$(pwd)/forward" && ./forward --help)
