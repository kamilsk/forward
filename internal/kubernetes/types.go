package kubernetes

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Forwarding defines the format for port forwarding input.
var Forwarding = regexp.MustCompile(`^\d+(?::\d+)?$`)

const (
	entrySeparator = "--"
	portSeparator  = ":"
)

// NewMapping returns port forwarding rules.
func NewMapping(rows ...string) (Mapping, error) {
	mapping := make(Mapping, len(rows))
	for _, row := range rows {
		if row == entrySeparator {
			continue
		}
		if !Forwarding.MatchString(row) {
			return nil, errors.Errorf("expected port forwarding in format [local:]remote, obtained %s", row)
		}
		ports := strings.Split(row, portSeparator)
		converted := make([]int16, 0, len(ports))
		for _, port := range ports {
			value, err := strconv.ParseInt(port, 10, 16)
			if err != nil {
				return nil, errors.Errorf("expected a valid 16-bit port number, obtained %s: %+v", port, err)
			}
			converted = append(converted, int16(value))
		}
		var (
			local  = Local(converted[0])
			remote = Remote(converted[0])
		)
		if len(ports) == 2 {
			remote = Remote(converted[1])
		}
		if _, exists := mapping[local]; exists {
			return nil, errors.Errorf("the local port number %d is already used for forwarding", local)
		}
		mapping[local] = remote
	}
	return mapping, nil
}

type (
	// Pod specifies a fully-qualified pod name.
	Pod string
	// Pods is a list of fully-qualified pod names.
	Pods []Pod
	// Port is a type for port number.
	Port int16
	// Local specifies local ports.
	Local Port
	// Remote specifies remote ports.
	Remote Port
	// Mapping specifies port forwarding rules.
	Mapping map[Local]Remote
)

// Like compares the fully-qualified pod name with the pattern.
func (pod Pod) Like(pattern string) bool {
	return strings.Contains(pod.String(), pattern)
}

// String returns string representation of the fully-qualified pod name.
func (pod Pod) String() string {
	return string(pod)
}

// AsStrings converts a list of fully-qualified pod names into a string list.
func (pods Pods) AsStrings() []string {
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

// String returns string representation of the port number.
func (port Port) String() string {
	return strconv.Itoa(int(port))
}
