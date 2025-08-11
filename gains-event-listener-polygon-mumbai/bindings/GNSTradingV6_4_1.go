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

// StorageInterfaceV5PendingMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type StorageInterfaceV5PendingMarketOrder struct {
	Trade StorageInterfaceV5Trade
	Block *big.Int
	WantedPrice      *big.Int
	SlippageP        *big.Int
	SpreadReductionP *big.Int
	TokenId          *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStorageInterfaceV5\",\"name\":\"_storageT\",\"type\":\"address\"},{\"internalType\":\"contractIGNSOracleRewardsV6_4_1\",\"name\":\"_oracleRewards\",\"type\":\"address\"},{\"internalType\":\"contractGNSPairInfosInterfaceV6\",\"name\":\"_pairInfos\",\"type\":\"address\"},{\"internalType\":\"contractGNSReferralsInterfaceV6_2\",\"name\":\"_referrals\",\"type\":\"address\"},{\"internalType\":\"contractGNSBorrowingFeesInterfaceV6_4\",\"name\":\"_borrowingFees\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPosDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marketOrdersTimeout\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"bypass\",\"type\":\"bool\"}],\"name\":\"BypassTriggerLinkUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structStorageInterfaceV5.Trade\",\"name\":\"trade\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slippageP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spreadReductionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structStorageInterfaceV5.PendingMarketOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"ChainlinkCallbackTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"CouldNotCloseTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"name\":\"Done\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"}],\"name\":\"MarketOrderInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"byPassesLinkCost\",\"type\":\"bool\"}],\"name\":\"NftOrderInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"OpenLimitCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"OpenLimitPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSl\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSlippageP\",\"type\":\"uint256\"}],\"name\":\"OpenLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSl\",\"type\":\"uint256\"}],\"name\":\"SlUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTp\",\"type\":\"uint256\"}],\"name\":\"TpUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_msgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowingFees\",\"outputs\":[{\"internalType\":\"contractGNSBorrowingFeesInterfaceV6_4\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bypassTriggerLink\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"cancelOpenLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"closeTradeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_order\",\"type\":\"uint256\"}],\"name\":\"closeTradeMarketTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"call_data\",\"type\":\"bytes\"}],\"name\":\"delegatedAction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"done\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"packed\",\"type\":\"uint256\"}],\"name\":\"executeNftOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketOrdersTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPosDai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialPosToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionSizeDai\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buy\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"}],\"internalType\":\"structStorageInterfaceV5.Trade\",\"name\":\"t\",\"type\":\"tuple\"},{\"internalType\":\"enumIGNSOracleRewardsV6_4_1.OpenLimitOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"slippageP\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"openTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_order\",\"type\":\"uint256\"}],\"name\":\"openTradeMarketTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleRewards\",\"outputs\":[{\"internalType\":\"contractIGNSOracleRewardsV6_4_1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairInfos\",\"outputs\":[{\"internalType\":\"contractGNSPairInfosInterfaceV6\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referrals\",\"outputs\":[{\"internalType\":\"contractGNSReferralsInterfaceV6_2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bypass\",\"type\":\"bool\"}],\"name\":\"setBypassTriggerLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMarketOrdersTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMaxPosDai\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageT\",\"outputs\":[{\"internalType\":\"contractStorageInterfaceV5\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSlippageP\",\"type\":\"uint256\"}],\"name\":\"updateOpenLimitOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newSl\",\"type\":\"uint256\"}],\"name\":\"updateSl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pairIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newTp\",\"type\":\"uint256\"}],\"name\":\"updateTp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// MsgSender is a free data retrieval call binding the contract method 0x119df25f.
//
// Solidity: function _msgSender() view returns(address)
func (_Bindings *BindingsCaller) MsgSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "_msgSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgSender is a free data retrieval call binding the contract method 0x119df25f.
//
// Solidity: function _msgSender() view returns(address)
func (_Bindings *BindingsSession) MsgSender() (common.Address, error) {
	return _Bindings.Contract.MsgSender(&_Bindings.CallOpts)
}

// MsgSender is a free data retrieval call binding the contract method 0x119df25f.
//
// Solidity: function _msgSender() view returns(address)
func (_Bindings *BindingsCallerSession) MsgSender() (common.Address, error) {
	return _Bindings.Contract.MsgSender(&_Bindings.CallOpts)
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

// BypassTriggerLink is a free data retrieval call binding the contract method 0x2c7c8f4e.
//
// Solidity: function bypassTriggerLink(address ) view returns(bool)
func (_Bindings *BindingsCaller) BypassTriggerLink(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "bypassTriggerLink", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BypassTriggerLink is a free data retrieval call binding the contract method 0x2c7c8f4e.
//
// Solidity: function bypassTriggerLink(address ) view returns(bool)
func (_Bindings *BindingsSession) BypassTriggerLink(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.BypassTriggerLink(&_Bindings.CallOpts, arg0)
}

// BypassTriggerLink is a free data retrieval call binding the contract method 0x2c7c8f4e.
//
// Solidity: function bypassTriggerLink(address ) view returns(bool)
func (_Bindings *BindingsCallerSession) BypassTriggerLink(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.BypassTriggerLink(&_Bindings.CallOpts, arg0)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(address)
func (_Bindings *BindingsCaller) Delegations(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "delegations", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(address)
func (_Bindings *BindingsSession) Delegations(arg0 common.Address) (common.Address, error) {
	return _Bindings.Contract.Delegations(&_Bindings.CallOpts, arg0)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(address)
func (_Bindings *BindingsCallerSession) Delegations(arg0 common.Address) (common.Address, error) {
	return _Bindings.Contract.Delegations(&_Bindings.CallOpts, arg0)
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

// MarketOrdersTimeout is a free data retrieval call binding the contract method 0x410c0b7c.
//
// Solidity: function marketOrdersTimeout() view returns(uint256)
func (_Bindings *BindingsCaller) MarketOrdersTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "marketOrdersTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketOrdersTimeout is a free data retrieval call binding the contract method 0x410c0b7c.
//
// Solidity: function marketOrdersTimeout() view returns(uint256)
func (_Bindings *BindingsSession) MarketOrdersTimeout() (*big.Int, error) {
	return _Bindings.Contract.MarketOrdersTimeout(&_Bindings.CallOpts)
}

// MarketOrdersTimeout is a free data retrieval call binding the contract method 0x410c0b7c.
//
// Solidity: function marketOrdersTimeout() view returns(uint256)
func (_Bindings *BindingsCallerSession) MarketOrdersTimeout() (*big.Int, error) {
	return _Bindings.Contract.MarketOrdersTimeout(&_Bindings.CallOpts)
}

// MaxPosDai is a free data retrieval call binding the contract method 0x279a165c.
//
// Solidity: function maxPosDai() view returns(uint256)
func (_Bindings *BindingsCaller) MaxPosDai(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "maxPosDai")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPosDai is a free data retrieval call binding the contract method 0x279a165c.
//
// Solidity: function maxPosDai() view returns(uint256)
func (_Bindings *BindingsSession) MaxPosDai() (*big.Int, error) {
	return _Bindings.Contract.MaxPosDai(&_Bindings.CallOpts)
}

// MaxPosDai is a free data retrieval call binding the contract method 0x279a165c.
//
// Solidity: function maxPosDai() view returns(uint256)
func (_Bindings *BindingsCallerSession) MaxPosDai() (*big.Int, error) {
	return _Bindings.Contract.MaxPosDai(&_Bindings.CallOpts)
}

// OracleRewards is a free data retrieval call binding the contract method 0x272300c9.
//
// Solidity: function oracleRewards() view returns(address)
func (_Bindings *BindingsCaller) OracleRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "oracleRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleRewards is a free data retrieval call binding the contract method 0x272300c9.
//
// Solidity: function oracleRewards() view returns(address)
func (_Bindings *BindingsSession) OracleRewards() (common.Address, error) {
	return _Bindings.Contract.OracleRewards(&_Bindings.CallOpts)
}

// OracleRewards is a free data retrieval call binding the contract method 0x272300c9.
//
// Solidity: function oracleRewards() view returns(address)
func (_Bindings *BindingsCallerSession) OracleRewards() (common.Address, error) {
	return _Bindings.Contract.OracleRewards(&_Bindings.CallOpts)
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

// CancelOpenLimitOrder is a paid mutator transaction binding the contract method 0xb9b6573a.
//
// Solidity: function cancelOpenLimitOrder(uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsTransactor) CancelOpenLimitOrder(opts *bind.TransactOpts, pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "cancelOpenLimitOrder", pairIndex, index)
}

// CancelOpenLimitOrder is a paid mutator transaction binding the contract method 0xb9b6573a.
//
// Solidity: function cancelOpenLimitOrder(uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsSession) CancelOpenLimitOrder(pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CancelOpenLimitOrder(&_Bindings.TransactOpts, pairIndex, index)
}

// CancelOpenLimitOrder is a paid mutator transaction binding the contract method 0xb9b6573a.
//
// Solidity: function cancelOpenLimitOrder(uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsTransactorSession) CancelOpenLimitOrder(pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CancelOpenLimitOrder(&_Bindings.TransactOpts, pairIndex, index)
}

// CloseTradeMarket is a paid mutator transaction binding the contract method 0xa2a3c0cb.
//
// Solidity: function closeTradeMarket(uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsTransactor) CloseTradeMarket(opts *bind.TransactOpts, pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "closeTradeMarket", pairIndex, index)
}

// CloseTradeMarket is a paid mutator transaction binding the contract method 0xa2a3c0cb.
//
// Solidity: function closeTradeMarket(uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsSession) CloseTradeMarket(pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CloseTradeMarket(&_Bindings.TransactOpts, pairIndex, index)
}

// CloseTradeMarket is a paid mutator transaction binding the contract method 0xa2a3c0cb.
//
// Solidity: function closeTradeMarket(uint256 pairIndex, uint256 index) returns()
func (_Bindings *BindingsTransactorSession) CloseTradeMarket(pairIndex *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CloseTradeMarket(&_Bindings.TransactOpts, pairIndex, index)
}

// CloseTradeMarketTimeout is a paid mutator transaction binding the contract method 0x990382cf.
//
// Solidity: function closeTradeMarketTimeout(uint256 _order) returns()
func (_Bindings *BindingsTransactor) CloseTradeMarketTimeout(opts *bind.TransactOpts, _order *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "closeTradeMarketTimeout", _order)
}

// CloseTradeMarketTimeout is a paid mutator transaction binding the contract method 0x990382cf.
//
// Solidity: function closeTradeMarketTimeout(uint256 _order) returns()
func (_Bindings *BindingsSession) CloseTradeMarketTimeout(_order *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CloseTradeMarketTimeout(&_Bindings.TransactOpts, _order)
}

// CloseTradeMarketTimeout is a paid mutator transaction binding the contract method 0x990382cf.
//
// Solidity: function closeTradeMarketTimeout(uint256 _order) returns()
func (_Bindings *BindingsTransactorSession) CloseTradeMarketTimeout(_order *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.CloseTradeMarketTimeout(&_Bindings.TransactOpts, _order)
}

// DelegatedAction is a paid mutator transaction binding the contract method 0x9a10cc32.
//
// Solidity: function delegatedAction(address trader, bytes call_data) returns(bytes)
func (_Bindings *BindingsTransactor) DelegatedAction(opts *bind.TransactOpts, trader common.Address, call_data []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "delegatedAction", trader, call_data)
}

// DelegatedAction is a paid mutator transaction binding the contract method 0x9a10cc32.
//
// Solidity: function delegatedAction(address trader, bytes call_data) returns(bytes)
func (_Bindings *BindingsSession) DelegatedAction(trader common.Address, call_data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.DelegatedAction(&_Bindings.TransactOpts, trader, call_data)
}

// DelegatedAction is a paid mutator transaction binding the contract method 0x9a10cc32.
//
// Solidity: function delegatedAction(address trader, bytes call_data) returns(bytes)
func (_Bindings *BindingsTransactorSession) DelegatedAction(trader common.Address, call_data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.DelegatedAction(&_Bindings.TransactOpts, trader, call_data)
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

// ExecuteNftOrder is a paid mutator transaction binding the contract method 0x6c53cc1c.
//
// Solidity: function executeNftOrder(uint256 packed) returns()
func (_Bindings *BindingsTransactor) ExecuteNftOrder(opts *bind.TransactOpts, packed *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "executeNftOrder", packed)
}

// ExecuteNftOrder is a paid mutator transaction binding the contract method 0x6c53cc1c.
//
// Solidity: function executeNftOrder(uint256 packed) returns()
func (_Bindings *BindingsSession) ExecuteNftOrder(packed *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteNftOrder(&_Bindings.TransactOpts, packed)
}

// ExecuteNftOrder is a paid mutator transaction binding the contract method 0x6c53cc1c.
//
// Solidity: function executeNftOrder(uint256 packed) returns()
func (_Bindings *BindingsTransactorSession) ExecuteNftOrder(packed *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ExecuteNftOrder(&_Bindings.TransactOpts, packed)
}

// OpenTrade is a paid mutator transaction binding the contract method 0xfb4b71bb.
//
// Solidity: function openTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, uint8 orderType, uint256 slippageP, address referrer) returns()
func (_Bindings *BindingsTransactor) OpenTrade(opts *bind.TransactOpts, t StorageInterfaceV5Trade, orderType uint8, slippageP *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "openTrade", t, orderType, slippageP, referrer)
}

// OpenTrade is a paid mutator transaction binding the contract method 0xfb4b71bb.
//
// Solidity: function openTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, uint8 orderType, uint256 slippageP, address referrer) returns()
func (_Bindings *BindingsSession) OpenTrade(t StorageInterfaceV5Trade, orderType uint8, slippageP *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.OpenTrade(&_Bindings.TransactOpts, t, orderType, slippageP, referrer)
}

// OpenTrade is a paid mutator transaction binding the contract method 0xfb4b71bb.
//
// Solidity: function openTrade((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256) t, uint8 orderType, uint256 slippageP, address referrer) returns()
func (_Bindings *BindingsTransactorSession) OpenTrade(t StorageInterfaceV5Trade, orderType uint8, slippageP *big.Int, referrer common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.OpenTrade(&_Bindings.TransactOpts, t, orderType, slippageP, referrer)
}

// OpenTradeMarketTimeout is a paid mutator transaction binding the contract method 0x876d3abd.
//
// Solidity: function openTradeMarketTimeout(uint256 _order) returns()
func (_Bindings *BindingsTransactor) OpenTradeMarketTimeout(opts *bind.TransactOpts, _order *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "openTradeMarketTimeout", _order)
}

// OpenTradeMarketTimeout is a paid mutator transaction binding the contract method 0x876d3abd.
//
// Solidity: function openTradeMarketTimeout(uint256 _order) returns()
func (_Bindings *BindingsSession) OpenTradeMarketTimeout(_order *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.OpenTradeMarketTimeout(&_Bindings.TransactOpts, _order)
}

// OpenTradeMarketTimeout is a paid mutator transaction binding the contract method 0x876d3abd.
//
// Solidity: function openTradeMarketTimeout(uint256 _order) returns()
func (_Bindings *BindingsTransactorSession) OpenTradeMarketTimeout(_order *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.OpenTradeMarketTimeout(&_Bindings.TransactOpts, _order)
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

// RemoveDelegate is a paid mutator transaction binding the contract method 0x36fb8b15.
//
// Solidity: function removeDelegate() returns()
func (_Bindings *BindingsTransactor) RemoveDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "removeDelegate")
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x36fb8b15.
//
// Solidity: function removeDelegate() returns()
func (_Bindings *BindingsSession) RemoveDelegate() (*types.Transaction, error) {
	return _Bindings.Contract.RemoveDelegate(&_Bindings.TransactOpts)
}

// RemoveDelegate is a paid mutator transaction binding the contract method 0x36fb8b15.
//
// Solidity: function removeDelegate() returns()
func (_Bindings *BindingsTransactorSession) RemoveDelegate() (*types.Transaction, error) {
	return _Bindings.Contract.RemoveDelegate(&_Bindings.TransactOpts)
}

// SetBypassTriggerLink is a paid mutator transaction binding the contract method 0x4278bafe.
//
// Solidity: function setBypassTriggerLink(address user, bool bypass) returns()
func (_Bindings *BindingsTransactor) SetBypassTriggerLink(opts *bind.TransactOpts, user common.Address, bypass bool) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setBypassTriggerLink", user, bypass)
}

// SetBypassTriggerLink is a paid mutator transaction binding the contract method 0x4278bafe.
//
// Solidity: function setBypassTriggerLink(address user, bool bypass) returns()
func (_Bindings *BindingsSession) SetBypassTriggerLink(user common.Address, bypass bool) (*types.Transaction, error) {
	return _Bindings.Contract.SetBypassTriggerLink(&_Bindings.TransactOpts, user, bypass)
}

// SetBypassTriggerLink is a paid mutator transaction binding the contract method 0x4278bafe.
//
// Solidity: function setBypassTriggerLink(address user, bool bypass) returns()
func (_Bindings *BindingsTransactorSession) SetBypassTriggerLink(user common.Address, bypass bool) (*types.Transaction, error) {
	return _Bindings.Contract.SetBypassTriggerLink(&_Bindings.TransactOpts, user, bypass)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address delegate) returns()
func (_Bindings *BindingsTransactor) SetDelegate(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDelegate", delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address delegate) returns()
func (_Bindings *BindingsSession) SetDelegate(delegate common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetDelegate(&_Bindings.TransactOpts, delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address delegate) returns()
func (_Bindings *BindingsTransactorSession) SetDelegate(delegate common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetDelegate(&_Bindings.TransactOpts, delegate)
}

// SetMarketOrdersTimeout is a paid mutator transaction binding the contract method 0xb81bfa05.
//
// Solidity: function setMarketOrdersTimeout(uint256 value) returns()
func (_Bindings *BindingsTransactor) SetMarketOrdersTimeout(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMarketOrdersTimeout", value)
}

// SetMarketOrdersTimeout is a paid mutator transaction binding the contract method 0xb81bfa05.
//
// Solidity: function setMarketOrdersTimeout(uint256 value) returns()
func (_Bindings *BindingsSession) SetMarketOrdersTimeout(value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMarketOrdersTimeout(&_Bindings.TransactOpts, value)
}

// SetMarketOrdersTimeout is a paid mutator transaction binding the contract method 0xb81bfa05.
//
// Solidity: function setMarketOrdersTimeout(uint256 value) returns()
func (_Bindings *BindingsTransactorSession) SetMarketOrdersTimeout(value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMarketOrdersTimeout(&_Bindings.TransactOpts, value)
}

// SetMaxPosDai is a paid mutator transaction binding the contract method 0x934d1cf4.
//
// Solidity: function setMaxPosDai(uint256 value) returns()
func (_Bindings *BindingsTransactor) SetMaxPosDai(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxPosDai", value)
}

// SetMaxPosDai is a paid mutator transaction binding the contract method 0x934d1cf4.
//
// Solidity: function setMaxPosDai(uint256 value) returns()
func (_Bindings *BindingsSession) SetMaxPosDai(value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxPosDai(&_Bindings.TransactOpts, value)
}

// SetMaxPosDai is a paid mutator transaction binding the contract method 0x934d1cf4.
//
// Solidity: function setMaxPosDai(uint256 value) returns()
func (_Bindings *BindingsTransactorSession) SetMaxPosDai(value *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxPosDai(&_Bindings.TransactOpts, value)
}

// UpdateOpenLimitOrder is a paid mutator transaction binding the contract method 0xc641558e.
//
// Solidity: function updateOpenLimitOrder(uint256 pairIndex, uint256 index, uint256 price, uint256 tp, uint256 sl, uint256 maxSlippageP) returns()
func (_Bindings *BindingsTransactor) UpdateOpenLimitOrder(opts *bind.TransactOpts, pairIndex *big.Int, index *big.Int, price *big.Int, tp *big.Int, sl *big.Int, maxSlippageP *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateOpenLimitOrder", pairIndex, index, price, tp, sl, maxSlippageP)
}

// UpdateOpenLimitOrder is a paid mutator transaction binding the contract method 0xc641558e.
//
// Solidity: function updateOpenLimitOrder(uint256 pairIndex, uint256 index, uint256 price, uint256 tp, uint256 sl, uint256 maxSlippageP) returns()
func (_Bindings *BindingsSession) UpdateOpenLimitOrder(pairIndex *big.Int, index *big.Int, price *big.Int, tp *big.Int, sl *big.Int, maxSlippageP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateOpenLimitOrder(&_Bindings.TransactOpts, pairIndex, index, price, tp, sl, maxSlippageP)
}

// UpdateOpenLimitOrder is a paid mutator transaction binding the contract method 0xc641558e.
//
// Solidity: function updateOpenLimitOrder(uint256 pairIndex, uint256 index, uint256 price, uint256 tp, uint256 sl, uint256 maxSlippageP) returns()
func (_Bindings *BindingsTransactorSession) UpdateOpenLimitOrder(pairIndex *big.Int, index *big.Int, price *big.Int, tp *big.Int, sl *big.Int, maxSlippageP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateOpenLimitOrder(&_Bindings.TransactOpts, pairIndex, index, price, tp, sl, maxSlippageP)
}

// UpdateSl is a paid mutator transaction binding the contract method 0xbe73fb99.
//
// Solidity: function updateSl(uint256 pairIndex, uint256 index, uint256 newSl) returns()
func (_Bindings *BindingsTransactor) UpdateSl(opts *bind.TransactOpts, pairIndex *big.Int, index *big.Int, newSl *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateSl", pairIndex, index, newSl)
}

// UpdateSl is a paid mutator transaction binding the contract method 0xbe73fb99.
//
// Solidity: function updateSl(uint256 pairIndex, uint256 index, uint256 newSl) returns()
func (_Bindings *BindingsSession) UpdateSl(pairIndex *big.Int, index *big.Int, newSl *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateSl(&_Bindings.TransactOpts, pairIndex, index, newSl)
}

// UpdateSl is a paid mutator transaction binding the contract method 0xbe73fb99.
//
// Solidity: function updateSl(uint256 pairIndex, uint256 index, uint256 newSl) returns()
func (_Bindings *BindingsTransactorSession) UpdateSl(pairIndex *big.Int, index *big.Int, newSl *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateSl(&_Bindings.TransactOpts, pairIndex, index, newSl)
}

// UpdateTp is a paid mutator transaction binding the contract method 0xd8defd15.
//
// Solidity: function updateTp(uint256 pairIndex, uint256 index, uint256 newTp) returns()
func (_Bindings *BindingsTransactor) UpdateTp(opts *bind.TransactOpts, pairIndex *big.Int, index *big.Int, newTp *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "updateTp", pairIndex, index, newTp)
}

// UpdateTp is a paid mutator transaction binding the contract method 0xd8defd15.
//
// Solidity: function updateTp(uint256 pairIndex, uint256 index, uint256 newTp) returns()
func (_Bindings *BindingsSession) UpdateTp(pairIndex *big.Int, index *big.Int, newTp *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateTp(&_Bindings.TransactOpts, pairIndex, index, newTp)
}

// UpdateTp is a paid mutator transaction binding the contract method 0xd8defd15.
//
// Solidity: function updateTp(uint256 pairIndex, uint256 index, uint256 newTp) returns()
func (_Bindings *BindingsTransactorSession) UpdateTp(pairIndex *big.Int, index *big.Int, newTp *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UpdateTp(&_Bindings.TransactOpts, pairIndex, index, newTp)
}

// BindingsBypassTriggerLinkUpdatedIterator is returned from FilterBypassTriggerLinkUpdated and is used to iterate over the raw logs and unpacked data for BypassTriggerLinkUpdated events raised by the Bindings contract.
type BindingsBypassTriggerLinkUpdatedIterator struct {
	Event *BindingsBypassTriggerLinkUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsBypassTriggerLinkUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBypassTriggerLinkUpdated)
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
		it.Event = new(BindingsBypassTriggerLinkUpdated)
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
func (it *BindingsBypassTriggerLinkUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBypassTriggerLinkUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBypassTriggerLinkUpdated represents a BypassTriggerLinkUpdated event raised by the Bindings contract.
type BindingsBypassTriggerLinkUpdated struct {
	User   common.Address
	Bypass bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBypassTriggerLinkUpdated is a free log retrieval operation binding the contract event 0x3c0f648064e21cceb91d918a80bed80a58c69474a28d50fe3d9975b72be97c44.
//
// Solidity: event BypassTriggerLinkUpdated(address user, bool bypass)
func (_Bindings *BindingsFilterer) FilterBypassTriggerLinkUpdated(opts *bind.FilterOpts) (*BindingsBypassTriggerLinkUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BypassTriggerLinkUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsBypassTriggerLinkUpdatedIterator{contract: _Bindings.contract, event: "BypassTriggerLinkUpdated", logs: logs, sub: sub}, nil
}

// WatchBypassTriggerLinkUpdated is a free log subscription operation binding the contract event 0x3c0f648064e21cceb91d918a80bed80a58c69474a28d50fe3d9975b72be97c44.
//
// Solidity: event BypassTriggerLinkUpdated(address user, bool bypass)
func (_Bindings *BindingsFilterer) WatchBypassTriggerLinkUpdated(opts *bind.WatchOpts, sink chan<- *BindingsBypassTriggerLinkUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BypassTriggerLinkUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBypassTriggerLinkUpdated)
				if err := _Bindings.contract.UnpackLog(event, "BypassTriggerLinkUpdated", log); err != nil {
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

// ParseBypassTriggerLinkUpdated is a log parse operation binding the contract event 0x3c0f648064e21cceb91d918a80bed80a58c69474a28d50fe3d9975b72be97c44.
//
// Solidity: event BypassTriggerLinkUpdated(address user, bool bypass)
func (_Bindings *BindingsFilterer) ParseBypassTriggerLinkUpdated(log types.Log) (*BindingsBypassTriggerLinkUpdated, error) {
	event := new(BindingsBypassTriggerLinkUpdated)
	if err := _Bindings.contract.UnpackLog(event, "BypassTriggerLinkUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsChainlinkCallbackTimeoutIterator is returned from FilterChainlinkCallbackTimeout and is used to iterate over the raw logs and unpacked data for ChainlinkCallbackTimeout events raised by the Bindings contract.
type BindingsChainlinkCallbackTimeoutIterator struct {
	Event *BindingsChainlinkCallbackTimeout // Event containing the contract specifics and raw log

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
func (it *BindingsChainlinkCallbackTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsChainlinkCallbackTimeout)
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
		it.Event = new(BindingsChainlinkCallbackTimeout)
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
func (it *BindingsChainlinkCallbackTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsChainlinkCallbackTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsChainlinkCallbackTimeout represents a ChainlinkCallbackTimeout event raised by the Bindings contract.
type BindingsChainlinkCallbackTimeout struct {
	OrderId *big.Int
	Order   StorageInterfaceV5PendingMarketOrder
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCallbackTimeout is a free log retrieval operation binding the contract event 0x3adaa586cdbe84dd24e45bd7dada6da933d7c2d1c7b4e4cd02fce033356decb1.
//
// Solidity: event ChainlinkCallbackTimeout(uint256 indexed orderId, ((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) order)
func (_Bindings *BindingsFilterer) FilterChainlinkCallbackTimeout(opts *bind.FilterOpts, orderId []*big.Int) (*BindingsChainlinkCallbackTimeoutIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ChainlinkCallbackTimeout", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsChainlinkCallbackTimeoutIterator{contract: _Bindings.contract, event: "ChainlinkCallbackTimeout", logs: logs, sub: sub}, nil
}

// WatchChainlinkCallbackTimeout is a free log subscription operation binding the contract event 0x3adaa586cdbe84dd24e45bd7dada6da933d7c2d1c7b4e4cd02fce033356decb1.
//
// Solidity: event ChainlinkCallbackTimeout(uint256 indexed orderId, ((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) order)
func (_Bindings *BindingsFilterer) WatchChainlinkCallbackTimeout(opts *bind.WatchOpts, sink chan<- *BindingsChainlinkCallbackTimeout, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ChainlinkCallbackTimeout", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsChainlinkCallbackTimeout)
				if err := _Bindings.contract.UnpackLog(event, "ChainlinkCallbackTimeout", log); err != nil {
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

// ParseChainlinkCallbackTimeout is a log parse operation binding the contract event 0x3adaa586cdbe84dd24e45bd7dada6da933d7c2d1c7b4e4cd02fce033356decb1.
//
// Solidity: event ChainlinkCallbackTimeout(uint256 indexed orderId, ((address,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256),uint256,uint256,uint256,uint256,uint256) order)
func (_Bindings *BindingsFilterer) ParseChainlinkCallbackTimeout(log types.Log) (*BindingsChainlinkCallbackTimeout, error) {
	event := new(BindingsChainlinkCallbackTimeout)
	if err := _Bindings.contract.UnpackLog(event, "ChainlinkCallbackTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsCouldNotCloseTradeIterator is returned from FilterCouldNotCloseTrade and is used to iterate over the raw logs and unpacked data for CouldNotCloseTrade events raised by the Bindings contract.
type BindingsCouldNotCloseTradeIterator struct {
	Event *BindingsCouldNotCloseTrade // Event containing the contract specifics and raw log

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
func (it *BindingsCouldNotCloseTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsCouldNotCloseTrade)
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
		it.Event = new(BindingsCouldNotCloseTrade)
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
func (it *BindingsCouldNotCloseTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsCouldNotCloseTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsCouldNotCloseTrade represents a CouldNotCloseTrade event raised by the Bindings contract.
type BindingsCouldNotCloseTrade struct {
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCouldNotCloseTrade is a free log retrieval operation binding the contract event 0x60e497734ddabcd7293fd91739aaf65cf525eb539c97be528125a235a89288d8.
//
// Solidity: event CouldNotCloseTrade(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) FilterCouldNotCloseTrade(opts *bind.FilterOpts, trader []common.Address, pairIndex []*big.Int) (*BindingsCouldNotCloseTradeIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "CouldNotCloseTrade", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsCouldNotCloseTradeIterator{contract: _Bindings.contract, event: "CouldNotCloseTrade", logs: logs, sub: sub}, nil
}

// WatchCouldNotCloseTrade is a free log subscription operation binding the contract event 0x60e497734ddabcd7293fd91739aaf65cf525eb539c97be528125a235a89288d8.
//
// Solidity: event CouldNotCloseTrade(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) WatchCouldNotCloseTrade(opts *bind.WatchOpts, sink chan<- *BindingsCouldNotCloseTrade, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "CouldNotCloseTrade", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsCouldNotCloseTrade)
				if err := _Bindings.contract.UnpackLog(event, "CouldNotCloseTrade", log); err != nil {
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

// ParseCouldNotCloseTrade is a log parse operation binding the contract event 0x60e497734ddabcd7293fd91739aaf65cf525eb539c97be528125a235a89288d8.
//
// Solidity: event CouldNotCloseTrade(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) ParseCouldNotCloseTrade(log types.Log) (*BindingsCouldNotCloseTrade, error) {
	event := new(BindingsCouldNotCloseTrade)
	if err := _Bindings.contract.UnpackLog(event, "CouldNotCloseTrade", log); err != nil {
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

// BindingsMarketOrderInitiatedIterator is returned from FilterMarketOrderInitiated and is used to iterate over the raw logs and unpacked data for MarketOrderInitiated events raised by the Bindings contract.
type BindingsMarketOrderInitiatedIterator struct {
	Event *BindingsMarketOrderInitiated // Event containing the contract specifics and raw log

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
func (it *BindingsMarketOrderInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMarketOrderInitiated)
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
		it.Event = new(BindingsMarketOrderInitiated)
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
func (it *BindingsMarketOrderInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMarketOrderInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMarketOrderInitiated represents a MarketOrderInitiated event raised by the Bindings contract.
type BindingsMarketOrderInitiated struct {
	OrderId   *big.Int
	Trader    common.Address
	PairIndex *big.Int
	Open      bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMarketOrderInitiated is a free log retrieval operation binding the contract event 0x3e544118c04e3bb18b669475695cd270ba0e41fb13177483f01c14222de62a86.
//
// Solidity: event MarketOrderInitiated(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, bool open)
func (_Bindings *BindingsFilterer) FilterMarketOrderInitiated(opts *bind.FilterOpts, orderId []*big.Int, trader []common.Address, pairIndex []*big.Int) (*BindingsMarketOrderInitiatedIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MarketOrderInitiated", orderIdRule, traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMarketOrderInitiatedIterator{contract: _Bindings.contract, event: "MarketOrderInitiated", logs: logs, sub: sub}, nil
}

// WatchMarketOrderInitiated is a free log subscription operation binding the contract event 0x3e544118c04e3bb18b669475695cd270ba0e41fb13177483f01c14222de62a86.
//
// Solidity: event MarketOrderInitiated(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, bool open)
func (_Bindings *BindingsFilterer) WatchMarketOrderInitiated(opts *bind.WatchOpts, sink chan<- *BindingsMarketOrderInitiated, orderId []*big.Int, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MarketOrderInitiated", orderIdRule, traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMarketOrderInitiated)
				if err := _Bindings.contract.UnpackLog(event, "MarketOrderInitiated", log); err != nil {
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

// ParseMarketOrderInitiated is a log parse operation binding the contract event 0x3e544118c04e3bb18b669475695cd270ba0e41fb13177483f01c14222de62a86.
//
// Solidity: event MarketOrderInitiated(uint256 indexed orderId, address indexed trader, uint256 indexed pairIndex, bool open)
func (_Bindings *BindingsFilterer) ParseMarketOrderInitiated(log types.Log) (*BindingsMarketOrderInitiated, error) {
	event := new(BindingsMarketOrderInitiated)
	if err := _Bindings.contract.UnpackLog(event, "MarketOrderInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsNftOrderInitiatedIterator is returned from FilterNftOrderInitiated and is used to iterate over the raw logs and unpacked data for NftOrderInitiated events raised by the Bindings contract.
type BindingsNftOrderInitiatedIterator struct {
	Event *BindingsNftOrderInitiated // Event containing the contract specifics and raw log

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
func (it *BindingsNftOrderInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNftOrderInitiated)
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
		it.Event = new(BindingsNftOrderInitiated)
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
func (it *BindingsNftOrderInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNftOrderInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNftOrderInitiated represents a NftOrderInitiated event raised by the Bindings contract.
type BindingsNftOrderInitiated struct {
	OrderId          *big.Int
	Trader           common.Address
	PairIndex        *big.Int
	ByPassesLinkCost bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNftOrderInitiated is a free log retrieval operation binding the contract event 0x50a583b02839381dff332433f1a37825291992d796b87483d7c51649ef504d43.
//
// Solidity: event NftOrderInitiated(uint256 orderId, address indexed trader, uint256 indexed pairIndex, bool byPassesLinkCost)
func (_Bindings *BindingsFilterer) FilterNftOrderInitiated(opts *bind.FilterOpts, trader []common.Address, pairIndex []*big.Int) (*BindingsNftOrderInitiatedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NftOrderInitiated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsNftOrderInitiatedIterator{contract: _Bindings.contract, event: "NftOrderInitiated", logs: logs, sub: sub}, nil
}

// WatchNftOrderInitiated is a free log subscription operation binding the contract event 0x50a583b02839381dff332433f1a37825291992d796b87483d7c51649ef504d43.
//
// Solidity: event NftOrderInitiated(uint256 orderId, address indexed trader, uint256 indexed pairIndex, bool byPassesLinkCost)
func (_Bindings *BindingsFilterer) WatchNftOrderInitiated(opts *bind.WatchOpts, sink chan<- *BindingsNftOrderInitiated, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NftOrderInitiated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNftOrderInitiated)
				if err := _Bindings.contract.UnpackLog(event, "NftOrderInitiated", log); err != nil {
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

// ParseNftOrderInitiated is a log parse operation binding the contract event 0x50a583b02839381dff332433f1a37825291992d796b87483d7c51649ef504d43.
//
// Solidity: event NftOrderInitiated(uint256 orderId, address indexed trader, uint256 indexed pairIndex, bool byPassesLinkCost)
func (_Bindings *BindingsFilterer) ParseNftOrderInitiated(log types.Log) (*BindingsNftOrderInitiated, error) {
	event := new(BindingsNftOrderInitiated)
	if err := _Bindings.contract.UnpackLog(event, "NftOrderInitiated", log); err != nil {
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

// BindingsOpenLimitCanceledIterator is returned from FilterOpenLimitCanceled and is used to iterate over the raw logs and unpacked data for OpenLimitCanceled events raised by the Bindings contract.
type BindingsOpenLimitCanceledIterator struct {
	Event *BindingsOpenLimitCanceled // Event containing the contract specifics and raw log

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
func (it *BindingsOpenLimitCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOpenLimitCanceled)
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
		it.Event = new(BindingsOpenLimitCanceled)
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
func (it *BindingsOpenLimitCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOpenLimitCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOpenLimitCanceled represents a OpenLimitCanceled event raised by the Bindings contract.
type BindingsOpenLimitCanceled struct {
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOpenLimitCanceled is a free log retrieval operation binding the contract event 0xf1b38881d7f4b2b12141c5f39c5124545d6112532eb6afbe9630cdbde3ee53e9.
//
// Solidity: event OpenLimitCanceled(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) FilterOpenLimitCanceled(opts *bind.FilterOpts, trader []common.Address, pairIndex []*big.Int) (*BindingsOpenLimitCanceledIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OpenLimitCanceled", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOpenLimitCanceledIterator{contract: _Bindings.contract, event: "OpenLimitCanceled", logs: logs, sub: sub}, nil
}

// WatchOpenLimitCanceled is a free log subscription operation binding the contract event 0xf1b38881d7f4b2b12141c5f39c5124545d6112532eb6afbe9630cdbde3ee53e9.
//
// Solidity: event OpenLimitCanceled(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) WatchOpenLimitCanceled(opts *bind.WatchOpts, sink chan<- *BindingsOpenLimitCanceled, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OpenLimitCanceled", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOpenLimitCanceled)
				if err := _Bindings.contract.UnpackLog(event, "OpenLimitCanceled", log); err != nil {
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

// ParseOpenLimitCanceled is a log parse operation binding the contract event 0xf1b38881d7f4b2b12141c5f39c5124545d6112532eb6afbe9630cdbde3ee53e9.
//
// Solidity: event OpenLimitCanceled(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) ParseOpenLimitCanceled(log types.Log) (*BindingsOpenLimitCanceled, error) {
	event := new(BindingsOpenLimitCanceled)
	if err := _Bindings.contract.UnpackLog(event, "OpenLimitCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOpenLimitPlacedIterator is returned from FilterOpenLimitPlaced and is used to iterate over the raw logs and unpacked data for OpenLimitPlaced events raised by the Bindings contract.
type BindingsOpenLimitPlacedIterator struct {
	Event *BindingsOpenLimitPlaced // Event containing the contract specifics and raw log

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
func (it *BindingsOpenLimitPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOpenLimitPlaced)
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
		it.Event = new(BindingsOpenLimitPlaced)
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
func (it *BindingsOpenLimitPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOpenLimitPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOpenLimitPlaced represents a OpenLimitPlaced event raised by the Bindings contract.
type BindingsOpenLimitPlaced struct {
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOpenLimitPlaced is a free log retrieval operation binding the contract event 0xdfabd6f206f17b7f2e1f9e0d33c40d30d1e8d7b6a4f520a03fdc1c1811059343.
//
// Solidity: event OpenLimitPlaced(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) FilterOpenLimitPlaced(opts *bind.FilterOpts, trader []common.Address, pairIndex []*big.Int) (*BindingsOpenLimitPlacedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OpenLimitPlaced", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOpenLimitPlacedIterator{contract: _Bindings.contract, event: "OpenLimitPlaced", logs: logs, sub: sub}, nil
}

// WatchOpenLimitPlaced is a free log subscription operation binding the contract event 0xdfabd6f206f17b7f2e1f9e0d33c40d30d1e8d7b6a4f520a03fdc1c1811059343.
//
// Solidity: event OpenLimitPlaced(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) WatchOpenLimitPlaced(opts *bind.WatchOpts, sink chan<- *BindingsOpenLimitPlaced, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OpenLimitPlaced", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOpenLimitPlaced)
				if err := _Bindings.contract.UnpackLog(event, "OpenLimitPlaced", log); err != nil {
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

// ParseOpenLimitPlaced is a log parse operation binding the contract event 0xdfabd6f206f17b7f2e1f9e0d33c40d30d1e8d7b6a4f520a03fdc1c1811059343.
//
// Solidity: event OpenLimitPlaced(address indexed trader, uint256 indexed pairIndex, uint256 index)
func (_Bindings *BindingsFilterer) ParseOpenLimitPlaced(log types.Log) (*BindingsOpenLimitPlaced, error) {
	event := new(BindingsOpenLimitPlaced)
	if err := _Bindings.contract.UnpackLog(event, "OpenLimitPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOpenLimitUpdatedIterator is returned from FilterOpenLimitUpdated and is used to iterate over the raw logs and unpacked data for OpenLimitUpdated events raised by the Bindings contract.
type BindingsOpenLimitUpdatedIterator struct {
	Event *BindingsOpenLimitUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsOpenLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOpenLimitUpdated)
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
		it.Event = new(BindingsOpenLimitUpdated)
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
func (it *BindingsOpenLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOpenLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOpenLimitUpdated represents a OpenLimitUpdated event raised by the Bindings contract.
type BindingsOpenLimitUpdated struct {
	Trader       common.Address
	PairIndex    *big.Int
	Index        *big.Int
	NewPrice     *big.Int
	NewTp        *big.Int
	NewSl        *big.Int
	MaxSlippageP *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOpenLimitUpdated is a free log retrieval operation binding the contract event 0x710a8db87f04e82a9de40076812593a965f4aa48693196d2144c07ff9710e890.
//
// Solidity: event OpenLimitUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newPrice, uint256 newTp, uint256 newSl, uint256 maxSlippageP)
func (_Bindings *BindingsFilterer) FilterOpenLimitUpdated(opts *bind.FilterOpts, trader []common.Address, pairIndex []*big.Int) (*BindingsOpenLimitUpdatedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OpenLimitUpdated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOpenLimitUpdatedIterator{contract: _Bindings.contract, event: "OpenLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchOpenLimitUpdated is a free log subscription operation binding the contract event 0x710a8db87f04e82a9de40076812593a965f4aa48693196d2144c07ff9710e890.
//
// Solidity: event OpenLimitUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newPrice, uint256 newTp, uint256 newSl, uint256 maxSlippageP)
func (_Bindings *BindingsFilterer) WatchOpenLimitUpdated(opts *bind.WatchOpts, sink chan<- *BindingsOpenLimitUpdated, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OpenLimitUpdated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOpenLimitUpdated)
				if err := _Bindings.contract.UnpackLog(event, "OpenLimitUpdated", log); err != nil {
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

// ParseOpenLimitUpdated is a log parse operation binding the contract event 0x710a8db87f04e82a9de40076812593a965f4aa48693196d2144c07ff9710e890.
//
// Solidity: event OpenLimitUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newPrice, uint256 newTp, uint256 newSl, uint256 maxSlippageP)
func (_Bindings *BindingsFilterer) ParseOpenLimitUpdated(log types.Log) (*BindingsOpenLimitUpdated, error) {
	event := new(BindingsOpenLimitUpdated)
	if err := _Bindings.contract.UnpackLog(event, "OpenLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bindings contract.
type BindingsPausedIterator struct {
	Event *BindingsPaused // Event containing the contract specifics and raw log

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
func (it *BindingsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPaused)
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
		it.Event = new(BindingsPaused)
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
func (it *BindingsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPaused represents a Paused event raised by the Bindings contract.
type BindingsPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Bindings *BindingsFilterer) FilterPaused(opts *bind.FilterOpts) (*BindingsPausedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BindingsPausedIterator{contract: _Bindings.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Bindings *BindingsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BindingsPaused) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPaused)
				if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: event Paused(bool paused)
func (_Bindings *BindingsFilterer) ParsePaused(log types.Log) (*BindingsPaused, error) {
	event := new(BindingsPaused)
	if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSlUpdatedIterator is returned from FilterSlUpdated and is used to iterate over the raw logs and unpacked data for SlUpdated events raised by the Bindings contract.
type BindingsSlUpdatedIterator struct {
	Event *BindingsSlUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsSlUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSlUpdated)
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
		it.Event = new(BindingsSlUpdated)
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
func (it *BindingsSlUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSlUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSlUpdated represents a SlUpdated event raised by the Bindings contract.
type BindingsSlUpdated struct {
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	NewSl     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlUpdated is a free log retrieval operation binding the contract event 0x1fc4a6c7ffe506697979b8ed54dc4135cd1ecd26a2745f70b760a2492222b316.
//
// Solidity: event SlUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newSl)
func (_Bindings *BindingsFilterer) FilterSlUpdated(opts *bind.FilterOpts, trader []common.Address, pairIndex []*big.Int) (*BindingsSlUpdatedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SlUpdated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSlUpdatedIterator{contract: _Bindings.contract, event: "SlUpdated", logs: logs, sub: sub}, nil
}

// WatchSlUpdated is a free log subscription operation binding the contract event 0x1fc4a6c7ffe506697979b8ed54dc4135cd1ecd26a2745f70b760a2492222b316.
//
// Solidity: event SlUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newSl)
func (_Bindings *BindingsFilterer) WatchSlUpdated(opts *bind.WatchOpts, sink chan<- *BindingsSlUpdated, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SlUpdated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSlUpdated)
				if err := _Bindings.contract.UnpackLog(event, "SlUpdated", log); err != nil {
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

// ParseSlUpdated is a log parse operation binding the contract event 0x1fc4a6c7ffe506697979b8ed54dc4135cd1ecd26a2745f70b760a2492222b316.
//
// Solidity: event SlUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newSl)
func (_Bindings *BindingsFilterer) ParseSlUpdated(log types.Log) (*BindingsSlUpdated, error) {
	event := new(BindingsSlUpdated)
	if err := _Bindings.contract.UnpackLog(event, "SlUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsTpUpdatedIterator is returned from FilterTpUpdated and is used to iterate over the raw logs and unpacked data for TpUpdated events raised by the Bindings contract.
type BindingsTpUpdatedIterator struct {
	Event *BindingsTpUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsTpUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTpUpdated)
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
		it.Event = new(BindingsTpUpdated)
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
func (it *BindingsTpUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTpUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTpUpdated represents a TpUpdated event raised by the Bindings contract.
type BindingsTpUpdated struct {
	Trader    common.Address
	PairIndex *big.Int
	Index     *big.Int
	NewTp     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTpUpdated is a free log retrieval operation binding the contract event 0x7e06a81c7a47891ccc7455b5ccb2ed850e32bb655ccda67eb3ebaaeed83242a4.
//
// Solidity: event TpUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newTp)
func (_Bindings *BindingsFilterer) FilterTpUpdated(opts *bind.FilterOpts, trader []common.Address, pairIndex []*big.Int) (*BindingsTpUpdatedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "TpUpdated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTpUpdatedIterator{contract: _Bindings.contract, event: "TpUpdated", logs: logs, sub: sub}, nil
}

// WatchTpUpdated is a free log subscription operation binding the contract event 0x7e06a81c7a47891ccc7455b5ccb2ed850e32bb655ccda67eb3ebaaeed83242a4.
//
// Solidity: event TpUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newTp)
func (_Bindings *BindingsFilterer) WatchTpUpdated(opts *bind.WatchOpts, sink chan<- *BindingsTpUpdated, trader []common.Address, pairIndex []*big.Int) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}
	var pairIndexRule []interface{}
	for _, pairIndexItem := range pairIndex {
		pairIndexRule = append(pairIndexRule, pairIndexItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "TpUpdated", traderRule, pairIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTpUpdated)
				if err := _Bindings.contract.UnpackLog(event, "TpUpdated", log); err != nil {
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

// ParseTpUpdated is a log parse operation binding the contract event 0x7e06a81c7a47891ccc7455b5ccb2ed850e32bb655ccda67eb3ebaaeed83242a4.
//
// Solidity: event TpUpdated(address indexed trader, uint256 indexed pairIndex, uint256 index, uint256 newTp)
func (_Bindings *BindingsFilterer) ParseTpUpdated(log types.Log) (*BindingsTpUpdated, error) {
	event := new(BindingsTpUpdated)
	if err := _Bindings.contract.UnpackLog(event, "TpUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
