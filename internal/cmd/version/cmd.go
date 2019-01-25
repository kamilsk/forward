package version

import (
	"runtime"

	"github.com/spf13/cobra"
)

// New returns new version command.
func New(commit, date, version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show tool version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf(
				"Version %s (commit: %s, build date: %s, go version: %s, compiler: %s, platform: %s/%s)\n",
				version, commit, date, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
		},
		Version: version,
	}
}
