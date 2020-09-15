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

type KuMsgRegisterSinger struct {
	chainTypes.KuMsg
}

func NewKuMsgRegisterSinger(auth sdk.AccAddress, singerAccount AccountID) KuMsgRegisterSinger {
	return KuMsgRegisterSinger{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgRegisterSinger{
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgRegisterSinger) ValidateBasic() error {
	msgData := MsgRegisterSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgPayAccess struct {
	chainTypes.KuMsg
}

func NewKuMsgPayAccess(auth sdk.AccAddress, singerAccount AccountID, amount Coin) KuMsgPayAccess {
	return KuMsgPayAccess{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(singerAccount, ModuleAccountID, chainTypes.Coins{amount}),
			msg.WithData(Cdc(), &MsgPayAccess{
				SingerAccount: singerAccount,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgPayAccess) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}

	msgData := MsgPayAccess{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Amount)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgActiveSinger struct {
	chainTypes.KuMsg
}

func NewKuMsgActiveSinger(auth sdk.AccAddress, systemAccount, singerAccount AccountID) KuMsgActiveSinger {
	return KuMsgActiveSinger{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgActiveSinger{
				SystemAccount: systemAccount,
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgActiveSinger) ValidateBasic() error {
	msgData := MsgActiveSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}
