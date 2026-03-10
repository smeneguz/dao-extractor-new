package governorbravo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"

	"github.com/dao-portal/extractor/modules/governor_bravo/contracts"
	"github.com/dao-portal/extractor/types"
)

// ProvidesInstance is implemented by generated ABI contract types.
type ProvidesInstance interface {
	Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract
}

// ContractInfo bundles a contract ABI wrapper, its persisted address and a
// live BoundContract instance.
type ContractInfo[T ProvidesInstance] struct {
	Contract T
	Address  *types.Address
	Instance *bind.BoundContract
}

// NewContractInfo creates a ContractInfo by binding the contract to the given
// backend at the address stored in *types.Address.
func NewContractInfo[T ProvidesInstance](
	backend bind.ContractBackend,
	contract T,
	address *types.Address,
) *ContractInfo[T] {
	contractAddress := common.HexToAddress(address.Address)
	instance := contract.Instance(backend, contractAddress)
	return &ContractInfo[T]{
		Contract: contract,
		Address:  address,
		Instance: instance,
	}
}

// AddressBytes returns the raw 20-byte representation of the contract address.
func (c *ContractInfo[T]) AddressBytes() []byte {
	return c.Instance.Address().Bytes()
}

// DAOInstance holds everything needed to process events for one DAO.
type DAOInstance struct {
	Config        DAOConfig
	DAO           *types.DAO
	Chain         *types.Blockchain
	GovernorBravo *ContractInfo[*contracts.GovernorBravo]
	DeferredOps   *types.HeightDeferredOperations
	DeployedAt    uint64 // earliest deployment block among this DAO's contracts
}

// FetchGovernorBravoProposalStatus is the payload stored in a deferred
// operation that will query the final proposal status.
type FetchGovernorBravoProposalStatus struct {
	ProposalID *big.Int `json:"id"`
}
