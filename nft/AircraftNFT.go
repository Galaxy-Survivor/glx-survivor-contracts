// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

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

// AircraftNFTMetaData contains all meta data concerning the AircraftNFT contract.
var AircraftNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"aircraftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"componentId\",\"type\":\"uint256\"}],\"name\":\"UpgradeAircraft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"aircraftId\",\"type\":\"uint256\"}],\"name\":\"event_MintedAircraft\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LUCKY_BOX_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_COMPONENT_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aircraftIdToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAircrafts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintAircraft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"mintComponent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintLuckyBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"aircraftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"componentId\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AircraftNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use AircraftNFTMetaData.ABI instead.
var AircraftNFTABI = AircraftNFTMetaData.ABI

// AircraftNFT is an auto generated Go binding around an Ethereum contract.
type AircraftNFT struct {
	AircraftNFTCaller     // Read-only binding to the contract
	AircraftNFTTransactor // Write-only binding to the contract
	AircraftNFTFilterer   // Log filterer for contract events
}

// AircraftNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type AircraftNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AircraftNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AircraftNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AircraftNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AircraftNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AircraftNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AircraftNFTSession struct {
	Contract     *AircraftNFT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AircraftNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AircraftNFTCallerSession struct {
	Contract *AircraftNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AircraftNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AircraftNFTTransactorSession struct {
	Contract     *AircraftNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AircraftNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type AircraftNFTRaw struct {
	Contract *AircraftNFT // Generic contract binding to access the raw methods on
}

// AircraftNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AircraftNFTCallerRaw struct {
	Contract *AircraftNFTCaller // Generic read-only contract binding to access the raw methods on
}

// AircraftNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AircraftNFTTransactorRaw struct {
	Contract *AircraftNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAircraftNFT creates a new instance of AircraftNFT, bound to a specific deployed contract.
func NewAircraftNFT(address common.Address, backend bind.ContractBackend) (*AircraftNFT, error) {
	contract, err := bindAircraftNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AircraftNFT{AircraftNFTCaller: AircraftNFTCaller{contract: contract}, AircraftNFTTransactor: AircraftNFTTransactor{contract: contract}, AircraftNFTFilterer: AircraftNFTFilterer{contract: contract}}, nil
}

// NewAircraftNFTCaller creates a new read-only instance of AircraftNFT, bound to a specific deployed contract.
func NewAircraftNFTCaller(address common.Address, caller bind.ContractCaller) (*AircraftNFTCaller, error) {
	contract, err := bindAircraftNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTCaller{contract: contract}, nil
}

// NewAircraftNFTTransactor creates a new write-only instance of AircraftNFT, bound to a specific deployed contract.
func NewAircraftNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*AircraftNFTTransactor, error) {
	contract, err := bindAircraftNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTTransactor{contract: contract}, nil
}

// NewAircraftNFTFilterer creates a new log filterer instance of AircraftNFT, bound to a specific deployed contract.
func NewAircraftNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*AircraftNFTFilterer, error) {
	contract, err := bindAircraftNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTFilterer{contract: contract}, nil
}

// bindAircraftNFT binds a generic wrapper to an already deployed contract.
func bindAircraftNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AircraftNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AircraftNFT *AircraftNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AircraftNFT.Contract.AircraftNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AircraftNFT *AircraftNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AircraftNFT.Contract.AircraftNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AircraftNFT *AircraftNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AircraftNFT.Contract.AircraftNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AircraftNFT *AircraftNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AircraftNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AircraftNFT *AircraftNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AircraftNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AircraftNFT *AircraftNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AircraftNFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AircraftNFT.Contract.DEFAULTADMINROLE(&_AircraftNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AircraftNFT.Contract.DEFAULTADMINROLE(&_AircraftNFT.CallOpts)
}

// LUCKYBOXID is a free data retrieval call binding the contract method 0xd2bcd2dd.
//
// Solidity: function LUCKY_BOX_ID() view returns(uint256)
func (_AircraftNFT *AircraftNFTCaller) LUCKYBOXID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "LUCKY_BOX_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LUCKYBOXID is a free data retrieval call binding the contract method 0xd2bcd2dd.
//
// Solidity: function LUCKY_BOX_ID() view returns(uint256)
func (_AircraftNFT *AircraftNFTSession) LUCKYBOXID() (*big.Int, error) {
	return _AircraftNFT.Contract.LUCKYBOXID(&_AircraftNFT.CallOpts)
}

// LUCKYBOXID is a free data retrieval call binding the contract method 0xd2bcd2dd.
//
// Solidity: function LUCKY_BOX_ID() view returns(uint256)
func (_AircraftNFT *AircraftNFTCallerSession) LUCKYBOXID() (*big.Int, error) {
	return _AircraftNFT.Contract.LUCKYBOXID(&_AircraftNFT.CallOpts)
}

// MAXCOMPONENTID is a free data retrieval call binding the contract method 0x6320ccfd.
//
// Solidity: function MAX_COMPONENT_ID() view returns(uint256)
func (_AircraftNFT *AircraftNFTCaller) MAXCOMPONENTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "MAX_COMPONENT_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOMPONENTID is a free data retrieval call binding the contract method 0x6320ccfd.
//
// Solidity: function MAX_COMPONENT_ID() view returns(uint256)
func (_AircraftNFT *AircraftNFTSession) MAXCOMPONENTID() (*big.Int, error) {
	return _AircraftNFT.Contract.MAXCOMPONENTID(&_AircraftNFT.CallOpts)
}

// MAXCOMPONENTID is a free data retrieval call binding the contract method 0x6320ccfd.
//
// Solidity: function MAX_COMPONENT_ID() view returns(uint256)
func (_AircraftNFT *AircraftNFTCallerSession) MAXCOMPONENTID() (*big.Int, error) {
	return _AircraftNFT.Contract.MAXCOMPONENTID(&_AircraftNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTSession) MINTERROLE() ([32]byte, error) {
	return _AircraftNFT.Contract.MINTERROLE(&_AircraftNFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTCallerSession) MINTERROLE() ([32]byte, error) {
	return _AircraftNFT.Contract.MINTERROLE(&_AircraftNFT.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTSession) PAUSERROLE() ([32]byte, error) {
	return _AircraftNFT.Contract.PAUSERROLE(&_AircraftNFT.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AircraftNFT *AircraftNFTCallerSession) PAUSERROLE() ([32]byte, error) {
	return _AircraftNFT.Contract.PAUSERROLE(&_AircraftNFT.CallOpts)
}

// AircraftIdToOwner is a free data retrieval call binding the contract method 0xd2411f57.
//
// Solidity: function aircraftIdToOwner(uint256 ) view returns(address)
func (_AircraftNFT *AircraftNFTCaller) AircraftIdToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "aircraftIdToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AircraftIdToOwner is a free data retrieval call binding the contract method 0xd2411f57.
//
// Solidity: function aircraftIdToOwner(uint256 ) view returns(address)
func (_AircraftNFT *AircraftNFTSession) AircraftIdToOwner(arg0 *big.Int) (common.Address, error) {
	return _AircraftNFT.Contract.AircraftIdToOwner(&_AircraftNFT.CallOpts, arg0)
}

// AircraftIdToOwner is a free data retrieval call binding the contract method 0xd2411f57.
//
// Solidity: function aircraftIdToOwner(uint256 ) view returns(address)
func (_AircraftNFT *AircraftNFTCallerSession) AircraftIdToOwner(arg0 *big.Int) (common.Address, error) {
	return _AircraftNFT.Contract.AircraftIdToOwner(&_AircraftNFT.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_AircraftNFT *AircraftNFTCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_AircraftNFT *AircraftNFTSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _AircraftNFT.Contract.BalanceOf(&_AircraftNFT.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_AircraftNFT *AircraftNFTCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _AircraftNFT.Contract.BalanceOf(&_AircraftNFT.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_AircraftNFT *AircraftNFTCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_AircraftNFT *AircraftNFTSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _AircraftNFT.Contract.BalanceOfBatch(&_AircraftNFT.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_AircraftNFT *AircraftNFTCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _AircraftNFT.Contract.BalanceOfBatch(&_AircraftNFT.CallOpts, accounts, ids)
}

// GetAircrafts is a free data retrieval call binding the contract method 0xd77fea6e.
//
// Solidity: function getAircrafts(address account) view returns(uint256[])
func (_AircraftNFT *AircraftNFTCaller) GetAircrafts(opts *bind.CallOpts, account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "getAircrafts", account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAircrafts is a free data retrieval call binding the contract method 0xd77fea6e.
//
// Solidity: function getAircrafts(address account) view returns(uint256[])
func (_AircraftNFT *AircraftNFTSession) GetAircrafts(account common.Address) ([]*big.Int, error) {
	return _AircraftNFT.Contract.GetAircrafts(&_AircraftNFT.CallOpts, account)
}

// GetAircrafts is a free data retrieval call binding the contract method 0xd77fea6e.
//
// Solidity: function getAircrafts(address account) view returns(uint256[])
func (_AircraftNFT *AircraftNFTCallerSession) GetAircrafts(account common.Address) ([]*big.Int, error) {
	return _AircraftNFT.Contract.GetAircrafts(&_AircraftNFT.CallOpts, account)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AircraftNFT *AircraftNFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AircraftNFT *AircraftNFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AircraftNFT.Contract.GetRoleAdmin(&_AircraftNFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AircraftNFT *AircraftNFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AircraftNFT.Contract.GetRoleAdmin(&_AircraftNFT.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AircraftNFT *AircraftNFTCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AircraftNFT *AircraftNFTSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AircraftNFT.Contract.GetRoleMember(&_AircraftNFT.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AircraftNFT *AircraftNFTCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AircraftNFT.Contract.GetRoleMember(&_AircraftNFT.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AircraftNFT *AircraftNFTCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AircraftNFT *AircraftNFTSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AircraftNFT.Contract.GetRoleMemberCount(&_AircraftNFT.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AircraftNFT *AircraftNFTCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AircraftNFT.Contract.GetRoleMemberCount(&_AircraftNFT.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AircraftNFT *AircraftNFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AircraftNFT *AircraftNFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AircraftNFT.Contract.HasRole(&_AircraftNFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AircraftNFT *AircraftNFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AircraftNFT.Contract.HasRole(&_AircraftNFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_AircraftNFT *AircraftNFTCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_AircraftNFT *AircraftNFTSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _AircraftNFT.Contract.IsApprovedForAll(&_AircraftNFT.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_AircraftNFT *AircraftNFTCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _AircraftNFT.Contract.IsApprovedForAll(&_AircraftNFT.CallOpts, account, operator)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AircraftNFT *AircraftNFTCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AircraftNFT *AircraftNFTSession) Paused() (bool, error) {
	return _AircraftNFT.Contract.Paused(&_AircraftNFT.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AircraftNFT *AircraftNFTCallerSession) Paused() (bool, error) {
	return _AircraftNFT.Contract.Paused(&_AircraftNFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AircraftNFT *AircraftNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AircraftNFT *AircraftNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AircraftNFT.Contract.SupportsInterface(&_AircraftNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AircraftNFT *AircraftNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AircraftNFT.Contract.SupportsInterface(&_AircraftNFT.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_AircraftNFT *AircraftNFTCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _AircraftNFT.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_AircraftNFT *AircraftNFTSession) Uri(arg0 *big.Int) (string, error) {
	return _AircraftNFT.Contract.Uri(&_AircraftNFT.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_AircraftNFT *AircraftNFTCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _AircraftNFT.Contract.Uri(&_AircraftNFT.CallOpts, arg0)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_AircraftNFT *AircraftNFTTransactor) AddMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "addMinter", minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_AircraftNFT *AircraftNFTSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.AddMinter(&_AircraftNFT.TransactOpts, minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.AddMinter(&_AircraftNFT.TransactOpts, minter)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.GrantRole(&_AircraftNFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.GrantRole(&_AircraftNFT.TransactOpts, role, account)
}

// MintAircraft is a paid mutator transaction binding the contract method 0x2c4a38ed.
//
// Solidity: function mintAircraft(address to, uint256 amount) returns()
func (_AircraftNFT *AircraftNFTTransactor) MintAircraft(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "mintAircraft", to, amount)
}

// MintAircraft is a paid mutator transaction binding the contract method 0x2c4a38ed.
//
// Solidity: function mintAircraft(address to, uint256 amount) returns()
func (_AircraftNFT *AircraftNFTSession) MintAircraft(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.MintAircraft(&_AircraftNFT.TransactOpts, to, amount)
}

// MintAircraft is a paid mutator transaction binding the contract method 0x2c4a38ed.
//
// Solidity: function mintAircraft(address to, uint256 amount) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) MintAircraft(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.MintAircraft(&_AircraftNFT.TransactOpts, to, amount)
}

// MintComponent is a paid mutator transaction binding the contract method 0x43556cf0.
//
// Solidity: function mintComponent(address to, uint256 id) returns()
func (_AircraftNFT *AircraftNFTTransactor) MintComponent(opts *bind.TransactOpts, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "mintComponent", to, id)
}

// MintComponent is a paid mutator transaction binding the contract method 0x43556cf0.
//
// Solidity: function mintComponent(address to, uint256 id) returns()
func (_AircraftNFT *AircraftNFTSession) MintComponent(to common.Address, id *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.MintComponent(&_AircraftNFT.TransactOpts, to, id)
}

// MintComponent is a paid mutator transaction binding the contract method 0x43556cf0.
//
// Solidity: function mintComponent(address to, uint256 id) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) MintComponent(to common.Address, id *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.MintComponent(&_AircraftNFT.TransactOpts, to, id)
}

// MintLuckyBox is a paid mutator transaction binding the contract method 0xa3dd3720.
//
// Solidity: function mintLuckyBox(address to, uint256 amount) returns()
func (_AircraftNFT *AircraftNFTTransactor) MintLuckyBox(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "mintLuckyBox", to, amount)
}

// MintLuckyBox is a paid mutator transaction binding the contract method 0xa3dd3720.
//
// Solidity: function mintLuckyBox(address to, uint256 amount) returns()
func (_AircraftNFT *AircraftNFTSession) MintLuckyBox(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.MintLuckyBox(&_AircraftNFT.TransactOpts, to, amount)
}

// MintLuckyBox is a paid mutator transaction binding the contract method 0xa3dd3720.
//
// Solidity: function mintLuckyBox(address to, uint256 amount) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) MintLuckyBox(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.MintLuckyBox(&_AircraftNFT.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AircraftNFT *AircraftNFTTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AircraftNFT *AircraftNFTSession) Pause() (*types.Transaction, error) {
	return _AircraftNFT.Contract.Pause(&_AircraftNFT.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AircraftNFT *AircraftNFTTransactorSession) Pause() (*types.Transaction, error) {
	return _AircraftNFT.Contract.Pause(&_AircraftNFT.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.RenounceRole(&_AircraftNFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.RenounceRole(&_AircraftNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.RevokeRole(&_AircraftNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AircraftNFT.Contract.RevokeRole(&_AircraftNFT.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_AircraftNFT *AircraftNFTTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_AircraftNFT *AircraftNFTSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _AircraftNFT.Contract.SafeBatchTransferFrom(&_AircraftNFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _AircraftNFT.Contract.SafeBatchTransferFrom(&_AircraftNFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_AircraftNFT *AircraftNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_AircraftNFT *AircraftNFTSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AircraftNFT.Contract.SafeTransferFrom(&_AircraftNFT.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AircraftNFT.Contract.SafeTransferFrom(&_AircraftNFT.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AircraftNFT *AircraftNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AircraftNFT *AircraftNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AircraftNFT.Contract.SetApprovalForAll(&_AircraftNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AircraftNFT.Contract.SetApprovalForAll(&_AircraftNFT.TransactOpts, operator, approved)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AircraftNFT *AircraftNFTTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AircraftNFT *AircraftNFTSession) Unpause() (*types.Transaction, error) {
	return _AircraftNFT.Contract.Unpause(&_AircraftNFT.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AircraftNFT *AircraftNFTTransactorSession) Unpause() (*types.Transaction, error) {
	return _AircraftNFT.Contract.Unpause(&_AircraftNFT.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x50fa8942.
//
// Solidity: function upgrade(address to, uint256 aircraftId, uint256 componentId) returns()
func (_AircraftNFT *AircraftNFTTransactor) Upgrade(opts *bind.TransactOpts, to common.Address, aircraftId *big.Int, componentId *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.contract.Transact(opts, "upgrade", to, aircraftId, componentId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x50fa8942.
//
// Solidity: function upgrade(address to, uint256 aircraftId, uint256 componentId) returns()
func (_AircraftNFT *AircraftNFTSession) Upgrade(to common.Address, aircraftId *big.Int, componentId *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.Upgrade(&_AircraftNFT.TransactOpts, to, aircraftId, componentId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x50fa8942.
//
// Solidity: function upgrade(address to, uint256 aircraftId, uint256 componentId) returns()
func (_AircraftNFT *AircraftNFTTransactorSession) Upgrade(to common.Address, aircraftId *big.Int, componentId *big.Int) (*types.Transaction, error) {
	return _AircraftNFT.Contract.Upgrade(&_AircraftNFT.TransactOpts, to, aircraftId, componentId)
}

// AircraftNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the AircraftNFT contract.
type AircraftNFTApprovalForAllIterator struct {
	Event *AircraftNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AircraftNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTApprovalForAll)
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
		it.Event = new(AircraftNFTApprovalForAll)
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
func (it *AircraftNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTApprovalForAll represents a ApprovalForAll event raised by the AircraftNFT contract.
type AircraftNFTApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_AircraftNFT *AircraftNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*AircraftNFTApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTApprovalForAllIterator{contract: _AircraftNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_AircraftNFT *AircraftNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AircraftNFTApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTApprovalForAll)
				if err := _AircraftNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_AircraftNFT *AircraftNFTFilterer) ParseApprovalForAll(log types.Log) (*AircraftNFTApprovalForAll, error) {
	event := new(AircraftNFTApprovalForAll)
	if err := _AircraftNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AircraftNFT contract.
type AircraftNFTPausedIterator struct {
	Event *AircraftNFTPaused // Event containing the contract specifics and raw log

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
func (it *AircraftNFTPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTPaused)
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
		it.Event = new(AircraftNFTPaused)
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
func (it *AircraftNFTPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTPaused represents a Paused event raised by the AircraftNFT contract.
type AircraftNFTPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AircraftNFT *AircraftNFTFilterer) FilterPaused(opts *bind.FilterOpts) (*AircraftNFTPausedIterator, error) {

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AircraftNFTPausedIterator{contract: _AircraftNFT.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AircraftNFT *AircraftNFTFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AircraftNFTPaused) (event.Subscription, error) {

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTPaused)
				if err := _AircraftNFT.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AircraftNFT *AircraftNFTFilterer) ParsePaused(log types.Log) (*AircraftNFTPaused, error) {
	event := new(AircraftNFTPaused)
	if err := _AircraftNFT.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AircraftNFT contract.
type AircraftNFTRoleAdminChangedIterator struct {
	Event *AircraftNFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AircraftNFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTRoleAdminChanged)
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
		it.Event = new(AircraftNFTRoleAdminChanged)
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
func (it *AircraftNFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTRoleAdminChanged represents a RoleAdminChanged event raised by the AircraftNFT contract.
type AircraftNFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AircraftNFT *AircraftNFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AircraftNFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTRoleAdminChangedIterator{contract: _AircraftNFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AircraftNFT *AircraftNFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AircraftNFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTRoleAdminChanged)
				if err := _AircraftNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AircraftNFT *AircraftNFTFilterer) ParseRoleAdminChanged(log types.Log) (*AircraftNFTRoleAdminChanged, error) {
	event := new(AircraftNFTRoleAdminChanged)
	if err := _AircraftNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AircraftNFT contract.
type AircraftNFTRoleGrantedIterator struct {
	Event *AircraftNFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *AircraftNFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTRoleGranted)
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
		it.Event = new(AircraftNFTRoleGranted)
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
func (it *AircraftNFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTRoleGranted represents a RoleGranted event raised by the AircraftNFT contract.
type AircraftNFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AircraftNFT *AircraftNFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AircraftNFTRoleGrantedIterator, error) {

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

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTRoleGrantedIterator{contract: _AircraftNFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AircraftNFT *AircraftNFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AircraftNFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTRoleGranted)
				if err := _AircraftNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AircraftNFT *AircraftNFTFilterer) ParseRoleGranted(log types.Log) (*AircraftNFTRoleGranted, error) {
	event := new(AircraftNFTRoleGranted)
	if err := _AircraftNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AircraftNFT contract.
type AircraftNFTRoleRevokedIterator struct {
	Event *AircraftNFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AircraftNFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTRoleRevoked)
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
		it.Event = new(AircraftNFTRoleRevoked)
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
func (it *AircraftNFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTRoleRevoked represents a RoleRevoked event raised by the AircraftNFT contract.
type AircraftNFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AircraftNFT *AircraftNFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AircraftNFTRoleRevokedIterator, error) {

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

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTRoleRevokedIterator{contract: _AircraftNFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AircraftNFT *AircraftNFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AircraftNFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTRoleRevoked)
				if err := _AircraftNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AircraftNFT *AircraftNFTFilterer) ParseRoleRevoked(log types.Log) (*AircraftNFTRoleRevoked, error) {
	event := new(AircraftNFTRoleRevoked)
	if err := _AircraftNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the AircraftNFT contract.
type AircraftNFTTransferBatchIterator struct {
	Event *AircraftNFTTransferBatch // Event containing the contract specifics and raw log

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
func (it *AircraftNFTTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTTransferBatch)
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
		it.Event = new(AircraftNFTTransferBatch)
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
func (it *AircraftNFTTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTTransferBatch represents a TransferBatch event raised by the AircraftNFT contract.
type AircraftNFTTransferBatch struct {
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
func (_AircraftNFT *AircraftNFTFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*AircraftNFTTransferBatchIterator, error) {

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

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTTransferBatchIterator{contract: _AircraftNFT.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_AircraftNFT *AircraftNFTFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *AircraftNFTTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTTransferBatch)
				if err := _AircraftNFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_AircraftNFT *AircraftNFTFilterer) ParseTransferBatch(log types.Log) (*AircraftNFTTransferBatch, error) {
	event := new(AircraftNFTTransferBatch)
	if err := _AircraftNFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the AircraftNFT contract.
type AircraftNFTTransferSingleIterator struct {
	Event *AircraftNFTTransferSingle // Event containing the contract specifics and raw log

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
func (it *AircraftNFTTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTTransferSingle)
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
		it.Event = new(AircraftNFTTransferSingle)
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
func (it *AircraftNFTTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTTransferSingle represents a TransferSingle event raised by the AircraftNFT contract.
type AircraftNFTTransferSingle struct {
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
func (_AircraftNFT *AircraftNFTFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*AircraftNFTTransferSingleIterator, error) {

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

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTTransferSingleIterator{contract: _AircraftNFT.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_AircraftNFT *AircraftNFTFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *AircraftNFTTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTTransferSingle)
				if err := _AircraftNFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_AircraftNFT *AircraftNFTFilterer) ParseTransferSingle(log types.Log) (*AircraftNFTTransferSingle, error) {
	event := new(AircraftNFTTransferSingle)
	if err := _AircraftNFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the AircraftNFT contract.
type AircraftNFTURIIterator struct {
	Event *AircraftNFTURI // Event containing the contract specifics and raw log

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
func (it *AircraftNFTURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTURI)
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
		it.Event = new(AircraftNFTURI)
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
func (it *AircraftNFTURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTURI represents a URI event raised by the AircraftNFT contract.
type AircraftNFTURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AircraftNFT *AircraftNFTFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*AircraftNFTURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTURIIterator{contract: _AircraftNFT.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_AircraftNFT *AircraftNFTFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *AircraftNFTURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTURI)
				if err := _AircraftNFT.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_AircraftNFT *AircraftNFTFilterer) ParseURI(log types.Log) (*AircraftNFTURI, error) {
	event := new(AircraftNFTURI)
	if err := _AircraftNFT.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AircraftNFT contract.
type AircraftNFTUnpausedIterator struct {
	Event *AircraftNFTUnpaused // Event containing the contract specifics and raw log

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
func (it *AircraftNFTUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTUnpaused)
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
		it.Event = new(AircraftNFTUnpaused)
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
func (it *AircraftNFTUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTUnpaused represents a Unpaused event raised by the AircraftNFT contract.
type AircraftNFTUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AircraftNFT *AircraftNFTFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AircraftNFTUnpausedIterator, error) {

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AircraftNFTUnpausedIterator{contract: _AircraftNFT.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AircraftNFT *AircraftNFTFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AircraftNFTUnpaused) (event.Subscription, error) {

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTUnpaused)
				if err := _AircraftNFT.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AircraftNFT *AircraftNFTFilterer) ParseUnpaused(log types.Log) (*AircraftNFTUnpaused, error) {
	event := new(AircraftNFTUnpaused)
	if err := _AircraftNFT.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTUpgradeAircraftIterator is returned from FilterUpgradeAircraft and is used to iterate over the raw logs and unpacked data for UpgradeAircraft events raised by the AircraftNFT contract.
type AircraftNFTUpgradeAircraftIterator struct {
	Event *AircraftNFTUpgradeAircraft // Event containing the contract specifics and raw log

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
func (it *AircraftNFTUpgradeAircraftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTUpgradeAircraft)
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
		it.Event = new(AircraftNFTUpgradeAircraft)
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
func (it *AircraftNFTUpgradeAircraftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTUpgradeAircraftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTUpgradeAircraft represents a UpgradeAircraft event raised by the AircraftNFT contract.
type AircraftNFTUpgradeAircraft struct {
	Operator    common.Address
	AircraftId  *big.Int
	ComponentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAircraft is a free log retrieval operation binding the contract event 0x386f8ebec58297b76dfa5d010bc797af0edeb58610b9d3a42160ff7592d7b933.
//
// Solidity: event UpgradeAircraft(address indexed operator, uint256 indexed aircraftId, uint256 indexed componentId)
func (_AircraftNFT *AircraftNFTFilterer) FilterUpgradeAircraft(opts *bind.FilterOpts, operator []common.Address, aircraftId []*big.Int, componentId []*big.Int) (*AircraftNFTUpgradeAircraftIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var aircraftIdRule []interface{}
	for _, aircraftIdItem := range aircraftId {
		aircraftIdRule = append(aircraftIdRule, aircraftIdItem)
	}
	var componentIdRule []interface{}
	for _, componentIdItem := range componentId {
		componentIdRule = append(componentIdRule, componentIdItem)
	}

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "UpgradeAircraft", operatorRule, aircraftIdRule, componentIdRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTUpgradeAircraftIterator{contract: _AircraftNFT.contract, event: "UpgradeAircraft", logs: logs, sub: sub}, nil
}

// WatchUpgradeAircraft is a free log subscription operation binding the contract event 0x386f8ebec58297b76dfa5d010bc797af0edeb58610b9d3a42160ff7592d7b933.
//
// Solidity: event UpgradeAircraft(address indexed operator, uint256 indexed aircraftId, uint256 indexed componentId)
func (_AircraftNFT *AircraftNFTFilterer) WatchUpgradeAircraft(opts *bind.WatchOpts, sink chan<- *AircraftNFTUpgradeAircraft, operator []common.Address, aircraftId []*big.Int, componentId []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var aircraftIdRule []interface{}
	for _, aircraftIdItem := range aircraftId {
		aircraftIdRule = append(aircraftIdRule, aircraftIdItem)
	}
	var componentIdRule []interface{}
	for _, componentIdItem := range componentId {
		componentIdRule = append(componentIdRule, componentIdItem)
	}

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "UpgradeAircraft", operatorRule, aircraftIdRule, componentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTUpgradeAircraft)
				if err := _AircraftNFT.contract.UnpackLog(event, "UpgradeAircraft", log); err != nil {
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

// ParseUpgradeAircraft is a log parse operation binding the contract event 0x386f8ebec58297b76dfa5d010bc797af0edeb58610b9d3a42160ff7592d7b933.
//
// Solidity: event UpgradeAircraft(address indexed operator, uint256 indexed aircraftId, uint256 indexed componentId)
func (_AircraftNFT *AircraftNFTFilterer) ParseUpgradeAircraft(log types.Log) (*AircraftNFTUpgradeAircraft, error) {
	event := new(AircraftNFTUpgradeAircraft)
	if err := _AircraftNFT.contract.UnpackLog(event, "UpgradeAircraft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AircraftNFTEventMintedAircraftIterator is returned from FilterEventMintedAircraft and is used to iterate over the raw logs and unpacked data for EventMintedAircraft events raised by the AircraftNFT contract.
type AircraftNFTEventMintedAircraftIterator struct {
	Event *AircraftNFTEventMintedAircraft // Event containing the contract specifics and raw log

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
func (it *AircraftNFTEventMintedAircraftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AircraftNFTEventMintedAircraft)
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
		it.Event = new(AircraftNFTEventMintedAircraft)
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
func (it *AircraftNFTEventMintedAircraftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AircraftNFTEventMintedAircraftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AircraftNFTEventMintedAircraft represents a EventMintedAircraft event raised by the AircraftNFT contract.
type AircraftNFTEventMintedAircraft struct {
	Operator   common.Address
	AircraftId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventMintedAircraft is a free log retrieval operation binding the contract event 0xf2b0a49cd859710f9fe79ae979f330607d532f61b1cd3942d11ca91dbda52af2.
//
// Solidity: event event_MintedAircraft(address indexed operator, uint256 indexed aircraftId)
func (_AircraftNFT *AircraftNFTFilterer) FilterEventMintedAircraft(opts *bind.FilterOpts, operator []common.Address, aircraftId []*big.Int) (*AircraftNFTEventMintedAircraftIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var aircraftIdRule []interface{}
	for _, aircraftIdItem := range aircraftId {
		aircraftIdRule = append(aircraftIdRule, aircraftIdItem)
	}

	logs, sub, err := _AircraftNFT.contract.FilterLogs(opts, "event_MintedAircraft", operatorRule, aircraftIdRule)
	if err != nil {
		return nil, err
	}
	return &AircraftNFTEventMintedAircraftIterator{contract: _AircraftNFT.contract, event: "event_MintedAircraft", logs: logs, sub: sub}, nil
}

// WatchEventMintedAircraft is a free log subscription operation binding the contract event 0xf2b0a49cd859710f9fe79ae979f330607d532f61b1cd3942d11ca91dbda52af2.
//
// Solidity: event event_MintedAircraft(address indexed operator, uint256 indexed aircraftId)
func (_AircraftNFT *AircraftNFTFilterer) WatchEventMintedAircraft(opts *bind.WatchOpts, sink chan<- *AircraftNFTEventMintedAircraft, operator []common.Address, aircraftId []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var aircraftIdRule []interface{}
	for _, aircraftIdItem := range aircraftId {
		aircraftIdRule = append(aircraftIdRule, aircraftIdItem)
	}

	logs, sub, err := _AircraftNFT.contract.WatchLogs(opts, "event_MintedAircraft", operatorRule, aircraftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AircraftNFTEventMintedAircraft)
				if err := _AircraftNFT.contract.UnpackLog(event, "event_MintedAircraft", log); err != nil {
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

// ParseEventMintedAircraft is a log parse operation binding the contract event 0xf2b0a49cd859710f9fe79ae979f330607d532f61b1cd3942d11ca91dbda52af2.
//
// Solidity: event event_MintedAircraft(address indexed operator, uint256 indexed aircraftId)
func (_AircraftNFT *AircraftNFTFilterer) ParseEventMintedAircraft(log types.Log) (*AircraftNFTEventMintedAircraft, error) {
	event := new(AircraftNFTEventMintedAircraft)
	if err := _AircraftNFT.contract.UnpackLog(event, "event_MintedAircraft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
