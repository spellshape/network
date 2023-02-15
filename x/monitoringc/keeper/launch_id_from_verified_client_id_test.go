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

func createNLaunchIDFromVerifiedClientID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LaunchIDFromVerifiedClientID {
	items := make([]types.LaunchIDFromVerifiedClientID, n)
	for i := range items {
		items[i].ClientID = strconv.Itoa(i)
		keeper.SetLaunchIDFromVerifiedClientID(ctx, items[i])
	}
	return items
}

func TestLaunchIDFromVerifiedClientIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get", func(t *testing.T) {
		items := createNLaunchIDFromVerifiedClientID(tk.MonitoringConsumerKeeper, ctx, 10)
		for _, item := range items {
			rst, found := tk.MonitoringConsumerKeeper.GetLaunchIDFromVerifiedClientID(ctx,
				item.ClientID,
			)
			require.True(t, found)
			require.Equal(t,
				nullify.Fill(&item),
				nullify.Fill(&rst),
			)
		}
	})
}

func TestLaunchIDFromVerifiedClientIDGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get all", func(t *testing.T) {
		items := createNLaunchIDFromVerifiedClientID(tk.MonitoringConsumerKeeper, ctx, 10)
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(tk.MonitoringConsumerKeeper.GetAllLaunchIDFromVerifiedClientID(ctx)),
		)
	})
}
