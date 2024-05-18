package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"kiichain/x/kiichain/types"
)

func (k msgServer) CreateTile(goCtx context.Context, msg *types.MsgCreateTile) (*types.MsgCreateTileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var tile = types.Tile{
		Creator: msg.Creator,
		Body:    msg.Body,
	}

	id := k.AppendTile(
		ctx,
		tile,
	)

	return &types.MsgCreateTileResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTile(goCtx context.Context, msg *types.MsgUpdateTile) (*types.MsgUpdateTileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var tile = types.Tile{
		Creator: msg.Creator,
		Id:      msg.Id,
		Body:    msg.Body,
	}

	// Checks that the element exists
	val, found := k.GetTile(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTile(ctx, tile)

	return &types.MsgUpdateTileResponse{}, nil
}

func (k msgServer) DeleteTile(goCtx context.Context, msg *types.MsgDeleteTile) (*types.MsgDeleteTileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTile(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTile(ctx, msg.Id)

	return &types.MsgDeleteTileResponse{}, nil
}
