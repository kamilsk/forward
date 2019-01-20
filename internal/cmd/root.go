package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kamilsk/forward/internal/cmd/completion"
	"github.com/kamilsk/forward/internal/cmd/version"
	"github.com/kamilsk/forward/internal/kubernetes"
	"github.com/kamilsk/forward/internal/kubernetes/cli"
	"github.com/kamilsk/forward/internal/process"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// New returns new root command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			kubectl := cli.New(process.New(ctx), cmd.OutOrStderr(), cmd.OutOrStdout())
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
				pattern, ports := args[0], convert(args[1:])
				{
					options, err := kubectl.Find(pattern)
					if err != nil {
						_, _ = fmt.Fprintf(os.Stderr,
							"Tried to find a pod by the pattern %q: %+v\n",
							pattern, err)
						continue
					}
					if len(options) == 0 {
						_, _ = fmt.Fprintf(os.Stderr,
							"Pod not found by the pattern %q\n",
							pattern)
						continue
					}
					pod := options.Default()
					if len(options) > 1 {
						pod = refine(options)
					}
					go func() {
						if err := kubectl.Forward(pod, ports); err != nil {
							_, _ = fmt.Fprintf(os.Stderr,
								"Tried to forward ports %+v for pod %s: %+v\n",
								options, pod, err)
						}
					}()
					time.Sleep(50 * time.Millisecond)
				}
			}
		},
	}
	cmd.AddCommand(completion.New(), version.New())
	return cmd
}

func handle(kubectl kubernetes.Interface, args []string) {
	is := regexp.MustCompile(`^\d+(?::\d+)?$`)
	entries := make([][]string, 0, len(args)/2)
	for len(args) > 0 {
		var name, port string
		entry := make([]string, 0, 4)
		name, args = args[0], args[1:]
		if name == "--" {
			continue
		}
		if is.MatchString(name) {
			panic("Please provide a pod name first")
		}
		entry = append(entry, name)
		for len(args) > 0 {
			port = args[0]
			if !is.MatchString(port) {
				break
			}
			args = args[1:]
			entry = append(entry, port)
		}
		if len(entry) == 1 {
			panic(fmt.Sprintf(
				"Please provide the %q pod's ports in format [local:]remote",
				entry[0]))
		}
		entries = append(entries, entry)
	}
	for _, args := range entries {
		pattern, ports := args[0], convert(args[1:])
		{
			options, err := kubectl.Find(pattern)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr,
					"Tried to find a pod by the pattern %q: %+v\n",
					pattern, err)
				continue
			}
			if len(options) == 0 {
				_, _ = fmt.Fprintf(os.Stderr,
					"Pod not found by the pattern %q\n",
					pattern)
				continue
			}
			pod := options.Default()
			if len(options) > 1 {
				pod = refine(options)
			}
			go func() {
				if err := kubectl.Forward(pod, ports); err != nil {
					_, _ = fmt.Fprintf(os.Stderr,
						"Tried to forward ports %+v for pod %s: %+v\n",
						ports, pod, err)
				}
			}()
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func convert(raw []string) kubernetes.Mapping {
	forwarding := make(map[kubernetes.Local]kubernetes.Remote)
	for _, row := range raw {
		ports := strings.Split(row, ":")
		if len(ports) != 1 && len(ports) != 2 {
			panic("please provide ports in format [local:]remote")
		}
		converted := make([]int16, 0, len(ports))
		for _, port := range ports {
			value, err := strconv.ParseInt(port, 10, 16)
			if err != nil {
				panic(err)
			}
			converted = append(converted, int16(value))
		}
		if len(ports) == 1 {
			forwarding[kubernetes.Local(converted[0])] = kubernetes.Remote(converted[0])
			continue
		}
		forwarding[kubernetes.Local(converted[0])] = kubernetes.Remote(converted[1])
	}
	return forwarding
}

func refine(options kubernetes.Pods) kubernetes.Pod {
	questions := []*survey.Question{
		{
			Name: "pod",
			Prompt: &survey.Select{
				Message: "Choose a pod:",
				Options: options.AsString(),
				Default: options.Default().String(),
			},
		},
	}
	answer := struct {
		Pod string `survey:"color"`
	}{}
	if err := survey.Ask(questions, &answer); err != nil {
		panic(err)
	}
	return kubernetes.Pod(answer.Pod)
}
