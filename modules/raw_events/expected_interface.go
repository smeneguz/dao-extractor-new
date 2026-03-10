package rawevents

import (
	"context"

	"github.com/dao-portal/extractor/types"
)

// Database defines the persistence operations required by this module.
type Database interface {
	InsertBlockchain(ctx context.Context, blockchain *types.Blockchain, allowConflict bool) (*types.Blockchain, error)
	InsertAddress(ctx context.Context, address *types.Address, allowConflict bool) (*types.Address, error)
	AssociateAddressToBlockchain(ctx context.Context, address *types.Address, blockchain *types.Blockchain) error
	InsertDAO(ctx context.Context, dao *types.DAO, allowConflict bool) (*types.DAO, error)
	StoreRawEvent(ctx context.Context, event *types.RawEvent) error
}
