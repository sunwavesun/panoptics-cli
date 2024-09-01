// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PanopticFactory

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

// PanopticFactoryMetaData contains all meta data concerning the PanopticFactory contract.
var PanopticFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"},{\"internalType\":\"contractSemiFungiblePositionManager\",\"name\":\"_SFPM\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"_univ3Factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolReference\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralReference\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidUniswapCallback\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UniswapPoolNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UniswapPoolNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractPanopticPool\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"uniswapPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractCollateralTracker\",\"name\":\"collateralTracker0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractCollateralTracker\",\"name\":\"collateralTracker1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rareNftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rarity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"PoolDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint96\",\"name\":\"salt\",\"type\":\"uint96\"}],\"name\":\"deployNewPool\",\"outputs\":[{\"internalType\":\"contractPanopticPool\",\"name\":\"newPoolContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"univ3pool\",\"type\":\"address\"}],\"name\":\"getPanopticPool\",\"outputs\":[{\"internalType\":\"contractPanopticPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint96\",\"name\":\"salt\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loops\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTargetRarity\",\"type\":\"uint256\"}],\"name\":\"minePoolAddress\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"bestSalt\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"highestRarity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PanopticFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PanopticFactoryMetaData.ABI instead.
var PanopticFactoryABI = PanopticFactoryMetaData.ABI

// PanopticFactory is an auto generated Go binding around an Ethereum contract.
type PanopticFactory struct {
	PanopticFactoryCaller     // Read-only binding to the contract
	PanopticFactoryTransactor // Write-only binding to the contract
	PanopticFactoryFilterer   // Log filterer for contract events
}

// PanopticFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PanopticFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanopticFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PanopticFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanopticFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PanopticFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanopticFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PanopticFactorySession struct {
	Contract     *PanopticFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PanopticFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PanopticFactoryCallerSession struct {
	Contract *PanopticFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PanopticFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PanopticFactoryTransactorSession struct {
	Contract     *PanopticFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PanopticFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PanopticFactoryRaw struct {
	Contract *PanopticFactory // Generic contract binding to access the raw methods on
}

// PanopticFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PanopticFactoryCallerRaw struct {
	Contract *PanopticFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PanopticFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PanopticFactoryTransactorRaw struct {
	Contract *PanopticFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPanopticFactory creates a new instance of PanopticFactory, bound to a specific deployed contract.
func NewPanopticFactory(address common.Address, backend bind.ContractBackend) (*PanopticFactory, error) {
	contract, err := bindPanopticFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PanopticFactory{PanopticFactoryCaller: PanopticFactoryCaller{contract: contract}, PanopticFactoryTransactor: PanopticFactoryTransactor{contract: contract}, PanopticFactoryFilterer: PanopticFactoryFilterer{contract: contract}}, nil
}

// NewPanopticFactoryCaller creates a new read-only instance of PanopticFactory, bound to a specific deployed contract.
func NewPanopticFactoryCaller(address common.Address, caller bind.ContractCaller) (*PanopticFactoryCaller, error) {
	contract, err := bindPanopticFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryCaller{contract: contract}, nil
}

// NewPanopticFactoryTransactor creates a new write-only instance of PanopticFactory, bound to a specific deployed contract.
func NewPanopticFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PanopticFactoryTransactor, error) {
	contract, err := bindPanopticFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryTransactor{contract: contract}, nil
}

// NewPanopticFactoryFilterer creates a new log filterer instance of PanopticFactory, bound to a specific deployed contract.
func NewPanopticFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PanopticFactoryFilterer, error) {
	contract, err := bindPanopticFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryFilterer{contract: contract}, nil
}

// bindPanopticFactory binds a generic wrapper to an already deployed contract.
func bindPanopticFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PanopticFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PanopticFactory *PanopticFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PanopticFactory.Contract.PanopticFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PanopticFactory *PanopticFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PanopticFactory.Contract.PanopticFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PanopticFactory *PanopticFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PanopticFactory.Contract.PanopticFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PanopticFactory *PanopticFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PanopticFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PanopticFactory *PanopticFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PanopticFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PanopticFactory *PanopticFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PanopticFactory.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PanopticFactory *PanopticFactoryCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PanopticFactory *PanopticFactorySession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _PanopticFactory.Contract.BalanceOf(&_PanopticFactory.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PanopticFactory *PanopticFactoryCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _PanopticFactory.Contract.BalanceOf(&_PanopticFactory.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PanopticFactory *PanopticFactoryCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PanopticFactory *PanopticFactorySession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _PanopticFactory.Contract.BalanceOfBatch(&_PanopticFactory.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PanopticFactory *PanopticFactoryCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _PanopticFactory.Contract.BalanceOfBatch(&_PanopticFactory.CallOpts, accounts, ids)
}

// FactoryOwner is a free data retrieval call binding the contract method 0x4273601c.
//
// Solidity: function factoryOwner() view returns(address)
func (_PanopticFactory *PanopticFactoryCaller) FactoryOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "factoryOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryOwner is a free data retrieval call binding the contract method 0x4273601c.
//
// Solidity: function factoryOwner() view returns(address)
func (_PanopticFactory *PanopticFactorySession) FactoryOwner() (common.Address, error) {
	return _PanopticFactory.Contract.FactoryOwner(&_PanopticFactory.CallOpts)
}

// FactoryOwner is a free data retrieval call binding the contract method 0x4273601c.
//
// Solidity: function factoryOwner() view returns(address)
func (_PanopticFactory *PanopticFactoryCallerSession) FactoryOwner() (common.Address, error) {
	return _PanopticFactory.Contract.FactoryOwner(&_PanopticFactory.CallOpts)
}

// GetPanopticPool is a free data retrieval call binding the contract method 0xf2626ead.
//
// Solidity: function getPanopticPool(address univ3pool) view returns(address)
func (_PanopticFactory *PanopticFactoryCaller) GetPanopticPool(opts *bind.CallOpts, univ3pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "getPanopticPool", univ3pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPanopticPool is a free data retrieval call binding the contract method 0xf2626ead.
//
// Solidity: function getPanopticPool(address univ3pool) view returns(address)
func (_PanopticFactory *PanopticFactorySession) GetPanopticPool(univ3pool common.Address) (common.Address, error) {
	return _PanopticFactory.Contract.GetPanopticPool(&_PanopticFactory.CallOpts, univ3pool)
}

// GetPanopticPool is a free data retrieval call binding the contract method 0xf2626ead.
//
// Solidity: function getPanopticPool(address univ3pool) view returns(address)
func (_PanopticFactory *PanopticFactoryCallerSession) GetPanopticPool(univ3pool common.Address) (common.Address, error) {
	return _PanopticFactory.Contract.GetPanopticPool(&_PanopticFactory.CallOpts, univ3pool)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PanopticFactory *PanopticFactoryCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PanopticFactory *PanopticFactorySession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _PanopticFactory.Contract.IsApprovedForAll(&_PanopticFactory.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PanopticFactory *PanopticFactoryCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _PanopticFactory.Contract.IsApprovedForAll(&_PanopticFactory.CallOpts, account, operator)
}

// MinePoolAddress is a free data retrieval call binding the contract method 0x5c8f50c0.
//
// Solidity: function minePoolAddress(address token0, address token1, uint24 fee, uint96 salt, address deployer, uint256 loops, uint256 minTargetRarity) view returns(uint96 bestSalt, uint256 highestRarity)
func (_PanopticFactory *PanopticFactoryCaller) MinePoolAddress(opts *bind.CallOpts, token0 common.Address, token1 common.Address, fee *big.Int, salt *big.Int, deployer common.Address, loops *big.Int, minTargetRarity *big.Int) (struct {
	BestSalt      *big.Int
	HighestRarity *big.Int
}, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "minePoolAddress", token0, token1, fee, salt, deployer, loops, minTargetRarity)

	outstruct := new(struct {
		BestSalt      *big.Int
		HighestRarity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BestSalt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HighestRarity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MinePoolAddress is a free data retrieval call binding the contract method 0x5c8f50c0.
//
// Solidity: function minePoolAddress(address token0, address token1, uint24 fee, uint96 salt, address deployer, uint256 loops, uint256 minTargetRarity) view returns(uint96 bestSalt, uint256 highestRarity)
func (_PanopticFactory *PanopticFactorySession) MinePoolAddress(token0 common.Address, token1 common.Address, fee *big.Int, salt *big.Int, deployer common.Address, loops *big.Int, minTargetRarity *big.Int) (struct {
	BestSalt      *big.Int
	HighestRarity *big.Int
}, error) {
	return _PanopticFactory.Contract.MinePoolAddress(&_PanopticFactory.CallOpts, token0, token1, fee, salt, deployer, loops, minTargetRarity)
}

// MinePoolAddress is a free data retrieval call binding the contract method 0x5c8f50c0.
//
// Solidity: function minePoolAddress(address token0, address token1, uint24 fee, uint96 salt, address deployer, uint256 loops, uint256 minTargetRarity) view returns(uint96 bestSalt, uint256 highestRarity)
func (_PanopticFactory *PanopticFactoryCallerSession) MinePoolAddress(token0 common.Address, token1 common.Address, fee *big.Int, salt *big.Int, deployer common.Address, loops *big.Int, minTargetRarity *big.Int) (struct {
	BestSalt      *big.Int
	HighestRarity *big.Int
}, error) {
	return _PanopticFactory.Contract.MinePoolAddress(&_PanopticFactory.CallOpts, token0, token1, fee, salt, deployer, loops, minTargetRarity)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PanopticFactory *PanopticFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PanopticFactory *PanopticFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PanopticFactory.Contract.SupportsInterface(&_PanopticFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PanopticFactory *PanopticFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PanopticFactory.Contract.SupportsInterface(&_PanopticFactory.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_PanopticFactory *PanopticFactoryCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PanopticFactory.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_PanopticFactory *PanopticFactorySession) Uri(arg0 *big.Int) (string, error) {
	return _PanopticFactory.Contract.Uri(&_PanopticFactory.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_PanopticFactory *PanopticFactoryCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _PanopticFactory.Contract.Uri(&_PanopticFactory.CallOpts, arg0)
}

// DeployNewPool is a paid mutator transaction binding the contract method 0xdaa4c697.
//
// Solidity: function deployNewPool(address token0, address token1, uint24 fee, uint96 salt) returns(address newPoolContract)
func (_PanopticFactory *PanopticFactoryTransactor) DeployNewPool(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _PanopticFactory.contract.Transact(opts, "deployNewPool", token0, token1, fee, salt)
}

// DeployNewPool is a paid mutator transaction binding the contract method 0xdaa4c697.
//
// Solidity: function deployNewPool(address token0, address token1, uint24 fee, uint96 salt) returns(address newPoolContract)
func (_PanopticFactory *PanopticFactorySession) DeployNewPool(token0 common.Address, token1 common.Address, fee *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _PanopticFactory.Contract.DeployNewPool(&_PanopticFactory.TransactOpts, token0, token1, fee, salt)
}

// DeployNewPool is a paid mutator transaction binding the contract method 0xdaa4c697.
//
// Solidity: function deployNewPool(address token0, address token1, uint24 fee, uint96 salt) returns(address newPoolContract)
func (_PanopticFactory *PanopticFactoryTransactorSession) DeployNewPool(token0 common.Address, token1 common.Address, fee *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _PanopticFactory.Contract.DeployNewPool(&_PanopticFactory.TransactOpts, token0, token1, fee, salt)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PanopticFactory *PanopticFactoryTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _PanopticFactory.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PanopticFactory *PanopticFactorySession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.Multicall(&_PanopticFactory.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_PanopticFactory *PanopticFactoryTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.Multicall(&_PanopticFactory.TransactOpts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PanopticFactory *PanopticFactoryTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PanopticFactory *PanopticFactorySession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SafeBatchTransferFrom(&_PanopticFactory.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PanopticFactory *PanopticFactoryTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SafeBatchTransferFrom(&_PanopticFactory.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PanopticFactory *PanopticFactoryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PanopticFactory *PanopticFactorySession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SafeTransferFrom(&_PanopticFactory.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PanopticFactory *PanopticFactoryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SafeTransferFrom(&_PanopticFactory.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PanopticFactory *PanopticFactoryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PanopticFactory.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PanopticFactory *PanopticFactorySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SetApprovalForAll(&_PanopticFactory.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PanopticFactory *PanopticFactoryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SetApprovalForAll(&_PanopticFactory.TransactOpts, operator, approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_PanopticFactory *PanopticFactoryTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PanopticFactory.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_PanopticFactory *PanopticFactorySession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SetOwner(&_PanopticFactory.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_PanopticFactory *PanopticFactoryTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PanopticFactory.Contract.SetOwner(&_PanopticFactory.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_PanopticFactory *PanopticFactoryTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_PanopticFactory *PanopticFactorySession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.UniswapV3MintCallback(&_PanopticFactory.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_PanopticFactory *PanopticFactoryTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _PanopticFactory.Contract.UniswapV3MintCallback(&_PanopticFactory.TransactOpts, amount0Owed, amount1Owed, data)
}

// PanopticFactoryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PanopticFactory contract.
type PanopticFactoryApprovalForAllIterator struct {
	Event *PanopticFactoryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PanopticFactoryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PanopticFactoryApprovalForAll)
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
		it.Event = new(PanopticFactoryApprovalForAll)
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
func (it *PanopticFactoryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PanopticFactoryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PanopticFactoryApprovalForAll represents a ApprovalForAll event raised by the PanopticFactory contract.
type PanopticFactoryApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PanopticFactory *PanopticFactoryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*PanopticFactoryApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PanopticFactory.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryApprovalForAllIterator{contract: _PanopticFactory.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PanopticFactory *PanopticFactoryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PanopticFactoryApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PanopticFactory.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PanopticFactoryApprovalForAll)
				if err := _PanopticFactory.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PanopticFactory *PanopticFactoryFilterer) ParseApprovalForAll(log types.Log) (*PanopticFactoryApprovalForAll, error) {
	event := new(PanopticFactoryApprovalForAll)
	if err := _PanopticFactory.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PanopticFactoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the PanopticFactory contract.
type PanopticFactoryOwnerChangedIterator struct {
	Event *PanopticFactoryOwnerChanged // Event containing the contract specifics and raw log

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
func (it *PanopticFactoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PanopticFactoryOwnerChanged)
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
		it.Event = new(PanopticFactoryOwnerChanged)
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
func (it *PanopticFactoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PanopticFactoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PanopticFactoryOwnerChanged represents a OwnerChanged event raised by the PanopticFactory contract.
type PanopticFactoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PanopticFactory *PanopticFactoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PanopticFactoryOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PanopticFactory.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryOwnerChangedIterator{contract: _PanopticFactory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PanopticFactory *PanopticFactoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *PanopticFactoryOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PanopticFactory.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PanopticFactoryOwnerChanged)
				if err := _PanopticFactory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PanopticFactory *PanopticFactoryFilterer) ParseOwnerChanged(log types.Log) (*PanopticFactoryOwnerChanged, error) {
	event := new(PanopticFactoryOwnerChanged)
	if err := _PanopticFactory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PanopticFactoryPoolDeployedIterator is returned from FilterPoolDeployed and is used to iterate over the raw logs and unpacked data for PoolDeployed events raised by the PanopticFactory contract.
type PanopticFactoryPoolDeployedIterator struct {
	Event *PanopticFactoryPoolDeployed // Event containing the contract specifics and raw log

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
func (it *PanopticFactoryPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PanopticFactoryPoolDeployed)
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
		it.Event = new(PanopticFactoryPoolDeployed)
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
func (it *PanopticFactoryPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PanopticFactoryPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PanopticFactoryPoolDeployed represents a PoolDeployed event raised by the PanopticFactory contract.
type PanopticFactoryPoolDeployed struct {
	PoolAddress        common.Address
	UniswapPool        common.Address
	CollateralTracker0 common.Address
	CollateralTracker1 common.Address
	RareNftId          *big.Int
	Rarity             *big.Int
	Amount0            *big.Int
	Amount1            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolDeployed is a free log retrieval operation binding the contract event 0xfd60c5699b7fabdc7ef5bdc78da781a026faa233e4bfc2969ee6ffb6540b29f3.
//
// Solidity: event PoolDeployed(address indexed poolAddress, address indexed uniswapPool, address collateralTracker0, address collateralTracker1, uint256 rareNftId, uint256 indexed rarity, uint256 amount0, uint256 amount1)
func (_PanopticFactory *PanopticFactoryFilterer) FilterPoolDeployed(opts *bind.FilterOpts, poolAddress []common.Address, uniswapPool []common.Address, rarity []*big.Int) (*PanopticFactoryPoolDeployedIterator, error) {

	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var uniswapPoolRule []interface{}
	for _, uniswapPoolItem := range uniswapPool {
		uniswapPoolRule = append(uniswapPoolRule, uniswapPoolItem)
	}

	var rarityRule []interface{}
	for _, rarityItem := range rarity {
		rarityRule = append(rarityRule, rarityItem)
	}

	logs, sub, err := _PanopticFactory.contract.FilterLogs(opts, "PoolDeployed", poolAddressRule, uniswapPoolRule, rarityRule)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryPoolDeployedIterator{contract: _PanopticFactory.contract, event: "PoolDeployed", logs: logs, sub: sub}, nil
}

// WatchPoolDeployed is a free log subscription operation binding the contract event 0xfd60c5699b7fabdc7ef5bdc78da781a026faa233e4bfc2969ee6ffb6540b29f3.
//
// Solidity: event PoolDeployed(address indexed poolAddress, address indexed uniswapPool, address collateralTracker0, address collateralTracker1, uint256 rareNftId, uint256 indexed rarity, uint256 amount0, uint256 amount1)
func (_PanopticFactory *PanopticFactoryFilterer) WatchPoolDeployed(opts *bind.WatchOpts, sink chan<- *PanopticFactoryPoolDeployed, poolAddress []common.Address, uniswapPool []common.Address, rarity []*big.Int) (event.Subscription, error) {

	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var uniswapPoolRule []interface{}
	for _, uniswapPoolItem := range uniswapPool {
		uniswapPoolRule = append(uniswapPoolRule, uniswapPoolItem)
	}

	var rarityRule []interface{}
	for _, rarityItem := range rarity {
		rarityRule = append(rarityRule, rarityItem)
	}

	logs, sub, err := _PanopticFactory.contract.WatchLogs(opts, "PoolDeployed", poolAddressRule, uniswapPoolRule, rarityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PanopticFactoryPoolDeployed)
				if err := _PanopticFactory.contract.UnpackLog(event, "PoolDeployed", log); err != nil {
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

// ParsePoolDeployed is a log parse operation binding the contract event 0xfd60c5699b7fabdc7ef5bdc78da781a026faa233e4bfc2969ee6ffb6540b29f3.
//
// Solidity: event PoolDeployed(address indexed poolAddress, address indexed uniswapPool, address collateralTracker0, address collateralTracker1, uint256 rareNftId, uint256 indexed rarity, uint256 amount0, uint256 amount1)
func (_PanopticFactory *PanopticFactoryFilterer) ParsePoolDeployed(log types.Log) (*PanopticFactoryPoolDeployed, error) {
	event := new(PanopticFactoryPoolDeployed)
	if err := _PanopticFactory.contract.UnpackLog(event, "PoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PanopticFactoryTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the PanopticFactory contract.
type PanopticFactoryTransferBatchIterator struct {
	Event *PanopticFactoryTransferBatch // Event containing the contract specifics and raw log

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
func (it *PanopticFactoryTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PanopticFactoryTransferBatch)
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
		it.Event = new(PanopticFactoryTransferBatch)
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
func (it *PanopticFactoryTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PanopticFactoryTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PanopticFactoryTransferBatch represents a TransferBatch event raised by the PanopticFactory contract.
type PanopticFactoryTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PanopticFactory *PanopticFactoryFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PanopticFactoryTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PanopticFactory.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryTransferBatchIterator{contract: _PanopticFactory.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PanopticFactory *PanopticFactoryFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *PanopticFactoryTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PanopticFactory.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PanopticFactoryTransferBatch)
				if err := _PanopticFactory.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PanopticFactory *PanopticFactoryFilterer) ParseTransferBatch(log types.Log) (*PanopticFactoryTransferBatch, error) {
	event := new(PanopticFactoryTransferBatch)
	if err := _PanopticFactory.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PanopticFactoryTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the PanopticFactory contract.
type PanopticFactoryTransferSingleIterator struct {
	Event *PanopticFactoryTransferSingle // Event containing the contract specifics and raw log

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
func (it *PanopticFactoryTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PanopticFactoryTransferSingle)
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
		it.Event = new(PanopticFactoryTransferSingle)
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
func (it *PanopticFactoryTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PanopticFactoryTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PanopticFactoryTransferSingle represents a TransferSingle event raised by the PanopticFactory contract.
type PanopticFactoryTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PanopticFactory *PanopticFactoryFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PanopticFactoryTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PanopticFactory.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryTransferSingleIterator{contract: _PanopticFactory.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PanopticFactory *PanopticFactoryFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *PanopticFactoryTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PanopticFactory.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PanopticFactoryTransferSingle)
				if err := _PanopticFactory.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PanopticFactory *PanopticFactoryFilterer) ParseTransferSingle(log types.Log) (*PanopticFactoryTransferSingle, error) {
	event := new(PanopticFactoryTransferSingle)
	if err := _PanopticFactory.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PanopticFactoryURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the PanopticFactory contract.
type PanopticFactoryURIIterator struct {
	Event *PanopticFactoryURI // Event containing the contract specifics and raw log

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
func (it *PanopticFactoryURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PanopticFactoryURI)
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
		it.Event = new(PanopticFactoryURI)
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
func (it *PanopticFactoryURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PanopticFactoryURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PanopticFactoryURI represents a URI event raised by the PanopticFactory contract.
type PanopticFactoryURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PanopticFactory *PanopticFactoryFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*PanopticFactoryURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PanopticFactory.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &PanopticFactoryURIIterator{contract: _PanopticFactory.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PanopticFactory *PanopticFactoryFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *PanopticFactoryURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PanopticFactory.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PanopticFactoryURI)
				if err := _PanopticFactory.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PanopticFactory *PanopticFactoryFilterer) ParseURI(log types.Log) (*PanopticFactoryURI, error) {
	event := new(PanopticFactoryURI)
	if err := _PanopticFactory.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
