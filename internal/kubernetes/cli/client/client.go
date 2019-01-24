package client

import (
	"context"
	"fmt"
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

// Start starts the process in the background.
func (client *client) Start(stderr, stdout io.Writer, command string, args ...string) error {
	cmd := exec.CommandContext(client.ctx, command, args...)
	cmd.Stderr, cmd.Stdout = stderr, stdout
	if err := cmd.Start(); err != nil {
		return errors.Wrapf(err, "tried to start the specified command %s with args %+v in background", command, args)
	}
	go client.store(cmd).wait(cmd)
	return nil
}

func (client *client) remove(cmd *exec.Cmd) *client {
	client.mux.Lock()
	defer client.mux.Unlock()
	delete(client.idx, cmd.Process.Pid)
	return client
}

func (client *client) store(cmd *exec.Cmd) *client {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.idx[cmd.Process.Pid] = cmd
	return client
}

func (client *client) wait(cmd *exec.Cmd) {
	err := cmd.Wait()
	if err != nil {
		_, _ = fmt.Fprintf(cmd.Stderr, "An error occurred while waiting for the subprocess execution: %+v\n", err)
	}
	client.remove(cmd)
}
