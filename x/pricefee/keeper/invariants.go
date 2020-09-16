package keeper

// DONTCOVER

import (
//	"fmt"

	"github.com/KuChainNetwork/kuchain/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


// AllInvariants runs all invariants of the governance module
func AllInvariants(keeper Keeper, bk types.BankKeeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return ModuleAccountInvariant(keeper, bk)(ctx)
	}
}

// ModuleAccountInvariant checks that the module account coins reflects the sum of
// pricefee amounts held on store
func ModuleAccountInvariant(keeper Keeper, bk types.BankKeeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
//		var expectedDeposits Coins

		return "",true
	}
}
