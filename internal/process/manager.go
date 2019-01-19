package process

import (
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

// New returns new instance of process manager.
func New() *manager {
	return &manager{}
}

type manager struct{}

// Run starts the process synchronously.
func (*manager) Run(stdout, stderr io.Writer, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout, cmd.Stderr = stdout, stderr
	return errors.Wrapf(cmd.Run(), "tried to starts the specified command %s with args %v", command, args)
}
