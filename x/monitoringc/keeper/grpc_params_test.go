package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/x/monitoringc/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)

	t.Run("should allow params query", func(t *testing.T) {
		params := types.DefaultParams()
		tk.MonitoringConsumerKeeper.SetParams(ctx, params)

		response, err := tk.MonitoringConsumerKeeper.Params(wctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, &types.QueryParamsResponse{Params: params}, response)

		_, err = tk.MonitoringConsumerKeeper.Params(wctx, nil)
		require.Error(t, err)
	})
}
