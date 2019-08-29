module tools

go 1.12

replace github.com/pelletier/go-toml => github.com/kamilsk/go-toml v1.4.0-asd-patch

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.0-20190618115843-d350ce7f7a97

require (
	github.com/golang/mock v1.3.1
	github.com/golangci/golangci-lint v1.17.1
	github.com/goreleaser/godownloader v0.0.0-20190803193356-7ef626c90bb6
	golang.org/x/tools v0.0.0-20190829051458-42f498d34c4d
)
