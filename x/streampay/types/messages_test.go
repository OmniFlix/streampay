package types_test

import (
	"testing"
	"time"

	"github.com/OmniFlix/streampay/v2/x/streampay/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto/ed25519"
)

// TestMsgStreamSend tests if valid/invalid stream send messages are properly validated/invalidated
func TestMsgStreamSend(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())

	pk2 := ed25519.GenPrivKey().PubKey()
	addr2 := sdk.AccAddress(pk2.Address())

	// make a valid streamSend message
	streamSendMsg := func(
		after func(msg types.MsgStreamSend) types.MsgStreamSend,
		streamType types.StreamType,
		cancellable bool,
		periods []*types.Period,
	) types.MsgStreamSend {
		validMsg := *types.NewMsgStreamSend(
			addr1.String(),
			addr2.String(),
			sdk.NewInt64Coin("uspay", 100_000_000),
			streamType,
			time.Second*100,
			periods,
			false,
			sdk.NewInt64Coin("uspay", 1_000_000),
		)

		return after(validMsg)
	}
	msg := streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
		return msg
	}, types.TypeContinuous, false, nil)

	require.Equal(t, msg.Route(), types.RouterKey)
	require.Equal(t, msg.Type(), "stream_send")
	signers := msg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name  string
		msg   types.MsgStreamSend
		valid bool
	}{
		{
			name: "valid msg",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				return msg
			}, types.TypeContinuous, false, nil),
			valid: true,
		},
		{
			name: "empty sender",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				msg.Sender = ""
				return msg
			}, types.TypeContinuous, false, nil),
			valid: false,
		},
		{
			name: "invalid recipient",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				msg.Recipient = "cwiubfiewnfoew"
				return msg
			}, types.TypeContinuous, false, nil),
			valid: false,
		},
		{
			name: "invalid amount",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				msg.Amount = sdk.Coin{}
				return msg
			}, types.TypeContinuous, false, nil),
			valid: false,
		},
		{
			name: "invalid duration",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				msg.Duration = 0
				return msg
			}, types.TypeContinuous, false, nil),
			valid: false,
		},
		{
			name: "negative duration",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				msg.Duration = -100
				return msg
			}, types.TypeContinuous, false, nil),
			valid: false,
		},
		{
			name: "nil periods data for periodic stream type",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				msg.Duration = 100
				return msg
			}, types.TypePeriodic, false, nil),
			valid: false,
		},
		{
			name: "invalid periods  data for periodic stream type",
			msg: streamSendMsg(func(msg types.MsgStreamSend) types.MsgStreamSend {
				msg.Duration = 100
				return msg
			}, types.TypePeriodic, false, []*types.Period{
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   10_000_000,
					Duration: 10,
				},
			}),
			valid: false,
		},
	}

	for _, test := range tests {
		if test.valid {
			require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
		}
	}
}
