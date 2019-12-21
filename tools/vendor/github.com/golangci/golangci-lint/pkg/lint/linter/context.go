package linter

import (
	"golang.org/x/tools/go/packages"

	"github.com/golangci/golangci-lint/internal/pkgcache"
	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/fsutils"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis/load"
	"github.com/golangci/golangci-lint/pkg/logutils"
)

type Context struct {
	// Packages are deduplicated (test and normal packages) packages
	Packages []*packages.Package

	// OriginalPackages aren't deduplicated: they contain both normal and test
	// version for each of packages
	OriginalPackages []*packages.Package

	Cfg       *config.Config
	FileCache *fsutils.FileCache
	LineCache *fsutils.LineCache
	Log       logutils.Log

	PkgCache  *pkgcache.Cache
	LoadGuard *load.Guard
}

func (c *Context) Settings() *config.LintersSettings {
	return &c.Cfg.LintersSettings
}
