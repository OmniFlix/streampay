package keeper_test

import (
	"time"

	"github.com/OmniFlix/streampay/v2/x/streampay/types"
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
		paymentFee            sdk.Coin
		valid                 bool
		expectedMessageEvents int
	}{
		{
			sender:                suite.TestAccs[0].String(),
			recipient:             suite.TestAccs[1].String(),
			amount:                sdk.NewInt64Coin("uspay", 100_000_000),
			streamType:            types.TypeContinuous,
			duration:              time.Second * 100,
			periods:               nil,
			cancellable:           false,
			paymentFee:            sdk.NewInt64Coin("uspay", 1_000_000),
			valid:                 true,
			expectedMessageEvents: 1,
		},
		{
			sender:                suite.TestAccs[0].String(),
			recipient:             suite.TestAccs[1].String(),
			amount:                sdk.NewInt64Coin("uspay", 100_000_000),
			streamType:            types.TypeContinuous,
			duration:              time.Second * 100,
			periods:               nil,
			cancellable:           false,
			paymentFee:            sdk.NewInt64Coin("uspay", 10_000),
			valid:                 false,
			expectedMessageEvents: 0,
		},
		{
			sender:                suite.TestAccs[0].String(),
			recipient:             suite.TestAccs[1].String(),
			amount:                sdk.NewInt64Coin("uspay", 100_000_000),
			streamType:            types.TypeContinuous,
			duration:              time.Second * 100,
			periods:               nil,
			cancellable:           false,
			paymentFee:            sdk.NewInt64Coin("uspy", 1_000_000),
			valid:                 false,
			expectedMessageEvents: 0,
		},
		{
			sender:                suite.TestAccs[0].String(),
			recipient:             suite.TestAccs[1].String(),
			amount:                sdk.NewInt64Coin("uspay", 1_000_000),
			streamType:            types.TypeContinuous,
			duration:              time.Second * 100,
			periods:               nil,
			cancellable:           false,
			paymentFee:            sdk.Coin{Denom: "", Amount: sdk.NewInt(-1)},
			valid:                 false,
			expectedMessageEvents: 0,
		},
		{
			sender:                suite.TestAccs[0].String(),
			recipient:             suite.TestAccs[1].String(),
			amount:                sdk.Coin{Denom: "", Amount: sdk.NewInt(-1)},
			streamType:            types.TypeContinuous,
			duration:              time.Second * 100,
			periods:               nil,
			cancellable:           false,
			paymentFee:            sdk.NewInt64Coin("uspay", 1_000_000),
			valid:                 false,
			expectedMessageEvents: 0,
		},
		{
			sender:                suite.TestAccs[0].String(),
			recipient:             suite.TestAccs[1].String(),
			amount:                sdk.NewInt64Coin("uspay", 100_000_000),
			streamType:            types.TypeDelayed,
			duration:              time.Second * 100,
			periods:               nil,
			cancellable:           false,
			paymentFee:            sdk.NewInt64Coin("uspay", 1_000_000),
			valid:                 true,
			expectedMessageEvents: 1,
		},
		{
			sender:     suite.TestAccs[0].String(),
			recipient:  suite.TestAccs[1].String(),
			amount:     sdk.NewInt64Coin("uspay", 10_000_000),
			streamType: types.TypePeriodic,
			duration:   time.Second * 100,
			periods: []*types.Period{
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
			},
			cancellable:           true,
			paymentFee:            sdk.NewInt64Coin("uspay", 100_000),
			valid:                 true,
			expectedMessageEvents: 1,
		},
		{
			sender:     suite.TestAccs[0].String(),
			recipient:  suite.TestAccs[1].String(),
			amount:     sdk.NewInt64Coin("uspay", 100_000_000),
			streamType: types.TypePeriodic,
			duration:   time.Second * 100,
			periods: []*types.Period{
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
				{
					Amount:   1_000_000,
					Duration: 10,
				},
			},
			cancellable:           true,
			paymentFee:            sdk.NewInt64Coin("uspay", 1_000_000),
			valid:                 false,
			expectedMessageEvents: 0,
		},
	} {
		suite.Run("create stream payment", func() {
			ctx := suite.Ctx.WithEventManager(sdk.NewEventManager())
			suite.Require().Equal(0, len(ctx.EventManager().Events()))
			// Test stream send message
			msg := types.NewMsgStreamSend(
				tc.sender,
				tc.recipient,
				tc.amount,
				tc.streamType,
				tc.duration,
				tc.periods,
				tc.cancellable,
				tc.paymentFee,
			)
			err := msg.ValidateBasic()
			if err == nil {
				_, err = suite.msgServer.StreamSend(
					sdk.WrapSDKContext(ctx),
					msg,
				)
			}
			if tc.valid {
				suite.Require().NoError(err)
			}
			// Ensure current number and type of event is emitted
			suite.AssertEventEmitted(ctx, types.EventTypeCreateStreamPayment, tc.expectedMessageEvents)
		})
	}
}

func (suite *KeeperTestSuite) TestStopStreamMsg() {
	suite.CreateDefaultStreamPayment(true)
	sp2 := suite.CreateStreamPayment(types.TypeDelayed, false)
	sp3 := suite.CreateStreamPayment(types.TypeContinuous, true)
	for _, tc := range []struct {
		streamId              string
		sender                string
		valid                 bool
		expectedMessageEvents int
	}{
		{
			streamId:              suite.defaultStreamPaymentId,
			sender:                suite.TestAccs[0].String(),
			valid:                 true,
			expectedMessageEvents: 1,
		},
		{
			streamId:              sp2,
			sender:                suite.TestAccs[0].String(),
			valid:                 false,
			expectedMessageEvents: 0,
		},
		{
			streamId:              sp3,
			sender:                suite.TestAccs[1].String(),
			valid:                 false,
			expectedMessageEvents: 0,
		},
	} {
		suite.Run("stop stream", func() {
			ctx := suite.Ctx.WithEventManager(sdk.NewEventManager())
			suite.Require().Equal(0, len(ctx.EventManager().Events()))
			// Test stream send message
			_, err := suite.msgServer.StopStream(
				sdk.WrapSDKContext(ctx),
				types.NewMsgStopStream(
					tc.streamId,
					tc.sender,
				),
			)
			if tc.valid {
				suite.Require().NoError(err)
			}
			// Ensure current number and type of event is emitted
			suite.AssertEventEmitted(ctx, types.EventTypeStopStreamPayment, tc.expectedMessageEvents)
			suite.AssertEventEmitted(ctx, types.EventTypeEndStreamPayment, tc.expectedMessageEvents)
		})
	}
}

func (suite *KeeperTestSuite) TestClaimStreamedAmountMsg() {
	suite.CreateDefaultStreamPayment(false)
	for _, tc := range []struct {
		streamId              string
		recipient             string
		valid                 bool
		expectedMessageEvents int
	}{
		{
			streamId:              suite.defaultStreamPaymentId,
			recipient:             suite.TestAccs[1].String(),
			valid:                 true,
			expectedMessageEvents: 1,
		},
		{
			streamId:              suite.defaultStreamPaymentId,
			recipient:             suite.TestAccs[0].String(),
			valid:                 false,
			expectedMessageEvents: 0,
		},
	} {
		suite.Run("claim streamed amount", func() {
			ctx := suite.Ctx.WithEventManager(sdk.NewEventManager())
			suite.Require().Equal(0, len(ctx.EventManager().Events()))
			// Test stream send message
			_, err := suite.msgServer.ClaimStreamedAmount(
				sdk.WrapSDKContext(ctx),
				types.NewMsgClaimStreamedAmount(
					tc.streamId,
					tc.recipient,
				),
			)
			if tc.valid {
				suite.Require().NoError(err)
			}
			// Ensure current number and type of event is emitted
			suite.AssertEventEmitted(ctx, types.EventTypeClaimStreamedAmount, tc.expectedMessageEvents)
		})
	}
}

func (suite *KeeperTestSuite) TestUpdateParams() {
	testCases := []struct {
		name      string
		request   *types.MsgUpdateParams
		expectErr bool
	}{
		{
			name: "set invalid authority",
			request: &types.MsgUpdateParams{
				Authority: "foo",
			},
			expectErr: true,
		},
		{
			name: "set invalid payment fee param",
			request: &types.MsgUpdateParams{
				Authority: suite.App.StreamPayKeeper.GetAuthority(),
				Params: types.Params{
					StreamPaymentFeePercentage: sdk.NewDec(1),
				},
			},
			expectErr: true,
		},
		{
			name: "set invalid payment gee",
			request: &types.MsgUpdateParams{
				Authority: suite.App.StreamPayKeeper.GetAuthority(),
				Params: types.Params{
					StreamPaymentFeePercentage: sdk.NewDecWithPrec(-5, 2),
				},
			},
			expectErr: true,
		},
		{
			name: "set full valid params",
			request: &types.MsgUpdateParams{
				Authority: suite.App.StreamPayKeeper.GetAuthority(),
				Params: types.Params{
					StreamPaymentFeePercentage: types.DefaultStreamPaymentFeePercentage,
				},
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			_, err := suite.msgServer.UpdateParams(suite.Ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
