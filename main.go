package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/google/gops/agent"
	"github.com/kamilsk/forward/internal/cmd"
	executor "github.com/kamilsk/forward/internal/executor/cli"
	provider "github.com/kamilsk/forward/internal/kubernetes/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		cancel()
		signal.Stop(c)
		fmt.Println()
		os.Exit(0)
	}()
	go func() { _ = agent.Listen(agent.Options{ShutdownCleanup: true}) }()
	go func() { _ = http.ListenAndServe(":1234", nil) }()

	if err := cmd.New(provider.New(executor.New(ctx), os.Stderr, os.Stdin)).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
