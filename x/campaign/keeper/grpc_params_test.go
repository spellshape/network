package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/testutil/sample"
	"github.com/spellshape/network/x/campaign/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := sample.CampaignParams(r)

	t.Run("should allow query for params", func(t *testing.T) {
		tk.CampaignKeeper.SetParams(ctx, params)
		response, err := tk.CampaignKeeper.Params(wctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, &types.QueryParamsResponse{Params: params}, response)

		_, err = tk.CampaignKeeper.Params(wctx, nil)
		require.Error(t, err)
	})
}
