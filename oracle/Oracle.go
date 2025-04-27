// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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
	_ = abi.ConvertType
)

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseCurrencyUnit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AssetSourceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCurrencyUnit\",\"type\":\"uint256\"}],\"name\":\"BaseCurrencySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"FallbackOracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CURRENCY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"getAssetsPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getSourceOfAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"}],\"name\":\"setAssetSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"setFallbackOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Oracle *OracleCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Oracle *OracleSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Oracle.Contract.ADDRESSESPROVIDER(&_Oracle.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Oracle *OracleCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Oracle.Contract.ADDRESSESPROVIDER(&_Oracle.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_Oracle *OracleCaller) BASECURRENCY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "BASE_CURRENCY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_Oracle *OracleSession) BASECURRENCY() (common.Address, error) {
	return _Oracle.Contract.BASECURRENCY(&_Oracle.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_Oracle *OracleCallerSession) BASECURRENCY() (common.Address, error) {
	return _Oracle.Contract.BASECURRENCY(&_Oracle.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_Oracle *OracleCaller) BASECURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "BASE_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_Oracle *OracleSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _Oracle.Contract.BASECURRENCYUNIT(&_Oracle.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_Oracle *OracleCallerSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _Oracle.Contract.BASECURRENCYUNIT(&_Oracle.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Oracle *OracleCaller) GetAssetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getAssetPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Oracle *OracleSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetAssetPrice(&_Oracle.CallOpts, asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Oracle *OracleCallerSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetAssetPrice(&_Oracle.CallOpts, asset)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_Oracle *OracleCaller) GetAssetsPrices(opts *bind.CallOpts, assets []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getAssetsPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_Oracle *OracleSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _Oracle.Contract.GetAssetsPrices(&_Oracle.CallOpts, assets)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_Oracle *OracleCallerSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _Oracle.Contract.GetAssetsPrices(&_Oracle.CallOpts, assets)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_Oracle *OracleCaller) GetFallbackOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getFallbackOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_Oracle *OracleSession) GetFallbackOracle() (common.Address, error) {
	return _Oracle.Contract.GetFallbackOracle(&_Oracle.CallOpts)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_Oracle *OracleCallerSession) GetFallbackOracle() (common.Address, error) {
	return _Oracle.Contract.GetFallbackOracle(&_Oracle.CallOpts)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_Oracle *OracleCaller) GetSourceOfAsset(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getSourceOfAsset", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_Oracle *OracleSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _Oracle.Contract.GetSourceOfAsset(&_Oracle.CallOpts, asset)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_Oracle *OracleCallerSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _Oracle.Contract.GetSourceOfAsset(&_Oracle.CallOpts, asset)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_Oracle *OracleTransactor) SetAssetSources(opts *bind.TransactOpts, assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setAssetSources", assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_Oracle *OracleSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssetSources(&_Oracle.TransactOpts, assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_Oracle *OracleTransactorSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssetSources(&_Oracle.TransactOpts, assets, sources)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_Oracle *OracleTransactor) SetFallbackOracle(opts *bind.TransactOpts, fallbackOracle common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setFallbackOracle", fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_Oracle *OracleSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetFallbackOracle(&_Oracle.TransactOpts, fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_Oracle *OracleTransactorSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetFallbackOracle(&_Oracle.TransactOpts, fallbackOracle)
}

// OracleAssetSourceUpdatedIterator is returned from FilterAssetSourceUpdated and is used to iterate over the raw logs and unpacked data for AssetSourceUpdated events raised by the Oracle contract.
type OracleAssetSourceUpdatedIterator struct {
	Event *OracleAssetSourceUpdated // Event containing the contract specifics and raw log

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
func (it *OracleAssetSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAssetSourceUpdated)
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
		it.Event = new(OracleAssetSourceUpdated)
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
func (it *OracleAssetSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAssetSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAssetSourceUpdated represents a AssetSourceUpdated event raised by the Oracle contract.
type OracleAssetSourceUpdated struct {
	Asset  common.Address
	Source common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetSourceUpdated is a free log retrieval operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_Oracle *OracleFilterer) FilterAssetSourceUpdated(opts *bind.FilterOpts, asset []common.Address, source []common.Address) (*OracleAssetSourceUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &OracleAssetSourceUpdatedIterator{contract: _Oracle.contract, event: "AssetSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetSourceUpdated is a free log subscription operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_Oracle *OracleFilterer) WatchAssetSourceUpdated(opts *bind.WatchOpts, sink chan<- *OracleAssetSourceUpdated, asset []common.Address, source []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAssetSourceUpdated)
				if err := _Oracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
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

// ParseAssetSourceUpdated is a log parse operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_Oracle *OracleFilterer) ParseAssetSourceUpdated(log types.Log) (*OracleAssetSourceUpdated, error) {
	event := new(OracleAssetSourceUpdated)
	if err := _Oracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleBaseCurrencySetIterator is returned from FilterBaseCurrencySet and is used to iterate over the raw logs and unpacked data for BaseCurrencySet events raised by the Oracle contract.
type OracleBaseCurrencySetIterator struct {
	Event *OracleBaseCurrencySet // Event containing the contract specifics and raw log

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
func (it *OracleBaseCurrencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleBaseCurrencySet)
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
		it.Event = new(OracleBaseCurrencySet)
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
func (it *OracleBaseCurrencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleBaseCurrencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleBaseCurrencySet represents a BaseCurrencySet event raised by the Oracle contract.
type OracleBaseCurrencySet struct {
	BaseCurrency     common.Address
	BaseCurrencyUnit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBaseCurrencySet is a free log retrieval operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_Oracle *OracleFilterer) FilterBaseCurrencySet(opts *bind.FilterOpts, baseCurrency []common.Address) (*OracleBaseCurrencySetIterator, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &OracleBaseCurrencySetIterator{contract: _Oracle.contract, event: "BaseCurrencySet", logs: logs, sub: sub}, nil
}

// WatchBaseCurrencySet is a free log subscription operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_Oracle *OracleFilterer) WatchBaseCurrencySet(opts *bind.WatchOpts, sink chan<- *OracleBaseCurrencySet, baseCurrency []common.Address) (event.Subscription, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleBaseCurrencySet)
				if err := _Oracle.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
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

// ParseBaseCurrencySet is a log parse operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_Oracle *OracleFilterer) ParseBaseCurrencySet(log types.Log) (*OracleBaseCurrencySet, error) {
	event := new(OracleBaseCurrencySet)
	if err := _Oracle.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFallbackOracleUpdatedIterator is returned from FilterFallbackOracleUpdated and is used to iterate over the raw logs and unpacked data for FallbackOracleUpdated events raised by the Oracle contract.
type OracleFallbackOracleUpdatedIterator struct {
	Event *OracleFallbackOracleUpdated // Event containing the contract specifics and raw log

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
func (it *OracleFallbackOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFallbackOracleUpdated)
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
		it.Event = new(OracleFallbackOracleUpdated)
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
func (it *OracleFallbackOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFallbackOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFallbackOracleUpdated represents a FallbackOracleUpdated event raised by the Oracle contract.
type OracleFallbackOracleUpdated struct {
	FallbackOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFallbackOracleUpdated is a free log retrieval operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_Oracle *OracleFilterer) FilterFallbackOracleUpdated(opts *bind.FilterOpts, fallbackOracle []common.Address) (*OracleFallbackOracleUpdatedIterator, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return &OracleFallbackOracleUpdatedIterator{contract: _Oracle.contract, event: "FallbackOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchFallbackOracleUpdated is a free log subscription operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_Oracle *OracleFilterer) WatchFallbackOracleUpdated(opts *bind.WatchOpts, sink chan<- *OracleFallbackOracleUpdated, fallbackOracle []common.Address) (event.Subscription, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFallbackOracleUpdated)
				if err := _Oracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
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

// ParseFallbackOracleUpdated is a log parse operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_Oracle *OracleFilterer) ParseFallbackOracleUpdated(log types.Log) (*OracleFallbackOracleUpdated, error) {
	event := new(OracleFallbackOracleUpdated)
	if err := _Oracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
