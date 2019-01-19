package cli

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/kamilsk/forward/internal/kubernetes"
	"github.com/pkg/errors"
)

const kubectl = "kubectl"

// New returns new instance of Kubernetes provider above CLI.
func New(cli ProcessManager) *provider {
	return &provider{cli}
}

type provider struct {
	cli ProcessManager
}

// Find tries to find pods suitable by the pattern.
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
func (*provider) Forward(kubernetes.Pod, kubernetes.Mapping) {
	panic("implement me")
}

func (provider *provider) pods() ([]kubernetes.Pod, error) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	if err := provider.cli.Run(stdout, stderr, kubectl, "get", "pod"); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)
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
