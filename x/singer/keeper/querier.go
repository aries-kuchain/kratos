package keeper

import (
	//	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type QueryResResolve struct {
	Value string `json:"value"`
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err error) {
	// name := path[0]

	// value := keeper.ResolveName(ctx, name)

	// if value == "" {
	// 	return []byte{}, nil //sdk.ErrUnknownRequest("could not resolve name")
	// }

	// bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResResolve{value})
	// if err2 != nil {
	// 	panic("could not marshal result to JSON")
	// }

	return nil, nil
}
