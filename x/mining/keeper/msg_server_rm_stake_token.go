package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oldfurya/furya/x/mining/types"
	sudotypes "github.com/oldfurya/furya/x/sudo/types"
)

func (k msgServer) RmStakeToken(goCtx context.Context, msg *types.MsgRmStakeToken) (*types.MsgRmStakeTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.RemoveStakeToken(ctx, msg.Denom)

	return &types.MsgRmStakeTokenResponse{}, nil
}
