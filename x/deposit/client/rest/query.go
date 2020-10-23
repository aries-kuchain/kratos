package rest

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
	"net/http"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/deposit/alldeposit",
		allDepositHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/deposit/aberrantdeposit",
		aberrantDepositHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/deposit/cashreadydeposit",
		cashReadyDepositHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/deposit/{depositID}/depositinfo",
		depositInfoHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/deposit/{ownerAccount}/deposit",
		userDepositHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/deposit/{depositID}/depositspv",
		depositSpvHandlerFn(cliCtx),
	).Methods("GET")
}

func allDepositHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryAllDeposit(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllDeposit))
}

func depositInfoHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryDepositInfo(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDepositInfo))
}

func userDepositHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryUserDeposit(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryUserDeposit))
}

func cashReadyDepositHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryAllDeposit(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCashReadyDeposit))
}

func aberrantDepositHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryAllDeposit(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAberrantDeposit))
}

func depositSpvHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryDepositInfo(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDepositSpv))
}