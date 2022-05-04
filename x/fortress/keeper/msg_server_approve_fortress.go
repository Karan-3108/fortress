package keeper

import (
	"context"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveFortress(goCtx context.Context, msg *types.MsgApproveFortress) (*types.MsgApproveFortressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fortress, found := k.GetFortress(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}

	// TODO: for some reason the error doesn't get printed to the terminal
	if fortress.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongFortressState, "%v", fortress.State)
	}

	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(fortress.Borrower)
	amount, err := sdk.ParseCoinsNormalized(fortress.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongFortressState, "Cannot parse coins in fortress amount")
	}

	k.bankKeeper.SendCoins(ctx, lender, borrower, amount)

	fortress.Lender = msg.Creator
	fortress.State = "approved"

	k.SetFortress(ctx, fortress)

	return &types.MsgApproveFortressResponse{}, nil
}
