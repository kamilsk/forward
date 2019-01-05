package main

import (
	"fmt"
	"os"
	"os/signal"

	grmon "github.com/bcicen/grmon/agent"
	"github.com/google/gops/agent"
	"github.com/kamilsk/forward/internal/cmd"
)

func main() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		signal.Stop(c)
		fmt.Println()
		os.Exit(0)
	}()
	go func() { _ = agent.Listen(agent.Options{ShutdownCleanup: true}) }()
	go func() { grmon.Start() }()

	if err := cmd.New().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
