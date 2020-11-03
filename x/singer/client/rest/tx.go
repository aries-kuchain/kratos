package rest

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/KuChainNetwork/kuchain/chain/client/txutil"
	rest "github.com/KuChainNetwork/kuchain/chain/types"
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
}

type (
	SingerRequest struct {
		BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
	}
	SingerMortgageRequest struct {
		BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
		Amount string       `json:"amount" yaml:"amount"`
	}
	ActiveSingerRequest struct {
		BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
		SystemAccount string       `json:"system_account" yaml:"system_account"`
		SingerAccount string       `json:"singer_account" yaml:"singer_account"`
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

		msg := types.NewKuMsgPayAccess(singerAccAddress, singerAccount,amount)
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

		msg := types.NewKuMsgBTCMortgage(singerAccAddress, singerAccount,amount)
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

		msg := types.NewKuMsgActiveSinger(systemAccAddress,systemAccount, singerAccount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}