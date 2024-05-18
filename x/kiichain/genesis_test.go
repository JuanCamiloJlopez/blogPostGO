package kiichain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "kiichain/testutil/keeper"
	"kiichain/testutil/nullify"
	"kiichain/x/kiichain"
	"kiichain/x/kiichain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TileList: []types.Tile{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TileCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KiichainKeeper(t)
	kiichain.InitGenesis(ctx, *k, genesisState)
	got := kiichain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TileList, got.TileList)
	require.Equal(t, genesisState.TileCount, got.TileCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
