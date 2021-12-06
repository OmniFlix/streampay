package types

import (
	"github.com/OmniFlix/payment-stream/x/paymentstream/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStreamSend{}, "OmniFlix/payment-stream/MsgStreamSend", nil)
	cdc.RegisterInterface((*exported.PaymentStreamI)(nil), nil)
	cdc.RegisterConcrete(&PaymentStream{}, "OmniFlix/payment-stream/PaymentStream", nil)

	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStreamSend{},
	)

	registry.RegisterImplementations((*exported.PaymentStreamI)(nil),
		&PaymentStream{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
