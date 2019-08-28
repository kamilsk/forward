//go:generate mockgen -package $GOPACKAGE -destination mock_contract_test.go github.com/kamilsk/forward/internal/kubernetes/api API
//go:generate echo generated at $PWD/$GOFILE ($GOPACKAGE)
package api_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/forward/internal/kubernetes"
	. "github.com/kamilsk/forward/internal/kubernetes/api"
)

func TestProvider_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		pattern string
		api     func() API
	}{
		{"not implemented", "postgres",
			func() API {
				mock := NewMockAPI(ctrl)
				return mock
			},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Panics(t, func() { _, _ = New(tc.api()).Find(tc.pattern) })
		})
	}
}

func TestProvider_Forward(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name  string
		pod   kubernetes.Pod
		ports kubernetes.Mapping
		api   func() API
	}{
		{"not implemented", "postgres", kubernetes.Mapping{5432: 5432},
			func() API {
				mock := NewMockAPI(ctrl)
				return mock
			},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Panics(t, func() { _ = New(tc.api()).Forward(tc.pod, tc.ports) })
		})
	}
}
