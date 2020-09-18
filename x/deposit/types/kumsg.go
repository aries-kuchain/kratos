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

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgCreateDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgCreateDeposit(auth sdk.AccAddress, ownerAccountID AccountID, amount Coin) KuMsgCreateDeposit {
	return KuMsgCreateDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgCreateDeposit{
				OwnerAccount: ownerAccountID,
				Amount:       amount,
			}),
		),
	}
}

func (msg KuMsgCreateDeposit) ValidateBasic() error {
	msgData := MsgCreateDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgCreateLegalCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgCreateLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin, symbol chainTypes.Name) KuMsgCreateLegalCoin {
	return KuMsgCreateLegalCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgCreateLegalCoin{
				SystemAccount: systemAccountID,
				MaxSupply:     amount,
				Symbol:        symbol,
			}),
		),
	}
}

func (msg KuMsgCreateLegalCoin) ValidateBasic() error {
	msgData := MsgCreateLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgProhibitLegalCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgProhibitLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin) KuMsgProhibitLegalCoin {
	return KuMsgProhibitLegalCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgProhibitLegalCoin{
				SystemAccount: systemAccountID,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgProhibitLegalCoin) ValidateBasic() error {
	msgData := MsgProhibitLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgPermintLegalCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgPermintLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin) KuMsgPermintLegalCoin {
	return KuMsgPermintLegalCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgProhibitLegalCoin{
				SystemAccount: systemAccountID,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgPermintLegalCoin) ValidateBasic() error {
	msgData := MsgPermintLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}
