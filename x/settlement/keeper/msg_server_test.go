package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/Finschia/finschia-sdk/types"
	keepertest "github.com/Finschia/finschia/testutil/keeper"
	"github.com/Finschia/finschia/x/settlement/keeper"
	"github.com/Finschia/finschia/x/settlement/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SettlementKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
