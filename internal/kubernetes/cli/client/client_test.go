package client_test

import (
	"context"
	"io"
	"io/ioutil"
	"testing"
	"time"

	. "github.com/kamilsk/forward/internal/kubernetes/cli/client"
	"github.com/stretchr/testify/assert"
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

func TestClient_Start(t *testing.T) {
	tests := []struct {
		name           string
		command        string
		args           []string
		stderr, stdout io.Writer
		assert         func(assert.TestingT, error, ...interface{}) bool
	}{
		{"normal case", "echo", []string{"Hello,", "World!"}, ioutil.Discard, ioutil.Discard, assert.NoError},
		{"error case", "unknown", nil, ioutil.Discard, ioutil.Discard, assert.Error},
		{"background error", "./fixtures/failure.sh", []string{"5"}, ioutil.Discard, ioutil.Discard, assert.NoError},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			tc.assert(t, New(ctx).Start(tc.stderr, tc.stdout, tc.command, tc.args...))
			time.Sleep(10 * time.Millisecond)
			cancel()
			time.Sleep(10 * time.Millisecond)
		})
	}
}
