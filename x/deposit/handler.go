package deposit

import (
	//	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdk "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/keeper"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(k keeper.Keeper) msg.Handler {
	return func(ctx chainTypes.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

// Handle a message to buy name
func handleMsgSetStore(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgSetStore) sdk.Result {
	keeper.Setvalue(ctx, msg.Name, msg.Value, msg.Owner)
	return sdk.Result{}
}
