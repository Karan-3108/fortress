package keeper

import (
	"context"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestFortress(goCtx context.Context, msg *types.MsgRequestFortress) (*types.MsgRequestFortressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create a new Loan with the following user input
	var fortress = types.Fortress{
		Amount:     msg.Amount,
		Fee:        msg.Fee,
		Collateral: msg.Collateral,
		Deadline:   msg.Deadline,
		State:      "requested",
		Borrower:   msg.Creator,
	}

	// TODO: collateral has to be more than the amount (+fee?)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	borrower, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Get the collateral as sdk.Coins
	collateral, err := sdk.ParseCoinsNormalized(fortress.Collateral)
	if err != nil {
		panic(err)
	}

	// Use the module account as escrow account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	// Add the fortress to the keeper
	k.AppendFortress(
		ctx,
		fortress,
	)

	return &types.MsgRequestFortressResponse{}, nil
}
