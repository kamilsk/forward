package cli

import "io"

// ProcessManager defines behavior of abstract process manager.
type ProcessManager interface {
	// Run starts the process synchronously.
	Run(stdout, stderr io.Writer, command string, args ...string) error
}
