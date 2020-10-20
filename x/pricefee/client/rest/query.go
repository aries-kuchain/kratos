package rest

import (
	"fmt"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/pricefee/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/depositfee/{ownerAccount}/fee",
		depositFeeHandlerFn(cliCtx),
	).Methods("GET")
}

func depositFeeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return queryFeeInfo(cliCtx, fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFeeInfo))
}

func queryFeeInfo(cliCtx context.CLIContext, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32account := vars["ownerAccount"]
		ownerAccount, _ := chainTypes.NewAccountIDFromStr(bech32account)

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQuerySingerInfoParams(ownerAccount)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(endpoint, bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
