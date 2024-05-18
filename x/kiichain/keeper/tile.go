package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kiichain/x/kiichain/types"
)

// GetTileCount get the total number of tile
func (k Keeper) GetTileCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TileCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTileCount set the total number of tile
func (k Keeper) SetTileCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TileCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTile appends a tile in the store with a new id and update the count
func (k Keeper) AppendTile(
	ctx sdk.Context,
	tile types.Tile,
) uint64 {
	// Create the tile
	count := k.GetTileCount(ctx)

	// Set the ID of the appended value
	tile.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TileKey))
	appendedValue := k.cdc.MustMarshal(&tile)
	store.Set(GetTileIDBytes(tile.Id), appendedValue)

	// Update tile count
	k.SetTileCount(ctx, count+1)

	return count
}

// SetTile set a specific tile in the store
func (k Keeper) SetTile(ctx sdk.Context, tile types.Tile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TileKey))
	b := k.cdc.MustMarshal(&tile)
	store.Set(GetTileIDBytes(tile.Id), b)
}

// GetTile returns a tile from its id
func (k Keeper) GetTile(ctx sdk.Context, id uint64) (val types.Tile, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TileKey))
	b := store.Get(GetTileIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTile removes a tile from the store
func (k Keeper) RemoveTile(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TileKey))
	store.Delete(GetTileIDBytes(id))
}

// GetAllTile returns all tile
func (k Keeper) GetAllTile(ctx sdk.Context) (list []types.Tile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TileKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tile
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTileIDBytes returns the byte representation of the ID
func GetTileIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTileIDFromBytes returns ID in uint64 format from a byte array
func GetTileIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
