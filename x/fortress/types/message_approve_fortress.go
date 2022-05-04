package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveFortress = "approve_fortress"

var _ sdk.Msg = &MsgApproveFortress{}

func NewMsgApproveFortress(creator string, id uint64) *MsgApproveFortress {
	return &MsgApproveFortress{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgApproveFortress) Route() string {
	return RouterKey
}

func (msg *MsgApproveFortress) Type() string {
	return TypeMsgApproveFortress
}

func (msg *MsgApproveFortress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveFortress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveFortress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
