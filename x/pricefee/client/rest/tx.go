package rest

import(
	"fmt"
	"net/http"

	"github.com/KuChainNetwork/kuchain/chain/client/txutil"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	rest "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/pricefee/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	ctx := txutil.NewKuCLICtx(cliCtx)

	r.HandleFunc(
		"/pricefee/openfee",
		postOpenFeeHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/pricefee/prestorefee",
		postPreStoreFeeHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/pricefee/claimfee",
		postClaimFeeHandlerFn(ctx),
	).Methods("POST")
	r.HandleFunc(
		"/pricefee/setprice",
		postSetPriceHandlerFn(ctx),
	).Methods("POST")
}

type (
	OpenFeeRequest struct {
		BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
		Owner string       `json:"owner" yaml:"owner"`
	}
	PrestoreFeeRequest struct {
		BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
		Owner string       `json:"owner" yaml:"owner"`
		Amount string       `json:"amount" yaml:"amount"`
	}
	SetPriceRequest struct {
		BaseReq      rest.BaseReq `json:"base_req" yaml:"base_req"`
 		SystemAccount string       `json:"system_account" yaml:"system_account"`
		Base string       `json:"base" yaml:"base"`
		Quote string       `json:"quote" yaml:"quote"`
		Remark string       `json:"remark" yaml:"remark"`
	}
)


func postOpenFeeHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req OpenFeeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		owner, err := chainTypes.NewAccountIDFromStr(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("owner accountID error, %v", err))
			return
		}

		ownerAccAddress, err := txutil.QueryAccountAuth(cliCtx, owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", owner, err))
			return
		}

		msg := types.NewKuMsgOpenFee(ownerAccAddress, owner)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postPreStoreFeeHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PrestoreFeeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		owner, err := chainTypes.NewAccountIDFromStr(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("owner accountID error, %v", err))
			return
		}

		amount, err := chainTypes.ParseCoin(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("amount parse error, %v", err))
			return
		}

		ownerAccAddress, err := txutil.QueryAccountAuth(cliCtx, owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", owner, err))
			return
		}

		msg := types.NewKuMsgPrestoreFee(ownerAccAddress, owner,amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postClaimFeeHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PrestoreFeeRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		owner, err := chainTypes.NewAccountIDFromStr(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("owner accountID error, %v", err))
			return
		}

		amount, err := chainTypes.ParseCoin(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("amount parse error, %v", err))
			return
		}

		ownerAccAddress, err := txutil.QueryAccountAuth(cliCtx, owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", owner, err))
			return
		}

		msg := types.NewKuMsgClaimFee(ownerAccAddress, owner,amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}


func postSetPriceHandlerFn(cliCtx txutil.KuCLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SetPriceRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()

		systemAccount, err := chainTypes.NewAccountIDFromStr(req.SystemAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("systemAccount accountID error, %v", err))
			return
		}

		base, err := chainTypes.ParseCoin(req.Base)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("base parse error, %v", err))
			return
		}

		quote, err := chainTypes.ParseCoin(req.Quote)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("quote parse error, %v", err))
			return
		}

		systemAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query account %s auth error, %v", systemAccount, err))
			return
		}

		msg := types.NewKuMsgSetPrice(systemAccAddress, systemAccount,base,quote,req.Remark)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		txutil.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}