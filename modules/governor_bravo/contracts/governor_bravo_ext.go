package contracts

import (
	"context"
	"math/big"

	"github.com/dao-portal/extractor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (governorBravo *GovernorBravo) GetEventName(log *ethtypes.Log) string {
	if len(log.Topics) == 0 {
		return ""
	}

	switch log.Topics[0] {
	case governorBravo.abi.Events[GovernorBravoProposalCreatedEventName].ID:
		return GovernorBravoProposalCreatedEventName
	case governorBravo.abi.Events[GovernorBravoProposalCanceledEventName].ID:
		return GovernorBravoProposalCanceledEventName
	case governorBravo.abi.Events[GovernorBravoProposalExecutedEventName].ID:
		return GovernorBravoProposalExecutedEventName
	case governorBravo.abi.Events[GovernorBravoProposalQueuedEventName].ID:
		return GovernorBravoProposalQueuedEventName
	case governorBravo.abi.Events[GovernorBravoProposalThresholdSetEventName].ID:
		return GovernorBravoProposalThresholdSetEventName
	case governorBravo.abi.Events[GovernorBravoVoteCastEventName].ID:
		return GovernorBravoVoteCastEventName
	case governorBravo.abi.Events[GovernorBravoVotingDelaySetEventName].ID:
		return GovernorBravoVotingDelaySetEventName
	case governorBravo.abi.Events[GovernorBravoVotingPeriodSetEventName].ID:
		return GovernorBravoVotingPeriodSetEventName
	default:
		return ""
	}
}

func (governorBravo *GovernorBravo) Implementation(
	ctx context.Context,
	instance *bind.BoundContract,
	height *big.Int,
) (common.Address, error) {
	output, err := instance.CallRaw(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: height,
	}, governorBravo.PackImplementation())
	if err != nil {
		return common.Address{}, err
	}

	return governorBravo.UnpackImplementation(output)
}

func (governorBravo *GovernorBravo) QuorumVotes(
	ctx context.Context,
	instance *bind.BoundContract,
	height *big.Int,
) (*big.Int, error) {
	result, err := instance.CallRaw(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: height,
	}, governorBravo.PackQuorumVotes())
	if err != nil {
		return nil, err
	}
	return governorBravo.UnpackQuorumVotes(result)
}

func (governorBravo *GovernorBravo) GetProposalStatus(
	ctx context.Context,
	instance *bind.BoundContract,
	proposalID *big.Int,
	height *big.Int,
) (types.ProposalStatus, error) {
	result, err := instance.CallRaw(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: height,
	}, governorBravo.PackState(proposalID))
	if err != nil {
		return types.ProposalStatusUnknown, err
	}
	state, err := governorBravo.UnpackState(result)
	if err != nil {
		return types.ProposalStatusUnknown, err
	}

	switch state {
	case 0: // Pending
		return types.ProposalStatusPending, nil
	case 1: // Active
		return types.ProposalStatusActive, nil
	case 2: // Canceled
		return types.ProposalStatusCanceled, nil
	case 3: // Defeated
		return types.ProposalStatusDefeated, nil
	case 4: // Succeeded
		return types.ProposalStatusVoteClosed, nil
	case 5: // Queued
		return types.ProposalStatusVoteClosed, nil
	case 6: // Expired
		return types.ProposalStatusExpired, nil
	case 7: // Executed
		return types.ProposalStatusExecuted, nil
	default:
		return types.ProposalStatusUnknown, nil
	}
}
