// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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

// AdminMetaData contains all meta data concerning the Admin contract.
var AdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAicat\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aicat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"}],\"name\":\"buyAICAT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"buyArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"catArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getBuyArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBuyArrayLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getReqLpArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReqLpArrayLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpAmountTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpCatArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpUsdtArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidityAmount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reqLpArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"setBuyArray\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"setReqLpArray\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usdtArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AdminABI is the input ABI used to generate the binding from.
// Deprecated: Use AdminMetaData.ABI instead.
var AdminABI = AdminMetaData.ABI

// Admin is an auto generated Go binding around an Ethereum contract.
type Admin struct {
	AdminCaller     // Read-only binding to the contract
	AdminTransactor // Write-only binding to the contract
	AdminFilterer   // Log filterer for contract events
}

// AdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminSession struct {
	Contract     *Admin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminCallerSession struct {
	Contract *AdminCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminTransactorSession struct {
	Contract     *AdminTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminRaw struct {
	Contract *Admin // Generic contract binding to access the raw methods on
}

// AdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminCallerRaw struct {
	Contract *AdminCaller // Generic read-only contract binding to access the raw methods on
}

// AdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminTransactorRaw struct {
	Contract *AdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdmin creates a new instance of Admin, bound to a specific deployed contract.
func NewAdmin(address common.Address, backend bind.ContractBackend) (*Admin, error) {
	contract, err := bindAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Admin{AdminCaller: AdminCaller{contract: contract}, AdminTransactor: AdminTransactor{contract: contract}, AdminFilterer: AdminFilterer{contract: contract}}, nil
}

// NewAdminCaller creates a new read-only instance of Admin, bound to a specific deployed contract.
func NewAdminCaller(address common.Address, caller bind.ContractCaller) (*AdminCaller, error) {
	contract, err := bindAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminCaller{contract: contract}, nil
}

// NewAdminTransactor creates a new write-only instance of Admin, bound to a specific deployed contract.
func NewAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminTransactor, error) {
	contract, err := bindAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminTransactor{contract: contract}, nil
}

// NewAdminFilterer creates a new log filterer instance of Admin, bound to a specific deployed contract.
func NewAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminFilterer, error) {
	contract, err := bindAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminFilterer{contract: contract}, nil
}

// bindAdmin binds a generic wrapper to an already deployed contract.
func bindAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.AdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.AdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Admin *AdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Admin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Admin *AdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Admin *AdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Admin.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminSession) ADMINROLE() ([32]byte, error) {
	return _Admin.Contract.ADMINROLE(&_Admin.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) ADMINROLE() ([32]byte, error) {
	return _Admin.Contract.ADMINROLE(&_Admin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Admin.Contract.DEFAULTADMINROLE(&_Admin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Admin.Contract.DEFAULTADMINROLE(&_Admin.CallOpts)
}

// USERROLE is a free data retrieval call binding the contract method 0x13119161.
//
// Solidity: function USER_ROLE() view returns(bytes32)
func (_Admin *AdminCaller) USERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "USER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// USERROLE is a free data retrieval call binding the contract method 0x13119161.
//
// Solidity: function USER_ROLE() view returns(bytes32)
func (_Admin *AdminSession) USERROLE() ([32]byte, error) {
	return _Admin.Contract.USERROLE(&_Admin.CallOpts)
}

// USERROLE is a free data retrieval call binding the contract method 0x13119161.
//
// Solidity: function USER_ROLE() view returns(bytes32)
func (_Admin *AdminCallerSession) USERROLE() ([32]byte, error) {
	return _Admin.Contract.USERROLE(&_Admin.CallOpts)
}

// Aicat is a free data retrieval call binding the contract method 0x5c767d9a.
//
// Solidity: function aicat() view returns(address)
func (_Admin *AdminCaller) Aicat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "aicat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aicat is a free data retrieval call binding the contract method 0x5c767d9a.
//
// Solidity: function aicat() view returns(address)
func (_Admin *AdminSession) Aicat() (common.Address, error) {
	return _Admin.Contract.Aicat(&_Admin.CallOpts)
}

// Aicat is a free data retrieval call binding the contract method 0x5c767d9a.
//
// Solidity: function aicat() view returns(address)
func (_Admin *AdminCallerSession) Aicat() (common.Address, error) {
	return _Admin.Contract.Aicat(&_Admin.CallOpts)
}

// BuyArray is a free data retrieval call binding the contract method 0xad1961bc.
//
// Solidity: function buyArray(uint256 ) view returns(uint256)
func (_Admin *AdminCaller) BuyArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "buyArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyArray is a free data retrieval call binding the contract method 0xad1961bc.
//
// Solidity: function buyArray(uint256 ) view returns(uint256)
func (_Admin *AdminSession) BuyArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.BuyArray(&_Admin.CallOpts, arg0)
}

// BuyArray is a free data retrieval call binding the contract method 0xad1961bc.
//
// Solidity: function buyArray(uint256 ) view returns(uint256)
func (_Admin *AdminCallerSession) BuyArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.BuyArray(&_Admin.CallOpts, arg0)
}

// CatArray is a free data retrieval call binding the contract method 0xcbf10fb6.
//
// Solidity: function catArray(uint256 ) view returns(uint256)
func (_Admin *AdminCaller) CatArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "catArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CatArray is a free data retrieval call binding the contract method 0xcbf10fb6.
//
// Solidity: function catArray(uint256 ) view returns(uint256)
func (_Admin *AdminSession) CatArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.CatArray(&_Admin.CallOpts, arg0)
}

// CatArray is a free data retrieval call binding the contract method 0xcbf10fb6.
//
// Solidity: function catArray(uint256 ) view returns(uint256)
func (_Admin *AdminCallerSession) CatArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.CatArray(&_Admin.CallOpts, arg0)
}

// GetBuyArray is a free data retrieval call binding the contract method 0x3bfc6177.
//
// Solidity: function getBuyArray(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Admin *AdminCaller) GetBuyArray(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getBuyArray", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBuyArray is a free data retrieval call binding the contract method 0x3bfc6177.
//
// Solidity: function getBuyArray(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Admin *AdminSession) GetBuyArray(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Admin.Contract.GetBuyArray(&_Admin.CallOpts, startIndex, endIndex)
}

// GetBuyArray is a free data retrieval call binding the contract method 0x3bfc6177.
//
// Solidity: function getBuyArray(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Admin *AdminCallerSession) GetBuyArray(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Admin.Contract.GetBuyArray(&_Admin.CallOpts, startIndex, endIndex)
}

// GetBuyArrayLength is a free data retrieval call binding the contract method 0x59c2f8fe.
//
// Solidity: function getBuyArrayLength() view returns(uint256)
func (_Admin *AdminCaller) GetBuyArrayLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getBuyArrayLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyArrayLength is a free data retrieval call binding the contract method 0x59c2f8fe.
//
// Solidity: function getBuyArrayLength() view returns(uint256)
func (_Admin *AdminSession) GetBuyArrayLength() (*big.Int, error) {
	return _Admin.Contract.GetBuyArrayLength(&_Admin.CallOpts)
}

// GetBuyArrayLength is a free data retrieval call binding the contract method 0x59c2f8fe.
//
// Solidity: function getBuyArrayLength() view returns(uint256)
func (_Admin *AdminCallerSession) GetBuyArrayLength() (*big.Int, error) {
	return _Admin.Contract.GetBuyArrayLength(&_Admin.CallOpts)
}

// GetReqLpArray is a free data retrieval call binding the contract method 0x6bc09544.
//
// Solidity: function getReqLpArray(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Admin *AdminCaller) GetReqLpArray(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getReqLpArray", startIndex, endIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetReqLpArray is a free data retrieval call binding the contract method 0x6bc09544.
//
// Solidity: function getReqLpArray(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Admin *AdminSession) GetReqLpArray(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Admin.Contract.GetReqLpArray(&_Admin.CallOpts, startIndex, endIndex)
}

// GetReqLpArray is a free data retrieval call binding the contract method 0x6bc09544.
//
// Solidity: function getReqLpArray(uint256 startIndex, uint256 endIndex) view returns(uint256[])
func (_Admin *AdminCallerSession) GetReqLpArray(startIndex *big.Int, endIndex *big.Int) ([]*big.Int, error) {
	return _Admin.Contract.GetReqLpArray(&_Admin.CallOpts, startIndex, endIndex)
}

// GetReqLpArrayLength is a free data retrieval call binding the contract method 0x63917b68.
//
// Solidity: function getReqLpArrayLength() view returns(uint256)
func (_Admin *AdminCaller) GetReqLpArrayLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getReqLpArrayLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReqLpArrayLength is a free data retrieval call binding the contract method 0x63917b68.
//
// Solidity: function getReqLpArrayLength() view returns(uint256)
func (_Admin *AdminSession) GetReqLpArrayLength() (*big.Int, error) {
	return _Admin.Contract.GetReqLpArrayLength(&_Admin.CallOpts)
}

// GetReqLpArrayLength is a free data retrieval call binding the contract method 0x63917b68.
//
// Solidity: function getReqLpArrayLength() view returns(uint256)
func (_Admin *AdminCallerSession) GetReqLpArrayLength() (*big.Int, error) {
	return _Admin.Contract.GetReqLpArrayLength(&_Admin.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Admin *AdminCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Admin *AdminSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Admin.Contract.GetRoleAdmin(&_Admin.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Admin *AdminCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Admin.Contract.GetRoleAdmin(&_Admin.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Admin *AdminCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Admin *AdminSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Admin.Contract.HasRole(&_Admin.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Admin *AdminCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Admin.Contract.HasRole(&_Admin.CallOpts, role, account)
}

// LpAmount is a free data retrieval call binding the contract method 0x6fc9d181.
//
// Solidity: function lpAmount(address ) view returns(uint256)
func (_Admin *AdminCaller) LpAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "lpAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpAmount is a free data retrieval call binding the contract method 0x6fc9d181.
//
// Solidity: function lpAmount(address ) view returns(uint256)
func (_Admin *AdminSession) LpAmount(arg0 common.Address) (*big.Int, error) {
	return _Admin.Contract.LpAmount(&_Admin.CallOpts, arg0)
}

// LpAmount is a free data retrieval call binding the contract method 0x6fc9d181.
//
// Solidity: function lpAmount(address ) view returns(uint256)
func (_Admin *AdminCallerSession) LpAmount(arg0 common.Address) (*big.Int, error) {
	return _Admin.Contract.LpAmount(&_Admin.CallOpts, arg0)
}

// LpAmountTotal is a free data retrieval call binding the contract method 0xaa1bc6b5.
//
// Solidity: function lpAmountTotal() view returns(uint256)
func (_Admin *AdminCaller) LpAmountTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "lpAmountTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpAmountTotal is a free data retrieval call binding the contract method 0xaa1bc6b5.
//
// Solidity: function lpAmountTotal() view returns(uint256)
func (_Admin *AdminSession) LpAmountTotal() (*big.Int, error) {
	return _Admin.Contract.LpAmountTotal(&_Admin.CallOpts)
}

// LpAmountTotal is a free data retrieval call binding the contract method 0xaa1bc6b5.
//
// Solidity: function lpAmountTotal() view returns(uint256)
func (_Admin *AdminCallerSession) LpAmountTotal() (*big.Int, error) {
	return _Admin.Contract.LpAmountTotal(&_Admin.CallOpts)
}

// LpArray is a free data retrieval call binding the contract method 0x3a9072a9.
//
// Solidity: function lpArray(uint256 ) view returns(uint256)
func (_Admin *AdminCaller) LpArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "lpArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpArray is a free data retrieval call binding the contract method 0x3a9072a9.
//
// Solidity: function lpArray(uint256 ) view returns(uint256)
func (_Admin *AdminSession) LpArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.LpArray(&_Admin.CallOpts, arg0)
}

// LpArray is a free data retrieval call binding the contract method 0x3a9072a9.
//
// Solidity: function lpArray(uint256 ) view returns(uint256)
func (_Admin *AdminCallerSession) LpArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.LpArray(&_Admin.CallOpts, arg0)
}

// LpCatArray is a free data retrieval call binding the contract method 0x572e7abd.
//
// Solidity: function lpCatArray(uint256 ) view returns(uint256)
func (_Admin *AdminCaller) LpCatArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "lpCatArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpCatArray is a free data retrieval call binding the contract method 0x572e7abd.
//
// Solidity: function lpCatArray(uint256 ) view returns(uint256)
func (_Admin *AdminSession) LpCatArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.LpCatArray(&_Admin.CallOpts, arg0)
}

// LpCatArray is a free data retrieval call binding the contract method 0x572e7abd.
//
// Solidity: function lpCatArray(uint256 ) view returns(uint256)
func (_Admin *AdminCallerSession) LpCatArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.LpCatArray(&_Admin.CallOpts, arg0)
}

// LpUsdtArray is a free data retrieval call binding the contract method 0x6d799c2d.
//
// Solidity: function lpUsdtArray(uint256 ) view returns(uint256)
func (_Admin *AdminCaller) LpUsdtArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "lpUsdtArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpUsdtArray is a free data retrieval call binding the contract method 0x6d799c2d.
//
// Solidity: function lpUsdtArray(uint256 ) view returns(uint256)
func (_Admin *AdminSession) LpUsdtArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.LpUsdtArray(&_Admin.CallOpts, arg0)
}

// LpUsdtArray is a free data retrieval call binding the contract method 0x6d799c2d.
//
// Solidity: function lpUsdtArray(uint256 ) view returns(uint256)
func (_Admin *AdminCallerSession) LpUsdtArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.LpUsdtArray(&_Admin.CallOpts, arg0)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_Admin *AdminCaller) Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_Admin *AdminSession) Pair() (common.Address, error) {
	return _Admin.Contract.Pair(&_Admin.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_Admin *AdminCallerSession) Pair() (common.Address, error) {
	return _Admin.Contract.Pair(&_Admin.CallOpts)
}

// ReqLpArray is a free data retrieval call binding the contract method 0xa3b49d7a.
//
// Solidity: function reqLpArray(uint256 ) view returns(uint256)
func (_Admin *AdminCaller) ReqLpArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "reqLpArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReqLpArray is a free data retrieval call binding the contract method 0xa3b49d7a.
//
// Solidity: function reqLpArray(uint256 ) view returns(uint256)
func (_Admin *AdminSession) ReqLpArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.ReqLpArray(&_Admin.CallOpts, arg0)
}

// ReqLpArray is a free data retrieval call binding the contract method 0xa3b49d7a.
//
// Solidity: function reqLpArray(uint256 ) view returns(uint256)
func (_Admin *AdminCallerSession) ReqLpArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.ReqLpArray(&_Admin.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Admin *AdminCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Admin *AdminSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Admin.Contract.SupportsInterface(&_Admin.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Admin *AdminCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Admin.Contract.SupportsInterface(&_Admin.CallOpts, interfaceId)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Admin *AdminCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Admin *AdminSession) UniswapV2Router() (common.Address, error) {
	return _Admin.Contract.UniswapV2Router(&_Admin.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Admin *AdminCallerSession) UniswapV2Router() (common.Address, error) {
	return _Admin.Contract.UniswapV2Router(&_Admin.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Admin *AdminCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Admin *AdminSession) Usdt() (common.Address, error) {
	return _Admin.Contract.Usdt(&_Admin.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Admin *AdminCallerSession) Usdt() (common.Address, error) {
	return _Admin.Contract.Usdt(&_Admin.CallOpts)
}

// UsdtArray is a free data retrieval call binding the contract method 0x317c3410.
//
// Solidity: function usdtArray(uint256 ) view returns(uint256)
func (_Admin *AdminCaller) UsdtArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Admin.contract.Call(opts, &out, "usdtArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdtArray is a free data retrieval call binding the contract method 0x317c3410.
//
// Solidity: function usdtArray(uint256 ) view returns(uint256)
func (_Admin *AdminSession) UsdtArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.UsdtArray(&_Admin.CallOpts, arg0)
}

// UsdtArray is a free data retrieval call binding the contract method 0x317c3410.
//
// Solidity: function usdtArray(uint256 ) view returns(uint256)
func (_Admin *AdminCallerSession) UsdtArray(arg0 *big.Int) (*big.Int, error) {
	return _Admin.Contract.UsdtArray(&_Admin.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x95164bae.
//
// Solidity: function addLiquidity(uint256 orderId, address user, uint256 amountUsdt, uint256 amountAicat) returns()
func (_Admin *AdminTransactor) AddLiquidity(opts *bind.TransactOpts, orderId *big.Int, user common.Address, amountUsdt *big.Int, amountAicat *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "addLiquidity", orderId, user, amountUsdt, amountAicat)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x95164bae.
//
// Solidity: function addLiquidity(uint256 orderId, address user, uint256 amountUsdt, uint256 amountAicat) returns()
func (_Admin *AdminSession) AddLiquidity(orderId *big.Int, user common.Address, amountUsdt *big.Int, amountAicat *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.AddLiquidity(&_Admin.TransactOpts, orderId, user, amountUsdt, amountAicat)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x95164bae.
//
// Solidity: function addLiquidity(uint256 orderId, address user, uint256 amountUsdt, uint256 amountAicat) returns()
func (_Admin *AdminTransactorSession) AddLiquidity(orderId *big.Int, user common.Address, amountUsdt *big.Int, amountAicat *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.AddLiquidity(&_Admin.TransactOpts, orderId, user, amountUsdt, amountAicat)
}

// BuyAICAT is a paid mutator transaction binding the contract method 0x3c3b0ecf.
//
// Solidity: function buyAICAT(uint256 orderId, uint256 usdtAmount) returns()
func (_Admin *AdminTransactor) BuyAICAT(opts *bind.TransactOpts, orderId *big.Int, usdtAmount *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "buyAICAT", orderId, usdtAmount)
}

// BuyAICAT is a paid mutator transaction binding the contract method 0x3c3b0ecf.
//
// Solidity: function buyAICAT(uint256 orderId, uint256 usdtAmount) returns()
func (_Admin *AdminSession) BuyAICAT(orderId *big.Int, usdtAmount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.BuyAICAT(&_Admin.TransactOpts, orderId, usdtAmount)
}

// BuyAICAT is a paid mutator transaction binding the contract method 0x3c3b0ecf.
//
// Solidity: function buyAICAT(uint256 orderId, uint256 usdtAmount) returns()
func (_Admin *AdminTransactorSession) BuyAICAT(orderId *big.Int, usdtAmount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.BuyAICAT(&_Admin.TransactOpts, orderId, usdtAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Admin *AdminSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.GrantRole(&_Admin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.GrantRole(&_Admin.TransactOpts, role, account)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc4ccdeea.
//
// Solidity: function removeLiquidity(uint256 orderId, address user, uint256 liquidityAmount) returns()
func (_Admin *AdminTransactor) RemoveLiquidity(opts *bind.TransactOpts, orderId *big.Int, user common.Address, liquidityAmount *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "removeLiquidity", orderId, user, liquidityAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc4ccdeea.
//
// Solidity: function removeLiquidity(uint256 orderId, address user, uint256 liquidityAmount) returns()
func (_Admin *AdminSession) RemoveLiquidity(orderId *big.Int, user common.Address, liquidityAmount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.RemoveLiquidity(&_Admin.TransactOpts, orderId, user, liquidityAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xc4ccdeea.
//
// Solidity: function removeLiquidity(uint256 orderId, address user, uint256 liquidityAmount) returns()
func (_Admin *AdminTransactorSession) RemoveLiquidity(orderId *big.Int, user common.Address, liquidityAmount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.RemoveLiquidity(&_Admin.TransactOpts, orderId, user, liquidityAmount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Admin *AdminTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Admin *AdminSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RenounceRole(&_Admin.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Admin *AdminTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RenounceRole(&_Admin.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Admin *AdminSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RevokeRole(&_Admin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Admin *AdminTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Admin.Contract.RevokeRole(&_Admin.TransactOpts, role, account)
}

// SetBuyArray is a paid mutator transaction binding the contract method 0x26a31e97.
//
// Solidity: function setBuyArray(uint256 orderId) payable returns()
func (_Admin *AdminTransactor) SetBuyArray(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "setBuyArray", orderId)
}

// SetBuyArray is a paid mutator transaction binding the contract method 0x26a31e97.
//
// Solidity: function setBuyArray(uint256 orderId) payable returns()
func (_Admin *AdminSession) SetBuyArray(orderId *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetBuyArray(&_Admin.TransactOpts, orderId)
}

// SetBuyArray is a paid mutator transaction binding the contract method 0x26a31e97.
//
// Solidity: function setBuyArray(uint256 orderId) payable returns()
func (_Admin *AdminTransactorSession) SetBuyArray(orderId *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetBuyArray(&_Admin.TransactOpts, orderId)
}

// SetReqLpArray is a paid mutator transaction binding the contract method 0xf5dc7465.
//
// Solidity: function setReqLpArray(uint256 orderId) payable returns()
func (_Admin *AdminTransactor) SetReqLpArray(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "setReqLpArray", orderId)
}

// SetReqLpArray is a paid mutator transaction binding the contract method 0xf5dc7465.
//
// Solidity: function setReqLpArray(uint256 orderId) payable returns()
func (_Admin *AdminSession) SetReqLpArray(orderId *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetReqLpArray(&_Admin.TransactOpts, orderId)
}

// SetReqLpArray is a paid mutator transaction binding the contract method 0xf5dc7465.
//
// Solidity: function setReqLpArray(uint256 orderId) payable returns()
func (_Admin *AdminTransactorSession) SetReqLpArray(orderId *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.SetReqLpArray(&_Admin.TransactOpts, orderId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_Admin *AdminTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "withdraw", to, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_Admin *AdminSession) Withdraw(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.Withdraw(&_Admin.TransactOpts, to, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_Admin *AdminTransactorSession) Withdraw(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.Withdraw(&_Admin.TransactOpts, to, token, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Admin *AdminTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Admin.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Admin *AdminSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.WithdrawETH(&_Admin.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Admin *AdminTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Admin.Contract.WithdrawETH(&_Admin.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Admin *AdminTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Admin.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Admin *AdminSession) Receive() (*types.Transaction, error) {
	return _Admin.Contract.Receive(&_Admin.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Admin *AdminTransactorSession) Receive() (*types.Transaction, error) {
	return _Admin.Contract.Receive(&_Admin.TransactOpts)
}

// AdminRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Admin contract.
type AdminRoleAdminChangedIterator struct {
	Event *AdminRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AdminRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRoleAdminChanged)
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
		it.Event = new(AdminRoleAdminChanged)
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
func (it *AdminRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRoleAdminChanged represents a RoleAdminChanged event raised by the Admin contract.
type AdminRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Admin *AdminFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AdminRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AdminRoleAdminChangedIterator{contract: _Admin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Admin *AdminFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AdminRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRoleAdminChanged)
				if err := _Admin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Admin *AdminFilterer) ParseRoleAdminChanged(log types.Log) (*AdminRoleAdminChanged, error) {
	event := new(AdminRoleAdminChanged)
	if err := _Admin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Admin contract.
type AdminRoleGrantedIterator struct {
	Event *AdminRoleGranted // Event containing the contract specifics and raw log

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
func (it *AdminRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRoleGranted)
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
		it.Event = new(AdminRoleGranted)
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
func (it *AdminRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRoleGranted represents a RoleGranted event raised by the Admin contract.
type AdminRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminRoleGrantedIterator{contract: _Admin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AdminRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRoleGranted)
				if err := _Admin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) ParseRoleGranted(log types.Log) (*AdminRoleGranted, error) {
	event := new(AdminRoleGranted)
	if err := _Admin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Admin contract.
type AdminRoleRevokedIterator struct {
	Event *AdminRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AdminRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminRoleRevoked)
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
		it.Event = new(AdminRoleRevoked)
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
func (it *AdminRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminRoleRevoked represents a RoleRevoked event raised by the Admin contract.
type AdminRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdminRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Admin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdminRoleRevokedIterator{contract: _Admin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AdminRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Admin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminRoleRevoked)
				if err := _Admin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Admin *AdminFilterer) ParseRoleRevoked(log types.Log) (*AdminRoleRevoked, error) {
	event := new(AdminRoleRevoked)
	if err := _Admin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
