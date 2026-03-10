package etherscan

import (
	"context"

	clitypes "github.com/dao-portal/flux/cli/types"
	fluxtypes "github.com/dao-portal/flux/types"

	"github.com/dao-portal/extractor/types"
)

func InjectClient(cliCtx *clitypes.CliContext, client *Client) {
	cliCtx.WithGlobalObject(types.EterscanClientKey, client)
}

func GetClient(cliCtx *clitypes.CliContext) *Client {
	client, ok := cliCtx.GetGlobalObject(types.EterscanClientKey)
	if !ok {
		panic("can't get eterscan client, it is not set")
	}
	ethClient, ok := client.(*Client)
	if !ok {
		panic("can't get eterscan client, it is not of the right type")
	}

	return ethClient
}

func GetFromContext(ctx context.Context) *Client {
	return GetClient(ctx.Value(fluxtypes.CliContextKey).(*clitypes.CliContext))
}
