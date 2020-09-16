package simulation

import (
	//	"bytes"
	"fmt"

	tmkv "github.com/tendermint/tendermint/libs/kv"

	// "github.com/KuChainNetwork/kuchain/x/staking/types"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding staking type
func DecodeStore(cdc *codec.Codec, kvA, kvB tmkv.Pair) string {
	switch {

	default:
		panic(fmt.Sprintf("invalid staking key prefix %X", kvA.Key[:1]))
	}
}
