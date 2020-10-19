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
}

// HTTP request handler to query a delegator delegations
func singerInfoHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryDelegator(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySingerInfo))
}
