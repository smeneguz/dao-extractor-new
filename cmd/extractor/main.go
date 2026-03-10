package main

import (
	"fmt"
	"os"

	fluxcli "github.com/dao-portal/flux/cli"
	fluxclitypes "github.com/dao-portal/flux/cli/types"
	fluxpostgresql "github.com/dao-portal/flux/database/postgresql"
	"gopkg.in/yaml.v3"

	"github.com/dao-portal/extractor/cmd/abi"
	"github.com/dao-portal/extractor/cmd/decode"
	"github.com/dao-portal/extractor/cmd/export"
	"github.com/dao-portal/extractor/cmd/pricefeed"
	"github.com/dao-portal/extractor/database/postgresql"
	"github.com/dao-portal/extractor/etherscan"
	evmfallback "github.com/dao-portal/extractor/flux/evm/node/fallback"
	evmnode "github.com/dao-portal/extractor/flux/evm/node/rpc"
	governorbravo "github.com/dao-portal/extractor/modules/governor_bravo"
	ozgovernor "github.com/dao-portal/extractor/modules/oz_governor"
	rawevents "github.com/dao-portal/extractor/modules/raw_events"
	tokenevents "github.com/dao-portal/extractor/modules/token_events"
	"github.com/dao-portal/extractor/types"
)

func main() {
	cliCtx := fluxclitypes.NewCliContext("dao-portal")

	// Register supported nodes
	cliCtx.NodesManager.RegisterNode(evmnode.NodeType, evmnode.NodeBuilder)
	cliCtx.NodesManager.RegisterNode(evmfallback.NodeType, evmfallback.NodeBuilder)

	// Register our custom database
	cliCtx.DatabasesManager.RegisterDatabase(fluxpostgresql.DatabaseType, postgresql.Builder)

	// Register our custom modules
	cliCtx.ModulesManager.RegisterModule(governorbravo.ModuleName, governorbravo.Builder)
	cliCtx.ModulesManager.RegisterModule(ozgovernor.ModuleName, ozgovernor.Builder)
	cliCtx.ModulesManager.RegisterModule(rawevents.ModuleName, rawevents.Builder)
	cliCtx.ModulesManager.RegisterModule(tokenevents.ModuleName, tokenevents.Builder)

	// Register function to load our config
	cliCtx.WithRawConfigLoadedHook(loadCustomConfig)

	// Start the application
	rooCmd := fluxcli.NewDefaultIndexerCLI(cliCtx)

	// Add our custom commands
	rooCmd.AddCommand(abi.NewABICmd(rooCmd))
	rooCmd.AddCommand(pricefeed.NewPriceFeedCmd(rooCmd))
	rooCmd.AddCommand(export.NewExportCmd(rooCmd))
	rooCmd.AddCommand(decode.NewDecodeCmd(rooCmd))

	if err := rooCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func loadCustomConfig(cliCtx *fluxclitypes.CliContext, rawConfig []byte) error {
	var config types.Config
	err := yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		return fmt.Errorf("unmarshal custom config: %w", err)
	}

	// Validate the configurations
	err = config.Validate()
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	eterscanClient := etherscan.NewClient(&config.EtherscanConfig)
	etherscan.InjectClient(cliCtx, eterscanClient)

	return nil
}
