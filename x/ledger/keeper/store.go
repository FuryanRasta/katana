package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

func (k Keeper) AddPool(ctx sdk.Context, denom string, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)
	pool, ok := k.GetPool(ctx, denom)
	if !ok {
		pool = types.NewPool(denom)
	}

	pool.Addrs[addr] = true
	b := k.cdc.MustMarshal(&pool)
	store.Set([]byte(denom), b)
}

func (k Keeper) IsPoolExist(ctx sdk.Context, denom string, addr string) bool {
	pool, ok := k.GetPool(ctx, denom)
	if !ok {
		return false
	}

	return pool.Addrs[addr]
}

func (k Keeper) GetPool(ctx sdk.Context, denom string) (pool types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return pool, false
	}

	k.cdc.MustUnmarshal(b, &pool)
	return pool, true
}

func (k Keeper) RemovePool(ctx sdk.Context, denom string, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)
	pool, ok := k.GetPool(ctx, denom)
	if !ok {
		return
	}

	delete(pool.Addrs, addr)
	b := k.cdc.MustMarshal(&pool)
	store.Set([]byte(denom), b)
}

func (k Keeper) IsBondedPoolExist(ctx sdk.Context, denom string, addr string) bool {
	pool, ok := k.GetBondedPool(ctx, denom)
	if !ok {
		return false
	}

	return pool.Addrs[addr]
}

func (k Keeper) GetBondedPool(ctx sdk.Context, denom string) (pool types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return pool, false
	}

	k.cdc.MustUnmarshal(b, &pool)
	return pool, true
}

func (k Keeper) RemoveBondedPool(ctx sdk.Context, denom string, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	pool, ok := k.GetBondedPool(ctx, denom)
	if !ok {
		return
	}

	delete(pool.Addrs, addr)
	b := k.cdc.MustMarshal(&pool)
	store.Set([]byte(denom), b)
}

func (k Keeper) AddBondedPool(ctx sdk.Context, denom string, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	pool, ok := k.GetBondedPool(ctx, denom)
	if !ok {
		pool = types.NewPool(denom)
	}

	pool.Addrs[addr] = true
	b := k.cdc.MustMarshal(&pool)
	store.Set([]byte(denom), b)
}

func (k Keeper) SetBondPipeline(ctx sdk.Context, pipe types.BondPipeline) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondPipelinePrefix)
	b := k.cdc.MustMarshal(&pipe)
	store.Set([]byte(pipe.Denom+pipe.Pool), b)
}

func (k Keeper) GetBondPipeLine(ctx sdk.Context, denom string, pool string) (val types.BondPipeline, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondPipelinePrefix)

	b := store.Get([]byte(denom + pool))
	if b == nil {
		return types.NewBondPipeline(denom, pool), false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetEraUnbondLimit(ctx sdk.Context, denom string, limit uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	eul := &types.EraUnbondLimit{
		Denom: denom,
		Limit: limit,
	}
	b := k.cdc.MustMarshal(eul)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetEraUnbondLimit(ctx sdk.Context, denom string) (val types.EraUnbondLimit, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetChainBondingDuration(ctx sdk.Context, denom string, era uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainBondingDurationPrefix)
	cbd := &types.ChainBondingDuration{
		Denom: denom,
		Era:   era,
	}
	b := k.cdc.MustMarshal(cbd)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetChainBondingDuration(ctx sdk.Context, denom string) (val types.ChainBondingDuration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainBondingDurationPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetPoolDetail(ctx sdk.Context, denom string, pool string, subAccounts []string, threshold uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)
	cbd := &types.PoolDetail{
		Denom:       denom,
		Pool:        pool,
		SubAccounts: subAccounts,
		Threshold:   threshold,
	}
	b := k.cdc.MustMarshal(cbd)
	store.Set([]byte(denom+pool), b)
}

func (k Keeper) GetPoolDetail(ctx sdk.Context, denom string, pool string) (val types.PoolDetail, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)

	b := store.Get([]byte(denom + pool))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetLeastBond(ctx sdk.Context, denom string, amount sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LeastBondPrefix)
	lb := &types.LeastBond{
		Denom:  denom,
		Amount: amount,
	}
	b := k.cdc.MustMarshal(lb)
	store.Set([]byte(denom), b)
}

func (k Keeper) LeastBond(ctx sdk.Context, denom string) (val *types.LeastBond, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.LeastBondPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetCurrentEraSnapshot(ctx sdk.Context, shot types.EraSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapshotPrefix)
	b := k.cdc.MustMarshal(&shot)
	store.Set([]byte(shot.Denom), b)
}

func (k Keeper) CurrentEraSnapshots(ctx sdk.Context, denom string) types.EraSnapshot {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapshotPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return types.NewEraSnapshot(denom)
	}

	var val types.EraSnapshot
	k.cdc.MustUnmarshal(b, &val)
	return val
}

func (k Keeper) ClearCurrentEraSnapshots(ctx sdk.Context, denom string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapshotPrefix)

	cess := &types.EraSnapshot{
		Denom:   denom,
		ShotIds: [][]byte{},
	}
	b := k.cdc.MustMarshal(cess)
	store.Set([]byte(denom), b)
}

func (k Keeper) SetSnapshot(ctx sdk.Context, shotId []byte, shot types.BondSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SnapshotPrefix)
	b := k.cdc.MustMarshal(&shot)
	store.Set(shotId, b)
}

func (k Keeper) Snapshot(ctx sdk.Context, shotId []byte) (val types.BondSnapshot, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SnapshotPrefix)

	b := store.Get(shotId)
	if b == nil {
		return types.BondSnapshot{}, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetEraSnapshot(ctx sdk.Context, era uint32, shot types.EraSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraSnapshotPrefix)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	b := k.cdc.MustMarshal(&shot)
	key := append([]byte(shot.Denom), bera...)
	store.Set(key, b)
}

func (k Keeper) EraSnapshot(ctx sdk.Context, denom string, era uint32) (val types.EraSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraSnapshotPrefix)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)
	b := store.Get(key)
	if b == nil {
		return types.NewEraSnapshot(denom)
	}

	k.cdc.MustUnmarshal(b, &val)
	return
}

func (k Keeper) SetChainEra(ctx sdk.Context, denom string, era uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainEraPrefix)
	ce := &types.ChainEra{
		Denom: denom,
		Era:   era,
	}

	b := k.cdc.MustMarshal(ce)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetChainEra(ctx sdk.Context, denom string) (val types.ChainEra, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainEraPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetCommission(ctx sdk.Context, commission sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	b, _ := commission.Marshal()
	store.Set(types.CommissionPrefix, b)
}

func (k Keeper) Commission(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.CommissionPrefix)
	if b == nil {
		return sdk.ZeroDec()
	}

	var val sdk.Dec
	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}

	return val
}

func (k Keeper) SetReceiver(ctx sdk.Context, receiver sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ReceiverPrefix, receiver)
}

func (k Keeper) GetReceiver(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.ReceiverPrefix)
}

func (k Keeper) SetTotalExpectedActive(ctx sdk.Context, denom string, era uint32, active sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalExpectedActivePrefix)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)

	b, _ := active.Marshal()
	store.Set(key, b)
}

func (k Keeper) TotalExpectedActive(ctx sdk.Context, denom string, era uint32) (val sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalExpectedActivePrefix)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)
	b := store.Get(key)
	if b == nil {
		return sdk.NewInt(0)
	}

	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}

	return val
}

func (k Keeper) SetPoolUnbond(ctx sdk.Context, pu types.PoolUnbond) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondPrefix)
	b := k.cdc.MustMarshal(&pu)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, pu.Era)
	key := append([]byte(pu.Denom+pu.Pool), bera...)
	store.Set(key, b)
}

func (k Keeper) GetPoolUnbond(ctx sdk.Context, denom string, pool string, era uint32) (val types.PoolUnbond, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondPrefix)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom+pool), bera...)
	b := store.Get(key)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetUnbondFee(ctx sdk.Context, value sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	uf := types.UnbondFee{
		Value: value,
	}

	b := k.cdc.MustMarshal(&uf)
	store.Set(types.UnbondFeePrefix, b)
}

func (k Keeper) GetUnbondFee(ctx sdk.Context) (val types.UnbondFee, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.UnbondFeePrefix)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetUnbondCommission(ctx sdk.Context, value sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	b, err := value.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.UnbondCommissionPrefix, b)
}

func (k Keeper) GetUnbondCommission(ctx sdk.Context) (val sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.UnbondCommissionPrefix)
	if b == nil {
		return sdk.ZeroDec()
	}

	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}

	return
}

func (k Keeper) SetAccountUnbond(ctx sdk.Context, unbond types.AccountUnbond) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountUnbondPrefix)

	b := k.cdc.MustMarshal(&unbond)
	store.Set([]byte(unbond.Denom+unbond.Unbonder), b)
}

func (k Keeper) GetAccountUnbond(ctx sdk.Context, denom, unbonder string) (val types.AccountUnbond, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountUnbondPrefix)
	b := store.Get([]byte(denom + unbonder))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetBondRecord(ctx sdk.Context, br types.BondRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondRecordPrefix)
	b := k.cdc.MustMarshal(&br)
	store.Set([]byte(br.Denom+br.Blockhash+br.Txhash), b)
}

func (k Keeper) GetBondRecord(ctx sdk.Context, denom, blockhash, txhash string) (val types.BondRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondRecordPrefix)
	b := store.Get([]byte(denom + blockhash + txhash))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
