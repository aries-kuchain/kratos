package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	yaml "gopkg.in/yaml.v2"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type DepositStatus int32

// staking constants
const (
	Init            DepositStatus = 1
	SingerReady     DepositStatus = 2
	AddressReady    DepositStatus = 3
	DepositSpvReady DepositStatus = 4
	Active          DepositStatus = 5
	CashReady       DepositStatus = 6
	Cashing         DepositStatus = 7
	CashOut         DepositStatus = 8
	Finish          DepositStatus = 9
	Aberrant    DepositStatus = 10
	WrongDepositSPV  DepositStatus = 11
	WrongSingerSPV  DepositStatus = 12

	DepositTypeInit            = "Init"
	DepositTypeSingerReady     = "SingerReady"
	DepositTypeAddressReady    = "DepositAddressReady"
	DepositTypeDepositSpvReady = "DepositSpvReady"
	DepositTypeActive          = "Active"
	DepositTypeCashReady       = "CashReady"
	DepositTypeCashing         = "Cashing"
	DepositTypeCashOut         = "CashOut"
	DepositTypeFinish          = "Finish"
	DepositTypeAberrant          = "Aberrant"
	DepositTypeWrongDepositSPV          = "WrongDepositSPV"
	DepositTypeWrongSingerSPV          = "WrongSingerSPV"
)

// Equal compares two BondStatus instances
func (b DepositStatus) Equal(b2 DepositStatus) bool {
	return byte(b) == byte(b2)
}

func (b DepositStatus) String() string {
	switch b {
	case Init:
		return DepositTypeInit
	case SingerReady:
		return DepositTypeSingerReady
	case AddressReady:
		return DepositTypeAddressReady
	case DepositSpvReady:
		return DepositTypeDepositSpvReady
	case Active:
		return DepositTypeActive
	case CashReady:
		return DepositTypeCashReady
	case Cashing:
		return DepositTypeCashing
	case CashOut:
		return DepositTypeCashOut
	case Finish:
		return DepositTypeFinish
	case Aberrant:
		return DepositTypeAberrant
	default:
		panic("invalid deposit status")
	}
}

type DepositInfo struct {
	DepositID       string //用户名，时间，数额等等
	Owner           AccountID
	Asset           Coin
	Singers         []AccountID
	DepositAddress  string
	WithDrawAddress string
	Status          DepositStatus
	CurrentFee sdk.Int
	TotalFee sdk.Int
	DepositChangeTime time.Time
}

func NewDepositInfo(depositID string, owner AccountID, asset Coin) DepositInfo {
	//	depositID := fmt.Sprintf("%s-%s-%s",owner.String(),asset.String(), time.Now().UTC().Format("2006-01-02 15:04:05") )
	return DepositInfo{
		DepositID: depositID,
		Owner:     owner,
		Asset:     asset,
		Status:    Init,
	}
}

// return the redelegation
func MustMarshalDepositInfo(cdc *codec.Codec, depositInfo DepositInfo) []byte {
	return cdc.MustMarshalBinaryBare(&depositInfo)
}

// unmarshal a redelegation from a store value
func MustUnmarshalDepositInfo(cdc *codec.Codec, value []byte) DepositInfo {
	depositInfo, err := UnmarshalDepositInfo(cdc, value)
	if err != nil {
		panic(err)
	}
	return depositInfo
}

// unmarshal a redelegation from a store value
func UnmarshalDepositInfo(cdc *codec.Codec, value []byte) (v DepositInfo, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

// String implements the Stringer interface for a DepositInfo object.
func (v DepositInfo) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

func (v *DepositInfo) SetSingers(pickedSingers singerTypes.SingerInfos) {
	for _,singerInfo := range pickedSingers {
		v.Singers = append(v.Singers ,singerInfo.SingerAccount)
	}
	v.Status = SingerReady
}

//--------------------------------------------------------
type LegalCoinStatus uint32

const (
	Permint  LegalCoinStatus = 1
	Prohibit LegalCoinStatus = 2
)

type LegalCoin struct {
	Symbol  Name `json:"symbol" yaml:"symbol"`
	Asset  Coin `json:"asset" yaml:"asset"`
	Status LegalCoinStatus `json:"status" yaml:"status"`
}

func NewLegalCoin(asset Coin,symbol Name) LegalCoin {
	return LegalCoin{
		Asset:  asset,
		Status: Permint,
		Symbol:symbol,
	}
}

// return the redelegation
func MustMarshalLegalCoin(cdc *codec.Codec, depositInfo LegalCoin) []byte {
	return cdc.MustMarshalBinaryBare(&depositInfo)
}

// unmarshal a redelegation from a store value
func MustUnmarshalLegalCoin(cdc *codec.Codec, value []byte) LegalCoin {
	depositInfo, err := UnmarshalLegalCoin(cdc, value)
	if err != nil {
		panic(err)
	}
	return depositInfo
}

// unmarshal a redelegation from a store value
func UnmarshalLegalCoin(cdc *codec.Codec, value []byte) (v LegalCoin, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

// String implements the Stringer interface for a DepositInfo object.
func (v LegalCoin) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}
