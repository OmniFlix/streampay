package keeper_test

import (
	"time"

	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TestStreamSendMsg tests TypeMsgMint message is emitted on a successful mint
func (suite *KeeperTestSuite) TestStreamSendMsg() {
	for _, tc := range []struct {
		sender                string
		recipient             string
		amount                sdk.Coin
		streamType            types.StreamType
		duration              time.Duration
		periods               []*types.Period
		cancellable           bool
		fee                   sdk.Coin
		valid                 bool
		expectedMessageEvents int
	}{
		{
			sender:                suite.TestAccs[0].String(),
			recipient:             suite.TestAccs[1].String(),
			amount:                sdk.NewInt64Coin(types.DefaultParams().StreamPaymentFee.Denom, 100_000_000),
			streamType:            types.TypeContinuous,
			duration:              time.Second * 100,
			periods:               nil,
			cancellable:           false,
			fee:                   sdk.NewInt64Coin(types.DefaultParams().StreamPaymentFee.Denom, 10_000_000),
			valid:                 true,
			expectedMessageEvents: 1,
		},
	} {
		suite.Run("valid case", func() {
			ctx := suite.Ctx.WithEventManager(sdk.NewEventManager())
			suite.Require().Equal(0, len(ctx.EventManager().Events()))
			// Test stream send message
			suite.msgServer.StreamSend(
				sdk.WrapSDKContext(ctx),
				types.NewMsgStreamSend(
					tc.sender,
					tc.recipient,
					tc.amount,
					tc.streamType,
					tc.duration,
					tc.periods,
					tc.cancellable,
					tc.fee,
				),
			)
			// Ensure current number and type of event is emitted
			suite.AssertEventEmitted(ctx, types.EventTypeCreateStreamPayment, tc.expectedMessageEvents)
		})
	}
}
