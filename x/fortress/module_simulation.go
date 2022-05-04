package fortress

import (
	"math/rand"

	"github.com/Karan-3108/Fortress/testutil/sample"
	fortresssimulation "github.com/Karan-3108/Fortress/x/fortress/simulation"
	"github.com/Karan-3108/Fortress/x/fortress/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = fortresssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestFortress = "op_weight_msg_request_fortress"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestFortress int = 100

	opWeightMsgApproveFortress = "op_weight_msg_approve_fortress"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveFortress int = 100

	opWeightMsgRepayFortress = "op_weight_msg_repay_fortress"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayFortress int = 100

	opWeightMsgLiquidateFortress = "op_weight_msg_liquidate_fortress"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiquidateFortress int = 100

	opWeightMsgCancelFortress = "op_weight_msg_cancel_fortress"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelFortress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	fortressGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&fortressGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRequestFortress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestFortress, &weightMsgRequestFortress, nil,
		func(_ *rand.Rand) {
			weightMsgRequestFortress = defaultWeightMsgRequestFortress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestFortress,
		fortresssimulation.SimulateMsgRequestFortress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveFortress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveFortress, &weightMsgApproveFortress, nil,
		func(_ *rand.Rand) {
			weightMsgApproveFortress = defaultWeightMsgApproveFortress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveFortress,
		fortresssimulation.SimulateMsgApproveFortress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayFortress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRepayFortress, &weightMsgRepayFortress, nil,
		func(_ *rand.Rand) {
			weightMsgRepayFortress = defaultWeightMsgRepayFortress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayFortress,
		fortresssimulation.SimulateMsgRepayFortress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiquidateFortress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLiquidateFortress, &weightMsgLiquidateFortress, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidateFortress = defaultWeightMsgLiquidateFortress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidateFortress,
		fortresssimulation.SimulateMsgLiquidateFortress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelFortress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelFortress, &weightMsgCancelFortress, nil,
		func(_ *rand.Rand) {
			weightMsgCancelFortress = defaultWeightMsgCancelFortress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelFortress,
		fortresssimulation.SimulateMsgCancelFortress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
