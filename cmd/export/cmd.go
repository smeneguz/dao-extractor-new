package export

import (
	"github.com/spf13/cobra"
)

// NewExportCmd creates the export command group.
func NewExportCmd(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export indexed data to CSV files",
	}
	cmd.AddCommand(newCSVCmd(rootCmd))
	return cmd
}
