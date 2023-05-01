package keeper_test

import (
	"testing"

	testkeeper "github.com/Finschia/finschia/testutil/keeper"
	"github.com/Finschia/finschia/x/settlement/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SettlementKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
