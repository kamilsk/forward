package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/forward/internal/kubernetes/api/client"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}
