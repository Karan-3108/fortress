package keeper

import (
	"context"
	"fmt"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelFortress(goCtx context.Context, msg *types.MsgCancelFortress) (*types.MsgCancelFortressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fortress, found := k.GetFortress(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if fortress.Borrower != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the borrower")
	}

	if fortress.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongFortressState, "%v", fortress.State)
	}

	borrower, _ := sdk.AccAddressFromBech32(fortress.Borrower)
	collateral, _ := sdk.ParseCoinsNormalized(fortress.Collateral)
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)

	fortress.State = "cancelled"

	k.SetFortress(ctx, fortress)

	return &types.MsgCancelFortressResponse{}, nil
}
