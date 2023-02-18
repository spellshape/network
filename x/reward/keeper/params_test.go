package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/spellshape/network/testutil/keeper"
	"github.com/spellshape/network/x/reward/types"
)

func TestGetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	params := types.DefaultParams()

	tk.RewardKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, tk.RewardKeeper.GetParams(ctx))
}
