package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTile = "create_tile"
	TypeMsgUpdateTile = "update_tile"
	TypeMsgDeleteTile = "delete_tile"
)

var _ sdk.Msg = &MsgCreateTile{}

func NewMsgCreateTile(creator string, body string) *MsgCreateTile {
	return &MsgCreateTile{
		Creator: creator,
		Body:    body,
	}
}

func (msg *MsgCreateTile) Route() string {
	return RouterKey
}

func (msg *MsgCreateTile) Type() string {
	return TypeMsgCreateTile
}

func (msg *MsgCreateTile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTile{}

func NewMsgUpdateTile(creator string, id uint64, body string) *MsgUpdateTile {
	return &MsgUpdateTile{
		Id:      id,
		Creator: creator,
		Body:    body,
	}
}

func (msg *MsgUpdateTile) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTile) Type() string {
	return TypeMsgUpdateTile
}

func (msg *MsgUpdateTile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTile{}

func NewMsgDeleteTile(creator string, id uint64) *MsgDeleteTile {
	return &MsgDeleteTile{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTile) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTile) Type() string {
	return TypeMsgDeleteTile
}

func (msg *MsgDeleteTile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
