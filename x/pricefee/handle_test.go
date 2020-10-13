package pricefee_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/smartystreets/goconvey/convey"

	abci "github.com/tendermint/tendermint/abci/types"

	// "encoding/hex"

	"github.com/KuChainNetwork/kuchain/chain/config"
	"github.com/KuChainNetwork/kuchain/chain/constants"
	//"github.com/KuChainNetwork/kuchain/chain/msg"
	"github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/test/simapp"
	priceFeeTypes "github.com/KuChainNetwork/kuchain/x/pricefee/types"

	// stakingTypes "github.com/KuChainNetwork/kuchain/x/staking/types"
	// "github.com/tendermint/tendermint/crypto"
	// "github.com/tendermint/tendermint/crypto/ed25519"
)

func newTestApp(wallet *simapp.Wallet) (addAlice, addJack, addValidator sdk.AccAddress, accAlice, accJack, accValidator,addAccount types.AccountID, app *simapp.SimApp) {
	addAlice = wallet.NewAccAddress()
	addJack = wallet.NewAccAddress()
	addValidator = wallet.NewAccAddress()

	accAlice = types.MustAccountID("alice@ok")
	accJack = types.MustAccountID("jack@ok")
	accValidator = types.MustAccountID("validator@ok")

	accSystem := types.MustAccountID("test@sys")
	addAccount = types.NewAccountIDFromAccAdd(addAlice)

	resInt, succ := sdk.NewIntFromString("100000000000000000000000")
	if !succ {
		resInt = sdk.NewInt(10000000000000000)
	}
	otherCoinDenom := types.CoinDenom(types.MustName("foo"), types.MustName("coin"))
	initAsset := types.NewCoin(constants.DefaultBondDenom, resInt)

	asset1 := types.NewCoins(
		types.NewInt64Coin(otherCoinDenom, 67),
		initAsset)

	asset2 := types.NewCoins(
		types.NewInt64Coin(otherCoinDenom, 67),
		types.NewInt64Coin(constants.DefaultBondDenom, 10000000))

	genAlice := simapp.NewSimGenesisAccount(accAlice, addAlice).WithAsset(asset1)
	genJack := simapp.NewSimGenesisAccount(accJack, addJack).WithAsset(asset1)
	genValidator := simapp.NewSimGenesisAccount(accValidator, addValidator).WithAsset(asset2)
	genAddAccount := simapp.NewSimGenesisAccount(addAccount, addAlice).WithAsset(asset1)
	gensysAccount := simapp.NewSimGenesisAccount(accSystem, addAlice).WithAsset(asset1)


	genAccs := simapp.NewGenesisAccounts(wallet.GetRootAuth(), genAlice, genJack, genValidator,genAddAccount,gensysAccount)
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

	return addAlice, addJack, addValidator, accAlice, accJack, accValidator,addAccount, app
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

func preStoreFee(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin,passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	msg := priceFeeTypes.NewKuMsgPrestoreFee(addAlice, accAlice,amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("preStoreFee error log", "err", err)
	return err
}

func claimFee(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount types.Coin,passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgClaimFee(auth sdk.AccAddress, owner AccountID, amount Coin) KuMsgClaimFee
	msg := priceFeeTypes.NewKuMsgClaimFee(addAlice, accAlice,amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("claimFee error log", "err", err)
	return err
}

func setPrice(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, amount1,amout2 types.Coin,passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)
	//NewKuMsgSetPrice(auth sdk.AccAddress, systemAccount AccountID, base,quote Coin,remark string) KuMsgSetPrice
	msg := priceFeeTypes.NewKuMsgSetPrice(addAlice, accAlice,amount1,amout2,"")

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("setPrice error log", "err", err)
	return err
}

func TestPriceFeeHandler(t *testing.T) {
	config.SealChainConfig()
	wallet := simapp.NewWallet()
	Convey("TestOpenFee", t, func() {
		addAlice, _, _, accAlice, _, _, addAccount,app := newTestApp(wallet)
		err:= openFee(t, wallet, app,addAlice,accAlice,true)
		So(err, ShouldBeNil)
		err= openFee(t, wallet, app,addAlice,accAlice,false)
		So(err, ShouldNotBeNil)
		//address 
		err= openFee(t, wallet, app,addAlice,addAccount,true)
		So(err, ShouldBeNil)
	})
	Convey("TestPreStoreFee", t, func() {
		addAlice, _, _, accAlice, _, _, addAccount,app := newTestApp(wallet)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		err:= preStoreFee(t, wallet, app,addAlice,accAlice,amout1,false)
		So(err, ShouldNotBeNil)
		err= openFee(t, wallet, app,addAlice,accAlice,true)
		So(err, ShouldBeNil)
		err= preStoreFee(t, wallet, app,addAlice,accAlice,amout1,true)
		So(err, ShouldBeNil)
		err= preStoreFee(t, wallet, app,addAlice,addAccount,amout1,false)
		So(err, ShouldNotBeNil)
		resInt, succ := sdk.NewIntFromString("100000000000000000000000")
		if !succ {
			resInt = sdk.NewInt(10000000000000000)
		}
		bigAsset := types.NewCoin(constants.DefaultBondDenom, resInt)
		err= preStoreFee(t, wallet, app,addAlice,accAlice,bigAsset,false)
		So(err, ShouldNotBeNil)
	})
	Convey("TestPreStoreFee", t, func() {
		addAlice, _, _, accAlice, _, _, addAccount,app := newTestApp(wallet)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		amout2 := types.NewInt64Coin(constants.DefaultBondDenom, 100000000)
		err := openFee(t, wallet, app,addAlice,accAlice,true)
		So(err, ShouldBeNil)
		err= preStoreFee(t, wallet, app,addAlice,accAlice,amout1,true)
		So(err, ShouldBeNil)
		err = claimFee(t, wallet, app,addAlice,addAccount,amout1,false)
		So(err, ShouldNotBeNil)
		err = claimFee(t, wallet, app,addAlice,accAlice,amout2,false)
		So(err, ShouldNotBeNil)
		err = claimFee(t, wallet, app,addAlice,accAlice,amout1,true)
		So(err, ShouldBeNil)
	})
	Convey("TestPreStoreFee", t, func() {
		addAlice, _, _, accAlice, _, _, _,app := newTestApp(wallet)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		otherCoinDenom := types.CoinDenom(types.MustName("foo"), types.MustName("coin"))
		amout2 := types.NewInt64Coin(otherCoinDenom, 100000000)
		err := setPrice(t, wallet, app,addAlice,accAlice,amout1,amout2,false)
		So(err, ShouldNotBeNil)
		accSystem := types.MustAccountID("test@sys")
		err = setPrice(t, wallet, app,addAlice,accSystem,amout1,amout2,true)
		So(err, ShouldBeNil)
		err = setPrice(t, wallet, app,addAlice,accSystem,amout1,amout1,false)
		So(err, ShouldNotBeNil)
		err = setPrice(t, wallet, app,addAlice,accSystem,amout1,amout2,true)
		So(err, ShouldBeNil)
	})
}