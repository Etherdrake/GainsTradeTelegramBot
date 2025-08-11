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

// GFarmTradingStorageV5OpenLimitOrder is an auto generated low-level Go binding around an user-defined struct.
type GFarmTradingStorageV5OpenLimitOrder struct {
	Trader           common.Address
	PairIndex        *big.Int
	Index            *big.Int
	PositionSize     *big.Int
	SpreadReductionP *big.Int
	Buy              bool
	Leverage         *big.Int
	Tp               *big.Int
	Sl               *big.Int
	MinPrice         *big.Int
	MaxPrice         *big.Int
	Block            *big.Int
	TokenId          *big.Int
}

// GFarmTradingStorageV5PendingMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type GFarmTradingStorageV5PendingMarketOrder struct {
	Trade GFarmTradingStorageV5Trade
	Block *big.Int
	WantedPrice      *big.Int
	SlippageP        *big.Int
	SpreadReductionP *big.Int
	TokenId          *big.Int
}

// GFarmTradingStorageV5PendingNftOrder is an auto generated low-level Go binding around an user-defined struct.
type GFarmTradingStorageV5PendingNftOrder struct {
	NftHolder common.Address
	NftId     *big.Int
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	OrderType uint8
}

// GFarmTradingStorageV5Trade is an auto generated low-level Go binding around an user-defined struct.
type GFarmTradingStorageV5Trade struct {
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

// GFarmTradingStorageV5TradeInfo is an auto generated low-level Go binding around an user-defined struct.
type GFarmTradingStorageV5TradeInfo struct {
	TokenId           *big.Int
	TokenPriceDai     *big.Int
	OpenInterestDai   *big.Int
	TpLastUpdated     *big.Int
	SlLastUpdated     *big.Int
	BeingMarketClosed bool
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractNftInterfaceV5[5]\",\"name\":\"nfts\",\"type\":\"address[5]\"}],\"name\":\"NftsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NumberUpdatedPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"name\":\"SpreadReductionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"SupportedTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"TradingContractAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"TradingContractRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trading\",\"type\":\"address\"}],\"name\":\"addTradingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callbacks\",\"outputs\":[{\"internalType\":\"contractPausableInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractTokenInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultLeverageUnlocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devFeesDai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devFeesToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"distributeLpRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"}],\"name\":\"firstEmptyOpenLimitIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"}],\"name\":\"firstEmptyTradeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getLeverageUnlocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOpenLimitOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.OpenLimitOrder\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenLimitOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.OpenLimitOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getPendingOrderIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getReferral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSpreadReductionsArray\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govFeesDai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govFeesToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_leveragedPositionSize\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_dai\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_fullFee\",\"type\":\"bool\"}],\"name\":\"handleDevGovFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_mint\",\"type\":\"bool\"}],\"name\":\"handleTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"hasOpenLimitOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseNftRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseReferralRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTradingContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkErc677\",\"outputs\":[{\"internalType\":\"contractTokenInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGainP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPendingMarketOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSlP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTradesPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTradesPerPair\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftLastSuccess\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftSuccessTimelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nfts\",\"outputs\":[{\"internalType\":\"contractNftInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openInterestDai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openLimitOrderIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openLimitOrders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openLimitOrdersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openTrades\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openTradesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openTradesInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPriceDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openInterestDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tpLastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slLastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"beingMarketClosed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pairTraders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"}],\"name\":\"pairTradersArray\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pairTradersId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingMarketCloseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingMarketOpenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingOrderIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"pendingOrderIdsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractPoolInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceAggregator\",\"outputs\":[{\"internalType\":\"contractAggregatorInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trading\",\"type\":\"address\"}],\"name\":\"removeTradingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reqID_pendingMarketOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reqID_pendingNftOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nftHolder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"enumGFarmTradingStorageV5.LimitOrder\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_callbacks\",\"type\":\"address\"}],\"name\":\"setCallbacks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lev\",\"type\":\"uint256\"}],\"name\":\"setDefaultLeverageUnlocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dev\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newLeverage\",\"type\":\"uint256\"}],\"name\":\"setLeverageUnlocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxGainP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newMaxOpenInterest\",\"type\":\"uint256\"}],\"name\":\"setMaxOpenInterestDai\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPendingMarketOrders\",\"type\":\"uint256\"}],\"name\":\"setMaxPendingMarketOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxSlP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTradesPerBlock\",\"type\":\"uint256\"}],\"name\":\"setMaxTradesPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTradesPerPair\",\"type\":\"uint256\"}],\"name\":\"setMaxTradesPerPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"setNftSuccessTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setPriceAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_r\",\"type\":\"uint256[5]\"}],\"name\":\"setSpreadReductionsP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenDaiRouter\",\"type\":\"address\"}],\"name\":\"setTokenDaiRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trading\",\"type\":\"address\"}],\"name\":\"setTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spreadReductionsP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.OpenLimitOrder\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"storeOpenLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.PendingMarketOrder\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_open\",\"type\":\"bool\"}],\"name\":\"storePendingMarketOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nftHolder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"enumGFarmTradingStorageV5.LimitOrder\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"internalType\":\"structGFarmTradingStorageV5.PendingNftOrder\",\"name\":\"_nftOrder\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"storePendingNftOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"storeReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.Trade\",\"name\":\"_trade\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPriceDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openInterestDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tpLastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slLastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"beingMarketClosed\",\"type\":\"bool\"}],\"internalType\":\"structGFarmTradingStorageV5.TradeInfo\",\"name\":\"_tradeInfo\",\"type\":\"tuple\"}],\"name\":\"storeTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractTokenInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDaiRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensBurned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"traders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"leverageUnlocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralRewardsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tradesPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trading\",\"outputs\":[{\"internalType\":\"contractPausableInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferDai\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_leveragedPosDai\",\"type\":\"uint256\"}],\"name\":\"transferLinkToAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"unregisterOpenLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_open\",\"type\":\"bool\"}],\"name\":\"unregisterPendingMarketOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_order\",\"type\":\"uint256\"}],\"name\":\"unregisterPendingNftOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"unregisterTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNftInterfaceV5[5]\",\"name\":\"_nfts\",\"type\":\"address[5]\"}],\"name\":\"updateNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.OpenLimitOrder\",\"name\":\"_o\",\"type\":\"tuple\"}],\"name\":\"updateOpenLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newSl\",\"type\":\"uint256\"}],\"name\":\"updateSl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTokenInterfaceV5\",\"name\":\"_newToken\",\"type\":\"address\"}],\"name\":\"updateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newTp\",\"type\":\"uint256\"}],\"name\":\"updateTp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structGFarmTradingStorageV5.Trade\",\"name\":\"_t\",\"type\":\"tuple\"}],\"name\":\"updateTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) MINTERROLE() ([32]byte, error) {
	return _Bindings.Contract.MINTERROLE(&_Bindings.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) MINTERROLE() ([32]byte, error) {
	return _Bindings.Contract.MINTERROLE(&_Bindings.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Bindings *BindingsCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Bindings *BindingsSession) PRECISION() (*big.Int, error) {
	return _Bindings.Contract.PRECISION(&_Bindings.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Bindings *BindingsCallerSession) PRECISION() (*big.Int, error) {
	return _Bindings.Contract.PRECISION(&_Bindings.CallOpts)
}

// Callbacks is a free data retrieval call binding the contract method 0x00b12783.
//
// Solidity: function callbacks() view returns(address)
func (_Bindings *BindingsCaller) Callbacks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "callbacks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Callbacks is a free data retrieval call binding the contract method 0x00b12783.
//
// Solidity: function callbacks() view returns(address)
func (_Bindings *BindingsSession) Callbacks() (common.Address, error) {
	return _Bindings.Contract.Callbacks(&_Bindings.CallOpts)
}

// Callbacks is a free data retrieval call binding the contract method 0x00b12783.
//
// Solidity: function callbacks() view returns(address)
func (_Bindings *BindingsCallerSession) Callbacks() (common.Address, error) {
	return _Bindings.Contract.Callbacks(&_Bindings.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_Bindings *BindingsCaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_Bindings *BindingsSession) Dai() (common.Address, error) {
	return _Bindings.Contract.Dai(&_Bindings.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_Bindings *BindingsCallerSession) Dai() (common.Address, error) {
	return _Bindings.Contract.Dai(&_Bindings.CallOpts)
}

// DefaultLeverageUnlocked is a free data retrieval call binding the contract method 0x8bbb644c.
//
// Solidity: function defaultLeverageUnlocked() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultLeverageUnlocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultLeverageUnlocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultLeverageUnlocked is a free data retrieval call binding the contract method 0x8bbb644c.
//
// Solidity: function defaultLeverageUnlocked() view returns(uint256)
func (_Bindings *BindingsSession) DefaultLeverageUnlocked() (*big.Int, error) {
	return _Bindings.Contract.DefaultLeverageUnlocked(&_Bindings.CallOpts)
}

// DefaultLeverageUnlocked is a free data retrieval call binding the contract method 0x8bbb644c.
//
// Solidity: function defaultLeverageUnlocked() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultLeverageUnlocked() (*big.Int, error) {
	return _Bindings.Contract.DefaultLeverageUnlocked(&_Bindings.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_Bindings *BindingsCaller) Dev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "dev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_Bindings *BindingsSession) Dev() (common.Address, error) {
	return _Bindings.Contract.Dev(&_Bindings.CallOpts)
}

// Dev is a free data retrieval call binding the contract method 0x91cca3db.
//
// Solidity: function dev() view returns(address)
func (_Bindings *BindingsCallerSession) Dev() (common.Address, error) {
	return _Bindings.Contract.Dev(&_Bindings.CallOpts)
}

// DevFeesDai is a free data retrieval call binding the contract method 0xea5c3c33.
//
// Solidity: function devFeesDai() view returns(uint256)
func (_Bindings *BindingsCaller) DevFeesDai(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "devFeesDai")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevFeesDai is a free data retrieval call binding the contract method 0xea5c3c33.
//
// Solidity: function devFeesDai() view returns(uint256)
func (_Bindings *BindingsSession) DevFeesDai() (*big.Int, error) {
	return _Bindings.Contract.DevFeesDai(&_Bindings.CallOpts)
}

// DevFeesDai is a free data retrieval call binding the contract method 0xea5c3c33.
//
// Solidity: function devFeesDai() view returns(uint256)
func (_Bindings *BindingsCallerSession) DevFeesDai() (*big.Int, error) {
	return _Bindings.Contract.DevFeesDai(&_Bindings.CallOpts)
}

// DevFeesToken is a free data retrieval call binding the contract method 0xee92a4ff.
//
// Solidity: function devFeesToken() view returns(uint256)
func (_Bindings *BindingsCaller) DevFeesToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "devFeesToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevFeesToken is a free data retrieval call binding the contract method 0xee92a4ff.
//
// Solidity: function devFeesToken() view returns(uint256)
func (_Bindings *BindingsSession) DevFeesToken() (*big.Int, error) {
	return _Bindings.Contract.DevFeesToken(&_Bindings.CallOpts)
}

// DevFeesToken is a free data retrieval call binding the contract method 0xee92a4ff.
//
// Solidity: function devFeesToken() view returns(uint256)
func (_Bindings *BindingsCallerSession) DevFeesToken() (*big.Int, error) {
	return _Bindings.Contract.DevFeesToken(&_Bindings.CallOpts)
}

// FirstEmptyOpenLimitIndex is a free data retrieval call binding the contract method 0x44d44f5a.
//
// Solidity: function firstEmptyOpenLimitIndex(address trader, uint256 pairIndex) view returns(uint256 index)
func (_Bindings *BindingsCaller) FirstEmptyOpenLimitIndex(opts *bind.CallOpts, trader common.Address, pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "firstEmptyOpenLimitIndex", trader, pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstEmptyOpenLimitIndex is a free data retrieval call binding the contract method 0x44d44f5a.
//
// Solidity: function firstEmptyOpenLimitIndex(address trader, uint256 pairIndex) view returns(uint256 index)
func (_Bindings *BindingsSession) FirstEmptyOpenLimitIndex(trader common.Address, pairIndex *big.Int) (*big.Int, error) {
	return _Bindings.Contract.FirstEmptyOpenLimitIndex(&_Bindings.CallOpts, trader, pairIndex)
}

// FirstEmptyOpenLimitIndex is a free data retrieval call binding the contract method 0x44d44f5a.
//
// Solidity: function firstEmptyOpenLimitIndex(address trader, uint256 pairIndex) view returns(uint256 index)
func (_Bindings *BindingsCallerSession) FirstEmptyOpenLimitIndex(trader common.Address, pairIndex *big.Int) (*big.Int, error) {
	return _Bindings.Contract.FirstEmptyOpenLimitIndex(&_Bindings.CallOpts, trader, pairIndex)
}

// FirstEmptyTradeIndex is a free data retrieval call binding the contract method 0x292c1617.
//
// Solidity: function firstEmptyTradeIndex(address trader, uint256 pairIndex) view returns(uint256 index)
func (_Bindings *BindingsCaller) FirstEmptyTradeIndex(opts *bind.CallOpts, trader common.Address, pairIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "firstEmptyTradeIndex", trader, pairIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstEmptyTradeIndex is a free data retrieval call binding the contract method 0x292c1617.
//
// Solidity: function firstEmptyTradeIndex(address trader, uint256 pairIndex) view returns(uint256 index)
func (_Bindings *BindingsSession) FirstEmptyTradeIndex(trader common.Address, pairIndex *big.Int) (*big.Int, error) {
	return _Bindings.Contract.FirstEmptyTradeIndex(&_Bindings.CallOpts, trader, pairIndex)
}

// FirstEmptyTradeIndex is a free data retrieval call binding the contract method 0x292c1617.
//
// Solidity: function firstEmptyTradeIndex(address trader, uint256 pairIndex) view returns(uint256 index)
func (_Bindings *BindingsCallerSession) FirstEmptyTradeIndex(trader common.Address, pairIndex *big.Int) (*big.Int, error) {
	return _Bindings.Contract.FirstEmptyTradeIndex(&_Bindings.CallOpts, trader, pairIndex)
}

// GetLeverageUnlocked is a free data retrieval call binding the contract method 0x68ca199d.
//
// Solidity: function getLeverageUnlocked(address _trader) view returns(uint256)
func (_Bindings *BindingsCaller) GetLeverageUnlocked(opts *bind.CallOpts, _trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getLeverageUnlocked", _trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLeverageUnlocked is a free data retrieval call binding the contract method 0x68ca199d.
//
// Solidity: function getLeverageUnlocked(address _trader) view returns(uint256)
func (_Bindings *BindingsSession) GetLeverageUnlocked(_trader common.Address) (*big.Int, error) {
	return _Bindings.Contract.GetLeverageUnlocked(&_Bindings.CallOpts, _trader)
}

// GetLeverageUnlocked is a free data retrieval call binding the contract method 0x68ca199d.
//
// Solidity: function getLeverageUnlocked(address _trader) view returns(uint256)
func (_Bindings *BindingsCallerSession) GetLeverageUnlocked(_trader common.Address) (*big.Int, error) {
	return _Bindings.Contract.GetLeverageUnlocked(&_Bindings.CallOpts, _trader)
}

// GetOpenLimitOrder is a free data retrieval call binding the contract method 0xb8878a2c.
//
// Solidity: function getOpenLimitOrder(address _trader, uint256 _pairIndex, uint256 _index) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Bindings *BindingsCaller) GetOpenLimitOrder(opts *bind.CallOpts, _trader common.Address, _pairIndex *big.Int, _index *big.Int) (GFarmTradingStorageV5OpenLimitOrder, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getOpenLimitOrder", _trader, _pairIndex, _index)

	if err != nil {
		return *new(GFarmTradingStorageV5OpenLimitOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(GFarmTradingStorageV5OpenLimitOrder)).(*GFarmTradingStorageV5OpenLimitOrder)

	return out0, err

}

// GetOpenLimitOrder is a free data retrieval call binding the contract method 0xb8878a2c.
//
// Solidity: function getOpenLimitOrder(address _trader, uint256 _pairIndex, uint256 _index) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Bindings *BindingsSession) GetOpenLimitOrder(_trader common.Address, _pairIndex *big.Int, _index *big.Int) (GFarmTradingStorageV5OpenLimitOrder, error) {
	return _Bindings.Contract.GetOpenLimitOrder(&_Bindings.CallOpts, _trader, _pairIndex, _index)
}

// GetOpenLimitOrder is a free data retrieval call binding the contract method 0xb8878a2c.
//
// Solidity: function getOpenLimitOrder(address _trader, uint256 _pairIndex, uint256 _index) view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Bindings *BindingsCallerSession) GetOpenLimitOrder(_trader common.Address, _pairIndex *big.Int, _index *big.Int) (GFarmTradingStorageV5OpenLimitOrder, error) {
	return _Bindings.Contract.GetOpenLimitOrder(&_Bindings.CallOpts, _trader, _pairIndex, _index)
}

// GetOpenLimitOrders is a free data retrieval call binding the contract method 0x8a0bd702.
//
// Solidity: function getOpenLimitOrders() view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Bindings *BindingsCaller) GetOpenLimitOrders(opts *bind.CallOpts) ([]GFarmTradingStorageV5OpenLimitOrder, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getOpenLimitOrders")

	if err != nil {
		return *new([]GFarmTradingStorageV5OpenLimitOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]GFarmTradingStorageV5OpenLimitOrder)).(*[]GFarmTradingStorageV5OpenLimitOrder)

	return out0, err

}

// GetOpenLimitOrders is a free data retrieval call binding the contract method 0x8a0bd702.
//
// Solidity: function getOpenLimitOrders() view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Bindings *BindingsSession) GetOpenLimitOrders() ([]GFarmTradingStorageV5OpenLimitOrder, error) {
	return _Bindings.Contract.GetOpenLimitOrders(&_Bindings.CallOpts)
}

// GetOpenLimitOrders is a free data retrieval call binding the contract method 0x8a0bd702.
//
// Solidity: function getOpenLimitOrders() view returns((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_Bindings *BindingsCallerSession) GetOpenLimitOrders() ([]GFarmTradingStorageV5OpenLimitOrder, error) {
	return _Bindings.Contract.GetOpenLimitOrders(&_Bindings.CallOpts)
}

// GetPendingOrderIds is a free data retrieval call binding the contract method 0x5c76ac9e.
//
// Solidity: function getPendingOrderIds(address _trader) view returns(uint256[])
func (_Bindings *BindingsCaller) GetPendingOrderIds(opts *bind.CallOpts, _trader common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getPendingOrderIds", _trader)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingOrderIds is a free data retrieval call binding the contract method 0x5c76ac9e.
//
// Solidity: function getPendingOrderIds(address _trader) view returns(uint256[])
func (_Bindings *BindingsSession) GetPendingOrderIds(_trader common.Address) ([]*big.Int, error) {
	return _Bindings.Contract.GetPendingOrderIds(&_Bindings.CallOpts, _trader)
}

// GetPendingOrderIds is a free data retrieval call binding the contract method 0x5c76ac9e.
//
// Solidity: function getPendingOrderIds(address _trader) view returns(uint256[])
func (_Bindings *BindingsCallerSession) GetPendingOrderIds(_trader common.Address) ([]*big.Int, error) {
	return _Bindings.Contract.GetPendingOrderIds(&_Bindings.CallOpts, _trader)
}

// GetReferral is a free data retrieval call binding the contract method 0x3b0f0f2f.
//
// Solidity: function getReferral(address _trader) view returns(address)
func (_Bindings *BindingsCaller) GetReferral(opts *bind.CallOpts, _trader common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getReferral", _trader)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReferral is a free data retrieval call binding the contract method 0x3b0f0f2f.
//
// Solidity: function getReferral(address _trader) view returns(address)
func (_Bindings *BindingsSession) GetReferral(_trader common.Address) (common.Address, error) {
	return _Bindings.Contract.GetReferral(&_Bindings.CallOpts, _trader)
}

// GetReferral is a free data retrieval call binding the contract method 0x3b0f0f2f.
//
// Solidity: function getReferral(address _trader) view returns(address)
func (_Bindings *BindingsCallerSession) GetReferral(_trader common.Address) (common.Address, error) {
	return _Bindings.Contract.GetReferral(&_Bindings.CallOpts, _trader)
}

// GetSpreadReductionsArray is a free data retrieval call binding the contract method 0x56a6796f.
//
// Solidity: function getSpreadReductionsArray() view returns(uint256[5])
func (_Bindings *BindingsCaller) GetSpreadReductionsArray(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getSpreadReductionsArray")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSpreadReductionsArray is a free data retrieval call binding the contract method 0x56a6796f.
//
// Solidity: function getSpreadReductionsArray() view returns(uint256[5])
func (_Bindings *BindingsSession) GetSpreadReductionsArray() ([5]*big.Int, error) {
	return _Bindings.Contract.GetSpreadReductionsArray(&_Bindings.CallOpts)
}

// GetSpreadReductionsArray is a free data retrieval call binding the contract method 0x56a6796f.
//
// Solidity: function getSpreadReductionsArray() view returns(uint256[5])
func (_Bindings *BindingsCallerSession) GetSpreadReductionsArray() ([5]*big.Int, error) {
	return _Bindings.Contract.GetSpreadReductionsArray(&_Bindings.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_Bindings *BindingsCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_Bindings *BindingsSession) GetSupportedTokens() ([]common.Address, error) {
	return _Bindings.Contract.GetSupportedTokens(&_Bindings.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_Bindings *BindingsCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _Bindings.Contract.GetSupportedTokens(&_Bindings.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Bindings *BindingsCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Bindings *BindingsSession) Gov() (common.Address, error) {
	return _Bindings.Contract.Gov(&_Bindings.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Bindings *BindingsCallerSession) Gov() (common.Address, error) {
	return _Bindings.Contract.Gov(&_Bindings.CallOpts)
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

// GovFeesToken is a free data retrieval call binding the contract method 0xfdb59413.
//
// Solidity: function govFeesToken() view returns(uint256)
func (_Bindings *BindingsCaller) GovFeesToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "govFeesToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GovFeesToken is a free data retrieval call binding the contract method 0xfdb59413.
//
// Solidity: function govFeesToken() view returns(uint256)
func (_Bindings *BindingsSession) GovFeesToken() (*big.Int, error) {
	return _Bindings.Contract.GovFeesToken(&_Bindings.CallOpts)
}

// GovFeesToken is a free data retrieval call binding the contract method 0xfdb59413.
//
// Solidity: function govFeesToken() view returns(uint256)
func (_Bindings *BindingsCallerSession) GovFeesToken() (*big.Int, error) {
	return _Bindings.Contract.GovFeesToken(&_Bindings.CallOpts)
}

// HasOpenLimitOrder is a free data retrieval call binding the contract method 0xdcdf339a.
//
// Solidity: function hasOpenLimitOrder(address trader, uint256 pairIndex, uint256 index) view returns(bool)
func (_Bindings *BindingsCaller) HasOpenLimitOrder(opts *bind.CallOpts, trader common.Address, pairIndex *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "hasOpenLimitOrder", trader, pairIndex, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOpenLimitOrder is a free data retrieval call binding the contract method 0xdcdf339a.
//
// Solidity: function hasOpenLimitOrder(address trader, uint256 pairIndex, uint256 index) view returns(bool)
func (_Bindings *BindingsSession) HasOpenLimitOrder(trader common.Address, pairIndex *big.Int, index *big.Int) (bool, error) {
	return _Bindings.Contract.HasOpenLimitOrder(&_Bindings.CallOpts, trader, pairIndex, index)
}

// HasOpenLimitOrder is a free data retrieval call binding the contract method 0xdcdf339a.
//
// Solidity: function hasOpenLimitOrder(address trader, uint256 pairIndex, uint256 index) view returns(bool)
func (_Bindings *BindingsCallerSession) HasOpenLimitOrder(trader common.Address, pairIndex *big.Int, index *big.Int) (bool, error) {
	return _Bindings.Contract.HasOpenLimitOrder(&_Bindings.CallOpts, trader, pairIndex, index)
}

// IsTradingContract is a free data retrieval call binding the contract method 0x6d81d981.
//
// Solidity: function isTradingContract(address ) view returns(bool)
func (_Bindings *BindingsCaller) IsTradingContract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isTradingContract", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTradingContract is a free data retrieval call binding the contract method 0x6d81d981.
//
// Solidity: function isTradingContract(address ) view returns(bool)
func (_Bindings *BindingsSession) IsTradingContract(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.IsTradingContract(&_Bindings.CallOpts, arg0)
}

// IsTradingContract is a free data retrieval call binding the contract method 0x6d81d981.
//
// Solidity: function isTradingContract(address ) view returns(bool)
func (_Bindings *BindingsCallerSession) IsTradingContract(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.IsTradingContract(&_Bindings.CallOpts, arg0)
}

// LinkErc677 is a free data retrieval call binding the contract method 0xfece4eac.
//
// Solidity: function linkErc677() view returns(address)
func (_Bindings *BindingsCaller) LinkErc677(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "linkErc677")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkErc677 is a free data retrieval call binding the contract method 0xfece4eac.
//
// Solidity: function linkErc677() view returns(address)
func (_Bindings *BindingsSession) LinkErc677() (common.Address, error) {
	return _Bindings.Contract.LinkErc677(&_Bindings.CallOpts)
}

// LinkErc677 is a free data retrieval call binding the contract method 0xfece4eac.
//
// Solidity: function linkErc677() view returns(address)
func (_Bindings *BindingsCallerSession) LinkErc677() (common.Address, error) {
	return _Bindings.Contract.LinkErc677(&_Bindings.CallOpts)
}

// MaxGainP is a free data retrieval call binding the contract method 0xc122fea6.
//
// Solidity: function maxGainP() view returns(uint256)
func (_Bindings *BindingsCaller) MaxGainP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "maxGainP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGainP is a free data retrieval call binding the contract method 0xc122fea6.
//
// Solidity: function maxGainP() view returns(uint256)
func (_Bindings *BindingsSession) MaxGainP() (*big.Int, error) {
	return _Bindings.Contract.MaxGainP(&_Bindings.CallOpts)
}

// MaxGainP is a free data retrieval call binding the contract method 0xc122fea6.
//
// Solidity: function maxGainP() view returns(uint256)
func (_Bindings *BindingsCallerSession) MaxGainP() (*big.Int, error) {
	return _Bindings.Contract.MaxGainP(&_Bindings.CallOpts)
}

// MaxPendingMarketOrders is a free data retrieval call binding the contract method 0x8049fde7.
//
// Solidity: function maxPendingMarketOrders() view returns(uint256)
func (_Bindings *BindingsCaller) MaxPendingMarketOrders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "maxPendingMarketOrders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPendingMarketOrders is a free data retrieval call binding the contract method 0x8049fde7.
//
// Solidity: function maxPendingMarketOrders() view returns(uint256)
func (_Bindings *BindingsSession) MaxPendingMarketOrders() (*big.Int, error) {
	return _Bindings.Contract.MaxPendingMarketOrders(&_Bindings.CallOpts)
}

// MaxPendingMarketOrders is a free data retrieval call binding the contract method 0x8049fde7.
//
// Solidity: function maxPendingMarketOrders() view returns(uint256)
func (_Bindings *BindingsCallerSession) MaxPendingMarketOrders() (*big.Int, error) {
	return _Bindings.Contract.MaxPendingMarketOrders(&_Bindings.CallOpts)
}

// MaxSlP is a free data retrieval call binding the contract method 0x5a912551.
//
// Solidity: function maxSlP() view returns(uint256)
func (_Bindings *BindingsCaller) MaxSlP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "maxSlP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSlP is a free data retrieval call binding the contract method 0x5a912551.
//
// Solidity: function maxSlP() view returns(uint256)
func (_Bindings *BindingsSession) MaxSlP() (*big.Int, error) {
	return _Bindings.Contract.MaxSlP(&_Bindings.CallOpts)
}

// MaxSlP is a free data retrieval call binding the contract method 0x5a912551.
//
// Solidity: function maxSlP() view returns(uint256)
func (_Bindings *BindingsCallerSession) MaxSlP() (*big.Int, error) {
	return _Bindings.Contract.MaxSlP(&_Bindings.CallOpts)
}

// MaxTradesPerBlock is a free data retrieval call binding the contract method 0x796d82d9.
//
// Solidity: function maxTradesPerBlock() view returns(uint256)
func (_Bindings *BindingsCaller) MaxTradesPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "maxTradesPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTradesPerBlock is a free data retrieval call binding the contract method 0x796d82d9.
//
// Solidity: function maxTradesPerBlock() view returns(uint256)
func (_Bindings *BindingsSession) MaxTradesPerBlock() (*big.Int, error) {
	return _Bindings.Contract.MaxTradesPerBlock(&_Bindings.CallOpts)
}

// MaxTradesPerBlock is a free data retrieval call binding the contract method 0x796d82d9.
//
// Solidity: function maxTradesPerBlock() view returns(uint256)
func (_Bindings *BindingsCallerSession) MaxTradesPerBlock() (*big.Int, error) {
	return _Bindings.Contract.MaxTradesPerBlock(&_Bindings.CallOpts)
}

// MaxTradesPerPair is a free data retrieval call binding the contract method 0xf65d9dbe.
//
// Solidity: function maxTradesPerPair() view returns(uint256)
func (_Bindings *BindingsCaller) MaxTradesPerPair(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "maxTradesPerPair")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTradesPerPair is a free data retrieval call binding the contract method 0xf65d9dbe.
//
// Solidity: function maxTradesPerPair() view returns(uint256)
func (_Bindings *BindingsSession) MaxTradesPerPair() (*big.Int, error) {
	return _Bindings.Contract.MaxTradesPerPair(&_Bindings.CallOpts)
}

// MaxTradesPerPair is a free data retrieval call binding the contract method 0xf65d9dbe.
//
// Solidity: function maxTradesPerPair() view returns(uint256)
func (_Bindings *BindingsCallerSession) MaxTradesPerPair() (*big.Int, error) {
	return _Bindings.Contract.MaxTradesPerPair(&_Bindings.CallOpts)
}

// NftLastSuccess is a free data retrieval call binding the contract method 0xaa3b7d77.
//
// Solidity: function nftLastSuccess(uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) NftLastSuccess(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "nftLastSuccess", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftLastSuccess is a free data retrieval call binding the contract method 0xaa3b7d77.
//
// Solidity: function nftLastSuccess(uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) NftLastSuccess(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.NftLastSuccess(&_Bindings.CallOpts, arg0)
}

// NftLastSuccess is a free data retrieval call binding the contract method 0xaa3b7d77.
//
// Solidity: function nftLastSuccess(uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) NftLastSuccess(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.NftLastSuccess(&_Bindings.CallOpts, arg0)
}

// NftRewards is a free data retrieval call binding the contract method 0x5ad9f993.
//
// Solidity: function nftRewards() view returns(uint256)
func (_Bindings *BindingsCaller) NftRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "nftRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftRewards is a free data retrieval call binding the contract method 0x5ad9f993.
//
// Solidity: function nftRewards() view returns(uint256)
func (_Bindings *BindingsSession) NftRewards() (*big.Int, error) {
	return _Bindings.Contract.NftRewards(&_Bindings.CallOpts)
}

// NftRewards is a free data retrieval call binding the contract method 0x5ad9f993.
//
// Solidity: function nftRewards() view returns(uint256)
func (_Bindings *BindingsCallerSession) NftRewards() (*big.Int, error) {
	return _Bindings.Contract.NftRewards(&_Bindings.CallOpts)
}

// NftSuccessTimelock is a free data retrieval call binding the contract method 0xf8d24b48.
//
// Solidity: function nftSuccessTimelock() view returns(uint256)
func (_Bindings *BindingsCaller) NftSuccessTimelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "nftSuccessTimelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftSuccessTimelock is a free data retrieval call binding the contract method 0xf8d24b48.
//
// Solidity: function nftSuccessTimelock() view returns(uint256)
func (_Bindings *BindingsSession) NftSuccessTimelock() (*big.Int, error) {
	return _Bindings.Contract.NftSuccessTimelock(&_Bindings.CallOpts)
}

// NftSuccessTimelock is a free data retrieval call binding the contract method 0xf8d24b48.
//
// Solidity: function nftSuccessTimelock() view returns(uint256)
func (_Bindings *BindingsCallerSession) NftSuccessTimelock() (*big.Int, error) {
	return _Bindings.Contract.NftSuccessTimelock(&_Bindings.CallOpts)
}

// Nfts is a free data retrieval call binding the contract method 0x265aa621.
//
// Solidity: function nfts(uint256 ) view returns(address)
func (_Bindings *BindingsCaller) Nfts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "nfts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nfts is a free data retrieval call binding the contract method 0x265aa621.
//
// Solidity: function nfts(uint256 ) view returns(address)
func (_Bindings *BindingsSession) Nfts(arg0 *big.Int) (common.Address, error) {
	return _Bindings.Contract.Nfts(&_Bindings.CallOpts, arg0)
}

// Nfts is a free data retrieval call binding the contract method 0x265aa621.
//
// Solidity: function nfts(uint256 ) view returns(address)
func (_Bindings *BindingsCallerSession) Nfts(arg0 *big.Int) (common.Address, error) {
	return _Bindings.Contract.Nfts(&_Bindings.CallOpts, arg0)
}

// OpenInterestDai is a free data retrieval call binding the contract method 0x28daca21.
//
// Solidity: function openInterestDai(uint256 , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) OpenInterestDai(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "openInterestDai", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenInterestDai is a free data retrieval call binding the contract method 0x28daca21.
//
// Solidity: function openInterestDai(uint256 , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) OpenInterestDai(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenInterestDai(&_Bindings.CallOpts, arg0, arg1)
}

// OpenInterestDai is a free data retrieval call binding the contract method 0x28daca21.
//
// Solidity: function openInterestDai(uint256 , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) OpenInterestDai(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenInterestDai(&_Bindings.CallOpts, arg0, arg1)
}

// OpenLimitOrderIds is a free data retrieval call binding the contract method 0xeda00d54.
//
// Solidity: function openLimitOrderIds(address , uint256 , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) OpenLimitOrderIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "openLimitOrderIds", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenLimitOrderIds is a free data retrieval call binding the contract method 0xeda00d54.
//
// Solidity: function openLimitOrderIds(address , uint256 , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) OpenLimitOrderIds(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenLimitOrderIds(&_Bindings.CallOpts, arg0, arg1, arg2)
}

// OpenLimitOrderIds is a free data retrieval call binding the contract method 0xeda00d54.
//
// Solidity: function openLimitOrderIds(address , uint256 , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) OpenLimitOrderIds(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenLimitOrderIds(&_Bindings.CallOpts, arg0, arg1, arg2)
}

// OpenLimitOrders is a free data retrieval call binding the contract method 0x32149592.
//
// Solidity: function openLimitOrders(uint256 ) view returns(address trader, uint256 pairIndex, uint256 index, uint256 positionSize, uint256 spreadReductionP, bool buy, uint256 leverage, uint256 tp, uint256 sl, uint256 minPrice, uint256 maxPrice, uint256 block, uint256 tokenId)
func (_Bindings *BindingsCaller) OpenLimitOrders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Trader           common.Address
	PairIndex        *big.Int
	Index            *big.Int
	PositionSize     *big.Int
	SpreadReductionP *big.Int
	Buy              bool
	Leverage         *big.Int
	Tp               *big.Int
	Sl               *big.Int
	MinPrice         *big.Int
	MaxPrice         *big.Int
	Block            *big.Int
	TokenId          *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "openLimitOrders", arg0)

	outstruct := new(struct {
		Trader           common.Address
		PairIndex        *big.Int
		Index            *big.Int
		PositionSize     *big.Int
		SpreadReductionP *big.Int
		Buy              bool
		Leverage         *big.Int
		Tp               *big.Int
		Sl               *big.Int
		MinPrice         *big.Int
		MaxPrice         *big.Int
		Block            *big.Int
		TokenId          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Trader = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PairIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PositionSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SpreadReductionP = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Buy = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Leverage = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Tp = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Sl = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MinPrice = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.MaxPrice = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OpenLimitOrders is a free data retrieval call binding the contract method 0x32149592.
//
// Solidity: function openLimitOrders(uint256 ) view returns(address trader, uint256 pairIndex, uint256 index, uint256 positionSize, uint256 spreadReductionP, bool buy, uint256 leverage, uint256 tp, uint256 sl, uint256 minPrice, uint256 maxPrice, uint256 block, uint256 tokenId)
func (_Bindings *BindingsSession) OpenLimitOrders(arg0 *big.Int) (struct {
	Trader           common.Address
	PairIndex        *big.Int
	Index            *big.Int
	PositionSize     *big.Int
	SpreadReductionP *big.Int
	Buy              bool
	Leverage         *big.Int
	Tp               *big.Int
	Sl               *big.Int
	MinPrice         *big.Int
	MaxPrice         *big.Int
	Block            *big.Int
	TokenId          *big.Int
}, error) {
	return _Bindings.Contract.OpenLimitOrders(&_Bindings.CallOpts, arg0)
}

// OpenLimitOrders is a free data retrieval call binding the contract method 0x32149592.
//
// Solidity: function openLimitOrders(uint256 ) view returns(address trader, uint256 pairIndex, uint256 index, uint256 positionSize, uint256 spreadReductionP, bool buy, uint256 leverage, uint256 tp, uint256 sl, uint256 minPrice, uint256 maxPrice, uint256 block, uint256 tokenId)
func (_Bindings *BindingsCallerSession) OpenLimitOrders(arg0 *big.Int) (struct {
	Trader           common.Address
	PairIndex        *big.Int
	Index            *big.Int
	PositionSize     *big.Int
	SpreadReductionP *big.Int
	Buy              bool
	Leverage         *big.Int
	Tp               *big.Int
	Sl               *big.Int
	MinPrice         *big.Int
	MaxPrice         *big.Int
	Block            *big.Int
	TokenId          *big.Int
}, error) {
	return _Bindings.Contract.OpenLimitOrders(&_Bindings.CallOpts, arg0)
}

// OpenLimitOrdersCount is a free data retrieval call binding the contract method 0x97e7995a.
//
// Solidity: function openLimitOrdersCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) OpenLimitOrdersCount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "openLimitOrdersCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenLimitOrdersCount is a free data retrieval call binding the contract method 0x97e7995a.
//
// Solidity: function openLimitOrdersCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) OpenLimitOrdersCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenLimitOrdersCount(&_Bindings.CallOpts, arg0, arg1)
}

// OpenLimitOrdersCount is a free data retrieval call binding the contract method 0x97e7995a.
//
// Solidity: function openLimitOrdersCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) OpenLimitOrdersCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenLimitOrdersCount(&_Bindings.CallOpts, arg0, arg1)
}

// OpenTrades is a free data retrieval call binding the contract method 0xa3a80ffe.
//
// Solidity: function openTrades(address , uint256 , uint256 ) view returns(address trader, uint256 pairIndex, uint256 index, uint256 initialPosToken, uint256 positionSizeDai, uint256 openPrice, bool buy, uint256 leverage, uint256 tp, uint256 sl)
func (_Bindings *BindingsCaller) OpenTrades(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
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
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "openTrades", arg0, arg1, arg2)

	outstruct := new(struct {
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
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Trader = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PairIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InitialPosToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PositionSizeDai = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OpenPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Buy = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Leverage = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Tp = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Sl = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OpenTrades is a free data retrieval call binding the contract method 0xa3a80ffe.
//
// Solidity: function openTrades(address , uint256 , uint256 ) view returns(address trader, uint256 pairIndex, uint256 index, uint256 initialPosToken, uint256 positionSizeDai, uint256 openPrice, bool buy, uint256 leverage, uint256 tp, uint256 sl)
func (_Bindings *BindingsSession) OpenTrades(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
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
}, error) {
	return _Bindings.Contract.OpenTrades(&_Bindings.CallOpts, arg0, arg1, arg2)
}

// OpenTrades is a free data retrieval call binding the contract method 0xa3a80ffe.
//
// Solidity: function openTrades(address , uint256 , uint256 ) view returns(address trader, uint256 pairIndex, uint256 index, uint256 initialPosToken, uint256 positionSizeDai, uint256 openPrice, bool buy, uint256 leverage, uint256 tp, uint256 sl)
func (_Bindings *BindingsCallerSession) OpenTrades(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
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
}, error) {
	return _Bindings.Contract.OpenTrades(&_Bindings.CallOpts, arg0, arg1, arg2)
}

// OpenTradesCount is a free data retrieval call binding the contract method 0x1c8636b4.
//
// Solidity: function openTradesCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) OpenTradesCount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "openTradesCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenTradesCount is a free data retrieval call binding the contract method 0x1c8636b4.
//
// Solidity: function openTradesCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) OpenTradesCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenTradesCount(&_Bindings.CallOpts, arg0, arg1)
}

// OpenTradesCount is a free data retrieval call binding the contract method 0x1c8636b4.
//
// Solidity: function openTradesCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) OpenTradesCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.OpenTradesCount(&_Bindings.CallOpts, arg0, arg1)
}

// OpenTradesInfo is a free data retrieval call binding the contract method 0x8c8ff1d5.
//
// Solidity: function openTradesInfo(address , uint256 , uint256 ) view returns(uint256 tokenId, uint256 tokenPriceDai, uint256 openInterestDai, uint256 tpLastUpdated, uint256 slLastUpdated, bool beingMarketClosed)
func (_Bindings *BindingsCaller) OpenTradesInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	TokenId           *big.Int
	TokenPriceDai     *big.Int
	OpenInterestDai   *big.Int
	TpLastUpdated     *big.Int
	SlLastUpdated     *big.Int
	BeingMarketClosed bool
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "openTradesInfo", arg0, arg1, arg2)

	outstruct := new(struct {
		TokenId           *big.Int
		TokenPriceDai     *big.Int
		OpenInterestDai   *big.Int
		TpLastUpdated     *big.Int
		SlLastUpdated     *big.Int
		BeingMarketClosed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenPriceDai = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OpenInterestDai = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TpLastUpdated = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SlLastUpdated = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BeingMarketClosed = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// OpenTradesInfo is a free data retrieval call binding the contract method 0x8c8ff1d5.
//
// Solidity: function openTradesInfo(address , uint256 , uint256 ) view returns(uint256 tokenId, uint256 tokenPriceDai, uint256 openInterestDai, uint256 tpLastUpdated, uint256 slLastUpdated, bool beingMarketClosed)
func (_Bindings *BindingsSession) OpenTradesInfo(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	TokenId           *big.Int
	TokenPriceDai     *big.Int
	OpenInterestDai   *big.Int
	TpLastUpdated     *big.Int
	SlLastUpdated     *big.Int
	BeingMarketClosed bool
}, error) {
	return _Bindings.Contract.OpenTradesInfo(&_Bindings.CallOpts, arg0, arg1, arg2)
}

// OpenTradesInfo is a free data retrieval call binding the contract method 0x8c8ff1d5.
//
// Solidity: function openTradesInfo(address , uint256 , uint256 ) view returns(uint256 tokenId, uint256 tokenPriceDai, uint256 openInterestDai, uint256 tpLastUpdated, uint256 slLastUpdated, bool beingMarketClosed)
func (_Bindings *BindingsCallerSession) OpenTradesInfo(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	TokenId           *big.Int
	TokenPriceDai     *big.Int
	OpenInterestDai   *big.Int
	TpLastUpdated     *big.Int
	SlLastUpdated     *big.Int
	BeingMarketClosed bool
}, error) {
	return _Bindings.Contract.OpenTradesInfo(&_Bindings.CallOpts, arg0, arg1, arg2)
}

// PairTraders is a free data retrieval call binding the contract method 0xe3cbe3a7.
//
// Solidity: function pairTraders(uint256 , uint256 ) view returns(address)
func (_Bindings *BindingsCaller) PairTraders(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pairTraders", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairTraders is a free data retrieval call binding the contract method 0xe3cbe3a7.
//
// Solidity: function pairTraders(uint256 , uint256 ) view returns(address)
func (_Bindings *BindingsSession) PairTraders(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Bindings.Contract.PairTraders(&_Bindings.CallOpts, arg0, arg1)
}

// PairTraders is a free data retrieval call binding the contract method 0xe3cbe3a7.
//
// Solidity: function pairTraders(uint256 , uint256 ) view returns(address)
func (_Bindings *BindingsCallerSession) PairTraders(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Bindings.Contract.PairTraders(&_Bindings.CallOpts, arg0, arg1)
}

// PairTradersArray is a free data retrieval call binding the contract method 0x5fbfe8cc.
//
// Solidity: function pairTradersArray(uint256 _pairIndex) view returns(address[])
func (_Bindings *BindingsCaller) PairTradersArray(opts *bind.CallOpts, _pairIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pairTradersArray", _pairIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PairTradersArray is a free data retrieval call binding the contract method 0x5fbfe8cc.
//
// Solidity: function pairTradersArray(uint256 _pairIndex) view returns(address[])
func (_Bindings *BindingsSession) PairTradersArray(_pairIndex *big.Int) ([]common.Address, error) {
	return _Bindings.Contract.PairTradersArray(&_Bindings.CallOpts, _pairIndex)
}

// PairTradersArray is a free data retrieval call binding the contract method 0x5fbfe8cc.
//
// Solidity: function pairTradersArray(uint256 _pairIndex) view returns(address[])
func (_Bindings *BindingsCallerSession) PairTradersArray(_pairIndex *big.Int) ([]common.Address, error) {
	return _Bindings.Contract.PairTradersArray(&_Bindings.CallOpts, _pairIndex)
}

// PairTradersId is a free data retrieval call binding the contract method 0x4415b369.
//
// Solidity: function pairTradersId(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) PairTradersId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pairTradersId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairTradersId is a free data retrieval call binding the contract method 0x4415b369.
//
// Solidity: function pairTradersId(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) PairTradersId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PairTradersId(&_Bindings.CallOpts, arg0, arg1)
}

// PairTradersId is a free data retrieval call binding the contract method 0x4415b369.
//
// Solidity: function pairTradersId(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) PairTradersId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PairTradersId(&_Bindings.CallOpts, arg0, arg1)
}

// PendingMarketCloseCount is a free data retrieval call binding the contract method 0xb7682d54.
//
// Solidity: function pendingMarketCloseCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) PendingMarketCloseCount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pendingMarketCloseCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingMarketCloseCount is a free data retrieval call binding the contract method 0xb7682d54.
//
// Solidity: function pendingMarketCloseCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) PendingMarketCloseCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PendingMarketCloseCount(&_Bindings.CallOpts, arg0, arg1)
}

// PendingMarketCloseCount is a free data retrieval call binding the contract method 0xb7682d54.
//
// Solidity: function pendingMarketCloseCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) PendingMarketCloseCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PendingMarketCloseCount(&_Bindings.CallOpts, arg0, arg1)
}

// PendingMarketOpenCount is a free data retrieval call binding the contract method 0x9c8912f1.
//
// Solidity: function pendingMarketOpenCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) PendingMarketOpenCount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pendingMarketOpenCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingMarketOpenCount is a free data retrieval call binding the contract method 0x9c8912f1.
//
// Solidity: function pendingMarketOpenCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) PendingMarketOpenCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PendingMarketOpenCount(&_Bindings.CallOpts, arg0, arg1)
}

// PendingMarketOpenCount is a free data retrieval call binding the contract method 0x9c8912f1.
//
// Solidity: function pendingMarketOpenCount(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) PendingMarketOpenCount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PendingMarketOpenCount(&_Bindings.CallOpts, arg0, arg1)
}

// PendingOrderIds is a free data retrieval call binding the contract method 0xbcc9f1d3.
//
// Solidity: function pendingOrderIds(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) PendingOrderIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pendingOrderIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingOrderIds is a free data retrieval call binding the contract method 0xbcc9f1d3.
//
// Solidity: function pendingOrderIds(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) PendingOrderIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PendingOrderIds(&_Bindings.CallOpts, arg0, arg1)
}

// PendingOrderIds is a free data retrieval call binding the contract method 0xbcc9f1d3.
//
// Solidity: function pendingOrderIds(address , uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) PendingOrderIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.PendingOrderIds(&_Bindings.CallOpts, arg0, arg1)
}

// PendingOrderIdsCount is a free data retrieval call binding the contract method 0xf7b3c6fd.
//
// Solidity: function pendingOrderIdsCount(address _trader) view returns(uint256)
func (_Bindings *BindingsCaller) PendingOrderIdsCount(opts *bind.CallOpts, _trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pendingOrderIdsCount", _trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingOrderIdsCount is a free data retrieval call binding the contract method 0xf7b3c6fd.
//
// Solidity: function pendingOrderIdsCount(address _trader) view returns(uint256)
func (_Bindings *BindingsSession) PendingOrderIdsCount(_trader common.Address) (*big.Int, error) {
	return _Bindings.Contract.PendingOrderIdsCount(&_Bindings.CallOpts, _trader)
}

// PendingOrderIdsCount is a free data retrieval call binding the contract method 0xf7b3c6fd.
//
// Solidity: function pendingOrderIdsCount(address _trader) view returns(uint256)
func (_Bindings *BindingsCallerSession) PendingOrderIdsCount(_trader common.Address) (*big.Int, error) {
	return _Bindings.Contract.PendingOrderIdsCount(&_Bindings.CallOpts, _trader)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Bindings *BindingsCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Bindings *BindingsSession) Pool() (common.Address, error) {
	return _Bindings.Contract.Pool(&_Bindings.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Bindings *BindingsCallerSession) Pool() (common.Address, error) {
	return _Bindings.Contract.Pool(&_Bindings.CallOpts)
}

// PriceAggregator is a free data retrieval call binding the contract method 0x3078fff5.
//
// Solidity: function priceAggregator() view returns(address)
func (_Bindings *BindingsCaller) PriceAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "priceAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceAggregator is a free data retrieval call binding the contract method 0x3078fff5.
//
// Solidity: function priceAggregator() view returns(address)
func (_Bindings *BindingsSession) PriceAggregator() (common.Address, error) {
	return _Bindings.Contract.PriceAggregator(&_Bindings.CallOpts)
}

// PriceAggregator is a free data retrieval call binding the contract method 0x3078fff5.
//
// Solidity: function priceAggregator() view returns(address)
func (_Bindings *BindingsCallerSession) PriceAggregator() (common.Address, error) {
	return _Bindings.Contract.PriceAggregator(&_Bindings.CallOpts)
}

// ReqIDPendingMarketOrder is a free data retrieval call binding the contract method 0xb66e6b05.
//
// Solidity: function reqID_pendingMarketOrder(uint256 ) view returns((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) trade, uint256 block, uint256 wantedPrice, uint256 slippageP, uint256 spreadReductionP, uint256 tokenId)
func (_Bindings *BindingsCaller) ReqIDPendingMarketOrder(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Trade            GFarmTradingStorageV5Trade
	Block            *big.Int
	WantedPrice      *big.Int
	SlippageP        *big.Int
	SpreadReductionP *big.Int
	TokenId          *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "reqID_pendingMarketOrder", arg0)

	outstruct := new(struct {
		Trade GFarmTradingStorageV5Trade
		Block *big.Int
		WantedPrice      *big.Int
		SlippageP        *big.Int
		SpreadReductionP *big.Int
		TokenId          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Trade = *abi.ConvertType(out[0], new(GFarmTradingStorageV5Trade)).(*GFarmTradingStorageV5Trade)
	outstruct.Block = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WantedPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SlippageP = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SpreadReductionP = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReqIDPendingMarketOrder is a free data retrieval call binding the contract method 0xb66e6b05.
//
// Solidity: function reqID_pendingMarketOrder(uint256 ) view returns((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) trade, uint256 block, uint256 wantedPrice, uint256 slippageP, uint256 spreadReductionP, uint256 tokenId)
func (_Bindings *BindingsSession) ReqIDPendingMarketOrder(arg0 *big.Int) (struct {
	Trade            GFarmTradingStorageV5Trade
	Block            *big.Int
	WantedPrice      *big.Int
	SlippageP        *big.Int
	SpreadReductionP *big.Int
	TokenId          *big.Int
}, error) {
	return _Bindings.Contract.ReqIDPendingMarketOrder(&_Bindings.CallOpts, arg0)
}

// ReqIDPendingMarketOrder is a free data retrieval call binding the contract method 0xb66e6b05.
//
// Solidity: function reqID_pendingMarketOrder(uint256 ) view returns((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) trade, uint256 block, uint256 wantedPrice, uint256 slippageP, uint256 spreadReductionP, uint256 tokenId)
func (_Bindings *BindingsCallerSession) ReqIDPendingMarketOrder(arg0 *big.Int) (struct {
	Trade            GFarmTradingStorageV5Trade
	Block            *big.Int
	WantedPrice      *big.Int
	SlippageP        *big.Int
	SpreadReductionP *big.Int
	TokenId          *big.Int
}, error) {
	return _Bindings.Contract.ReqIDPendingMarketOrder(&_Bindings.CallOpts, arg0)
}

// ReqIDPendingNftOrder is a free data retrieval call binding the contract method 0x99794d35.
//
// Solidity: function reqID_pendingNftOrder(uint256 ) view returns(address nftHolder, uint256 nftId, address trader, uint256 pairIndex, uint256 index, uint8 orderType)
func (_Bindings *BindingsCaller) ReqIDPendingNftOrder(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftHolder common.Address
	NftId     *big.Int
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	OrderType uint8
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "reqID_pendingNftOrder", arg0)

	outstruct := new(struct {
		NftHolder common.Address
		NftId     *big.Int
		Trader    common.Address
		PairIndex *big.Int
		Index     *big.Int
		OrderType uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftHolder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Trader = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.PairIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OrderType = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// ReqIDPendingNftOrder is a free data retrieval call binding the contract method 0x99794d35.
//
// Solidity: function reqID_pendingNftOrder(uint256 ) view returns(address nftHolder, uint256 nftId, address trader, uint256 pairIndex, uint256 index, uint8 orderType)
func (_Bindings *BindingsSession) ReqIDPendingNftOrder(arg0 *big.Int) (struct {
	NftHolder common.Address
	NftId     *big.Int
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	OrderType uint8
}, error) {
	return _Bindings.Contract.ReqIDPendingNftOrder(&_Bindings.CallOpts, arg0)
}

// ReqIDPendingNftOrder is a free data retrieval call binding the contract method 0x99794d35.
//
// Solidity: function reqID_pendingNftOrder(uint256 ) view returns(address nftHolder, uint256 nftId, address trader, uint256 pairIndex, uint256 index, uint8 orderType)
func (_Bindings *BindingsCallerSession) ReqIDPendingNftOrder(arg0 *big.Int) (struct {
	NftHolder common.Address
	NftId     *big.Int
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	OrderType uint8
}, error) {
	return _Bindings.Contract.ReqIDPendingNftOrder(&_Bindings.CallOpts, arg0)
}

// SpreadReductionsP is a free data retrieval call binding the contract method 0x3dae10a0.
//
// Solidity: function spreadReductionsP(uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) SpreadReductionsP(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "spreadReductionsP", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadReductionsP is a free data retrieval call binding the contract method 0x3dae10a0.
//
// Solidity: function spreadReductionsP(uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) SpreadReductionsP(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.SpreadReductionsP(&_Bindings.CallOpts, arg0)
}

// SpreadReductionsP is a free data retrieval call binding the contract method 0x3dae10a0.
//
// Solidity: function spreadReductionsP(uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) SpreadReductionsP(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.SpreadReductionsP(&_Bindings.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0xc6255626.
//
// Solidity: function supportedTokens(uint256 ) view returns(address)
func (_Bindings *BindingsCaller) SupportedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0xc6255626.
//
// Solidity: function supportedTokens(uint256 ) view returns(address)
func (_Bindings *BindingsSession) SupportedTokens(arg0 *big.Int) (common.Address, error) {
	return _Bindings.Contract.SupportedTokens(&_Bindings.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0xc6255626.
//
// Solidity: function supportedTokens(uint256 ) view returns(address)
func (_Bindings *BindingsCallerSession) SupportedTokens(arg0 *big.Int) (common.Address, error) {
	return _Bindings.Contract.SupportedTokens(&_Bindings.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bindings *BindingsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bindings *BindingsSession) Token() (common.Address, error) {
	return _Bindings.Contract.Token(&_Bindings.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bindings *BindingsCallerSession) Token() (common.Address, error) {
	return _Bindings.Contract.Token(&_Bindings.CallOpts)
}

// TokenDaiRouter is a free data retrieval call binding the contract method 0xf2c13bdf.
//
// Solidity: function tokenDaiRouter() view returns(address)
func (_Bindings *BindingsCaller) TokenDaiRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tokenDaiRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenDaiRouter is a free data retrieval call binding the contract method 0xf2c13bdf.
//
// Solidity: function tokenDaiRouter() view returns(address)
func (_Bindings *BindingsSession) TokenDaiRouter() (common.Address, error) {
	return _Bindings.Contract.TokenDaiRouter(&_Bindings.CallOpts)
}

// TokenDaiRouter is a free data retrieval call binding the contract method 0xf2c13bdf.
//
// Solidity: function tokenDaiRouter() view returns(address)
func (_Bindings *BindingsCallerSession) TokenDaiRouter() (common.Address, error) {
	return _Bindings.Contract.TokenDaiRouter(&_Bindings.CallOpts)
}

// TokensBurned is a free data retrieval call binding the contract method 0xe7873b58.
//
// Solidity: function tokensBurned() view returns(uint256)
func (_Bindings *BindingsCaller) TokensBurned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tokensBurned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensBurned is a free data retrieval call binding the contract method 0xe7873b58.
//
// Solidity: function tokensBurned() view returns(uint256)
func (_Bindings *BindingsSession) TokensBurned() (*big.Int, error) {
	return _Bindings.Contract.TokensBurned(&_Bindings.CallOpts)
}

// TokensBurned is a free data retrieval call binding the contract method 0xe7873b58.
//
// Solidity: function tokensBurned() view returns(uint256)
func (_Bindings *BindingsCallerSession) TokensBurned() (*big.Int, error) {
	return _Bindings.Contract.TokensBurned(&_Bindings.CallOpts)
}

// TokensMinted is a free data retrieval call binding the contract method 0x6de9f32b.
//
// Solidity: function tokensMinted() view returns(uint256)
func (_Bindings *BindingsCaller) TokensMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tokensMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensMinted is a free data retrieval call binding the contract method 0x6de9f32b.
//
// Solidity: function tokensMinted() view returns(uint256)
func (_Bindings *BindingsSession) TokensMinted() (*big.Int, error) {
	return _Bindings.Contract.TokensMinted(&_Bindings.CallOpts)
}

// TokensMinted is a free data retrieval call binding the contract method 0x6de9f32b.
//
// Solidity: function tokensMinted() view returns(uint256)
func (_Bindings *BindingsCallerSession) TokensMinted() (*big.Int, error) {
	return _Bindings.Contract.TokensMinted(&_Bindings.CallOpts)
}

// Traders is a free data retrieval call binding the contract method 0x92a88fa2.
//
// Solidity: function traders(address ) view returns(uint256 leverageUnlocked, address referral, uint256 referralRewardsTotal)
func (_Bindings *BindingsCaller) Traders(opts *bind.CallOpts, arg0 common.Address) (struct {
	LeverageUnlocked     *big.Int
	Referral             common.Address
	ReferralRewardsTotal *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "traders", arg0)

	outstruct := new(struct {
		LeverageUnlocked     *big.Int
		Referral             common.Address
		ReferralRewardsTotal *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LeverageUnlocked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Referral = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReferralRewardsTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Traders is a free data retrieval call binding the contract method 0x92a88fa2.
//
// Solidity: function traders(address ) view returns(uint256 leverageUnlocked, address referral, uint256 referralRewardsTotal)
func (_Bindings *BindingsSession) Traders(arg0 common.Address) (struct {
	LeverageUnlocked     *big.Int
	Referral             common.Address
	ReferralRewardsTotal *big.Int
}, error) {
	return _Bindings.Contract.Traders(&_Bindings.CallOpts, arg0)
}

// Traders is a free data retrieval call binding the contract method 0x92a88fa2.
//
// Solidity: function traders(address ) view returns(uint256 leverageUnlocked, address referral, uint256 referralRewardsTotal)
func (_Bindings *BindingsCallerSession) Traders(arg0 common.Address) (struct {
	LeverageUnlocked     *big.Int
	Referral             common.Address
	ReferralRewardsTotal *big.Int
}, error) {
	return _Bindings.Contract.Traders(&_Bindings.CallOpts, arg0)
}

// TradesPerBlock is a free data retrieval call binding the contract method 0x52e5398f.
//
// Solidity: function tradesPerBlock(uint256 ) view returns(uint256)
func (_Bindings *BindingsCaller) TradesPerBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "tradesPerBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradesPerBlock is a free data retrieval call binding the contract method 0x52e5398f.
//
// Solidity: function tradesPerBlock(uint256 ) view returns(uint256)
func (_Bindings *BindingsSession) TradesPerBlock(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.TradesPerBlock(&_Bindings.CallOpts, arg0)
}

// TradesPerBlock is a free data retrieval call binding the contract method 0x52e5398f.
//
// Solidity: function tradesPerBlock(uint256 ) view returns(uint256)
func (_Bindings *BindingsCallerSession) TradesPerBlock(arg0 *big.Int) (*big.Int, error) {
	return _Bindings.Contract.TradesPerBlock(&_Bindings.CallOpts, arg0)
}

// Trading is a free data retrieval call binding the contract method 0xec44acf2.
//
// Solidity: function trading() view returns(address)
func (_Bindings *BindingsCaller) Trading(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "trading")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Trading is a free data retrieval call binding the contract method 0xec44acf2.
//
// Solidity: function trading() view returns(address)
func (_Bindings *BindingsSession) Trading() (common.Address, error) {
	return _Bindings.Contract.Trading(&_Bindings.CallOpts)
}

// Trading is a free data retrieval call binding the contract method 0xec44acf2.
//
// Solidity: function trading() view returns(address)
func (_Bindings *BindingsCallerSession) Trading() (common.Address, error) {
	return _Bindings.Contract.Trading(&_Bindings.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Bindings *BindingsCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Bindings *BindingsSession) Vault() (common.Address, error) {
	return _Bindings.Contract.Vault(&_Bindings.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Bindings *BindingsCallerSession) Vault() (common.Address, error) {
	return _Bindings.Contract.Vault(&_Bindings.CallOpts)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _token) returns()
func (_Bindings *BindingsTransactor) AddSupportedToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "addSupportedToken", _token)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _token) returns()
func (_Bindings *BindingsSession) AddSupportedToken(_token common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddSupportedToken(&_Bindings.TransactOpts, _token)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _token) returns()
func (_Bindings *BindingsTransactorSession) AddSupportedToken(_token common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddSupportedToken(&_Bindings.TransactOpts, _token)
}

// AddTradingContract is a paid mutator transaction binding the contract method 0xf823e2ae.
//
// Solidity: function addTradingContract(address _trading) returns()
func (_Bindings *BindingsTransactor) AddTradingContract(opts *bind.TransactOpts, _trading common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "addTradingContract", _trading)
}

// AddTradingContract is a paid mutator transaction binding the contract method 0xf823e2ae.
//
// Solidity: function addTradingContract(address _trading) returns()
func (_Bindings *BindingsSession) AddTradingContract(_trading common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddTradingContract(&_Bindings.TransactOpts, _trading)
}

// AddTradingContract is a paid mutator transaction binding the contract method 0xf823e2ae.
//
// Solidity: function addTradingContract(address _trading) returns()
func (_Bindings *BindingsTransactorSession) AddTradingContract(_trading common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.AddTradingContract(&_Bindings.TransactOpts, _trading)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Bindings *BindingsTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Bindings *BindingsSession) ClaimFees() (*types.Transaction, error) {
	return _Bindings.Contract.ClaimFees(&_Bindings.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Bindings *BindingsTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _Bindings.Contract.ClaimFees(&_Bindings.TransactOpts)
}

// DistributeLpRewards is a paid mutator transaction binding the contract method 0xe3f04eba.
//
// Solidity: function distributeLpRewards(uint256 _amount) returns()
func (_Bindings *BindingsTransactor) DistributeLpRewards(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "distributeLpRewards", _amount)
}

// DistributeLpRewards is a paid mutator transaction binding the contract method 0xe3f04eba.
//
// Solidity: function distributeLpRewards(uint256 _amount) returns()
func (_Bindings *BindingsSession) DistributeLpRewards(_amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.DistributeLpRewards(&_Bindings.TransactOpts, _amount)
}

// DistributeLpRewards is a paid mutator transaction binding the contract method 0xe3f04eba.
//
// Solidity: function distributeLpRewards(uint256 _amount) returns()
func (_Bindings *BindingsTransactorSession) DistributeLpRewards(_amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.DistributeLpRewards(&_Bindings.TransactOpts, _amount)
}

// HandleDevGovFees is a paid mutator transaction binding the contract method 0xcfbacc14.
//
// Solidity: function handleDevGovFees(uint256 _pairIndex, uint256 _leveragedPositionSize, bool _dai, bool _fullFee) returns(uint256 fee)
func (_Bindings *BindingsTransactor) HandleDevGovFees(opts *bind.TransactOpts, _pairIndex *big.Int, _leveragedPositionSize *big.Int, _dai bool, _fullFee bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "handleDevGovFees", _pairIndex, _leveragedPositionSize, _dai, _fullFee)
}

// HandleDevGovFees is a paid mutator transaction binding the contract method 0xcfbacc14.
//
// Solidity: function handleDevGovFees(uint256 _pairIndex, uint256 _leveragedPositionSize, bool _dai, bool _fullFee) returns(uint256 fee)
func (_Bindings *BindingsSession) HandleDevGovFees(_pairIndex *big.Int, _leveragedPositionSize *big.Int, _dai bool, _fullFee bool) (*types.Transaction, error) {
	return _Bindings.Contract.HandleDevGovFees(&_Bindings.TransactOpts, _pairIndex, _leveragedPositionSize, _dai, _fullFee)
}

// HandleDevGovFees is a paid mutator transaction binding the contract method 0xcfbacc14.
//
// Solidity: function handleDevGovFees(uint256 _pairIndex, uint256 _leveragedPositionSize, bool _dai, bool _fullFee) returns(uint256 fee)
func (_Bindings *BindingsTransactorSession) HandleDevGovFees(_pairIndex *big.Int, _leveragedPositionSize *big.Int, _dai bool, _fullFee bool) (*types.Transaction, error) {
	return _Bindings.Contract.HandleDevGovFees(&_Bindings.TransactOpts, _pairIndex, _leveragedPositionSize, _dai, _fullFee)
}

// HandleTokens is a paid mutator transaction binding the contract method 0x5378143e.
//
// Solidity: function handleTokens(address _a, uint256 _amount, bool _mint) returns()
func (_Bindings *BindingsTransactor) HandleTokens(opts *bind.TransactOpts, _a common.Address, _amount *big.Int, _mint bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "handleTokens", _a, _amount, _mint)
}

// HandleTokens is a paid mutator transaction binding the contract method 0x5378143e.
//
// Solidity: function handleTokens(address _a, uint256 _amount, bool _mint) returns()
func (_Bindings *BindingsSession) HandleTokens(_a common.Address, _amount *big.Int, _mint bool) (*types.Transaction, error) {
	return _Bindings.Contract.HandleTokens(&_Bindings.TransactOpts, _a, _amount, _mint)
}

// HandleTokens is a paid mutator transaction binding the contract method 0x5378143e.
//
// Solidity: function handleTokens(address _a, uint256 _amount, bool _mint) returns()
func (_Bindings *BindingsTransactorSession) HandleTokens(_a common.Address, _amount *big.Int, _mint bool) (*types.Transaction, error) {
	return _Bindings.Contract.HandleTokens(&_Bindings.TransactOpts, _a, _amount, _mint)
}

// IncreaseNftRewards is a paid mutator transaction binding the contract method 0x7395d79e.
//
// Solidity: function increaseNftRewards(uint256 _nftId, uint256 _amount) returns()
func (_Bindings *BindingsTransactor) IncreaseNftRewards(opts *bind.TransactOpts, _nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "increaseNftRewards", _nftId, _amount)
}

// IncreaseNftRewards is a paid mutator transaction binding the contract method 0x7395d79e.
//
// Solidity: function increaseNftRewards(uint256 _nftId, uint256 _amount) returns()
func (_Bindings *BindingsSession) IncreaseNftRewards(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.IncreaseNftRewards(&_Bindings.TransactOpts, _nftId, _amount)
}

// IncreaseNftRewards is a paid mutator transaction binding the contract method 0x7395d79e.
//
// Solidity: function increaseNftRewards(uint256 _nftId, uint256 _amount) returns()
func (_Bindings *BindingsTransactorSession) IncreaseNftRewards(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.IncreaseNftRewards(&_Bindings.TransactOpts, _nftId, _amount)
}

// IncreaseReferralRewards is a paid mutator transaction binding the contract method 0x21f7a30b.
//
// Solidity: function increaseReferralRewards(address _referral, uint256 _amount) returns()
func (_Bindings *BindingsTransactor) IncreaseReferralRewards(opts *bind.TransactOpts, _referral common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "increaseReferralRewards", _referral, _amount)
}

// IncreaseReferralRewards is a paid mutator transaction binding the contract method 0x21f7a30b.
//
// Solidity: function increaseReferralRewards(address _referral, uint256 _amount) returns()
func (_Bindings *BindingsSession) IncreaseReferralRewards(_referral common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.IncreaseReferralRewards(&_Bindings.TransactOpts, _referral, _amount)
}

// IncreaseReferralRewards is a paid mutator transaction binding the contract method 0x21f7a30b.
//
// Solidity: function increaseReferralRewards(address _referral, uint256 _amount) returns()
func (_Bindings *BindingsTransactorSession) IncreaseReferralRewards(_referral common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.IncreaseReferralRewards(&_Bindings.TransactOpts, _referral, _amount)
}

// RemoveTradingContract is a paid mutator transaction binding the contract method 0x5b621e32.
//
// Solidity: function removeTradingContract(address _trading) returns()
func (_Bindings *BindingsTransactor) RemoveTradingContract(opts *bind.TransactOpts, _trading common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "removeTradingContract", _trading)
}

// RemoveTradingContract is a paid mutator transaction binding the contract method 0x5b621e32.
//
// Solidity: function removeTradingContract(address _trading) returns()
func (_Bindings *BindingsSession) RemoveTradingContract(_trading common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveTradingContract(&_Bindings.TransactOpts, _trading)
}

// RemoveTradingContract is a paid mutator transaction binding the contract method 0x5b621e32.
//
// Solidity: function removeTradingContract(address _trading) returns()
func (_Bindings *BindingsTransactorSession) RemoveTradingContract(_trading common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveTradingContract(&_Bindings.TransactOpts, _trading)
}

// SetCallbacks is a paid mutator transaction binding the contract method 0xfe0fc8d6.
//
// Solidity: function setCallbacks(address _callbacks) returns()
func (_Bindings *BindingsTransactor) SetCallbacks(opts *bind.TransactOpts, _callbacks common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setCallbacks", _callbacks)
}

// SetCallbacks is a paid mutator transaction binding the contract method 0xfe0fc8d6.
//
// Solidity: function setCallbacks(address _callbacks) returns()
func (_Bindings *BindingsSession) SetCallbacks(_callbacks common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetCallbacks(&_Bindings.TransactOpts, _callbacks)
}

// SetCallbacks is a paid mutator transaction binding the contract method 0xfe0fc8d6.
//
// Solidity: function setCallbacks(address _callbacks) returns()
func (_Bindings *BindingsTransactorSession) SetCallbacks(_callbacks common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetCallbacks(&_Bindings.TransactOpts, _callbacks)
}

// SetDefaultLeverageUnlocked is a paid mutator transaction binding the contract method 0x43b8fab5.
//
// Solidity: function setDefaultLeverageUnlocked(uint256 _lev) returns()
func (_Bindings *BindingsTransactor) SetDefaultLeverageUnlocked(opts *bind.TransactOpts, _lev *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultLeverageUnlocked", _lev)
}

// SetDefaultLeverageUnlocked is a paid mutator transaction binding the contract method 0x43b8fab5.
//
// Solidity: function setDefaultLeverageUnlocked(uint256 _lev) returns()
func (_Bindings *BindingsSession) SetDefaultLeverageUnlocked(_lev *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultLeverageUnlocked(&_Bindings.TransactOpts, _lev)
}

// SetDefaultLeverageUnlocked is a paid mutator transaction binding the contract method 0x43b8fab5.
//
// Solidity: function setDefaultLeverageUnlocked(uint256 _lev) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultLeverageUnlocked(_lev *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultLeverageUnlocked(&_Bindings.TransactOpts, _lev)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _dev) returns()
func (_Bindings *BindingsTransactor) SetDev(opts *bind.TransactOpts, _dev common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDev", _dev)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _dev) returns()
func (_Bindings *BindingsSession) SetDev(_dev common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetDev(&_Bindings.TransactOpts, _dev)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _dev) returns()
func (_Bindings *BindingsTransactorSession) SetDev(_dev common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetDev(&_Bindings.TransactOpts, _dev)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Bindings *BindingsTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Bindings *BindingsSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetGov(&_Bindings.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Bindings *BindingsTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetGov(&_Bindings.TransactOpts, _gov)
}

// SetLeverageUnlocked is a paid mutator transaction binding the contract method 0x28b1f887.
//
// Solidity: function setLeverageUnlocked(address _trader, uint256 _newLeverage) returns()
func (_Bindings *BindingsTransactor) SetLeverageUnlocked(opts *bind.TransactOpts, _trader common.Address, _newLeverage *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setLeverageUnlocked", _trader, _newLeverage)
}

// SetLeverageUnlocked is a paid mutator transaction binding the contract method 0x28b1f887.
//
// Solidity: function setLeverageUnlocked(address _trader, uint256 _newLeverage) returns()
func (_Bindings *BindingsSession) SetLeverageUnlocked(_trader common.Address, _newLeverage *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetLeverageUnlocked(&_Bindings.TransactOpts, _trader, _newLeverage)
}

// SetLeverageUnlocked is a paid mutator transaction binding the contract method 0x28b1f887.
//
// Solidity: function setLeverageUnlocked(address _trader, uint256 _newLeverage) returns()
func (_Bindings *BindingsTransactorSession) SetLeverageUnlocked(_trader common.Address, _newLeverage *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetLeverageUnlocked(&_Bindings.TransactOpts, _trader, _newLeverage)
}

// SetMaxGainP is a paid mutator transaction binding the contract method 0x13b36ff6.
//
// Solidity: function setMaxGainP(uint256 _max) returns()
func (_Bindings *BindingsTransactor) SetMaxGainP(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxGainP", _max)
}

// SetMaxGainP is a paid mutator transaction binding the contract method 0x13b36ff6.
//
// Solidity: function setMaxGainP(uint256 _max) returns()
func (_Bindings *BindingsSession) SetMaxGainP(_max *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxGainP(&_Bindings.TransactOpts, _max)
}

// SetMaxGainP is a paid mutator transaction binding the contract method 0x13b36ff6.
//
// Solidity: function setMaxGainP(uint256 _max) returns()
func (_Bindings *BindingsTransactorSession) SetMaxGainP(_max *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxGainP(&_Bindings.TransactOpts, _max)
}

// SetMaxOpenInterestDai is a paid mutator transaction binding the contract method 0xe378b367.
//
// Solidity: function setMaxOpenInterestDai(uint256 _pairIndex, uint256 _newMaxOpenInterest) returns()
func (_Bindings *BindingsTransactor) SetMaxOpenInterestDai(opts *bind.TransactOpts, _pairIndex *big.Int, _newMaxOpenInterest *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxOpenInterestDai", _pairIndex, _newMaxOpenInterest)
}

// SetMaxOpenInterestDai is a paid mutator transaction binding the contract method 0xe378b367.
//
// Solidity: function setMaxOpenInterestDai(uint256 _pairIndex, uint256 _newMaxOpenInterest) returns()
func (_Bindings *BindingsSession) SetMaxOpenInterestDai(_pairIndex *big.Int, _newMaxOpenInterest *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxOpenInterestDai(&_Bindings.TransactOpts, _pairIndex, _newMaxOpenInterest)
}

// SetMaxOpenInterestDai is a paid mutator transaction binding the contract method 0xe378b367.
//
// Solidity: function setMaxOpenInterestDai(uint256 _pairIndex, uint256 _newMaxOpenInterest) returns()
func (_Bindings *BindingsTransactorSession) SetMaxOpenInterestDai(_pairIndex *big.Int, _newMaxOpenInterest *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxOpenInterestDai(&_Bindings.TransactOpts, _pairIndex, _newMaxOpenInterest)
}

// SetMaxPendingMarketOrders is a paid mutator transaction binding the contract method 0xecf56a1f.
//
// Solidity: function setMaxPendingMarketOrders(uint256 _maxPendingMarketOrders) returns()
func (_Bindings *BindingsTransactor) SetMaxPendingMarketOrders(opts *bind.TransactOpts, _maxPendingMarketOrders *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxPendingMarketOrders", _maxPendingMarketOrders)
}

// SetMaxPendingMarketOrders is a paid mutator transaction binding the contract method 0xecf56a1f.
//
// Solidity: function setMaxPendingMarketOrders(uint256 _maxPendingMarketOrders) returns()
func (_Bindings *BindingsSession) SetMaxPendingMarketOrders(_maxPendingMarketOrders *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxPendingMarketOrders(&_Bindings.TransactOpts, _maxPendingMarketOrders)
}

// SetMaxPendingMarketOrders is a paid mutator transaction binding the contract method 0xecf56a1f.
//
// Solidity: function setMaxPendingMarketOrders(uint256 _maxPendingMarketOrders) returns()
func (_Bindings *BindingsTransactorSession) SetMaxPendingMarketOrders(_maxPendingMarketOrders *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxPendingMarketOrders(&_Bindings.TransactOpts, _maxPendingMarketOrders)
}

// SetMaxSlP is a paid mutator transaction binding the contract method 0x099fa934.
//
// Solidity: function setMaxSlP(uint256 _max) returns()
func (_Bindings *BindingsTransactor) SetMaxSlP(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxSlP", _max)
}

// SetMaxSlP is a paid mutator transaction binding the contract method 0x099fa934.
//
// Solidity: function setMaxSlP(uint256 _max) returns()
func (_Bindings *BindingsSession) SetMaxSlP(_max *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxSlP(&_Bindings.TransactOpts, _max)
}

// SetMaxSlP is a paid mutator transaction binding the contract method 0x099fa934.
//
// Solidity: function setMaxSlP(uint256 _max) returns()
func (_Bindings *BindingsTransactorSession) SetMaxSlP(_max *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxSlP(&_Bindings.TransactOpts, _max)
}

// SetMaxTradesPerBlock is a paid mutator transaction binding the contract method 0x8bf7f5f9.
//
// Solidity: function setMaxTradesPerBlock(uint256 _maxTradesPerBlock) returns()
func (_Bindings *BindingsTransactor) SetMaxTradesPerBlock(opts *bind.TransactOpts, _maxTradesPerBlock *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxTradesPerBlock", _maxTradesPerBlock)
}

// SetMaxTradesPerBlock is a paid mutator transaction binding the contract method 0x8bf7f5f9.
//
// Solidity: function setMaxTradesPerBlock(uint256 _maxTradesPerBlock) returns()
func (_Bindings *BindingsSession) SetMaxTradesPerBlock(_maxTradesPerBlock *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxTradesPerBlock(&_Bindings.TransactOpts, _maxTradesPerBlock)
}

// SetMaxTradesPerBlock is a paid mutator transaction binding the contract method 0x8bf7f5f9.
//
// Solidity: function setMaxTradesPerBlock(uint256 _maxTradesPerBlock) returns()
func (_Bindings *BindingsTransactorSession) SetMaxTradesPerBlock(_maxTradesPerBlock *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxTradesPerBlock(&_Bindings.TransactOpts, _maxTradesPerBlock)
}

// SetMaxTradesPerPair is a paid mutator transaction binding the contract method 0xdc2123ff.
//
// Solidity: function setMaxTradesPerPair(uint256 _maxTradesPerPair) returns()
func (_Bindings *BindingsTransactor) SetMaxTradesPerPair(opts *bind.TransactOpts, _maxTradesPerPair *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxTradesPerPair", _maxTradesPerPair)
}

// SetMaxTradesPerPair is a paid mutator transaction binding the contract method 0xdc2123ff.
//
// Solidity: function setMaxTradesPerPair(uint256 _maxTradesPerPair) returns()
func (_Bindings *BindingsSession) SetMaxTradesPerPair(_maxTradesPerPair *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxTradesPerPair(&_Bindings.TransactOpts, _maxTradesPerPair)
}

// SetMaxTradesPerPair is a paid mutator transaction binding the contract method 0xdc2123ff.
//
// Solidity: function setMaxTradesPerPair(uint256 _maxTradesPerPair) returns()
func (_Bindings *BindingsTransactorSession) SetMaxTradesPerPair(_maxTradesPerPair *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxTradesPerPair(&_Bindings.TransactOpts, _maxTradesPerPair)
}

// SetNftSuccessTimelock is a paid mutator transaction binding the contract method 0xe212991a.
//
// Solidity: function setNftSuccessTimelock(uint256 _blocks) returns()
func (_Bindings *BindingsTransactor) SetNftSuccessTimelock(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setNftSuccessTimelock", _blocks)
}

// SetNftSuccessTimelock is a paid mutator transaction binding the contract method 0xe212991a.
//
// Solidity: function setNftSuccessTimelock(uint256 _blocks) returns()
func (_Bindings *BindingsSession) SetNftSuccessTimelock(_blocks *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetNftSuccessTimelock(&_Bindings.TransactOpts, _blocks)
}

// SetNftSuccessTimelock is a paid mutator transaction binding the contract method 0xe212991a.
//
// Solidity: function setNftSuccessTimelock(uint256 _blocks) returns()
func (_Bindings *BindingsTransactorSession) SetNftSuccessTimelock(_blocks *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetNftSuccessTimelock(&_Bindings.TransactOpts, _blocks)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_Bindings *BindingsTransactor) SetPool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPool", _pool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_Bindings *BindingsSession) SetPool(_pool common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPool(&_Bindings.TransactOpts, _pool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _pool) returns()
func (_Bindings *BindingsTransactorSession) SetPool(_pool common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPool(&_Bindings.TransactOpts, _pool)
}

// SetPriceAggregator is a paid mutator transaction binding the contract method 0x5070e837.
//
// Solidity: function setPriceAggregator(address _aggregator) returns()
func (_Bindings *BindingsTransactor) SetPriceAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPriceAggregator", _aggregator)
}

// SetPriceAggregator is a paid mutator transaction binding the contract method 0x5070e837.
//
// Solidity: function setPriceAggregator(address _aggregator) returns()
func (_Bindings *BindingsSession) SetPriceAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPriceAggregator(&_Bindings.TransactOpts, _aggregator)
}

// SetPriceAggregator is a paid mutator transaction binding the contract method 0x5070e837.
//
// Solidity: function setPriceAggregator(address _aggregator) returns()
func (_Bindings *BindingsTransactorSession) SetPriceAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPriceAggregator(&_Bindings.TransactOpts, _aggregator)
}

// SetSpreadReductionsP is a paid mutator transaction binding the contract method 0xf013d278.
//
// Solidity: function setSpreadReductionsP(uint256[5] _r) returns()
func (_Bindings *BindingsTransactor) SetSpreadReductionsP(opts *bind.TransactOpts, _r [5]*big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setSpreadReductionsP", _r)
}

// SetSpreadReductionsP is a paid mutator transaction binding the contract method 0xf013d278.
//
// Solidity: function setSpreadReductionsP(uint256[5] _r) returns()
func (_Bindings *BindingsSession) SetSpreadReductionsP(_r [5]*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetSpreadReductionsP(&_Bindings.TransactOpts, _r)
}

// SetSpreadReductionsP is a paid mutator transaction binding the contract method 0xf013d278.
//
// Solidity: function setSpreadReductionsP(uint256[5] _r) returns()
func (_Bindings *BindingsTransactorSession) SetSpreadReductionsP(_r [5]*big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetSpreadReductionsP(&_Bindings.TransactOpts, _r)
}

// SetTokenDaiRouter is a paid mutator transaction binding the contract method 0x4daf04b1.
//
// Solidity: function setTokenDaiRouter(address _tokenDaiRouter) returns()
func (_Bindings *BindingsTransactor) SetTokenDaiRouter(opts *bind.TransactOpts, _tokenDaiRouter common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setTokenDaiRouter", _tokenDaiRouter)
}

// SetTokenDaiRouter is a paid mutator transaction binding the contract method 0x4daf04b1.
//
// Solidity: function setTokenDaiRouter(address _tokenDaiRouter) returns()
func (_Bindings *BindingsSession) SetTokenDaiRouter(_tokenDaiRouter common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetTokenDaiRouter(&_Bindings.TransactOpts, _tokenDaiRouter)
}

// SetTokenDaiRouter is a paid mutator transaction binding the contract method 0x4daf04b1.
//
// Solidity: function setTokenDaiRouter(address _tokenDaiRouter) returns()
func (_Bindings *BindingsTransactorSession) SetTokenDaiRouter(_tokenDaiRouter common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetTokenDaiRouter(&_Bindings.TransactOpts, _tokenDaiRouter)
}

// SetTrading is a paid mutator transaction binding the contract method 0xa553906c.
//
// Solidity: function setTrading(address _trading) returns()
func (_Bindings *BindingsTransactor) SetTrading(opts *bind.TransactOpts, _trading common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setTrading", _trading)
}

// SetTrading is a paid mutator transaction binding the contract method 0xa553906c.
//
// Solidity: function setTrading(address _trading) returns()
func (_Bindings *BindingsSession) SetTrading(_trading common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetTrading(&_Bindings.TransactOpts, _trading)
}

// SetTrading is a paid mutator transaction binding the contract method 0xa553906c.
//
// Solidity: function setTrading(address _trading) returns()
func (_Bindings *BindingsTransactorSession) SetTrading(_trading common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetTrading(&_Bindings.TransactOpts, _trading)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_Bindings *BindingsTransactor) SetVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setVault", _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_Bindings *BindingsSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetVault(&_Bindings.TransactOpts, _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_Bindings *BindingsTransactorSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetVault(&_Bindings.TransactOpts, _vault)
}

// StoreOpenLimitOrder is a paid mutator transaction binding the contract method 0xd3b5fe70.
//
// Solidity: function storeOpenLimitOrder((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) o) returns()
func (_Bindings *BindingsTransactor) StoreOpenLimitOrder(opts *bind.TransactOpts, o GFarmTradingStorageV5OpenLimitOrder) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "storeOpenLimitOrder", o)
}

// StoreOpenLimitOrder is a paid mutator transaction binding the contract method 0xd3b5fe70.
//
// Solidity: function storeOpenLimitOrder((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) o) returns()
func (_Bindings *BindingsSession) StoreOpenLimitOrder(o GFarmTradingStorageV5OpenLimitOrder) (*types.Transaction, error) {
	return _Bindings.Contract.StoreOpenLimitOrder(&_Bindings.TransactOpts, o)
}

// StoreOpenLimitOrder is a paid mutator transaction binding the contract method 0xd3b5fe70.
//
// Solidity: function storeOpenLimitOrder((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) o) returns()
func (_Bindings *BindingsTransactorSession) StoreOpenLimitOrder(o GFarmTradingStorageV5OpenLimitOrder) (*types.Transaction, error) {
	return _Bindings.Contract.StoreOpenLimitOrder(&_Bindings.TransactOpts, o)
}

// StorePendingMarketOrder is a paid mutator transaction binding the contract method 0xc76faf78.
//
// Solidity: function storePendingMarketOrder(((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) _order, uint256 _id, bool _open) returns()
func (_Bindings *BindingsTransactor) StorePendingMarketOrder(opts *bind.TransactOpts, _order GFarmTradingStorageV5PendingMarketOrder, _id *big.Int, _open bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "storePendingMarketOrder", _order, _id, _open)
}

// StorePendingMarketOrder is a paid mutator transaction binding the contract method 0xc76faf78.
//
// Solidity: function storePendingMarketOrder(((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) _order, uint256 _id, bool _open) returns()
func (_Bindings *BindingsSession) StorePendingMarketOrder(_order GFarmTradingStorageV5PendingMarketOrder, _id *big.Int, _open bool) (*types.Transaction, error) {
	return _Bindings.Contract.StorePendingMarketOrder(&_Bindings.TransactOpts, _order, _id, _open)
}

// StorePendingMarketOrder is a paid mutator transaction binding the contract method 0xc76faf78.
//
// Solidity: function storePendingMarketOrder(((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) _order, uint256 _id, bool _open) returns()
func (_Bindings *BindingsTransactorSession) StorePendingMarketOrder(_order GFarmTradingStorageV5PendingMarketOrder, _id *big.Int, _open bool) (*types.Transaction, error) {
	return _Bindings.Contract.StorePendingMarketOrder(&_Bindings.TransactOpts, _order, _id, _open)
}

// StorePendingNftOrder is a paid mutator transaction binding the contract method 0x817fa1a4.
//
// Solidity: function storePendingNftOrder((address,uint256,address,uint256,uint256,uint8) _nftOrder, uint256 _orderId) returns()
func (_Bindings *BindingsTransactor) StorePendingNftOrder(opts *bind.TransactOpts, _nftOrder GFarmTradingStorageV5PendingNftOrder, _orderId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "storePendingNftOrder", _nftOrder, _orderId)
}

// StorePendingNftOrder is a paid mutator transaction binding the contract method 0x817fa1a4.
//
// Solidity: function storePendingNftOrder((address,uint256,address,uint256,uint256,uint8) _nftOrder, uint256 _orderId) returns()
func (_Bindings *BindingsSession) StorePendingNftOrder(_nftOrder GFarmTradingStorageV5PendingNftOrder, _orderId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.StorePendingNftOrder(&_Bindings.TransactOpts, _nftOrder, _orderId)
}

// StorePendingNftOrder is a paid mutator transaction binding the contract method 0x817fa1a4.
//
// Solidity: function storePendingNftOrder((address,uint256,address,uint256,uint256,uint8) _nftOrder, uint256 _orderId) returns()
func (_Bindings *BindingsTransactorSession) StorePendingNftOrder(_nftOrder GFarmTradingStorageV5PendingNftOrder, _orderId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.StorePendingNftOrder(&_Bindings.TransactOpts, _nftOrder, _orderId)
}

// StoreReferral is a paid mutator transaction binding the contract method 0x9d1a1073.
//
// Solidity: function storeReferral(address _trader, address _referral) returns()
func (_Bindings *BindingsTransactor) StoreReferral(opts *bind.TransactOpts, _trader common.Address, _referral common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "storeReferral", _trader, _referral)
}

// StoreReferral is a paid mutator transaction binding the contract method 0x9d1a1073.
//
// Solidity: function storeReferral(address _trader, address _referral) returns()
func (_Bindings *BindingsSession) StoreReferral(_trader common.Address, _referral common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.StoreReferral(&_Bindings.TransactOpts, _trader, _referral)
}

// StoreReferral is a paid mutator transaction binding the contract method 0x9d1a1073.
//
// Solidity: function storeReferral(address _trader, address _referral) returns()
func (_Bindings *BindingsTransactorSession) StoreReferral(_trader common.Address, _referral common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.StoreReferral(&_Bindings.TransactOpts, _trader, _referral)
}

// StoreTrade is a paid mutator transaction binding the contract method 0xc311ab31.
//
// Solidity: function storeTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) _trade, (uint256,uint256,uint256,uint256,uint256,bool) _tradeInfo) returns()
func (_Bindings *BindingsTransactor) StoreTrade(opts *bind.TransactOpts, _trade GFarmTradingStorageV5Trade, _tradeInfo GFarmTradingStorageV5TradeInfo) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "storeTrade", _trade, _tradeInfo)
}

// StoreTrade is a paid mutator transaction binding the contract method 0xc311ab31.
//
// Solidity: function storeTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) _trade, (uint256,uint256,uint256,uint256,uint256,bool) _tradeInfo) returns()
func (_Bindings *BindingsSession) StoreTrade(_trade GFarmTradingStorageV5Trade, _tradeInfo GFarmTradingStorageV5TradeInfo) (*types.Transaction, error) {
	return _Bindings.Contract.StoreTrade(&_Bindings.TransactOpts, _trade, _tradeInfo)
}

// StoreTrade is a paid mutator transaction binding the contract method 0xc311ab31.
//
// Solidity: function storeTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) _trade, (uint256,uint256,uint256,uint256,uint256,bool) _tradeInfo) returns()
func (_Bindings *BindingsTransactorSession) StoreTrade(_trade GFarmTradingStorageV5Trade, _tradeInfo GFarmTradingStorageV5TradeInfo) (*types.Transaction, error) {
	return _Bindings.Contract.StoreTrade(&_Bindings.TransactOpts, _trade, _tradeInfo)
}

// TransferDai is a paid mutator transaction binding the contract method 0x6e70e7b6.
//
// Solidity: function transferDai(address _from, address _to, uint256 _amount) returns()
func (_Bindings *BindingsTransactor) TransferDai(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferDai", _from, _to, _amount)
}

// TransferDai is a paid mutator transaction binding the contract method 0x6e70e7b6.
//
// Solidity: function transferDai(address _from, address _to, uint256 _amount) returns()
func (_Bindings *BindingsSession) TransferDai(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferDai(&_Bindings.TransactOpts, _from, _to, _amount)
}

// TransferDai is a paid mutator transaction binding the contract method 0x6e70e7b6.
//
// Solidity: function transferDai(address _from, address _to, uint256 _amount) returns()
func (_Bindings *BindingsTransactorSession) TransferDai(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferDai(&_Bindings.TransactOpts, _from, _to, _amount)
}

// TransferLinkToAggregator is a paid mutator transaction binding the contract method 0x6690a806.
//
// Solidity: function transferLinkToAggregator(address _from, uint256 _pairIndex, uint256 _leveragedPosDai) returns()
func (_Bindings *BindingsTransactor) TransferLinkToAggregator(opts *bind.TransactOpts, _from common.Address, _pairIndex *big.Int, _leveragedPosDai *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferLinkToAggregator", _from, _pairIndex, _leveragedPosDai)
}

// TransferLinkToAggregator is a paid mutator transaction binding the contract method 0x6690a806.
//
// Solidity: function transferLinkToAggregator(address _from, uint256 _pairIndex, uint256 _leveragedPosDai) returns()
func (_Bindings *BindingsSession) TransferLinkToAggregator(_from common.Address, _pairIndex *big.Int, _leveragedPosDai *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferLinkToAggregator(&_Bindings.TransactOpts, _from, _pairIndex, _leveragedPosDai)
}

// TransferLinkToAggregator is a paid mutator transaction binding the contract method 0x6690a806.
//
// Solidity: function transferLinkToAggregator(address _from, uint256 _pairIndex, uint256 _leveragedPosDai) returns()
func (_Bindings *BindingsTransactorSession) TransferLinkToAggregator(_from common.Address, _pairIndex *big.Int, _leveragedPosDai *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferLinkToAggregator(&_Bindings.TransactOpts, _from, _pairIndex, _leveragedPosDai)
}

// UnregisterOpenLimitOrder is a paid mutator transaction binding the contract method 0xb4b0f567.
//
// Solidity: function unregisterOpenLimitOrder(address _trader, uint256 _pairIndex, uint256 _index) returns()
func (_Bindings *BindingsTransactor) UnregisterOpenLimitOrder(opts *bind.TransactOpts, _trader common.Address, _pairIndex *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unregisterOpenLimitOrder", _trader, _pairIndex, _index)
}

// UnregisterOpenLimitOrder is a paid mutator transaction binding the contract method 0xb4b0f567.
//
// Solidity: function unregisterOpenLimitOrder(address _trader, uint256 _pairIndex, uint256 _index) returns()
func (_Bindings *BindingsSession) UnregisterOpenLimitOrder(_trader common.Address, _pairIndex *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterOpenLimitOrder(&_Bindings.TransactOpts, _trader, _pairIndex, _index)
}

// UnregisterOpenLimitOrder is a paid mutator transaction binding the contract method 0xb4b0f567.
//
// Solidity: function unregisterOpenLimitOrder(address _trader, uint256 _pairIndex, uint256 _index) returns()
func (_Bindings *BindingsTransactorSession) UnregisterOpenLimitOrder(_trader common.Address, _pairIndex *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterOpenLimitOrder(&_Bindings.TransactOpts, _trader, _pairIndex, _index)
}

// UnregisterPendingMarketOrder is a paid mutator transaction binding the contract method 0xcab94580.
//
// Solidity: function unregisterPendingMarketOrder(uint256 _id, bool _open) returns()
func (_Bindings *BindingsTransactor) UnregisterPendingMarketOrder(opts *bind.TransactOpts, _id *big.Int, _open bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unregisterPendingMarketOrder", _id, _open)
}

// UnregisterPendingMarketOrder is a paid mutator transaction binding the contract method 0xcab94580.
//
// Solidity: function unregisterPendingMarketOrder(uint256 _id, bool _open) returns()
func (_Bindings *BindingsSession) UnregisterPendingMarketOrder(_id *big.Int, _open bool) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterPendingMarketOrder(&_Bindings.TransactOpts, _id, _open)
}

// UnregisterPendingMarketOrder is a paid mutator transaction binding the contract method 0xcab94580.
//
// Solidity: function unregisterPendingMarketOrder(uint256 _id, bool _open) returns()
func (_Bindings *BindingsTransactorSession) UnregisterPendingMarketOrder(_id *big.Int, _open bool) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterPendingMarketOrder(&_Bindings.TransactOpts, _id, _open)
}

// UnregisterPendingNftOrder is a paid mutator transaction binding the contract method 0x201d8bdb.
//
// Solidity: function unregisterPendingNftOrder(uint256 _order) returns()
func (_Bindings *BindingsTransactor) UnregisterPendingNftOrder(opts *bind.TransactOpts, _order *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unregisterPendingNftOrder", _order)
}

// UnregisterPendingNftOrder is a paid mutator transaction binding the contract method 0x201d8bdb.
//
// Solidity: function unregisterPendingNftOrder(uint256 _order) returns()
func (_Bindings *BindingsSession) UnregisterPendingNftOrder(_order *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterPendingNftOrder(&_Bindings.TransactOpts, _order)
}

// UnregisterPendingNftOrder is a paid mutator transaction binding the contract method 0x201d8bdb.
//
// Solidity: function unregisterPendingNftOrder(uint256 _order) returns()
func (_Bindings *BindingsTransactorSession) UnregisterPendingNftOrder(_order *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterPendingNftOrder(&_Bindings.TransactOpts, _order)
}

// UnregisterTrade is a paid mutator transaction binding the contract method 0x90d112b9.
//
// Solidity: function unregisterTrade(address trader, uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsTransactor) UnregisterTrade(opts *bind.TransactOpts, trader common.Address, pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unregisterTrade", trader, pairIndex, index)
}

// UnregisterTrade is a paid mutator transaction binding the contract method 0x90d112b9.
//
// Solidity: function unregisterTrade(address trader, uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsSession) UnregisterTrade(trader common.Address, pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterTrade(&_Bindings.TransactOpts, trader, pairIndex, index)
}

// UnregisterTrade is a paid mutator transaction binding the contract method 0x90d112b9.
//
// Solidity: function unregisterTrade(address trader, uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsTransactorSession) UnregisterTrade(trader common.Address, pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnregisterTrade(&_Bindings.TransactOpts, trader, pairIndex, index)
}

// UpdateNfts is a paid mutator transaction binding the contract method 0x06aef7de.
//
// Solidity: function updateNfts(address[5] _nfts) returns()
func (_Bindings *BindingsTransactor) UpdateNfts(opts *bind.TransactOpts, _nfts [5]common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateNfts", _nfts)
}

// UpdateNfts is a paid mutator transaction binding the contract method 0x06aef7de.
//
// Solidity: function updateNfts(address[5] _nfts) returns()
func (_Bindings *BindingsSession) UpdateNfts(_nfts [5]common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateNfts(&_Bindings.TransactOpts, _nfts)
}

// UpdateNfts is a paid mutator transaction binding the contract method 0x06aef7de.
//
// Solidity: function updateNfts(address[5] _nfts) returns()
func (_Bindings *BindingsTransactorSession) UpdateNfts(_nfts [5]common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateNfts(&_Bindings.TransactOpts, _nfts)
}

// UpdateOpenLimitOrder is a paid mutator transaction binding the contract method 0xb767608a.
//
// Solidity: function updateOpenLimitOrder((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _o) returns()
func (_Bindings *BindingsTransactor) UpdateOpenLimitOrder(opts *bind.TransactOpts, _o GFarmTradingStorageV5OpenLimitOrder) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateOpenLimitOrder", _o)
}

// UpdateOpenLimitOrder is a paid mutator transaction binding the contract method 0xb767608a.
//
// Solidity: function updateOpenLimitOrder((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _o) returns()
func (_Bindings *BindingsSession) UpdateOpenLimitOrder(_o GFarmTradingStorageV5OpenLimitOrder) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateOpenLimitOrder(&_Bindings.TransactOpts, _o)
}

// UpdateOpenLimitOrder is a paid mutator transaction binding the contract method 0xb767608a.
//
// Solidity: function updateOpenLimitOrder((address,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) _o) returns()
func (_Bindings *BindingsTransactorSession) UpdateOpenLimitOrder(_o GFarmTradingStorageV5OpenLimitOrder) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateOpenLimitOrder(&_Bindings.TransactOpts, _o)
}

// UpdateSl is a paid mutator transaction binding the contract method 0xa3544181.
//
// Solidity: function updateSl(address _trader, uint256 _pairIndex, uint256 _index, uint256 _newSl) returns()
func (_Bindings *BindingsTransactor) UpdateSl(opts *bind.TransactOpts, _trader common.Address, _pairIndex *big.Int, _index *big.Int, _newSl *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateSl", _trader, _pairIndex, _index, _newSl)
}

// UpdateSl is a paid mutator transaction binding the contract method 0xa3544181.
//
// Solidity: function updateSl(address _trader, uint256 _pairIndex, uint256 _index, uint256 _newSl) returns()
func (_Bindings *BindingsSession) UpdateSl(_trader common.Address, _pairIndex *big.Int, _index *big.Int, _newSl *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateSl(&_Bindings.TransactOpts, _trader, _pairIndex, _index, _newSl)
}

// UpdateSl is a paid mutator transaction binding the contract method 0xa3544181.
//
// Solidity: function updateSl(address _trader, uint256 _pairIndex, uint256 _index, uint256 _newSl) returns()
func (_Bindings *BindingsTransactorSession) UpdateSl(_trader common.Address, _pairIndex *big.Int, _index *big.Int, _newSl *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateSl(&_Bindings.TransactOpts, _trader, _pairIndex, _index, _newSl)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x7ad3def2.
//
// Solidity: function updateToken(address _newToken) returns()
func (_Bindings *BindingsTransactor) UpdateToken(opts *bind.TransactOpts, _newToken common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateToken", _newToken)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x7ad3def2.
//
// Solidity: function updateToken(address _newToken) returns()
func (_Bindings *BindingsSession) UpdateToken(_newToken common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateToken(&_Bindings.TransactOpts, _newToken)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x7ad3def2.
//
// Solidity: function updateToken(address _newToken) returns()
func (_Bindings *BindingsTransactorSession) UpdateToken(_newToken common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateToken(&_Bindings.TransactOpts, _newToken)
}

// UpdateTp is a paid mutator transaction binding the contract method 0x7fdb96f3.
//
// Solidity: function updateTp(address _trader, uint256 _pairIndex, uint256 _index, uint256 _newTp) returns()
func (_Bindings *BindingsTransactor) UpdateTp(opts *bind.TransactOpts, _trader common.Address, _pairIndex *big.Int, _index *big.Int, _newTp *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateTp", _trader, _pairIndex, _index, _newTp)
}

// UpdateTp is a paid mutator transaction binding the contract method 0x7fdb96f3.
//
// Solidity: function updateTp(address _trader, uint256 _pairIndex, uint256 _index, uint256 _newTp) returns()
func (_Bindings *BindingsSession) UpdateTp(_trader common.Address, _pairIndex *big.Int, _index *big.Int, _newTp *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateTp(&_Bindings.TransactOpts, _trader, _pairIndex, _index, _newTp)
}

// UpdateTp is a paid mutator transaction binding the contract method 0x7fdb96f3.
//
// Solidity: function updateTp(address _trader, uint256 _pairIndex, uint256 _index, uint256 _newTp) returns()
func (_Bindings *BindingsTransactorSession) UpdateTp(_trader common.Address, _pairIndex *big.Int, _index *big.Int, _newTp *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateTp(&_Bindings.TransactOpts, _trader, _pairIndex, _index, _newTp)
}

// UpdateTrade is a paid mutator transaction binding the contract method 0xd4bda908.
//
// Solidity: function updateTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) _t) returns()
func (_Bindings *BindingsTransactor) UpdateTrade(opts *bind.TransactOpts, _t GFarmTradingStorageV5Trade) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateTrade", _t)
}

// UpdateTrade is a paid mutator transaction binding the contract method 0xd4bda908.
//
// Solidity: function updateTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) _t) returns()
func (_Bindings *BindingsSession) UpdateTrade(_t GFarmTradingStorageV5Trade) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateTrade(&_Bindings.TransactOpts, _t)
}

// UpdateTrade is a paid mutator transaction binding the contract method 0xd4bda908.
//
// Solidity: function updateTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) _t) returns()
func (_Bindings *BindingsTransactorSession) UpdateTrade(_t GFarmTradingStorageV5Trade) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateTrade(&_Bindings.TransactOpts, _t)
}

// BindingsAddressUpdatedIterator is returned from FilterAddressUpdated and is used to iterate over the raw logs and unpacked data for AddressUpdated events raised by the Bindings contract.
type BindingsAddressUpdatedIterator struct {
	Event *BindingsAddressUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAddressUpdated)
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
		it.Event = new(BindingsAddressUpdated)
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
func (it *BindingsAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAddressUpdated represents a AddressUpdated event raised by the Bindings contract.
type BindingsAddressUpdated struct {
	Name string
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressUpdated is a free log retrieval operation binding the contract event 0x943e9d45a11aaae5d87503e3bc248665d9807856e5cf2bdb4a988bee44422781.
//
// Solidity: event AddressUpdated(string name, address a)
func (_Bindings *BindingsFilterer) FilterAddressUpdated(opts *bind.FilterOpts) (*BindingsAddressUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AddressUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsAddressUpdatedIterator{contract: _Bindings.contract, event: "AddressUpdated", logs: logs, sub: sub}, nil
}

// WatchAddressUpdated is a free log subscription operation binding the contract event 0x943e9d45a11aaae5d87503e3bc248665d9807856e5cf2bdb4a988bee44422781.
//
// Solidity: event AddressUpdated(string name, address a)
func (_Bindings *BindingsFilterer) WatchAddressUpdated(opts *bind.WatchOpts, sink chan<- *BindingsAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAddressUpdated)
				if err := _Bindings.contract.UnpackLog(event, "AddressUpdated", log); err != nil {
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

// ParseAddressUpdated is a log parse operation binding the contract event 0x943e9d45a11aaae5d87503e3bc248665d9807856e5cf2bdb4a988bee44422781.
//
// Solidity: event AddressUpdated(string name, address a)
func (_Bindings *BindingsFilterer) ParseAddressUpdated(log types.Log) (*BindingsAddressUpdated, error) {
	event := new(BindingsAddressUpdated)
	if err := _Bindings.contract.UnpackLog(event, "AddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsNftsUpdatedIterator is returned from FilterNftsUpdated and is used to iterate over the raw logs and unpacked data for NftsUpdated events raised by the Bindings contract.
type BindingsNftsUpdatedIterator struct {
	Event *BindingsNftsUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsNftsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNftsUpdated)
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
		it.Event = new(BindingsNftsUpdated)
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
func (it *BindingsNftsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNftsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNftsUpdated represents a NftsUpdated event raised by the Bindings contract.
type BindingsNftsUpdated struct {
	Nfts [5]common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNftsUpdated is a free log retrieval operation binding the contract event 0x89a3bfeda71ede5f8735aa09127c97cd56aa3803356c52f7c30caf418238b464.
//
// Solidity: event NftsUpdated(address[5] nfts)
func (_Bindings *BindingsFilterer) FilterNftsUpdated(opts *bind.FilterOpts) (*BindingsNftsUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NftsUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsNftsUpdatedIterator{contract: _Bindings.contract, event: "NftsUpdated", logs: logs, sub: sub}, nil
}

// WatchNftsUpdated is a free log subscription operation binding the contract event 0x89a3bfeda71ede5f8735aa09127c97cd56aa3803356c52f7c30caf418238b464.
//
// Solidity: event NftsUpdated(address[5] nfts)
func (_Bindings *BindingsFilterer) WatchNftsUpdated(opts *bind.WatchOpts, sink chan<- *BindingsNftsUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NftsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNftsUpdated)
				if err := _Bindings.contract.UnpackLog(event, "NftsUpdated", log); err != nil {
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

// ParseNftsUpdated is a log parse operation binding the contract event 0x89a3bfeda71ede5f8735aa09127c97cd56aa3803356c52f7c30caf418238b464.
//
// Solidity: event NftsUpdated(address[5] nfts)
func (_Bindings *BindingsFilterer) ParseNftsUpdated(log types.Log) (*BindingsNftsUpdated, error) {
	event := new(BindingsNftsUpdated)
	if err := _Bindings.contract.UnpackLog(event, "NftsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsNumberUpdatedIterator is returned from FilterNumberUpdated and is used to iterate over the raw logs and unpacked data for NumberUpdated events raised by the Bindings contract.
type BindingsNumberUpdatedIterator struct {
	Event *BindingsNumberUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNumberUpdated)
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
		it.Event = new(BindingsNumberUpdated)
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
func (it *BindingsNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNumberUpdated represents a NumberUpdated event raised by the Bindings contract.
type BindingsNumberUpdated struct {
	Name  string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNumberUpdated is a free log retrieval operation binding the contract event 0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab.
//
// Solidity: event NumberUpdated(string name, uint256 value)
func (_Bindings *BindingsFilterer) FilterNumberUpdated(opts *bind.FilterOpts) (*BindingsNumberUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NumberUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsNumberUpdatedIterator{contract: _Bindings.contract, event: "NumberUpdated", logs: logs, sub: sub}, nil
}

// WatchNumberUpdated is a free log subscription operation binding the contract event 0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab.
//
// Solidity: event NumberUpdated(string name, uint256 value)
func (_Bindings *BindingsFilterer) WatchNumberUpdated(opts *bind.WatchOpts, sink chan<- *BindingsNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNumberUpdated)
				if err := _Bindings.contract.UnpackLog(event, "NumberUpdated", log); err != nil {
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

// ParseNumberUpdated is a log parse operation binding the contract event 0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab.
//
// Solidity: event NumberUpdated(string name, uint256 value)
func (_Bindings *BindingsFilterer) ParseNumberUpdated(log types.Log) (*BindingsNumberUpdated, error) {
	event := new(BindingsNumberUpdated)
	if err := _Bindings.contract.UnpackLog(event, "NumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsNumberUpdatedPairIterator is returned from FilterNumberUpdatedPair and is used to iterate over the raw logs and unpacked data for NumberUpdatedPair events raised by the Bindings contract.
type BindingsNumberUpdatedPairIterator struct {
	Event *BindingsNumberUpdatedPair // Event containing the contract specifics and raw log

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
func (it *BindingsNumberUpdatedPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNumberUpdatedPair)
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
		it.Event = new(BindingsNumberUpdatedPair)
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
func (it *BindingsNumberUpdatedPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNumberUpdatedPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNumberUpdatedPair represents a NumberUpdatedPair event raised by the Bindings contract.
type BindingsNumberUpdatedPair struct {
	Name      string
	PairIndex *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNumberUpdatedPair is a free log retrieval operation binding the contract event 0x957cd67ea49f3e1ee953cdd23a1ba1f11d973e842411220d5c24ef0e24e1c956.
//
// Solidity: event NumberUpdatedPair(string name, uint256 pairIndex, uint256 value)
func (_Bindings *BindingsFilterer) FilterNumberUpdatedPair(opts *bind.FilterOpts) (*BindingsNumberUpdatedPairIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NumberUpdatedPair")
	if err != nil {
		return nil, err
	}
	return &BindingsNumberUpdatedPairIterator{contract: _Bindings.contract, event: "NumberUpdatedPair", logs: logs, sub: sub}, nil
}

// WatchNumberUpdatedPair is a free log subscription operation binding the contract event 0x957cd67ea49f3e1ee953cdd23a1ba1f11d973e842411220d5c24ef0e24e1c956.
//
// Solidity: event NumberUpdatedPair(string name, uint256 pairIndex, uint256 value)
func (_Bindings *BindingsFilterer) WatchNumberUpdatedPair(opts *bind.WatchOpts, sink chan<- *BindingsNumberUpdatedPair) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NumberUpdatedPair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNumberUpdatedPair)
				if err := _Bindings.contract.UnpackLog(event, "NumberUpdatedPair", log); err != nil {
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

// ParseNumberUpdatedPair is a log parse operation binding the contract event 0x957cd67ea49f3e1ee953cdd23a1ba1f11d973e842411220d5c24ef0e24e1c956.
//
// Solidity: event NumberUpdatedPair(string name, uint256 pairIndex, uint256 value)
func (_Bindings *BindingsFilterer) ParseNumberUpdatedPair(log types.Log) (*BindingsNumberUpdatedPair, error) {
	event := new(BindingsNumberUpdatedPair)
	if err := _Bindings.contract.UnpackLog(event, "NumberUpdatedPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSpreadReductionsUpdatedIterator is returned from FilterSpreadReductionsUpdated and is used to iterate over the raw logs and unpacked data for SpreadReductionsUpdated events raised by the Bindings contract.
type BindingsSpreadReductionsUpdatedIterator struct {
	Event *BindingsSpreadReductionsUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsSpreadReductionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSpreadReductionsUpdated)
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
		it.Event = new(BindingsSpreadReductionsUpdated)
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
func (it *BindingsSpreadReductionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSpreadReductionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSpreadReductionsUpdated represents a SpreadReductionsUpdated event raised by the Bindings contract.
type BindingsSpreadReductionsUpdated struct {
	Arg0 [5]*big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSpreadReductionsUpdated is a free log retrieval operation binding the contract event 0xdaae5e047a75b6f38ad4e279cdefe910fc35d5a99a91b8303bd948c0a999372a.
//
// Solidity: event SpreadReductionsUpdated(uint256[5] arg0)
func (_Bindings *BindingsFilterer) FilterSpreadReductionsUpdated(opts *bind.FilterOpts) (*BindingsSpreadReductionsUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SpreadReductionsUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsSpreadReductionsUpdatedIterator{contract: _Bindings.contract, event: "SpreadReductionsUpdated", logs: logs, sub: sub}, nil
}

// WatchSpreadReductionsUpdated is a free log subscription operation binding the contract event 0xdaae5e047a75b6f38ad4e279cdefe910fc35d5a99a91b8303bd948c0a999372a.
//
// Solidity: event SpreadReductionsUpdated(uint256[5] arg0)
func (_Bindings *BindingsFilterer) WatchSpreadReductionsUpdated(opts *bind.WatchOpts, sink chan<- *BindingsSpreadReductionsUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SpreadReductionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSpreadReductionsUpdated)
				if err := _Bindings.contract.UnpackLog(event, "SpreadReductionsUpdated", log); err != nil {
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

// ParseSpreadReductionsUpdated is a log parse operation binding the contract event 0xdaae5e047a75b6f38ad4e279cdefe910fc35d5a99a91b8303bd948c0a999372a.
//
// Solidity: event SpreadReductionsUpdated(uint256[5] arg0)
func (_Bindings *BindingsFilterer) ParseSpreadReductionsUpdated(log types.Log) (*BindingsSpreadReductionsUpdated, error) {
	event := new(BindingsSpreadReductionsUpdated)
	if err := _Bindings.contract.UnpackLog(event, "SpreadReductionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSupportedTokenAddedIterator is returned from FilterSupportedTokenAdded and is used to iterate over the raw logs and unpacked data for SupportedTokenAdded events raised by the Bindings contract.
type BindingsSupportedTokenAddedIterator struct {
	Event *BindingsSupportedTokenAdded // Event containing the contract specifics and raw log

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
func (it *BindingsSupportedTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSupportedTokenAdded)
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
		it.Event = new(BindingsSupportedTokenAdded)
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
func (it *BindingsSupportedTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSupportedTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSupportedTokenAdded represents a SupportedTokenAdded event raised by the Bindings contract.
type BindingsSupportedTokenAdded struct {
	A   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSupportedTokenAdded is a free log retrieval operation binding the contract event 0xd1be2e90bd3d24839d9dd94ad871068e1f9688b02fa43f2a62c9975dfa9de2d7.
//
// Solidity: event SupportedTokenAdded(address a)
func (_Bindings *BindingsFilterer) FilterSupportedTokenAdded(opts *bind.FilterOpts) (*BindingsSupportedTokenAddedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SupportedTokenAdded")
	if err != nil {
		return nil, err
	}
	return &BindingsSupportedTokenAddedIterator{contract: _Bindings.contract, event: "SupportedTokenAdded", logs: logs, sub: sub}, nil
}

// WatchSupportedTokenAdded is a free log subscription operation binding the contract event 0xd1be2e90bd3d24839d9dd94ad871068e1f9688b02fa43f2a62c9975dfa9de2d7.
//
// Solidity: event SupportedTokenAdded(address a)
func (_Bindings *BindingsFilterer) WatchSupportedTokenAdded(opts *bind.WatchOpts, sink chan<- *BindingsSupportedTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SupportedTokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSupportedTokenAdded)
				if err := _Bindings.contract.UnpackLog(event, "SupportedTokenAdded", log); err != nil {
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

// ParseSupportedTokenAdded is a log parse operation binding the contract event 0xd1be2e90bd3d24839d9dd94ad871068e1f9688b02fa43f2a62c9975dfa9de2d7.
//
// Solidity: event SupportedTokenAdded(address a)
func (_Bindings *BindingsFilterer) ParseSupportedTokenAdded(log types.Log) (*BindingsSupportedTokenAdded, error) {
	event := new(BindingsSupportedTokenAdded)
	if err := _Bindings.contract.UnpackLog(event, "SupportedTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTradingContractAddedIterator is returned from FilterTradingContractAdded and is used to iterate over the raw logs and unpacked data for TradingContractAdded events raised by the Bindings contract.
type BindingsTradingContractAddedIterator struct {
	Event *BindingsTradingContractAdded // Event containing the contract specifics and raw log

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
func (it *BindingsTradingContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTradingContractAdded)
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
		it.Event = new(BindingsTradingContractAdded)
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
func (it *BindingsTradingContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTradingContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTradingContractAdded represents a TradingContractAdded event raised by the Bindings contract.
type BindingsTradingContractAdded struct {
	A   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTradingContractAdded is a free log retrieval operation binding the contract event 0xb692fa98cbec9bd705572b4c7e0927e3f7698e9fe526dcef6f40f3d4d4980ca3.
//
// Solidity: event TradingContractAdded(address a)
func (_Bindings *BindingsFilterer) FilterTradingContractAdded(opts *bind.FilterOpts) (*BindingsTradingContractAddedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TradingContractAdded")
	if err != nil {
		return nil, err
	}
	return &BindingsTradingContractAddedIterator{contract: _Bindings.contract, event: "TradingContractAdded", logs: logs, sub: sub}, nil
}

// WatchTradingContractAdded is a free log subscription operation binding the contract event 0xb692fa98cbec9bd705572b4c7e0927e3f7698e9fe526dcef6f40f3d4d4980ca3.
//
// Solidity: event TradingContractAdded(address a)
func (_Bindings *BindingsFilterer) WatchTradingContractAdded(opts *bind.WatchOpts, sink chan<- *BindingsTradingContractAdded) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TradingContractAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTradingContractAdded)
				if err := _Bindings.contract.UnpackLog(event, "TradingContractAdded", log); err != nil {
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

// ParseTradingContractAdded is a log parse operation binding the contract event 0xb692fa98cbec9bd705572b4c7e0927e3f7698e9fe526dcef6f40f3d4d4980ca3.
//
// Solidity: event TradingContractAdded(address a)
func (_Bindings *BindingsFilterer) ParseTradingContractAdded(log types.Log) (*BindingsTradingContractAdded, error) {
	event := new(BindingsTradingContractAdded)
	if err := _Bindings.contract.UnpackLog(event, "TradingContractAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTradingContractRemovedIterator is returned from FilterTradingContractRemoved and is used to iterate over the raw logs and unpacked data for TradingContractRemoved events raised by the Bindings contract.
type BindingsTradingContractRemovedIterator struct {
	Event *BindingsTradingContractRemoved // Event containing the contract specifics and raw log

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
func (it *BindingsTradingContractRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTradingContractRemoved)
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
		it.Event = new(BindingsTradingContractRemoved)
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
func (it *BindingsTradingContractRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTradingContractRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTradingContractRemoved represents a TradingContractRemoved event raised by the Bindings contract.
type BindingsTradingContractRemoved struct {
	A   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTradingContractRemoved is a free log retrieval operation binding the contract event 0xf9cf924d7d98e6e0a27ec523cd7214b0bbdfaba197018b33b41a702f5238190e.
//
// Solidity: event TradingContractRemoved(address a)
func (_Bindings *BindingsFilterer) FilterTradingContractRemoved(opts *bind.FilterOpts) (*BindingsTradingContractRemovedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TradingContractRemoved")
	if err != nil {
		return nil, err
	}
	return &BindingsTradingContractRemovedIterator{contract: _Bindings.contract, event: "TradingContractRemoved", logs: logs, sub: sub}, nil
}

// WatchTradingContractRemoved is a free log subscription operation binding the contract event 0xf9cf924d7d98e6e0a27ec523cd7214b0bbdfaba197018b33b41a702f5238190e.
//
// Solidity: event TradingContractRemoved(address a)
func (_Bindings *BindingsFilterer) WatchTradingContractRemoved(opts *bind.WatchOpts, sink chan<- *BindingsTradingContractRemoved) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TradingContractRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTradingContractRemoved)
				if err := _Bindings.contract.UnpackLog(event, "TradingContractRemoved", log); err != nil {
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

// ParseTradingContractRemoved is a log parse operation binding the contract event 0xf9cf924d7d98e6e0a27ec523cd7214b0bbdfaba197018b33b41a702f5238190e.
//
// Solidity: event TradingContractRemoved(address a)
func (_Bindings *BindingsFilterer) ParseTradingContractRemoved(log types.Log) (*BindingsTradingContractRemoved, error) {
	event := new(BindingsTradingContractRemoved)
	if err := _Bindings.contract.UnpackLog(event, "TradingContractRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
