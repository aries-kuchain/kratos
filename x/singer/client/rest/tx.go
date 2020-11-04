package rest

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/chain/client/txutil"
	rest "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"net/http"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	ctx := txutil.NewKuCLICtx(cliCtx)

	r.HandleFunc(
		"/singer/registersinger",
		postRegisterSingerHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/singer/payaccess",
		postPayAccessHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/singer/paymortgage",
		postPayMortgageHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/singer/activesinger",
		postActiveSingerHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/singer/setaddress",
		postSetAddressHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/singer/activedeposit",
		postActiveDepositHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/singer/submitspv",
		postSubmitSpvHandlerFn(ctx),
	).Methods("POST")
}

type (
	SingerRequest struct {
		BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
	}
	SingerMortgageRequest struct {
		BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
		Amount        string       `json:"amount" yaml:"amount"`
	}
	ActiveSingerRequest struct {
		BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
		SystemAccount string       `json:"system_account" yaml:"system_account"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
	}
	SetAddressRequest struct {
		BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
		DepositID     string       `json:"deposit_id" yaml:"deposit_id"`
		Address       string       `json:"address" yaml:"address"`
	}
	SingerDepositRequest struct {
		BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
		DepositID     string       `json:"deposit_id" yaml:"deposit_id"`
	}
	SubmitSpvRequest struct {
		BaseReq            rest.BaseReq `json:"base_req" yaml:"base_req"`
		DepositID          string       `json:"deposit_id" yaml:"deposit_id"`
		SpvSubmiter        string       `json:"spv_submiter" yaml:"spv_submiter"`
		Version            string       `json:"version" yaml:"version"`
		TxInputVector      string       `json:"tx_input_vector" yaml:"tx_input_vector"`
		TxOutputVector     string       `json:"tx_output_vector" yaml:"tx_output_vector"`
		TxLockTime         string       `json:"tx_lock_time" yaml:"tx_lock_time"`
		FundingOutputIndex int          `json:"funding_output_index" yaml:"funding_output_index"`
		MerkleProof        string       `json:"merkle_proof" yaml:"merkle_proof"`
		TxIndexInBlock     int          `json:"tx_index_in_block" yaml:"tx_index_in_block"`
		BitcoinHeaders     string       `json:"bit_coin_headers" yaml:"bit_coin_headers"`
	}
)

func postRegisterSingerHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SingerRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		singerAccount, err := rest.NewAccountIDFromStr(req.SingerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		singerAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", singerAccount, err))
			return
		}

		msg := types.NewKuMsgRegisterSinger(singerAccAddress, singerAccount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postPayAccessHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SingerMortgageRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		singerAccount, err := rest.NewAccountIDFromStr(req.SingerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		singerAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", singerAccount, err))
			return
		}

		amount, err := rest.ParseCoin(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("amount parse error, %v", err))
			return
		}

		msg := types.NewKuMsgPayAccess(singerAccAddress, singerAccount, amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postPayMortgageHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SingerMortgageRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		singerAccount, err := rest.NewAccountIDFromStr(req.SingerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		singerAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", singerAccount, err))
			return
		}

		amount, err := rest.ParseCoin(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("amount parse error, %v", err))
			return
		}

		msg := types.NewKuMsgBTCMortgage(singerAccAddress, singerAccount, amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postActiveSingerHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ActiveSingerRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		singerAccount, err := rest.NewAccountIDFromStr(req.SingerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}
		systemAccount, err := rest.NewAccountIDFromStr(req.SystemAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		systemAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", systemAccount, err))
			return
		}

		msg := types.NewKuMsgActiveSinger(systemAccAddress, systemAccount, singerAccount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postSetAddressHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SetAddressRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		singerAccount, err := rest.NewAccountIDFromStr(req.SingerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		singerAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", singerAccount, err))
			return
		}

		msg := types.NewKuMsgMsgSetBtcAddress(singerAccAddress, singerAccount, req.DepositID, req.Address)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postActiveDepositHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SingerDepositRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		singerAccount, err := rest.NewAccountIDFromStr(req.SingerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		singerAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", singerAccount, err))
			return
		}

		msg := types.NewKuMsgActiveDeposit(singerAccAddress, singerAccount, req.DepositID)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postSubmitSpvHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SubmitSpvRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		spvSubmiter, err := rest.NewAccountIDFromStr(req.SpvSubmiter)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("spvSubmiter accountID error, %v", err))
			return
		}

		submiterAccAddress, err := txutil.QueryAccountAuth(cliCtx, spvSubmiter)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", spvSubmiter, err))
			return
		}
		//(depositID string, spvSubminter AccountID, version, txInputVector, txOutputVector, txLockTime, merkleProof, bitcoinHeaders string, fundingOupputIndex, txIndexInBlock int) SpvInfo {
		spvInfo := types.NewSpvInfo(req.DepositID, spvSubmiter, req.Version, req.TxInputVector, req.TxOutputVector, req.TxLockTime, req.MerkleProof,
			req.BitcoinHeaders, req.FundingOutputIndex, req.TxIndexInBlock)

		msg := types.NewKuMsgSubmitSpv(submiterAccAddress, spvInfo)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
