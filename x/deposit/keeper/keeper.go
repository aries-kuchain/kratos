package keeper

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//abci "github.com/tendermint/tendermint/abci/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/external"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc            *codec.Codec // The wire codec for binary encoding/decoding.
	bankKeeper     types.BankKeeper
	accountKeeper  types.AccountKeeper
	supplyKeeper   types.SupplyKeeper
	pricefeeKeeper types.PriceFeeKeeper
	singerKeeper   types.SingerKeeper
	paramstore     external.ParamsSubspace
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, bk types.BankKeeper, ak types.AccountKeeper,
	sk types.SupplyKeeper, pfk types.PriceFeeKeeper, srk types.SingerKeeper, ps external.ParamsSubspace,
) Keeper {
	return Keeper{
		storeKey:       storeKey,
		cdc:            cdc,
		bankKeeper:     bk,
		accountKeeper:  ak,
		supplyKeeper:   sk,
		singerKeeper:   srk,
		pricefeeKeeper: pfk,
		paramstore:     ps.WithKeyTable(types.ParamKeyTable()),
	}
}

// RegisterInvariants registers the bank module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, bk Keeper) {
	ir.RegisterRoute(types.ModuleName, "nonnegative-outstanding",
		NonnegativeBalanceInvariant(bk))
}

// NonnegativeBalanceInvariant checks that all accounts in the application have non-negative balances
func NonnegativeBalanceInvariant(bk Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			msg   string
			count int
		)

		bk.IterateAllBalances(ctx, func(addr sdk.AccAddress, balance sdk.Coin) bool {
			if balance.IsNegative() {
				count++
				msg += fmt.Sprintf("\t%s has a negative balance of %s\n", addr, balance)
			}

			return false
		})

		broken := count != 0

		return sdk.FormatInvariant(
			types.ModuleName, "nonnegative-outstanding",
			fmt.Sprintf("amount of negative balances found %d\n%s", count, msg),
		), broken
	}
}

// IterateAllBalances iterates over all the balances of all accounts and
// denominations that are provided to a callback. If true is returned from the
// callback, iteration is halted.
func (k Keeper) IterateAllBalances(ctx sdk.Context, cb func(sdk.AccAddress, sdk.Coin) bool) {
	// store := ctx.KVStore(k.storeKey)
	// balancesStore := prefix.NewStore(store, types.BalancesPrefix)

	// iterator := balancesStore.Iterator(nil, nil)
	// defer iterator.Close()

	// for ; iterator.Valid(); iterator.Next() {
	// 	address := types.AddressFromBalancesStore(iterator.Key())

	// 	var balance sdk.Coin
	// 	k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &balance)

	// 	if cb(address, balance) {
	// 		break
	// 	}
	// }
}
