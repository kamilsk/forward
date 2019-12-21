module github.com/kamilsk/forward/tools

go 1.11

require (
	github.com/golang/mock v1.3.1
	github.com/golangci/golangci-lint v1.21.0
	github.com/goreleaser/godownloader v0.0.0-20190803193356-7ef626c90bb6
	github.com/goreleaser/goreleaser v0.110.0
	github.com/spf13/afero v1.2.2 // indirect
	golang.org/x/tools v0.0.0-20191010075000-0337d82405ff
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.0-20190921135421-dca3d7403570
