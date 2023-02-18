package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	spntypes "github.com/spellshape/network/pkg/types"
	campaignkeeper "github.com/spellshape/network/x/campaign/keeper"
	campaigntypes "github.com/spellshape/network/x/campaign/types"
	launchkeeper "github.com/spellshape/network/x/launch/keeper"
	launchtypes "github.com/spellshape/network/x/launch/types"
	monitoringptypes "github.com/spellshape/network/x/monitoringp/types"
	participationkeeper "github.com/spellshape/network/x/participation/keeper"
	participationtypes "github.com/spellshape/network/x/participation/types"
	profilekeeper "github.com/spellshape/network/x/profile/keeper"
	rewardkeeper "github.com/spellshape/network/x/reward/keeper"
	rewardtypes "github.com/spellshape/network/x/reward/types"
)

// NewTestSetupWithMonitoringp returns a test keepers struct and servers struct with the monitoring provider module
func NewTestSetupWithMonitoringp(t testing.TB) (sdk.Context, TestKeepers, TestMsgServers) {
	return NewTestSetupWithIBCMocksMonitoringp(t, []Connection{}, []Channel{})
}

// NewTestSetupWithIBCMocksMonitoringp returns a keeper of the monitoring provider module for testing purpose with mocks for IBC keepers
func NewTestSetupWithIBCMocksMonitoringp(
	t testing.TB,
	connectionMock []Connection,
	channelMock []Channel,
) (sdk.Context, TestKeepers, TestMsgServers) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingKeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	distrKeeper := initializer.Distribution(authKeeper, bankKeeper, stakingKeeper, paramKeeper)
	upgradeKeeper := initializer.Upgrade()
	ibcKeeper := initializer.IBC(paramKeeper, stakingKeeper, *capabilityKeeper, upgradeKeeper)
	monitoringProviderKeeper := initializer.Monitoringp(
		stakingKeeper,
		*ibcKeeper,
		*capabilityKeeper,
		paramKeeper,
		connectionMock,
		channelMock,
	)
	fundraisingKeeper := initializer.Fundraising(paramKeeper, authKeeper, bankKeeper, distrKeeper)
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, distrKeeper, paramKeeper)
	rewardKeeper := initializer.Reward(authKeeper, bankKeeper, profileKeeper, launchKeeper, paramKeeper)
	campaignKeeper := initializer.Campaign(launchKeeper, profileKeeper, bankKeeper, distrKeeper, *rewardKeeper, paramKeeper, fundraisingKeeper)
	participationKeeper := initializer.Participation(paramKeeper, fundraisingKeeper, stakingKeeper)
	launchKeeper.SetCampaignKeeper(campaignKeeper)

	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time:   ExampleTimestamp,
		Height: ExampleHeight,
	}, false, log.NewNopLogger())

	// Initialize community pool
	distrKeeper.SetFeePool(ctx, distrtypes.InitialFeePool())

	// Initialize params
	distrKeeper.SetParams(ctx, distrtypes.DefaultParams())
	stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams())
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())
	rewardKeeper.SetParams(ctx, rewardtypes.DefaultParams())
	campaignKeeper.SetParams(ctx, campaigntypes.DefaultParams())
	fundraisingKeeper.SetParams(ctx, fundraisingtypes.DefaultParams())
	participationKeeper.SetParams(ctx, participationtypes.DefaultParams())
	monitoringProviderKeeper.SetParams(ctx, monitoringptypes.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	profileSrv := profilekeeper.NewMsgServerImpl(*profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(*launchKeeper)
	campaignSrv := campaignkeeper.NewMsgServerImpl(*campaignKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(*rewardKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(*participationKeeper)

	// set max shares - only set during app InitGenesis
	campaignKeeper.SetTotalShares(ctx, spntypes.TotalShareNumber)

	return ctx, TestKeepers{
			T:                        t,
			CampaignKeeper:           campaignKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringProviderKeeper: monitoringProviderKeeper,
			BankKeeper:               bankKeeper,
			IBCKeeper:                ibcKeeper,
			StakingKeeper:            stakingKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
		}, TestMsgServers{
			T:                t,
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			CampaignSrv:      campaignSrv,
			RewardSrv:        rewardSrv,
			ParticipationSrv: participationSrv,
		}
}
