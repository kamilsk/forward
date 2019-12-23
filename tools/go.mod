module github.com/kamilsk/forward/tools

go 1.11

require (
	github.com/golang/mock v1.3.1
	github.com/golangci/golangci-lint v1.21.0
	github.com/goreleaser/godownloader v0.1.0
	github.com/goreleaser/goreleaser v0.118.2
	golang.org/x/tools v0.2.2
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.1
