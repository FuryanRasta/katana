package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warmage-sports/katana/x/ledger/types"
	sudotypes "github.com/warmage-sports/katana/x/sudo/types"
)

func (k msgServer) SetEraSeconds(goCtx context.Context, msg *types.MsgSetEraSeconds) (*types.MsgSetEraSecondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	rparams, found := k.Keeper.GetRParams(ctx, msg.Denom)
	if !found {
		rparams.Denom = msg.Denom
	}
	rparams.EraSeconds = msg.EraSeconds
	rparams.BondingDuration = msg.BondingDuration
	rparams.Offset = msg.Offset

	k.Keeper.SetRParams(ctx, rparams)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRParamsChanged,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyGasPrice, rparams.GasPrice),
			sdk.NewAttribute(types.AttributeKeyEraSeconds, fmt.Sprintf("%d", rparams.EraSeconds)),
			sdk.NewAttribute(types.AttributeKeyOffset, fmt.Sprintf("%d", rparams.Offset)),
			sdk.NewAttribute(types.AttributeKeyBondingDuration, fmt.Sprintf("%d", rparams.BondingDuration)),
			sdk.NewAttribute(types.AttributeKeyLeastBond, rparams.LeastBond),
		),
	)
	return &types.MsgSetEraSecondsResponse{}, nil
}
