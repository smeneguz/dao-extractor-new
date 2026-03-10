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

// TimelockMetaData contains all meta data concerning the Timelock contract.
var TimelockMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"CancelTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ExecuteTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"NewDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"QueueTransaction\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"GRACE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAXIMUM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"cancelTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"queueTransaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"queuedTransactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay_\",\"type\":\"uint256\"}],\"name\":\"setDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingAdmin_\",\"type\":\"address\"}],\"name\":\"setPendingAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Timelock",
}

// Timelock is an auto generated Go binding around an Ethereum contract.
type Timelock struct {
	abi abi.ABI
}

// NewTimelock creates a new instance of Timelock.
func NewTimelock() *Timelock {
	parsed, err := TimelockMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Timelock{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Timelock) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address admin_, uint256 delay_) returns()
func (timelock *Timelock) PackConstructor(admin_ common.Address, delay_ *big.Int) []byte {
	enc, err := timelock.abi.Pack("", admin_, delay_)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGRACEPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1a287e2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (timelock *Timelock) PackGRACEPERIOD() []byte {
	enc, err := timelock.abi.Pack("GRACE_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGRACEPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1a287e2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (timelock *Timelock) TryPackGRACEPERIOD() ([]byte, error) {
	return timelock.abi.Pack("GRACE_PERIOD")
}

// UnpackGRACEPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc1a287e2.
//
// Solidity: function GRACE_PERIOD() view returns(uint256)
func (timelock *Timelock) UnpackGRACEPERIOD(data []byte) (*big.Int, error) {
	out, err := timelock.abi.Unpack("GRACE_PERIOD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMAXIMUMDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d645fab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAXIMUM_DELAY() view returns(uint256)
func (timelock *Timelock) PackMAXIMUMDELAY() []byte {
	enc, err := timelock.abi.Pack("MAXIMUM_DELAY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXIMUMDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d645fab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAXIMUM_DELAY() view returns(uint256)
func (timelock *Timelock) TryPackMAXIMUMDELAY() ([]byte, error) {
	return timelock.abi.Pack("MAXIMUM_DELAY")
}

// UnpackMAXIMUMDELAY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d645fab.
//
// Solidity: function MAXIMUM_DELAY() view returns(uint256)
func (timelock *Timelock) UnpackMAXIMUMDELAY(data []byte) (*big.Int, error) {
	out, err := timelock.abi.Unpack("MAXIMUM_DELAY", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMINIMUMDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1b43ae5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MINIMUM_DELAY() view returns(uint256)
func (timelock *Timelock) PackMINIMUMDELAY() []byte {
	enc, err := timelock.abi.Pack("MINIMUM_DELAY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINIMUMDELAY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1b43ae5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MINIMUM_DELAY() view returns(uint256)
func (timelock *Timelock) TryPackMINIMUMDELAY() ([]byte, error) {
	return timelock.abi.Pack("MINIMUM_DELAY")
}

// UnpackMINIMUMDELAY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb1b43ae5.
//
// Solidity: function MINIMUM_DELAY() view returns(uint256)
func (timelock *Timelock) UnpackMINIMUMDELAY(data []byte) (*big.Int, error) {
	out, err := timelock.abi.Unpack("MINIMUM_DELAY", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackAcceptAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e18b681.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function acceptAdmin() returns()
func (timelock *Timelock) PackAcceptAdmin() []byte {
	enc, err := timelock.abi.Pack("acceptAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAcceptAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e18b681.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function acceptAdmin() returns()
func (timelock *Timelock) TryPackAcceptAdmin() ([]byte, error) {
	return timelock.abi.Pack("acceptAdmin")
}

// PackAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf851a440.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function admin() view returns(address)
func (timelock *Timelock) PackAdmin() []byte {
	enc, err := timelock.abi.Pack("admin")
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
func (timelock *Timelock) TryPackAdmin() ([]byte, error) {
	return timelock.abi.Pack("admin")
}

// UnpackAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (timelock *Timelock) UnpackAdmin(data []byte) (common.Address, error) {
	out, err := timelock.abi.Unpack("admin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCancelTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x591fcdfe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns()
func (timelock *Timelock) PackCancelTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) []byte {
	enc, err := timelock.abi.Pack("cancelTransaction", target, value, signature, data, eta)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x591fcdfe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns()
func (timelock *Timelock) TryPackCancelTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) ([]byte, error) {
	return timelock.abi.Pack("cancelTransaction", target, value, signature, data, eta)
}

// PackDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6a42b8f8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function delay() view returns(uint256)
func (timelock *Timelock) PackDelay() []byte {
	enc, err := timelock.abi.Pack("delay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6a42b8f8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function delay() view returns(uint256)
func (timelock *Timelock) TryPackDelay() ([]byte, error) {
	return timelock.abi.Pack("delay")
}

// UnpackDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (timelock *Timelock) UnpackDelay(data []byte) (*big.Int, error) {
	out, err := timelock.abi.Unpack("delay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExecuteTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0825f38f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) payable returns(bytes)
func (timelock *Timelock) PackExecuteTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) []byte {
	enc, err := timelock.abi.Pack("executeTransaction", target, value, signature, data, eta)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0825f38f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) payable returns(bytes)
func (timelock *Timelock) TryPackExecuteTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) ([]byte, error) {
	return timelock.abi.Pack("executeTransaction", target, value, signature, data, eta)
}

// UnpackExecuteTransaction is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0825f38f.
//
// Solidity: function executeTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) payable returns(bytes)
func (timelock *Timelock) UnpackExecuteTransaction(data []byte) ([]byte, error) {
	out, err := timelock.abi.Unpack("executeTransaction", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x26782247.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingAdmin() view returns(address)
func (timelock *Timelock) PackPendingAdmin() []byte {
	enc, err := timelock.abi.Pack("pendingAdmin")
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
func (timelock *Timelock) TryPackPendingAdmin() ([]byte, error) {
	return timelock.abi.Pack("pendingAdmin")
}

// UnpackPendingAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (timelock *Timelock) UnpackPendingAdmin(data []byte) (common.Address, error) {
	out, err := timelock.abi.Unpack("pendingAdmin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackQueueTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a66f901.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function queueTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns(bytes32)
func (timelock *Timelock) PackQueueTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) []byte {
	enc, err := timelock.abi.Pack("queueTransaction", target, value, signature, data, eta)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQueueTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a66f901.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function queueTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns(bytes32)
func (timelock *Timelock) TryPackQueueTransaction(target common.Address, value *big.Int, signature string, data []byte, eta *big.Int) ([]byte, error) {
	return timelock.abi.Pack("queueTransaction", target, value, signature, data, eta)
}

// UnpackQueueTransaction is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a66f901.
//
// Solidity: function queueTransaction(address target, uint256 value, string signature, bytes data, uint256 eta) returns(bytes32)
func (timelock *Timelock) UnpackQueueTransaction(data []byte) ([32]byte, error) {
	out, err := timelock.abi.Unpack("queueTransaction", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackQueuedTransactions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2b06537.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (timelock *Timelock) PackQueuedTransactions(arg0 [32]byte) []byte {
	enc, err := timelock.abi.Pack("queuedTransactions", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQueuedTransactions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2b06537.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (timelock *Timelock) TryPackQueuedTransactions(arg0 [32]byte) ([]byte, error) {
	return timelock.abi.Pack("queuedTransactions", arg0)
}

// UnpackQueuedTransactions is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(bool)
func (timelock *Timelock) UnpackQueuedTransactions(data []byte) (bool, error) {
	out, err := timelock.abi.Unpack("queuedTransactions", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe177246e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (timelock *Timelock) PackSetDelay(delay *big.Int) []byte {
	enc, err := timelock.abi.Pack("setDelay", delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe177246e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setDelay(uint256 delay_) returns()
func (timelock *Timelock) TryPackSetDelay(delay *big.Int) ([]byte, error) {
	return timelock.abi.Pack("setDelay", delay)
}

// PackSetPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4dd18bf5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setPendingAdmin(address pendingAdmin_) returns()
func (timelock *Timelock) PackSetPendingAdmin(pendingAdmin common.Address) []byte {
	enc, err := timelock.abi.Pack("setPendingAdmin", pendingAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPendingAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4dd18bf5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setPendingAdmin(address pendingAdmin_) returns()
func (timelock *Timelock) TryPackSetPendingAdmin(pendingAdmin common.Address) ([]byte, error) {
	return timelock.abi.Pack("setPendingAdmin", pendingAdmin)
}

// TimelockCancelTransaction represents a CancelTransaction event raised by the Timelock contract.
type TimelockCancelTransaction struct {
	TxHash    [32]byte
	Target    common.Address
	Value     *big.Int
	Signature string
	Data      []byte
	Eta       *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const TimelockCancelTransactionEventName = "CancelTransaction"

// ContractEventName returns the user-defined event name.
func (TimelockCancelTransaction) ContractEventName() string {
	return TimelockCancelTransactionEventName
}

// UnpackCancelTransactionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CancelTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (timelock *Timelock) UnpackCancelTransactionEvent(log *types.Log) (*TimelockCancelTransaction, error) {
	event := "CancelTransaction"
	if len(log.Topics) == 0 || log.Topics[0] != timelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockCancelTransaction)
	if len(log.Data) > 0 {
		if err := timelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelock.abi.Events[event].Inputs {
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

// TimelockExecuteTransaction represents a ExecuteTransaction event raised by the Timelock contract.
type TimelockExecuteTransaction struct {
	TxHash    [32]byte
	Target    common.Address
	Value     *big.Int
	Signature string
	Data      []byte
	Eta       *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const TimelockExecuteTransactionEventName = "ExecuteTransaction"

// ContractEventName returns the user-defined event name.
func (TimelockExecuteTransaction) ContractEventName() string {
	return TimelockExecuteTransactionEventName
}

// UnpackExecuteTransactionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExecuteTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (timelock *Timelock) UnpackExecuteTransactionEvent(log *types.Log) (*TimelockExecuteTransaction, error) {
	event := "ExecuteTransaction"
	if len(log.Topics) == 0 || log.Topics[0] != timelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockExecuteTransaction)
	if len(log.Data) > 0 {
		if err := timelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelock.abi.Events[event].Inputs {
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

// TimelockNewAdmin represents a NewAdmin event raised by the Timelock contract.
type TimelockNewAdmin struct {
	NewAdmin common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const TimelockNewAdminEventName = "NewAdmin"

// ContractEventName returns the user-defined event name.
func (TimelockNewAdmin) ContractEventName() string {
	return TimelockNewAdminEventName
}

// UnpackNewAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewAdmin(address indexed newAdmin)
func (timelock *Timelock) UnpackNewAdminEvent(log *types.Log) (*TimelockNewAdmin, error) {
	event := "NewAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != timelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockNewAdmin)
	if len(log.Data) > 0 {
		if err := timelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelock.abi.Events[event].Inputs {
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

// TimelockNewDelay represents a NewDelay event raised by the Timelock contract.
type TimelockNewDelay struct {
	NewDelay *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const TimelockNewDelayEventName = "NewDelay"

// ContractEventName returns the user-defined event name.
func (TimelockNewDelay) ContractEventName() string {
	return TimelockNewDelayEventName
}

// UnpackNewDelayEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewDelay(uint256 indexed newDelay)
func (timelock *Timelock) UnpackNewDelayEvent(log *types.Log) (*TimelockNewDelay, error) {
	event := "NewDelay"
	if len(log.Topics) == 0 || log.Topics[0] != timelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockNewDelay)
	if len(log.Data) > 0 {
		if err := timelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelock.abi.Events[event].Inputs {
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

// TimelockNewPendingAdmin represents a NewPendingAdmin event raised by the Timelock contract.
type TimelockNewPendingAdmin struct {
	NewPendingAdmin common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const TimelockNewPendingAdminEventName = "NewPendingAdmin"

// ContractEventName returns the user-defined event name.
func (TimelockNewPendingAdmin) ContractEventName() string {
	return TimelockNewPendingAdminEventName
}

// UnpackNewPendingAdminEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewPendingAdmin(address indexed newPendingAdmin)
func (timelock *Timelock) UnpackNewPendingAdminEvent(log *types.Log) (*TimelockNewPendingAdmin, error) {
	event := "NewPendingAdmin"
	if len(log.Topics) == 0 || log.Topics[0] != timelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockNewPendingAdmin)
	if len(log.Data) > 0 {
		if err := timelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelock.abi.Events[event].Inputs {
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

// TimelockQueueTransaction represents a QueueTransaction event raised by the Timelock contract.
type TimelockQueueTransaction struct {
	TxHash    [32]byte
	Target    common.Address
	Value     *big.Int
	Signature string
	Data      []byte
	Eta       *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const TimelockQueueTransactionEventName = "QueueTransaction"

// ContractEventName returns the user-defined event name.
func (TimelockQueueTransaction) ContractEventName() string {
	return TimelockQueueTransactionEventName
}

// UnpackQueueTransactionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QueueTransaction(bytes32 indexed txHash, address indexed target, uint256 value, string signature, bytes data, uint256 eta)
func (timelock *Timelock) UnpackQueueTransactionEvent(log *types.Log) (*TimelockQueueTransaction, error) {
	event := "QueueTransaction"
	if len(log.Topics) == 0 || log.Topics[0] != timelock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimelockQueueTransaction)
	if len(log.Data) > 0 {
		if err := timelock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timelock.abi.Events[event].Inputs {
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
