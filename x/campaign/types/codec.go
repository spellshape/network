package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInitializeMainnet{}, "campaign/InitializeMainnet", nil)
	cdc.RegisterConcrete(&MsgUnredeemVouchers{}, "campaign/UnredeemVouchers", nil)
	cdc.RegisterConcrete(&MsgMintVouchers{}, "campaign/MintVouchers", nil)
	cdc.RegisterConcrete(&MsgUpdateTotalShares{}, "campaign/UpdateTotalShares", nil)
	cdc.RegisterConcrete(&MsgUpdateTotalSupply{}, "campaign/UpdateTotalSupply", nil)
	cdc.RegisterConcrete(&MsgCreateCampaign{}, "campaign/CreateCampaign", nil)
	cdc.RegisterConcrete(&MsgAddShares{}, "campaign/AddShares", nil)
	cdc.RegisterConcrete(&MsgAddVestingOptions{}, "campaign/AddVestingOptions", nil)
	cdc.RegisterConcrete(&MsgRedeemVouchers{}, "campaign/RedeemVouchers", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRedeemVouchers{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnredeemVouchers{},
		&MsgInitializeMainnet{},
		&MsgUpdateTotalShares{},
		&MsgUpdateTotalSupply{},
		&MsgCreateCampaign{},
		&MsgAddShares{},
		&MsgAddVestingOptions{},
		&MsgMintVouchers{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)