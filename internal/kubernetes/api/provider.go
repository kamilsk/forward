package api

import "github.com/kamilsk/forward/internal/kubernetes"

// New returns new instance of Kubernetes provider above API.
func New(api API) *provider {
	return &provider{api}
}

type provider struct {
	api API
}

// Find tries to find pods suitable by the pattern.
func (*provider) Find(string) (kubernetes.Pods, error) {
	panic("implement me")
}

// Forward initiates the port forwarding process.
func (*provider) Forward(kubernetes.Pod, kubernetes.Mapping) error {
	panic("implement me")
}
