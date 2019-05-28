package cobra

import (
	"github.com/spf13/cobra"
)

const (
	bashFormat = "bash"
	zshFormat  = "zsh"
)

// NewCompletionCommand returns new completion command.
func NewCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Print Bash or Zsh completion",
		Long:  "Print Bash or Zsh completion.",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   bashFormat,
			Short: "Print Bash completion",
			Long:  "Print Bash completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				root := cmd
				for {
					if !root.HasParent() {
						break
					}
					root = root.Parent()
				}
				return root.GenBashCompletion(cmd.OutOrStdout())
			},
		},
		&cobra.Command{
			Use:   zshFormat,
			Short: "Print Zsh completion",
			Long:  "Print Zsh completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				root := cmd
				for {
					if !root.HasParent() {
						break
					}
					root = root.Parent()
				}
				return root.GenZshCompletion(cmd.OutOrStdout())
			},
		},
	)
	return cmd
}
