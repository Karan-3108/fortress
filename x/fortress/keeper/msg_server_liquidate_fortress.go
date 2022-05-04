package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) LiquidateFortress(goCtx context.Context, msg *types.MsgLiquidateFortress) (*types.MsgLiquidateFortressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fortress, found := k.GetFortress(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if fortress.Lender != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot liquidate: not the lender")
	}

	if fortress.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongFortressState, "%v", fortress.State)
	}

	lender, _ := sdk.AccAddressFromBech32(fortress.Lender)
	collateral, _ := sdk.ParseCoinsNormalized(fortress.Collateral)

	deadline, err := strconv.ParseInt(fortress.Deadline, 10, 64)
	if err != nil {
		panic(err)
	}

	if ctx.BlockHeight() < deadline {
		return nil, sdkerrors.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)

	fortress.State = "liquidated"

	k.SetFortress(ctx, fortress)

	return &types.MsgLiquidateFortressResponse{}, nil
}
