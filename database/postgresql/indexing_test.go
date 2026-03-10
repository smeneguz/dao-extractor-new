package postgresql_test

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/dao-portal/extractor/database/postgresql"
	"github.com/dao-portal/extractor/types"
)

func (suite *DbTestSuite) TestStoreHeightDeferredOperation() {
	testCases := []struct {
		name      string
		setup     func(ctx context.Context)
		op        *types.HeightDeferredOperation
		shouldErr bool
		check     func()
	}{
		{
			name: "store a new operation without payload correctly",
			op: types.NewHeightDeferredOperation(
				"creatorKey",
				"type",
				big.NewInt(10),
				nil,
			),
			shouldErr: false,
			check: func() {
				var row postgresql.HeightDeferredOperationRow
				err := suite.database.SQL.Get(&row,
					`SELECT * FROM height_deferred_operations WHERE creator_key = $1 AND type = $2 AND height = $3`,
					"creatorKey", "type", 10,
				)
				suite.Require().NoError(err)

				op := row.ToOperation()
				suite.Require().Equal("creatorKey", op.CreatorKey)
				suite.Require().Equal(big.NewInt(10), op.Height)
				suite.Require().Equal("type", op.Type)
				suite.Require().Nil(op.Payload)
			},
		},
		{
			name: "store a new operation with payload correctly",
			op: types.NewHeightDeferredOperation(
				"creatorKey",
				"type",
				big.NewInt(10),
				json.RawMessage(`{"key": "value"}`),
			),
			shouldErr: false,
			check: func() {
				var row postgresql.HeightDeferredOperationRow
				err := suite.database.SQL.Get(&row,
					`SELECT * FROM height_deferred_operations WHERE creator_key = $1 AND type = $2 AND height = $3`,
					"creatorKey", "type", 10,
				)
				suite.Require().NoError(err)

				op := row.ToOperation()
				suite.Require().Equal("creatorKey", op.CreatorKey)
				suite.Require().Equal(big.NewInt(10), op.Height)
				suite.Require().Equal("type", op.Type)
				suite.Require().Equal(json.RawMessage(`{"key": "value"}`), op.Payload)
			},
		},
		{
			name: "store a new operation with invalid payload fails",
			op: types.NewHeightDeferredOperation(
				"creatorKey",
				"type",
				big.NewInt(10),
				json.RawMessage(`{"key": "value"`),
			),
			shouldErr: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupSuite()
			if tc.setup != nil {
				tc.setup(suite.T().Context())
			}

			err := suite.database.StoreHeightDeferredOperation(suite.T().Context(), tc.op)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check()
				}
			}
		})
	}
}
