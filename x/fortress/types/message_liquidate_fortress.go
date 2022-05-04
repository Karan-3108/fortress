package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLiquidateFortress = "liquidate_fortress"

var _ sdk.Msg = &MsgLiquidateFortress{}

func NewMsgLiquidateFortress(creator string, id uint64) *MsgLiquidateFortress {
	return &MsgLiquidateFortress{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgLiquidateFortress) Route() string {
	return RouterKey
}

func (msg *MsgLiquidateFortress) Type() string {
	return TypeMsgLiquidateFortress
}

func (msg *MsgLiquidateFortress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLiquidateFortress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLiquidateFortress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
