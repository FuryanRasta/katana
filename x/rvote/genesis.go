package rvote

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warmage-sports/katana/x/rvote/keeper"
	"github.com/warmage-sports/katana/x/rvote/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetProposalLife(ctx, genState.ProposalLife)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.ProposalLife = k.ProposalLife(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
