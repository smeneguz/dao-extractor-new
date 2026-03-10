package abi

import "github.com/spf13/cobra"

func NewABICmd(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "abi",
		Short: "ABI related commands",
	}

	cmd.AddCommand(NewDownloadCmd())
	cmd.AddCommand(NewFetchAllCmd())

	return cmd
}
