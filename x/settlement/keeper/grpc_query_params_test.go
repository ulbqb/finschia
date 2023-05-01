package keeper_test

import (
	"testing"

	sdk "github.com/Finschia/finschia-sdk/types"
	testkeeper "github.com/Finschia/finschia/testutil/keeper"
	"github.com/Finschia/finschia/x/settlement/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.SettlementKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
