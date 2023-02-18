package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/testutil/nullify"
	"github.com/spellshape/network/x/reward/keeper"
	"github.com/spellshape/network/x/reward/types"
)

func createNRewardPool(k *keeper.Keeper, ctx sdk.Context, n int) []types.RewardPool {
	items := make([]types.RewardPool, n)
	for i := range items {
		items[i].LaunchID = uint64(i)
		k.SetRewardPool(ctx, items[i])
	}
	return items
}

func TestRewardPoolGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRewardPool(tk.RewardKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.RewardKeeper.GetRewardPool(ctx,
			item.LaunchID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestRewardPoolRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRewardPool(tk.RewardKeeper, ctx, 10)
	for _, item := range items {
		tk.RewardKeeper.RemoveRewardPool(ctx,
			item.LaunchID,
		)
		_, found := tk.RewardKeeper.GetRewardPool(ctx,
			item.LaunchID,
		)
		require.False(t, found)
	}
}

func TestRewardPoolGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNRewardPool(tk.RewardKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(tk.RewardKeeper.GetAllRewardPool(ctx)),
	)
}
