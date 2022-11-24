package idc_test

import (
	"testing"

	keepertest "IDC/testutil/keeper"
	"IDC/testutil/nullify"
	"IDC/x/idc"
	"IDC/x/idc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdcKeeper(t)
	idc.InitGenesis(ctx, *k, genesisState)
	got := idc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
