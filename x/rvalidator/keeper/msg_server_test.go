package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/oldfurya/furya/testutil/keeper"
	"github.com/oldfurya/furya/x/rvalidator/keeper"
	"github.com/oldfurya/furya/x/rvalidator/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RvalidatorKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
