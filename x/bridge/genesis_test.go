package bridge_test

import (
	"testing"

	keepertest "github.com/warmage-sports/katana/testutil/keeper"
	"github.com/warmage-sports/katana/testutil/nullify"
	"github.com/warmage-sports/katana/x/bridge"
	"github.com/warmage-sports/katana/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BridgeKeeper(t)
	bridge.InitGenesis(ctx, *k, genesisState)
	got := bridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
