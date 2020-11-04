package rest

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/chain/client/txutil"
	rest "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"net/http"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	ctx := txutil.NewKuCLICtx(cliCtx)

	r.HandleFunc(
		"/deposit/createlegalcoin",
		postCreateLegalCoinHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/deposit/createdeposit",
		postCreateDepositHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/deposit/submitspv",
		postSubmitSpvHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/deposit/transferdeposit",
		postTransferDepositHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/deposit/deposittocoin",
		postDepositToCoinHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/deposit/depositclaimcoin",
		postDepositClaimCoinHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/deposit/finishdeposit",
		postFinishDepositHandlerFn(ctx),
	).Methods("POST")
}

type (
	CreateLegalCoinRequest struct {
		BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
		SystemAccount string       `json:"system_account" yaml:"system_account"`
		Amount        string       `json:"amount" yaml:"amount"`
		Symbol        string       `json:"symbol" yaml:"symbol"`
	}
	CreateDepositRequest struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
		Owner   string       `json:"owner" yaml:"owner"`
		Amount  string       `json:"amount" yaml:"amount"`
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
	TransferDepositRequest struct {
		BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
		DepositID string       `json:"deposit_id" yaml:"deposit_id"`
		From      string       `json:"from" yaml:"from"`
		To        string       `json:"to" yaml:"to"`
		Memo      string       `json:"memo" yaml:"memo"`
	}
	DepositOwnerRequest struct {
		BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
		DepositID string       `json:"deposit_id" yaml:"deposit_id"`
		Owner     string       `json:"owner" yaml:"owner"`
	}
	DepositClaimRequest struct {
		BaseReq   rest.BaseReq `json:"base_req" yaml:"base_req"`
		DepositID string       `json:"deposit_id" yaml:"deposit_id"`
		Owner     string       `json:"owner" yaml:"owner"`
		Amount    string       `json:"amount" yaml:"amount"`
		Address   string       `json:"address" yaml:"address"`
	}
)

func postCreateLegalCoinHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateLegalCoinRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		systemAccount, err := rest.NewAccountIDFromStr(req.SystemAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		amount, err := rest.ParseCoin(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("amount parse error, %v", err))
			return
		}

		symbol, err := rest.NewName(req.Symbol)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("symbol parse error, %v", err))
			return
		}

		systemAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", systemAccount, err))
			return
		}

		msg := types.NewKuMsgCreateLegalCoin(systemAccAddress, systemAccount, amount, symbol)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postCreateDepositHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateDepositRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		ownerAccount, err := rest.NewAccountIDFromStr(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("singerAccount accountID error, %v", err))
			return
		}

		amount, err := rest.ParseCoin(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("amount parse error, %v", err))
			return
		}

		ownerAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", ownerAccount, err))
			return
		}

		msg := types.NewKuMsgCreateDeposit(ownerAccAddress, ownerAccount, amount)
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
		spvInfo := singerTypes.NewSpvInfo(req.DepositID, spvSubmiter, req.Version, req.TxInputVector, req.TxOutputVector, req.TxLockTime, req.MerkleProof,
			req.BitcoinHeaders, req.FundingOutputIndex, req.TxIndexInBlock)

		msg := types.NewKuMsgSubmitSpv(submiterAccAddress, spvInfo)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postTransferDepositHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req TransferDepositRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		fromAccount, err := rest.NewAccountIDFromStr(req.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("from accountID error, %v", err))
			return
		}

		toAccount, err := rest.NewAccountIDFromStr(req.To)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("to accountID error, %v", err))
			return
		}

		fromAccAddress, err := txutil.QueryAccountAuth(cliCtx, fromAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", fromAccount, err))
			return
		}

		msg := types.NewKuMsgTransferDeposit(fromAccAddress, req.DepositID, fromAccount, toAccount, req.Memo)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postDepositToCoinHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DepositOwnerRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		ownerAccount, err := rest.NewAccountIDFromStr(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("from accountID error, %v", err))
			return
		}

		ownerAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", ownerAccount, err))
			return
		}

		msg := types.NewKuMsgDepositToCoin(ownerAccAddress, req.DepositID, ownerAccount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postDepositClaimCoinHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DepositClaimRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		ownerAccount, err := rest.NewAccountIDFromStr(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("ownerAccount accountID error, %v", err))
			return
		}

		amount, err := rest.ParseCoin(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("amount parse error, %v", err))
			return
		}

		ownerAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", ownerAccount, err))
			return
		}

		msg := types.NewKuMsgDepositClaimCoin(ownerAccAddress, req.DepositID, ownerAccount, amount, req.Address)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postFinishDepositHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DepositOwnerRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		ownerAccount, err := rest.NewAccountIDFromStr(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("from accountID error, %v", err))
			return
		}

		ownerAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", ownerAccount, err))
			return
		}

		msg := types.NewKuMsgFinishDeposit(ownerAccAddress, req.DepositID, ownerAccount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
