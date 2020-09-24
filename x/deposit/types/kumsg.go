package types

import (
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
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
//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgSubmitSpv struct {
	chainTypes.KuMsg
}

func NewKuMsgSubmitSpv(auth sdk.AccAddress,spvInfo singerTypes.SpvInfo ) KuMsgSubmitSpv {
	return KuMsgSubmitSpv{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgSubmitSpv{
				SpvInfo:spvInfo,
			}),
		),
	}
}

func (msg KuMsgSubmitSpv) ValidateBasic() error {
	msgData := MsgSubmitSpv{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}
//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgTransferDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgTransferDeposit(auth sdk.AccAddress,depositID string,from,to AccountID,memo string ) KuMsgTransferDeposit {
	return KuMsgTransferDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgTransferDeposit{
					DepositID:depositID,
					From:from,
					To:to,
					Memo:memo,
			}),
		),
	}
}

func (msg KuMsgTransferDeposit) ValidateBasic() error {
	msgData := MsgTransferDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}
//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgDepositToCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgDepositToCoin(auth sdk.AccAddress,depositID string,owner AccountID ) KuMsgDepositToCoin {
	return KuMsgDepositToCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgDepositToCoin{
					DepositID:depositID,
					Owner:owner,
			}),
		),
	}
}

func (msg KuMsgDepositToCoin) ValidateBasic() error {
	msgData := MsgDepositToCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}
//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgDepositClaimCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgDepositClaimCoin(auth sdk.AccAddress,depositID string,owner AccountID,asset Coin,claimAddress []byte ) KuMsgDepositClaimCoin {
	return KuMsgDepositClaimCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(owner, ModuleAccountID, chainTypes.Coins{asset}),
			msg.WithData(Cdc(), &MsgDepositClaimCoin{
					DepositID:depositID,
					Owner:owner,
					Asset:asset,
					ClaimAddress:claimAddress,
			}),
		),
	}
}

func (msg KuMsgDepositClaimCoin) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}

	msgData := MsgDepositClaimCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Asset)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}
//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgFinishDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgFinishDeposit(auth sdk.AccAddress,depositID string,owner AccountID ) KuMsgFinishDeposit {
	return KuMsgFinishDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgFinishDeposit{
					DepositID:depositID,
					Owner:owner,
			}),
		),
	}
}

func (msg KuMsgFinishDeposit) ValidateBasic() error {

	msgData := MsgFinishDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}