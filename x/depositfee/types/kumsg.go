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
//------------------------------------------------------------------------------------------------------------------------------------
type KuMsgPrestoreFee struct {
	chainTypes.KuMsg
}

func NewKuMsgPrestoreFee(auth sdk.AccAddress, owner AccountID,amount Coin) KuMsgPrestoreFee {
	return KuMsgPrestoreFee{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(owner, ModuleAccountID, chainTypes.Coins{amount}),
			msg.WithData(Cdc(), &MsgPrestoreFee{
				Owner: owner,
				Amount:amount,
			}),
		),
	}
}

func (msg KuMsgPrestoreFee) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}
	msgData := MsgPrestoreFee{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Amount)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}
//------------------------------------------------------------------------------------------------------------------------------------
type KuMsgClaimFee struct {
	chainTypes.KuMsg
}

func NewKuMsgClaimFee(auth sdk.AccAddress, owner AccountID, amount Coin) KuMsgClaimFee {
	return KuMsgClaimFee{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgClaimFee{
				Owner: owner,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgClaimFee) ValidateBasic() error {
	msgData := MsgClaimFee{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}