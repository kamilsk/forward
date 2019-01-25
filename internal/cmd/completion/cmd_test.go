package completion_test

import (
	"bytes"
	"testing"

	. "github.com/kamilsk/forward/internal/cmd/completion"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCompletion(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{"Bash", "bash", "# bash completion for test"},
		{"Zsh", "zsh", "#compdef test"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			cmd := &cobra.Command{Use: "test"}
			cmd.AddCommand(New())
			assert.Len(t, cmd.Commands(), 1)
			cmd = cmd.Commands()[0]
			cmd.SetOutput(buf)
			assert.NoError(t, cmd.Flag("format").Value.Set(tc.format))
			assert.NoError(t, cmd.RunE(cmd, nil))
			assert.Contains(t, buf.String(), tc.expected)
		})
	}
}
