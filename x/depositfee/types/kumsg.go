package types

import (
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//	"github.com/tendermint/tendermint/crypto"
)

var (
	RouterKeyName = chainTypes.MustName(RouterKey)
)

type KuMsgOpenFee struct {
	chainTypes.KuMsg
}

func NewKuMsgRegisterSinger(auth sdk.AccAddress, owner AccountID) KuMsgOpenFee {
	return KuMsgOpenFee{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgOpenFee{
				Owner: owner,
			}),
		),
	}
}

func (msg KuMsgOpenFee) ValidateBasic() error {
	msgData := MsgOpenFee{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}