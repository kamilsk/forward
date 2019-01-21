package kubernetes

// Interface defines a top-level behavior of Kubernetes provider (API or CLI).
type Interface interface {
	// Find tries to find pods suitable by the pattern.
	Find(string) (Pods, error)
	// Forward initiates the port forwarding process.
	Forward(Pod, Mapping) error
}
