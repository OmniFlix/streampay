package types

import (
	"github.com/OmniFlix/streampay/v2/x/streampay/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

const (
	AminoTypeStreamSendMsg          = "OmniFlix/streampay/MsgStreamSend"
	AminoTypeStopStreamMsg          = "OmniFlix/streampay/MsgStopStream"
	AminoTypeClaimStreamedAmountMsg = "OmniFlix/streampay/MsgClaimStream"
	AminoTypeStreamPayment          = "OmniFlix/streampay/StreamPayment"
	AminoTypeUpdateParamsMsg        = "OmniFlix/streampay/MsgUpdateParams"
	AminoTypeParams                 = "OmniFlix/streampay/Params"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgStreamSend{}, AminoTypeStreamSendMsg)
	legacy.RegisterAminoMsg(cdc, &MsgStopStream{}, AminoTypeStopStreamMsg)
	legacy.RegisterAminoMsg(cdc, &MsgClaimStreamedAmount{}, AminoTypeClaimStreamedAmountMsg)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, AminoTypeUpdateParamsMsg)

	cdc.RegisterInterface((*exported.StreamPaymentI)(nil), nil)
	cdc.RegisterConcrete(&StreamPayment{}, AminoTypeStreamPayment, nil)
	cdc.RegisterConcrete(&Params{}, AminoTypeParams, nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgStreamSend{},
		&MsgStopStream{},
		&MsgClaimStreamedAmount{},
		&MsgUpdateParams{},
	)

	registry.RegisterImplementations((*exported.StreamPaymentI)(nil),
		&StreamPayment{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
