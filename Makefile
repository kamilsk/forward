GO_TEST_COVERAGE_MODE     ?= count
GO_TEST_COVERAGE_FILENAME ?= cover.out
PACKAGES                  ?= go list ./... | grep -v vendor
SHELL                     ?= /bin/bash -euo pipefail

.PHONY: test-with-coverage-profile
test-with-coverage-profile:   #| Runs tests with coverage and collects the result.
                              #| Uses: GO_TEST_COVERAGE_MODE, GO_TEST_COVERAGE_FILENAME, PACKAGES.
	echo 'mode: ${GO_TEST_COVERAGE_MODE}' > '${GO_TEST_COVERAGE_FILENAME}'
	for package in $$($(PACKAGES)); do \
	    go test -covermode '${GO_TEST_COVERAGE_MODE}' \
	            -coverprofile "coverage_$${package##*/}.out" \
	            "$${package}"; \
	    if [ -f "coverage_$${package##*/}.out" ]; then \
	        sed '1d' "coverage_$${package##*/}.out" >> '${GO_TEST_COVERAGE_FILENAME}'; \
	        rm "coverage_$${package##*/}.out"; \
	    fi \
	done
