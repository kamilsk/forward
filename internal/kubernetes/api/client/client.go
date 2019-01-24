package client

// New returns new instance of Kubernetes API client.
func New() *client {
	return &client{}
}

type client struct{}
