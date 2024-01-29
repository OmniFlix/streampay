package keeper_test

import (
	"testing"

	"github.com/OmniFlix/streampay/v2/app/apptesting"
	"github.com/OmniFlix/streampay/v2/x/streampay/keeper"
	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient            types.QueryClient
	msgServer              types.MsgServer
	defaultStreamPaymentId string
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	fundAccsAmount := sdk.NewCoins(
		sdk.NewCoin("uspay", sdk.NewInt(10_000_000_000)),
	)
	for _, acc := range suite.TestAccs {
		suite.FundAcc(acc, fundAccsAmount)
	}
	_ = suite.App.StreamPayKeeper.SetParams(suite.Ctx, types.DefaultParams())

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.StreamPayKeeper)
}

func (suite *KeeperTestSuite) CreateDefaultStreamPayment(cancellable bool) {
	ctx := sdk.WrapSDKContext(suite.Ctx)
	res, _ := suite.msgServer.StreamSend(ctx, &types.MsgStreamSend{
		Sender:      suite.TestAccs[0].String(),
		Recipient:   suite.TestAccs[1].String(),
		Amount:      sdk.NewInt64Coin("uspay", 100_000_000),
		StreamType:  types.TypeContinuous,
		Duration:    100,
		Periods:     nil,
		Cancellable: cancellable,
		PaymentFee:  sdk.NewInt64Coin("uspay", 1_000_000),
	})
	suite.defaultStreamPaymentId = res.StreamId
}

func (suite *KeeperTestSuite) CreateStreamPayment(streamType types.StreamType, cancellable bool) string {
	ctx := sdk.WrapSDKContext(suite.Ctx)
	res, _ := suite.msgServer.StreamSend(ctx, &types.MsgStreamSend{
		Sender:      suite.TestAccs[0].String(),
		Recipient:   suite.TestAccs[1].String(),
		Amount:      sdk.NewInt64Coin("uspay", 100_000_000),
		StreamType:  streamType,
		Duration:    100,
		Periods:     nil,
		Cancellable: cancellable,
		PaymentFee:  sdk.NewInt64Coin("uspay", 1_000_000),
	})
	return res.StreamId
}

func (suite *KeeperTestSuite) TestParams() {
	testCases := []struct {
		name      string
		input     types.Params
		expectErr bool
	}{
		{
			name: "set invalid fee",
			input: types.Params{
				StreamPaymentFeePercentage: sdk.NewDec(1),
			},
			expectErr: true,
		},
		{
			name: "set invalid fee",
			input: types.Params{
				StreamPaymentFeePercentage: sdk.NewDec(-1),
			},
			expectErr: true,
		},
		{
			name: "set full valid params",
			input: types.Params{
				StreamPaymentFeePercentage: types.DefaultStreamPaymentFeePercentage,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			expected := suite.App.StreamPayKeeper.GetParams(suite.Ctx)
			err := suite.App.StreamPayKeeper.SetParams(suite.Ctx, tc.input)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				expected = tc.input
				suite.Require().NoError(err)
			}

			p := suite.App.StreamPayKeeper.GetParams(suite.Ctx)
			suite.Require().Equal(expected, p)
		})
	}
}
