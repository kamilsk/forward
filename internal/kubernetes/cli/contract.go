package cli

import "io"

// ProcessManager defines behavior of abstract process manager.
type ProcessManager interface {
	// Run starts the process synchronously.
	Run(stderr, stdout io.Writer, command string, args ...string) error
	// Start starts the process in the background.
	Start(stderr, stdout io.Writer, command string, args ...string) error
}
