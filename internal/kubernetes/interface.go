package kubernetes

import (
	"strconv"
	"strings"
)

// Interface defines a top-level behavior of Kubernetes provider (API or CLI).
type Interface interface {
	// Find tries to find pods suitable by the pattern.
	Find(string) (Pods, error)
	// Forward initiates the port forwarding process.
	Forward(Pod, Mapping) error
}

type (
	// Port is a type for port number.
	Port int16
	// Local specifies local ports.
	Local Port
	// Remote specifies remote ports.
	Remote Port
	// Pod specifies a fully-qualified pod name.
	Pod string
	// Pods is a list of fully-qualified pod names.
	Pods []Pod
	// Mapping specifies port forwarding rules.
	Mapping map[Local]Remote
)

// String returns string representation of the port number.
func (port Port) String() string {
	return strconv.Itoa(int(port))
}

// Like compares the fully-qualified pod name with the pattern.
func (pod Pod) Like(pattern string) bool {
	return strings.Contains(pod.String(), pattern)
}

// String returns string representation of the fully-qualified pod name.
func (pod Pod) String() string {
	return string(pod)
}

// AsString converts a list of fully-qualified pod names
func (pods Pods) AsString() []string {
	converted := make([]string, 0, len(pods))
	for _, pod := range pods {
		converted = append(converted, pod.String())
	}
	return converted
}

// Default returns a default fully-qualified pod name from the list.
func (pods Pods) Default() Pod {
	var pod Pod
	if len(pods) > 0 {
		pod = pods[0]
	}
	return pod
}
