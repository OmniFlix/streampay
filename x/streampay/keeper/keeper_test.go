package keeper_test

import (
	"testing"

	"github.com/OmniFlix/streampay/app/apptesting"
	"github.com/OmniFlix/streampay/x/streampay/keeper"
	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
	msgServer   types.MsgServer
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	fundAccsAmount := sdk.NewCoins(sdk.NewCoin(types.DefaultParams().StreamPaymentFee.Denom, types.DefaultParams().StreamPaymentFee.Amount.MulRaw(100)))
	for _, acc := range suite.TestAccs {
		suite.FundAcc(acc, fundAccsAmount)
	}

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.StreamPayKeeper)
}
