package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
//	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/xuyp1991/cosaccount/x/easystore/types"
	"github.com/KuChainNetwork/kuchain/x/depositfee/types"
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"
)
// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {

	storeKey  sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc  *codec.Codec // The wire codec for binary encoding/decoding.
}

// Sets the entire Whois metadata struct for a name
func (k Keeper) SetStoredata(ctx sdk.Context, name string, storedata types.Storedata) {
	if storedata.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(&storedata))
}

// Gets the entire Whois metadata struct for a name
func (k Keeper) GetStoreData(ctx sdk.Context, name string) types.Storedata {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(name)) {
		return types.NewStoredata()
	}
	bz := store.Get([]byte(name))
	var storedata types.Storedata
	k.cdc.MustUnmarshalBinaryBare(bz, &storedata)
	return storedata
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetStoreData(ctx, name).Value
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetStoreData(ctx, name).Owner.Empty()
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetStoreData(ctx, name).Owner
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	storedata := k.GetStoreData(ctx, name)
	storedata.Owner = owner
	k.SetStoredata(ctx, name, storedata)
}

// SetOwner - sets the current owner of a name
func (k Keeper) Setvalue(ctx sdk.Context, name string, value string,owner sdk.AccAddress) {
	storedata := k.GetStoreData(ctx, name)
	storedata.Owner = owner
	storedata.Value = value
	k.SetStoredata(ctx, name, storedata)
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper( storeKey sdk.StoreKey, cdc  *codec.Codec) Keeper {
	return Keeper{
	//	coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
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

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryValue:
			return nil,nil//queryResolve(ctx, path[1:], req, keeper)

		default:
			return nil, nil//sdk.ErrUnknownRequest("unknown bank query endpoint")
		}
	}
}
