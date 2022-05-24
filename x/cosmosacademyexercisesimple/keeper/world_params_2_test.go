package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/b9lab/cosmos-academy-exercise-simple/testutil/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
)

func createTestWorldParams2(keeper *keeper.Keeper, ctx sdk.Context) types.WorldParams2 {
	item := types.WorldParams2{
		Name:       "lunarland",
		Gravity:    34,
		LandSupply: 72333,
	}
	keeper.SetWorldParams2(ctx, item)
	return item
}

func TestWorldParams2GetAsWorldParams(t *testing.T) {
	keeper, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	createTestWorldParams2(keeper, ctx)
	rst, found := keeper.GetWorldParams(ctx)
	require.True(t, found)
	require.Equal(t, types.WorldParams{
		Name:       "lunarland",
		Gravity:    34,
		LandSupply: 72333,
	}, rst)
}
