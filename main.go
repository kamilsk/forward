package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"

	"github.com/google/gops/agent"
	"go.octolab.org/toolkit/cli/cobra"

	"github.com/kamilsk/forward/internal/cmd"
	"github.com/kamilsk/forward/internal/kubernetes/cli"
	"github.com/kamilsk/forward/internal/kubernetes/cli/client"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		cancel()
		time.Sleep(50 * time.Millisecond) // add a possibility of shutting down subprocesses
		// TODO make it better using a callback

		signal.Stop(c)
		fmt.Println()
		os.Exit(0)
	}()
	go func() { _ = agent.Listen(agent.Options{ShutdownCleanup: true}) }()
	go func() { _ = http.ListenAndServe(":1234", nil) }()

	root := cmd.New(cli.New(client.New(ctx), os.Stderr, os.Stdout))
	root.AddCommand(cobra.NewCompletionCommand(), cobra.NewVersionCommand(version, date, commit))
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
