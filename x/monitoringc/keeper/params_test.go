package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/x/monitoringc/types"
)

func TestGetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get", func(t *testing.T) {
		params := types.DefaultParams()
		tk.MonitoringConsumerKeeper.SetParams(ctx, params)
		require.EqualValues(t, params, tk.MonitoringConsumerKeeper.GetParams(ctx))

		params = types.NewParams()
		tk.MonitoringConsumerKeeper.SetParams(ctx, params)
		require.EqualValues(t, params, tk.MonitoringConsumerKeeper.GetParams(ctx))
	})
}
