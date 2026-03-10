package pricefeed

import (
	"github.com/spf13/cobra"
)

// NewPriceFeedCmd creates the price-feed command group.
func NewPriceFeedCmd(rootCmd *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price-feed",
		Short: "Fetch and store historical token prices",
	}
	cmd.AddCommand(newFetchCmd(rootCmd))
	return cmd
}
