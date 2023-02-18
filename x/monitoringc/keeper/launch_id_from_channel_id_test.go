package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/testutil/nullify"
	"github.com/spellshape/network/x/monitoringc/keeper"
	"github.com/spellshape/network/x/monitoringc/types"
)

func createNLaunchIDFromChannelID(ctx sdk.Context, keeper *keeper.Keeper, n int) []types.LaunchIDFromChannelID {
	items := make([]types.LaunchIDFromChannelID, n)
	for i := range items {
		items[i].ChannelID = strconv.Itoa(i)
		keeper.SetLaunchIDFromChannelID(ctx, items[i])
	}
	return items
}

func TestLaunchIDFromChannelIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get", func(t *testing.T) {
		items := createNLaunchIDFromChannelID(ctx, tk.MonitoringConsumerKeeper, 10)
		for _, item := range items {
			rst, found := tk.MonitoringConsumerKeeper.GetLaunchIDFromChannelID(ctx,
				item.ChannelID,
			)
			require.True(t, found)
			require.Equal(t,
				nullify.Fill(&item),
				nullify.Fill(&rst),
			)
		}
	})
}

func TestLaunchIDFromChannelIDGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get all", func(t *testing.T) {
		items := createNLaunchIDFromChannelID(ctx, tk.MonitoringConsumerKeeper, 10)
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(tk.MonitoringConsumerKeeper.GetAllLaunchIDFromChannelID(ctx)),
		)
	})
}
