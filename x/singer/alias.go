package singer

// nolint

import (
	"github.com/KuChainNetwork/kuchain/x/singer/keeper"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
)

const (
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
)

var (
	NewKeeper     = keeper.NewKeeper
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper = keeper.Keeper
)

var (
	Cdc = types.Cdc
)