// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coin

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

// CoinMetaData contains all meta data concerning the Coin contract.
var CoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506105838061005c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c8063075461721461004e57806327e235e31461006c57806340c10f191461009c578063d0679d34146100b8575b5f80fd5b6100566100d4565b6040516100639190610363565b60405180910390f35b610086600480360381019061008191906103aa565b6100f7565b60405161009391906103ed565b60405180910390f35b6100b660048036038101906100b19190610430565b61010c565b005b6100d260048036038101906100cd9190610430565b6101b9565b005b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610162575f80fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546101ae919061049b565b925050819055505050565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205481111561027a578060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546040517fcf4791810000000000000000000000000000000000000000000000000000000081526004016102719291906104ce565b60405180910390fd5b8060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546102c691906104f5565b925050819055508060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610319919061049b565b925050819055505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61034d82610324565b9050919050565b61035d81610343565b82525050565b5f6020820190506103765f830184610354565b92915050565b5f80fd5b61038981610343565b8114610393575f80fd5b50565b5f813590506103a481610380565b92915050565b5f602082840312156103bf576103be61037c565b5b5f6103cc84828501610396565b91505092915050565b5f819050919050565b6103e7816103d5565b82525050565b5f6020820190506104005f8301846103de565b92915050565b61040f816103d5565b8114610419575f80fd5b50565b5f8135905061042a81610406565b92915050565b5f80604083850312156104465761044561037c565b5b5f61045385828601610396565b92505060206104648582860161041c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6104a5826103d5565b91506104b0836103d5565b92508282019050808211156104c8576104c761046e565b5b92915050565b5f6040820190506104e15f8301856103de565b6104ee60208301846103de565b9392505050565b5f6104ff826103d5565b915061050a836103d5565b92508282039050818111156105225761052161046e565b5b9291505056fea26469706673582212204b4a11c402e40e62d51d06842d243d584268bcc1ab238d7f1f91615d1fcdf06e64736f6c637827302e382e32322d646576656c6f702e323032332e392e352b636f6d6d69742e31366165373663610058",
}

// CoinABI is the input ABI used to generate the binding from.
// Deprecated: Use CoinMetaData.ABI instead.
var CoinABI = CoinMetaData.ABI

// CoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoinMetaData.Bin instead.
var CoinBin = CoinMetaData.Bin

// DeployCoin deploys a new Ethereum contract, binding an instance of Coin to it.
func DeployCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Coin, error) {
	parsed, err := CoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Coin{CoinCaller: CoinCaller{contract: contract}, CoinTransactor: CoinTransactor{contract: contract}, CoinFilterer: CoinFilterer{contract: contract}}, nil
}

// Coin is an auto generated Go binding around an Ethereum contract.
type Coin struct {
	CoinCaller     // Read-only binding to the contract
	CoinTransactor // Write-only binding to the contract
	CoinFilterer   // Log filterer for contract events
}

// CoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinSession struct {
	Contract     *Coin             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinCallerSession struct {
	Contract *CoinCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinTransactorSession struct {
	Contract     *CoinTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinRaw struct {
	Contract *Coin // Generic contract binding to access the raw methods on
}

// CoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinCallerRaw struct {
	Contract *CoinCaller // Generic read-only contract binding to access the raw methods on
}

// CoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinTransactorRaw struct {
	Contract *CoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoin creates a new instance of Coin, bound to a specific deployed contract.
func NewCoin(address common.Address, backend bind.ContractBackend) (*Coin, error) {
	contract, err := bindCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Coin{CoinCaller: CoinCaller{contract: contract}, CoinTransactor: CoinTransactor{contract: contract}, CoinFilterer: CoinFilterer{contract: contract}}, nil
}

// NewCoinCaller creates a new read-only instance of Coin, bound to a specific deployed contract.
func NewCoinCaller(address common.Address, caller bind.ContractCaller) (*CoinCaller, error) {
	contract, err := bindCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinCaller{contract: contract}, nil
}

// NewCoinTransactor creates a new write-only instance of Coin, bound to a specific deployed contract.
func NewCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinTransactor, error) {
	contract, err := bindCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinTransactor{contract: contract}, nil
}

// NewCoinFilterer creates a new log filterer instance of Coin, bound to a specific deployed contract.
func NewCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinFilterer, error) {
	contract, err := bindCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinFilterer{contract: contract}, nil
}

// bindCoin binds a generic wrapper to an already deployed contract.
func bindCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coin *CoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coin.Contract.CoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coin *CoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coin.Contract.CoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coin *CoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coin.Contract.CoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Coin *CoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Coin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Coin *CoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Coin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Coin *CoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Coin.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Coin *CoinCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Coin.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Coin *CoinSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Coin.Contract.Balances(&_Coin.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Coin *CoinCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Coin.Contract.Balances(&_Coin.CallOpts, arg0)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Coin *CoinCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Coin.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Coin *CoinSession) Minter() (common.Address, error) {
	return _Coin.Contract.Minter(&_Coin.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Coin *CoinCallerSession) Minter() (common.Address, error) {
	return _Coin.Contract.Minter(&_Coin.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactor) Mint(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.contract.Transact(opts, "mint", receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Coin *CoinSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Mint(&_Coin.TransactOpts, receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactorSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Mint(&_Coin.TransactOpts, receiver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactor) Send(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.contract.Transact(opts, "send", receiver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address receiver, uint256 amount) returns()
func (_Coin *CoinSession) Send(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Send(&_Coin.TransactOpts, receiver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address receiver, uint256 amount) returns()
func (_Coin *CoinTransactorSession) Send(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Coin.Contract.Send(&_Coin.TransactOpts, receiver, amount)
}
