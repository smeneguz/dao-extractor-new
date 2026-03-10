package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/big"
	"time"

	dbtypes "github.com/dao-portal/extractor/database/postgresql/types"
	"github.com/dao-portal/extractor/types"
)

type ProposalRow struct {
	DBID types.ProposalDBID `db:"id"`

	DaoID      types.DAOID        `db:"dao_id"`
	ProposalID dbtypes.BigNumeric `db:"proposal_id"`
	ChainID    types.BlockchainID `db:"chain_id"`

	CreatorAddressID  types.AddressID `db:"creator_address_id"`
	ContractAddressID types.AddressID `db:"contract_address_id"`

	Title       string               `db:"title"`
	Description *string              `db:"description"`
	Type        *string              `db:"type"`
	Status      types.ProposalStatus `db:"status"`

	CreationHeight dbtypes.BigInt `db:"creation_height"`
	CreationTxHash string         `db:"creation_tx_hash"`
	CreationTime   time.Time      `db:"creation_ts"`

	GasUsed dbtypes.BigInt `db:"gas_used"`
	GasFees dbtypes.DBCoin `db:"gas_fees"`

	StartHeight dbtypes.BigInt `db:"start_height"`
	StartTime   *time.Time     `db:"start_time"`

	EndHeight dbtypes.BigInt `db:"end_height"`
	EndTime   *time.Time     `db:"end_time"`

	Quorum dbtypes.BigNumeric `db:"quorum"`

	ExtraMetadata json.RawMessage `db:"extra_metadata"`
}

func (r *ProposalRow) ToProposal() *types.Proposal {
	return types.NewProposal(
		r.ProposalID.Int,
		r.DaoID,
		r.ChainID,
	).SetCreatorAddressID(r.CreatorAddressID).
		SetDBID(r.DBID).
		SetContractAddressID(r.ContractAddressID).
		SetDetails(r.Title, r.Description).
		SetType(r.Type).
		SetStatus(r.Status).
		SetCreationDetails(r.CreationHeight.Int, r.CreationTxHash, r.CreationTime).
		SetGasInfo(r.GasUsed.Int, *r.GasFees.Coin).
		SetStartDetails(r.StartHeight.Int, r.StartTime).
		SetEndDetails(r.EndHeight.Int, r.EndTime).
		SetQuorum(r.Quorum.Int).
		SetExtraMetadata(r.ExtraMetadata)
}

// StoreProposal stores a proposal in the database.
// If the proposal already exists, it will update the existing record.
func (db *DB) StoreProposal(ctx context.Context, proposal *types.Proposal) (*types.Proposal, error) {
	stmt := `
	INSERT INTO proposals(
		dao_id, proposal_id, chain_id, 
		creator_address_id, contract_address_id, 
		title, description, type, 
		status, 
		creation_height, creation_tx_hash, creation_ts, 
		gas_used, gas_fees, 
		start_height, start_time, 
		end_height, end_time, 
		quorum, 
		extra_metadata)
	VALUES (
		$1, $2, $3,
		$4, $5, 
		$6, $7, $8,
		$9,
		$10, $11, $12,
		$13, $14,
		$15, $16,
		$17, $18, 
		$19,
		$20)
	ON CONFLICT (dao_id, proposal_id, chain_id) DO UPDATE SET
		creator_address_id = $4,
		contract_address_id = $5,
		title = $6,
		description = $7,
		type = $8,
		status = $9,
		creation_height = $10,
		creation_tx_hash = $11,
		creation_ts = $12,
		gas_used = $13,
		gas_fees = $14,
		start_height = $15,
		start_time = $16,
		end_height = $17,
		end_time = $18,
		quorum = $19,
		extra_metadata = $20
	RETURNING id
	`

	var proposalDBID types.ProposalDBID
	err := db.SQL.QueryRowContext(ctx, stmt,
		proposal.DaoID,
		dbtypes.NewBigInt(proposal.ProposalID),
		proposal.ChainID,
		proposal.CreatorAddressID,
		proposal.ContractAddressID,
		proposal.Title,
		proposal.Description,
		proposal.Type,
		proposal.Status,
		dbtypes.NewBigInt(proposal.CreationHeight),
		proposal.CreationTxHash,
		proposal.CreationTime,
		dbtypes.NewBigInt(proposal.GasUsed),
		dbtypes.NewDBCoin(&proposal.GasFees),
		dbtypes.NewBigInt(proposal.StartHeight),
		proposal.StartTime,
		dbtypes.NewBigInt(proposal.EndHeight),
		proposal.EndTime,
		dbtypes.NewBigInt(proposal.Quorum),
		proposal.ExtraMetadata,
	).Scan(&proposalDBID)
	if err != nil {
		return nil, err
	}

	return proposal.SetDBID(proposalDBID), nil
}

func (db *DB) getProposalByStmt(
	ctx context.Context, stmt string, args ...any,
) (*types.Proposal, error) {
	var proposalRow ProposalRow
	err := db.SQL.GetContext(ctx, &proposalRow, stmt, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	proposal := proposalRow.ToProposal()
	if proposal == nil {
		return nil, nil
	}

	return proposal, nil
}

func (db *DB) GetProposalByProposalID(
	ctx context.Context,
	daoID types.DAOID,
	chainID types.BlockchainID,
	contractAddressID types.AddressID,
	proposalID *big.Int,
) (*types.Proposal, error) {
	return db.getProposalByStmt(ctx, `
		SELECT *
		FROM proposals
		WHERE dao_id = $1 AND proposal_id = $2 AND chain_id = $3 AND contract_address_id = $4
		`,
		daoID, proposalID.Text(10), chainID, contractAddressID,
	)
}

// -------------------------------------------------------------------------------------------------------------------
// ---- ProposalFinalization
// -------------------------------------------------------------------------------------------------------------------

type ProposalFinalizationRow struct {
	DBID types.ProposalFinalizationDBID `db:"id"`

	// ProposalID is the ID of the proposal that has been finalized.
	ProposalDBID types.ProposalDBID `db:"proposal_id"`
	// TxHash is the hash of the transaction that has been used to finalize the proposal.
	TxHash string `db:"tx_hash"`
	// Height is the height at which the finalization has been executed.
	Height dbtypes.BigInt `db:"height"`
	// GasUsed is the amount of gas consumed to finalize the proposal.
	GasUsed dbtypes.BigInt `db:"gas_used"`
	// GasFees are the fees paid for the gas used to finalize the proposal.
	GasFees dbtypes.DBCoin `db:"gas_fees"`
	// Timestamp is the timestamp of the block that has included the transaction.
	Timestamp time.Time `db:"ts"`

	// The status that this event triggered.
	StatusTriggered types.ProposalStatus `db:"status_triggered"`
	// ExtraMetadata is any extra metadata that has been attached to the finalization.
	ExtraMetadata json.RawMessage `db:"extra_metadata"`
}

func (db *DB) StoreProposalFinalization(ctx context.Context, finalization *types.ProposalFinalization) (*types.ProposalFinalization, error) {
	stmt := `
	INSERT INTO proposal_finalizations(
		proposal_id, tx_hash, height, gas_used, gas_fees, ts, 
		status_triggered, extra_metadata)
	VALUES (
		$1, $2, $3, $4, $5, $6, 
		$7, $8)
	ON CONFLICT (proposal_id) DO UPDATE SET
		tx_hash = $2,
		height = $3,
		gas_used = $4,
		gas_fees = $5,
		ts = $6,
		status_triggered = $7,
		extra_metadata = $8
	RETURNING id
	`

	var proposalFinalizationDBID types.ProposalFinalizationDBID
	err := db.SQL.QueryRowContext(ctx, stmt,
		finalization.ProposalDBID,
		finalization.TxHash,
		dbtypes.NewBigInt(finalization.Height),
		dbtypes.NewBigInt(finalization.GasUsed),
		dbtypes.NewDBCoin(&finalization.GasFees),
		finalization.Timestamp,
		finalization.StatusTriggered,
		finalization.ExtraMetadata,
	).Scan(&proposalFinalizationDBID)
	if err != nil {
		return nil, err
	}

	return finalization.SetDBID(proposalFinalizationDBID), nil
}

// -------------------------------------------------------------------------------------------------------------------
// ---- VoteAction
// -------------------------------------------------------------------------------------------------------------------

type VoteActionRow struct {
	DBID types.VoteActionDBID `db:"id"`

	ProposalDBID       types.ProposalDBID `db:"proposal_id"`
	SenderAddressID    types.AddressID    `db:"sender_address_id"`
	ContractAddressID  types.AddressID    `db:"contract_address_id"`
	DelegatorAddressID *types.AddressID   `db:"delegator_address_id"`

	TxHash    string         `db:"tx_hash"`
	Height    dbtypes.BigInt `db:"height"`
	Timestamp time.Time      `db:"ts"`

	ActionType    types.VoteActionType `db:"action_type"`
	Vote          *int8                `db:"vote"`
	VotingPower   *dbtypes.BigNumeric  `db:"voting_power"`
	ExtraMetadata json.RawMessage      `db:"extra_metadata"`
}

func (r *VoteActionRow) ToVoteAction() *types.VoteAction {
	var votingPower *big.Int
	if r.VotingPower != nil {
		votingPower = r.VotingPower.Int
	}
	return &types.VoteAction{
		DBID:               r.DBID,
		ProposalDBID:       r.ProposalDBID,
		SenderAddressID:    r.SenderAddressID,
		ContractAddressID:  r.ContractAddressID,
		DelegatorAddressID: r.DelegatorAddressID,
		TxHash:             r.TxHash,
		Height:             r.Height.Int,
		Timestamp:          r.Timestamp,
		ActionType:         r.ActionType,
		Vote:               r.Vote,
		VotingPower:        votingPower,
		ExtraMetadata:      r.ExtraMetadata,
	}
}

func (db *DB) StoreVoteAction(ctx context.Context, voteAction *types.VoteAction) (*types.VoteAction, error) {
	stmt := `
	INSERT INTO vote_actions(
		proposal_id, sender_address_id, contract_address_id, delegator_address_id, 
		tx_hash, height, ts, 
		action_type, vote, voting_power, extra_metadata)
	VALUES (
		$1, $2, $3, $4, 
		$5, $6, $7, 
		$8, $9, $10, $11)
	ON CONFLICT (tx_hash, proposal_id, sender_address_id) DO UPDATE SET
		ts = $7,
		action_type = $8,
		vote = $9,
		voting_power = $10,
		extra_metadata = $11
	RETURNING id
	`
	var voteActionDBID types.VoteActionDBID
	err := db.SQL.QueryRowContext(ctx, stmt,
		voteAction.ProposalDBID,
		voteAction.SenderAddressID,
		voteAction.ContractAddressID,
		voteAction.DelegatorAddressID,
		voteAction.TxHash,
		dbtypes.NewBigInt(voteAction.Height),
		voteAction.Timestamp.UTC(),
		voteAction.ActionType,
		voteAction.Vote,
		dbtypes.NewBigNumeric(voteAction.VotingPower),
		voteAction.ExtraMetadata,
	).Scan(&voteActionDBID)
	if err != nil {
		return nil, err
	}

	return voteAction.SetDBID(voteActionDBID), nil
}
