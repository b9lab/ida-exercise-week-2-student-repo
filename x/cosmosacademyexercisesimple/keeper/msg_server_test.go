package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/b9lab/cosmos-academy-exercise-simple/testutil/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/keeper"
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosacademyexercisesimpleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
