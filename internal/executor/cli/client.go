package cli

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/pkg/errors"
)

// New returns new instance of process manager.
func New(ctx context.Context) *manager {
	return &manager{ctx, sync.RWMutex{}, make(map[int]*exec.Cmd)}
}

type manager struct {
	ctx context.Context

	mux sync.RWMutex
	idx map[int]*exec.Cmd
}

// Run starts the process synchronously.
func (manager *manager) Run(stderr, stdout io.Writer, command string, args ...string) error {
	cmd := exec.CommandContext(manager.ctx, command, args...)
	cmd.Stderr, cmd.Stdout = stderr, stdout
	return errors.Wrapf(cmd.Run(), "tried to start the specified command %s with args %+v", command, args)
}

// Start starts the process in the background.
func (manager *manager) Start(stderr, stdout io.Writer, command string, args ...string) error {
	cmd := exec.CommandContext(manager.ctx, command, args...)
	cmd.Stderr, cmd.Stdout = stderr, stdout
	if err := cmd.Start(); err != nil {
		return errors.Wrapf(err, "tried to start the specified command %s with args %+v in background", command, args)
	}
	go manager.store(cmd).wait(cmd)
	return nil
}

func (manager *manager) remove(cmd *exec.Cmd) *manager {
	manager.mux.Lock()
	defer manager.mux.Unlock()
	delete(manager.idx, cmd.Process.Pid)
	return manager
}

func (manager *manager) store(cmd *exec.Cmd) *manager {
	manager.mux.Lock()
	defer manager.mux.Unlock()
	manager.idx[cmd.Process.Pid] = cmd
	return manager
}

func (manager *manager) wait(cmd *exec.Cmd) {
	err := cmd.Wait()
	if err != nil {
		_, _ = fmt.Fprintf(cmd.Stderr, "An error occurred while waiting for the subprocess execution: %+v\n", err)
	}
	manager.remove(cmd)
}
