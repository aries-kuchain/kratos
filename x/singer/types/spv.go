package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
)

type SpvInfo struct {
	DepositID          string    `json:"deposit_id" yaml:"deposit_id"`
	SpvSubmiter        AccountID `json:"spv_submiter" yaml:"spv_submiter"`
	Version            string    `json:"version" yaml:"version"`
	TxInputVector      string    `json:"tx_input_vector" yaml:"tx_input_vector"`
	TxOutputVector     string    `json:"tx_output_vector" yaml:"tx_output_vector"`
	TxLockTime         string    `json:"tx_lock_time" yaml:"tx_lock_time"`
	FundingOutputIndex int       `json:"funding_output_index" yaml:"funding_output_index"`
	MerkleProof        string    `json:"merkle_proof" yaml:"merkle_proof"`
	TxIndexInBlock     int       `json:"tx_index_in_blocl" yaml:"tx_index_in_blocl"`
	BitcoinHeaders     string    `json:"bit_coin_headers" yaml:"bit_coin_headers"`
}

func NewSpvInfo(depositID string, spvSubminter AccountID, version, txInputVector, txOutputVector, txLockTime, merkleProof, bitcoinHeaders string, fundingOupputIndex, txIndexInBlock int) SpvInfo {
	return SpvInfo{
		DepositID:          depositID,
		SpvSubmiter:        spvSubminter,
		Version:            version,
		TxInputVector:      txInputVector,
		TxOutputVector:     txOutputVector,
		TxLockTime:         txLockTime,
		FundingOutputIndex: fundingOupputIndex,
		MerkleProof:        merkleProof,
		TxIndexInBlock:     txIndexInBlock,
		BitcoinHeaders:     bitcoinHeaders,
	}
}

// return the redelegation
func MustMarshalSpvInfo(cdc *codec.Codec, SpvInfo SpvInfo) []byte {
	return cdc.MustMarshalBinaryBare(&SpvInfo)
}

// unmarshal a redelegation from a store value
func MustUnmarshalSpvInfo(cdc *codec.Codec, value []byte) SpvInfo {
	SpvInfo, err := UnmarshalSpvInfo(cdc, value)
	if err != nil {
		panic(err)
	}
	return SpvInfo
}

// unmarshal a redelegation from a store value
func UnmarshalSpvInfo(cdc *codec.Codec, value []byte) (v SpvInfo, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

func (v SpvInfo) String() string {
	return fmt.Sprintf(`DepositID:%s\n
		SpvSubmiter:%s\n
		Version:%s\n
		TxInputVector:%s\n
		TxOutputVector:%s\n
		TxLockTime:%s\n
		FundingOutputIndex:%d\n
		MerkleProof:%s\n
		TxIndexInBlock:%s\n
		BitcoinHeaders:%s\n
		`, v.DepositID,
		v.SpvSubmiter.String(),
		v.Version,
		v.TxInputVector,
		v.TxOutputVector,
		v.TxLockTime,
		v.FundingOutputIndex,
		v.MerkleProof,
		v.TxIndexInBlock,
		v.BitcoinHeaders,
	)
}
