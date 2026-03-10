package ozgovernor

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	fluxevmtypes "github.com/dao-portal/extractor/flux/evm/types"
	"github.com/dao-portal/extractor/modules/oz_governor/contracts"
	"github.com/dao-portal/extractor/types"
)

func (m *Module) handleOZGovernorEvents(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	eventName := inst.Governor.Contract.GetEventName(log.EthLog)
	switch eventName {
	case contracts.OZGovernorProposalCreatedEventName:
		return m.handleProposalCreatedEvent(ctx, block, tx, log, inst)
	case contracts.OZGovernorProposalCanceledEventName:
		return m.handleProposalCanceledEvent(ctx, block, tx, log, inst)
	case contracts.OZGovernorProposalExecutedEventName:
		return m.handleProposalExecutedEvent(ctx, block, tx, log, inst)
	case contracts.OZGovernorVoteCastEventName:
		return m.handleVoteCastEvent(ctx, block, tx, log, inst)
	case contracts.OZGovernorVoteCastWithParamsEventName:
		return m.handleVoteCastWithParamsEvent(ctx, block, tx, log, inst)
	case contracts.OZGovernorProposalQueuedEventName:
		return m.handleProposalQueuedEvent(ctx, block, tx, log, inst)
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
	proposalCreated, err := inst.Governor.Contract.UnpackProposalCreatedEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal created event: %w", err)
	}

	quorum, err := inst.Governor.Contract.Quorum(
		ctx, inst.Governor.Instance, block.GetHeight().ToBigInt(),
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

	proposal := types.NewProposal(proposalCreated.ProposalId, inst.DAO.ID, inst.Chain.ID)
	proposal.SetDetails(fmt.Sprintf("Proposal #%s", proposalCreated.ProposalId.String()), &proposalCreated.Description).
		SetCreationDetails(block.GetHeight().ToBigInt(), tx.Hash.NormalizedHex(), block.GetTimeStamp()).
		SetContractAddressID(inst.Governor.Address.ID).
		SetCreatorAddressID(proposerAddr.ID).
		SetStatus(types.ProposalStatusActive).
		SetStartDetails(proposalCreated.VoteStart, nil).
		SetEndDetails(proposalCreated.VoteEnd, nil).
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
	marshaledPayload, err := json.Marshal(&FetchOZGovernorProposalStatus{
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
	proposalCanceled, err := inst.Governor.Contract.UnpackProposalCanceledEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal canceled event: %w", err)
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.Governor.Address.ID, proposalCanceled.ProposalId,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", proposalCanceled.ProposalId.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for canceled event, skipping (will be processed when proposal is indexed)")
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
	proposalExecuted, err := inst.Governor.Contract.UnpackProposalExecutedEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal executed event: %w", err)
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.Governor.Address.ID, proposalExecuted.ProposalId,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", proposalExecuted.ProposalId.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for executed event, skipping (will be processed when proposal is indexed)")
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
	voteCast, err := inst.Governor.Contract.UnpackVoteCastEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack vote cast event: %w", err)
	}

	return m.processVote(ctx, block, tx, inst, voteCast.Voter, voteCast.ProposalId, voteCast.Support, voteCast.Weight, voteCast.Reason, nil)
}

func (m *Module) handleVoteCastWithParamsEvent(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	log *fluxevmtypes.LogEntry,
	inst *DAOInstance,
) error {
	voteCast, err := inst.Governor.Contract.UnpackVoteCastWithParamsEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack vote cast with params event: %w", err)
	}

	return m.processVote(ctx, block, tx, inst, voteCast.Voter, voteCast.ProposalId, voteCast.Support, voteCast.Weight, voteCast.Reason, voteCast.Params)
}

func (m *Module) processVote(
	ctx context.Context,
	block *fluxevmtypes.Block,
	tx *fluxevmtypes.Tx,
	inst *DAOInstance,
	voter interface{ Hex() string },
	proposalID *big.Int,
	support uint8,
	weight *big.Int,
	reason string,
	params []byte,
) error {
	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.Governor.Address.ID, proposalID,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", proposalID.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for vote event, skipping (will be processed when proposal is indexed)")
		return nil
	}

	voterAddr, err := m.db.InsertAddress(
		ctx,
		types.NewAddress(voter.Hex(), "", false, types.AddressEncodingTypeHex),
		true,
	)
	if err != nil {
		return fmt.Errorf("insert voter address %s: %w", voter.Hex(), err)
	}

	voteAction := types.NewVoteAction(proposal.DBID, voterAddr.ID, inst.Governor.Address.ID).
		SetExecutionDetails(block.GetHeight().ToBigInt(), tx.GetHash(), block.GetTimeStamp()).
		SetVotingPower(weight)

	if reason != "" || len(params) > 0 {
		meta := map[string]interface{}{}
		if reason != "" {
			meta["reason"] = reason
		}
		if len(params) > 0 {
			meta["params"] = params
		}
		voteMeta, _ := json.Marshal(meta)
		voteAction.SetExtraMetadata(voteMeta)
	}

	switch support {
	case 0: // Against
		voteAction.SetVote(0)
	case 1: // For
		voteAction.SetVote(1)
	case 2: // Abstain
		voteAction.SetAbstain()
	default:
		return fmt.Errorf("invalid vote support value: %d", support)
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
	proposalQueued, err := inst.Governor.Contract.UnpackProposalQueuedEvent(log.EthLog)
	if err != nil {
		return fmt.Errorf("unpack proposal queued event: %w", err)
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.Governor.Address.ID, proposalQueued.ProposalId,
	)
	if err != nil {
		return fmt.Errorf("get proposal by id: %w", err)
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", proposalQueued.ProposalId.String()).
			Uint64("height", uint64(block.GetHeight())).
			Msg("proposal not found for queued event, skipping (will be processed when proposal is indexed)")
		return nil
	}

	if _, err = m.db.StoreProposal(ctx, proposal.SetStatus(types.ProposalStatusVoteClosed)); err != nil {
		return fmt.Errorf("set proposal %d to vote_closed: %w", proposal.DBID, err)
	}

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("tx_hash", tx.Hash.NormalizedHex()).
		Str("proposal_id", proposal.ProposalID.String()).
		Str("eta", proposalQueued.EtaSeconds.String()).
		Msg("proposal queued")

	return nil
}
