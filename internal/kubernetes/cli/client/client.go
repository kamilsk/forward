package client

import (
	"context"
	"io"
	"os/exec"
	"sync"

	"github.com/pkg/errors"
)

// New returns new instance of process manager.
func New(ctx context.Context) *client {
	return &client{ctx, sync.RWMutex{}, make(map[int]*exec.Cmd)}
}

type client struct {
	ctx context.Context

	mux sync.RWMutex
	idx map[int]*exec.Cmd
}

// Run starts the process synchronously.
func (client *client) Run(stderr, stdout io.Writer, command string, args ...string) error {
	cmd := exec.CommandContext(client.ctx, command, args...)
	cmd.Stderr, cmd.Stdout = stderr, stdout
	return errors.Wrapf(cmd.Run(), "tried to start the specified command %s with args %+v", command, args)
}
