package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelFortress = "cancel_fortress"

var _ sdk.Msg = &MsgCancelFortress{}

func NewMsgCancelFortress(creator string, id uint64) *MsgCancelFortress {
	return &MsgCancelFortress{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelFortress) Route() string {
	return RouterKey
}

func (msg *MsgCancelFortress) Type() string {
	return TypeMsgCancelFortress
}

func (msg *MsgCancelFortress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelFortress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelFortress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
