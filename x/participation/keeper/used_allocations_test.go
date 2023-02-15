package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/testutil/nullify"
	"github.com/spellshape/network/testutil/sample"
	"github.com/spellshape/network/x/participation/keeper"
	"github.com/spellshape/network/x/participation/types"
)

func createNUsedAllocations(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UsedAllocations {
	items := make([]types.UsedAllocations, n)
	for i := range items {
		items[i].Address = sample.Address(r)
		items[i].NumAllocations = sample.Int(r)

		keeper.SetUsedAllocations(ctx, items[i])
	}
	return items
}

func TestUsedAllocationsGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get", func(t *testing.T) {
		items := createNUsedAllocations(tk.ParticipationKeeper, sdkCtx, 10)
		for _, item := range items {
			rst, found := tk.ParticipationKeeper.GetUsedAllocations(sdkCtx,
				item.Address,
			)
			require.True(t, found)
			require.Equal(t,
				nullify.Fill(&item),
				nullify.Fill(&rst),
			)
		}
	})
}

func TestUsedAllocationsGetAll(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get all", func(t *testing.T) {
		items := createNUsedAllocations(tk.ParticipationKeeper, sdkCtx, 10)
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(tk.ParticipationKeeper.GetAllUsedAllocations(sdkCtx)),
		)
	})
}
