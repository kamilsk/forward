package process

import (
	"context"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

// New returns new instance of process manager.
func New(ctx context.Context) *manager {
	return &manager{ctx}
}

type manager struct {
	ctx context.Context
}

// Start starts the process asynchronously.
func (manager *manager) Start(stderr, stdout io.Writer, command string, args ...string) {
	cmd := exec.CommandContext(manager.ctx, command, args...)
	cmd.Stderr, cmd.Stdout = stderr, stdout
}

// Run starts the process synchronously.
func (manager *manager) Run(stderr, stdout io.Writer, command string, args ...string) error {
	cmd := exec.CommandContext(manager.ctx, command, args...)
	cmd.Stderr, cmd.Stdout = stderr, stdout
	return errors.Wrapf(cmd.Run(), "tried to starts the specified command %s with args %v", command, args)
}
