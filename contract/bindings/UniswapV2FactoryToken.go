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
	_ = abi.ConvertType
)

// UniswapV2FactoryTokenMetaData contains all meta data concerning the UniswapV2FactoryToken contract.
var UniswapV2FactoryTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapV2FactoryTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2FactoryTokenMetaData.ABI instead.
var UniswapV2FactoryTokenABI = UniswapV2FactoryTokenMetaData.ABI

// UniswapV2FactoryToken is an auto generated Go binding around an Ethereum contract.
type UniswapV2FactoryToken struct {
	UniswapV2FactoryTokenCaller     // Read-only binding to the contract
	UniswapV2FactoryTokenTransactor // Write-only binding to the contract
	UniswapV2FactoryTokenFilterer   // Log filterer for contract events
}

// UniswapV2FactoryTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2FactoryTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2FactoryTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2FactoryTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2FactoryTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2FactoryTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2FactoryTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2FactoryTokenSession struct {
	Contract     *UniswapV2FactoryToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UniswapV2FactoryTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2FactoryTokenCallerSession struct {
	Contract *UniswapV2FactoryTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// UniswapV2FactoryTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2FactoryTokenTransactorSession struct {
	Contract     *UniswapV2FactoryTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// UniswapV2FactoryTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2FactoryTokenRaw struct {
	Contract *UniswapV2FactoryToken // Generic contract binding to access the raw methods on
}

// UniswapV2FactoryTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2FactoryTokenCallerRaw struct {
	Contract *UniswapV2FactoryTokenCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2FactoryTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2FactoryTokenTransactorRaw struct {
	Contract *UniswapV2FactoryTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2FactoryToken creates a new instance of UniswapV2FactoryToken, bound to a specific deployed contract.
func NewUniswapV2FactoryToken(address common.Address, backend bind.ContractBackend) (*UniswapV2FactoryToken, error) {
	contract, err := bindUniswapV2FactoryToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryToken{UniswapV2FactoryTokenCaller: UniswapV2FactoryTokenCaller{contract: contract}, UniswapV2FactoryTokenTransactor: UniswapV2FactoryTokenTransactor{contract: contract}, UniswapV2FactoryTokenFilterer: UniswapV2FactoryTokenFilterer{contract: contract}}, nil
}

// NewUniswapV2FactoryTokenCaller creates a new read-only instance of UniswapV2FactoryToken, bound to a specific deployed contract.
func NewUniswapV2FactoryTokenCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2FactoryTokenCaller, error) {
	contract, err := bindUniswapV2FactoryToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryTokenCaller{contract: contract}, nil
}

// NewUniswapV2FactoryTokenTransactor creates a new write-only instance of UniswapV2FactoryToken, bound to a specific deployed contract.
func NewUniswapV2FactoryTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2FactoryTokenTransactor, error) {
	contract, err := bindUniswapV2FactoryToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryTokenTransactor{contract: contract}, nil
}

// NewUniswapV2FactoryTokenFilterer creates a new log filterer instance of UniswapV2FactoryToken, bound to a specific deployed contract.
func NewUniswapV2FactoryTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2FactoryTokenFilterer, error) {
	contract, err := bindUniswapV2FactoryToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryTokenFilterer{contract: contract}, nil
}

// bindUniswapV2FactoryToken binds a generic wrapper to an already deployed contract.
func bindUniswapV2FactoryToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV2FactoryTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2FactoryToken.Contract.UniswapV2FactoryTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.UniswapV2FactoryTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.UniswapV2FactoryTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2FactoryToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2FactoryToken.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.AllPairs(&_UniswapV2FactoryToken.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.AllPairs(&_UniswapV2FactoryToken.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2FactoryToken.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) AllPairsLength() (*big.Int, error) {
	return _UniswapV2FactoryToken.Contract.AllPairsLength(&_UniswapV2FactoryToken.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCallerSession) AllPairsLength() (*big.Int, error) {
	return _UniswapV2FactoryToken.Contract.AllPairsLength(&_UniswapV2FactoryToken.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2FactoryToken.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) FeeTo() (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.FeeTo(&_UniswapV2FactoryToken.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCallerSession) FeeTo() (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.FeeTo(&_UniswapV2FactoryToken.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2FactoryToken.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) FeeToSetter() (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.FeeToSetter(&_UniswapV2FactoryToken.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCallerSession) FeeToSetter() (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.FeeToSetter(&_UniswapV2FactoryToken.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2FactoryToken.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.GetPair(&_UniswapV2FactoryToken.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _UniswapV2FactoryToken.Contract.GetPair(&_UniswapV2FactoryToken.CallOpts, arg0, arg1)
}

// InitCodeHash is a free data retrieval call binding the contract method 0xdb4c545e.
//
// Solidity: function initCodeHash() view returns(bytes32)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCaller) InitCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV2FactoryToken.contract.Call(opts, &out, "initCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InitCodeHash is a free data retrieval call binding the contract method 0xdb4c545e.
//
// Solidity: function initCodeHash() view returns(bytes32)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) InitCodeHash() ([32]byte, error) {
	return _UniswapV2FactoryToken.Contract.InitCodeHash(&_UniswapV2FactoryToken.CallOpts)
}

// InitCodeHash is a free data retrieval call binding the contract method 0xdb4c545e.
//
// Solidity: function initCodeHash() view returns(bytes32)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenCallerSession) InitCodeHash() ([32]byte, error) {
	return _UniswapV2FactoryToken.Contract.InitCodeHash(&_UniswapV2FactoryToken.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.CreatePair(&_UniswapV2FactoryToken.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.CreatePair(&_UniswapV2FactoryToken.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.SetFeeTo(&_UniswapV2FactoryToken.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.SetFeeTo(&_UniswapV2FactoryToken.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.SetFeeToSetter(&_UniswapV2FactoryToken.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _UniswapV2FactoryToken.Contract.SetFeeToSetter(&_UniswapV2FactoryToken.TransactOpts, _feeToSetter)
}

// UniswapV2FactoryTokenPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the UniswapV2FactoryToken contract.
type UniswapV2FactoryTokenPairCreatedIterator struct {
	Event *UniswapV2FactoryTokenPairCreated // Event containing the contract specifics and raw log

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
func (it *UniswapV2FactoryTokenPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapV2FactoryTokenPairCreated)
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
		it.Event = new(UniswapV2FactoryTokenPairCreated)
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
func (it *UniswapV2FactoryTokenPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapV2FactoryTokenPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapV2FactoryTokenPairCreated represents a PairCreated event raised by the UniswapV2FactoryToken contract.
type UniswapV2FactoryTokenPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*UniswapV2FactoryTokenPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _UniswapV2FactoryToken.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryTokenPairCreatedIterator{contract: _UniswapV2FactoryToken.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *UniswapV2FactoryTokenPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _UniswapV2FactoryToken.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapV2FactoryTokenPairCreated)
				if err := _UniswapV2FactoryToken.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_UniswapV2FactoryToken *UniswapV2FactoryTokenFilterer) ParsePairCreated(log types.Log) (*UniswapV2FactoryTokenPairCreated, error) {
	event := new(UniswapV2FactoryTokenPairCreated)
	if err := _UniswapV2FactoryToken.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
