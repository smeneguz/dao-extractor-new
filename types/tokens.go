package types

type TokenID uint64

type TokenType string

const (
	// TokenTypeNative represents a native token
	// e.g. ETH, BNB, MATIC, etc.
	TokenTypeNative TokenType = "native"
	// TokenTypeContract represents token that has been implemented
	// trough a smart contract.
	// e.g. ERC20, ERC721, ERC1155, etc.
	TokenTypeContract TokenType = "contract"
)

// Token represents a generic token
type Token interface {
	GetType() TokenType
	GetID() TokenID
	GetSymbol() string
	GetName() string
}

type BaseToken struct {
	id     TokenID
	symbol string
	name   string
}

func NewBaseToken(
	symbol string,
	name string,
) BaseToken {
	return BaseToken{
		id:     0,
		symbol: symbol,
		name:   name,
	}
}

// WithID sets the token ID.
func (t *BaseToken) WithID(id TokenID) *BaseToken {
	t.id = id
	return t
}

func (t *BaseToken) GetID() TokenID {
	return t.id
}

func (t *BaseToken) GetSymbol() string {
	return t.symbol
}

func (t *BaseToken) GetName() string {
	return t.name
}

// -----------------------------------------------------------------------------
// ---- NativeToken
// -----------------------------------------------------------------------------

// NativeToken represents a native token
type NativeToken struct {
	BaseToken
	// BlockchainID represents the ID of the blockchain where the token has been issued
	BlockchainID BlockchainID
	// Denom represents the denomination of the token (e.g. uatom, wei, etc.)
	Denom string
	// Decimals represents the number of decimals used by the token
	Decimals uint8
}

func NewNativeToken(baseToken BaseToken, blockchainID BlockchainID, denom string, decimals uint8) *NativeToken {
	return &NativeToken{
		BaseToken:    baseToken,
		BlockchainID: blockchainID,
		Denom:        denom,
		Decimals:     decimals,
	}
}

var _ Token = (*NativeToken)(nil)

func (t *NativeToken) GetType() TokenType {
	return TokenTypeNative
}

// -----------------------------------------------------------------------------
// ---- ContractToken
// -----------------------------------------------------------------------------

type TokenContractStandard string

const (
	TokenContractStandardERC20   TokenContractStandard = "erc20"
	TokenContractStandardERC721  TokenContractStandard = "erc721"
	TokenContractStandardERC1155 TokenContractStandard = "erc1155"
)

// ContractToken represents a token that has been implemented
// trough a smart contract.
// e.g. ERC20, ERC721, ERC1155, etc.
type ContractToken struct {
	BaseToken
	// BlockchainID represents the ID of the blockchain where the token has been issued
	BlockchainID BlockchainID
	// ContractAddressID represents the ID of the contract that represents the token
	ContractAddressID AddressID
	// Standard represents the standard used by the contract to represent the token
	Standard TokenContractStandard
	// Decimals represents the number of decimals used by the token
	Decimals uint8
}

var _ Token = (*ContractToken)(nil)

func NewContractToken(
	baseToken BaseToken,
	blockchainID BlockchainID,
	contractAddressID AddressID,
	standard TokenContractStandard,
	decimals uint8,
) *ContractToken {
	return &ContractToken{
		BaseToken:         baseToken,
		BlockchainID:      blockchainID,
		ContractAddressID: contractAddressID,
		Standard:          standard,
		Decimals:          decimals,
	}
}

func (t *ContractToken) GetType() TokenType {
	return TokenTypeContract
}
