package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warmage-sports/katana/x/bridge/types"
	sudoTypes "github.com/warmage-sports/katana/x/sudo/types"
)

func (k msgServer) RmChainId(goCtx context.Context, msg *types.MsgRmChainId) (*types.MsgRmChainIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.RmChainId(ctx, uint8(msg.ChainId))

	return &types.MsgRmChainIdResponse{}, nil
}
