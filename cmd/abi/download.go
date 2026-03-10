package abi

import (
	"fmt"

	"github.com/dao-portal/extractor/etherscan"
	evmtypes "github.com/dao-portal/extractor/types/evm"
	fluxclitypes "github.com/dao-portal/flux/cli/types"
	"github.com/spf13/cobra"
)

const (
	dbFlagKey = "db"
)

func NewDownloadCmd() *cobra.Command {
	var dbFlag string
	var chainID string
	cmd := &cobra.Command{
		Use:     "download [contract-address]",
		Short:   "Download the ABI of the provided contract",
		Example: "download 0x1234567890123456789012345678901234567890 --db=my-db --chain-id=1",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("contract address is required")
			}
			contract := args[0]

			return downloadABI(cmd, contract, chainID, dbFlag)
		},
	}
	cmd.Flags().StringVar(&dbFlag, dbFlagKey, "", "database where the ABI will be stored")
	cmd.Flags().StringVar(&chainID, "chain-id", "1", "chain ID of the contract")

	return cmd
}

func downloadABI(cmd *cobra.Command, contract string, chainID string, dbID string) error {
	cliCtx := fluxclitypes.GetCliContext(cmd)
	cfg, err := cliCtx.LoadConfig()
	if err != nil {
		return err
	}

	var abiDB ABIDatabase
	if dbID != "" {
		db, err := cliCtx.DatabasesManager.GetDatabaseByID(cmd.Context(), cfg, dbID)
		if err != nil {
			return err
		}
		abiDBI, ok := db.(ABIDatabase)
		if !ok {
			return fmt.Errorf("database %s is not an ABI database", dbID)
		}
		abiDB = abiDBI
	}

	ctx := cmd.Context()
	etherscanClient := etherscan.GetClient(cliCtx)

	abi, err := etherscanClient.GetContractAbi(ctx, chainID, contract)
	if err != nil {
		return err
	}
	// Print the downloaded ABI to the console
	fmt.Println(string(abi))

	// If the database flag is not set, return
	if abiDB == nil {
		return nil
	}

	// We have a database flag, so we need to store the ABI in the database
	err = abiDB.SaveAbi(ctx, &evmtypes.Abi{
		ContractAddress: contract,
		ChainID:         chainID,
		ABI:             abi,
	})
	if err != nil {
		return fmt.Errorf("save ABI to database: %w", err)
	}

	return nil
}
