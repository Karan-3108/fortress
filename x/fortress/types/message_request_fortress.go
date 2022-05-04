package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestFortress = "request_fortress"

var _ sdk.Msg = &MsgRequestFortress{}

func NewMsgRequestFortress(creator string, amount string, fee string, collateral string, deadline string) *MsgRequestFortress {
	return &MsgRequestFortress{
		Creator:    creator,
		Amount:     amount,
		Fee:        fee,
		Collateral: collateral,
		Deadline:   deadline,
	}
}

func (msg *MsgRequestFortress) Route() string {
	return RouterKey
}

func (msg *MsgRequestFortress) Type() string {
	return TypeMsgRequestFortress
}

func (msg *MsgRequestFortress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestFortress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestFortress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(msg.Collateral)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}
	if amount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	if !fee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
	}
	if !collateral.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coins object")
	}
	if collateral.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is empty")
	}
	return nil
}
