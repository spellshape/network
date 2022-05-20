package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/claim/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		items[i].MissionID = uint64(i)
		items[i].Weight = sdk.NewDec(r.Int63())
		keeper.SetMission(ctx, items[i])
	}
	return items
}

func TestMissionGet(t *testing.T) {
	k, ctx := keepertest.ClaimKeeper(t)
	items := createNMission(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetMission(ctx, item.MissionID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMissionRemove(t *testing.T) {
	k, ctx := keepertest.ClaimKeeper(t)
	items := createNMission(k, ctx, 10)
	for _, item := range items {
		k.RemoveMission(ctx, item.MissionID)
		_, found := k.GetMission(ctx, item.MissionID)
		require.False(t, found)
	}
}

func TestMissionGetAll(t *testing.T) {
	k, ctx := keepertest.ClaimKeeper(t)
	items := createNMission(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllMission(ctx)),
	)
}