package keeper

import (
	"github.com/b9lab/cosmos-academy-exercise-simple/x/cosmosacademyexercisesimple/types"
)

var _ types.QueryServer = Keeper{}
