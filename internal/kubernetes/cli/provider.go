package cli

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/kamilsk/forward/internal/kubernetes"
	"github.com/pkg/errors"
)

const kubectl = "kubectl"

// New returns new instance of Kubernetes provider above CLI.
func New(cli ProcessManager, stderr, stdout io.Writer) *provider {
	return &provider{cli, stderr, stdout}
}

type provider struct {
	cli            ProcessManager
	stderr, stdout io.Writer
}

// Find tries to find suitable pods by the pattern.
func (provider *provider) Find(pattern string) (kubernetes.Pods, error) {
	pods, err := provider.pods()
	if err != nil {
		return nil, err
	}
	options := make([]kubernetes.Pod, 0, len(pods))
	for _, pod := range pods {
		if pod.Like(pattern) {
			options = append(options, pod)
		}
	}
	return options, nil
}

// Forward initiates the port forwarding process.
func (provider *provider) Forward(pod kubernetes.Pod, ports kubernetes.Mapping) error {
	args := make([]string, 0, len(ports)+1)
	args = append(args, "port-forward", pod.String())
	for local, remote := range ports {
		args = append(args, strings.Join([]string{local.String(), remote.String()}, ":"))
	}
	return provider.cli.Start(provider.stderr, provider.stdout, kubectl, args...)
}

func (provider *provider) pods() ([]kubernetes.Pod, error) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	if err := provider.cli.Run(stdout, stderr, kubectl, "get", "pod"); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	if !scanner.Scan() && scanner.Err() != nil {
		return nil, errors.Wrap(scanner.Err(), "tried to skip header")
	}
	pods := make([]kubernetes.Pod, 0, 10)
	for scanner.Scan() {
		cols := strings.Split(scanner.Text(), " ")
		if len(cols) < 1 {
			return nil, errors.New("unexpected cols count")
		}
		pods = append(pods, kubernetes.Pod(cols[0]))
	}
	return pods, errors.Wrap(scanner.Err(), "tried to read pod names")
}
