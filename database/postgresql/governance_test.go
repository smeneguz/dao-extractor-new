package postgresql_test

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/dao-portal/extractor/database/postgresql"
	"github.com/dao-portal/extractor/types"
)

func (suite *DbTestSuite) TestStoreProposal() {
	testStartTime := time.Date(2025, time.August, 9, 19, 0, 0, 0, time.UTC)
	testEndTime := time.Date(2025, time.August, 16, 19, 0, 0, 0, time.UTC)
	testDescription := "Proposal description"
	testType := "Proposal type"

	testCases := []struct {
		name      string
		setup     func(context.Context)
		proposal  *types.Proposal
		shouldErr bool
		check     func()
	}{
		{
			name: "store proposal correctly",
			setup: func(ctx context.Context) {
				// Add a test dao
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("test-dao", "Test DAO"), false)
				suite.Require().NoError(err)

				// Add a test blockchain
				_, err = suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", "EVM"), false)
				suite.Require().NoError(err)

				// Add a test user address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x12345678", "Test Address", false, "hex"), false)
				suite.Require().NoError(err)
				// Add a test contract address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x123456789", "Test Contract Address", true, "hex"), false)
				suite.Require().NoError(err)
			},
			proposal: types.NewProposal(big.NewInt(1), 1, 1).
				SetCreatorAddressID(1).
				SetContractAddressID(2).
				SetDetails("Proposal title", &testDescription).
				SetType(&testType).
				SetStatus(types.ProposalStatusActive).
				SetCreationDetails(big.NewInt(1), "tx-hash", testStartTime).
				SetGasInfo(big.NewInt(1000), types.NewCoin("ETH", big.NewInt(100))).
				SetStartDetails(big.NewInt(1), &testStartTime).
				SetEndDetails(big.NewInt(3), &testEndTime).
				SetQuorum(big.NewInt(100)).
				SetExtraMetadata(json.RawMessage(`{"key": "value"}`)),
			shouldErr: false,
			check: func() {
				var proposal postgresql.ProposalRow
				err := suite.database.SQL.Get(&proposal, `
					SELECT * FROM proposals
					WHERE id = $1
				`, 1)

				suite.Require().NoError(err)
				suite.Require().Equal(1, proposal.DaoID)
				suite.Require().Equal(big.NewInt(1), proposal.ProposalID.Int)
				suite.Require().Equal(1, proposal.ChainID)
				suite.Require().Equal(1, proposal.CreatorAddressID)
				suite.Require().Equal(2, proposal.ContractAddressID)
				suite.Require().Equal("Proposal title", proposal.Title)
				suite.Require().Equal(testDescription, *proposal.Description)
				suite.Require().Equal(testType, *proposal.Type)
				suite.Require().Equal(types.ProposalStatusActive, proposal.Status)
				suite.Require().Equal(big.NewInt(1), proposal.CreationHeight.Int)
				suite.Require().Equal("tx-hash", proposal.CreationTxHash)
				suite.Require().Equal(testStartTime, proposal.CreationTime.UTC())
				suite.Require().Equal(uint64(1000), proposal.GasUsed)
				suite.Require().Equal(types.NewCoin("ETH", big.NewInt(100)), proposal.GasFees.Coin)
				suite.Require().Equal(big.NewInt(1), proposal.StartHeight.Int)
				suite.Require().Equal(testStartTime, proposal.StartTime.UTC())
				suite.Require().Equal(big.NewInt(3), proposal.EndHeight.Int)
				suite.Require().Equal(testEndTime, proposal.EndTime.UTC())
				suite.Require().Equal(big.NewInt(100), proposal.Quorum.Int)
				suite.Require().Equal(`{"key": "value"}`, string(proposal.ExtraMetadata))
			},
		},
		{
			name: "on conflict proposal update correctly",
			setup: func(ctx context.Context) {
				// Add a test dao
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("test-dao", "Test DAO"), false)
				suite.Require().NoError(err)

				// Add a test blockchain
				_, err = suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", "EVM"), false)
				suite.Require().NoError(err)

				// Add a test user address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x12345678", "Test Address", false, "hex"), false)
				suite.Require().NoError(err)
				// Add a test contract address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x123456789", "Test Contract Address", true, "hex"), false)
				suite.Require().NoError(err)

				proposal := types.NewProposal(big.NewInt(1), 1, 1).
					SetCreatorAddressID(1).
					SetContractAddressID(2).
					SetDetails("Proposal title", &testDescription).
					SetType(&testType).
					SetStatus(types.ProposalStatusActive).
					SetCreationDetails(big.NewInt(1), "tx-hash", testStartTime).
					SetGasInfo(big.NewInt(1000), types.NewCoin("ETH", big.NewInt(100))).
					SetStartDetails(big.NewInt(1), &testStartTime).
					SetEndDetails(big.NewInt(3), &testEndTime).
					SetQuorum(big.NewInt(100)).
					SetExtraMetadata(json.RawMessage(`{"key": "value"}`))
				_, err = suite.database.StoreProposal(ctx, proposal)
				suite.Require().NoError(err)
			},
			proposal: types.NewProposal(big.NewInt(1), 1, 1).
				SetCreatorAddressID(1).
				SetContractAddressID(2).
				SetDetails("Proposal title new", &testDescription).
				SetType(&testType).
				SetStatus(types.ProposalStatusCanceled).
				SetCreationDetails(big.NewInt(2), "tx-hash-new", testStartTime).
				SetGasInfo(big.NewInt(10000), types.NewCoin("ETH2", big.NewInt(150))).
				SetStartDetails(big.NewInt(12), &testStartTime).
				SetEndDetails(big.NewInt(30), &testEndTime).
				SetQuorum(big.NewInt(200)).
				SetExtraMetadata(json.RawMessage(`{"key": "value-new"}`)),
			shouldErr: false,
			check: func() {
				var proposal postgresql.ProposalRow
				err := suite.database.SQL.Get(&proposal, `
					SELECT * FROM proposals
					WHERE id = $1
				`, 1)

				suite.Require().NoError(err)
				suite.Require().Equal(1, proposal.DaoID)
				suite.Require().Equal(big.NewInt(1), proposal.ProposalID.Int)
				suite.Require().Equal(1, proposal.ChainID)
				suite.Require().Equal(1, proposal.CreatorAddressID)
				suite.Require().Equal(2, proposal.ContractAddressID)
				suite.Require().Equal("Proposal title new", proposal.Title)
				suite.Require().Equal(testDescription, *proposal.Description)
				suite.Require().Equal(testType, *proposal.Type)
				suite.Require().Equal(types.ProposalStatusCanceled, proposal.Status)
				suite.Require().Equal(big.NewInt(2), proposal.CreationHeight.Int)
				suite.Require().Equal("tx-hash-new", proposal.CreationTxHash)
				suite.Require().Equal(testStartTime, proposal.CreationTime.UTC())
				suite.Require().Equal(uint64(10000), proposal.GasUsed)
				suite.Require().Equal(types.NewCoin("ETH2", big.NewInt(150)), proposal.GasFees.Coin)
				suite.Require().Equal(big.NewInt(12), proposal.StartHeight.Int)
				suite.Require().Equal(testStartTime, proposal.StartTime.UTC())
				suite.Require().Equal(big.NewInt(30), proposal.EndHeight.Int)
				suite.Require().Equal(testEndTime, proposal.EndTime.UTC())
				suite.Require().Equal(big.NewInt(200), proposal.Quorum.Int)
				suite.Require().Equal(`{"key": "value-new"}`, string(proposal.ExtraMetadata))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			proposal, err := suite.database.StoreProposal(suite.T().Context(), tc.proposal)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.proposal, proposal)
			}
		})
	}
}

func (suite *DbTestSuite) TestStoreProposalFinalization() {
	testStartTime := time.Date(2025, time.August, 9, 19, 0, 0, 0, time.UTC)
	testEndTime := time.Date(2025, time.August, 16, 19, 0, 0, 0, time.UTC)
	testDescription := "Proposal description"
	testType := "Proposal type"

	testCases := []struct {
		name         string
		setup        func(context.Context)
		finalization *types.ProposalFinalization
		shouldErr    bool
		check        func()
	}{
		{
			name: "store proposal finalization correctly",
			setup: func(ctx context.Context) {
				// Add a test dao
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("test-dao", "Test DAO"), false)
				suite.Require().NoError(err)

				// Add a test blockchain
				_, err = suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", "EVM"), false)
				suite.Require().NoError(err)

				// Add a test user address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x12345678", "Test Address", false, "hex"), false)
				suite.Require().NoError(err)
				// Add a test contract address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x123456789", "Test Contract Address", true, "hex"), false)
				suite.Require().NoError(err)

				proposal := types.NewProposal(big.NewInt(1), 1, 1).
					SetCreatorAddressID(1).
					SetContractAddressID(2).
					SetDetails("Proposal title", &testDescription).
					SetType(&testType).
					SetStatus(types.ProposalStatusActive).
					SetCreationDetails(big.NewInt(1), "tx-hash", testStartTime).
					SetGasInfo(big.NewInt(1000), types.NewCoin("ETH", big.NewInt(100))).
					SetStartDetails(big.NewInt(1), &testStartTime).
					SetEndDetails(big.NewInt(3), &testEndTime).
					SetQuorum(big.NewInt(100)).
					SetExtraMetadata(json.RawMessage(`{"key": "value"}`))
				_, err = suite.database.StoreProposal(ctx, proposal)
				suite.Require().NoError(err)
			},
			finalization: types.NewProposalFinalization(1).
				SetExecutionDetails(big.NewInt(1), "tx-hash", testStartTime).
				SetGasInfo(big.NewInt(1000), types.NewCoin("ETH", big.NewInt(100))).
				SetStatusTriggered(types.ProposalStatusExecuted).
				SetExtraMetadata(json.RawMessage(`{"key": "value"}`)),
			shouldErr: false,
			check: func() {
				var proposalFinalization postgresql.ProposalFinalizationRow
				err := suite.database.SQL.Get(&proposalFinalization, `
					SELECT * FROM proposal_finalizations
					WHERE id = $1
				`, 1)

				suite.Require().NoError(err)
				suite.Require().Equal(1, proposalFinalization.ProposalDBID)
				suite.Require().Equal("tx-hash", proposalFinalization.TxHash)
				suite.Require().Equal(big.NewInt(1), proposalFinalization.Height.Int)
				suite.Require().Equal(big.NewInt(1000), proposalFinalization.GasUsed.Int)
				suite.Require().Equal(types.NewCoin("ETH", big.NewInt(100)), proposalFinalization.GasFees.Coin)
				suite.Require().Equal(testStartTime, proposalFinalization.Timestamp.UTC())
				suite.Require().Equal(types.ProposalStatusExecuted, proposalFinalization.StatusTriggered)
				suite.Require().Equal(`{"key": "value"}`, string(proposalFinalization.ExtraMetadata))
			},
		},
		{
			name: "on conflict proposal finalization update correctly",
			setup: func(ctx context.Context) {
				// Add a test dao
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("test-dao", "Test DAO"), false)
				suite.Require().NoError(err)

				// Add a test blockchain
				_, err = suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", "EVM"), false)
				suite.Require().NoError(err)

				// Add a test user address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x12345678", "Test Address", false, "hex"), false)
				suite.Require().NoError(err)
				// Add a test contract address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x123456789", "Test Contract Address", true, "hex"), false)
				suite.Require().NoError(err)

				proposal := types.NewProposal(big.NewInt(1), 1, 1).
					SetCreatorAddressID(1).
					SetContractAddressID(2).
					SetDetails("Proposal title", &testDescription).
					SetType(&testType).
					SetStatus(types.ProposalStatusActive).
					SetCreationDetails(big.NewInt(1), "tx-hash", testStartTime).
					SetGasInfo(big.NewInt(1000), types.NewCoin("ETH", big.NewInt(100))).
					SetStartDetails(big.NewInt(1), &testStartTime).
					SetEndDetails(big.NewInt(3), &testEndTime).
					SetQuorum(big.NewInt(100)).
					SetExtraMetadata(json.RawMessage(`{"key": "value"}`))
				_, err = suite.database.StoreProposal(ctx, proposal)
				suite.Require().NoError(err)

				finalization := types.NewProposalFinalization(1).
					SetExecutionDetails(big.NewInt(1), "tx-hash", testStartTime).
					SetGasInfo(big.NewInt(1000), types.NewCoin("ETH", big.NewInt(100))).
					SetStatusTriggered(types.ProposalStatusExecuted).
					SetExtraMetadata(json.RawMessage(`{"key": "value"}`))
				_, err = suite.database.StoreProposalFinalization(ctx, finalization)
				suite.Require().NoError(err)

			},
			finalization: types.NewProposalFinalization(1).
				SetExecutionDetails(big.NewInt(10), "tx-hash-new", testStartTime).
				SetGasInfo(big.NewInt(10000), types.NewCoin("ETH2", big.NewInt(1000))).
				SetStatusTriggered(types.ProposalStatusCanceled).
				SetExtraMetadata(json.RawMessage(`{"key": "value-new"}`)),
			shouldErr: false,
			check: func() {
				var proposalFinalization postgresql.ProposalFinalizationRow
				err := suite.database.SQL.Get(&proposalFinalization, `
					SELECT * FROM proposal_finalizations
					WHERE id = $1
				`, 1)

				suite.Require().NoError(err)
				suite.Require().Equal(1, proposalFinalization.ProposalDBID)
				suite.Require().Equal("tx-hash-new", proposalFinalization.TxHash)
				suite.Require().Equal(big.NewInt(10), proposalFinalization.Height.Int)
				suite.Require().Equal(big.NewInt(10000), proposalFinalization.GasUsed.Int)
				suite.Require().Equal(types.NewCoin("ETH2", big.NewInt(1000)), proposalFinalization.GasFees.Coin)
				suite.Require().Equal(testStartTime, proposalFinalization.Timestamp.UTC())
				suite.Require().Equal(types.ProposalStatusCanceled, proposalFinalization.StatusTriggered)
				suite.Require().Equal(`{"key": "value-new"}`, string(proposalFinalization.ExtraMetadata))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			finalization, err := suite.database.StoreProposalFinalization(suite.T().Context(), tc.finalization)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.finalization, finalization)
			}
		})
	}
}

func (suite *DbTestSuite) TestStoreVoteAction() {
	testStartTime := time.Date(2025, time.August, 9, 19, 0, 0, 0, time.UTC)
	testEndTime := time.Date(2025, time.August, 16, 19, 0, 0, 0, time.UTC)
	testDescription := "Proposal description"
	testType := "Proposal type"

	testCases := []struct {
		name       string
		setup      func(context.Context)
		voteAction *types.VoteAction
		shouldErr  bool
		check      func()
	}{
		{
			name: "store vote action correctly",
			setup: func(ctx context.Context) {
				// Add a test dao
				_, err := suite.database.InsertDAO(ctx, types.NewDAO("test-dao", "Test DAO"), false)
				suite.Require().NoError(err)

				// Add a test blockchain
				_, err = suite.database.InsertBlockchain(ctx, types.NewBlockchain("1", "Ethereum", "EVM"), false)
				suite.Require().NoError(err)

				// Add a test user address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x12345678", "Test Address", false, "hex"), false)
				suite.Require().NoError(err)
				// Add a test contract address
				_, err = suite.database.InsertAddress(ctx, types.NewAddress("0x123456789", "Test Contract Address", true, "hex"), false)
				suite.Require().NoError(err)

				proposal := types.NewProposal(big.NewInt(1), 1, 1).
					SetCreatorAddressID(1).
					SetContractAddressID(2).
					SetDetails("Proposal title", &testDescription).
					SetType(&testType).
					SetStatus(types.ProposalStatusActive).
					SetCreationDetails(big.NewInt(1), "tx-hash", testStartTime).
					SetGasInfo(big.NewInt(1000), types.NewCoin("ETH", big.NewInt(100))).
					SetStartDetails(big.NewInt(1), &testStartTime).
					SetEndDetails(big.NewInt(3), &testEndTime).
					SetQuorum(big.NewInt(100)).
					SetExtraMetadata(json.RawMessage(`{"key": "value"}`))
				_, err = suite.database.StoreProposal(ctx, proposal)
				suite.Require().NoError(err)
			},
			voteAction: types.NewVoteAction(1, 1, 2).
				SetExecutionDetails(big.NewInt(1), "tx-hash", testStartTime).
				SetVote(1).
				SetVotingPower(big.NewInt(1000)).
				SetExtraMetadata(json.RawMessage(`{"key": "value"}`)),
			check: func() {
				var voteAction postgresql.VoteActionRow
				err := suite.database.SQL.Get(&voteAction, `
					SELECT * FROM vote_actions
					WHERE id = $1
				`, 1)

				suite.Require().NoError(err)
				suite.Require().Equal(types.ProposalDBID(1), voteAction.ProposalDBID)
				suite.Require().Equal(types.AddressID(1), voteAction.SenderAddressID)
				suite.Require().Equal(types.AddressID(2), voteAction.ContractAddressID)
				suite.Require().Nil(voteAction.DelegatorAddressID)
				suite.Require().Equal("tx-hash", voteAction.TxHash)
				suite.Require().Equal(big.NewInt(1000), voteAction.VotingPower.Int)
				suite.Require().Equal(big.NewInt(1), voteAction.Height.Int)
				suite.Require().Equal(testStartTime, voteAction.Timestamp.UTC())
				suite.Require().Equal(types.VoteActionTypeVote, voteAction.ActionType)
				suite.Require().Equal(int8(1), *voteAction.Vote)
				suite.Require().Equal(`{"key": "value"}`, string(voteAction.ExtraMetadata))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			_, err := suite.database.StoreVoteAction(suite.T().Context(), tc.voteAction)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check()
				}
			}
		})
	}
}
