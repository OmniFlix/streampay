package types

import (
	"github.com/OmniFlix/streampay/v2/x/streampay/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

const (
	AminoTypeStreamSendMsg          = "OmniFlix/streampay/MsgStreamSend"
	AminoTypeStopStreamMsg          = "OmniFlix/streampay/MsgStopStream"
	AminoTypeClaimStreamedAmountMsg = "OmniFlix/streampay/MsgClaimStreamedAmount"
	AminoTypeStreamPayment          = "OmniFlix/streampay/StreamPayment"
	AminoTypeUpdateParamsMsg        = "OmniFlix/streampay/MsgUpdateParams"
	AminoTypeParams                 = "OmniFlix/streampay/Params"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStreamSend{}, AminoTypeStreamSendMsg, nil)
	cdc.RegisterConcrete(&MsgStopStream{}, AminoTypeStopStreamMsg, nil)
	cdc.RegisterConcrete(&MsgClaimStreamedAmount{}, AminoTypeClaimStreamedAmountMsg, nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, AminoTypeUpdateParamsMsg, nil)

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

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()

	// Register all Amino interfaces and concrete types on the authz Amino codec
	// so that this can later be used to properly serialize MsgGrant and MsgExec
	// instances.
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
