package types

import ()

const (
	EventTypeRegisterSinger  = "register_singer"
	EventTypePayAccess       = "pay_access"
	EventTypePayMortgage     = "pay_mortgage"
	EventTypeActiveSinger    = "active_singer"
	EventTypeClaimMortgage   = "claim_mortgage"
	EventTypeClaimAccess     = "claim_mortgage"
	EventTypeLogoutSinger    = "logout_singer"
	EventTypeSetAddress      = "set_address"
	EventTypeActiveDeposit   = "active_deposit"
	EventTypeSubmitSpv       = "submit_spv"
	EventTypeFinishDeposit   = "finish_deposit"
	EventTypeWrongDepositSpv = "wrong_deposit_spv"

	AttributeKeySingerAccount = "singer_account"
	AttributeKeyAmount        = "amount"
	AttributeKeyDepositID     = "deposit_id"
)
