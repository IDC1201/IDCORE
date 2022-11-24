package keeper_test

import (
	"context"
	"testing"

	keepertest "IDC/testutil/keeper"
	"IDC/x/idc/keeper"
	"IDC/x/idc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdcKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
