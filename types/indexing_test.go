package types_test

import (
	"math/big"
	"testing"

	"github.com/dao-portal/extractor/types"
	"github.com/stretchr/testify/assert"
)

func TestNewHeightDeferredOperations(t *testing.T) {
	ops := []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(30), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(20), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(10), nil),
	}
	heightDeferredOperations := types.NewHeightDeferredOperations(ops)
	if len(heightDeferredOperations.Operations) != 3 {
		t.Errorf("Expected 3 operations, got %d", len(heightDeferredOperations.Operations))
	}

	ops = []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(10), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(20), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(30), nil),
	}
	// Ensure the operations are sorted by height
	assert.Equal(t, ops, heightDeferredOperations.Operations)
}

func TestHeightDeferredOperations_AppendOperation(t *testing.T) {
	ops := []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(30), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(20), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(10), nil),
	}
	heightDeferredOperations := types.NewHeightDeferredOperations(ops)
	heightDeferredOperations.AppendOperation(types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(4), nil))

	ops = []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(4), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(10), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(20), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(30), nil),
	}
	// Ensure the operations are sorted by height
	assert.Equal(t, ops, heightDeferredOperations.Operations)
}

func TestHeightDeferredOperations_PopOperationsBeforeHeight(t *testing.T) {
	ops := []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(30), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(20), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(10), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(8), nil),
	}
	heightDeferredOperations := types.NewHeightDeferredOperations(ops)

	// Pop the operations with height <= 10
	op := heightDeferredOperations.PopOperationsBeforeHeight(big.NewInt(5))
	assert.Empty(t, op)

	// Pop the operations with height <= 10
	op = heightDeferredOperations.PopOperationsBeforeHeight(big.NewInt(10))
	assert.Equal(t, []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(8), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(10), nil),
	}, op)

	// Check that the remaining operations are correct and sorted
	assert.Equal(t, []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(20), nil),
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(30), nil),
	}, heightDeferredOperations.Operations)

	// Pop the operations with height <= 20
	op = heightDeferredOperations.PopOperationsBeforeHeight(big.NewInt(20))
	assert.Equal(t, []types.HeightDeferredOperation{
		*types.NewHeightDeferredOperation("creatorKey", "type", big.NewInt(20), nil),
	}, op)
}
