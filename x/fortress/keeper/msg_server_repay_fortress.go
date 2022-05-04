package keeper

import (
	"context"
	"fmt"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RepayFortress(goCtx context.Context, msg *types.MsgRepayFortress) (*types.MsgRepayFortressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fortress, found := k.GetFortress(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if fortress.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongFortressState, "%v", fortress.State)
	}

	lender, _ := sdk.AccAddressFromBech32(fortress.Lender)
	borrower, _ := sdk.AccAddressFromBech32(fortress.Borrower)

	if msg.Creator != fortress.Borrower {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not the borrower")
	}

	amount, err := sdk.ParseCoinsNormalized(fortress.Amount)
	fee, err := sdk.ParseCoinsNormalized(fortress.Fee)
	collateral, err := sdk.ParseCoinsNormalized(fortress.Collateral)

	err = k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongFortressState, "Cannot send coins")
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongFortressState, "Cannot send coins")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongFortressState, "Cannot send coins")
	}

	fortress.State = "repayed"

	k.SetFortress(ctx, fortress)

	return &types.MsgRepayFortressResponse{}, nil
}
