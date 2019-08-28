//go:generate mockgen -package $GOPACKAGE -destination mock_contract_test.go github.com/kamilsk/forward/internal/kubernetes/cli CLI
//go:generate echo generated at $PWD/$GOFILE ($GOPACKAGE)
package cli_test

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/forward/internal/kubernetes"
	. "github.com/kamilsk/forward/internal/kubernetes/cli"
)

const (
	validPods = `
NAME                                 READY     STATUS    RESTARTS   AGE
prefix-postgresql-7595dd6b9c-jg8mn   1/1       Running   0          6d
prefix-rabbitmq-85d8ff44df-b4lmt     1/1       Running   0          6d
prefix-redis-76bbdf658b-bnc7b        1/1       Running   0          6d
postgresql-c9b6dd5957-nm8gj          1/1       Running   0          6d
rabbitmq-fd44ff8d58-tml4b            1/1       Running   0          6d
redis-b856fdbb67-b7cnb               1/1       Running   0          6d`
	invalidPods = `
NAME                                 READY     STATUS    RESTARTS   AGE
    `
)

func TestProvider_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		pattern string
		cli     func() CLI
		assert  func(kubernetes.Pods, error)
	}{
		{"multiple options", "postgres",
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().
					Run(gomock.Any(), gomock.Any(), "kubectl", "get", "pod").
					Return(nil).
					Do(func(stderr, stdout io.Writer, command string, args ...string) {
						_, _ = stdout.Write([]byte(strings.TrimLeft(validPods, "\n")))
					})
				return mock
			},
			func(pods kubernetes.Pods, err error) {
				assert.NoError(t, err)
				assert.Equal(t, kubernetes.Pods{"prefix-postgresql-7595dd6b9c-jg8mn", "postgresql-c9b6dd5957-nm8gj"}, pods)
			},
		},
		{"single option", "prefix-postgres",
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().
					Run(gomock.Any(), gomock.Any(), "kubectl", "get", "pod").
					Return(nil).
					Do(func(stderr, stdout io.Writer, command string, args ...string) {
						_, _ = stdout.Write([]byte(strings.TrimLeft(validPods, "\n")))
					})
				return mock
			},
			func(pods kubernetes.Pods, err error) {
				assert.NoError(t, err)
				assert.Equal(t, kubernetes.Pods{"prefix-postgresql-7595dd6b9c-jg8mn"}, pods)
			},
		},
		{"no options", "unknown-pod",
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().
					Run(gomock.Any(), gomock.Any(), "kubectl", "get", "pod").
					Return(nil).
					Do(func(stderr, stdout io.Writer, command string, args ...string) {
						_, _ = stdout.Write([]byte(strings.TrimLeft(validPods, "\n")))
					})
				return mock
			},
			func(pods kubernetes.Pods, err error) {
				assert.NoError(t, err)
				assert.Empty(t, pods)
			},
		},
		{"run error", "postgres",
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().Run(gomock.Any(), gomock.Any(), "kubectl", "get", "pod").Return(errors.New("test"))
				return mock
			},
			func(pods kubernetes.Pods, err error) {
				assert.Error(t, err)
				assert.Empty(t, pods)
			},
		},
		{"scan error", "postgres",
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().
					Run(gomock.Any(), gomock.Any(), "kubectl", "get", "pod").
					Return(nil).
					Do(func(stderr, stdout io.Writer, command string, args ...string) {
						_, _ = stdout.Write([]byte(strings.Repeat(" ", bufio.MaxScanTokenSize)))
					})
				return mock
			},
			func(pods kubernetes.Pods, err error) {
				assert.Error(t, err)
				assert.Empty(t, pods)
			},
		},
		{"unexpected cols count", "postgres",
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().
					Run(gomock.Any(), gomock.Any(), "kubectl", "get", "pod").
					Return(nil).
					Do(func(stderr, stdout io.Writer, command string, args ...string) {
						_, _ = stdout.Write([]byte(strings.TrimLeft(invalidPods, "\n")))
					})
				return mock
			},
			func(pods kubernetes.Pods, err error) {
				assert.Error(t, err)
				assert.Empty(t, pods)
			},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.assert(New(tc.cli(), ioutil.Discard, ioutil.Discard).Find(tc.pattern))
		})
	}
}

func TestProvider_Forward(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name   string
		pod    kubernetes.Pod
		ports  kubernetes.Mapping
		cli    func() CLI
		assert func(assert.TestingT, error, ...interface{}) bool
	}{
		{"normal case", "postgresql-c9b6dd5957-nm8gj", kubernetes.Mapping{5432: 5432},
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().
					Run(gomock.Any(), gomock.Any(), "kubectl", "port-forward", "postgresql-c9b6dd5957-nm8gj", "5432:5432").
					Return(nil)
				return mock
			},
			assert.NoError,
		},
		{"error case", "postgresql-c9b6dd5957-nm8gj", kubernetes.Mapping{5432: 5432},
			func() CLI {
				mock := NewMockCLI(ctrl)
				mock.EXPECT().
					Run(gomock.Any(), gomock.Any(), "kubectl", "port-forward", "postgresql-c9b6dd5957-nm8gj", "5432:5432").
					Return(errors.New("test"))
				return mock
			},
			assert.Error,
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			tc.assert(t, New(tc.cli(), ioutil.Discard, ioutil.Discard).Forward(tc.pod, tc.ports))
		})
	}
}
