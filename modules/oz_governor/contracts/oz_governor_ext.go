package contracts

import (
	"context"
	"math/big"

	"github.com/dao-portal/extractor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (gov *OZGovernor) GetEventName(log *ethtypes.Log) string {
	if len(log.Topics) == 0 {
		return ""
	}

	switch log.Topics[0] {
	case gov.abi.Events[OZGovernorProposalCreatedEventName].ID:
		return OZGovernorProposalCreatedEventName
	case gov.abi.Events[OZGovernorProposalCanceledEventName].ID:
		return OZGovernorProposalCanceledEventName
	case gov.abi.Events[OZGovernorProposalExecutedEventName].ID:
		return OZGovernorProposalExecutedEventName
	case gov.abi.Events[OZGovernorProposalQueuedEventName].ID:
		return OZGovernorProposalQueuedEventName
	case gov.abi.Events[OZGovernorVoteCastEventName].ID:
		return OZGovernorVoteCastEventName
	case gov.abi.Events[OZGovernorVoteCastWithParamsEventName].ID:
		return OZGovernorVoteCastWithParamsEventName
	default:
		return ""
	}
}

func (gov *OZGovernor) Quorum(
	ctx context.Context,
	instance *bind.BoundContract,
	timepoint *big.Int,
) (*big.Int, error) {
	result, err := instance.CallRaw(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: timepoint,
	}, gov.PackQuorum(timepoint))
	if err != nil {
		return nil, err
	}
	return gov.UnpackQuorum(result)
}

func (gov *OZGovernor) GetProposalStatus(
	ctx context.Context,
	instance *bind.BoundContract,
	proposalID *big.Int,
	height *big.Int,
) (types.ProposalStatus, error) {
	result, err := instance.CallRaw(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: height,
	}, gov.PackState(proposalID))
	if err != nil {
		return types.ProposalStatusUnknown, err
	}
	state, err := gov.UnpackState(result)
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
