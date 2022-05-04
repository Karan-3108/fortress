package fortress

import (
	"github.com/Karan-3108/Fortress/x/fortress/keeper"
	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the fortress
	for _, elem := range genState.FortressList {
		k.SetFortress(ctx, elem)
	}

	// Set fortress count
	k.SetFortressCount(ctx, genState.FortressCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.FortressList = k.GetAllFortress(ctx)
	genesis.FortressCount = k.GetFortressCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
