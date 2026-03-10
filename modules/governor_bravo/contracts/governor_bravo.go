// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// GovernorBravoDelegateStorageV1Receipt is an auto generated low-level Go binding around an user-defined struct.
type GovernorBravoDelegateStorageV1Receipt struct {
	HasVoted bool
	Support  uint8
	Votes    *big.Int
}

// GovernorBravoMetaData contains all meta data concerning the GovernorBravo contract.
var GovernorBravoMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_PROPOSAL_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTING_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_VOTING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_PROPOSAL_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_VOTING_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_VOTING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalCount\",\"type\":\"uint256\"}],\"name\":\"_initiate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"_setProposalThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"_setVotingDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"_setVotingPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteBySig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getActions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structGovernorBravoDelegateStorageV1.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"timelock_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uni_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingPeriod_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingDelay_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalThreshold_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestProposalIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalMaxOperations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"quorumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumGovernorBravoDelegateStorageV1.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"contractTimelockInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uni\",\"outputs\":[{\"internalType\":\"contractUniInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "GovernorBravo",
}

// GovernorBravo is an auto generated Go binding around an Ethereum contract.
type GovernorBravo struct {
	abi abi.ABI
}

// NewGovernorBravo creates a new instance of GovernorBravo.
func NewGovernorBravo() *GovernorBravo {
	parsed, err := GovernorBravoMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &GovernorBravo{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *GovernorBravo) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeaaa7cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (governorBravo *GovernorBravo) PackBALLOTTYPEHASH() []byte {
	enc, err := governorBravo.abi.Pack("BALLOT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeaaa7cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (governorBravo *GovernorBravo) TryPackBALLOTTYPEHASH() ([]byte, error) {
	return governorBravo.abi.Pack("BALLOT_TYPEHASH")
}

// UnpackBALLOTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (governorBravo *GovernorBravo) UnpackBALLOTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := governorBravo.abi.Unpack("BALLOT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackDOMAINTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20606b70.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (governorBravo *GovernorBravo) PackDOMAINTYPEHASH() []byte {
	enc, err := governorBravo.abi.Pack("DOMAIN_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDOMAINTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20606b70.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (governorBravo *GovernorBravo) TryPackDOMAINTYPEHASH() ([]byte, error) {
	return governorBravo.abi.Pack("DOMAIN_TYPEHASH")
}

// UnpackDOMAINTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (governorBravo *GovernorBravo) UnpackDOMAINTYPEHASH(data []byte) ([32]byte, error) {
	out, err := governorBravo.abi.Unpack("DOMAIN_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackMAXPROPOSALTHRESHOLD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25fd935a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAX_PROPOSAL_THRESHOLD() view returns(uint256)
func (governorBravo *GovernorBravo) PackMAXPROPOSALTHRESHOLD() []byte {
	enc, err := governorBravo.abi.Pack("MAX_PROPOSAL_THRESHOLD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXPROPOSALTHRESHOLD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25fd935a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAX_PROPOSAL_THRESHOLD() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackMAXPROPOSALTHRESHOLD() ([]byte, error) {
	return governorBravo.abi.Pack("MAX_PROPOSAL_THRESHOLD")
}

// UnpackMAXPROPOSALTHRESHOLD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x25fd935a.
//
// Solidity: function MAX_PROPOSAL_THRESHOLD() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackMAXPROPOSALTHRESHOLD(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("MAX_PROPOSAL_THRESHOLD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMAXVOTINGDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1126263.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAX_VOTING_DELAY() view returns(uint256)
func (governorBravo *GovernorBravo) PackMAXVOTINGDELAY() []byte {
	enc, err := governorBravo.abi.Pack("MAX_VOTING_DELAY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXVOTINGDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1126263.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAX_VOTING_DELAY() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackMAXVOTINGDELAY() ([]byte, error) {
	return governorBravo.abi.Pack("MAX_VOTING_DELAY")
}

// UnpackMAXVOTINGDELAY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb1126263.
//
// Solidity: function MAX_VOTING_DELAY() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackMAXVOTINGDELAY(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("MAX_VOTING_DELAY", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMAXVOTINGPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa64e024a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAX_VOTING_PERIOD() view returns(uint256)
func (governorBravo *GovernorBravo) PackMAXVOTINGPERIOD() []byte {
	enc, err := governorBravo.abi.Pack("MAX_VOTING_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXVOTINGPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa64e024a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAX_VOTING_PERIOD() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackMAXVOTINGPERIOD() ([]byte, error) {
	return governorBravo.abi.Pack("MAX_VOTING_PERIOD")
}

// UnpackMAXVOTINGPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa64e024a.
//
// Solidity: function MAX_VOTING_PERIOD() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackMAXVOTINGPERIOD(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("MAX_VOTING_PERIOD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMINPROPOSALTHRESHOLD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x791f5d23.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_PROPOSAL_THRESHOLD() view returns(uint256)
func (governorBravo *GovernorBravo) PackMINPROPOSALTHRESHOLD() []byte {
	enc, err := governorBravo.abi.Pack("MIN_PROPOSAL_THRESHOLD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINPROPOSALTHRESHOLD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x791f5d23.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_PROPOSAL_THRESHOLD() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackMINPROPOSALTHRESHOLD() ([]byte, error) {
	return governorBravo.abi.Pack("MIN_PROPOSAL_THRESHOLD")
}

// UnpackMINPROPOSALTHRESHOLD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x791f5d23.
//
// Solidity: function MIN_PROPOSAL_THRESHOLD() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackMINPROPOSALTHRESHOLD(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("MIN_PROPOSAL_THRESHOLD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMINVOTINGDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe48083fe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_VOTING_DELAY() view returns(uint256)
func (governorBravo *GovernorBravo) PackMINVOTINGDELAY() []byte {
	enc, err := governorBravo.abi.Pack("MIN_VOTING_DELAY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINVOTINGDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe48083fe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_VOTING_DELAY() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackMINVOTINGDELAY() ([]byte, error) {
	return governorBravo.abi.Pack("MIN_VOTING_DELAY")
}

// UnpackMINVOTINGDELAY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe48083fe.
//
// Solidity: function MIN_VOTING_DELAY() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackMINVOTINGDELAY(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("MIN_VOTING_DELAY", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMINVOTINGPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x215809ca.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_VOTING_PERIOD() view returns(uint256)
func (governorBravo *GovernorBravo) PackMINVOTINGPERIOD() []byte {
	enc, err := governorBravo.abi.Pack("MIN_VOTING_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINVOTINGPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x215809ca.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_VOTING_PERIOD() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackMINVOTINGPERIOD() ([]byte, error) {
	return governorBravo.abi.Pack("MIN_VOTING_PERIOD")
}

// UnpackMINVOTINGPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x215809ca.
//
// Solidity: function MIN_VOTING_PERIOD() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackMINVOTINGPERIOD(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("MIN_VOTING_PERIOD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackAcceptAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe9c714f2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function _acceptAdmin() returns()
func (governorBravo *GovernorBravo) PackAcceptAdmin() []byte {
	enc, err := governorBravo.abi.Pack("_acceptAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe9c714f2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function _acceptAdmin() returns()
func (governorBravo *GovernorBravo) TryPackAcceptAdmin() ([]byte, error) {
	return governorBravo.abi.Pack("_acceptAdmin")
}

// PackInitiate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x501b6ad3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function _initiate(uint256 proposalCount) returns()
func (governorBravo *GovernorBravo) PackInitiate(proposalCount *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("_initiate", proposalCount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitiate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x501b6ad3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function _initiate(uint256 proposalCount) returns()
func (governorBravo *GovernorBravo) TryPackInitiate(proposalCount *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("_initiate", proposalCount)
}

// PackSetPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb71d1a0c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (governorBravo *GovernorBravo) PackSetPendingAdmin(newPendingAdmin common.Address) []byte {
	enc, err := governorBravo.abi.Pack("_setPendingAdmin", newPendingAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb71d1a0c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (governorBravo *GovernorBravo) TryPackSetPendingAdmin(newPendingAdmin common.Address) ([]byte, error) {
	return governorBravo.abi.Pack("_setPendingAdmin", newPendingAdmin)
}

// PackSetProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17ba1b8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function _setProposalThreshold(uint256 newProposalThreshold) returns()
func (governorBravo *GovernorBravo) PackSetProposalThreshold(newProposalThreshold *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("_setProposalThreshold", newProposalThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17ba1b8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function _setProposalThreshold(uint256 newProposalThreshold) returns()
func (governorBravo *GovernorBravo) TryPackSetProposalThreshold(newProposalThreshold *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("_setProposalThreshold", newProposalThreshold)
}

// PackSetVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1dfb1b5a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function _setVotingDelay(uint256 newVotingDelay) returns()
func (governorBravo *GovernorBravo) PackSetVotingDelay(newVotingDelay *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("_setVotingDelay", newVotingDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1dfb1b5a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function _setVotingDelay(uint256 newVotingDelay) returns()
func (governorBravo *GovernorBravo) TryPackSetVotingDelay(newVotingDelay *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("_setVotingDelay", newVotingDelay)
}

// PackSetVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ea2d98c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function _setVotingPeriod(uint256 newVotingPeriod) returns()
func (governorBravo *GovernorBravo) PackSetVotingPeriod(newVotingPeriod *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("_setVotingPeriod", newVotingPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ea2d98c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function _setVotingPeriod(uint256 newVotingPeriod) returns()
func (governorBravo *GovernorBravo) TryPackSetVotingPeriod(newVotingPeriod *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("_setVotingPeriod", newVotingPeriod)
}

// PackAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf851a440.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function admin() view returns(address)
func (governorBravo *GovernorBravo) PackAdmin() []byte {
	enc, err := governorBravo.abi.Pack("admin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf851a440.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function admin() view returns(address)
func (governorBravo *GovernorBravo) TryPackAdmin() ([]byte, error) {
	return governorBravo.abi.Pack("admin")
}

// UnpackAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (governorBravo *GovernorBravo) UnpackAdmin(data []byte) (common.Address, error) {
	out, err := governorBravo.abi.Unpack("admin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40e58ee5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (governorBravo *GovernorBravo) PackCancel(proposalId *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("cancel", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40e58ee5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (governorBravo *GovernorBravo) TryPackCancel(proposalId *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("cancel", proposalId)
}

// PackCastVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56781388.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns()
func (governorBravo *GovernorBravo) PackCastVote(proposalId *big.Int, support uint8) []byte {
	enc, err := governorBravo.abi.Pack("castVote", proposalId, support)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56781388.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns()
func (governorBravo *GovernorBravo) TryPackCastVote(proposalId *big.Int, support uint8) ([]byte, error) {
	return governorBravo.abi.Pack("castVote", proposalId, support)
}

// PackCastVoteBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bccf4fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns()
func (governorBravo *GovernorBravo) PackCastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := governorBravo.abi.Pack("castVoteBySig", proposalId, support, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bccf4fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns()
func (governorBravo *GovernorBravo) TryPackCastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return governorBravo.abi.Pack("castVoteBySig", proposalId, support, v, r, s)
}

// PackCastVoteWithReason is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b3c71d3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (governorBravo *GovernorBravo) PackCastVoteWithReason(proposalId *big.Int, support uint8, reason string) []byte {
	enc, err := governorBravo.abi.Pack("castVoteWithReason", proposalId, support, reason)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteWithReason is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b3c71d3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (governorBravo *GovernorBravo) TryPackCastVoteWithReason(proposalId *big.Int, support uint8, reason string) ([]byte, error) {
	return governorBravo.abi.Pack("castVoteWithReason", proposalId, support, reason)
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe0d94c1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (governorBravo *GovernorBravo) PackExecute(proposalId *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("execute", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe0d94c1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (governorBravo *GovernorBravo) TryPackExecute(proposalId *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("execute", proposalId)
}

// PackGetActions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x328dd982.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (governorBravo *GovernorBravo) PackGetActions(proposalId *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("getActions", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetActions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x328dd982.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (governorBravo *GovernorBravo) TryPackGetActions(proposalId *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("getActions", proposalId)
}

// GetActionsOutput serves as a container for the return parameters of contract
// method GetActions.
type GetActionsOutput struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}

// UnpackGetActions is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (governorBravo *GovernorBravo) UnpackGetActions(data []byte) (GetActionsOutput, error) {
	out, err := governorBravo.abi.Unpack("getActions", data)
	outstruct := new(GetActionsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Targets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Calldatas = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)
	return *outstruct, nil
}

// PackGetReceipt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe23a9a52.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (governorBravo *GovernorBravo) PackGetReceipt(proposalId *big.Int, voter common.Address) []byte {
	enc, err := governorBravo.abi.Pack("getReceipt", proposalId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetReceipt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe23a9a52.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (governorBravo *GovernorBravo) TryPackGetReceipt(proposalId *big.Int, voter common.Address) ([]byte, error) {
	return governorBravo.abi.Pack("getReceipt", proposalId, voter)
}

// UnpackGetReceipt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (governorBravo *GovernorBravo) UnpackGetReceipt(data []byte) (GovernorBravoDelegateStorageV1Receipt, error) {
	out, err := governorBravo.abi.Unpack("getReceipt", data)
	if err != nil {
		return *new(GovernorBravoDelegateStorageV1Receipt), err
	}
	out0 := *abi.ConvertType(out[0], new(GovernorBravoDelegateStorageV1Receipt)).(*GovernorBravoDelegateStorageV1Receipt)
	return out0, nil
}

// PackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function implementation() view returns(address)
func (governorBravo *GovernorBravo) PackImplementation() []byte {
	enc, err := governorBravo.abi.Pack("implementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function implementation() view returns(address)
func (governorBravo *GovernorBravo) TryPackImplementation() ([]byte, error) {
	return governorBravo.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (governorBravo *GovernorBravo) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := governorBravo.abi.Unpack("implementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialProposalId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc4eee42.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialProposalId() view returns(uint256)
func (governorBravo *GovernorBravo) PackInitialProposalId() []byte {
	enc, err := governorBravo.abi.Pack("initialProposalId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialProposalId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc4eee42.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialProposalId() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackInitialProposalId() ([]byte, error) {
	return governorBravo.abi.Pack("initialProposalId")
}

// UnpackInitialProposalId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc4eee42.
//
// Solidity: function initialProposalId() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackInitialProposalId(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("initialProposalId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd13f90b4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address timelock_, address uni_, uint256 votingPeriod_, uint256 votingDelay_, uint256 proposalThreshold_) returns()
func (governorBravo *GovernorBravo) PackInitialize(timelock common.Address, uni common.Address, votingPeriod *big.Int, votingDelay *big.Int, proposalThreshold *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("initialize", timelock, uni, votingPeriod, votingDelay, proposalThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd13f90b4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address timelock_, address uni_, uint256 votingPeriod_, uint256 votingDelay_, uint256 proposalThreshold_) returns()
func (governorBravo *GovernorBravo) TryPackInitialize(timelock common.Address, uni common.Address, votingPeriod *big.Int, votingDelay *big.Int, proposalThreshold *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("initialize", timelock, uni, votingPeriod, votingDelay, proposalThreshold)
}

// PackLatestProposalIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17977c61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (governorBravo *GovernorBravo) PackLatestProposalIds(arg0 common.Address) []byte {
	enc, err := governorBravo.abi.Pack("latestProposalIds", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLatestProposalIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17977c61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (governorBravo *GovernorBravo) TryPackLatestProposalIds(arg0 common.Address) ([]byte, error) {
	return governorBravo.abi.Pack("latestProposalIds", arg0)
}

// UnpackLatestProposalIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (governorBravo *GovernorBravo) UnpackLatestProposalIds(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("latestProposalIds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (governorBravo *GovernorBravo) PackName() []byte {
	enc, err := governorBravo.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (governorBravo *GovernorBravo) TryPackName() ([]byte, error) {
	return governorBravo.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (governorBravo *GovernorBravo) UnpackName(data []byte) (string, error) {
	out, err := governorBravo.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x26782247.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingAdmin() view returns(address)
func (governorBravo *GovernorBravo) PackPendingAdmin() []byte {
	enc, err := governorBravo.abi.Pack("pendingAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x26782247.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingAdmin() view returns(address)
func (governorBravo *GovernorBravo) TryPackPendingAdmin() ([]byte, error) {
	return governorBravo.abi.Pack("pendingAdmin")
}

// UnpackPendingAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (governorBravo *GovernorBravo) UnpackPendingAdmin(data []byte) (common.Address, error) {
	out, err := governorBravo.abi.Unpack("pendingAdmin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProposalCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda35c664.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalCount() view returns(uint256)
func (governorBravo *GovernorBravo) PackProposalCount() []byte {
	enc, err := governorBravo.abi.Pack("proposalCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda35c664.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalCount() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackProposalCount() ([]byte, error) {
	return governorBravo.abi.Pack("proposalCount")
}

// UnpackProposalCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackProposalCount(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("proposalCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalMaxOperations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bdbe4d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalMaxOperations() view returns(uint256)
func (governorBravo *GovernorBravo) PackProposalMaxOperations() []byte {
	enc, err := governorBravo.abi.Pack("proposalMaxOperations")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalMaxOperations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bdbe4d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalMaxOperations() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackProposalMaxOperations() ([]byte, error) {
	return governorBravo.abi.Pack("proposalMaxOperations")
}

// UnpackProposalMaxOperations is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7bdbe4d0.
//
// Solidity: function proposalMaxOperations() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackProposalMaxOperations(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("proposalMaxOperations", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb58131b0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (governorBravo *GovernorBravo) PackProposalThreshold() []byte {
	enc, err := governorBravo.abi.Pack("proposalThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb58131b0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackProposalThreshold() ([]byte, error) {
	return governorBravo.abi.Pack("proposalThreshold")
}

// UnpackProposalThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackProposalThreshold(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("proposalThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x013cf08b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, uint256 abstainVotes, bool canceled, bool executed)
func (governorBravo *GovernorBravo) PackProposals(arg0 *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("proposals", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x013cf08b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, uint256 abstainVotes, bool canceled, bool executed)
func (governorBravo *GovernorBravo) TryPackProposals(arg0 *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("proposals", arg0)
}

// ProposalsOutput serves as a container for the return parameters of contract
// method Proposals.
type ProposalsOutput struct {
	Id           *big.Int
	Proposer     common.Address
	Eta          *big.Int
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	AbstainVotes *big.Int
	Canceled     bool
	Executed     bool
}

// UnpackProposals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, address proposer, uint256 eta, uint256 startBlock, uint256 endBlock, uint256 forVotes, uint256 againstVotes, uint256 abstainVotes, bool canceled, bool executed)
func (governorBravo *GovernorBravo) UnpackProposals(data []byte) (ProposalsOutput, error) {
	out, err := governorBravo.abi.Unpack("proposals", data)
	outstruct := new(ProposalsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Id = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Proposer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Eta = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.StartBlock = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.EndBlock = abi.ConvertType(out[4], new(big.Int)).(*big.Int)
	outstruct.ForVotes = abi.ConvertType(out[5], new(big.Int)).(*big.Int)
	outstruct.AgainstVotes = abi.ConvertType(out[6], new(big.Int)).(*big.Int)
	outstruct.AbstainVotes = abi.ConvertType(out[7], new(big.Int)).(*big.Int)
	outstruct.Canceled = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Executed = *abi.ConvertType(out[9], new(bool)).(*bool)
	return *outstruct, nil
}

// PackPropose is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda95691a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (governorBravo *GovernorBravo) PackPropose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) []byte {
	enc, err := governorBravo.abi.Pack("propose", targets, values, signatures, calldatas, description)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPropose is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda95691a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (governorBravo *GovernorBravo) TryPackPropose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) ([]byte, error) {
	return governorBravo.abi.Pack("propose", targets, values, signatures, calldatas, description)
}

// UnpackPropose is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (governorBravo *GovernorBravo) UnpackPropose(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("propose", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQueue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddf0b009.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function queue(uint256 proposalId) returns()
func (governorBravo *GovernorBravo) PackQueue(proposalId *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("queue", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQueue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddf0b009.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function queue(uint256 proposalId) returns()
func (governorBravo *GovernorBravo) TryPackQueue(proposalId *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("queue", proposalId)
}

// PackQuorumVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24bc1a64.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorumVotes() view returns(uint256)
func (governorBravo *GovernorBravo) PackQuorumVotes() []byte {
	enc, err := governorBravo.abi.Pack("quorumVotes")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorumVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24bc1a64.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorumVotes() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackQuorumVotes() ([]byte, error) {
	return governorBravo.abi.Pack("quorumVotes")
}

// UnpackQuorumVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x24bc1a64.
//
// Solidity: function quorumVotes() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackQuorumVotes(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("quorumVotes", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e4f49e6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (governorBravo *GovernorBravo) PackState(proposalId *big.Int) []byte {
	enc, err := governorBravo.abi.Pack("state", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e4f49e6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (governorBravo *GovernorBravo) TryPackState(proposalId *big.Int) ([]byte, error) {
	return governorBravo.abi.Pack("state", proposalId)
}

// UnpackState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (governorBravo *GovernorBravo) UnpackState(data []byte) (uint8, error) {
	out, err := governorBravo.abi.Unpack("state", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd33219b4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelock() view returns(address)
func (governorBravo *GovernorBravo) PackTimelock() []byte {
	enc, err := governorBravo.abi.Pack("timelock")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd33219b4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timelock() view returns(address)
func (governorBravo *GovernorBravo) TryPackTimelock() ([]byte, error) {
	return governorBravo.abi.Pack("timelock")
}

// UnpackTimelock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (governorBravo *GovernorBravo) UnpackTimelock(data []byte) (common.Address, error) {
	out, err := governorBravo.abi.Unpack("timelock", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUni is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xedc9af95.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uni() view returns(address)
func (governorBravo *GovernorBravo) PackUni() []byte {
	enc, err := governorBravo.abi.Pack("uni")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUni is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xedc9af95.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uni() view returns(address)
func (governorBravo *GovernorBravo) TryPackUni() ([]byte, error) {
	return governorBravo.abi.Pack("uni")
}

// UnpackUni is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xedc9af95.
//
// Solidity: function uni() view returns(address)
func (governorBravo *GovernorBravo) UnpackUni(data []byte) (common.Address, error) {
	out, err := governorBravo.abi.Unpack("uni", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3932abb1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function votingDelay() view returns(uint256)
func (governorBravo *GovernorBravo) PackVotingDelay() []byte {
	enc, err := governorBravo.abi.Pack("votingDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3932abb1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function votingDelay() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackVotingDelay() ([]byte, error) {
	return governorBravo.abi.Pack("votingDelay")
}

// UnpackVotingDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackVotingDelay(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("votingDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02a251a3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function votingPeriod() view returns(uint256)
func (governorBravo *GovernorBravo) PackVotingPeriod() []byte {
	enc, err := governorBravo.abi.Pack("votingPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02a251a3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function votingPeriod() view returns(uint256)
func (governorBravo *GovernorBravo) TryPackVotingPeriod() ([]byte, error) {
	return governorBravo.abi.Pack("votingPeriod")
}

// UnpackVotingPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (governorBravo *GovernorBravo) UnpackVotingPeriod(data []byte) (*big.Int, error) {
	out, err := governorBravo.abi.Unpack("votingPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// GovernorBravoNewAdmin represents a NewAdmin event raised by the GovernorBravo contract.
type GovernorBravoNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const GovernorBravoNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (GovernorBravoNewAdmin) ContractEventName() string {
	return GovernorBravoNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (governorBravo *GovernorBravo) UnpackNewAdminEvent(log *types.Log) (*GovernorBravoNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoNewAdmin)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoNewImplementation represents a NewImplementation event raised by the GovernorBravo contract.
type GovernorBravoNewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const GovernorBravoNewImplementationEventName = "NewImplementation"

// ContractEventName returns the user-defined event name.
func (GovernorBravoNewImplementation) ContractEventName() string {
	return GovernorBravoNewImplementationEventName
}

// UnpackNewImplementationEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (governorBravo *GovernorBravo) UnpackNewImplementationEvent(log *types.Log) (*GovernorBravoNewImplementation, error) {
	event := "NewImplementation"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoNewImplementation)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoNewPendingAdmin represents a NewPendingAdmin event raised by the GovernorBravo contract.
type GovernorBravoNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const GovernorBravoNewPendingAdminEventName = "NewPendingAdmin"

// ContractEventName returns the user-defined event name.
func (GovernorBravoNewPendingAdmin) ContractEventName() string {
	return GovernorBravoNewPendingAdminEventName
}

// UnpackNewPendingAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (governorBravo *GovernorBravo) UnpackNewPendingAdminEvent(log *types.Log) (*GovernorBravoNewPendingAdmin, error) {
	event := "NewPendingAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoNewPendingAdmin)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoProposalCanceled represents a ProposalCanceled event raised by the GovernorBravo contract.
type GovernorBravoProposalCanceled struct {
	Id  *big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const GovernorBravoProposalCanceledEventName = "ProposalCanceled"

// ContractEventName returns the user-defined event name.
func (GovernorBravoProposalCanceled) ContractEventName() string {
	return GovernorBravoProposalCanceledEventName
}

// UnpackProposalCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCanceled(uint256 id)
func (governorBravo *GovernorBravo) UnpackProposalCanceledEvent(log *types.Log) (*GovernorBravoProposalCanceled, error) {
	event := "ProposalCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoProposalCanceled)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoProposalCreated represents a ProposalCreated event raised by the GovernorBravo contract.
type GovernorBravoProposalCreated struct {
	Id          *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	StartBlock  *big.Int
	EndBlock    *big.Int
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const GovernorBravoProposalCreatedEventName = "ProposalCreated"

// ContractEventName returns the user-defined event name.
func (GovernorBravoProposalCreated) ContractEventName() string {
	return GovernorBravoProposalCreatedEventName
}

// UnpackProposalCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCreated(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (governorBravo *GovernorBravo) UnpackProposalCreatedEvent(log *types.Log) (*GovernorBravoProposalCreated, error) {
	event := "ProposalCreated"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoProposalCreated)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoProposalExecuted represents a ProposalExecuted event raised by the GovernorBravo contract.
type GovernorBravoProposalExecuted struct {
	Id  *big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const GovernorBravoProposalExecutedEventName = "ProposalExecuted"

// ContractEventName returns the user-defined event name.
func (GovernorBravoProposalExecuted) ContractEventName() string {
	return GovernorBravoProposalExecutedEventName
}

// UnpackProposalExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalExecuted(uint256 id)
func (governorBravo *GovernorBravo) UnpackProposalExecutedEvent(log *types.Log) (*GovernorBravoProposalExecuted, error) {
	event := "ProposalExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoProposalExecuted)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoProposalQueued represents a ProposalQueued event raised by the GovernorBravo contract.
type GovernorBravoProposalQueued struct {
	Id  *big.Int
	Eta *big.Int
	Raw *types.Log // Blockchain specific contextual infos
}

const GovernorBravoProposalQueuedEventName = "ProposalQueued"

// ContractEventName returns the user-defined event name.
func (GovernorBravoProposalQueued) ContractEventName() string {
	return GovernorBravoProposalQueuedEventName
}

// UnpackProposalQueuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalQueued(uint256 id, uint256 eta)
func (governorBravo *GovernorBravo) UnpackProposalQueuedEvent(log *types.Log) (*GovernorBravoProposalQueued, error) {
	event := "ProposalQueued"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoProposalQueued)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoProposalThresholdSet represents a ProposalThresholdSet event raised by the GovernorBravo contract.
type GovernorBravoProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  *types.Log // Blockchain specific contextual infos
}

const GovernorBravoProposalThresholdSetEventName = "ProposalThresholdSet"

// ContractEventName returns the user-defined event name.
func (GovernorBravoProposalThresholdSet) ContractEventName() string {
	return GovernorBravoProposalThresholdSetEventName
}

// UnpackProposalThresholdSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (governorBravo *GovernorBravo) UnpackProposalThresholdSetEvent(log *types.Log) (*GovernorBravoProposalThresholdSet, error) {
	event := "ProposalThresholdSet"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoProposalThresholdSet)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoVoteCast represents a VoteCast event raised by the GovernorBravo contract.
type GovernorBravoVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Votes      *big.Int
	Reason     string
	Raw        *types.Log // Blockchain specific contextual infos
}

const GovernorBravoVoteCastEventName = "VoteCast"

// ContractEventName returns the user-defined event name.
func (GovernorBravoVoteCast) ContractEventName() string {
	return GovernorBravoVoteCastEventName
}

// UnpackVoteCastEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 votes, string reason)
func (governorBravo *GovernorBravo) UnpackVoteCastEvent(log *types.Log) (*GovernorBravoVoteCast, error) {
	event := "VoteCast"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoVoteCast)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoVotingDelaySet represents a VotingDelaySet event raised by the GovernorBravo contract.
type GovernorBravoVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const GovernorBravoVotingDelaySetEventName = "VotingDelaySet"

// ContractEventName returns the user-defined event name.
func (GovernorBravoVotingDelaySet) ContractEventName() string {
	return GovernorBravoVotingDelaySetEventName
}

// UnpackVotingDelaySetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (governorBravo *GovernorBravo) UnpackVotingDelaySetEvent(log *types.Log) (*GovernorBravoVotingDelaySet, error) {
	event := "VotingDelaySet"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoVotingDelaySet)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GovernorBravoVotingPeriodSet represents a VotingPeriodSet event raised by the GovernorBravo contract.
type GovernorBravoVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const GovernorBravoVotingPeriodSetEventName = "VotingPeriodSet"

// ContractEventName returns the user-defined event name.
func (GovernorBravoVotingPeriodSet) ContractEventName() string {
	return GovernorBravoVotingPeriodSetEventName
}

// UnpackVotingPeriodSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (governorBravo *GovernorBravo) UnpackVotingPeriodSetEvent(log *types.Log) (*GovernorBravoVotingPeriodSet, error) {
	event := "VotingPeriodSet"
	if len(log.Topics) == 0 || log.Topics[0] != governorBravo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GovernorBravoVotingPeriodSet)
	if len(log.Data) > 0 {
		if err := governorBravo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range governorBravo.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
