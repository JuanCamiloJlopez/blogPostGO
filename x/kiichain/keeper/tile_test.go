package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "kiichain/testutil/keeper"
	"kiichain/testutil/nullify"
	"kiichain/x/kiichain/keeper"
	"kiichain/x/kiichain/types"
)

func createNTile(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Tile {
	items := make([]types.Tile, n)
	for i := range items {
		items[i].Id = keeper.AppendTile(ctx, items[i])
	}
	return items
}

func TestTileGet(t *testing.T) {
	keeper, ctx := keepertest.KiichainKeeper(t)
	items := createNTile(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTile(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTileRemove(t *testing.T) {
	keeper, ctx := keepertest.KiichainKeeper(t)
	items := createNTile(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTile(ctx, item.Id)
		_, found := keeper.GetTile(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTileGetAll(t *testing.T) {
	keeper, ctx := keepertest.KiichainKeeper(t)
	items := createNTile(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTile(ctx)),
	)
}

func TestTileCount(t *testing.T) {
	keeper, ctx := keepertest.KiichainKeeper(t)
	items := createNTile(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTileCount(ctx))
}
