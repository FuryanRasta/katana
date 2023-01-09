package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warmage-sports/katana/x/bridge/types"
	sudoTypes "github.com/warmage-sports/katana/x/sudo/types"
)

func (k msgServer) AddChainId(goCtx context.Context, msg *types.MsgAddChainId) (*types.MsgAddChainIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.AddChainId(ctx, uint8(msg.ChainId))
	return &types.MsgAddChainIdResponse{}, nil
}
