package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestFortress{}, "fortress/RequestFortress", nil)
	cdc.RegisterConcrete(&MsgApproveFortress{}, "fortress/ApproveFortress", nil)
	cdc.RegisterConcrete(&MsgRepayFortress{}, "fortress/RepayFortress", nil)
	cdc.RegisterConcrete(&MsgLiquidateFortress{}, "fortress/LiquidateFortress", nil)
	cdc.RegisterConcrete(&MsgCancelFortress{}, "fortress/CancelFortress", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestFortress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveFortress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRepayFortress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLiquidateFortress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelFortress{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
