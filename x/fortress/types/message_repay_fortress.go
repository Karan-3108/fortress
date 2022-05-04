package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRepayFortress = "repay_fortress"

var _ sdk.Msg = &MsgRepayFortress{}

func NewMsgRepayFortress(creator string, id uint64) *MsgRepayFortress {
	return &MsgRepayFortress{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRepayFortress) Route() string {
	return RouterKey
}

func (msg *MsgRepayFortress) Type() string {
	return TypeMsgRepayFortress
}

func (msg *MsgRepayFortress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRepayFortress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRepayFortress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
