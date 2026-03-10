package types

import (
	"math/big"
	"time"
)

// GovernanceParamChange represents a governance parameter change event
// (VotingDelaySet, VotingPeriodSet, ProposalThresholdSet).
type GovernanceParamChange struct {
	DBID            uint64
	DaoID           DAOID
	ChainID         string
	ContractAddress string
	ParamName       string   // "voting_delay", "voting_period", "proposal_threshold"
	OldValue        *big.Int
	NewValue        *big.Int
	TxHash          string
	BlockHeight     uint64
	LogIndex        uint
	Timestamp       time.Time
}
