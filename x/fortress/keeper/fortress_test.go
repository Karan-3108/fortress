package keeper_test

import (
	"testing"

	keepertest "github.com/Karan-3108/Fortress/testutil/keeper"
	"github.com/Karan-3108/Fortress/testutil/nullify"
	"github.com/Karan-3108/Fortress/x/fortress/keeper"
	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNFortress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Fortress {
	items := make([]types.Fortress, n)
	for i := range items {
		items[i].Id = keeper.AppendFortress(ctx, items[i])
	}
	return items
}

func TestFortressGet(t *testing.T) {
	keeper, ctx := keepertest.FortressKeeper(t)
	items := createNFortress(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFortress(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFortressRemove(t *testing.T) {
	keeper, ctx := keepertest.FortressKeeper(t)
	items := createNFortress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFortress(ctx, item.Id)
		_, found := keeper.GetFortress(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFortressGetAll(t *testing.T) {
	keeper, ctx := keepertest.FortressKeeper(t)
	items := createNFortress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFortress(ctx)),
	)
}

func TestFortressCount(t *testing.T) {
	keeper, ctx := keepertest.FortressKeeper(t)
	items := createNFortress(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFortressCount(ctx))
}
