package evm

import (
	"bytes"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Abi struct {
	sync.Mutex

	// ChainID represents the ID of the chain where the contract has been deployed
	ChainID string
	// ContractAddress represents the address of the contract
	ContractAddress string
	// ABI represents the JSON ABI of the contract
	ABI []byte

	// decodedABI is the ABI parsed from the ABI bytes, this is cached to avoid
	// parsing the ABI multiple times.
	decodedABI *abi.ABI
}

func NewAbi(chainID string, address string, ABI []byte) *Abi {
	return &Abi{
		ChainID:         chainID,
		ContractAddress: address,
		ABI:             ABI,
		decodedABI:      nil,
	}
}

func (a *Abi) DecodeAbi() (*abi.ABI, error) {
	a.Lock()
	defer a.Unlock()

	if a.decodedABI != nil {
		return a.decodedABI, nil
	}

	parsedABI, err := abi.JSON(bytes.NewReader(a.ABI))
	if err != nil {
		return nil, err
	}

	a.decodedABI = &parsedABI
	return a.decodedABI, nil
}
