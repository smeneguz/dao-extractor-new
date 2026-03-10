package types

// -------------------------------------------------------------------------------------------------------------------
// ---- Blockchain structs
// -------------------------------------------------------------------------------------------------------------------

// BlockchainID represents a unique ID for a blockchain.
type BlockchainID uint64

type ChainType string

const (
	ChainTypeEVM ChainType = "EVM"
)

// ChainID represents a blockchain chain ID like "1" for Ethereum.
type ChainID string

// Blockchain represents a blockchain.
type Blockchain struct {
	// Unique ID of the blockchain.
	ID BlockchainID `json:"id"`
	// Chain ID of the blockchain.
	ChainID ChainID `json:"chain_id"`
	// Name of the blockchain.
	Name string `json:"name"`
	// Type of the blockchain.
	Type ChainType `json:"type"`
}

func NewBlockchain(chainID ChainID, name string, blockchainType ChainType) *Blockchain {
	return &Blockchain{
		ID:      0,
		ChainID: chainID,
		Name:    name,
		Type:    blockchainType,
	}
}

// WithID sets the ID of the blockchain.
func (b *Blockchain) WithID(id BlockchainID) *Blockchain {
	b.ID = id
	return b
}

// -------------------------------------------------------------------------------------------------------------------
// ---- Address structs
// -------------------------------------------------------------------------------------------------------------------

type AddressID uint64

type AddressEncoding string

const (
	AddressEncodingTypeHex    AddressEncoding = "hex"
	AddressEncodingTypeBech32 AddressEncoding = "bech32"
)

// Address represents an blockchain address.
type Address struct {
	// Unique ID of the address.
	ID AddressID `json:"id"`
	// The string representation of the address.
	Address string `json:"address"`
	// An human-readable label for the address.
	Label string `json:"label"`
	// Whether the address is a contract.
	IsContract bool `json:"is_contract"`
	// The encoding type of the address.
	Encoding AddressEncoding `json:"encoding"`
}

func NewAddress(address string, label string, isContract bool, encoding AddressEncoding) *Address {
	return &Address{
		ID:         0,
		Address:    address,
		Label:      label,
		IsContract: isContract,
		Encoding:   encoding,
	}
}

func (a *Address) WithID(id AddressID) *Address {
	a.ID = id
	return a
}
