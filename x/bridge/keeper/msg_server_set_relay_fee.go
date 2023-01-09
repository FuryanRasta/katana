package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warmage-sports/katana/x/bridge/types"
	sudoTypes "github.com/warmage-sports/katana/x/sudo/types"
)

func (k msgServer) SetRelayFee(goCtx context.Context, msg *types.MsgSetRelayFee) (*types.MsgSetRelayFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	chainId := uint8(msg.ChainId)
	if !k.Keeper.HasChainId(ctx, chainId) {
		return nil, types.ErrChainIdNotSupport
	}

	k.Keeper.SetRelayFee(ctx, chainId, msg.Value)

	return &types.MsgSetRelayFeeResponse{}, nil
}
