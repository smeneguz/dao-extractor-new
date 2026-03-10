package governorbravo

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	fluxevmtypes "github.com/dao-portal/extractor/flux/evm/types"
	"github.com/dao-portal/extractor/modules/governor_bravo/contracts"
	"github.com/dao-portal/extractor/types"
)

func (m *Module) handleGovernorBravoEvents(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	eventName := inst.GovernorBravo.Contract.GetEventName(log.EthLog)
	switch eventName {
	case contracts.GovernorBravoProposalCreatedEventName:
		return m.handleProposalCreatedEvent(ctx, block, tx, log, inst)
	case contracts.GovernorBravoProposalCanceledEventName:
		return m.handleProposalCanceledEvent(ctx, block, tx, log, inst)
	case contracts.GovernorBravoProposalExecutedEventName:
		return m.handleProposalExecutedEvent(ctx, block, tx, log, inst)
	case contracts.GovernorBravoVoteCastEventName:
		return m.handleVoteCastEvent(ctx, block, tx, log, inst)
	case contracts.GovernorBravoProposalQueuedEventName:
		return m.handleProposalQueuedEvent(ctx, block, tx, log, inst)
	case contracts.GovernorBravoVotingDelaySetEventName:
		return m.handleVotingDelaySetEvent(ctx, block, tx, log, inst)
	case contracts.GovernorBravoVotingPeriodSetEventName:
		return m.handleVotingPeriodSetEvent(ctx, block, tx, log, inst)
	case contracts.GovernorBravoProposalThresholdSetEventName:
		return m.handleProposalThresholdSetEvent(ctx, block, tx, log, inst)
	default:
		return nil
	}
}

func (m *Module) handleProposalCreatedEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	proposalCreated, err := inst.GovernorBravo.Contract.UnpackProposalCreatedEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal created event: %w", err)
	}

	quorum, err := inst.GovernorBravo.Contract.QuorumVotes(
		ctx, inst.GovernorBravo.Instance, block.GetHeight().ToBigInt(),
	)
	if err != nil {
		return fmt.Errorf("get quorum at height %d: %w", block.GetHeight(), err)
	}

	gasFee, gasUsed, err := m.evmNode.GetTxGasFeesAndUsed(ctx, tx)
	if err != nil {
		return fmt.Errorf("get tx gas fees: %w", err)
	}
	txFees := types.NewCoin("gwei", gasFee)

	proposerAddr := types.NewAddress(proposalCreated.Proposer.Hex(), "", false, types.AddressEncodingTypeHex)
	proposerAddr, err = m.db.InsertAddress(ctx, proposerAddr, true)
	if err != nil {
		return fmt.Errorf("insert proposer address %s: %w", proposerAddr.Address, err)
	}

	// Store proposal actions (targets, values, signatures, calldatas) for future KPI analysis.
	proposalMeta, _ := json.Marshal(map[string]interface{}{
		"targets":    proposalCreated.Targets,
		"values":     proposalCreated.Values,
		"signatures": proposalCreated.Signatures,
		"calldatas":  proposalCreated.Calldatas,
	})

	proposal := types.NewProposal(proposalCreated.Id, inst.DAO.ID, inst.Chain.ID)
	proposal.SetDetails(fmt.Sprintf("Proposal #%s", proposalCreated.Id.String()), &proposalCreated.Description).
		SetCreationDetails(block.GetHeight().ToBigInt(), tx.Hash.NormalizedHex(), block.GetTimeStamp()).
		SetContractAddressID(inst.GovernorBravo.Address.ID).
		SetCreatorAddressID(proposerAddr.ID).
		SetStatus(types.ProposalStatusActive).
		SetStartDetails(proposalCreated.StartBlock, nil).
		SetEndDetails(proposalCreated.EndBlock, nil).
		SetQuorum(quorum).
		SetGasInfo(gasUsed, txFees).
		SetExtraMetadata(proposalMeta)

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("tx_hash", tx.Hash.NormalizedHex()).
		Str("proposal_id", proposal.ProposalID.Text(10)).
		Str("proposer", proposerAddr.Address).
		Msg("proposal created")

	if _, err = m.db.StoreProposal(ctx, proposal); err != nil {
		return fmt.Errorf("store proposal: %w", err)
	}

	// Schedule deferred status update after voting period ends.
	marshaledPayload, err := json.Marshal(&FetchGovernorBravoProposalStatus{
		ProposalID: proposal.ProposalID,
	})
	if err != nil {
		return fmt.Errorf("marshal deferred payload: %w", err)
	}

	creatorKey := operationCreatorKeyForDAO(inst.Config.Symbol)
	op := types.NewHeightDeferredOperation(
		creatorKey,
		OperationFetchProposalStatus,
		new(big.Int).Add(proposal.EndHeight, big.NewInt(1)),
		json.RawMessage(marshaledPayload),
	)

	inst.DeferredOps.AppendOperation(op)

	if err = m.db.StoreHeightDeferredOperation(ctx, op); err != nil {
		return fmt.Errorf("store deferred operation: %w", err)
	}

	return nil
}

func (m *Module) handleProposalCanceledEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	proposalCanceled, err := inst.GovernorBravo.Contract.UnpackProposalCanceledEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal canceled event: %w", err)
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.GovernorBravo.Address.ID, proposalCanceled.Id,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", proposalCanceled.Id.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for canceled event, skipping")
		return nil
	}

	txGasFees, txGasUsed, err := m.evmNode.GetTxGasFeesAndUsed(ctx, tx)
	if err != nil {
		return fmt.Errorf("get tx gas fees: %w", err)
	}
	txFees := types.NewCoin("gwei", txGasFees)

	if _, err = m.db.StoreProposal(ctx, proposal.SetStatus(types.ProposalStatusCanceled)); err != nil {
		return fmt.Errorf("set proposal %d to canceled: %w", proposal.DBID, err)
	}

	fin := types.NewProposalFinalization(proposal.DBID).
		SetExecutionDetails(block.GetHeight().ToBigInt(), tx.GetHash(), block.GetTimeStamp()).
		SetGasInfo(txGasUsed, txFees).
		SetStatusTriggered(types.ProposalStatusCanceled)

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("tx_hash", tx.Hash.NormalizedHex()).
		Str("proposal_id", proposal.ProposalID.String()).
		Msg("proposal canceled")

	if _, err = m.db.StoreProposalFinalization(ctx, fin); err != nil {
		return fmt.Errorf("store finalization: %w", err)
	}

	return nil
}

func (m *Module) handleProposalExecutedEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	proposalExecuted, err := inst.GovernorBravo.Contract.UnpackProposalExecutedEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal executed event: %w", err)
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.GovernorBravo.Address.ID, proposalExecuted.Id,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", proposalExecuted.Id.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for executed event, skipping")
		return nil
	}

	txGasFees, txGasUsed, err := m.evmNode.GetTxGasFeesAndUsed(ctx, tx)
	if err != nil {
		return fmt.Errorf("get tx gas fees: %w", err)
	}
	txFees := types.NewCoin("gwei", txGasFees)

	if _, err = m.db.StoreProposal(ctx, proposal.SetStatus(types.ProposalStatusExecuted)); err != nil {
		return fmt.Errorf("set proposal %d to executed: %w", proposal.DBID, err)
	}

	fin := types.NewProposalFinalization(proposal.DBID).
		SetExecutionDetails(block.GetHeight().ToBigInt(), tx.GetHash(), block.GetTimeStamp()).
		SetGasInfo(txGasUsed, txFees).
		SetStatusTriggered(types.ProposalStatusExecuted)

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("tx_hash", tx.Hash.NormalizedHex()).
		Str("proposal_id", proposal.ProposalID.String()).
		Msg("proposal executed")

	if _, err = m.db.StoreProposalFinalization(ctx, fin); err != nil {
		return fmt.Errorf("store finalization: %w", err)
	}

	return nil
}

func (m *Module) handleVoteCastEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	voteCast, err := inst.GovernorBravo.Contract.UnpackVoteCastEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack vote cast event: %w", err)
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.GovernorBravo.Address.ID, voteCast.ProposalId,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", voteCast.ProposalId.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for vote event, skipping")
		return nil
	}

	voterAddr, err := m.db.InsertAddress(
		ctx,
		types.NewAddress(voteCast.Voter.Hex(), "", false, types.AddressEncodingTypeHex),
		true,
	)
	if err != nil {
		return fmt.Errorf("insert voter address %s: %w", voteCast.Voter.Hex(), err)
	}

	voteAction := types.NewVoteAction(proposal.DBID, voterAddr.ID, inst.GovernorBravo.Address.ID).
		SetExecutionDetails(block.GetHeight().ToBigInt(), tx.GetHash(), block.GetTimeStamp()).
		SetVotingPower(voteCast.Votes)

	if voteCast.Reason != "" {
		voteMeta, _ := json.Marshal(map[string]string{"reason": voteCast.Reason})
		voteAction.SetExtraMetadata(voteMeta)
	}

	switch voteCast.Support {
	case 0: // Against
		voteAction.SetVote(0)
	case 1: // For
		voteAction.SetVote(1)
	case 2: // Abstain
		voteAction.SetAbstain()
	default:
		return fmt.Errorf("invalid vote support value: %d", voteCast.Support)
	}

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("tx_hash", tx.Hash.NormalizedHex()).
		Str("proposal_id", proposal.ProposalID.String()).
		Str("voter", voterAddr.Address).
		Msg("vote cast")

	if _, err = m.db.StoreVoteAction(ctx, voteAction); err != nil {
		return fmt.Errorf("store vote action: %w", err)
	}

	return nil
}

func (m *Module) handleProposalQueuedEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	proposalQueued, err := inst.GovernorBravo.Contract.UnpackProposalQueuedEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal queued event: %w", err)
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.GovernorBravo.Address.ID, proposalQueued.Id,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", proposalQueued.Id.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for queued event, skipping")
		return nil
	}

	if _, err = m.db.StoreProposal(ctx, proposal.SetStatus(types.ProposalStatusVoteClosed)); err != nil {
		return fmt.Errorf("set proposal %d to vote_closed: %w", proposal.DBID, err)
	}

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("tx_hash", tx.Hash.NormalizedHex()).
		Str("proposal_id", proposal.ProposalID.String()).
		Str("eta", proposalQueued.Eta.String()).
		Msg("proposal queued")

	return nil
}

func (m *Module) handleVotingDelaySetEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	event, err := inst.GovernorBravo.Contract.UnpackVotingDelaySetEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack VotingDelaySet: %w", err)
	}

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("old", event.OldVotingDelay.String()).
		Str("new", event.NewVotingDelay.String()).
		Msg("voting delay changed")

	return m.db.StoreGovernanceParamChange(ctx, &types.GovernanceParamChange{
		DaoID:           inst.DAO.ID,
		ChainID:         string(inst.Chain.ChainID),
		ContractAddress: inst.GovernorBravo.Address.Address,
		ParamName:       "voting_delay",
		OldValue:        event.OldVotingDelay,
		NewValue:        event.NewVotingDelay,
		TxHash:          tx.Hash.NormalizedHex(),
		BlockHeight:     uint64(block.GetHeight()),
		LogIndex:        log.Index,
		Timestamp:       block.GetTimeStamp(),
	})
}

func (m *Module) handleVotingPeriodSetEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	event, err := inst.GovernorBravo.Contract.UnpackVotingPeriodSetEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack VotingPeriodSet: %w", err)
	}

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("old", event.OldVotingPeriod.String()).
		Str("new", event.NewVotingPeriod.String()).
		Msg("voting period changed")

	return m.db.StoreGovernanceParamChange(ctx, &types.GovernanceParamChange{
		DaoID:           inst.DAO.ID,
		ChainID:         string(inst.Chain.ChainID),
		ContractAddress: inst.GovernorBravo.Address.Address,
		ParamName:       "voting_period",
		OldValue:        event.OldVotingPeriod,
		NewValue:        event.NewVotingPeriod,
		TxHash:          tx.Hash.NormalizedHex(),
		BlockHeight:     uint64(block.GetHeight()),
		LogIndex:        log.Index,
		Timestamp:       block.GetTimeStamp(),
	})
}

func (m *Module) handleProposalThresholdSetEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	event, err := inst.GovernorBravo.Contract.UnpackProposalThresholdSetEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack ProposalThresholdSet: %w", err)
	}

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("old", event.OldProposalThreshold.String()).
		Str("new", event.NewProposalThreshold.String()).
		Msg("proposal threshold changed")

	return m.db.StoreGovernanceParamChange(ctx, &types.GovernanceParamChange{
		DaoID:           inst.DAO.ID,
		ChainID:         string(inst.Chain.ChainID),
		ContractAddress: inst.GovernorBravo.Address.Address,
		ParamName:       "proposal_threshold",
		OldValue:        event.OldProposalThreshold,
		NewValue:        event.NewProposalThreshold,
		TxHash:          tx.Hash.NormalizedHex(),
		BlockHeight:     uint64(block.GetHeight()),
		LogIndex:        log.Index,
		Timestamp:       block.GetTimeStamp(),
	})
}
