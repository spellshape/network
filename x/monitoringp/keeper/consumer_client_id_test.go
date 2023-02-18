package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/testutil/nullify"
	"github.com/spellshape/network/x/monitoringp/keeper"
	"github.com/spellshape/network/x/monitoringp/types"
)

func createTestConsumerClientID(ctx sdk.Context, keeper *keeper.Keeper) types.ConsumerClientID {
	item := types.ConsumerClientID{}
	keeper.SetConsumerClientID(ctx, item)
	return item
}

func TestConsumerClientIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

	t.Run("should allow get", func(t *testing.T) {
		item := createTestConsumerClientID(ctx, tk.MonitoringProviderKeeper)
		rst, found := tk.MonitoringProviderKeeper.GetConsumerClientID(ctx)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	})
}
