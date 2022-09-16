package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/b9lab/other-world/testutil/keeper"
	"github.com/b9lab/other-world/x/otherworld/keeper"
	"github.com/b9lab/other-world/x/otherworld/types"
)

func createTestWorldParams2(keeper *keeper.Keeper, ctx sdk.Context) types.WorldParams {
	item := types.WorldParams{
		Name:       "Moon",
		Gravity:    1625,
		LandSupply: 10_000,
	}
	keeper.SetWorldParams2(ctx, item)
	return item
}

func TestWorldParamsGet2(t *testing.T) {
	keeper, ctx := keepertest.OtherworldKeeper(t)
	item := createTestWorldParams2(keeper, ctx)
	rst, found := keeper.GetWorldParams2(ctx)
	require.True(t, found)
	require.EqualValues(t, item, rst)
}
