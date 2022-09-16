package types_test

import (
	"testing"

	"github.com/b9lab/other-world/x/otherworld/types"
	"github.com/stretchr/testify/require"
)

func TestCanCreateParams(t *testing.T) {
	worldParams := types.WorldParams{
		Name:       "Moon",
		Gravity:    1625,
		LandSupply: 10_000,
	}
	require.Equal(t, "Moon", worldParams.Name)
	require.Equal(t, uint64(1625), worldParams.Gravity)
	require.Equal(t, uint64(10_000), worldParams.LandSupply)
}
