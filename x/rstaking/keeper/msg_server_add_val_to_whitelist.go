package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddValToWhitelist(goCtx context.Context, msg *types.MsgAddValToWhitelist) (*types.MsgAddValToWhitelistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	valAddress, err := sdk.ValAddressFromBech32(msg.ValAddress)
	if err != nil {
		return nil, err
	}
	k.AddValAddressToWhitelist(ctx, valAddress)

	return &types.MsgAddValToWhitelistResponse{}, nil
}