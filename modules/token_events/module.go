package tokenevents

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
	"github.com/dao-portal/extractor/types"
	"github.com/dao-portal/extractor/utils"
)

// ModuleName is the identifier used in flux configuration.
const ModuleName = "token-events"

// DAOInstance holds everything needed to process token events for one DAO.
type DAOInstance struct {
	Config     DAOConfig
	DAO        *types.DAO
	Chain      *types.Blockchain
	DeployedAt uint64 // earliest deployment block among this DAO's token contracts
}

// Module captures ERC20 Transfer, DelegateChanged, and DelegateVotesChanged
// events from governance token contracts.
type Module struct {
	logger log.Logger
	db     Database

	// addressToDAO maps a token contract address to its DAOInstance.
	addressToDAO map[common.Address]*DAOInstance

	// addressToTokenAddr maps a go-ethereum address to its hex string
	// for storage without repeated conversions.
	addressToTokenStr map[common.Address]string

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
		return nil, fmt.Errorf("database must implement the token-events Database interface")
	}

	logger := fluxutils.LoggerFromContext(ctx).With().Str("module", ModuleName).Logger()

	cfg, err := ParseConfig(rawConfig)
	if err != nil {
		return nil, err
	}

	chainID := utils.HexChainIDToDecimal(node.GetChainID())

	mod := &Module{
		logger:            logger,
		db:                db,
		addressToDAO:      make(map[common.Address]*DAOInstance),
		addressToTokenStr: make(map[common.Address]string),
	}

	for _, daoCfg := range cfg.DAOs {
		// Only initialise DAOs that belong to this indexer's chain.
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
			Int("contracts", len(daoCfg.Contracts)).
			Msg("registered token-events DAO")
	}

	if len(mod.daoInstances) == 0 {
		logger.Warn().Str("chain_id", chainID).Msg("no token-events DAOs configured for this chain")
	}

	return fluxadapter.NewBlockHandleAdapter(mod), nil
}

// initDAO creates and persists all the DB entities for a single DAO.
func (m *Module) initDAO(ctx context.Context, daoCfg DAOConfig, chainID string) (*DAOInstance, error) {
	chainName := utils.ChainNameFromID(chainID)

	// Upsert blockchain record.
	blockchain := types.NewBlockchain(types.ChainID(chainID), chainName, types.ChainTypeEVM)
	blockchain, err := m.db.InsertBlockchain(ctx, blockchain, true)
	if err != nil {
		return nil, fmt.Errorf("insert blockchain: %w", err)
	}

	// Upsert DAO record.
	dao := types.NewDAO(types.DAOSymbol(strings.ToUpper(daoCfg.Symbol)), daoCfg.Name)
	dao, err = m.db.InsertDAO(ctx, dao, true)
	if err != nil {
		return nil, fmt.Errorf("insert DAO %s: %w", daoCfg.Symbol, err)
	}

	var earliestBlock uint64

	for i := range daoCfg.Contracts {
		c := &daoCfg.Contracts[i]
		if earliestBlock == 0 || (c.DeployedAt > 0 && c.DeployedAt < earliestBlock) {
			earliestBlock = c.DeployedAt
		}

		// Persist every contract address.
		addr := types.NewAddress(c.Address, c.Name, true, types.AddressEncodingTypeHex)
		addr, err = m.db.InsertAddress(ctx, addr, true)
		if err != nil {
			return nil, fmt.Errorf("insert address %s: %w", c.Address, err)
		}
		if err = m.db.AssociateAddressToBlockchain(ctx, addr, blockchain); err != nil {
			return nil, fmt.Errorf("associate address %s to blockchain: %w", c.Address, err)
		}

		// Register address in routing map.
		ethAddr := common.HexToAddress(c.Address)
		m.addressToTokenStr[ethAddr] = strings.ToLower(c.Address)
	}

	inst := &DAOInstance{
		Config:     daoCfg,
		DAO:        dao,
		Chain:      blockchain,
		DeployedAt: earliestBlock,
	}

	// Register all contract addresses in the module's routing map.
	for i := range daoCfg.Contracts {
		ethAddr := common.HexToAddress(daoCfg.Contracts[i].Address)
		m.addressToDAO[ethAddr] = inst
	}

	return inst, nil
}

