package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/x/ledger/types"
	relayertypes "github.com/stafihub/stafihub/x/relayers/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) LiquidityUnbond(goCtx context.Context, msg *types.MsgLiquidityUnbond) (*types.MsgLiquidityUnbondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom := msg.Value.Denom
	ce, ok := k.Keeper.GetChainEra(ctx, denom)
	if !ok {
		return nil, types.ErrChainEraNotFound
	}

	du, ok := k.Keeper.GetChainBondingDuration(ctx, denom)
	if !ok {
		return nil, types.ErrBondingDurationNotSet
	}

	receiver := k.Keeper.GetReceiver(ctx)
	if receiver == nil {
		return nil, types.ErrNoReceiver
	}

	unbonder, _ := sdk.AccAddressFromBech32(msg.Creator)
	rbalance := k.bankKeeper.GetBalance(ctx, unbonder, denom)
	if rbalance.IsLT(msg.Value) {
		return nil, sdkerrors.ErrInsufficientFunds
	}

	pipe, ok := k.Keeper.GetBondPipeline(ctx, denom, msg.Pool)
	if !ok {
		pipe = types.NewBondPipeline(denom, msg.Pool)
	}

	cms := k.Keeper.GetUnbondCommission(ctx)
	cmsFee := cms.MulInt(msg.Value.Amount).TruncateInt()
	leftValue := msg.Value.SubAmount(cmsFee)
	balance := k.RtokenToToken(ctx, leftValue.Denom, leftValue.Amount)
	pipe.Chunk.Active = pipe.Chunk.Active.Sub(balance)
	if pipe.Chunk.Active.IsNegative() {
		return nil, sdkerrors.ErrInsufficientFunds
	}
	pipe.Chunk.Unbond = pipe.Chunk.Unbond.Add(balance)

	unlockEra := ce.Era + du.GetEra()
	chunk := types.UserUnlockChunk{Pool: msg.Pool, UnlockEra: unlockEra, Value: balance, Recipient: msg.Recipient}
	unbonds, ok := k.Keeper.GetAccountUnbond(ctx, denom, msg.Creator)
	if !ok {
		unbonds = types.NewAccountUnbond(denom, msg.Creator, []types.UserUnlockChunk{chunk})
	} else {
		unbonds.Chunks = append(unbonds.Chunks, chunk)
	}

	unbonding := types.NewUnbonding(msg.Creator, msg.Recipient, balance)
	poolUnbonds, ok := k.Keeper.GetPoolUnbond(ctx, denom, msg.Pool, unlockEra)
	if !ok {
		poolUnbonds = types.NewPoolUnbond(denom, msg.Pool, unlockEra, []types.Unbonding{unbonding})
	} else {
		eul, _ := k.Keeper.GetEraUnbondLimit(ctx, denom)
		if uint32(len(poolUnbonds.Unbondings)) > eul.Limit {
			return nil, types.ErrPoolLimitReached
		}
		poolUnbonds.Unbondings = append(poolUnbonds.Unbondings, unbonding)
	}

	unbondFee, ok := k.Keeper.GetUnbondFee(ctx, denom)
	if ok && unbondFee.Value.IsPositive() {
		feeBal := k.bankKeeper.GetBalance(ctx, unbonder, unbondFee.Value.Denom)
		if feeBal.IsLT(unbondFee.Value) {
			return nil, sdkerrors.ErrInsufficientFunds
		}

		if err := k.bankKeeper.SendCoins(ctx, unbonder, receiver, sdk.Coins{unbondFee.Value}); err != nil {
			panic(err)
		}
	}

	if cmsFee.LT(sdk.ZeroInt()) {
		cmsFeeCoin := sdk.NewCoin(denom, cmsFee)
		if err := k.bankKeeper.SendCoins(ctx, unbonder, receiver, sdk.Coins{cmsFeeCoin}); err != nil {
			panic(err)
		}
	}

	burnCoins := sdk.Coins{leftValue}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, unbonder, types.ModuleName, burnCoins); err != nil {
		panic(err)
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins); err != nil {
		panic(err)
	}

	k.Keeper.SetBondPipeline(ctx, pipe)
	k.Keeper.SetAccountUnbond(ctx, unbonds)
	k.Keeper.SetPoolUnbond(ctx, poolUnbonds)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLiquidityUnbond,
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
			sdk.NewAttribute(types.AttributeKeyUnbonder, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyPool, msg.Pool),
			sdk.NewAttribute(types.AttributeKeyUnBondAmount, leftValue.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, balance.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Recipient),
		),
	)

	return &types.MsgLiquidityUnbondResponse{}, nil
}

func (k msgServer) SubmitSignature(goCtx context.Context, msg *types.MsgSubmitSignature) (*types.MsgSubmitSignatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.relayerKeeper.IsRelayer(ctx, msg.Denom, msg.Creator) {
		return nil, relayertypes.ErrProposerNotRelayer
	}

	if _, ok := k.GetBondedPool(ctx, msg.Denom); !ok {
		return nil, types.ErrPoolNotBonded
	}

	detail, ok := k.GetPoolDetail(ctx, msg.Denom, msg.Pool)
	if !ok {
		return nil, types.ErrPoolDetailNotFound
	}

	ce, ok := k.GetChainEra(ctx, msg.Denom)
	if !ok {
		return nil, types.ErrChainEraNotFound
	}
	if msg.Era > ce.Era {
		return nil, types.ErrInvalidEra
	}

	sig, ok := k.Keeper.GetSignature(ctx, msg.Denom, msg.Era, msg.Pool, msg.TxType, msg.PropId)
	if !ok {
		sig = types.NewSignature(msg.Denom, msg.Era, msg.Pool, msg.TxType, msg.PropId)
	}

	for _, sig := range sig.Sigs {
		if sig == msg.Signature {
			return nil, types.ErrSignatureRepeated
		}
	}
	sig.Sigs = append(sig.Sigs, msg.Signature)
	k.Keeper.SetSignature(ctx, sig)

	if uint32(len(sig.Sigs)) == detail.Threshold {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSignatureEnough,
				sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
				sdk.NewAttribute(types.AttributeKeyEra, strconv.FormatUint(uint64(msg.Era), 10)),
				sdk.NewAttribute(types.AttributeKeyPool, msg.Pool),
				sdk.NewAttribute(types.AttributeKeyTxType, msg.TxType.String()),
				sdk.NewAttribute(types.AttributeKeyPropId, msg.PropId),
			),
		)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSignatureSubmitted,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyEra, strconv.FormatUint(uint64(msg.Era), 10)),
			sdk.NewAttribute(types.AttributeKeyPool, msg.Pool),
			sdk.NewAttribute(types.AttributeKeyTxType, msg.TxType.String()),
			sdk.NewAttribute(types.AttributeKeyPropId, msg.PropId),
			sdk.NewAttribute(sdk.AttributeKeySignature, msg.Signature),
		),
	)

	return &types.MsgSubmitSignatureResponse{}, nil
}
