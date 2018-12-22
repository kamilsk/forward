package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"time"

	survey "gopkg.in/AlecAivazis/survey.v1"
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
	flag.Parse()
	handle(flag.Args())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		args := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(args) == 1 && args[0] == "" {
			continue
		}
		if len(args) < 2 {
			_, _ = fmt.Fprintln(os.Stderr, "Please enter a short or full pod name and its ports in format [local:]remote")
			continue
		}
		pod, ports := args[0], convert(args[1:])
		options := find(pod)
		if len(options) == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "Pod not found by criteria %q\n", pod)
			continue
		}
		pod = options[0]
		if len(options) > 1 {
			pod = define(options)
		}
		go forward(pod, ports)
		time.Sleep(50 * time.Millisecond)
	}
}

func handle(args []string) {
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
			panic(fmt.Sprintf("Please provide the %q pod's ports in format [local:]remote", entry[0]))
		}
		entries = append(entries, entry)
	}
	for _, args := range entries {
		pod, ports := args[0], convert(args[1:])
		options := find(pod)
		if len(options) == 0 {
			panic(fmt.Sprintf("Pod not found by criteria %q", pod))
		}
		pod = options[0]
		if len(options) > 1 {
			pod = define(options)
		}
		go forward(pod, ports)
		time.Sleep(50 * time.Millisecond)
	}
}

func convert(raw []string) map[int16]int16 {
	forwarding := make(map[int16]int16)
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
			forwarding[converted[0]] = converted[0]
			continue
		}
		forwarding[converted[0]] = converted[1]
	}
	return forwarding
}

func define(options []string) string {
	questions := []*survey.Question{
		{
			Name: "pod",
			Prompt: &survey.Select{
				Message: "Choose a pod:",
				Options: options,
				Default: options[0],
			},
		},
	}
	answer := struct {
		Pod string `survey:"color"`
	}{}
	if err := survey.Ask(questions, &answer); err != nil {
		panic(err)
	}
	return answer.Pod
}

func find(name string) []string {
	options := make([]string, 0, 4)
	for _, pod := range pods() {
		if strings.Contains(pod, name) {
			options = append(options, pod)
		}
	}
	return options
}

func forward(pod string, ports map[int16]int16) {
	args := make([]string, 0, len(ports)+1)
	args = append(args, "port-forward", pod)
	for local, remote := range ports {
		args = append(args, strings.Join([]string{strconv.Itoa(int(local)), strconv.Itoa(int(remote))}, ":"))
	}
	cmd := exec.Command("kubectl", args...)
	cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func pods() []string {
	buf := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)
	cmd := exec.Command("kubectl", "get", "pod")
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	_ = scanner.Scan() // skip header "NAME READY STATUS RESTARTS AGE"
	pods := make([]string, 0, 10)
	for scanner.Scan() {
		cols := strings.Split(scanner.Text(), " ")
		if len(cols) < 1 {
			panic("unexpected cols count")
		}
		pods = append(pods, cols[0])
	}
	return pods
}
