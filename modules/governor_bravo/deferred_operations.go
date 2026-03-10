package governorbravo

import (
	"context"
	"fmt"

	fluxevmtypes "github.com/dao-portal/extractor/flux/evm/types"
	"github.com/dao-portal/extractor/types"
)

func (m *Module) handleDeferredOperations(
	ctx context.Context,
	block *fluxevmtypes.Block,
	inst *DAOInstance,
) error {
	height := block.GetHeight().ToBigInt()

	ops := inst.DeferredOps.PopOperationsBeforeHeight(height)
	if len(ops) == 0 {
		return nil
	}

	for _, op := range ops {
		switch op.Type {
		case OperationFetchProposalStatus:
			if err := m.handleFetchProposalStatus(ctx, &op, inst); err != nil {
				m.logger.Error().Err(err).
					Str("dao", inst.DAO.Name).
					Msg("handle fetch proposal status")
				return err
			}
		}
	}

	return nil
}

func (m *Module) handleFetchProposalStatus(
	ctx context.Context,
	op *types.HeightDeferredOperation,
	inst *DAOInstance,
) error {
	m.logger.Debug().
		Str("dao", inst.DAO.Name).
		Str("op_type", op.Type).
		Str("op_height", op.Height.Text(10)).
		Msg("handling fetch proposal status")

	var payload FetchGovernorBravoProposalStatus
	if err := op.DecodePayload(&payload); err != nil {
		return err
	}

	proposal, err := m.db.GetProposalByProposalID(
		ctx, inst.DAO.ID, inst.Chain.ID, inst.GovernorBravo.Address.ID, payload.ProposalID,
	)
	if err != nil {
		return err
	}
	if proposal == nil {
		m.logger.Warn().
			Str("dao", inst.DAO.Name).
			Str("proposal_id", payload.ProposalID.String()).
			Msg("proposal not found for deferred status fetch, skipping")
		return nil
	}

	proposalStatus, err := inst.GovernorBravo.Contract.GetProposalStatus(
		ctx, inst.GovernorBravo.Instance, payload.ProposalID, op.Height,
	)
	if err != nil {
		return fmt.Errorf("get proposal status from contract: %w", err)
	}

	if _, err = m.db.StoreProposal(ctx, proposal.SetStatus(proposalStatus)); err != nil {
		return err
	}

	if err = m.db.RemoveHeightDeferredOperations(ctx, op); err != nil {
		return err
	}

	m.logger.Info().
		Str("dao", inst.DAO.Name).
		Str("proposal_id", proposal.ProposalID.String()).
		Str("new_status", string(proposal.Status)).
		Msg("updated proposal status")

	return nil
}
