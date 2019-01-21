package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kamilsk/forward/internal/cmd/completion"
	"github.com/kamilsk/forward/internal/cmd/version"
	"github.com/kamilsk/forward/internal/kubernetes"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// New returns new root command.
func New(kubectl kubernetes.Interface) *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			handle(kubectl, args)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				args := strings.Split(strings.TrimSpace(scanner.Text()), " ")
				if len(args) == 1 && args[0] == "" {
					continue
				}
				if len(args) < 2 {
					_, _ = fmt.Fprintln(os.Stderr,
						"Please enter a short or full pod name and its ports in format [local:]remote")
					continue
				}
				process(kubectl, args[0], args[1:]...)
			}
		},
	}
	cmd.AddCommand(completion.New(), version.New())
	return cmd
}

func handle(kubectl kubernetes.Interface, args []string) {
	entries := make([][]string, 0, len(args)/2)
	for len(args) > 0 {
		var name, port string
		entry := make([]string, 0, 4)
		name, args = args[0], args[1:]
		if name == "--" {
			continue
		}
		if kubernetes.Forwarding.MatchString(name) {
			panic("Please provide a pod name first")
		}
		entry = append(entry, name)
		for len(args) > 0 {
			port = args[0]
			if !kubernetes.Forwarding.MatchString(port) {
				break
			}
			args = args[1:]
			entry = append(entry, port)
		}
		if len(entry) == 1 {
			panic(fmt.Sprintf("Please provide the %q pod's ports in format [local:]remote", entry[0]))
		}
		entries = append(entries, entry)
	}
	for _, args := range entries {
		process(kubectl, args[0], args[1:]...)
	}
}

func process(kubectl kubernetes.Interface, pattern string, args ...string) {
	ports, err := kubernetes.NewMapping(args[1:]...)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "An error occurred while parsing the ports: %+v", err)
		return
	}
	options, err := kubectl.Find(pattern)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "An error occurred while finding pods by the pattern %q: %+v\n", pattern, err)
		return
	}
	if len(options) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Pods not found by the pattern %q\n", pattern)
		return
	}
	pod := options.Default()
	if len(options) > 1 {
		choice, err := refine(options.Default().String(), options.AsStrings()...)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "An error occurred while choosing the pod: %+v\n", err)
			return
		}
		pod = kubernetes.Pod(choice)
	}
	go func() {
		if err := kubectl.Forward(pod, ports); err != nil {
			_, _ = fmt.Fprintf(os.Stderr,
				"An error occurred while forwarding ports %+v for the pod %s: %+v\n",
				options, pod, err)
		}
	}()
	time.Sleep(50 * time.Millisecond)
}

func refine(defaults string, options ...string) (string, error) {
	questions := []*survey.Question{
		{
			Name: "pod",
			Prompt: &survey.Select{
				Message: "Choose a pod:",
				Options: options,
				Default: defaults,
			},
		},
	}
	answer := struct {
		Pod string `survey:"pod"`
	}{}
	if err := survey.Ask(questions, &answer); err != nil {
		return answer.Pod, err
	}
	return answer.Pod, nil
}
