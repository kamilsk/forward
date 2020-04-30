// +build go1.12

package main

import (
	"fmt"
	"runtime/debug"
)

//nolint:gochecknoinits
func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == unknown {
		version = info.Main.Version
		commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
	}
}
