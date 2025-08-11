// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GNSTradingCallbacksV641AggregatorAnswer is an auto generated low-level Go binding around an user-defined struct.
type GNSTradingCallbacksV641AggregatorAnswer struct {
	OrderId *big.Int
	Price   *big.Int
	SpreadP *big.Int
	Open    *big.Int
	High    *big.Int
	Low     *big.Int
}

// GNSTradingCallbacksV641LastUpdated is an auto generated low-level Go binding around an user-defined struct.
type GNSTradingCallbacksV641LastUpdated struct {
	Tp      uint32
	Sl      uint32
	Limit   uint32
	Created uint32
}

// GNSTradingCallbacksV641SimplifiedTradeId is an auto generated low-level Go binding around an user-defined struct.
type GNSTradingCallbacksV641SimplifiedTradeId struct {
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	TradeType uint8
}

// GNSTradingCallbacksV641TradeData is an auto generated low-level Go binding around an user-defined struct.
type GNSTradingCallbacksV641TradeData struct {
	MaxSlippageP *big.Int
	Placeholder  *big.Int
}

// StorageInterfaceV5Trade is an auto generated low-level Go binding around an user-defined struct.
type StorageInterfaceV5Trade struct {
	Trader          common.Address
	PairIndex       *big.Int
	Index           *big.Int
	InitialPosToken *big.Int
	PositionSizeDai *big.Int
	OpenPrice       *big.Int
	Buy             bool
	Leverage        *big.Int
	Tp              *big.Int
	Sl              *big.Int
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Forbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongParams\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeValueDai\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeValueDai\",\"type\":\"uint256\"}],\"name\":\"BorrowingFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"daiVaultFeeP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpFeeP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sssFeeP\",\"type\":\"uint256\"}],\"name\":\"ClosingFeeSharesPUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueDai\",\"type\":\"uint256\"}],\"name\":\"DaiVaultFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"Done\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueDai\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"distributed\",\"type\":\"bool\"}],\"name\":\"GovFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueDai\",\"type\":\"uint256\"}],\"name\":\"GovFeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structStorageInterfaceV5.Trade\",\"name\":\"t\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftHolder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumStorageInterfaceV5.LimitOrder\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceImpactP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"percentProfit\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"daiSentToTrader\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exactExecution\",\"type\":\"bool\"}],\"name\":\"LimitExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumGNSTradingCallbacksV6_4_1.CancelReason\",\"name\":\"cancelReason\",\"type\":\"uint8\"}],\"name\":\"MarketCloseCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structStorageInterfaceV5.Trade\",\"name\":\"t\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceImpactP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"percentProfit\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"daiSentToTrader\",\"type\":\"uint256\"}],\"name\":\"MarketExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumGNSTradingCallbacksV6_4_1.CancelReason\",\"name\":\"cancelReason\",\"type\":\"uint8\"}],\"name\":\"MarketOpenCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftHolder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumStorageInterfaceV5.LimitOrder\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumGNSTradingCallbacksV6_4_1.CancelReason\",\"name\":\"cancelReason\",\"type\":\"uint8\"}],\"name\":\"NftOrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"}],\"name\":\"PairMaxLeverageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueDai\",\"type\":\"uint256\"}],\"name\":\"ReferralFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueDai\",\"type\":\"uint256\"}],\"name\":\"SssFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueDai\",\"type\":\"uint256\"}],\"name\":\"TriggerFeeCharged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"borrowingFees\",\"outputs\":[{\"internalType\":\"contractGNSBorrowingFeesInterfaceV6_4\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canExecuteTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimGovFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"open\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"high\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"low\",\"type\":\"uint256\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.AggregatorAnswer\",\"name\":\"a\",\"type\":\"tuple\"}],\"name\":\"closeTradeMarketCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiVaultFeeP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"open\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"high\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"low\",\"type\":\"uint256\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.AggregatorAnswer\",\"name\":\"a\",\"type\":\"tuple\"}],\"name\":\"executeNftCloseOrderCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"open\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"high\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"low\",\"type\":\"uint256\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.AggregatorAnswer\",\"name\":\"a\",\"type\":\"tuple\"}],\"name\":\"executeNftOpenOrderCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPairsMaxLeverage\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govFeesDai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractStorageInterfaceV5\",\"name\":\"_storageT\",\"type\":\"address\"},{\"internalType\":\"contractIGNSOracleRewardsV6_4_1\",\"name\":\"_nftRewards\",\"type\":\"address\"},{\"internalType\":\"contractGNSPairInfosInterfaceV6\",\"name\":\"_pairInfos\",\"type\":\"address\"},{\"internalType\":\"contractGNSReferralsInterfaceV6_2\",\"name\":\"_referrals\",\"type\":\"address\"},{\"internalType\":\"contractGNSStakingInterfaceV6_4_1\",\"name\":\"_staking\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultToApprove\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_daiVaultFeeP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lpFeeP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sssFeeP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_canExecuteTimeout\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGNSBorrowingFeesInterfaceV6_4\",\"name\":\"_borrowingFees\",\"type\":\"address\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGNSStakingInterfaceV6_4_1\",\"name\":\"_staking\",\"type\":\"address\"},{\"internalType\":\"contractIGNSOracleRewardsV6_4_1\",\"name\":\"_oracleRewards\",\"type\":\"address\"}],\"name\":\"initializeV4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpFeeP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftRewards\",\"outputs\":[{\"internalType\":\"contractIGNSOracleRewardsV6_4_1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"open\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"high\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"low\",\"type\":\"uint256\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.AggregatorAnswer\",\"name\":\"a\",\"type\":\"tuple\"}],\"name\":\"openTradeMarketCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairInfos\",\"outputs\":[{\"internalType\":\"contractGNSPairInfosInterfaceV6\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pairMaxLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referrals\",\"outputs\":[{\"internalType\":\"contractGNSReferralsInterfaceV6_2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_daiVaultFeeP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lpFeeP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sssFeeP\",\"type\":\"uint256\"}],\"name\":\"setClosingFeeSharesP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverage\",\"type\":\"uint256\"}],\"name\":\"setPairMaxLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"setPairMaxLeverageArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"enumGNSTradingCallbacksV6_4_1.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.SimplifiedTradeId\",\"name\":\"_id\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint40\",\"name\":\"maxSlippageP\",\"type\":\"uint40\"},{\"internalType\":\"uint216\",\"name\":\"_placeholder\",\"type\":\"uint216\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.TradeData\",\"name\":\"_tradeData\",\"type\":\"tuple\"}],\"name\":\"setTradeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"enumGNSTradingCallbacksV6_4_1.TradeType\",\"name\":\"tradeType\",\"type\":\"uint8\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.SimplifiedTradeId\",\"name\":\"_id\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"tp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sl\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"created\",\"type\":\"uint32\"}],\"internalType\":\"structGNSTradingCallbacksV6_4_1.LastUpdated\",\"name\":\"_lastUpdated\",\"type\":\"tuple\"}],\"name\":\"setTradeLastUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sssFeeP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractGNSStakingInterfaceV6_4_1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageT\",\"outputs\":[{\"internalType\":\"contractStorageInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumGNSTradingCallbacksV6_4_1.TradeType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"tradeData\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"maxSlippageP\",\"type\":\"uint40\"},{\"internalType\":\"uint216\",\"name\":\"_placeholder\",\"type\":\"uint216\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumGNSTradingCallbacksV6_4_1.TradeType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"tradeLastUpdated\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"tp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sl\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"created\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BindingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// BorrowingFees is a free data retrieval call binding the contract method 0x531cc965.
//
// Solidity: function borrowingFees() view returns(address)
func (_Bindings *BindingsCaller) BorrowingFees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "borrowingFees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowingFees is a free data retrieval call binding the contract method 0x531cc965.
//
// Solidity: function borrowingFees() view returns(address)
func (_Bindings *BindingsSession) BorrowingFees() (common.Address, error) {
	return _Bindings.Contract.BorrowingFees(&_Bindings.CallOpts)
}

// BorrowingFees is a free data retrieval call binding the contract method 0x531cc965.
//
// Solidity: function borrowingFees() view returns(address)
func (_Bindings *BindingsCallerSession) BorrowingFees() (common.Address, error) {
	return _Bindings.Contract.BorrowingFees(&_Bindings.CallOpts)
}

// CanExecuteTimeout is a free data retrieval call binding the contract method 0x676c5164.
//
// Solidity: function canExecuteTimeout() view returns(uint256)
func (_Bindings *BindingsCaller) CanExecuteTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "canExecuteTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CanExecuteTimeout is a free data retrieval call binding the contract method 0x676c5164.
//
// Solidity: function canExecuteTimeout() view returns(uint256)
func (_Bindings *BindingsSession) CanExecuteTimeout() (*big.Int, error) {
	return _Bindings.Contract.CanExecuteTimeout(&_Bindings.CallOpts)
}

// CanExecuteTimeout is a free data retrieval call binding the contract method 0x676c5164.
//
// Solidity: function canExecuteTimeout() view returns(uint256)
func (_Bindings *BindingsCallerSession) CanExecuteTimeout() (*big.Int, error) {
	return _Bindings.Contract.CanExecuteTimeout(&_Bindings.CallOpts)
}

// DaiVaultFeeP is a free data retrieval call binding the contract method 0xa38ea9a8.
//
// Solidity: function daiVaultFeeP() view returns(uint256)
func (_Bindings *BindingsCaller) DaiVaultFeeP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "daiVaultFeeP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaiVaultFeeP is a free data retrieval call binding the contract method 0xa38ea9a8.
//
// Solidity: function daiVaultFeeP() view returns(uint256)
func (_Bindings *BindingsSession) DaiVaultFeeP() (*big.Int, error) {
	return _Bindings.Contract.DaiVaultFeeP(&_Bindings.CallOpts)
}

// DaiVaultFeeP is a free data retrieval call binding the contract method 0xa38ea9a8.
//
// Solidity: function daiVaultFeeP() view returns(uint256)
func (_Bindings *BindingsCallerSession) DaiVaultFeeP() (*big.Int, error) {
	return _Bindings.Contract.DaiVaultFeeP(&_Bindings.CallOpts)
}

// GetAllPairsMaxLeverage is a free data retrieval call binding the contract method 0x37edb138.
//
// Solidity: function getAllPairsMaxLeverage() view returns(uint256[])
func (_Bindings *BindingsCaller) GetAllPairsMaxLeverage(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getAllPairsMaxLeverage")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllPairsMaxLeverage is a free data retrieval call binding the contract method 0x37edb138.
//
// Solidity: function getAllPairsMaxLeverage() view returns(uint256[])
func (_Bindings *BindingsSession) GetAllPairsMaxLeverage() ([]*big.Int, error) {
	return _Bindings.Contract.GetAllPairsMaxLeverage(&_Bindings.CallOpts)
}

// GetAllPairsMaxLeverage is a free data retrieval call binding the contract method 0x37edb138.
//
// Solidity: function getAllPairsMaxLeverage() view returns(uint256[])
func (_Bindings *BindingsCallerSession) GetAllPairsMaxLeverage() ([]*big.Int, error) {
	return _Bindings.Contract.GetAllPairsMaxLeverage(&_Bindings.CallOpts)
}

// GovFeesDai is a free data retrieval call binding the contract method 0x18407fe0.
//
// Solidity: function govFeesDai() view returns(uint256)
func (_Bindings *BindingsCaller) GovFeesDai(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "govFeesDai")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GovFeesDai is a free data retrieval call binding the contract method 0x18407fe0.
//
// Solidity: function govFeesDai() view returns(uint256)
func (_Bindings *BindingsSession) GovFeesDai() (*big.Int, error) {
	return _Bindings.Contract.GovFeesDai(&_Bindings.CallOpts)
}

// GovFeesDai is a free data retrieval call binding the contract method 0x18407fe0.
//
// Solidity: function govFeesDai() view returns(uint256)
func (_Bindings *BindingsCallerSession) GovFeesDai() (*big.Int, error) {
	return _Bindings.Contract.GovFeesDai(&_Bindings.CallOpts)
}

// IsDone is a free data retrieval call binding the contract method 0x8f062227.
//
// Solidity: function isDone() view returns(bool)
func (_Bindings *BindingsCaller) IsDone(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isDone")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDone is a free data retrieval call binding the contract method 0x8f062227.
//
// Solidity: function isDone() view returns(bool)
func (_Bindings *BindingsSession) IsDone() (bool, error) {
	return _Bindings.Contract.IsDone(&_Bindings.CallOpts)
}

// IsDone is a free data retrieval call binding the contract method 0x8f062227.
//
// Solidity: function isDone() view returns(bool)
func (_Bindings *BindingsCallerSession) IsDone() (bool, error) {
	return _Bindings.Contract.IsDone(&_Bindings.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsSession) IsPaused() (bool, error) {
	return _Bindings.Contract.IsPaused(&_Bindings.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsCallerSession) IsPaused() (bool, error) {
	return _Bindings.Contract.IsPaused(&_Bindings.CallOpts)
}

// LpFeeP is a free data retrieval call binding the contract method 0xc68365df.
//
// Solidity: function lpFeeP() view returns(uint256)
func (_Bindings *BindingsCaller) LpFeeP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "lpFeeP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpFeeP is a free data retrieval call binding the contract method 0xc68365df.
//
// Solidity: function lpFeeP() view returns(uint256)
func (_Bindings *BindingsSession) LpFeeP() (*big.Int, error) {
	return _Bindings.Contract.LpFeeP(&_Bindings.CallOpts)
}

// LpFeeP is a free data retrieval call binding the contract method 0xc68365df.
//
// Solidity: function lpFeeP() view returns(uint256)
func (_Bindings *BindingsCallerSession) LpFeeP() (*big.Int, error) {
	return _Bindings.Contract.LpFeeP(&_Bindings.CallOpts)
}

// NftRewards is a free data retrieval call binding the contract method 0x5ad9f993.
//
// Solidity: function nftRewards() view returns(address)
func (_Bindings *BindingsCaller) NftRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "nftRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftRewards is a free data retrieval call binding the contract method 0x5ad9f993.
//
// Solidity: function nftRewards() view returns(address)
func (_Bindings *BindingsSession) NftRewards() (common.Address, error) {
	return _Bindings.Contract.NftRewards(&_Bindings.CallOpts)
}

// NftRewards is a free data retrieval call binding the contract method 0x5ad9f993.
//
// Solidity: function nftRewards() view returns(address)
func (_Bindings *BindingsCallerSession) NftRewards() (common.Address, error) {
	return _Bindings.Contract.NftRewards(&_Bindings.CallOpts)
}

// PairInfos is a free data retrieval call binding the contract method 0x1346b0ff.
//
// Solidity: function pairInfos() view returns(address)
func (_Bindings *BindingsCaller) PairInfos(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pairInfos")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairInfos is a free data retrieval call binding the contract method 0x1346b0ff.
//
// Solidity: function pairInfos() view returns(address)
func (_Bindings *BindingsSession) PairInfos() (common.Address, error) {
	return _Bindings.Contract.PairInfos(&_Bindings.CallOpts)
}

// PairInfos is a free data retrieval call binding the contract method 0x1346b0ff.
//
// Solidity: function pairInfos() view returns(address)
func (_Bindings *BindingsCallerSession) PairInfos() (common.Address, error) {
	return _Bindings.Contract.PairInfos(&_Bindings.CallOpts)
}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) PairMaxLeverage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pairMaxLeverage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) PairMaxLeverage(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PairMaxLeverage(&_Bindings.CallOpts, arg0)
}

// PairMaxLeverage is a free data retrieval call binding the contract method 0x281b693c.
//
// Solidity: function pairMaxLeverage(uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) PairMaxLeverage(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PairMaxLeverage(&_Bindings.CallOpts, arg0)
}

// Referrals is a free data retrieval call binding the contract method 0xd3dc7539.
//
// Solidity: function referrals() view returns(address)
func (_Bindings *BindingsCaller) Referrals(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "referrals")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Referrals is a free data retrieval call binding the contract method 0xd3dc7539.
//
// Solidity: function referrals() view returns(address)
func (_Bindings *BindingsSession) Referrals() (common.Address, error) {
	return _Bindings.Contract.Referrals(&_Bindings.CallOpts)
}

// Referrals is a free data retrieval call binding the contract method 0xd3dc7539.
//
// Solidity: function referrals() view returns(address)
func (_Bindings *BindingsCallerSession) Referrals() (common.Address, error) {
	return _Bindings.Contract.Referrals(&_Bindings.CallOpts)
}

// SssFeeP is a free data retrieval call binding the contract method 0x7bb6c11c.
//
// Solidity: function sssFeeP() view returns(uint256)
func (_Bindings *BindingsCaller) SssFeeP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "sssFeeP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SssFeeP is a free data retrieval call binding the contract method 0x7bb6c11c.
//
// Solidity: function sssFeeP() view returns(uint256)
func (_Bindings *BindingsSession) SssFeeP() (*big.Int, error) {
	return _Bindings.Contract.SssFeeP(&_Bindings.CallOpts)
}

// SssFeeP is a free data retrieval call binding the contract method 0x7bb6c11c.
//
// Solidity: function sssFeeP() view returns(uint256)
func (_Bindings *BindingsCallerSession) SssFeeP() (*big.Int, error) {
	return _Bindings.Contract.SssFeeP(&_Bindings.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Bindings *BindingsCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Bindings *BindingsSession) Staking() (common.Address, error) {
	return _Bindings.Contract.Staking(&_Bindings.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Bindings *BindingsCallerSession) Staking() (common.Address, error) {
	return _Bindings.Contract.Staking(&_Bindings.CallOpts)
}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_Bindings *BindingsCaller) StorageT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "storageT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_Bindings *BindingsSession) StorageT() (common.Address, error) {
	return _Bindings.Contract.StorageT(&_Bindings.CallOpts)
}

// StorageT is a free data retrieval call binding the contract method 0x16fff074.
//
// Solidity: function storageT() view returns(address)
func (_Bindings *BindingsCallerSession) StorageT() (common.Address, error) {
	return _Bindings.Contract.StorageT(&_Bindings.CallOpts)
}

// TradeData is a free data retrieval call binding the contract method 0x46f93a63.
//
// Solidity: function tradeData(address , uint256 , uint256 , uint8 ) view returns(uint40 maxSlippageP, uint216 _placeholder)
func (_Bindings *BindingsCaller) TradeData(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 uint8) (struct {
	MaxSlippageP *big.Int
	Placeholder  *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tradeData", arg0, arg1, arg2, arg3)

	outstruct := new(struct {
		MaxSlippageP *big.Int
		Placeholder  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxSlippageP = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Placeholder = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TradeData is a free data retrieval call binding the contract method 0x46f93a63.
//
// Solidity: function tradeData(address , uint256 , uint256 , uint8 ) view returns(uint40 maxSlippageP, uint216 _placeholder)
func (_Bindings *BindingsSession) TradeData(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 uint8) (struct {
	MaxSlippageP *big.Int
	Placeholder  *big.Int
}, error) {
	return _Bindings.Contract.TradeData(&_Bindings.CallOpts, arg0, arg1, arg2, arg3)
}

// TradeData is a free data retrieval call binding the contract method 0x46f93a63.
//
// Solidity: function tradeData(address , uint256 , uint256 , uint8 ) view returns(uint40 maxSlippageP, uint216 _placeholder)
func (_Bindings *BindingsCallerSession) TradeData(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 uint8) (struct {
	MaxSlippageP *big.Int
	Placeholder  *big.Int
}, error) {
	return _Bindings.Contract.TradeData(&_Bindings.CallOpts, arg0, arg1, arg2, arg3)
}

// TradeLastUpdated is a free data retrieval call binding the contract method 0x9bd24d83.
//
// Solidity: function tradeLastUpdated(address , uint256 , uint256 , uint8 ) view returns(uint32 tp, uint32 sl, uint32 limit, uint32 created)
func (_Bindings *BindingsCaller) TradeLastUpdated(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 uint8) (struct {
	Tp      uint32
	Sl      uint32
	Limit   uint32
	Created uint32
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tradeLastUpdated", arg0, arg1, arg2, arg3)

	outstruct := new(struct {
		Tp      uint32
		Sl      uint32
		Limit   uint32
		Created uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Sl = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Limit = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Created = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// TradeLastUpdated is a free data retrieval call binding the contract method 0x9bd24d83.
//
// Solidity: function tradeLastUpdated(address , uint256 , uint256 , uint8 ) view returns(uint32 tp, uint32 sl, uint32 limit, uint32 created)
func (_Bindings *BindingsSession) TradeLastUpdated(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 uint8) (struct {
	Tp      uint32
	Sl      uint32
	Limit   uint32
	Created uint32
}, error) {
	return _Bindings.Contract.TradeLastUpdated(&_Bindings.CallOpts, arg0, arg1, arg2, arg3)
}

// TradeLastUpdated is a free data retrieval call binding the contract method 0x9bd24d83.
//
// Solidity: function tradeLastUpdated(address , uint256 , uint256 , uint8 ) view returns(uint32 tp, uint32 sl, uint32 limit, uint32 created)
func (_Bindings *BindingsCallerSession) TradeLastUpdated(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 uint8) (struct {
	Tp      uint32
	Sl      uint32
	Limit   uint32
	Created uint32
}, error) {
	return _Bindings.Contract.TradeLastUpdated(&_Bindings.CallOpts, arg0, arg1, arg2, arg3)
}

// ClaimGovFees is a paid mutator transaction binding the contract method 0x0241e050.
//
// Solidity: function claimGovFees() returns()
func (_Bindings *BindingsTransactor) ClaimGovFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "claimGovFees")
}

// ClaimGovFees is a paid mutator transaction binding the contract method 0x0241e050.
//
// Solidity: function claimGovFees() returns()
func (_Bindings *BindingsSession) ClaimGovFees() (*types.Transaction, error) {
	return _Bindings.Contract.ClaimGovFees(&_Bindings.TransactOpts)
}

// ClaimGovFees is a paid mutator transaction binding the contract method 0x0241e050.
//
// Solidity: function claimGovFees() returns()
func (_Bindings *BindingsTransactorSession) ClaimGovFees() (*types.Transaction, error) {
	return _Bindings.Contract.ClaimGovFees(&_Bindings.TransactOpts)
}

// CloseTradeMarketCallback is a paid mutator transaction binding the contract method 0x7d490caf.
//
// Solidity: function closeTradeMarketCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactor) CloseTradeMarketCallback(opts *bind.TransactOpts, a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "closeTradeMarketCallback", a)
}

// CloseTradeMarketCallback is a paid mutator transaction binding the contract method 0x7d490caf.
//
// Solidity: function closeTradeMarketCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsSession) CloseTradeMarketCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.CloseTradeMarketCallback(&_Bindings.TransactOpts, a)
}

// CloseTradeMarketCallback is a paid mutator transaction binding the contract method 0x7d490caf.
//
// Solidity: function closeTradeMarketCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactorSession) CloseTradeMarketCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.CloseTradeMarketCallback(&_Bindings.TransactOpts, a)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_Bindings *BindingsTransactor) Done(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "done")
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_Bindings *BindingsSession) Done() (*types.Transaction, error) {
	return _Bindings.Contract.Done(&_Bindings.TransactOpts)
}

// Done is a paid mutator transaction binding the contract method 0xae8421e1.
//
// Solidity: function done() returns()
func (_Bindings *BindingsTransactorSession) Done() (*types.Transaction, error) {
	return _Bindings.Contract.Done(&_Bindings.TransactOpts)
}

// ExecuteNftCloseOrderCallback is a paid mutator transaction binding the contract method 0x23f7197b.
//
// Solidity: function executeNftCloseOrderCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactor) ExecuteNftCloseOrderCallback(opts *bind.TransactOpts, a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "executeNftCloseOrderCallback", a)
}

// ExecuteNftCloseOrderCallback is a paid mutator transaction binding the contract method 0x23f7197b.
//
// Solidity: function executeNftCloseOrderCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsSession) ExecuteNftCloseOrderCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteNftCloseOrderCallback(&_Bindings.TransactOpts, a)
}

// ExecuteNftCloseOrderCallback is a paid mutator transaction binding the contract method 0x23f7197b.
//
// Solidity: function executeNftCloseOrderCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactorSession) ExecuteNftCloseOrderCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteNftCloseOrderCallback(&_Bindings.TransactOpts, a)
}

// ExecuteNftOpenOrderCallback is a paid mutator transaction binding the contract method 0xb47ee408.
//
// Solidity: function executeNftOpenOrderCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactor) ExecuteNftOpenOrderCallback(opts *bind.TransactOpts, a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "executeNftOpenOrderCallback", a)
}

// ExecuteNftOpenOrderCallback is a paid mutator transaction binding the contract method 0xb47ee408.
//
// Solidity: function executeNftOpenOrderCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsSession) ExecuteNftOpenOrderCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteNftOpenOrderCallback(&_Bindings.TransactOpts, a)
}

// ExecuteNftOpenOrderCallback is a paid mutator transaction binding the contract method 0xb47ee408.
//
// Solidity: function executeNftOpenOrderCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactorSession) ExecuteNftOpenOrderCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteNftOpenOrderCallback(&_Bindings.TransactOpts, a)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address _storageT, address _nftRewards, address _pairInfos, address _referrals, address _staking, address vaultToApprove, uint256 _daiVaultFeeP, uint256 _lpFeeP, uint256 _sssFeeP, uint256 _canExecuteTimeout) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, _storageT common.Address, _nftRewards common.Address, _pairInfos common.Address, _referrals common.Address, _staking common.Address, vaultToApprove common.Address, _daiVaultFeeP *big.Int, _lpFeeP *big.Int, _sssFeeP *big.Int, _canExecuteTimeout *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", _storageT, _nftRewards, _pairInfos, _referrals, _staking, vaultToApprove, _daiVaultFeeP, _lpFeeP, _sssFeeP, _canExecuteTimeout)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address _storageT, address _nftRewards, address _pairInfos, address _referrals, address _staking, address vaultToApprove, uint256 _daiVaultFeeP, uint256 _lpFeeP, uint256 _sssFeeP, uint256 _canExecuteTimeout) returns()
func (_Bindings *BindingsSession) Initialize(_storageT common.Address, _nftRewards common.Address, _pairInfos common.Address, _referrals common.Address, _staking common.Address, vaultToApprove common.Address, _daiVaultFeeP *big.Int, _lpFeeP *big.Int, _sssFeeP *big.Int, _canExecuteTimeout *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, _storageT, _nftRewards, _pairInfos, _referrals, _staking, vaultToApprove, _daiVaultFeeP, _lpFeeP, _sssFeeP, _canExecuteTimeout)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address _storageT, address _nftRewards, address _pairInfos, address _referrals, address _staking, address vaultToApprove, uint256 _daiVaultFeeP, uint256 _lpFeeP, uint256 _sssFeeP, uint256 _canExecuteTimeout) returns()
func (_Bindings *BindingsTransactorSession) Initialize(_storageT common.Address, _nftRewards common.Address, _pairInfos common.Address, _referrals common.Address, _staking common.Address, vaultToApprove common.Address, _daiVaultFeeP *big.Int, _lpFeeP *big.Int, _sssFeeP *big.Int, _canExecuteTimeout *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, _storageT, _nftRewards, _pairInfos, _referrals, _staking, vaultToApprove, _daiVaultFeeP, _lpFeeP, _sssFeeP, _canExecuteTimeout)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address _borrowingFees) returns()
func (_Bindings *BindingsTransactor) InitializeV2(opts *bind.TransactOpts, _borrowingFees common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initializeV2", _borrowingFees)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address _borrowingFees) returns()
func (_Bindings *BindingsSession) InitializeV2(_borrowingFees common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.InitializeV2(&_Bindings.TransactOpts, _borrowingFees)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x29b6eca9.
//
// Solidity: function initializeV2(address _borrowingFees) returns()
func (_Bindings *BindingsTransactorSession) InitializeV2(_borrowingFees common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.InitializeV2(&_Bindings.TransactOpts, _borrowingFees)
}

// InitializeV4 is a paid mutator transaction binding the contract method 0x3a3ec5b0.
//
// Solidity: function initializeV4(address _staking, address _oracleRewards) returns()
func (_Bindings *BindingsTransactor) InitializeV4(opts *bind.TransactOpts, _staking common.Address, _oracleRewards common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initializeV4", _staking, _oracleRewards)
}

// InitializeV4 is a paid mutator transaction binding the contract method 0x3a3ec5b0.
//
// Solidity: function initializeV4(address _staking, address _oracleRewards) returns()
func (_Bindings *BindingsSession) InitializeV4(_staking common.Address, _oracleRewards common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.InitializeV4(&_Bindings.TransactOpts, _staking, _oracleRewards)
}

// InitializeV4 is a paid mutator transaction binding the contract method 0x3a3ec5b0.
//
// Solidity: function initializeV4(address _staking, address _oracleRewards) returns()
func (_Bindings *BindingsTransactorSession) InitializeV4(_staking common.Address, _oracleRewards common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.InitializeV4(&_Bindings.TransactOpts, _staking, _oracleRewards)
}

// OpenTradeMarketCallback is a paid mutator transaction binding the contract method 0x61ba6c3d.
//
// Solidity: function openTradeMarketCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactor) OpenTradeMarketCallback(opts *bind.TransactOpts, a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "openTradeMarketCallback", a)
}

// OpenTradeMarketCallback is a paid mutator transaction binding the contract method 0x61ba6c3d.
//
// Solidity: function openTradeMarketCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsSession) OpenTradeMarketCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.OpenTradeMarketCallback(&_Bindings.TransactOpts, a)
}

// OpenTradeMarketCallback is a paid mutator transaction binding the contract method 0x61ba6c3d.
//
// Solidity: function openTradeMarketCallback((uint256,uint256,uint256,uint256,uint256,uint256) a) returns()
func (_Bindings *BindingsTransactorSession) OpenTradeMarketCallback(a GNSTradingCallbacksV641AggregatorAnswer) (*types.Transaction, error) {
	return _Bindings.Contract.OpenTradeMarketCallback(&_Bindings.TransactOpts, a)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsSession) Pause() (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsTransactorSession) Pause() (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts)
}

// SetClosingFeeSharesP is a paid mutator transaction binding the contract method 0xc68078b9.
//
// Solidity: function setClosingFeeSharesP(uint256 _daiVaultFeeP, uint256 _lpFeeP, uint256 _sssFeeP) returns()
func (_Bindings *BindingsTransactor) SetClosingFeeSharesP(opts *bind.TransactOpts, _daiVaultFeeP *big.Int, _lpFeeP *big.Int, _sssFeeP *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setClosingFeeSharesP", _daiVaultFeeP, _lpFeeP, _sssFeeP)
}

// SetClosingFeeSharesP is a paid mutator transaction binding the contract method 0xc68078b9.
//
// Solidity: function setClosingFeeSharesP(uint256 _daiVaultFeeP, uint256 _lpFeeP, uint256 _sssFeeP) returns()
func (_Bindings *BindingsSession) SetClosingFeeSharesP(_daiVaultFeeP *big.Int, _lpFeeP *big.Int, _sssFeeP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetClosingFeeSharesP(&_Bindings.TransactOpts, _daiVaultFeeP, _lpFeeP, _sssFeeP)
}

// SetClosingFeeSharesP is a paid mutator transaction binding the contract method 0xc68078b9.
//
// Solidity: function setClosingFeeSharesP(uint256 _daiVaultFeeP, uint256 _lpFeeP, uint256 _sssFeeP) returns()
func (_Bindings *BindingsTransactorSession) SetClosingFeeSharesP(_daiVaultFeeP *big.Int, _lpFeeP *big.Int, _sssFeeP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetClosingFeeSharesP(&_Bindings.TransactOpts, _daiVaultFeeP, _lpFeeP, _sssFeeP)
}

// SetPairMaxLeverage is a paid mutator transaction binding the contract method 0x48829122.
//
// Solidity: function setPairMaxLeverage(uint256 pairIndex, uint256 maxLeverage) returns()
func (_Bindings *BindingsTransactor) SetPairMaxLeverage(opts *bind.TransactOpts, pairIndex *big.Int, maxLeverage *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPairMaxLeverage", pairIndex, maxLeverage)
}

// SetPairMaxLeverage is a paid mutator transaction binding the contract method 0x48829122.
//
// Solidity: function setPairMaxLeverage(uint256 pairIndex, uint256 maxLeverage) returns()
func (_Bindings *BindingsSession) SetPairMaxLeverage(pairIndex *big.Int, maxLeverage *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPairMaxLeverage(&_Bindings.TransactOpts, pairIndex, maxLeverage)
}

// SetPairMaxLeverage is a paid mutator transaction binding the contract method 0x48829122.
//
// Solidity: function setPairMaxLeverage(uint256 pairIndex, uint256 maxLeverage) returns()
func (_Bindings *BindingsTransactorSession) SetPairMaxLeverage(pairIndex *big.Int, maxLeverage *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPairMaxLeverage(&_Bindings.TransactOpts, pairIndex, maxLeverage)
}

// SetPairMaxLeverageArray is a paid mutator transaction binding the contract method 0x00e0fa89.
//
// Solidity: function setPairMaxLeverageArray(uint256[] indices, uint256[] values) returns()
func (_Bindings *BindingsTransactor) SetPairMaxLeverageArray(opts *bind.TransactOpts, indices []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPairMaxLeverageArray", indices, values)
}

// SetPairMaxLeverageArray is a paid mutator transaction binding the contract method 0x00e0fa89.
//
// Solidity: function setPairMaxLeverageArray(uint256[] indices, uint256[] values) returns()
func (_Bindings *BindingsSession) SetPairMaxLeverageArray(indices []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPairMaxLeverageArray(&_Bindings.TransactOpts, indices, values)
}

// SetPairMaxLeverageArray is a paid mutator transaction binding the contract method 0x00e0fa89.
//
// Solidity: function setPairMaxLeverageArray(uint256[] indices, uint256[] values) returns()
func (_Bindings *BindingsTransactorSession) SetPairMaxLeverageArray(indices []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPairMaxLeverageArray(&_Bindings.TransactOpts, indices, values)
}

// SetTradeData is a paid mutator transaction binding the contract method 0x21c86196.
//
// Solidity: function setTradeData((address,uint256,uint256,uint8) _id, (uint40,uint216) _tradeData) returns()
func (_Bindings *BindingsTransactor) SetTradeData(opts *bind.TransactOpts, _id GNSTradingCallbacksV641SimplifiedTradeId, _tradeData GNSTradingCallbacksV641TradeData) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setTradeData", _id, _tradeData)
}

// SetTradeData is a paid mutator transaction binding the contract method 0x21c86196.
//
// Solidity: function setTradeData((address,uint256,uint256,uint8) _id, (uint40,uint216) _tradeData) returns()
func (_Bindings *BindingsSession) SetTradeData(_id GNSTradingCallbacksV641SimplifiedTradeId, _tradeData GNSTradingCallbacksV641TradeData) (*types.Transaction, error) {
	return _Bindings.Contract.SetTradeData(&_Bindings.TransactOpts, _id, _tradeData)
}

// SetTradeData is a paid mutator transaction binding the contract method 0x21c86196.
//
// Solidity: function setTradeData((address,uint256,uint256,uint8) _id, (uint40,uint216) _tradeData) returns()
func (_Bindings *BindingsTransactorSession) SetTradeData(_id GNSTradingCallbacksV641SimplifiedTradeId, _tradeData GNSTradingCallbacksV641TradeData) (*types.Transaction, error) {
	return _Bindings.Contract.SetTradeData(&_Bindings.TransactOpts, _id, _tradeData)
}

// SetTradeLastUpdated is a paid mutator transaction binding the contract method 0x7e09ae8a.
//
// Solidity: function setTradeLastUpdated((address,uint256,uint256,uint8) _id, (uint32,uint32,uint32,uint32) _lastUpdated) returns()
func (_Bindings *BindingsTransactor) SetTradeLastUpdated(opts *bind.TransactOpts, _id GNSTradingCallbacksV641SimplifiedTradeId, _lastUpdated GNSTradingCallbacksV641LastUpdated) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setTradeLastUpdated", _id, _lastUpdated)
}

// SetTradeLastUpdated is a paid mutator transaction binding the contract method 0x7e09ae8a.
//
// Solidity: function setTradeLastUpdated((address,uint256,uint256,uint8) _id, (uint32,uint32,uint32,uint32) _lastUpdated) returns()
func (_Bindings *BindingsSession) SetTradeLastUpdated(_id GNSTradingCallbacksV641SimplifiedTradeId, _lastUpdated GNSTradingCallbacksV641LastUpdated) (*types.Transaction, error) {
	return _Bindings.Contract.SetTradeLastUpdated(&_Bindings.TransactOpts, _id, _lastUpdated)
}

// SetTradeLastUpdated is a paid mutator transaction binding the contract method 0x7e09ae8a.
//
// Solidity: function setTradeLastUpdated((address,uint256,uint256,uint8) _id, (uint32,uint32,uint32,uint32) _lastUpdated) returns()
func (_Bindings *BindingsTransactorSession) SetTradeLastUpdated(_id GNSTradingCallbacksV641SimplifiedTradeId, _lastUpdated GNSTradingCallbacksV641LastUpdated) (*types.Transaction, error) {
	return _Bindings.Contract.SetTradeLastUpdated(&_Bindings.TransactOpts, _id, _lastUpdated)
}

// BindingsBorrowingFeeChargedIterator is returned from FilterBorrowingFeeCharged and is used to iterate over the raw logs and unpacked data for BorrowingFeeCharged events raised by the Bindings contract.
type BindingsBorrowingFeeChargedIterator struct {
	Event *BindingsBorrowingFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsBorrowingFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBorrowingFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsBorrowingFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsBorrowingFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBorrowingFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBorrowingFeeCharged represents a BorrowingFeeCharged event raised by the Bindings contract.
type BindingsBorrowingFeeCharged struct {
	Trader        common.Address
	TradeValueDai *big.Int
	FeeValueDai   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrowingFeeCharged is a free log retrieval operation binding the contract event 0xe7d34775bf6fd7b34e703a903ef79ab16166ebdffce96a66f4d2f84b6263bb29.
//
// Solidity: event BorrowingFeeCharged(address indexed trader, uint256 tradeValueDai, uint256 feeValueDai)
func (_Bindings *BindingsFilterer) FilterBorrowingFeeCharged(opts *bind.FilterOpts, trader []common.Address) (*BindingsBorrowingFeeChargedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BorrowingFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBorrowingFeeChargedIterator{contract: _Bindings.contract, event: "BorrowingFeeCharged", logs: logs, sub: sub}, nil
}

// WatchBorrowingFeeCharged is a free log subscription operation binding the contract event 0xe7d34775bf6fd7b34e703a903ef79ab16166ebdffce96a66f4d2f84b6263bb29.
//
// Solidity: event BorrowingFeeCharged(address indexed trader, uint256 tradeValueDai, uint256 feeValueDai)
func (_Bindings *BindingsFilterer) WatchBorrowingFeeCharged(opts *bind.WatchOpts, sink chan<- *BindingsBorrowingFeeCharged, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BorrowingFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBorrowingFeeCharged)
				if err := _Bindings.contract.UnpackLog(event, "BorrowingFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrowingFeeCharged is a log parse operation binding the contract event 0xe7d34775bf6fd7b34e703a903ef79ab16166ebdffce96a66f4d2f84b6263bb29.
//
// Solidity: event BorrowingFeeCharged(address indexed trader, uint256 tradeValueDai, uint256 feeValueDai)
func (_Bindings *BindingsFilterer) ParseBorrowingFeeCharged(log types.Log) (*BindingsBorrowingFeeCharged, error) {
	event := new(BindingsBorrowingFeeCharged)
	if err := _Bindings.contract.UnpackLog(event, "BorrowingFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsClosingFeeSharesPUpdatedIterator is returned from FilterClosingFeeSharesPUpdated and is used to iterate over the raw logs and unpacked data for ClosingFeeSharesPUpdated events raised by the Bindings contract.
type BindingsClosingFeeSharesPUpdatedIterator struct {
	Event *BindingsClosingFeeSharesPUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsClosingFeeSharesPUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsClosingFeeSharesPUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsClosingFeeSharesPUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsClosingFeeSharesPUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsClosingFeeSharesPUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsClosingFeeSharesPUpdated represents a ClosingFeeSharesPUpdated event raised by the Bindings contract.
type BindingsClosingFeeSharesPUpdated struct {
	DaiVaultFeeP *big.Int
	LpFeeP       *big.Int
	SssFeeP      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClosingFeeSharesPUpdated is a free log retrieval operation binding the contract event 0x0caa98ed9a1605da290817d1f67b1b83c63f9229abeca5123df5d90581c49558.
//
// Solidity: event ClosingFeeSharesPUpdated(uint256 daiVaultFeeP, uint256 lpFeeP, uint256 sssFeeP)
func (_Bindings *BindingsFilterer) FilterClosingFeeSharesPUpdated(opts *bind.FilterOpts) (*BindingsClosingFeeSharesPUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ClosingFeeSharesPUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsClosingFeeSharesPUpdatedIterator{contract: _Bindings.contract, event: "ClosingFeeSharesPUpdated", logs: logs, sub: sub}, nil
}

// WatchClosingFeeSharesPUpdated is a free log subscription operation binding the contract event 0x0caa98ed9a1605da290817d1f67b1b83c63f9229abeca5123df5d90581c49558.
//
// Solidity: event ClosingFeeSharesPUpdated(uint256 daiVaultFeeP, uint256 lpFeeP, uint256 sssFeeP)
func (_Bindings *BindingsFilterer) WatchClosingFeeSharesPUpdated(opts *bind.WatchOpts, sink chan<- *BindingsClosingFeeSharesPUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ClosingFeeSharesPUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsClosingFeeSharesPUpdated)
				if err := _Bindings.contract.UnpackLog(event, "ClosingFeeSharesPUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosingFeeSharesPUpdated is a log parse operation binding the contract event 0x0caa98ed9a1605da290817d1f67b1b83c63f9229abeca5123df5d90581c49558.
//
// Solidity: event ClosingFeeSharesPUpdated(uint256 daiVaultFeeP, uint256 lpFeeP, uint256 sssFeeP)
func (_Bindings *BindingsFilterer) ParseClosingFeeSharesPUpdated(log types.Log) (*BindingsClosingFeeSharesPUpdated, error) {
	event := new(BindingsClosingFeeSharesPUpdated)
	if err := _Bindings.contract.UnpackLog(event, "ClosingFeeSharesPUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDaiVaultFeeChargedIterator is returned from FilterDaiVaultFeeCharged and is used to iterate over the raw logs and unpacked data for DaiVaultFeeCharged events raised by the Bindings contract.
type BindingsDaiVaultFeeChargedIterator struct {
	Event *BindingsDaiVaultFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsDaiVaultFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDaiVaultFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsDaiVaultFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsDaiVaultFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDaiVaultFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDaiVaultFeeCharged represents a DaiVaultFeeCharged event raised by the Bindings contract.
type BindingsDaiVaultFeeCharged struct {
	Trader   common.Address
	ValueDai *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDaiVaultFeeCharged is a free log retrieval operation binding the contract event 0x60c73da98faf96842eabd77a0c73964cd189dbaf2c9ae90923a3fed137f30e3e.
//
// Solidity: event DaiVaultFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) FilterDaiVaultFeeCharged(opts *bind.FilterOpts, trader []common.Address) (*BindingsDaiVaultFeeChargedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DaiVaultFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsDaiVaultFeeChargedIterator{contract: _Bindings.contract, event: "DaiVaultFeeCharged", logs: logs, sub: sub}, nil
}

// WatchDaiVaultFeeCharged is a free log subscription operation binding the contract event 0x60c73da98faf96842eabd77a0c73964cd189dbaf2c9ae90923a3fed137f30e3e.
//
// Solidity: event DaiVaultFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) WatchDaiVaultFeeCharged(opts *bind.WatchOpts, sink chan<- *BindingsDaiVaultFeeCharged, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DaiVaultFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDaiVaultFeeCharged)
				if err := _Bindings.contract.UnpackLog(event, "DaiVaultFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDaiVaultFeeCharged is a log parse operation binding the contract event 0x60c73da98faf96842eabd77a0c73964cd189dbaf2c9ae90923a3fed137f30e3e.
//
// Solidity: event DaiVaultFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) ParseDaiVaultFeeCharged(log types.Log) (*BindingsDaiVaultFeeCharged, error) {
	event := new(BindingsDaiVaultFeeCharged)
	if err := _Bindings.contract.UnpackLog(event, "DaiVaultFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDoneIterator is returned from FilterDone and is used to iterate over the raw logs and unpacked data for Done events raised by the Bindings contract.
type BindingsDoneIterator struct {
	Event *BindingsDone // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDone)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsDone)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDone represents a Done event raised by the Bindings contract.
type BindingsDone struct {
	Done bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDone is a free log retrieval operation binding the contract event 0xc3a6f986261de9467c2838c6df8ef74f9107855159205600c0bc7a14cdfd3888.
//
// Solidity: event Done(bool done)
func (_Bindings *BindingsFilterer) FilterDone(opts *bind.FilterOpts) (*BindingsDoneIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Done")
	if err != nil {
		return nil, err
	}
	return &BindingsDoneIterator{contract: _Bindings.contract, event: "Done", logs: logs, sub: sub}, nil
}

// WatchDone is a free log subscription operation binding the contract event 0xc3a6f986261de9467c2838c6df8ef74f9107855159205600c0bc7a14cdfd3888.
//
// Solidity: event Done(bool done)
func (_Bindings *BindingsFilterer) WatchDone(opts *bind.WatchOpts, sink chan<- *BindingsDone) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Done")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDone)
				if err := _Bindings.contract.UnpackLog(event, "Done", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDone is a log parse operation binding the contract event 0xc3a6f986261de9467c2838c6df8ef74f9107855159205600c0bc7a14cdfd3888.
//
// Solidity: event Done(bool done)
func (_Bindings *BindingsFilterer) ParseDone(log types.Log) (*BindingsDone, error) {
	event := new(BindingsDone)
	if err := _Bindings.contract.UnpackLog(event, "Done", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsGovFeeChargedIterator is returned from FilterGovFeeCharged and is used to iterate over the raw logs and unpacked data for GovFeeCharged events raised by the Bindings contract.
type BindingsGovFeeChargedIterator struct {
	Event *BindingsGovFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsGovFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsGovFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsGovFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsGovFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsGovFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsGovFeeCharged represents a GovFeeCharged event raised by the Bindings contract.
type BindingsGovFeeCharged struct {
	Trader      common.Address
	ValueDai    *big.Int
	Distributed bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovFeeCharged is a free log retrieval operation binding the contract event 0xccd80d359a6fbe0bfa5cbb1ecf0854adbe8c67b4ed6bf10d3c0d78c2be0f48cb.
//
// Solidity: event GovFeeCharged(address indexed trader, uint256 valueDai, bool distributed)
func (_Bindings *BindingsFilterer) FilterGovFeeCharged(opts *bind.FilterOpts, trader []common.Address) (*BindingsGovFeeChargedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "GovFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsGovFeeChargedIterator{contract: _Bindings.contract, event: "GovFeeCharged", logs: logs, sub: sub}, nil
}

// WatchGovFeeCharged is a free log subscription operation binding the contract event 0xccd80d359a6fbe0bfa5cbb1ecf0854adbe8c67b4ed6bf10d3c0d78c2be0f48cb.
//
// Solidity: event GovFeeCharged(address indexed trader, uint256 valueDai, bool distributed)
func (_Bindings *BindingsFilterer) WatchGovFeeCharged(opts *bind.WatchOpts, sink chan<- *BindingsGovFeeCharged, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "GovFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsGovFeeCharged)
				if err := _Bindings.contract.UnpackLog(event, "GovFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovFeeCharged is a log parse operation binding the contract event 0xccd80d359a6fbe0bfa5cbb1ecf0854adbe8c67b4ed6bf10d3c0d78c2be0f48cb.
//
// Solidity: event GovFeeCharged(address indexed trader, uint256 valueDai, bool distributed)
func (_Bindings *BindingsFilterer) ParseGovFeeCharged(log types.Log) (*BindingsGovFeeCharged, error) {
	event := new(BindingsGovFeeCharged)
	if err := _Bindings.contract.UnpackLog(event, "GovFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsGovFeesClaimedIterator is returned from FilterGovFeesClaimed and is used to iterate over the raw logs and unpacked data for GovFeesClaimed events raised by the Bindings contract.
type BindingsGovFeesClaimedIterator struct {
	Event *BindingsGovFeesClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsGovFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsGovFeesClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsGovFeesClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsGovFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsGovFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsGovFeesClaimed represents a GovFeesClaimed event raised by the Bindings contract.
type BindingsGovFeesClaimed struct {
	ValueDai *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGovFeesClaimed is a free log retrieval operation binding the contract event 0x39b06677afbdb5c3b9934c9ce55728be609a055c398ecd957e8d19a5d3d80a5b.
//
// Solidity: event GovFeesClaimed(uint256 valueDai)
func (_Bindings *BindingsFilterer) FilterGovFeesClaimed(opts *bind.FilterOpts) (*BindingsGovFeesClaimedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "GovFeesClaimed")
	if err != nil {
		return nil, err
	}
	return &BindingsGovFeesClaimedIterator{contract: _Bindings.contract, event: "GovFeesClaimed", logs: logs, sub: sub}, nil
}

// WatchGovFeesClaimed is a free log subscription operation binding the contract event 0x39b06677afbdb5c3b9934c9ce55728be609a055c398ecd957e8d19a5d3d80a5b.
//
// Solidity: event GovFeesClaimed(uint256 valueDai)
func (_Bindings *BindingsFilterer) WatchGovFeesClaimed(opts *bind.WatchOpts, sink chan<- *BindingsGovFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "GovFeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsGovFeesClaimed)
				if err := _Bindings.contract.UnpackLog(event, "GovFeesClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovFeesClaimed is a log parse operation binding the contract event 0x39b06677afbdb5c3b9934c9ce55728be609a055c398ecd957e8d19a5d3d80a5b.
//
// Solidity: event GovFeesClaimed(uint256 valueDai)
func (_Bindings *BindingsFilterer) ParseGovFeesClaimed(log types.Log) (*BindingsGovFeesClaimed, error) {
	event := new(BindingsGovFeesClaimed)
	if err := _Bindings.contract.UnpackLog(event, "GovFeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsLimitExecutedIterator is returned from FilterLimitExecuted and is used to iterate over the raw logs and unpacked data for LimitExecuted events raised by the Bindings contract.
type BindingsLimitExecutedIterator struct {
	Event *BindingsLimitExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsLimitExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsLimitExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsLimitExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsLimitExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsLimitExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsLimitExecuted represents a LimitExecuted event raised by the Bindings contract.
type BindingsLimitExecuted struct {
	OrderId         *big.Int
	LimitIndex *big.Int
	T          StorageInterfaceV5Trade
	NftHolder  common.Address
	OrderType       uint8
	Price           *big.Int
	PriceImpactP    *big.Int
	PositionSizeDai *big.Int
	PercentProfit   *big.Int
	DaiSentToTrader *big.Int
	ExactExecution  bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLimitExecuted is a free log retrieval operation binding the contract event 0x1ab0771256522e5114b583b488c490436d6f8fe02b1e1c9697443e8704c4e840.
//
// Solidity: event LimitExecuted(uint256 indexed orderId, uint256 limitIndex, (address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, address indexed nftHolder, uint8 orderType, uint256 price, uint256 priceImpactP, uint256 positionSizeDai, int256 percentProfit, uint256 daiSentToTrader, bool exactExecution)
func (_Bindings *BindingsFilterer) FilterLimitExecuted(opts *bind.FilterOpts, orderId []*big.Int, nftHolder []common.Address) (*BindingsLimitExecutedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	var nftHolderRule []interface{}
	for _, nftHolderItem := range nftHolder {
		nftHolderRule = append(nftHolderRule, nftHolderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "LimitExecuted", orderIdRule, nftHolderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsLimitExecutedIterator{contract: _Bindings.contract, event: "LimitExecuted", logs: logs, sub: sub}, nil
}

// WatchLimitExecuted is a free log subscription operation binding the contract event 0x1ab0771256522e5114b583b488c490436d6f8fe02b1e1c9697443e8704c4e840.
//
// Solidity: event LimitExecuted(uint256 indexed orderId, uint256 limitIndex, (address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, address indexed nftHolder, uint8 orderType, uint256 price, uint256 priceImpactP, uint256 positionSizeDai, int256 percentProfit, uint256 daiSentToTrader, bool exactExecution)
func (_Bindings *BindingsFilterer) WatchLimitExecuted(opts *bind.WatchOpts, sink chan<- *BindingsLimitExecuted, orderId []*big.Int, nftHolder []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	var nftHolderRule []interface{}
	for _, nftHolderItem := range nftHolder {
		nftHolderRule = append(nftHolderRule, nftHolderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "LimitExecuted", orderIdRule, nftHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsLimitExecuted)
				if err := _Bindings.contract.UnpackLog(event, "LimitExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLimitExecuted is a log parse operation binding the contract event 0x1ab0771256522e5114b583b488c490436d6f8fe02b1e1c9697443e8704c4e840.
//
// Solidity: event LimitExecuted(uint256 indexed orderId, uint256 limitIndex, (address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, address indexed nftHolder, uint8 orderType, uint256 price, uint256 priceImpactP, uint256 positionSizeDai, int256 percentProfit, uint256 daiSentToTrader, bool exactExecution)
func (_Bindings *BindingsFilterer) ParseLimitExecuted(log types.Log) (*BindingsLimitExecuted, error) {
	event := new(BindingsLimitExecuted)
	if err := _Bindings.contract.UnpackLog(event, "LimitExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMarketCloseCanceledIterator is returned from FilterMarketCloseCanceled and is used to iterate over the raw logs and unpacked data for MarketCloseCanceled events raised by the Bindings contract.
type BindingsMarketCloseCanceledIterator struct {
	Event *BindingsMarketCloseCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsMarketCloseCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMarketCloseCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsMarketCloseCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsMarketCloseCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMarketCloseCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMarketCloseCanceled represents a MarketCloseCanceled event raised by the Bindings contract.
type BindingsMarketCloseCanceled struct {
	OrderId      *big.Int
	Trader       common.Address
	PairIndex    *big.Int
	Index        *big.Int
	CancelReason uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMarketCloseCanceled is a free log retrieval operation binding the contract event 0x293df767d6749666902026d2f6a2cc4e5f15cdede46402226c42ef4fdf27a17c.
//
// Solidity: event MarketCloseCanceled(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, uint256 index, uint8 cancelReason)
func (_Bindings *BindingsFilterer) FilterMarketCloseCanceled(opts *bind.FilterOpts, orderId []*big.Int, trader []common.Address, pairIndex []*big.Int) (*BindingsMarketCloseCanceledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MarketCloseCanceled", orderIdRule, traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMarketCloseCanceledIterator{contract: _Bindings.contract, event: "MarketCloseCanceled", logs: logs, sub: sub}, nil
}

// WatchMarketCloseCanceled is a free log subscription operation binding the contract event 0x293df767d6749666902026d2f6a2cc4e5f15cdede46402226c42ef4fdf27a17c.
//
// Solidity: event MarketCloseCanceled(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, uint256 index, uint8 cancelReason)
func (_Bindings *BindingsFilterer) WatchMarketCloseCanceled(opts *bind.WatchOpts, sink chan<- *BindingsMarketCloseCanceled, orderId []*big.Int, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MarketCloseCanceled", orderIdRule, traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMarketCloseCanceled)
				if err := _Bindings.contract.UnpackLog(event, "MarketCloseCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketCloseCanceled is a log parse operation binding the contract event 0x293df767d6749666902026d2f6a2cc4e5f15cdede46402226c42ef4fdf27a17c.
//
// Solidity: event MarketCloseCanceled(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, uint256 index, uint8 cancelReason)
func (_Bindings *BindingsFilterer) ParseMarketCloseCanceled(log types.Log) (*BindingsMarketCloseCanceled, error) {
	event := new(BindingsMarketCloseCanceled)
	if err := _Bindings.contract.UnpackLog(event, "MarketCloseCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMarketExecutedIterator is returned from FilterMarketExecuted and is used to iterate over the raw logs and unpacked data for MarketExecuted events raised by the Bindings contract.
type BindingsMarketExecutedIterator struct {
	Event *BindingsMarketExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsMarketExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMarketExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsMarketExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsMarketExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMarketExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMarketExecuted represents a MarketExecuted event raised by the Bindings contract.
type BindingsMarketExecuted struct {
	OrderId *big.Int
	T       StorageInterfaceV5Trade
	Open    bool
	Price           *big.Int
	PriceImpactP    *big.Int
	PositionSizeDai *big.Int
	PercentProfit   *big.Int
	DaiSentToTrader *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMarketExecuted is a free log retrieval operation binding the contract event 0x2739a12dffae5d66bd9e126a286078ed771840f2288f0afa5709ce38c3330997.
//
// Solidity: event MarketExecuted(uint256 indexed orderId, (address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, bool open, uint256 price, uint256 priceImpactP, uint256 positionSizeDai, int256 percentProfit, uint256 daiSentToTrader)
func (_Bindings *BindingsFilterer) FilterMarketExecuted(opts *bind.FilterOpts, orderId []*big.Int) (*BindingsMarketExecutedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MarketExecuted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMarketExecutedIterator{contract: _Bindings.contract, event: "MarketExecuted", logs: logs, sub: sub}, nil
}

// WatchMarketExecuted is a free log subscription operation binding the contract event 0x2739a12dffae5d66bd9e126a286078ed771840f2288f0afa5709ce38c3330997.
//
// Solidity: event MarketExecuted(uint256 indexed orderId, (address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, bool open, uint256 price, uint256 priceImpactP, uint256 positionSizeDai, int256 percentProfit, uint256 daiSentToTrader)
func (_Bindings *BindingsFilterer) WatchMarketExecuted(opts *bind.WatchOpts, sink chan<- *BindingsMarketExecuted, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MarketExecuted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMarketExecuted)
				if err := _Bindings.contract.UnpackLog(event, "MarketExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketExecuted is a log parse operation binding the contract event 0x2739a12dffae5d66bd9e126a286078ed771840f2288f0afa5709ce38c3330997.
//
// Solidity: event MarketExecuted(uint256 indexed orderId, (address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, bool open, uint256 price, uint256 priceImpactP, uint256 positionSizeDai, int256 percentProfit, uint256 daiSentToTrader)
func (_Bindings *BindingsFilterer) ParseMarketExecuted(log types.Log) (*BindingsMarketExecuted, error) {
	event := new(BindingsMarketExecuted)
	if err := _Bindings.contract.UnpackLog(event, "MarketExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMarketOpenCanceledIterator is returned from FilterMarketOpenCanceled and is used to iterate over the raw logs and unpacked data for MarketOpenCanceled events raised by the Bindings contract.
type BindingsMarketOpenCanceledIterator struct {
	Event *BindingsMarketOpenCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsMarketOpenCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMarketOpenCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsMarketOpenCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsMarketOpenCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMarketOpenCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMarketOpenCanceled represents a MarketOpenCanceled event raised by the Bindings contract.
type BindingsMarketOpenCanceled struct {
	OrderId      *big.Int
	Trader       common.Address
	PairIndex    *big.Int
	CancelReason uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMarketOpenCanceled is a free log retrieval operation binding the contract event 0x1dc3532663e5566091476fb5aba1e514ef733714c83d4feec5723de6f16c3269.
//
// Solidity: event MarketOpenCanceled(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, uint8 cancelReason)
func (_Bindings *BindingsFilterer) FilterMarketOpenCanceled(opts *bind.FilterOpts, orderId []*big.Int, trader []common.Address, pairIndex []*big.Int) (*BindingsMarketOpenCanceledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MarketOpenCanceled", orderIdRule, traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMarketOpenCanceledIterator{contract: _Bindings.contract, event: "MarketOpenCanceled", logs: logs, sub: sub}, nil
}

// WatchMarketOpenCanceled is a free log subscription operation binding the contract event 0x1dc3532663e5566091476fb5aba1e514ef733714c83d4feec5723de6f16c3269.
//
// Solidity: event MarketOpenCanceled(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, uint8 cancelReason)
func (_Bindings *BindingsFilterer) WatchMarketOpenCanceled(opts *bind.WatchOpts, sink chan<- *BindingsMarketOpenCanceled, orderId []*big.Int, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MarketOpenCanceled", orderIdRule, traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMarketOpenCanceled)
				if err := _Bindings.contract.UnpackLog(event, "MarketOpenCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketOpenCanceled is a log parse operation binding the contract event 0x1dc3532663e5566091476fb5aba1e514ef733714c83d4feec5723de6f16c3269.
//
// Solidity: event MarketOpenCanceled(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, uint8 cancelReason)
func (_Bindings *BindingsFilterer) ParseMarketOpenCanceled(log types.Log) (*BindingsMarketOpenCanceled, error) {
	event := new(BindingsMarketOpenCanceled)
	if err := _Bindings.contract.UnpackLog(event, "MarketOpenCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsNftOrderCanceledIterator is returned from FilterNftOrderCanceled and is used to iterate over the raw logs and unpacked data for NftOrderCanceled events raised by the Bindings contract.
type BindingsNftOrderCanceledIterator struct {
	Event *BindingsNftOrderCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsNftOrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNftOrderCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsNftOrderCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsNftOrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNftOrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNftOrderCanceled represents a NftOrderCanceled event raised by the Bindings contract.
type BindingsNftOrderCanceled struct {
	OrderId      *big.Int
	NftHolder    common.Address
	OrderType    uint8
	CancelReason uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNftOrderCanceled is a free log retrieval operation binding the contract event 0xe9681b5336d843735c62e93114e5a0f45912a84ae83fa3f3ed80ca5ad933dfc3.
//
// Solidity: event NftOrderCanceled(uint256 indexed orderId, address indexed nftHolder, uint8 orderType, uint8 cancelReason)
func (_Bindings *BindingsFilterer) FilterNftOrderCanceled(opts *bind.FilterOpts, orderId []*big.Int, nftHolder []common.Address) (*BindingsNftOrderCanceledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nftHolderRule []interface{}
	for _, nftHolderItem := range nftHolder {
		nftHolderRule = append(nftHolderRule, nftHolderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NftOrderCanceled", orderIdRule, nftHolderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsNftOrderCanceledIterator{contract: _Bindings.contract, event: "NftOrderCanceled", logs: logs, sub: sub}, nil
}

// WatchNftOrderCanceled is a free log subscription operation binding the contract event 0xe9681b5336d843735c62e93114e5a0f45912a84ae83fa3f3ed80ca5ad933dfc3.
//
// Solidity: event NftOrderCanceled(uint256 indexed orderId, address indexed nftHolder, uint8 orderType, uint8 cancelReason)
func (_Bindings *BindingsFilterer) WatchNftOrderCanceled(opts *bind.WatchOpts, sink chan<- *BindingsNftOrderCanceled, orderId []*big.Int, nftHolder []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nftHolderRule []interface{}
	for _, nftHolderItem := range nftHolder {
		nftHolderRule = append(nftHolderRule, nftHolderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NftOrderCanceled", orderIdRule, nftHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNftOrderCanceled)
				if err := _Bindings.contract.UnpackLog(event, "NftOrderCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNftOrderCanceled is a log parse operation binding the contract event 0xe9681b5336d843735c62e93114e5a0f45912a84ae83fa3f3ed80ca5ad933dfc3.
//
// Solidity: event NftOrderCanceled(uint256 indexed orderId, address indexed nftHolder, uint8 orderType, uint8 cancelReason)
func (_Bindings *BindingsFilterer) ParseNftOrderCanceled(log types.Log) (*BindingsNftOrderCanceled, error) {
	event := new(BindingsNftOrderCanceled)
	if err := _Bindings.contract.UnpackLog(event, "NftOrderCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPairMaxLeverageUpdatedIterator is returned from FilterPairMaxLeverageUpdated and is used to iterate over the raw logs and unpacked data for PairMaxLeverageUpdated events raised by the Bindings contract.
type BindingsPairMaxLeverageUpdatedIterator struct {
	Event *BindingsPairMaxLeverageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsPairMaxLeverageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPairMaxLeverageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsPairMaxLeverageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsPairMaxLeverageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPairMaxLeverageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPairMaxLeverageUpdated represents a PairMaxLeverageUpdated event raised by the Bindings contract.
type BindingsPairMaxLeverageUpdated struct {
	PairIndex   *big.Int
	MaxLeverage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPairMaxLeverageUpdated is a free log retrieval operation binding the contract event 0x95924bc10431f9a625a06fe5a27d55f4348510b2da42a18fe3bf2a6f2c4eab67.
//
// Solidity: event PairMaxLeverageUpdated(uint256 indexed pairIndex, uint256 maxLeverage)
func (_Bindings *BindingsFilterer) FilterPairMaxLeverageUpdated(opts *bind.FilterOpts, pairIndex []*big.Int) (*BindingsPairMaxLeverageUpdatedIterator, error) {

	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PairMaxLeverageUpdated", pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPairMaxLeverageUpdatedIterator{contract: _Bindings.contract, event: "PairMaxLeverageUpdated", logs: logs, sub: sub}, nil
}

// WatchPairMaxLeverageUpdated is a free log subscription operation binding the contract event 0x95924bc10431f9a625a06fe5a27d55f4348510b2da42a18fe3bf2a6f2c4eab67.
//
// Solidity: event PairMaxLeverageUpdated(uint256 indexed pairIndex, uint256 maxLeverage)
func (_Bindings *BindingsFilterer) WatchPairMaxLeverageUpdated(opts *bind.WatchOpts, sink chan<- *BindingsPairMaxLeverageUpdated, pairIndex []*big.Int) (event.Subscription, error) {

	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PairMaxLeverageUpdated", pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPairMaxLeverageUpdated)
				if err := _Bindings.contract.UnpackLog(event, "PairMaxLeverageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePairMaxLeverageUpdated is a log parse operation binding the contract event 0x95924bc10431f9a625a06fe5a27d55f4348510b2da42a18fe3bf2a6f2c4eab67.
//
// Solidity: event PairMaxLeverageUpdated(uint256 indexed pairIndex, uint256 maxLeverage)
func (_Bindings *BindingsFilterer) ParsePairMaxLeverageUpdated(log types.Log) (*BindingsPairMaxLeverageUpdated, error) {
	event := new(BindingsPairMaxLeverageUpdated)
	if err := _Bindings.contract.UnpackLog(event, "PairMaxLeverageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Bindings contract.
type BindingsPauseIterator struct {
	Event *BindingsPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPause represents a Pause event raised by the Bindings contract.
type BindingsPause struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x9422424b175dda897495a07b091ef74a3ef715cf6d866fc972954c1c7f459304.
//
// Solidity: event Pause(bool paused)
func (_Bindings *BindingsFilterer) FilterPause(opts *bind.FilterOpts) (*BindingsPauseIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &BindingsPauseIterator{contract: _Bindings.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x9422424b175dda897495a07b091ef74a3ef715cf6d866fc972954c1c7f459304.
//
// Solidity: event Pause(bool paused)
func (_Bindings *BindingsFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *BindingsPause) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPause)
				if err := _Bindings.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePause is a log parse operation binding the contract event 0x9422424b175dda897495a07b091ef74a3ef715cf6d866fc972954c1c7f459304.
//
// Solidity: event Pause(bool paused)
func (_Bindings *BindingsFilterer) ParsePause(log types.Log) (*BindingsPause, error) {
	event := new(BindingsPause)
	if err := _Bindings.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsReferralFeeChargedIterator is returned from FilterReferralFeeCharged and is used to iterate over the raw logs and unpacked data for ReferralFeeCharged events raised by the Bindings contract.
type BindingsReferralFeeChargedIterator struct {
	Event *BindingsReferralFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsReferralFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsReferralFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsReferralFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsReferralFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsReferralFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsReferralFeeCharged represents a ReferralFeeCharged event raised by the Bindings contract.
type BindingsReferralFeeCharged struct {
	Trader   common.Address
	ValueDai *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReferralFeeCharged is a free log retrieval operation binding the contract event 0x0f5273269f52308b9c40fafda3ca13cc42f715fcd795365e87f351f59e249313.
//
// Solidity: event ReferralFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) FilterReferralFeeCharged(opts *bind.FilterOpts, trader []common.Address) (*BindingsReferralFeeChargedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ReferralFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsReferralFeeChargedIterator{contract: _Bindings.contract, event: "ReferralFeeCharged", logs: logs, sub: sub}, nil
}

// WatchReferralFeeCharged is a free log subscription operation binding the contract event 0x0f5273269f52308b9c40fafda3ca13cc42f715fcd795365e87f351f59e249313.
//
// Solidity: event ReferralFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) WatchReferralFeeCharged(opts *bind.WatchOpts, sink chan<- *BindingsReferralFeeCharged, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ReferralFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsReferralFeeCharged)
				if err := _Bindings.contract.UnpackLog(event, "ReferralFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReferralFeeCharged is a log parse operation binding the contract event 0x0f5273269f52308b9c40fafda3ca13cc42f715fcd795365e87f351f59e249313.
//
// Solidity: event ReferralFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) ParseReferralFeeCharged(log types.Log) (*BindingsReferralFeeCharged, error) {
	event := new(BindingsReferralFeeCharged)
	if err := _Bindings.contract.UnpackLog(event, "ReferralFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSssFeeChargedIterator is returned from FilterSssFeeCharged and is used to iterate over the raw logs and unpacked data for SssFeeCharged events raised by the Bindings contract.
type BindingsSssFeeChargedIterator struct {
	Event *BindingsSssFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsSssFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSssFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsSssFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsSssFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSssFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSssFeeCharged represents a SssFeeCharged event raised by the Bindings contract.
type BindingsSssFeeCharged struct {
	Trader   common.Address
	ValueDai *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSssFeeCharged is a free log retrieval operation binding the contract event 0xd1e388cc27c5125a80cf538c12b26dc5a784071d324a81a736e4d17f238588e4.
//
// Solidity: event SssFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) FilterSssFeeCharged(opts *bind.FilterOpts, trader []common.Address) (*BindingsSssFeeChargedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SssFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSssFeeChargedIterator{contract: _Bindings.contract, event: "SssFeeCharged", logs: logs, sub: sub}, nil
}

// WatchSssFeeCharged is a free log subscription operation binding the contract event 0xd1e388cc27c5125a80cf538c12b26dc5a784071d324a81a736e4d17f238588e4.
//
// Solidity: event SssFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) WatchSssFeeCharged(opts *bind.WatchOpts, sink chan<- *BindingsSssFeeCharged, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SssFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSssFeeCharged)
				if err := _Bindings.contract.UnpackLog(event, "SssFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSssFeeCharged is a log parse operation binding the contract event 0xd1e388cc27c5125a80cf538c12b26dc5a784071d324a81a736e4d17f238588e4.
//
// Solidity: event SssFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) ParseSssFeeCharged(log types.Log) (*BindingsSssFeeCharged, error) {
	event := new(BindingsSssFeeCharged)
	if err := _Bindings.contract.UnpackLog(event, "SssFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTriggerFeeChargedIterator is returned from FilterTriggerFeeCharged and is used to iterate over the raw logs and unpacked data for TriggerFeeCharged events raised by the Bindings contract.
type BindingsTriggerFeeChargedIterator struct {
	Event *BindingsTriggerFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsTriggerFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTriggerFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsTriggerFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsTriggerFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTriggerFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTriggerFeeCharged represents a TriggerFeeCharged event raised by the Bindings contract.
type BindingsTriggerFeeCharged struct {
	Trader   common.Address
	ValueDai *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTriggerFeeCharged is a free log retrieval operation binding the contract event 0x17fa86cf4833d28c6224a940e6bd001f2db0cb3d89d69727765679b3efee6559.
//
// Solidity: event TriggerFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) FilterTriggerFeeCharged(opts *bind.FilterOpts, trader []common.Address) (*BindingsTriggerFeeChargedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TriggerFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTriggerFeeChargedIterator{contract: _Bindings.contract, event: "TriggerFeeCharged", logs: logs, sub: sub}, nil
}

// WatchTriggerFeeCharged is a free log subscription operation binding the contract event 0x17fa86cf4833d28c6224a940e6bd001f2db0cb3d89d69727765679b3efee6559.
//
// Solidity: event TriggerFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) WatchTriggerFeeCharged(opts *bind.WatchOpts, sink chan<- *BindingsTriggerFeeCharged, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TriggerFeeCharged", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTriggerFeeCharged)
				if err := _Bindings.contract.UnpackLog(event, "TriggerFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerFeeCharged is a log parse operation binding the contract event 0x17fa86cf4833d28c6224a940e6bd001f2db0cb3d89d69727765679b3efee6559.
//
// Solidity: event TriggerFeeCharged(address indexed trader, uint256 valueDai)
func (_Bindings *BindingsFilterer) ParseTriggerFeeCharged(log types.Log) (*BindingsTriggerFeeCharged, error) {
	event := new(BindingsTriggerFeeCharged)
	if err := _Bindings.contract.UnpackLog(event, "TriggerFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
