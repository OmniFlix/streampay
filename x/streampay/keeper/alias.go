package keeper

import (
	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetPaymentStreamModuleAccount returns PaymentStream ModuleAccount
func (k Keeper) GetPaymentStreamModuleAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
}

func (k Keeper) TransferAmountFromModuleAccount(ctx sdk.Context, to sdk.AccAddress, amount sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, amount)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) TransferAmountToModuleAccount(ctx sdk.Context, fromAddress sdk.AccAddress, amount sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, fromAddress, types.ModuleName, amount)
	if err != nil {
		return err
	}
	return nil
}
