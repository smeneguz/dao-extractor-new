package types

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/dao-portal/extractor/utils"
)

type ProposalStatus string

const (
	ProposalStatusUnknown ProposalStatus = ""
	// ProposalStatusPending means the proposal has been created but is not yet open for voting.
	ProposalStatusPending ProposalStatus = "PENDING"
	// ProposalStatusActive means the proposal is currently open for voting.
	// The voting period has started and has not yet ended.
	ProposalStatusActive ProposalStatus = "ACTIVE"
	// ProposalStatusVoteClosed means the proposal is still active but the voting period has ended.
	ProposalStatusVoteClosed ProposalStatus = "VOTE_CLOSED"
	// ProposalStatusCanceled means the proposal has been withdrawn before the end of its voting period.
	ProposalStatusCanceled ProposalStatus = "CANCELED"
	// ProposalStatusExecuted means the proposal was successful.
	// It received the required number of votes to pass, and the on-chain action it proposed has been successfully carried out.
	ProposalStatusExecuted ProposalStatus = "EXECUTED"
	// ProposalStatusDefeated means the proposal failed to get enough positive votes to pass.
	ProposalStatusDefeated ProposalStatus = "DEFEATED"
	// ProposalStatusExpired means the proposal's voting period has ended, but it was neither executed nor defeated.
	// This typically happens when a proposal fails to meet the minimum participation requirement (quorum).
	ProposalStatusExpired ProposalStatus = "EXPIRED"
)

type ProposalDBID uint64

type Proposal struct {
	DBID ProposalDBID

	// Proposal ID is the unique identifier for the proposal on-chain.
	ProposalID *big.Int
	// DaoID is the ID of the DAO this proposal belongs to.
	DaoID DAOID
	// ChainID is the ID of the blockchain where the proposal originated.
	ChainID BlockchainID

	// Creation information
	// CreatorAddressID is the address ID of the proposal's creator.
	CreatorAddressID AddressID
	// ContractAddressID is the address ID of the contract that has been used to create the proposal.
	ContractAddressID AddressID

	// Proposal details
	Title       string
	Description *string
	// Type of proposal, this will be populated later on using an LLM to infer the type
	// from the description
	Type *string

	Status ProposalStatus

	// Proposal creation details
	CreationHeight *big.Int
	CreationTxHash string
	CreationTime   time.Time

	// GasUsed is the amount of gas consumed to create the proposal.
	GasUsed *big.Int
	// GasFees are the fees paid for the gas used to create the proposal.
	GasFees Coin

	// Proposal start details, this may be different from the creation height
	// if the proposal is a delayed start.
	StartHeight *big.Int
	StartTime   *time.Time

	// Proposal end details
	EndHeight *big.Int
	EndTime   *time.Time

	// Quorum is the voting power required to consider a proposal valid.
	Quorum *big.Int

	ExtraMetadata json.RawMessage
}

func NewProposal(proposalID *big.Int, daoID DAOID, chainID BlockchainID) *Proposal {
	return &Proposal{
		DBID:          0,
		DaoID:         daoID,
		ChainID:       chainID,
		ProposalID:    proposalID,
		ExtraMetadata: json.RawMessage("{}"),
	}
}

// SetDBID sets the ID that identifies the proposal in the database.
func (p *Proposal) SetDBID(id ProposalDBID) *Proposal {
	p.DBID = id
	return p
}

// SetProposalID sets the DB ID of the user address that created the proposal.
func (p *Proposal) SetCreatorAddressID(addressID AddressID) *Proposal {
	p.CreatorAddressID = addressID
	return p
}

// SetContractAddressID sets the DB ID of the contract that has been used to create the proposal.
func (p *Proposal) SetContractAddressID(addressID AddressID) *Proposal {
	p.ContractAddressID = addressID
	return p
}

// SetDetails sets the title and description of the proposal.
func (p *Proposal) SetDetails(title string, description *string) *Proposal {
	p.Title = title
	p.Description = description
	return p
}

// SetType sets the type of the proposal.
func (p *Proposal) SetType(type_ *string) *Proposal {
	p.Type = type_
	return p
}

// SetStatus sets the status of the proposal.
func (p *Proposal) SetStatus(status ProposalStatus) *Proposal {
	p.Status = status
	return p
}

// SetCreationDetails sets the creation details of the proposal.
func (p *Proposal) SetCreationDetails(height *big.Int, txHash string, time time.Time) *Proposal {
	p.CreationHeight = height
	p.CreationTxHash = txHash
	p.CreationTime = time.UTC()
	return p
}

// SetGasInfo sets the gas used and the fees paid for the gas used to create the proposal.
func (p *Proposal) SetGasInfo(gasUsed *big.Int, fees Coin) *Proposal {
	p.GasUsed = gasUsed
	p.GasFees = fees
	return p
}

// SetStartDetails sets the start details of the proposal.
func (p *Proposal) SetStartDetails(height *big.Int, time *time.Time) *Proposal {
	p.StartHeight = height
	p.StartTime = utils.ToUTC(time)
	return p
}

// SetEndDetails sets the end details of the proposal.
func (p *Proposal) SetEndDetails(height *big.Int, time *time.Time) *Proposal {
	p.EndHeight = height
	p.EndTime = utils.ToUTC(time)
	return p
}

// SetQuorum sets the quorum of the proposal.
func (p *Proposal) SetQuorum(quorum *big.Int) *Proposal {
	p.Quorum = quorum
	return p
}

// SetExtraMetadata sets the extra metadata of the proposal.
func (p *Proposal) SetExtraMetadata(metadata json.RawMessage) *Proposal {
	p.ExtraMetadata = metadata
	return p
}

// -------------------------------------------------------------------------------------------------------------------
// ---- ProposalFinalization
// -------------------------------------------------------------------------------------------------------------------

type ProposalFinalizationDBID uint64

type ProposalFinalization struct {
	DBID ProposalFinalizationDBID

	// ProposalID is the ID of the proposal that has been finalized.
	ProposalDBID ProposalDBID
	// TxHash is the hash of the transaction that has been used to finalize the proposal.
	TxHash string
	// Height is the height at which the finalization has been executed.
	Height *big.Int
	// GasUsed is the amount of gas consumed to finalize the proposal.
	GasUsed *big.Int
	// GasFees are the fees paid for the gas used to finalize the proposal.
	GasFees Coin
	// Timestamp is the timestamp of the block that has included the transaction.
	Timestamp time.Time

	// The status that this event triggered.
	StatusTriggered ProposalStatus
	// ExtraMetadata is any extra metadata that has been attached to the finalization.
	ExtraMetadata json.RawMessage
}

func NewProposalFinalization(proposalDBID ProposalDBID) *ProposalFinalization {
	return &ProposalFinalization{
		DBID:          0,
		ProposalDBID:  proposalDBID,
		ExtraMetadata: json.RawMessage("{}"),
	}
}

// SetDBID sets the ID that identifies the proposal finalization in the database.
func (f *ProposalFinalization) SetDBID(id ProposalFinalizationDBID) *ProposalFinalization {
	f.DBID = id
	return f
}

func (f *ProposalFinalization) SetExecutionDetails(height *big.Int, txHash string, time time.Time) *ProposalFinalization {
	f.Height = height
	f.TxHash = txHash
	f.Timestamp = time.UTC()
	return f
}

// SetGasInfo sets the gas used and the fees paid for the gas used to finalize the proposal.
func (f *ProposalFinalization) SetGasInfo(gasUsed *big.Int, fees Coin) *ProposalFinalization {
	f.GasUsed = gasUsed
	f.GasFees = fees
	return f
}

// SetStatus sets the status that triggered the finalization.
func (f *ProposalFinalization) SetStatusTriggered(status ProposalStatus) *ProposalFinalization {
	f.StatusTriggered = status
	return f
}

// SetExtraMetadata sets the extra metadata of the proposal finalization.
func (f *ProposalFinalization) SetExtraMetadata(metadata json.RawMessage) *ProposalFinalization {
	f.ExtraMetadata = metadata
	return f
}

// -------------------------------------------------------------------------------------------------------------------
// ---- VoteAction
// -------------------------------------------------------------------------------------------------------------------

type VoteActionDBID uint64

type VoteActionType string

const (
	// VoteActionTypeVote means the user has voted to a proposal.
	VoteActionTypeVote VoteActionType = "VOTE"
	// VoteActionTypeCancel means the user has withdrawn their vote.
	VoteActionTypeCancel VoteActionType = "CANCEL"
	// VoteActionTypeAbstain means the user has abstained from voting.
	VoteActionTypeAbstain VoteActionType = "ABSTAIN"
)

type VoteAction struct {
	DBID VoteActionDBID

	ProposalDBID       ProposalDBID
	SenderAddressID    AddressID
	ContractAddressID  AddressID
	DelegatorAddressID *AddressID

	TxHash    string
	Height    *big.Int
	Timestamp time.Time

	ActionType    VoteActionType
	Vote          *int8
	VotingPower   *big.Int
	ExtraMetadata json.RawMessage
}

// NewVoteAction creates a new vote action.
func NewVoteAction(proposalDBID ProposalDBID, senderAddressID AddressID, contractAddressID AddressID) *VoteAction {
	return &VoteAction{
		DBID:              0,
		ProposalDBID:      proposalDBID,
		SenderAddressID:   senderAddressID,
		ContractAddressID: contractAddressID,
		ExtraMetadata:     json.RawMessage("{}"),
	}
}

// SetDBID sets the ID that identifies the vote action in the database.
func (v *VoteAction) SetDBID(id VoteActionDBID) *VoteAction {
	v.DBID = id
	return v
}

// SetDelegatorAddressID sets the ID of the delegator address that has voted.
func (v *VoteAction) SetDelegatorAddressID(addressID AddressID) *VoteAction {
	v.DelegatorAddressID = &addressID
	return v
}

// SetExecutionDetails sets the execution details of the vote action.
func (v *VoteAction) SetExecutionDetails(height *big.Int, txHash string, time time.Time) *VoteAction {
	v.Height = height
	v.TxHash = txHash
	v.Timestamp = time.UTC()
	return v
}

// SetAbstain sets the vote action to abstain.
func (v *VoteAction) SetAbstain() *VoteAction {
	v.ActionType = VoteActionTypeAbstain
	v.Vote = nil
	return v
}

// SetCancel sets the vote action to cancel to signal that the user has withdrawn their vote.
func (v *VoteAction) SetCancel() *VoteAction {
	v.ActionType = VoteActionTypeCancel
	v.Vote = nil
	return v
}

// SetVote sets the vote action to vote with the vote value.
func (v *VoteAction) SetVote(vote int8) *VoteAction {
	v.ActionType = VoteActionTypeVote
	v.Vote = &vote
	return v
}

// SetVotingPower sets the voting power of the vote action.
func (v *VoteAction) SetVotingPower(votingPower *big.Int) *VoteAction {
	v.VotingPower = votingPower
	return v
}

// SetExtraMetadata sets the extra metadata of the vote action.
func (v *VoteAction) SetExtraMetadata(metadata json.RawMessage) *VoteAction {
	v.ExtraMetadata = metadata
	return v
}
