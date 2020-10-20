package deposit_test

import (
	"github.com/KuChainNetwork/kuchain/chain/config"
	"github.com/KuChainNetwork/kuchain/chain/constants"
	"github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/test/simapp"
	assetTypes "github.com/KuChainNetwork/kuchain/x/asset/types"
	depositTypes "github.com/KuChainNetwork/kuchain/x/deposit/types"
	priceFeeTypes "github.com/KuChainNetwork/kuchain/x/pricefee/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	. "github.com/smartystreets/goconvey/convey"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
	"time"
)

func newTestApp(wallet *simapp.Wallet) (addAlice, addJack, addValidator sdk.AccAddress, accAlice, accJack, accValidator types.AccountID, app *simapp.SimApp) {
	addAlice = wallet.NewAccAddress()
	addJack = wallet.NewAccAddress()
	addValidator = wallet.NewAccAddress()

	accAlice = types.MustAccountID("alice@ok")
	accJack = types.MustAccountID("jack@ok")
	accValidator = types.MustAccountID("validator@ok")

	resInt, succ := sdk.NewIntFromString("100000000000000000000000")
	if !succ {
		resInt = sdk.NewInt(10000000000000000)
	}
	otherCoinDenom := types.CoinDenom(types.MustName("foo"), types.MustName("coin"))
	initAsset := types.NewCoin(constants.DefaultBondDenom, resInt)

	asset1 := types.NewCoins(
		types.NewInt64Coin(otherCoinDenom, 100000000000000000),
		initAsset)

	asset2 := types.NewCoins(
		types.NewInt64Coin(otherCoinDenom, 67),
		types.NewInt64Coin(constants.DefaultBondDenom, 10000000))

	genAlice := simapp.NewSimGenesisAccount(accAlice, addAlice).WithAsset(asset1)
	genJack := simapp.NewSimGenesisAccount(accJack, addJack).WithAsset(asset1)
	genValidator := simapp.NewSimGenesisAccount(accValidator, addValidator).WithAsset(asset2)

	//-------------------------------------------------------------------------------------------------------------------------------------
	singera := types.MustAccountID("singera")
	singerb := types.MustAccountID("singerb")
	singerc := types.MustAccountID("singerc")
	singerd := types.MustAccountID("singerd")
	singere := types.MustAccountID("singere")
	singerf := types.MustAccountID("singerf")
	singerg := types.MustAccountID("singerg")
	accSystem := types.MustAccountID("test@sys")

	genSingera := simapp.NewSimGenesisAccount(singera, addAlice).WithAsset(asset1)
	genSingerb := simapp.NewSimGenesisAccount(singerb, addAlice).WithAsset(asset1)
	genSingerc := simapp.NewSimGenesisAccount(singerc, addAlice).WithAsset(asset1)
	genSingerd := simapp.NewSimGenesisAccount(singerd, addAlice).WithAsset(asset1)
	genSingere := simapp.NewSimGenesisAccount(singere, addAlice).WithAsset(asset1)
	genSingerf := simapp.NewSimGenesisAccount(singerf, addAlice).WithAsset(asset1)
	genSingerg := simapp.NewSimGenesisAccount(singerg, addAlice).WithAsset(asset1)
	genSystem := simapp.NewSimGenesisAccount(accSystem, addAlice).WithAsset(asset1)

	//-------------------------------------------------------------------------------------------------------------------------------------

	genAccs := simapp.NewGenesisAccounts(wallet.GetRootAuth(), genAlice, genJack, genValidator,
		genSingera,
		genSingerb,
		genSingerc,
		genSingerd,
		genSingere,
		genSingerf,
		genSingerg,
		genSystem,
	)
	app = simapp.SetupWithGenesisAccounts(genAccs)

	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	accountAlice := app.AccountKeeper().GetAccount(ctxCheck, accAlice)
	accountJack := app.AccountKeeper().GetAccount(ctxCheck, accJack)
	accountValidator := app.AccountKeeper().GetAccount(ctxCheck, accValidator)

	So(accountAlice, ShouldNotBeNil)
	So(genAlice.GetID().Eq(accountAlice.GetID()), ShouldBeTrue)
	So(genAlice.GetAuth().Equals(accountAlice.GetAuth()), ShouldBeTrue)

	So(accountJack, ShouldNotBeNil)
	So(genJack.GetID().Eq(accountJack.GetID()), ShouldBeTrue)
	So(genJack.GetAuth().Equals(accountJack.GetAuth()), ShouldBeTrue)

	So(accountValidator, ShouldNotBeNil)
	So(genValidator.GetID().Eq(accountValidator.GetID()), ShouldBeTrue)
	So(genValidator.GetAuth().Equals(accountValidator.GetAuth()), ShouldBeTrue)

	return addAlice, addJack, addValidator, accAlice, accJack, accValidator, app
}

//  ready to work

func regesterSinger(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)

	//NewKuMsgCreateLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin, symbol chainTypes.Name) KuMsgCreateLegalCoin
	msg := singerTypes.NewKuMsgRegisterSinger(addAlice, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("regesterSinger error log", "err", err)
	return err
}

func payAccesss(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)

	msg := singerTypes.NewKuMsgPayAccess(addAlice, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("payAccesss error log", "err", err)
	return err
}

func activeSinger(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice, accJack types.AccountID, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)

	msg := singerTypes.NewKuMsgActiveSinger(addAlice, accAlice, accJack)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("payAccesss error log", "err", err)
	return err
}

func payMortgage(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgBTCMortgage(auth sdk.AccAddress, singerAccount AccountID, amount Coin) KuMsgBTCMortgage
	msg := singerTypes.NewKuMsgBTCMortgage(addAlice, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("payMortgage error log", "err", err)
	return err
}

func createLeginCoin(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, symbol types.Name, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgCreateLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin, symbol chainTypes.Name) KuMsgCreateLegalCoin
	msg := depositTypes.NewKuMsgCreateLegalCoin(addAlice, accAlice, amount, symbol)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("createLeginCoin error log", "err", err)
	return err
}

func prohibitLegalCoin(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgProhibitLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin) KuMsgProhibitLegalCoin
	msg := depositTypes.NewKuMsgProhibitLegalCoin(addAlice, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("prohibitLegalCoin error log", "err", err)
	return err
}

func permintLegalCoin(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgPermintLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin) KuMsgPermintLegalCoin
	msg := depositTypes.NewKuMsgPermintLegalCoin(addAlice, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("permintLegalCoin error log", "err", err)
	return err
}

func setPrice(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount1, amout2 types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgSetPrice(auth sdk.AccAddress, systemAccount AccountID, base,quote Coin,remark string) KuMsgSetPrice
	msg := priceFeeTypes.NewKuMsgSetPrice(addAlice, accAlice, amount1, amout2, "")

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("setPrice error log", "err", err)
	return err
}

//singer ready price ready  create-legalcoin
func readyForDeposit(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress) error {
	symbol := types.MustName("btc")
	otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
	leginCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(100000000000000000))
	accSystem := types.MustAccountID("test@sys")
	err := createLeginCoin(t, wallet, app, addAlice, accSystem, leginCoin, symbol, true)
	So(err, ShouldBeNil)
	amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 100)
	amountPrice := types.NewCoin(otherCoinDenom, sdk.NewInt(100))
	err = setPrice(t, wallet, app, addAlice, accSystem, amountPrice, amout1, true)
	So(err, ShouldBeNil)
	//------------------------------------------------------------------------------------
	singera := types.MustAccountID("singera")
	singerb := types.MustAccountID("singerb")
	singerc := types.MustAccountID("singerc")
	singerd := types.MustAccountID("singerd")
	singere := types.MustAccountID("singere")
	singerf := types.MustAccountID("singerf")
	singerg := types.MustAccountID("singerg")

	err = regesterSinger(t, wallet, app, addAlice, singera, true)
	So(err, ShouldBeNil)
	err = regesterSinger(t, wallet, app, addAlice, singerb, true)
	So(err, ShouldBeNil)
	err = regesterSinger(t, wallet, app, addAlice, singerc, true)
	So(err, ShouldBeNil)
	err = regesterSinger(t, wallet, app, addAlice, singerd, true)
	So(err, ShouldBeNil)
	err = regesterSinger(t, wallet, app, addAlice, singere, true)
	So(err, ShouldBeNil)
	err = regesterSinger(t, wallet, app, addAlice, singerf, true)
	So(err, ShouldBeNil)
	err = regesterSinger(t, wallet, app, addAlice, singerg, true)
	So(err, ShouldBeNil)

	amout2 := types.NewInt64Coin(constants.DefaultBondDenom, 1000000000000000000)
	err = payAccesss(t, wallet, app, addAlice, singera, amout2, true)
	So(err, ShouldBeNil)
	err = payAccesss(t, wallet, app, addAlice, singerb, amout2, true)
	So(err, ShouldBeNil)
	err = payAccesss(t, wallet, app, addAlice, singerc, amout2, true)
	So(err, ShouldBeNil)
	err = payAccesss(t, wallet, app, addAlice, singerd, amout2, true)
	So(err, ShouldBeNil)
	err = payAccesss(t, wallet, app, addAlice, singere, amout2, true)
	So(err, ShouldBeNil)
	err = payAccesss(t, wallet, app, addAlice, singerf, amout2, true)
	So(err, ShouldBeNil)
	err = payAccesss(t, wallet, app, addAlice, singerg, amout2, true)
	So(err, ShouldBeNil)

	err = activeSinger(t, wallet, app, addAlice, accSystem, singera, true)
	So(err, ShouldBeNil)
	err = activeSinger(t, wallet, app, addAlice, accSystem, singerb, true)
	So(err, ShouldBeNil)
	err = activeSinger(t, wallet, app, addAlice, accSystem, singerc, true)
	So(err, ShouldBeNil)
	err = activeSinger(t, wallet, app, addAlice, accSystem, singerd, true)
	So(err, ShouldBeNil)
	err = activeSinger(t, wallet, app, addAlice, accSystem, singere, true)
	So(err, ShouldBeNil)
	err = activeSinger(t, wallet, app, addAlice, accSystem, singerf, true)
	So(err, ShouldBeNil)
	err = activeSinger(t, wallet, app, addAlice, accSystem, singerg, true)
	So(err, ShouldBeNil)

	amoutMortgage := types.NewInt64Coin(constants.DefaultBondDenom, 1000000)
	err = payMortgage(t, wallet, app, addAlice, singera, amoutMortgage, true)
	So(err, ShouldBeNil)
	err = payMortgage(t, wallet, app, addAlice, singerb, amoutMortgage, true)
	So(err, ShouldBeNil)
	err = payMortgage(t, wallet, app, addAlice, singerc, amoutMortgage, true)
	So(err, ShouldBeNil)
	err = payMortgage(t, wallet, app, addAlice, singerd, amoutMortgage, true)
	So(err, ShouldBeNil)
	err = payMortgage(t, wallet, app, addAlice, singere, amoutMortgage, true)
	So(err, ShouldBeNil)
	err = payMortgage(t, wallet, app, addAlice, singerf, amoutMortgage, true)
	So(err, ShouldBeNil)
	err = payMortgage(t, wallet, app, addAlice, singerg, amoutMortgage, true)
	So(err, ShouldBeNil)
	//------------------------------------------------------------------------------------
	return nil
}

func openFee(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	msg := priceFeeTypes.NewKuMsgOpenFee(addAlice, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("openFee error log", "err", err)
	return err
}

func preStoreFee(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	msg := priceFeeTypes.NewKuMsgPrestoreFee(addAlice, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("preStoreFee error log", "err", err)
	return err
}

func createDeposit(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, passed bool) (err error, depositId string, singers []types.AccountID) {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgCreateDeposit(auth sdk.AccAddress, ownerAccountID AccountID, amount Coin)
	msg := depositTypes.NewKuMsgCreateDeposit(addAlice, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	blockTime := time.Now()
	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: blockTime}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("createDeposit error log", "err", err)
	keeper := app.DepositKeeper()
	allDeposit := keeper.GetAllDepositInfo(ctxCheck)
	if len(allDeposit) > 0 {
		Index := len(allDeposit) - 1
		depositId = allDeposit[Index].DepositID
		return err, depositId, allDeposit[Index].Singers
	}
	return err, "", nil

}

func setDepositAddress(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, btcAddress string, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	msg := singerTypes.NewKuMsgMsgSetBtcAddress(addAlice, accAlice, depositID, btcAddress)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("setDepositAddress error log", "err", err)
	return err
}

func submitDepositSpv(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, spvInfo singerTypes.SpvInfo, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgSubmitSpv(auth sdk.AccAddress,spvInfo singerTypes.SpvInfo ) KuMsgSubmitSpv
	msg := depositTypes.NewKuMsgSubmitSpv(addAlice, spvInfo)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("submitDepositSpv error log", "err", err)
	return err
}

func activeDeposit(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgActiveDeposit(auth sdk.AccAddress, singerAccount AccountID,depositID string) KuMsgActiveDeposit
	msg := singerTypes.NewKuMsgActiveDeposit(addAlice, accAlice, depositID)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("activeDeposit error log", "err", err)
	return err
}

func transferDeposit(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, from, to types.AccountID, depositID string, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgTransferDeposit(auth sdk.AccAddress,depositID string,from,to AccountID,memo string ) KuMsgTransferDeposit
	msg := depositTypes.NewKuMsgTransferDeposit(addAlice, depositID, from, to, "")

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, from, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("transferDeposit error log", "err", err)
	return err
}

func depositToCoin(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	msg := depositTypes.NewKuMsgDepositToCoin(addAlice, depositID, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("depositToCoin error log", "err", err)
	return err
}

func coinPowerToCoin(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewMsgExercise(auth types.AccAddress, id types.AccountID, amount types.Coin)
	msg := assetTypes.NewMsgExercise(addAlice, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("coinPowerToCoin error log", "err", err)
	return err
}

func transferCoin(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, from, to types.AccountID, amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewMsgTransfer(auth types.AccAddress, from types.AccountID, to types.AccountID, amount Coins)
	msg := assetTypes.NewMsgTransfer(addAlice, from, to, types.Coins{amount})

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, from, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("transferCoin error log", "err", err)
	return err
}

func depositClaimCoin(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, amount types.Coin, claimAddress string, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgDepositClaimCoin(auth sdk.AccAddress,depositID string,owner AccountID,asset Coin,claimAddress string )
	msg := depositTypes.NewKuMsgDepositClaimCoin(addAlice, depositID, accAlice, amount, claimAddress)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("depositToCoin error log", "err", err)
	return err
}

func submitSingerSpv(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, spvInfo singerTypes.SpvInfo, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgSubmitSpv(auth sdk.AccAddress,spvInfo SpvInfo ) KuMsgSubmitSpv
	msg := singerTypes.NewKuMsgSubmitSpv(addAlice, spvInfo)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("submitSingerSpv error log", "err", err)
	return err
}

func finishDeposit(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgFinishDeposit(auth sdk.AccAddress,depositID string,owner AccountID ) KuMsgFinishDeposit
	msg := depositTypes.NewKuMsgFinishDeposit(addAlice, depositID, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("finishDeposit error log", "err", err)
	return err
}

func signBlock(app *simapp.SimApp) {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	blockTime := time.Now().Add(app.DepositKeeper().WaitTime(ctxCheck) * 2)
	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: blockTime}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})
	app.EndBlock(abci.RequestEndBlock{})
	app.Commit()
}

func depositTimeOut(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgWaitTimeout(auth sdk.AccAddress,depositID string,owner AccountID ) KuMsgWaitTimeout {
	msg := depositTypes.NewKuMsgWaitTimeout(addAlice, depositID, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	blockTime := time.Now().Add(app.DepositKeeper().WaitTime(ctxCheck) * 2)
	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: blockTime}
	ctxCheck.Logger().Info("time", "blockTime", blockTime)

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("depositTimeOut error log", "err", err)
	return err
}

func singerTimeOut(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgWaitTimeout(auth sdk.AccAddress,depositID string,singerAccount AccountID ) KuMsgWaitTimeout
	msg := singerTypes.NewKuMsgWaitTimeout(addAlice, depositID, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	blockTime := time.Now().Add(app.DepositKeeper().WaitTime(ctxCheck) * 2)
	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: blockTime}
	ctxCheck.Logger().Info("time", "blockTime", blockTime)

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("singerTimeOut error log", "err", err)
	return err
}

func singerReportWrongSpv(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgReportSpvWrong(auth sdk.AccAddress,depositID string,singerAccount AccountID ) KuMsgReportSpvWrong
	msg := singerTypes.NewKuMsgReportSpvWrong(addAlice, depositID, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("singerReportWrongSpv error log", "err", err)
	return err
}

func depositReportWrongSpv(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	// NewKuMsgReportWrongSpv(auth sdk.AccAddress,depositID string,owner AccountID ) KuMsgReportWrongSpv
	msg := depositTypes.NewKuMsgReportWrongSpv(addAlice, depositID, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("depositReportWrongSpv error log", "err", err)
	return err
}

func claimAberrangDeposit(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, amount types.Coin, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	// NewKuMsgClaimAberrant(auth sdk.AccAddress,depositID string,claimAccount AccountID,amount Coin ) KuMsgClaimAberrant
	msg := depositTypes.NewKuMsgClaimAberrant(addAlice, depositID, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("claimAberrangDeposit error log", "err", err)
	return err
}

func judgeSpv(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, spvIsRight, feeToSinger, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgJudgeDepositSpv(auth sdk.AccAddress,depositID string,systemAccount AccountID ,spvIsRight bool,feeToSinger bool)
	msg := depositTypes.NewKuMsgJudgeDepositSpv(addAlice, depositID, accAlice, spvIsRight, feeToSinger)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	blockTime := time.Now().Add(app.DepositKeeper().WaitTime(ctxCheck) * 2)
	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: blockTime}
	ctxCheck.Logger().Info("time", "blockTime", blockTime)

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("judgeSpv error log", "err", err)
	return err
}

func forceCashReadyDeposit(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	msg := depositTypes.NewKuMsgCashReadyDeposit(addAlice, depositID, accAlice)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	blockTime := time.Now().Add(app.DepositKeeper().WaitTime(ctxCheck) * 2)
	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: blockTime}
	ctxCheck.Logger().Info("time", "blockTime", blockTime)

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("forceCashReadyDeposit error log", "err", err)
	return err
}

func claimDepositMortgage(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, depositID string, amount types.Coin, passed bool) error {
	signBlock(app)
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	// NewKuMsgClaimMortgage(auth sdk.AccAddress,depositID string,claimAccount AccountID,amount Coin ) KuMsgClaimMortgage
	msg := depositTypes.NewKuMsgClaimMortgage(addAlice, depositID, accAlice, amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("claimDepositMortgage error log", "err", err)
	return err
}

func checkDepositStatus(app *simapp.SimApp, depositID string, checkStatus depositTypes.DepositStatus) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	keeper := app.DepositKeeper()
	depositInfo, found := keeper.GetDepositInfo(ctxCheck, depositID)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "deposit not found")
	}
	if depositInfo.Status == checkStatus {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "status not equal")
}

func checkSingerDepositStatus(app *simapp.SimApp, depositID string, checkStatus singerTypes.DepositStatus) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	keeper := app.SingerKeeper()
	depositInfo, found := keeper.GetDepositInfo(ctxCheck, depositID)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "singer deposit not found")
	}
	if depositInfo.Status == checkStatus {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "singer deposit status not equal")
}

func printSingerInfo(app *simapp.SimApp) {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	keeper := app.SingerKeeper()
	allSinger := keeper.GetAllSingerInfo(ctxCheck)
	ctxCheck.Logger().Info("allSinger info", "allSinger", allSinger)
}

func printDepositInfo(app *simapp.SimApp) {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	keeper := app.DepositKeeper()
	allDeposit := keeper.GetAllDepositInfo(ctxCheck)
	ctxCheck.Logger().Info("allDeposit info", "allDeposit", allDeposit)
}

func checkSingerMortgage(app *simapp.SimApp, singer types.AccountID, prestoreMortgage, lockMortgage sdk.Int) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	keeper := app.SingerKeeper()
	singerInfo, found := keeper.GetSingerInfo(ctxCheck, singer)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "singer  not found")
	}
	if singerInfo.LockMortgage.Equal(lockMortgage) && singerInfo.SignatureMortgage.Equal(prestoreMortgage) {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "singer mortgage does not match")
}

func checkFee(app *simapp.SimApp, owner types.AccountID, preStoreFee, lockFee sdk.Int) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})
	keeper := app.PriceFeeKeeper()
	feeInfo, found := keeper.GetFeeInfo(ctxCheck, owner)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "feeInfo  not found")
	}
	if feeInfo.PrestoreFee.Equal(preStoreFee) && feeInfo.LockedFee.Equal(lockFee) {
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "fee info does not match")
}

func TestDepositHandler(t *testing.T) {
	config.SealChainConfig()
	wallet := simapp.NewWallet()
	Convey("TestCreateLeginCoin", t, func() {
		addAlice, _, _, accAlice, _, _, app := newTestApp(wallet)
		accSystem := types.MustAccountID("test@sys")
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		leginCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(100))
		err := prohibitLegalCoin(t, wallet, app, addAlice, accSystem, leginCoin, false)
		So(err, ShouldNotBeNil)
		err = createLeginCoin(t, wallet, app, addAlice, accAlice, leginCoin, symbol, false)
		So(err, ShouldNotBeNil)
		err = permintLegalCoin(t, wallet, app, addAlice, accSystem, leginCoin, false)
		So(err, ShouldNotBeNil)
		err = createLeginCoin(t, wallet, app, addAlice, accSystem, leginCoin, symbol, true)
		So(err, ShouldBeNil)
		err = permintLegalCoin(t, wallet, app, addAlice, accSystem, leginCoin, true)
		So(err, ShouldBeNil)
		err = prohibitLegalCoin(t, wallet, app, addAlice, accAlice, leginCoin, false)
		So(err, ShouldNotBeNil)
		err = prohibitLegalCoin(t, wallet, app, addAlice, accSystem, leginCoin, true)
		So(err, ShouldBeNil)
		err = permintLegalCoin(t, wallet, app, addAlice, accAlice, leginCoin, false)
		So(err, ShouldNotBeNil)
		err = permintLegalCoin(t, wallet, app, addAlice, accSystem, leginCoin, true)
		So(err, ShouldBeNil)
	})
	Convey("TestNormalDeposit", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Open)
		So(err, ShouldBeNil)
		//check singer mortgage
		lockMortgage := depositCoin.Amount.QuoRaw(int64(2))
		preStoreMortgage := sdk.NewInt(1000000).Sub(lockMortgage)
		for _, singer := range singers {
			err = checkSingerMortgage(app, singer, preStoreMortgage, lockMortgage)
			So(err, ShouldBeNil)
		}
		lockedFee := depositCoin.Amount.MulRaw(int64(5)).QuoRaw(int64(100 * 3)).MulRaw(int64(3))
		preStoreFee := amout1.Amount.Sub(lockedFee)
		err = checkFee(app, accAlice, preStoreFee, lockedFee)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.SPVReady)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = activeDeposit(t, wallet, app, addAlice, singer, depositID, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.DepositActive)
		So(err, ShouldBeNil)
		singerGetFee := lockedFee.QuoRaw(int64(3))
		for _, singer := range singers {
			err = checkFee(app, singer, sdk.ZeroInt(), singerGetFee)
			So(err, ShouldBeNil)
		}
		err = checkFee(app, accAlice, preStoreFee, sdk.ZeroInt())
		So(err, ShouldBeNil)
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = depositToCoin(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		claimAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Cashing)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Cashing)
		So(err, ShouldBeNil)
		preStoreFee = preStoreFee.Sub(lockedFee)
		err = checkFee(app, accAlice, preStoreFee, lockedFee)
		So(err, ShouldBeNil)
		spvSinger := singers[0]
		singerSpv := singerTypes.NewSpvInfo(depositID, spvSinger, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitSingerSpv(t, wallet, app, addAlice, spvSinger, singerSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashOut)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.CashOut)
		So(err, ShouldBeNil)
		err = finishDeposit(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Close)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = checkSingerMortgage(app, singer, sdk.NewInt(1000000), sdk.ZeroInt())
			So(err, ShouldBeNil)
			err = checkFee(app, singer, lockedFee.MulRaw(int64(2)).QuoRaw(int64(3)), sdk.ZeroInt())
			So(err, ShouldBeNil)
		}
		err = checkFee(app, accAlice, preStoreFee, sdk.ZeroInt())
		So(err, ShouldBeNil)
	})
	Convey("TestTimeOutDeposit", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		lockedFee := depositCoin.Amount.MulRaw(int64(5)).QuoRaw(int64(100 * 3)).MulRaw(int64(3))
		preStoreFee := amout1.Amount.Sub(lockedFee)
		err = checkFee(app, accAlice, preStoreFee, lockedFee)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		err = depositTimeOut(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		singerGetFee := lockedFee.QuoRaw(int64(3))
		for _, singer := range singers {
			err = checkFee(app, singer, sdk.ZeroInt(), singerGetFee)
			So(err, ShouldBeNil)
		}
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = depositToCoin(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		claimAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, true)
		So(err, ShouldBeNil)
		preStoreFee = preStoreFee.Sub(lockedFee)
		err = checkFee(app, accAlice, preStoreFee, lockedFee)
		err = checkDepositStatus(app, depositID, depositTypes.Cashing)
		So(err, ShouldBeNil)
		spvSinger := singers[0]
		singerSpv := singerTypes.NewSpvInfo(depositID, spvSinger, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitSingerSpv(t, wallet, app, addAlice, spvSinger, singerSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashOut)
		So(err, ShouldBeNil)
		err = singerTimeOut(t, wallet, app, addAlice, spvSinger, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = checkSingerMortgage(app, singer, sdk.NewInt(1000000), sdk.ZeroInt())
			So(err, ShouldBeNil)
			err = checkFee(app, singer, lockedFee.MulRaw(int64(2)).QuoRaw(int64(3)), sdk.ZeroInt())
			So(err, ShouldBeNil)
		}
		err = checkFee(app, accAlice, preStoreFee, sdk.ZeroInt())
		So(err, ShouldBeNil)
	})
	Convey("TestjudgespvRightdeposit", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		lockedFee := depositCoin.Amount.MulRaw(int64(5)).QuoRaw(int64(100 * 3)).MulRaw(int64(3))
		preStoreFee := amout1.Amount.Sub(lockedFee)
		err = checkFee(app, accAlice, preStoreFee, lockedFee)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		err = singerReportWrongSpv(t, wallet, app, addAlice, singers[0], depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.WrongDepositSPV)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.WrongDepositSPV)
		So(err, ShouldBeNil)
		//judge spv right
		accSystem := types.MustAccountID("test@sys")
		err = judgeSpv(t, wallet, app, addAlice, accSystem, depositID, true, true, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		singerGetFee := lockedFee.QuoRaw(int64(3))
		for _, singer := range singers {
			err = checkFee(app, singer, sdk.ZeroInt(), singerGetFee)
			So(err, ShouldBeNil)
		}
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = depositToCoin(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		claimAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Cashing)
		So(err, ShouldBeNil)
		preStoreFee = preStoreFee.Sub(lockedFee)
		err = checkFee(app, accAlice, preStoreFee, lockedFee)
		spvSinger := singers[0]
		singerSpv := singerTypes.NewSpvInfo(depositID, spvSinger, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitSingerSpv(t, wallet, app, addAlice, spvSinger, singerSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashOut)
		So(err, ShouldBeNil)
		err = depositReportWrongSpv(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.WrongSingerSPV)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.WrongSingerSPV)
		So(err, ShouldBeNil)
		err = judgeSpv(t, wallet, app, addAlice, accSystem, depositID, true, true, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = checkSingerMortgage(app, singer, sdk.NewInt(1000000), sdk.ZeroInt())
			So(err, ShouldBeNil)
			err = checkFee(app, singer, lockedFee.MulRaw(int64(2)).QuoRaw(int64(3)), sdk.ZeroInt())
			So(err, ShouldBeNil)
		}
		err = checkFee(app, accAlice, preStoreFee, sdk.ZeroInt())
		So(err, ShouldBeNil)
	})
	Convey("TestsingerWrongSPVClaimdeposit", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = activeDeposit(t, wallet, app, addAlice, singer, depositID, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = depositToCoin(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		claimAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Cashing)
		So(err, ShouldBeNil)
		spvSinger := singers[0]
		singerSpv := singerTypes.NewSpvInfo(depositID, spvSinger, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitSingerSpv(t, wallet, app, addAlice, spvSinger, singerSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashOut)
		So(err, ShouldBeNil)
		err = depositReportWrongSpv(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		accSystem := types.MustAccountID("test@sys")
		err = judgeSpv(t, wallet, app, addAlice, accSystem, depositID, false, false, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Aberrant)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Aberrant)
		So(err, ShouldBeNil)
		err = coinPowerToCoin(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = claimAberrangDeposit(t, wallet, app, addAlice, accAlice, depositID, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Close)
		So(err, ShouldBeNil)
		lockMortgage := depositCoin.Amount.QuoRaw(int64(2))
		preStoreMortgage := sdk.NewInt(1000000).Sub(lockMortgage)
		for _, singer := range singers {
			err = checkSingerMortgage(app, singer, preStoreMortgage, sdk.ZeroInt())
			So(err, ShouldBeNil)
			err = checkFee(app, singer, sdk.ZeroInt(), sdk.ZeroInt())
			So(err, ShouldBeNil)
		}
		err = checkFee(app, accAlice, amout1.Amount, sdk.ZeroInt())
		So(err, ShouldBeNil)
	})
	Convey("TestsingerNoSPVClaimdeposit", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = activeDeposit(t, wallet, app, addAlice, singer, depositID, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = depositToCoin(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		claimAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Cashing)
		So(err, ShouldBeNil)
		err = depositTimeOut(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Aberrant)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Aberrant)
		So(err, ShouldBeNil)
		err = coinPowerToCoin(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = claimAberrangDeposit(t, wallet, app, addAlice, accAlice, depositID, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Close)
		So(err, ShouldBeNil)
		lockMortgage := depositCoin.Amount.QuoRaw(int64(2))
		preStoreMortgage := sdk.NewInt(1000000).Sub(lockMortgage)
		for _, singer := range singers {
			err = checkSingerMortgage(app, singer, preStoreMortgage, sdk.ZeroInt())
			So(err, ShouldBeNil)
		}
	})
	Convey("TestForceDepositToCoin", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = activeDeposit(t, wallet, app, addAlice, singer, depositID, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = forceCashReadyDeposit(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		claimAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Cashing)
		So(err, ShouldBeNil)
		spvSinger := singers[0]
		singerSpv := singerTypes.NewSpvInfo(depositID, spvSinger, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitSingerSpv(t, wallet, app, addAlice, spvSinger, singerSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashOut)
		So(err, ShouldBeNil)
		err = finishDeposit(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
	})
	Convey("TestClaimCoin", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		for _, singer := range singers {
			err = activeDeposit(t, wallet, app, addAlice, singer, depositID, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = forceCashReadyDeposit(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		amoutPrice1 := types.NewInt64Coin(constants.DefaultBondDenom, 1000)
		amountPrice2 := types.NewCoin(otherCoinDenom, sdk.NewInt(100))
		accSystem := types.MustAccountID("test@sys")
		err = setPrice(t, wallet, app, addAlice, accSystem, amountPrice2, amoutPrice1, true)
		So(err, ShouldBeNil)
		err = claimDepositMortgage(t, wallet, app, addAlice, accAlice, depositID, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
	})
	Convey("TestDepositNoAddress", t, func() {
		addAlice, _, _, accAlice, _, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(10000))
		err, depositID, _ := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		err = depositTimeOut(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Close)
		So(err, ShouldBeNil)
	})
	Convey("TestDepositNoSPV", t, func() {
		addAlice, _, _, accAlice, _, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		lockedFee := depositCoin.Amount.MulRaw(int64(5)).QuoRaw(int64(100 * 3)).MulRaw(int64(3))
		preStoreFee := amout1.Amount.Sub(lockedFee)
		err = checkFee(app, accAlice, preStoreFee, lockedFee)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		spvSinger := singers[0]
		err = singerTimeOut(t, wallet, app, addAlice, spvSinger, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Close)
		So(err, ShouldBeNil)
		err = checkFee(app, accAlice, preStoreFee, sdk.ZeroInt())
		So(err, ShouldBeNil)
		singerFee := lockedFee.QuoRaw(int64(3))
		for _, singer := range singers {
			err = checkFee(app, singer, singerFee, sdk.ZeroInt())
			So(err, ShouldBeNil)
		}
	})
	Convey("TestDepositWrongSPV", t, func() {
		addAlice, _, _, accAlice, _, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		//get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		for _, singer := range singers {
			err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		err = singerReportWrongSpv(t, wallet, app, addAlice, singers[0], depositID, true)
		So(err, ShouldBeNil)
		//judge spv right
		accSystem := types.MustAccountID("test@sys")
		err = judgeSpv(t, wallet, app, addAlice, accSystem, depositID, false, true, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Close)
		So(err, ShouldBeNil)

		for _, singer := range singers {
			err = checkSingerMortgage(app, singer, sdk.NewInt(1000000), sdk.ZeroInt())
			So(err, ShouldBeNil)
		}
	})
	Convey("TestSomethingWrong", t, func() {
		addAlice, addJack, _, accAlice, accJack, _, app := newTestApp(wallet)
		err := readyForDeposit(t, wallet, app, addAlice)
		So(err, ShouldBeNil)
		// openfee prestorefee
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, depositID, singers := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.SingerReady)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Open)
		So(err, ShouldBeNil)
		printDepositInfo(app)
		err, _, _ = createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, false)
		So(err, ShouldNotBeNil)
		// //get deposit ID
		btcAddress := "bc1q6yrjchkkyp8yc4cqwhp0p9tysvm6luecxqt8l5"
		claimAddress := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"

		for i, singer := range singers {
			if i&1 == 0 {
				err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, btcAddress, true)
			} else {
				err = setDepositAddress(t, wallet, app, addAlice, singer, depositID, claimAddress, true)
			}
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldNotBeNil)

		err = setDepositAddress(t, wallet, app, addAlice, singers[1], depositID, btcAddress, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.AddressReady)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.AddressReady)
		So(err, ShouldBeNil)
		testByte := []byte("just for test")
		newSpv := singerTypes.NewSpvInfo(depositID, accAlice, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, true)
		So(err, ShouldBeNil)
		err = submitDepositSpv(t, wallet, app, addAlice, accAlice, newSpv, false)
		So(err, ShouldNotBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.DepositSpvReady)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.SPVReady)
		So(err, ShouldBeNil)
		for i, singer := range singers {
			if i&1 == 0 {
				err = activeDeposit(t, wallet, app, addAlice, singer, depositID, true)
			}
			So(err, ShouldBeNil)
		}
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldNotBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.DepositActive)
		So(err, ShouldNotBeNil)
		err = activeDeposit(t, wallet, app, addAlice, singers[1], depositID, true)
		err = checkDepositStatus(app, depositID, depositTypes.Active)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.DepositActive)
		So(err, ShouldBeNil)

		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = transferDeposit(t, wallet, app, addAlice, accAlice, accJack, depositID, false)
		So(err, ShouldNotBeNil)
		err = depositToCoin(t, wallet, app, addJack, accJack, depositID, true)
		So(err, ShouldBeNil)
		err = depositToCoin(t, wallet, app, addJack, accJack, depositID, false)
		So(err, ShouldNotBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashReady)
		So(err, ShouldBeNil)
		// coinpower to coin
		err = coinPowerToCoin(t, wallet, app, addJack, accJack, depositCoin, true)
		So(err, ShouldBeNil)
		err = transferCoin(t, wallet, app, addJack, accJack, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
		err = depositClaimCoin(t, wallet, app, addJack, accJack, depositID, depositCoin, claimAddress, false)
		So(err, ShouldNotBeNil)
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, true)
		So(err, ShouldBeNil)
		err = depositClaimCoin(t, wallet, app, addAlice, accAlice, depositID, depositCoin, claimAddress, false)
		So(err, ShouldNotBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Cashing)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Cashing)
		So(err, ShouldBeNil)
		spvSinger := singers[0]
		singerSpv := singerTypes.NewSpvInfo(depositID, spvSinger, testByte, testByte, testByte, testByte, testByte, testByte, 0, 0)
		err = submitSingerSpv(t, wallet, app, addAlice, spvSinger, singerSpv, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.CashOut)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.CashOut)
		So(err, ShouldBeNil)
		err = finishDeposit(t, wallet, app, addAlice, accAlice, depositID, true)
		So(err, ShouldBeNil)
		err = checkDepositStatus(app, depositID, depositTypes.Finish)
		So(err, ShouldBeNil)
		err = checkSingerDepositStatus(app, depositID, singerTypes.Close)
		So(err, ShouldBeNil)
	})
	Convey("TestSomethingWrongBeforeDeposit", t, func() {
		addAlice, _, _, accAlice, _, _, app := newTestApp(wallet)
		symbol := types.MustName("btc")
		otherCoinDenom := types.CoinDenom(depositTypes.ModuleAccountName, symbol)
		depositCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(1000000))
		err, _, _ := createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, false)
		So(err, ShouldNotBeNil)

		leginCoin := types.NewCoin(otherCoinDenom, sdk.NewInt(100000000000000000))
		accSystem := types.MustAccountID("test@sys")
		err = createLeginCoin(t, wallet, app, addAlice, accSystem, leginCoin, symbol, true)
		So(err, ShouldBeNil)
		err, _, _ = createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, false)
		So(err, ShouldNotBeNil)
		err = openFee(t, wallet, app, addAlice, accAlice, true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err = preStoreFee(t, wallet, app, addAlice, accAlice, amout1, true)
		So(err, ShouldBeNil)
		err, _, _ = createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, false)
		So(err, ShouldNotBeNil)
		//------------------------------------------------------------------------------------
		singera := types.MustAccountID("singera")
		singerb := types.MustAccountID("singerb")
		singerc := types.MustAccountID("singerc")
		singerd := types.MustAccountID("singerd")
		singere := types.MustAccountID("singere")
		singerf := types.MustAccountID("singerf")
		singerg := types.MustAccountID("singerg")

		err = regesterSinger(t, wallet, app, addAlice, singera, true)
		So(err, ShouldBeNil)
		err = regesterSinger(t, wallet, app, addAlice, singerb, true)
		So(err, ShouldBeNil)
		err = regesterSinger(t, wallet, app, addAlice, singerc, true)
		So(err, ShouldBeNil)
		err = regesterSinger(t, wallet, app, addAlice, singerd, true)
		So(err, ShouldBeNil)
		err = regesterSinger(t, wallet, app, addAlice, singere, true)
		So(err, ShouldBeNil)
		err = regesterSinger(t, wallet, app, addAlice, singerf, true)
		So(err, ShouldBeNil)
		err = regesterSinger(t, wallet, app, addAlice, singerg, true)
		So(err, ShouldBeNil)

		amout2 := types.NewInt64Coin(constants.DefaultBondDenom, 1000000000000000000)
		err = payAccesss(t, wallet, app, addAlice, singera, amout2, true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app, addAlice, singerb, amout2, true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app, addAlice, singerc, amout2, true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app, addAlice, singerd, amout2, true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app, addAlice, singere, amout2, true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app, addAlice, singerf, amout2, true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app, addAlice, singerg, amout2, true)
		So(err, ShouldBeNil)

		err = activeSinger(t, wallet, app, addAlice, accSystem, singera, true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app, addAlice, accSystem, singerb, true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app, addAlice, accSystem, singerc, true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app, addAlice, accSystem, singerd, true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app, addAlice, accSystem, singere, true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app, addAlice, accSystem, singerf, true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app, addAlice, accSystem, singerg, true)
		So(err, ShouldBeNil)

		amoutMortgage := types.NewInt64Coin(constants.DefaultBondDenom, 1000000)
		err = payMortgage(t, wallet, app, addAlice, singera, amoutMortgage, true)
		So(err, ShouldBeNil)
		err = payMortgage(t, wallet, app, addAlice, singerb, amoutMortgage, true)
		So(err, ShouldBeNil)
		err = payMortgage(t, wallet, app, addAlice, singerc, amoutMortgage, true)
		So(err, ShouldBeNil)
		err = payMortgage(t, wallet, app, addAlice, singerd, amoutMortgage, true)
		So(err, ShouldBeNil)
		err = payMortgage(t, wallet, app, addAlice, singere, amoutMortgage, true)
		So(err, ShouldBeNil)
		err = payMortgage(t, wallet, app, addAlice, singerf, amoutMortgage, true)
		So(err, ShouldBeNil)
		err = payMortgage(t, wallet, app, addAlice, singerg, amoutMortgage, true)
		So(err, ShouldBeNil)
		//------------------------------------------------------------------------------------
		err, _, _ = createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, false)
		So(err, ShouldNotBeNil)
		amoutPrice1 := types.NewInt64Coin(constants.DefaultBondDenom, 100)
		amountPrice := types.NewCoin(otherCoinDenom, sdk.NewInt(100))
		err = setPrice(t, wallet, app, addAlice, accSystem, amountPrice, amoutPrice1, true)
		So(err, ShouldBeNil)
		err = prohibitLegalCoin(t, wallet, app, addAlice, accSystem, leginCoin, true)
		So(err, ShouldBeNil)
		err, _, _ = createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, false)
		So(err, ShouldNotBeNil)
		err = permintLegalCoin(t, wallet, app, addAlice, accSystem, leginCoin, true)
		So(err, ShouldBeNil)
		err, _, _ = createDeposit(t, wallet, app, addAlice, accAlice, depositCoin, true)
		So(err, ShouldBeNil)
	})
}
