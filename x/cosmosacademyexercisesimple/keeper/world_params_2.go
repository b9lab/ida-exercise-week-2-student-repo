package keeper

import (
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetWorldParams set worldParams in the store
func (k Keeper) SetWorldParams2(ctx sdk.Context, worldParams types.WorldParams2) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParamsKey))
	b := k.cdc.MustMarshalBinaryBare(&worldParams)
	store.Set([]byte{0}, b)
}
