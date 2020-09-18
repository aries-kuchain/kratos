package deposit

import (
	"encoding/json"
	"github.com/KuChainNetwork/kuchain/x/deposit/keeper"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all auth state that must be provided at genesis
type GenesisState struct {
	Nullstring string `json:"null" yaml:"null"`
}

// NewGenesisState - Create a new genesis state
func NewGenesisState(tempstring string) GenesisState {
	return GenesisState{
		Nullstring: tempstring,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState("nothing")
}

// ValidateGenesis performs basic validation of bank genesis data returning an
// error for any failed validation criteria.
func (g GenesisState) ValidateGenesis(bz json.RawMessage) error {

	return nil
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

// ExportGenesis - output genesis parameters
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) GenesisState {

	return GenesisState{
		Nullstring: "",
	}
}

// InitGenesis - store genesis parameters
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data GenesisState, supplyKeeper types.SupplyKeeper,
) {
	if err := supplyKeeper.InitModuleAccount(ctx, types.ModuleName); err != nil {
		panic(err)
	}
}

const (
	keyCommunityTax        = "communitytax"
	keyBaseProposerReward  = "baseproposerreward"
	keyBonusProposerReward = "bonusproposerreward"
)
