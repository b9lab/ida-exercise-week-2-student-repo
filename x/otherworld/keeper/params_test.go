package keeper_test

import (
	"testing"

	testkeeper "github.com/b9lab/other-world/testutil/keeper"
	"github.com/b9lab/other-world/x/otherworld/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OtherworldKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
