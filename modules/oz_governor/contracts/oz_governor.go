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

// OZGovernorMetaData contains all meta data concerning the OZGovernor contract.
var OZGovernorMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "OZGovernor",
}

// OZGovernor is an auto generated Go binding around an Ethereum contract.
type OZGovernor struct {
	abi abi.ABI
}

// NewOZGovernor creates a new instance of OZGovernor.
func NewOZGovernor() *OZGovernor {
	parsed, err := OZGovernorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &OZGovernor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *OZGovernor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackProposalDeadline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc01f9e37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (oZGovernor *OZGovernor) PackProposalDeadline(proposalId *big.Int) []byte {
	enc, err := oZGovernor.abi.Pack("proposalDeadline", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalDeadline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc01f9e37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (oZGovernor *OZGovernor) TryPackProposalDeadline(proposalId *big.Int) ([]byte, error) {
	return oZGovernor.abi.Pack("proposalDeadline", proposalId)
}

// UnpackProposalDeadline is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (oZGovernor *OZGovernor) UnpackProposalDeadline(data []byte) (*big.Int, error) {
	out, err := oZGovernor.abi.Unpack("proposalDeadline", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143489d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (oZGovernor *OZGovernor) PackProposalProposer(proposalId *big.Int) []byte {
	enc, err := oZGovernor.abi.Pack("proposalProposer", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143489d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (oZGovernor *OZGovernor) TryPackProposalProposer(proposalId *big.Int) ([]byte, error) {
	return oZGovernor.abi.Pack("proposalProposer", proposalId)
}

// UnpackProposalProposer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (oZGovernor *OZGovernor) UnpackProposalProposer(data []byte) (common.Address, error) {
	out, err := oZGovernor.abi.Unpack("proposalProposer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProposalSnapshot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d63f693.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (oZGovernor *OZGovernor) PackProposalSnapshot(proposalId *big.Int) []byte {
	enc, err := oZGovernor.abi.Pack("proposalSnapshot", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalSnapshot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d63f693.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (oZGovernor *OZGovernor) TryPackProposalSnapshot(proposalId *big.Int) ([]byte, error) {
	return oZGovernor.abi.Pack("proposalSnapshot", proposalId)
}

// UnpackProposalSnapshot is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (oZGovernor *OZGovernor) UnpackProposalSnapshot(data []byte) (*big.Int, error) {
	out, err := oZGovernor.abi.Unpack("proposalSnapshot", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorum is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ce560a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (oZGovernor *OZGovernor) PackQuorum(timepoint *big.Int) []byte {
	enc, err := oZGovernor.abi.Pack("quorum", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorum is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ce560a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (oZGovernor *OZGovernor) TryPackQuorum(timepoint *big.Int) ([]byte, error) {
	return oZGovernor.abi.Pack("quorum", timepoint)
}

// UnpackQuorum is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (oZGovernor *OZGovernor) UnpackQuorum(data []byte) (*big.Int, error) {
	out, err := oZGovernor.abi.Unpack("quorum", data)
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
func (oZGovernor *OZGovernor) PackState(proposalId *big.Int) []byte {
	enc, err := oZGovernor.abi.Pack("state", proposalId)
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
func (oZGovernor *OZGovernor) TryPackState(proposalId *big.Int) ([]byte, error) {
	return oZGovernor.abi.Pack("state", proposalId)
}

// UnpackState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (oZGovernor *OZGovernor) UnpackState(data []byte) (uint8, error) {
	out, err := oZGovernor.abi.Unpack("state", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// OZGovernorProposalCanceled represents a ProposalCanceled event raised by the OZGovernor contract.
type OZGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const OZGovernorProposalCanceledEventName = "ProposalCanceled"

// ContractEventName returns the user-defined event name.
func (OZGovernorProposalCanceled) ContractEventName() string {
	return OZGovernorProposalCanceledEventName
}

// UnpackProposalCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (oZGovernor *OZGovernor) UnpackProposalCanceledEvent(log *types.Log) (*OZGovernorProposalCanceled, error) {
	event := "ProposalCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != oZGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OZGovernorProposalCanceled)
	if len(log.Data) > 0 {
		if err := oZGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range oZGovernor.abi.Events[event].Inputs {
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

// OZGovernorProposalCreated represents a ProposalCreated event raised by the OZGovernor contract.
type OZGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const OZGovernorProposalCreatedEventName = "ProposalCreated"

// ContractEventName returns the user-defined event name.
func (OZGovernorProposalCreated) ContractEventName() string {
	return OZGovernorProposalCreatedEventName
}

// UnpackProposalCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (oZGovernor *OZGovernor) UnpackProposalCreatedEvent(log *types.Log) (*OZGovernorProposalCreated, error) {
	event := "ProposalCreated"
	if len(log.Topics) == 0 || log.Topics[0] != oZGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OZGovernorProposalCreated)
	if len(log.Data) > 0 {
		if err := oZGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range oZGovernor.abi.Events[event].Inputs {
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

// OZGovernorProposalExecuted represents a ProposalExecuted event raised by the OZGovernor contract.
type OZGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const OZGovernorProposalExecutedEventName = "ProposalExecuted"

// ContractEventName returns the user-defined event name.
func (OZGovernorProposalExecuted) ContractEventName() string {
	return OZGovernorProposalExecutedEventName
}

// UnpackProposalExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (oZGovernor *OZGovernor) UnpackProposalExecutedEvent(log *types.Log) (*OZGovernorProposalExecuted, error) {
	event := "ProposalExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != oZGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OZGovernorProposalExecuted)
	if len(log.Data) > 0 {
		if err := oZGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range oZGovernor.abi.Events[event].Inputs {
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

// OZGovernorProposalQueued represents a ProposalQueued event raised by the OZGovernor contract.
type OZGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const OZGovernorProposalQueuedEventName = "ProposalQueued"

// ContractEventName returns the user-defined event name.
func (OZGovernorProposalQueued) ContractEventName() string {
	return OZGovernorProposalQueuedEventName
}

// UnpackProposalQueuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (oZGovernor *OZGovernor) UnpackProposalQueuedEvent(log *types.Log) (*OZGovernorProposalQueued, error) {
	event := "ProposalQueued"
	if len(log.Topics) == 0 || log.Topics[0] != oZGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OZGovernorProposalQueued)
	if len(log.Data) > 0 {
		if err := oZGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range oZGovernor.abi.Events[event].Inputs {
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

// OZGovernorVoteCast represents a VoteCast event raised by the OZGovernor contract.
type OZGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        *types.Log // Blockchain specific contextual infos
}

const OZGovernorVoteCastEventName = "VoteCast"

// ContractEventName returns the user-defined event name.
func (OZGovernorVoteCast) ContractEventName() string {
	return OZGovernorVoteCastEventName
}

// UnpackVoteCastEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (oZGovernor *OZGovernor) UnpackVoteCastEvent(log *types.Log) (*OZGovernorVoteCast, error) {
	event := "VoteCast"
	if len(log.Topics) == 0 || log.Topics[0] != oZGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OZGovernorVoteCast)
	if len(log.Data) > 0 {
		if err := oZGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range oZGovernor.abi.Events[event].Inputs {
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

// OZGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the OZGovernor contract.
type OZGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const OZGovernorVoteCastWithParamsEventName = "VoteCastWithParams"

// ContractEventName returns the user-defined event name.
func (OZGovernorVoteCastWithParams) ContractEventName() string {
	return OZGovernorVoteCastWithParamsEventName
}

// UnpackVoteCastWithParamsEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (oZGovernor *OZGovernor) UnpackVoteCastWithParamsEvent(log *types.Log) (*OZGovernorVoteCastWithParams, error) {
	event := "VoteCastWithParams"
	if len(log.Topics) == 0 || log.Topics[0] != oZGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OZGovernorVoteCastWithParams)
	if len(log.Data) > 0 {
		if err := oZGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range oZGovernor.abi.Events[event].Inputs {
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
