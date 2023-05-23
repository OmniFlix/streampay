package types

import (
	"github.com/OmniFlix/streampay/x/streampay/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

const (
	AminoTypeStreamSendMsg          = "OmniFlix/streampay/MsgStreamSend"
	AminoTypeStopStreamMsg          = "OmniFlix/streampay/MsgStopStream"
	AminoTypeClaimStreamedAmountMsg = "OmniFlix/streampay/MsgClaimStreamedAmount"
	AminoTypeStreamPayment          = "OmniFlix/streampay/StreamPayment"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStreamSend{}, AminoTypeStreamSendMsg, nil)
	cdc.RegisterConcrete(&MsgStopStream{}, AminoTypeStopStreamMsg, nil)
	cdc.RegisterConcrete(&MsgClaimStreamedAmount{}, AminoTypeClaimStreamedAmountMsg, nil)

	cdc.RegisterInterface((*exported.StreamPaymentI)(nil), nil)
	cdc.RegisterConcrete(&StreamPayment{}, AminoTypeStreamPayment, nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStreamSend{},
		&MsgStopStream{},
		&MsgClaimStreamedAmount{},
	)

	registry.RegisterImplementations((*exported.StreamPaymentI)(nil),
		&StreamPayment{},
	)

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
