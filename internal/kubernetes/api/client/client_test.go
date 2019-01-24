package client_test

import (
	"testing"

	. "github.com/kamilsk/forward/internal/kubernetes/api/client"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}
