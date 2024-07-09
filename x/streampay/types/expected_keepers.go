package types

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper Methods imported from account should be defined here
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, name string) sdk.ModuleAccountI
	GetModuleAddress(module string) sdk.AccAddress
	SetModuleAccount(context.Context, sdk.ModuleAccountI)
}

// BankKeeper Methods imported from bank should be defined here
type BankKeeper interface {
	BlockedAddr(recipient sdk.AccAddress) bool
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx context.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, fromModule string, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, fromAddr sdk.AccAddress, toModule string, amt sdk.Coins) error
}

type DistributionKeeper interface {
	FundCommunityPool(ctx context.Context, amount sdk.Coins, sender sdk.AccAddress) error
}
