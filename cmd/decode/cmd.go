package decode

import (
	"github.com/spf13/cobra"
)

// NewDecodeCmd creates the decode command group.
func NewDecodeCmd(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Decode raw events using contract ABIs",
	}
	cmd.AddCommand(newRawCmd(rootCmd))
	return cmd
}
