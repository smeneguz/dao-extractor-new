package types

import (
	"encoding/json"
	"math/big"
	"slices"
)

// HeightDeferredOperation is a struct that represents an operation that should be executed at a certain height.
type HeightDeferredOperation struct {
	// CreatorKey is the key of of who created the operation.
	CreatorKey string
	// Type is the operation type.
	Type string
	// Height is the height at which the operation should be performed.
	Height *big.Int
	// Payload is the operation payload.
	Payload json.RawMessage
}

func NewHeightDeferredOperation(
	creatorKey string, opType string, height *big.Int, payload json.RawMessage,
) *HeightDeferredOperation {
	return &HeightDeferredOperation{
		CreatorKey: creatorKey,
		Height:     height,
		Type:       opType,
		Payload:    payload,
	}
}

// DecodePayload decodes the payload of the operation into the provided destination.
func (o *HeightDeferredOperation) DecodePayload(des any) error {
	return json.Unmarshal(o.Payload, des)
}

// -----------------------------------------------------------------------------
// ---- HeightDeferredOperations
// -----------------------------------------------------------------------------

type HeightDeferredOperations struct {
	Operations []HeightDeferredOperation
}

func NewHeightDeferredOperations(ops []HeightDeferredOperation) *HeightDeferredOperations {
	o := &HeightDeferredOperations{
		Operations: ops,
	}
	// Sort the operations
	o.sort()

	return o
}

func (o *HeightDeferredOperations) sort() {
	slices.SortFunc(o.Operations, func(a, b HeightDeferredOperation) int {
		return a.Height.Cmp(b.Height)
	})
}

func (o *HeightDeferredOperations) PopOperationsBeforeHeight(height *big.Int) []HeightDeferredOperation {
	if len(o.Operations) == 0 || o.Operations[0].Height.Cmp(height) > 0 {
		return nil
	}

	var i int
	for i = 0; i < len(o.Operations); i++ {
		if o.Operations[i].Height.Cmp(height) >= 0 {
			break
		}
	}

	var popped []HeightDeferredOperation
	var newOps []HeightDeferredOperation
	if i == len(o.Operations) {
		popped = o.Operations
	} else {
		popped = o.Operations[:i+1]
		newOps = o.Operations[i+1:]
	}
	o.Operations = newOps

	return popped
}

func (o *HeightDeferredOperations) AppendOperation(op *HeightDeferredOperation) {
	o.Operations = append(o.Operations, *op)
	o.sort()
}
