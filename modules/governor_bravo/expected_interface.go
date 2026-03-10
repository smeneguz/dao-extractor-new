package governorbravo

import (
	"context"
	"math/big"

	"github.com/dao-portal/extractor/types"
	evmtypes "github.com/dao-portal/extractor/types/evm"
)

// Database defines the persistence operations required by this module.
type Database interface {
	StoreHeightDeferredOperation(ctx context.Context, op *types.HeightDeferredOperation) error
	GetHeightDeferredOperations(ctx context.Context, creatorKey string, opType string) ([]types.HeightDeferredOperation, error)
	RemoveHeightDeferredOperations(ctx context.Context, op *types.HeightDeferredOperation) error

	SaveAbi(ctx context.Context, abi *evmtypes.Abi) error
	GetAbi(ctx context.Context, chainID string, contractAddress string) (*evmtypes.Abi, error)

	InsertBlockchain(ctx context.Context, blockchain *types.Blockchain, allowConflict bool) (*types.Blockchain, error)
	InsertAddress(ctx context.Context, address *types.Address, allowConflict bool) (*types.Address, error)
	AssociateAddressToBlockchain(ctx context.Context, address *types.Address, blockchain *types.Blockchain) error

	InsertDAO(ctx context.Context, dao *types.DAO, allowConflict bool) (*types.DAO, error)

	StoreProposal(ctx context.Context, proposal *types.Proposal) (*types.Proposal, error)
	GetProposalByProposalID(ctx context.Context, daoID types.DAOID, chainID types.BlockchainID, contractAddressID types.AddressID, proposalID *big.Int) (*types.Proposal, error)
	StoreProposalFinalization(ctx context.Context, finalization *types.ProposalFinalization) (*types.ProposalFinalization, error)

	StoreVoteAction(ctx context.Context, voteAction *types.VoteAction) (*types.VoteAction, error)

	StoreGovernanceParamChange(ctx context.Context, p *types.GovernanceParamChange) error
}
