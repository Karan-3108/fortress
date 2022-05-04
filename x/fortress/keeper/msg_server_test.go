package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Karan-3108/Fortress/testutil/keeper"
	"github.com/Karan-3108/Fortress/x/fortress/keeper"
	"github.com/Karan-3108/Fortress/x/fortress/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FortressKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
