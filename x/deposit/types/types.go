package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName      = "deposit"
	StoreKey        = ModuleName
	RouterKey       = ModuleName
	keyCommunityTax = "just for test"
	QueryValue      = "deposit"
)

type QueryResResolve struct {
	Value string `json:"value"`
}

// Storedata
type Storedata struct {
	Value string         `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	Owner sdk.AccAddress `protobuf:"bytes,1,opt,name=value,proto3" json:"owner"`
}

// Returns a new Whois with the minprice as the price
func NewStoredata() Storedata {
	return Storedata{}
}

func (this QueryResResolve) String() string {
	return fmt.Sprintf("value:%s", this.Value)
}
