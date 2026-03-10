package abi

import (
	"context"

	evmtypes "github.com/dao-portal/extractor/types/evm"
)

type ABIDatabase interface {
	SaveAbi(ctx context.Context, abi *evmtypes.Abi) error
}
