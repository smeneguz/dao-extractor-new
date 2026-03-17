package ozgovernor

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/rs/zerolog"

	fluxdatabase "github.com/dao-portal/flux/database"
	fluxmodules "github.com/dao-portal/flux/modules"
	fluxadapter "github.com/dao-portal/flux/modules/adapter"
	fluxnode "github.com/dao-portal/flux/node"
	fluxutils "github.com/dao-portal/flux/utils"

	fluxevmtypes "github.com/dao-portal/extractor/flux/evm/types"
	"github.com/dao-portal/extractor/modules/oz_governor/contracts"
	"github.com/dao-portal/extractor/types"
	"github.com/dao-portal/extractor/utils"
)

// ModuleName is the identifier used in flux configuration.
const ModuleName = "oz-governor"

// Module handles OZ Governor events for one or more DAOs.
type Module struct {
	logger  log.Logger
	db      Database
	evmNode fluxevmtypes.EVMNode

	// addressToDAO maps a governor contract address to its DAOInstance.
	addressToDAO map[common.Address]*DAOInstance

	// daoInstances holds every DAO served by this module instance.
	daoInstances []*DAOInstance
}

var _ fluxadapter.BlockHandleModule[*fluxevmtypes.Block] = &Module{}

// GetName implements fluxmodules.Module.
func (m *Module) GetName() string {
	return ModuleName
}

// Builder is the factory function registered with flux's ModulesManager.
func Builder(
	ctx context.Context,
	fluxDB fluxdatabase.Database,
	node fluxnode.Node,
	rawConfig []byte,
) (fluxmodules.Module, error) {
	db, ok := fluxDB.(Database)
	if !ok {
		return nil, fmt.Errorf("database must implement the oz-governor Database interface")
	}

	evmNode, ok := node.(fluxevmtypes.EVMNode)
	if !ok {
		return nil, fmt.Errorf("node must implement EVMNode interface")
	}

	logger := fluxutils.LoggerFromContext(ctx).With().Str("module", ModuleName).Logger()

	cfg, err := ParseConfig(rawConfig)
	if err != nil {
		return nil, err
	}

	chainID := utils.HexChainIDToDecimal(node.GetChainID())

	mod := &Module{
		logger:       logger,
		db:           db,
		evmNode:      evmNode,
		addressToDAO: make(map[common.Address]*DAOInstance, len(cfg.DAOs)),
	}

	for _, daoCfg := range cfg.DAOs {
		if daoCfg.ChainID != chainID {
			logger.Debug().
				Str("dao", daoCfg.Symbol).
				Str("dao_chain", daoCfg.ChainID).
				Str("node_chain", chainID).
				Msg("skipping DAO: chain mismatch")
			continue
		}

		inst, err := mod.initDAO(ctx, daoCfg, chainID)
		if err != nil {
			return nil, fmt.Errorf("init DAO %s: %w", daoCfg.Symbol, err)
		}

		mod.daoInstances = append(mod.daoInstances, inst)
		logger.Info().
			Str("dao", daoCfg.Symbol).
			Str("governor", inst.Governor.Address.Address).
			Msg("registered OZ Governor DAO")
	}

	if len(mod.daoInstances) == 0 {
		logger.Warn().Str("chain_id", chainID).Msg("no OZ Governor DAOs configured for this chain")
	}

	return fluxadapter.NewBlockHandleAdapter(mod), nil
}

// initDAO creates and persists all the DB entities for a single DAO.
func (m *Module) initDAO(ctx context.Context, daoCfg DAOConfig, chainID string) (*DAOInstance, error) {
	chainName := utils.ChainNameFromID(chainID)

	blockchain := types.NewBlockchain(types.ChainID(chainID), chainName, types.ChainTypeEVM)
	blockchain, err := m.db.InsertBlockchain(ctx, blockchain, true)
	if err != nil {
		return nil, fmt.Errorf("insert blockchain: %w", err)
	}

	dao := types.NewDAO(types.DAOSymbol(strings.ToUpper(daoCfg.Symbol)), daoCfg.Name)
	dao, err = m.db.InsertDAO(ctx, dao, true)
	if err != nil {
		return nil, fmt.Errorf("insert DAO %s: %w", daoCfg.Symbol, err)
	}

	var governorCfg *DAOContractConfig
	var earliestBlock uint64

	for i := range daoCfg.Contracts {
		c := &daoCfg.Contracts[i]
		if earliestBlock == 0 || (c.DeployedAt > 0 && c.DeployedAt < earliestBlock) {
			earliestBlock = c.DeployedAt
		}
		if c.Type == "governor_proxy" || c.Type == "governor" {
			governorCfg = c
		}

		addr := types.NewAddress(c.Address, c.Name, true, types.AddressEncodingTypeHex)
		addr, err = m.db.InsertAddress(ctx, addr, true)
		if err != nil {
			return nil, fmt.Errorf("insert address %s: %w", c.Address, err)
		}
		if err = m.db.AssociateAddressToBlockchain(ctx, addr, blockchain); err != nil {
			return nil, fmt.Errorf("associate address %s to blockchain: %w", c.Address, err)
		}
	}

	if governorCfg == nil {
		return nil, fmt.Errorf("no governor_proxy or governor contract found")
	}

	governorAddr := types.NewAddress(governorCfg.Address, governorCfg.Name, true, types.AddressEncodingTypeHex)
	governorAddr, err = m.db.InsertAddress(ctx, governorAddr, true)
	if err != nil {
		return nil, fmt.Errorf("insert governor address: %w", err)
	}

	governor := NewContractInfo(
		m.evmNode.GetEthClient(),
		contracts.NewOZGovernor(),
		governorAddr,
	)

	ethAddr := common.HexToAddress(governorAddr.Address)
	m.addressToDAO[ethAddr] = nil // placeholder

	creatorKey := operationCreatorKeyForDAO(daoCfg.Symbol)
	ops, err := m.db.GetHeightDeferredOperations(ctx, creatorKey, OperationFetchProposalStatus)
	if err != nil {
		return nil, fmt.Errorf("load deferred operations: %w", err)
	}

	inst := &DAOInstance{
		Config:      daoCfg,
		DAO:         dao,
		Chain:       blockchain,
		Governor:    governor,
		DeferredOps: types.NewHeightDeferredOperations(ops),
		DeployedAt:  earliestBlock,
	}

	m.addressToDAO[ethAddr] = inst

	return inst, nil
}

func operationCreatorKeyForDAO(symbol string) string {
	return fmt.Sprintf("%s-%s", OperationCreatorKey, strings.ToLower(symbol))
}

