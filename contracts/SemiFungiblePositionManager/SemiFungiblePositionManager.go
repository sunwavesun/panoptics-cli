// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SemiFungiblePositionManager

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

// SemiFungiblePositionManagerMetaData contains all meta data concerning the SemiFungiblePositionManager contract.
var SemiFungiblePositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CastingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTick\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"parameterType\",\"type\":\"uint256\"}],\"name\":\"InvalidTokenIdParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUniswapCallback\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotATokenRoll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptionsBalanceZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceBoundFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TicksNotInitializable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnderOverFlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UniswapPoolNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeRecipient\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uniswapPool\",\"type\":\"address\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"positionSize\",\"type\":\"uint128\"}],\"name\":\"TokenizedPositionBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"positionSize\",\"type\":\"uint128\"}],\"name\":\"TokenizedPositionMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"positionSize\",\"type\":\"uint128\"}],\"name\":\"TokenizedPositionRolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"positionSize\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"slippageTickLimitLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"slippageTickLimitHigh\",\"type\":\"int24\"}],\"name\":\"burnTokenizedPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"totalCollected\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalSwapped\",\"type\":\"int256\"},{\"internalType\":\"int24\",\"name\":\"newTick\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"univ3pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getAccountFeesBase\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"feesBase0\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"feesBase1\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"univ3pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accountLiquidities\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"univ3pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"atTick\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"isLong\",\"type\":\"uint256\"}],\"name\":\"getAccountPremium\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"premiumToken0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"premiumToken1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"univ3pool\",\"type\":\"address\"}],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"poolId\",\"type\":\"uint64\"}],\"name\":\"getUniswapV3PoolFromId\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"UniswapV3Pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"initializeAMMPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approvedForAll\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"positionSize\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"slippageTickLimitLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"slippageTickLimitHigh\",\"type\":\"int24\"}],\"name\":\"mintTokenizedPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"totalCollected\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalSwapped\",\"type\":\"int256\"},{\"internalType\":\"int24\",\"name\":\"newTick\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"positionSize\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"slippageTickLimitLow\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"slippageTickLimitHigh\",\"type\":\"int24\"}],\"name\":\"rollTokenizedPositions\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"totalCollectedBurn\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalSwappedBurn\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalCollectedMint\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalSwappedMint\",\"type\":\"int256\"},{\"internalType\":\"int24\",\"name\":\"newTick\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SemiFungiblePositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SemiFungiblePositionManagerMetaData.ABI instead.
var SemiFungiblePositionManagerABI = SemiFungiblePositionManagerMetaData.ABI

// SemiFungiblePositionManager is an auto generated Go binding around an Ethereum contract.
type SemiFungiblePositionManager struct {
	SemiFungiblePositionManagerCaller     // Read-only binding to the contract
	SemiFungiblePositionManagerTransactor // Write-only binding to the contract
	SemiFungiblePositionManagerFilterer   // Log filterer for contract events
}

// SemiFungiblePositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SemiFungiblePositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemiFungiblePositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SemiFungiblePositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemiFungiblePositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SemiFungiblePositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SemiFungiblePositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SemiFungiblePositionManagerSession struct {
	Contract     *SemiFungiblePositionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SemiFungiblePositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SemiFungiblePositionManagerCallerSession struct {
	Contract *SemiFungiblePositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// SemiFungiblePositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SemiFungiblePositionManagerTransactorSession struct {
	Contract     *SemiFungiblePositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// SemiFungiblePositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SemiFungiblePositionManagerRaw struct {
	Contract *SemiFungiblePositionManager // Generic contract binding to access the raw methods on
}

// SemiFungiblePositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SemiFungiblePositionManagerCallerRaw struct {
	Contract *SemiFungiblePositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SemiFungiblePositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SemiFungiblePositionManagerTransactorRaw struct {
	Contract *SemiFungiblePositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSemiFungiblePositionManager creates a new instance of SemiFungiblePositionManager, bound to a specific deployed contract.
func NewSemiFungiblePositionManager(address common.Address, backend bind.ContractBackend) (*SemiFungiblePositionManager, error) {
	contract, err := bindSemiFungiblePositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManager{SemiFungiblePositionManagerCaller: SemiFungiblePositionManagerCaller{contract: contract}, SemiFungiblePositionManagerTransactor: SemiFungiblePositionManagerTransactor{contract: contract}, SemiFungiblePositionManagerFilterer: SemiFungiblePositionManagerFilterer{contract: contract}}, nil
}

// NewSemiFungiblePositionManagerCaller creates a new read-only instance of SemiFungiblePositionManager, bound to a specific deployed contract.
func NewSemiFungiblePositionManagerCaller(address common.Address, caller bind.ContractCaller) (*SemiFungiblePositionManagerCaller, error) {
	contract, err := bindSemiFungiblePositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerCaller{contract: contract}, nil
}

// NewSemiFungiblePositionManagerTransactor creates a new write-only instance of SemiFungiblePositionManager, bound to a specific deployed contract.
func NewSemiFungiblePositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SemiFungiblePositionManagerTransactor, error) {
	contract, err := bindSemiFungiblePositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerTransactor{contract: contract}, nil
}

// NewSemiFungiblePositionManagerFilterer creates a new log filterer instance of SemiFungiblePositionManager, bound to a specific deployed contract.
func NewSemiFungiblePositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SemiFungiblePositionManagerFilterer, error) {
	contract, err := bindSemiFungiblePositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerFilterer{contract: contract}, nil
}

// bindSemiFungiblePositionManager binds a generic wrapper to an already deployed contract.
func bindSemiFungiblePositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SemiFungiblePositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SemiFungiblePositionManager.Contract.SemiFungiblePositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SemiFungiblePositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SemiFungiblePositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SemiFungiblePositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 tokenId) view returns(uint256 balance)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) BalanceOf(opts *bind.CallOpts, account common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "balanceOf", account, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 tokenId) view returns(uint256 balance)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) BalanceOf(account common.Address, tokenId *big.Int) (*big.Int, error) {
	return _SemiFungiblePositionManager.Contract.BalanceOf(&_SemiFungiblePositionManager.CallOpts, account, tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 tokenId) view returns(uint256 balance)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) BalanceOf(account common.Address, tokenId *big.Int) (*big.Int, error) {
	return _SemiFungiblePositionManager.Contract.BalanceOf(&_SemiFungiblePositionManager.CallOpts, account, tokenId)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _SemiFungiblePositionManager.Contract.BalanceOfBatch(&_SemiFungiblePositionManager.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _SemiFungiblePositionManager.Contract.BalanceOfBatch(&_SemiFungiblePositionManager.CallOpts, owners, ids)
}

// GetAccountFeesBase is a free data retrieval call binding the contract method 0xa734bda8.
//
// Solidity: function getAccountFeesBase(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper) view returns(int128 feesBase0, int128 feesBase1)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) GetAccountFeesBase(opts *bind.CallOpts, univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	FeesBase0 *big.Int
	FeesBase1 *big.Int
}, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "getAccountFeesBase", univ3pool, owner, tokenType, tickLower, tickUpper)

	outstruct := new(struct {
		FeesBase0 *big.Int
		FeesBase1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeesBase0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeesBase1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccountFeesBase is a free data retrieval call binding the contract method 0xa734bda8.
//
// Solidity: function getAccountFeesBase(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper) view returns(int128 feesBase0, int128 feesBase1)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) GetAccountFeesBase(univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	FeesBase0 *big.Int
	FeesBase1 *big.Int
}, error) {
	return _SemiFungiblePositionManager.Contract.GetAccountFeesBase(&_SemiFungiblePositionManager.CallOpts, univ3pool, owner, tokenType, tickLower, tickUpper)
}

// GetAccountFeesBase is a free data retrieval call binding the contract method 0xa734bda8.
//
// Solidity: function getAccountFeesBase(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper) view returns(int128 feesBase0, int128 feesBase1)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) GetAccountFeesBase(univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	FeesBase0 *big.Int
	FeesBase1 *big.Int
}, error) {
	return _SemiFungiblePositionManager.Contract.GetAccountFeesBase(&_SemiFungiblePositionManager.CallOpts, univ3pool, owner, tokenType, tickLower, tickUpper)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0xe53dec38.
//
// Solidity: function getAccountLiquidity(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper) view returns(uint256 accountLiquidities)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) GetAccountLiquidity(opts *bind.CallOpts, univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "getAccountLiquidity", univ3pool, owner, tokenType, tickLower, tickUpper)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0xe53dec38.
//
// Solidity: function getAccountLiquidity(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper) view returns(uint256 accountLiquidities)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) GetAccountLiquidity(univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	return _SemiFungiblePositionManager.Contract.GetAccountLiquidity(&_SemiFungiblePositionManager.CallOpts, univ3pool, owner, tokenType, tickLower, tickUpper)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0xe53dec38.
//
// Solidity: function getAccountLiquidity(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper) view returns(uint256 accountLiquidities)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) GetAccountLiquidity(univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int) (*big.Int, error) {
	return _SemiFungiblePositionManager.Contract.GetAccountLiquidity(&_SemiFungiblePositionManager.CallOpts, univ3pool, owner, tokenType, tickLower, tickUpper)
}

// GetAccountPremium is a free data retrieval call binding the contract method 0x3cc3c1da.
//
// Solidity: function getAccountPremium(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper, int24 atTick, uint256 isLong) view returns(uint128 premiumToken0, uint128 premiumToken1)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) GetAccountPremium(opts *bind.CallOpts, univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int, atTick *big.Int, isLong *big.Int) (struct {
	PremiumToken0 *big.Int
	PremiumToken1 *big.Int
}, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "getAccountPremium", univ3pool, owner, tokenType, tickLower, tickUpper, atTick, isLong)

	outstruct := new(struct {
		PremiumToken0 *big.Int
		PremiumToken1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PremiumToken0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PremiumToken1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccountPremium is a free data retrieval call binding the contract method 0x3cc3c1da.
//
// Solidity: function getAccountPremium(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper, int24 atTick, uint256 isLong) view returns(uint128 premiumToken0, uint128 premiumToken1)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) GetAccountPremium(univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int, atTick *big.Int, isLong *big.Int) (struct {
	PremiumToken0 *big.Int
	PremiumToken1 *big.Int
}, error) {
	return _SemiFungiblePositionManager.Contract.GetAccountPremium(&_SemiFungiblePositionManager.CallOpts, univ3pool, owner, tokenType, tickLower, tickUpper, atTick, isLong)
}

// GetAccountPremium is a free data retrieval call binding the contract method 0x3cc3c1da.
//
// Solidity: function getAccountPremium(address univ3pool, address owner, uint256 tokenType, int24 tickLower, int24 tickUpper, int24 atTick, uint256 isLong) view returns(uint128 premiumToken0, uint128 premiumToken1)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) GetAccountPremium(univ3pool common.Address, owner common.Address, tokenType *big.Int, tickLower *big.Int, tickUpper *big.Int, atTick *big.Int, isLong *big.Int) (struct {
	PremiumToken0 *big.Int
	PremiumToken1 *big.Int
}, error) {
	return _SemiFungiblePositionManager.Contract.GetAccountPremium(&_SemiFungiblePositionManager.CallOpts, univ3pool, owner, tokenType, tickLower, tickUpper, atTick, isLong)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address univ3pool) view returns(uint64 poolId)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) GetPoolId(opts *bind.CallOpts, univ3pool common.Address) (uint64, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "getPoolId", univ3pool)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address univ3pool) view returns(uint64 poolId)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) GetPoolId(univ3pool common.Address) (uint64, error) {
	return _SemiFungiblePositionManager.Contract.GetPoolId(&_SemiFungiblePositionManager.CallOpts, univ3pool)
}

// GetPoolId is a free data retrieval call binding the contract method 0xcaa9a08d.
//
// Solidity: function getPoolId(address univ3pool) view returns(uint64 poolId)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) GetPoolId(univ3pool common.Address) (uint64, error) {
	return _SemiFungiblePositionManager.Contract.GetPoolId(&_SemiFungiblePositionManager.CallOpts, univ3pool)
}

// GetUniswapV3PoolFromId is a free data retrieval call binding the contract method 0x3f8e156e.
//
// Solidity: function getUniswapV3PoolFromId(uint64 poolId) view returns(address UniswapV3Pool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) GetUniswapV3PoolFromId(opts *bind.CallOpts, poolId uint64) (common.Address, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "getUniswapV3PoolFromId", poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUniswapV3PoolFromId is a free data retrieval call binding the contract method 0x3f8e156e.
//
// Solidity: function getUniswapV3PoolFromId(uint64 poolId) view returns(address UniswapV3Pool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) GetUniswapV3PoolFromId(poolId uint64) (common.Address, error) {
	return _SemiFungiblePositionManager.Contract.GetUniswapV3PoolFromId(&_SemiFungiblePositionManager.CallOpts, poolId)
}

// GetUniswapV3PoolFromId is a free data retrieval call binding the contract method 0x3f8e156e.
//
// Solidity: function getUniswapV3PoolFromId(uint64 poolId) view returns(address UniswapV3Pool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) GetUniswapV3PoolFromId(poolId uint64) (common.Address, error) {
	return _SemiFungiblePositionManager.Contract.GetUniswapV3PoolFromId(&_SemiFungiblePositionManager.CallOpts, poolId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool approvedForAll)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool approvedForAll)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SemiFungiblePositionManager.Contract.IsApprovedForAll(&_SemiFungiblePositionManager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool approvedForAll)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SemiFungiblePositionManager.Contract.IsApprovedForAll(&_SemiFungiblePositionManager.CallOpts, owner, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SemiFungiblePositionManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SemiFungiblePositionManager.Contract.SupportsInterface(&_SemiFungiblePositionManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SemiFungiblePositionManager.Contract.SupportsInterface(&_SemiFungiblePositionManager.CallOpts, interfaceId)
}

// BurnTokenizedPosition is a paid mutator transaction binding the contract method 0x266601a3.
//
// Solidity: function burnTokenizedPosition(uint256 tokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollected, int256 totalSwapped, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) BurnTokenizedPosition(opts *bind.TransactOpts, tokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "burnTokenizedPosition", tokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// BurnTokenizedPosition is a paid mutator transaction binding the contract method 0x266601a3.
//
// Solidity: function burnTokenizedPosition(uint256 tokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollected, int256 totalSwapped, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) BurnTokenizedPosition(tokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.BurnTokenizedPosition(&_SemiFungiblePositionManager.TransactOpts, tokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// BurnTokenizedPosition is a paid mutator transaction binding the contract method 0x266601a3.
//
// Solidity: function burnTokenizedPosition(uint256 tokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollected, int256 totalSwapped, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) BurnTokenizedPosition(tokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.BurnTokenizedPosition(&_SemiFungiblePositionManager.TransactOpts, tokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// InitializeAMMPool is a paid mutator transaction binding the contract method 0xc035e242.
//
// Solidity: function initializeAMMPool(address token0, address token1, uint24 fee) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) InitializeAMMPool(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "initializeAMMPool", token0, token1, fee)
}

// InitializeAMMPool is a paid mutator transaction binding the contract method 0xc035e242.
//
// Solidity: function initializeAMMPool(address token0, address token1, uint24 fee) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) InitializeAMMPool(token0 common.Address, token1 common.Address, fee *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.InitializeAMMPool(&_SemiFungiblePositionManager.TransactOpts, token0, token1, fee)
}

// InitializeAMMPool is a paid mutator transaction binding the contract method 0xc035e242.
//
// Solidity: function initializeAMMPool(address token0, address token1, uint24 fee) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) InitializeAMMPool(token0 common.Address, token1 common.Address, fee *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.InitializeAMMPool(&_SemiFungiblePositionManager.TransactOpts, token0, token1, fee)
}

// MintTokenizedPosition is a paid mutator transaction binding the contract method 0xfdf90cc1.
//
// Solidity: function mintTokenizedPosition(uint256 tokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollected, int256 totalSwapped, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) MintTokenizedPosition(opts *bind.TransactOpts, tokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "mintTokenizedPosition", tokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// MintTokenizedPosition is a paid mutator transaction binding the contract method 0xfdf90cc1.
//
// Solidity: function mintTokenizedPosition(uint256 tokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollected, int256 totalSwapped, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) MintTokenizedPosition(tokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.MintTokenizedPosition(&_SemiFungiblePositionManager.TransactOpts, tokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// MintTokenizedPosition is a paid mutator transaction binding the contract method 0xfdf90cc1.
//
// Solidity: function mintTokenizedPosition(uint256 tokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollected, int256 totalSwapped, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) MintTokenizedPosition(tokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.MintTokenizedPosition(&_SemiFungiblePositionManager.TransactOpts, tokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.Multicall(&_SemiFungiblePositionManager.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.Multicall(&_SemiFungiblePositionManager.TransactOpts, data)
}

// RollTokenizedPositions is a paid mutator transaction binding the contract method 0xa350438a.
//
// Solidity: function rollTokenizedPositions(uint256 oldTokenId, uint256 newTokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollectedBurn, int256 totalSwappedBurn, int256 totalCollectedMint, int256 totalSwappedMint, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) RollTokenizedPositions(opts *bind.TransactOpts, oldTokenId *big.Int, newTokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "rollTokenizedPositions", oldTokenId, newTokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// RollTokenizedPositions is a paid mutator transaction binding the contract method 0xa350438a.
//
// Solidity: function rollTokenizedPositions(uint256 oldTokenId, uint256 newTokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollectedBurn, int256 totalSwappedBurn, int256 totalCollectedMint, int256 totalSwappedMint, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) RollTokenizedPositions(oldTokenId *big.Int, newTokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.RollTokenizedPositions(&_SemiFungiblePositionManager.TransactOpts, oldTokenId, newTokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// RollTokenizedPositions is a paid mutator transaction binding the contract method 0xa350438a.
//
// Solidity: function rollTokenizedPositions(uint256 oldTokenId, uint256 newTokenId, uint128 positionSize, int24 slippageTickLimitLow, int24 slippageTickLimitHigh) returns(int256 totalCollectedBurn, int256 totalSwappedBurn, int256 totalCollectedMint, int256 totalSwappedMint, int24 newTick)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) RollTokenizedPositions(oldTokenId *big.Int, newTokenId *big.Int, positionSize *big.Int, slippageTickLimitLow *big.Int, slippageTickLimitHigh *big.Int) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.RollTokenizedPositions(&_SemiFungiblePositionManager.TransactOpts, oldTokenId, newTokenId, positionSize, slippageTickLimitLow, slippageTickLimitHigh)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SafeBatchTransferFrom(&_SemiFungiblePositionManager.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SafeBatchTransferFrom(&_SemiFungiblePositionManager.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SafeTransferFrom(&_SemiFungiblePositionManager.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SafeTransferFrom(&_SemiFungiblePositionManager.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SetApprovalForAll(&_SemiFungiblePositionManager.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.SetApprovalForAll(&_SemiFungiblePositionManager.TransactOpts, operator, approved)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.UniswapV3MintCallback(&_SemiFungiblePositionManager.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.UniswapV3MintCallback(&_SemiFungiblePositionManager.TransactOpts, amount0Owed, amount1Owed, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.UniswapV3SwapCallback(&_SemiFungiblePositionManager.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _SemiFungiblePositionManager.Contract.UniswapV3SwapCallback(&_SemiFungiblePositionManager.TransactOpts, amount0Delta, amount1Delta, data)
}

// SemiFungiblePositionManagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerApprovalForAllIterator struct {
	Event *SemiFungiblePositionManagerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SemiFungiblePositionManagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SemiFungiblePositionManagerApprovalForAll)
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
		it.Event = new(SemiFungiblePositionManagerApprovalForAll)
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
func (it *SemiFungiblePositionManagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SemiFungiblePositionManagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SemiFungiblePositionManagerApprovalForAll represents a ApprovalForAll event raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SemiFungiblePositionManagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerApprovalForAllIterator{contract: _SemiFungiblePositionManager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SemiFungiblePositionManagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SemiFungiblePositionManagerApprovalForAll)
				if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) ParseApprovalForAll(log types.Log) (*SemiFungiblePositionManagerApprovalForAll, error) {
	event := new(SemiFungiblePositionManagerApprovalForAll)
	if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SemiFungiblePositionManagerPoolInitializedIterator is returned from FilterPoolInitialized and is used to iterate over the raw logs and unpacked data for PoolInitialized events raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerPoolInitializedIterator struct {
	Event *SemiFungiblePositionManagerPoolInitialized // Event containing the contract specifics and raw log

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
func (it *SemiFungiblePositionManagerPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SemiFungiblePositionManagerPoolInitialized)
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
		it.Event = new(SemiFungiblePositionManagerPoolInitialized)
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
func (it *SemiFungiblePositionManagerPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SemiFungiblePositionManagerPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SemiFungiblePositionManagerPoolInitialized represents a PoolInitialized event raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerPoolInitialized struct {
	UniswapPool common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolInitialized is a free log retrieval operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed uniswapPool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) FilterPoolInitialized(opts *bind.FilterOpts, uniswapPool []common.Address) (*SemiFungiblePositionManagerPoolInitializedIterator, error) {

	var uniswapPoolRule []interface{}
	for _, uniswapPoolItem := range uniswapPool {
		uniswapPoolRule = append(uniswapPoolRule, uniswapPoolItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.FilterLogs(opts, "PoolInitialized", uniswapPoolRule)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerPoolInitializedIterator{contract: _SemiFungiblePositionManager.contract, event: "PoolInitialized", logs: logs, sub: sub}, nil
}

// WatchPoolInitialized is a free log subscription operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed uniswapPool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) WatchPoolInitialized(opts *bind.WatchOpts, sink chan<- *SemiFungiblePositionManagerPoolInitialized, uniswapPool []common.Address) (event.Subscription, error) {

	var uniswapPoolRule []interface{}
	for _, uniswapPoolItem := range uniswapPool {
		uniswapPoolRule = append(uniswapPoolRule, uniswapPoolItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.WatchLogs(opts, "PoolInitialized", uniswapPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SemiFungiblePositionManagerPoolInitialized)
				if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
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

// ParsePoolInitialized is a log parse operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed uniswapPool)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) ParsePoolInitialized(log types.Log) (*SemiFungiblePositionManagerPoolInitialized, error) {
	event := new(SemiFungiblePositionManagerPoolInitialized)
	if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SemiFungiblePositionManagerTokenizedPositionBurntIterator is returned from FilterTokenizedPositionBurnt and is used to iterate over the raw logs and unpacked data for TokenizedPositionBurnt events raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTokenizedPositionBurntIterator struct {
	Event *SemiFungiblePositionManagerTokenizedPositionBurnt // Event containing the contract specifics and raw log

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
func (it *SemiFungiblePositionManagerTokenizedPositionBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SemiFungiblePositionManagerTokenizedPositionBurnt)
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
		it.Event = new(SemiFungiblePositionManagerTokenizedPositionBurnt)
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
func (it *SemiFungiblePositionManagerTokenizedPositionBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SemiFungiblePositionManagerTokenizedPositionBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SemiFungiblePositionManagerTokenizedPositionBurnt represents a TokenizedPositionBurnt event raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTokenizedPositionBurnt struct {
	Recipient    common.Address
	TokenId      *big.Int
	PositionSize *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenizedPositionBurnt is a free log retrieval operation binding the contract event 0x23833bcc608a225445893342669a9bc035de4ef96cc1edd47cead470f1f7817a.
//
// Solidity: event TokenizedPositionBurnt(address indexed recipient, uint256 indexed tokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) FilterTokenizedPositionBurnt(opts *bind.FilterOpts, recipient []common.Address, tokenId []*big.Int) (*SemiFungiblePositionManagerTokenizedPositionBurntIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.FilterLogs(opts, "TokenizedPositionBurnt", recipientRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerTokenizedPositionBurntIterator{contract: _SemiFungiblePositionManager.contract, event: "TokenizedPositionBurnt", logs: logs, sub: sub}, nil
}

// WatchTokenizedPositionBurnt is a free log subscription operation binding the contract event 0x23833bcc608a225445893342669a9bc035de4ef96cc1edd47cead470f1f7817a.
//
// Solidity: event TokenizedPositionBurnt(address indexed recipient, uint256 indexed tokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) WatchTokenizedPositionBurnt(opts *bind.WatchOpts, sink chan<- *SemiFungiblePositionManagerTokenizedPositionBurnt, recipient []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.WatchLogs(opts, "TokenizedPositionBurnt", recipientRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SemiFungiblePositionManagerTokenizedPositionBurnt)
				if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TokenizedPositionBurnt", log); err != nil {
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

// ParseTokenizedPositionBurnt is a log parse operation binding the contract event 0x23833bcc608a225445893342669a9bc035de4ef96cc1edd47cead470f1f7817a.
//
// Solidity: event TokenizedPositionBurnt(address indexed recipient, uint256 indexed tokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) ParseTokenizedPositionBurnt(log types.Log) (*SemiFungiblePositionManagerTokenizedPositionBurnt, error) {
	event := new(SemiFungiblePositionManagerTokenizedPositionBurnt)
	if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TokenizedPositionBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SemiFungiblePositionManagerTokenizedPositionMintedIterator is returned from FilterTokenizedPositionMinted and is used to iterate over the raw logs and unpacked data for TokenizedPositionMinted events raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTokenizedPositionMintedIterator struct {
	Event *SemiFungiblePositionManagerTokenizedPositionMinted // Event containing the contract specifics and raw log

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
func (it *SemiFungiblePositionManagerTokenizedPositionMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SemiFungiblePositionManagerTokenizedPositionMinted)
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
		it.Event = new(SemiFungiblePositionManagerTokenizedPositionMinted)
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
func (it *SemiFungiblePositionManagerTokenizedPositionMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SemiFungiblePositionManagerTokenizedPositionMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SemiFungiblePositionManagerTokenizedPositionMinted represents a TokenizedPositionMinted event raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTokenizedPositionMinted struct {
	Caller       common.Address
	TokenId      *big.Int
	PositionSize *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenizedPositionMinted is a free log retrieval operation binding the contract event 0x16695fc037c20a8984b20ffc7aadd6ce10e62756f9ca42933a1aa6459ef064ca.
//
// Solidity: event TokenizedPositionMinted(address indexed caller, uint256 indexed tokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) FilterTokenizedPositionMinted(opts *bind.FilterOpts, caller []common.Address, tokenId []*big.Int) (*SemiFungiblePositionManagerTokenizedPositionMintedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.FilterLogs(opts, "TokenizedPositionMinted", callerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerTokenizedPositionMintedIterator{contract: _SemiFungiblePositionManager.contract, event: "TokenizedPositionMinted", logs: logs, sub: sub}, nil
}

// WatchTokenizedPositionMinted is a free log subscription operation binding the contract event 0x16695fc037c20a8984b20ffc7aadd6ce10e62756f9ca42933a1aa6459ef064ca.
//
// Solidity: event TokenizedPositionMinted(address indexed caller, uint256 indexed tokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) WatchTokenizedPositionMinted(opts *bind.WatchOpts, sink chan<- *SemiFungiblePositionManagerTokenizedPositionMinted, caller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.WatchLogs(opts, "TokenizedPositionMinted", callerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SemiFungiblePositionManagerTokenizedPositionMinted)
				if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TokenizedPositionMinted", log); err != nil {
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

// ParseTokenizedPositionMinted is a log parse operation binding the contract event 0x16695fc037c20a8984b20ffc7aadd6ce10e62756f9ca42933a1aa6459ef064ca.
//
// Solidity: event TokenizedPositionMinted(address indexed caller, uint256 indexed tokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) ParseTokenizedPositionMinted(log types.Log) (*SemiFungiblePositionManagerTokenizedPositionMinted, error) {
	event := new(SemiFungiblePositionManagerTokenizedPositionMinted)
	if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TokenizedPositionMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SemiFungiblePositionManagerTokenizedPositionRolledIterator is returned from FilterTokenizedPositionRolled and is used to iterate over the raw logs and unpacked data for TokenizedPositionRolled events raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTokenizedPositionRolledIterator struct {
	Event *SemiFungiblePositionManagerTokenizedPositionRolled // Event containing the contract specifics and raw log

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
func (it *SemiFungiblePositionManagerTokenizedPositionRolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SemiFungiblePositionManagerTokenizedPositionRolled)
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
		it.Event = new(SemiFungiblePositionManagerTokenizedPositionRolled)
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
func (it *SemiFungiblePositionManagerTokenizedPositionRolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SemiFungiblePositionManagerTokenizedPositionRolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SemiFungiblePositionManagerTokenizedPositionRolled represents a TokenizedPositionRolled event raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTokenizedPositionRolled struct {
	Recipient    common.Address
	OldTokenId   *big.Int
	NewTokenId   *big.Int
	PositionSize *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenizedPositionRolled is a free log retrieval operation binding the contract event 0x5bf352c233806ba7b9525da6b698c7a4b9a97512b4e222729d857fc7a0c280a2.
//
// Solidity: event TokenizedPositionRolled(address indexed recipient, uint256 indexed oldTokenId, uint256 indexed newTokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) FilterTokenizedPositionRolled(opts *bind.FilterOpts, recipient []common.Address, oldTokenId []*big.Int, newTokenId []*big.Int) (*SemiFungiblePositionManagerTokenizedPositionRolledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldTokenIdRule []interface{}
	for _, oldTokenIdItem := range oldTokenId {
		oldTokenIdRule = append(oldTokenIdRule, oldTokenIdItem)
	}
	var newTokenIdRule []interface{}
	for _, newTokenIdItem := range newTokenId {
		newTokenIdRule = append(newTokenIdRule, newTokenIdItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.FilterLogs(opts, "TokenizedPositionRolled", recipientRule, oldTokenIdRule, newTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerTokenizedPositionRolledIterator{contract: _SemiFungiblePositionManager.contract, event: "TokenizedPositionRolled", logs: logs, sub: sub}, nil
}

// WatchTokenizedPositionRolled is a free log subscription operation binding the contract event 0x5bf352c233806ba7b9525da6b698c7a4b9a97512b4e222729d857fc7a0c280a2.
//
// Solidity: event TokenizedPositionRolled(address indexed recipient, uint256 indexed oldTokenId, uint256 indexed newTokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) WatchTokenizedPositionRolled(opts *bind.WatchOpts, sink chan<- *SemiFungiblePositionManagerTokenizedPositionRolled, recipient []common.Address, oldTokenId []*big.Int, newTokenId []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldTokenIdRule []interface{}
	for _, oldTokenIdItem := range oldTokenId {
		oldTokenIdRule = append(oldTokenIdRule, oldTokenIdItem)
	}
	var newTokenIdRule []interface{}
	for _, newTokenIdItem := range newTokenId {
		newTokenIdRule = append(newTokenIdRule, newTokenIdItem)
	}

	logs, sub, err := _SemiFungiblePositionManager.contract.WatchLogs(opts, "TokenizedPositionRolled", recipientRule, oldTokenIdRule, newTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SemiFungiblePositionManagerTokenizedPositionRolled)
				if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TokenizedPositionRolled", log); err != nil {
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

// ParseTokenizedPositionRolled is a log parse operation binding the contract event 0x5bf352c233806ba7b9525da6b698c7a4b9a97512b4e222729d857fc7a0c280a2.
//
// Solidity: event TokenizedPositionRolled(address indexed recipient, uint256 indexed oldTokenId, uint256 indexed newTokenId, uint128 positionSize)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) ParseTokenizedPositionRolled(log types.Log) (*SemiFungiblePositionManagerTokenizedPositionRolled, error) {
	event := new(SemiFungiblePositionManagerTokenizedPositionRolled)
	if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TokenizedPositionRolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SemiFungiblePositionManagerTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTransferBatchIterator struct {
	Event *SemiFungiblePositionManagerTransferBatch // Event containing the contract specifics and raw log

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
func (it *SemiFungiblePositionManagerTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SemiFungiblePositionManagerTransferBatch)
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
		it.Event = new(SemiFungiblePositionManagerTransferBatch)
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
func (it *SemiFungiblePositionManagerTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SemiFungiblePositionManagerTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SemiFungiblePositionManagerTransferBatch represents a TransferBatch event raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*SemiFungiblePositionManagerTransferBatchIterator, error) {

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

	logs, sub, err := _SemiFungiblePositionManager.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerTransferBatchIterator{contract: _SemiFungiblePositionManager.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *SemiFungiblePositionManagerTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SemiFungiblePositionManager.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SemiFungiblePositionManagerTransferBatch)
				if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) ParseTransferBatch(log types.Log) (*SemiFungiblePositionManagerTransferBatch, error) {
	event := new(SemiFungiblePositionManagerTransferBatch)
	if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SemiFungiblePositionManagerTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTransferSingleIterator struct {
	Event *SemiFungiblePositionManagerTransferSingle // Event containing the contract specifics and raw log

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
func (it *SemiFungiblePositionManagerTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SemiFungiblePositionManagerTransferSingle)
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
		it.Event = new(SemiFungiblePositionManagerTransferSingle)
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
func (it *SemiFungiblePositionManagerTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SemiFungiblePositionManagerTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SemiFungiblePositionManagerTransferSingle represents a TransferSingle event raised by the SemiFungiblePositionManager contract.
type SemiFungiblePositionManagerTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*SemiFungiblePositionManagerTransferSingleIterator, error) {

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

	logs, sub, err := _SemiFungiblePositionManager.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SemiFungiblePositionManagerTransferSingleIterator{contract: _SemiFungiblePositionManager.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *SemiFungiblePositionManagerTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SemiFungiblePositionManager.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SemiFungiblePositionManagerTransferSingle)
				if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_SemiFungiblePositionManager *SemiFungiblePositionManagerFilterer) ParseTransferSingle(log types.Log) (*SemiFungiblePositionManagerTransferSingle, error) {
	event := new(SemiFungiblePositionManagerTransferSingle)
	if err := _SemiFungiblePositionManager.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
