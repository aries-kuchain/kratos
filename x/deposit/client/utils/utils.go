package utils

import (
	"io/ioutil"
	// chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	// rest "github.com/KuChainNetwork/kuchain/chain/types"
	// "github.com/KuChainNetwork/kuchain/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"
)


type (
	ParamsSpvJSON struct {
		Version string `json:"version" yaml:"version"`
		TxInputVector string `json:"tx_input_vector" yaml:"tx_input_vector"`
		TxOutputVector string `json:"tx_output_vector" yaml:"tx_output_vector"`
		TxLockTime string `json:"tx_lock_time" yaml:"tx_lock_time"`
		FundingOutputIndex int `json:"funding_output_index,string" yaml:"funding_output_index"`
		MerkleProof string `json:"merkle_proof" yaml:"merkle_proof"`
		TxIndexInBlock int `json:"tx_index_in_block,string" yaml:"tx_index_in_block"`
		BitcoinHeaders string `json:"bit_coin_headers" yaml:"bit_coin_headers"`
	}
)

func ParseParamsSpvJSON(cdc *codec.Codec, proposalFile string) (ParamsSpvJSON, error) {
	proposal := ParamsSpvJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

