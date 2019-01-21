package cli

import "io"

// CLI defines behavior of abstract process manager.
type CLI interface {
	// Run starts the process synchronously.
	Run(stderr, stdout io.Writer, command string, args ...string) error
	// Start starts the process in the background.
	Start(stderr, stdout io.Writer, command string, args ...string) error
}
