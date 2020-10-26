package types

const (
	EventTypeCreateDeposit    = "create_deposit"
	EventTypeSunmitSpv    = "submit_spv"
	EventTypeTransferDeposit   = "transfer_deposit"
	EventTypeDepositToCoin   = "deposit_to_coin"
	EventTypeClaimDeposit   = "claim_deposit"
	EventTypeFinishDeposit   = "finish_deposit"
	EventTypeTimeOutDeposit   = "timeout_deposit"
	EventTypeActiveDeposit    = "active_deposit"
	EventTypeAberrantDeposit    = "active_deposit"
	EventTypeWrongSingerSpv   = "wrong_singer_deposit"
	EventTypeClaimAberrantDeposit   = "claim_aberrant_deposit"
	EventTypeClaimMortgageDeposit   = "claim_mortgage_deposit"
	EventTypeJudgeSpv  = "judge_spv"


	AttributeKeyDepositID         = "deposit_id"
	AttributeKeyOwner         = "owner"
	AttributeKeySinger         = "singer"
	AttributeKeyAsset         = "asset"
	AttributeKeyMinStake         = "min_stake"
	AttributeKeyFee         = "fee"
	AttributeKeyRightSpv         = "right_spv"
	AttributeKeyFeeToSinger         = "fee_to_singer"
)