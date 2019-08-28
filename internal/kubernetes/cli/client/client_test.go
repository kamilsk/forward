package client_test

import (
	"context"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/forward/internal/kubernetes/cli/client"
)

func TestClient_Run(t *testing.T) {
	tests := []struct {
		name           string
		command        string
		args           []string
		stderr, stdout io.Writer
		assert         func(assert.TestingT, error, ...interface{}) bool
	}{
		{"normal case", "echo", []string{"Hello,", "World!"}, ioutil.Discard, ioutil.Discard, assert.NoError},
		{"error case", "unknown", nil, ioutil.Discard, ioutil.Discard, assert.Error},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.assert(t, New(context.Background()).Run(tc.stderr, tc.stdout, tc.command, tc.args...))
		})
	}
}
