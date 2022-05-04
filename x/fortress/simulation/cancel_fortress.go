package simulation

import (
	"math/rand"

	"github.com/Karan-3108/Fortress/x/fortress/keeper"
	"github.com/Karan-3108/Fortress/x/fortress/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCancelFortress(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelFortress{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CancelFortress simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CancelFortress simulation not implemented"), nil, nil
	}
}
