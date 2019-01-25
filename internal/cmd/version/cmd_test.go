package version_test

import (
	"bytes"
	"testing"

	. "github.com/kamilsk/forward/internal/cmd/version"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		name                  string
		commit, date, version string
		expected              string
	}{
		{"second version", "050fb2f", "Fri Jan 25", "2.0.0", "Version 2.0.0 (commit: 050fb2f, build date: Fri Jan 25"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			cmd := New(tc.commit, tc.date, tc.version)
			cmd.SetOutput(buf)
			cmd.Run(cmd, nil)
			assert.Contains(t, buf.String(), tc.expected)
		})
	}
}
