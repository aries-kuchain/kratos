package singer_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/smartystreets/goconvey/convey"

	abci "github.com/tendermint/tendermint/abci/types"

	// "encoding/hex"

	"github.com/KuChainNetwork/kuchain/chain/config"
	"github.com/KuChainNetwork/kuchain/chain/constants"
//	"github.com/KuChainNetwork/kuchain/chain/msg"
	"github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/test/simapp"
	// stakingTypes "github.com/KuChainNetwork/kuchain/x/staking/types"
	// "github.com/tendermint/tendermint/crypto"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	// "github.com/tendermint/tendermint/crypto/ed25519"
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
	singerd:= types.MustAccountID("singerd")
	singere := types.MustAccountID("singere")
	singerf := types.MustAccountID("singerf")
	singerg := types.MustAccountID("singerg")
	accSystem :=  types.MustAccountID("test@sys")

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

func regesterSinger(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)

	//NewKuMsgRegisterSinger(auth sdk.AccAddress, singerAccount AccountID) 
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

func payAccesss(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice types.AccountID,amount types.Coin, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)

	msg := singerTypes.NewKuMsgPayAccess(addAlice, accAlice,amount)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("payAccesss error log", "err", err)
	return err
}

func activeSinger(t *testing.T, wallet *simapp.Wallet, app *simapp.SimApp, addAlice sdk.AccAddress, accAlice,accJack types.AccountID, passed bool) error {
	ctxCheck := app.BaseApp.NewContext(true, abci.Header{Height: app.LastBlockHeight() + 1})

	origAuthSeq, origAuthNum, err := app.AccountKeeper().GetAuthSequence(ctxCheck, addAlice)
	So(err, ShouldBeNil)

	ctxCheck.Logger().Info("auth nums", "seq", origAuthSeq, "num", origAuthNum)

	msg := singerTypes.NewKuMsgActiveSinger(addAlice, accAlice,accJack)

	fee := types.Coins{types.NewInt64Coin(constants.DefaultBondDenom, 1000000)}
	header := abci.Header{Height: app.LastBlockHeight() + 1}

	_, _, err = simapp.SignCheckDeliver(t, app.Codec(), app.BaseApp,
		header, accAlice, fee,
		[]sdk.Msg{msg}, []uint64{origAuthNum}, []uint64{origAuthSeq},
		passed, passed, wallet.PrivKey(addAlice))
	ctxCheck.Logger().Info("payAccesss error log", "err", err)
	return err
}


func TestSingerHandler(t *testing.T) {
		config.SealChainConfig()
		wallet := simapp.NewWallet()
	Convey("TestRegisterSinger", t, func() {
		addAlice, addJack, _, accAlice, _, _, app := newTestApp(wallet)
		newTestApp(wallet)
		err := regesterSinger(t, wallet, app,addAlice,accAlice,true)
		So(err, ShouldBeNil)
		err = regesterSinger(t, wallet, app,addAlice,accAlice,false)
		So(err, ShouldNotBeNil)
		err = regesterSinger(t, wallet, app,addJack,accAlice,false)
		So(err, ShouldNotBeNil)
	})
	Convey("TestPayAccess", t, func() {
		addAlice, _, _, accAlice, _, _, app := newTestApp(wallet)
		newTestApp(wallet)
		err := regesterSinger(t, wallet, app,addAlice,accAlice,true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		amout2 := types.NewInt64Coin(constants.DefaultBondDenom, 1000000000000000000)
		otherCoinDenom := types.CoinDenom(types.MustName("foo"), types.MustName("coin"))
		otherAmount := types.NewInt64Coin(otherCoinDenom, 10000000)
		err = payAccesss(t, wallet, app,addAlice,accAlice,amout1,true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app,addAlice,accAlice,amout2,true)
		So(err, ShouldBeNil)
		err = payAccesss(t, wallet, app,addAlice,accAlice,otherAmount,false)
		So(err, ShouldNotBeNil)
	})
	Convey("TestActiveSinger", t, func() {
		addAlice, _, _, accAlice, _, _, app := newTestApp(wallet)
		newTestApp(wallet)
		err := regesterSinger(t, wallet, app,addAlice,accAlice,true)
		So(err, ShouldBeNil)
		amout1 := types.NewInt64Coin(constants.DefaultBondDenom, 10000000)
		amout2 := types.NewInt64Coin(constants.DefaultBondDenom, 1000000000000000000)
		err = payAccesss(t, wallet, app,addAlice,accAlice,amout1,true)
		So(err, ShouldBeNil)
		accSystem :=  types.MustAccountID("test@sys")
		err = activeSinger(t, wallet, app,addAlice,accSystem,accAlice,false)
		So(err, ShouldNotBeNil)
		err = payAccesss(t, wallet, app,addAlice,accAlice,amout2,true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app,addAlice,accAlice,accAlice,false)
		So(err, ShouldNotBeNil)
		err = activeSinger(t, wallet, app,addAlice,accSystem,accAlice,true)
		So(err, ShouldBeNil)
		err = activeSinger(t, wallet, app,addAlice,accSystem,accAlice,false)
		So(err, ShouldNotBeNil)
	})
}