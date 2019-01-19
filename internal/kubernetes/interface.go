package kubernetes

type (
	// Port is a type for port number.
	Port int16
	// Local specifies local ports.
	Local Port
	// Remote specifies remote ports.
	Remote Port
	// Pod specifies a fully-qualified pod name.
	Pod string
	// Mapping specifies port forwarding rules.
	Mapping map[Local]Remote
)

// Interface defines a top-level behavior of Kubernetes provider (API or CLI).
type Interface interface {
	// Find tries to find pods suitable by the pattern.
	Find(string) []Pod
	// Forward initiates the port forwarding process.
	Forward(Pod, Mapping)
}
