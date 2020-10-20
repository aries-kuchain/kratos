package rest

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
	"net/http"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/singer/{singerAccount}/singerinfo",
		singerInfoHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/singer/alldeposit",
		allDepositHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/singer/{depositID}/depositinfo",
		depositInfoHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/singer/{singerAccount}/singerdeposit",
		singerDepositHandlerFn(cliCtx),
	).Methods("GET")
}

func singerInfoHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return querySingerInfo(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySingerInfo))
}

func allDepositHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryAllDeposit(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllDeposit))
}

func depositInfoHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryDepositInfo(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDepositInfo))
}

func singerDepositHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return querySingerDeposit(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySingerDeposit))
}