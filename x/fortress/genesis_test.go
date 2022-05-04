package fortress_test

import (
	"testing"

	keepertest "github.com/Karan-3108/Fortress/testutil/keeper"
	"github.com/Karan-3108/Fortress/testutil/nullify"
	"github.com/Karan-3108/Fortress/x/fortress"
	"github.com/Karan-3108/Fortress/x/fortress/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FortressList: []types.Fortress{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		FortressCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FortressKeeper(t)
	fortress.InitGenesis(ctx, *k, genesisState)
	got := fortress.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FortressList, got.FortressList)
	require.Equal(t, genesisState.FortressCount, got.FortressCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
