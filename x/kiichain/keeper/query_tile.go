package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kiichain/x/kiichain/types"
)

func (k Keeper) TileAll(goCtx context.Context, req *types.QueryAllTileRequest) (*types.QueryAllTileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tiles []types.Tile
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	tileStore := prefix.NewStore(store, types.KeyPrefix(types.TileKey))

	pageRes, err := query.Paginate(tileStore, req.Pagination, func(key []byte, value []byte) error {
		var tile types.Tile
		if err := k.cdc.Unmarshal(value, &tile); err != nil {
			return err
		}

		tiles = append(tiles, tile)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTileResponse{Tile: tiles, Pagination: pageRes}, nil
}

func (k Keeper) Tile(goCtx context.Context, req *types.QueryGetTileRequest) (*types.QueryGetTileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	tile, found := k.GetTile(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTileResponse{Tile: tile}, nil
}
