// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// PancakePredictionV2BetInfo is an auto generated low-level Go binding around an user-defined struct.
type PancakePredictionV2BetInfo struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}

// PancakePredictionMetaData contains all meta data concerning the PancakePrediction contract.
var PancakePredictionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_intervalSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bufferSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oracleUpdateAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BetBear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BetBull\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"EndRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"LockRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdminAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bufferSeconds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"intervalSeconds\",\"type\":\"uint256\"}],\"name\":\"NewBufferAndIntervalSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBetAmount\",\"type\":\"uint256\"}],\"name\":\"NewMinBetAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"NewOperatorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oracleUpdateAllowance\",\"type\":\"uint256\"}],\"name\":\"NewOracleUpdateAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardBaseCalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryAmount\",\"type\":\"uint256\"}],\"name\":\"RewardsCalculated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"StartRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TreasuryClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TREASURY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"betBear\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"betBull\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bufferSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisLockOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisLockRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisStartOnce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisStartRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cursor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getUserRounds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"enumPancakePredictionV2.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"internalType\":\"structPancakePredictionV2.BetInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserRoundsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intervalSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"enumPancakePredictionV2.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleLatestRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleUpdateAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"refundable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"lockPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"closePrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lockOracleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closeOracleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bullAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bearAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardBaseCalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"oracleCalled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bufferSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_intervalSeconds\",\"type\":\"uint256\"}],\"name\":\"setBufferAndIntervalSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBetAmount\",\"type\":\"uint256\"}],\"name\":\"setMinBetAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorAddress\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_oracleUpdateAllowance\",\"type\":\"uint256\"}],\"name\":\"setOracleUpdateAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"}],\"name\":\"setTreasuryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userRounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PancakePredictionABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakePredictionMetaData.ABI instead.
var PancakePredictionABI = PancakePredictionMetaData.ABI

// PancakePrediction is an auto generated Go binding around an Ethereum contract.
type PancakePrediction struct {
	PancakePredictionCaller     // Read-only binding to the contract
	PancakePredictionTransactor // Write-only binding to the contract
	PancakePredictionFilterer   // Log filterer for contract events
}

// PancakePredictionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakePredictionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakePredictionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakePredictionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakePredictionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakePredictionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakePredictionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakePredictionSession struct {
	Contract     *PancakePrediction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PancakePredictionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakePredictionCallerSession struct {
	Contract *PancakePredictionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PancakePredictionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakePredictionTransactorSession struct {
	Contract     *PancakePredictionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PancakePredictionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakePredictionRaw struct {
	Contract *PancakePrediction // Generic contract binding to access the raw methods on
}

// PancakePredictionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakePredictionCallerRaw struct {
	Contract *PancakePredictionCaller // Generic read-only contract binding to access the raw methods on
}

// PancakePredictionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakePredictionTransactorRaw struct {
	Contract *PancakePredictionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakePrediction creates a new instance of PancakePrediction, bound to a specific deployed contract.
func NewPancakePrediction(address common.Address, backend bind.ContractBackend) (*PancakePrediction, error) {
	contract, err := bindPancakePrediction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakePrediction{PancakePredictionCaller: PancakePredictionCaller{contract: contract}, PancakePredictionTransactor: PancakePredictionTransactor{contract: contract}, PancakePredictionFilterer: PancakePredictionFilterer{contract: contract}}, nil
}

// NewPancakePredictionCaller creates a new read-only instance of PancakePrediction, bound to a specific deployed contract.
func NewPancakePredictionCaller(address common.Address, caller bind.ContractCaller) (*PancakePredictionCaller, error) {
	contract, err := bindPancakePrediction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionCaller{contract: contract}, nil
}

// NewPancakePredictionTransactor creates a new write-only instance of PancakePrediction, bound to a specific deployed contract.
func NewPancakePredictionTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakePredictionTransactor, error) {
	contract, err := bindPancakePrediction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionTransactor{contract: contract}, nil
}

// NewPancakePredictionFilterer creates a new log filterer instance of PancakePrediction, bound to a specific deployed contract.
func NewPancakePredictionFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakePredictionFilterer, error) {
	contract, err := bindPancakePrediction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionFilterer{contract: contract}, nil
}

// bindPancakePrediction binds a generic wrapper to an already deployed contract.
func bindPancakePrediction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakePredictionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakePrediction *PancakePredictionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakePrediction.Contract.PancakePredictionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakePrediction *PancakePredictionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.Contract.PancakePredictionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakePrediction *PancakePredictionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakePrediction.Contract.PancakePredictionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakePrediction *PancakePredictionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakePrediction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakePrediction *PancakePredictionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakePrediction *PancakePredictionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakePrediction.Contract.contract.Transact(opts, method, params...)
}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) MAXTREASURYFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "MAX_TREASURY_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) MAXTREASURYFEE() (*big.Int, error) {
	return _PancakePrediction.Contract.MAXTREASURYFEE(&_PancakePrediction.CallOpts)
}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) MAXTREASURYFEE() (*big.Int, error) {
	return _PancakePrediction.Contract.MAXTREASURYFEE(&_PancakePrediction.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_PancakePrediction *PancakePredictionCaller) AdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "adminAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_PancakePrediction *PancakePredictionSession) AdminAddress() (common.Address, error) {
	return _PancakePrediction.Contract.AdminAddress(&_PancakePrediction.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_PancakePrediction *PancakePredictionCallerSession) AdminAddress() (common.Address, error) {
	return _PancakePrediction.Contract.AdminAddress(&_PancakePrediction.CallOpts)
}

// BufferSeconds is a free data retrieval call binding the contract method 0xeaba2361.
//
// Solidity: function bufferSeconds() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) BufferSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "bufferSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferSeconds is a free data retrieval call binding the contract method 0xeaba2361.
//
// Solidity: function bufferSeconds() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) BufferSeconds() (*big.Int, error) {
	return _PancakePrediction.Contract.BufferSeconds(&_PancakePrediction.CallOpts)
}

// BufferSeconds is a free data retrieval call binding the contract method 0xeaba2361.
//
// Solidity: function bufferSeconds() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) BufferSeconds() (*big.Int, error) {
	return _PancakePrediction.Contract.BufferSeconds(&_PancakePrediction.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_PancakePrediction *PancakePredictionCaller) Claimable(opts *bind.CallOpts, epoch *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "claimable", epoch, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_PancakePrediction *PancakePredictionSession) Claimable(epoch *big.Int, user common.Address) (bool, error) {
	return _PancakePrediction.Contract.Claimable(&_PancakePrediction.CallOpts, epoch, user)
}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_PancakePrediction *PancakePredictionCallerSession) Claimable(epoch *big.Int, user common.Address) (bool, error) {
	return _PancakePrediction.Contract.Claimable(&_PancakePrediction.CallOpts, epoch, user)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) CurrentEpoch() (*big.Int, error) {
	return _PancakePrediction.Contract.CurrentEpoch(&_PancakePrediction.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) CurrentEpoch() (*big.Int, error) {
	return _PancakePrediction.Contract.CurrentEpoch(&_PancakePrediction.CallOpts)
}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_PancakePrediction *PancakePredictionCaller) GenesisLockOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "genesisLockOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_PancakePrediction *PancakePredictionSession) GenesisLockOnce() (bool, error) {
	return _PancakePrediction.Contract.GenesisLockOnce(&_PancakePrediction.CallOpts)
}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_PancakePrediction *PancakePredictionCallerSession) GenesisLockOnce() (bool, error) {
	return _PancakePrediction.Contract.GenesisLockOnce(&_PancakePrediction.CallOpts)
}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_PancakePrediction *PancakePredictionCaller) GenesisStartOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "genesisStartOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_PancakePrediction *PancakePredictionSession) GenesisStartOnce() (bool, error) {
	return _PancakePrediction.Contract.GenesisStartOnce(&_PancakePrediction.CallOpts)
}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_PancakePrediction *PancakePredictionCallerSession) GenesisStartOnce() (bool, error) {
	return _PancakePrediction.Contract.GenesisStartOnce(&_PancakePrediction.CallOpts)
}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], (uint8,uint256,bool)[], uint256)
func (_PancakePrediction *PancakePredictionCaller) GetUserRounds(opts *bind.CallOpts, user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, []PancakePredictionV2BetInfo, *big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "getUserRounds", user, cursor, size)

	if err != nil {
		return *new([]*big.Int), *new([]PancakePredictionV2BetInfo), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]PancakePredictionV2BetInfo)).(*[]PancakePredictionV2BetInfo)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], (uint8,uint256,bool)[], uint256)
func (_PancakePrediction *PancakePredictionSession) GetUserRounds(user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, []PancakePredictionV2BetInfo, *big.Int, error) {
	return _PancakePrediction.Contract.GetUserRounds(&_PancakePrediction.CallOpts, user, cursor, size)
}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], (uint8,uint256,bool)[], uint256)
func (_PancakePrediction *PancakePredictionCallerSession) GetUserRounds(user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, []PancakePredictionV2BetInfo, *big.Int, error) {
	return _PancakePrediction.Contract.GetUserRounds(&_PancakePrediction.CallOpts, user, cursor, size)
}

// GetUserRoundsLength is a free data retrieval call binding the contract method 0x273867d4.
//
// Solidity: function getUserRoundsLength(address user) view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) GetUserRoundsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "getUserRoundsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserRoundsLength is a free data retrieval call binding the contract method 0x273867d4.
//
// Solidity: function getUserRoundsLength(address user) view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) GetUserRoundsLength(user common.Address) (*big.Int, error) {
	return _PancakePrediction.Contract.GetUserRoundsLength(&_PancakePrediction.CallOpts, user)
}

// GetUserRoundsLength is a free data retrieval call binding the contract method 0x273867d4.
//
// Solidity: function getUserRoundsLength(address user) view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) GetUserRoundsLength(user common.Address) (*big.Int, error) {
	return _PancakePrediction.Contract.GetUserRoundsLength(&_PancakePrediction.CallOpts, user)
}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) IntervalSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "intervalSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) IntervalSeconds() (*big.Int, error) {
	return _PancakePrediction.Contract.IntervalSeconds(&_PancakePrediction.CallOpts)
}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) IntervalSeconds() (*big.Int, error) {
	return _PancakePrediction.Contract.IntervalSeconds(&_PancakePrediction.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_PancakePrediction *PancakePredictionCaller) Ledger(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "ledger", arg0, arg1)

	outstruct := new(struct {
		Position uint8
		Amount   *big.Int
		Claimed  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Position = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_PancakePrediction *PancakePredictionSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	return _PancakePrediction.Contract.Ledger(&_PancakePrediction.CallOpts, arg0, arg1)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_PancakePrediction *PancakePredictionCallerSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	return _PancakePrediction.Contract.Ledger(&_PancakePrediction.CallOpts, arg0, arg1)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) MinBetAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "minBetAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) MinBetAmount() (*big.Int, error) {
	return _PancakePrediction.Contract.MinBetAmount(&_PancakePrediction.CallOpts)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) MinBetAmount() (*big.Int, error) {
	return _PancakePrediction.Contract.MinBetAmount(&_PancakePrediction.CallOpts)
}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_PancakePrediction *PancakePredictionCaller) OperatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "operatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_PancakePrediction *PancakePredictionSession) OperatorAddress() (common.Address, error) {
	return _PancakePrediction.Contract.OperatorAddress(&_PancakePrediction.CallOpts)
}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_PancakePrediction *PancakePredictionCallerSession) OperatorAddress() (common.Address, error) {
	return _PancakePrediction.Contract.OperatorAddress(&_PancakePrediction.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PancakePrediction *PancakePredictionCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PancakePrediction *PancakePredictionSession) Oracle() (common.Address, error) {
	return _PancakePrediction.Contract.Oracle(&_PancakePrediction.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_PancakePrediction *PancakePredictionCallerSession) Oracle() (common.Address, error) {
	return _PancakePrediction.Contract.Oracle(&_PancakePrediction.CallOpts)
}

// OracleLatestRoundId is a free data retrieval call binding the contract method 0xec324703.
//
// Solidity: function oracleLatestRoundId() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) OracleLatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "oracleLatestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleLatestRoundId is a free data retrieval call binding the contract method 0xec324703.
//
// Solidity: function oracleLatestRoundId() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) OracleLatestRoundId() (*big.Int, error) {
	return _PancakePrediction.Contract.OracleLatestRoundId(&_PancakePrediction.CallOpts)
}

// OracleLatestRoundId is a free data retrieval call binding the contract method 0xec324703.
//
// Solidity: function oracleLatestRoundId() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) OracleLatestRoundId() (*big.Int, error) {
	return _PancakePrediction.Contract.OracleLatestRoundId(&_PancakePrediction.CallOpts)
}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) OracleUpdateAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "oracleUpdateAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) OracleUpdateAllowance() (*big.Int, error) {
	return _PancakePrediction.Contract.OracleUpdateAllowance(&_PancakePrediction.CallOpts)
}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) OracleUpdateAllowance() (*big.Int, error) {
	return _PancakePrediction.Contract.OracleUpdateAllowance(&_PancakePrediction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakePrediction *PancakePredictionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakePrediction *PancakePredictionSession) Owner() (common.Address, error) {
	return _PancakePrediction.Contract.Owner(&_PancakePrediction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakePrediction *PancakePredictionCallerSession) Owner() (common.Address, error) {
	return _PancakePrediction.Contract.Owner(&_PancakePrediction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PancakePrediction *PancakePredictionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PancakePrediction *PancakePredictionSession) Paused() (bool, error) {
	return _PancakePrediction.Contract.Paused(&_PancakePrediction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PancakePrediction *PancakePredictionCallerSession) Paused() (bool, error) {
	return _PancakePrediction.Contract.Paused(&_PancakePrediction.CallOpts)
}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_PancakePrediction *PancakePredictionCaller) Refundable(opts *bind.CallOpts, epoch *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "refundable", epoch, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_PancakePrediction *PancakePredictionSession) Refundable(epoch *big.Int, user common.Address) (bool, error) {
	return _PancakePrediction.Contract.Refundable(&_PancakePrediction.CallOpts, epoch, user)
}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_PancakePrediction *PancakePredictionCallerSession) Refundable(epoch *big.Int, user common.Address) (bool, error) {
	return _PancakePrediction.Contract.Refundable(&_PancakePrediction.CallOpts, epoch, user)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 lockOracleId, uint256 closeOracleId, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_PancakePrediction *PancakePredictionCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartTimestamp      *big.Int
	LockTimestamp       *big.Int
	CloseTimestamp      *big.Int
	LockPrice           *big.Int
	ClosePrice          *big.Int
	LockOracleId        *big.Int
	CloseOracleId       *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "rounds", arg0)

	outstruct := new(struct {
		Epoch               *big.Int
		StartTimestamp      *big.Int
		LockTimestamp       *big.Int
		CloseTimestamp      *big.Int
		LockPrice           *big.Int
		ClosePrice          *big.Int
		LockOracleId        *big.Int
		CloseOracleId       *big.Int
		TotalAmount         *big.Int
		BullAmount          *big.Int
		BearAmount          *big.Int
		RewardBaseCalAmount *big.Int
		RewardAmount        *big.Int
		OracleCalled        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CloseTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ClosePrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LockOracleId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CloseOracleId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.BullAmount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.BearAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.RewardBaseCalAmount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.RewardAmount = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.OracleCalled = *abi.ConvertType(out[13], new(bool)).(*bool)

	return *outstruct, err

}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 lockOracleId, uint256 closeOracleId, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_PancakePrediction *PancakePredictionSession) Rounds(arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartTimestamp      *big.Int
	LockTimestamp       *big.Int
	CloseTimestamp      *big.Int
	LockPrice           *big.Int
	ClosePrice          *big.Int
	LockOracleId        *big.Int
	CloseOracleId       *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	return _PancakePrediction.Contract.Rounds(&_PancakePrediction.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 lockOracleId, uint256 closeOracleId, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_PancakePrediction *PancakePredictionCallerSession) Rounds(arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartTimestamp      *big.Int
	LockTimestamp       *big.Int
	CloseTimestamp      *big.Int
	LockPrice           *big.Int
	ClosePrice          *big.Int
	LockOracleId        *big.Int
	CloseOracleId       *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	return _PancakePrediction.Contract.Rounds(&_PancakePrediction.CallOpts, arg0)
}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) TreasuryAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "treasuryAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) TreasuryAmount() (*big.Int, error) {
	return _PancakePrediction.Contract.TreasuryAmount(&_PancakePrediction.CallOpts)
}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) TreasuryAmount() (*big.Int, error) {
	return _PancakePrediction.Contract.TreasuryAmount(&_PancakePrediction.CallOpts)
}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) TreasuryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "treasuryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) TreasuryFee() (*big.Int, error) {
	return _PancakePrediction.Contract.TreasuryFee(&_PancakePrediction.CallOpts)
}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) TreasuryFee() (*big.Int, error) {
	return _PancakePrediction.Contract.TreasuryFee(&_PancakePrediction.CallOpts)
}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_PancakePrediction *PancakePredictionCaller) UserRounds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakePrediction.contract.Call(opts, &out, "userRounds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_PancakePrediction *PancakePredictionSession) UserRounds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PancakePrediction.Contract.UserRounds(&_PancakePrediction.CallOpts, arg0, arg1)
}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_PancakePrediction *PancakePredictionCallerSession) UserRounds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PancakePrediction.Contract.UserRounds(&_PancakePrediction.CallOpts, arg0, arg1)
}

// BetBear is a paid mutator transaction binding the contract method 0xaa6b873a.
//
// Solidity: function betBear(uint256 epoch) payable returns()
func (_PancakePrediction *PancakePredictionTransactor) BetBear(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "betBear", epoch)
}

// BetBear is a paid mutator transaction binding the contract method 0xaa6b873a.
//
// Solidity: function betBear(uint256 epoch) payable returns()
func (_PancakePrediction *PancakePredictionSession) BetBear(epoch *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.BetBear(&_PancakePrediction.TransactOpts, epoch)
}

// BetBear is a paid mutator transaction binding the contract method 0xaa6b873a.
//
// Solidity: function betBear(uint256 epoch) payable returns()
func (_PancakePrediction *PancakePredictionTransactorSession) BetBear(epoch *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.BetBear(&_PancakePrediction.TransactOpts, epoch)
}

// BetBull is a paid mutator transaction binding the contract method 0x57fb096f.
//
// Solidity: function betBull(uint256 epoch) payable returns()
func (_PancakePrediction *PancakePredictionTransactor) BetBull(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "betBull", epoch)
}

// BetBull is a paid mutator transaction binding the contract method 0x57fb096f.
//
// Solidity: function betBull(uint256 epoch) payable returns()
func (_PancakePrediction *PancakePredictionSession) BetBull(epoch *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.BetBull(&_PancakePrediction.TransactOpts, epoch)
}

// BetBull is a paid mutator transaction binding the contract method 0x57fb096f.
//
// Solidity: function betBull(uint256 epoch) payable returns()
func (_PancakePrediction *PancakePredictionTransactorSession) BetBull(epoch *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.BetBull(&_PancakePrediction.TransactOpts, epoch)
}

// Claim is a paid mutator transaction binding the contract method 0x6ba4c138.
//
// Solidity: function claim(uint256[] epochs) returns()
func (_PancakePrediction *PancakePredictionTransactor) Claim(opts *bind.TransactOpts, epochs []*big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "claim", epochs)
}

// Claim is a paid mutator transaction binding the contract method 0x6ba4c138.
//
// Solidity: function claim(uint256[] epochs) returns()
func (_PancakePrediction *PancakePredictionSession) Claim(epochs []*big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.Claim(&_PancakePrediction.TransactOpts, epochs)
}

// Claim is a paid mutator transaction binding the contract method 0x6ba4c138.
//
// Solidity: function claim(uint256[] epochs) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) Claim(epochs []*big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.Claim(&_PancakePrediction.TransactOpts, epochs)
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_PancakePrediction *PancakePredictionTransactor) ClaimTreasury(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "claimTreasury")
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_PancakePrediction *PancakePredictionSession) ClaimTreasury() (*types.Transaction, error) {
	return _PancakePrediction.Contract.ClaimTreasury(&_PancakePrediction.TransactOpts)
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_PancakePrediction *PancakePredictionTransactorSession) ClaimTreasury() (*types.Transaction, error) {
	return _PancakePrediction.Contract.ClaimTreasury(&_PancakePrediction.TransactOpts)
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_PancakePrediction *PancakePredictionTransactor) ExecuteRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "executeRound")
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_PancakePrediction *PancakePredictionSession) ExecuteRound() (*types.Transaction, error) {
	return _PancakePrediction.Contract.ExecuteRound(&_PancakePrediction.TransactOpts)
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_PancakePrediction *PancakePredictionTransactorSession) ExecuteRound() (*types.Transaction, error) {
	return _PancakePrediction.Contract.ExecuteRound(&_PancakePrediction.TransactOpts)
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_PancakePrediction *PancakePredictionTransactor) GenesisLockRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "genesisLockRound")
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_PancakePrediction *PancakePredictionSession) GenesisLockRound() (*types.Transaction, error) {
	return _PancakePrediction.Contract.GenesisLockRound(&_PancakePrediction.TransactOpts)
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_PancakePrediction *PancakePredictionTransactorSession) GenesisLockRound() (*types.Transaction, error) {
	return _PancakePrediction.Contract.GenesisLockRound(&_PancakePrediction.TransactOpts)
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_PancakePrediction *PancakePredictionTransactor) GenesisStartRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "genesisStartRound")
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_PancakePrediction *PancakePredictionSession) GenesisStartRound() (*types.Transaction, error) {
	return _PancakePrediction.Contract.GenesisStartRound(&_PancakePrediction.TransactOpts)
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_PancakePrediction *PancakePredictionTransactorSession) GenesisStartRound() (*types.Transaction, error) {
	return _PancakePrediction.Contract.GenesisStartRound(&_PancakePrediction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PancakePrediction *PancakePredictionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PancakePrediction *PancakePredictionSession) Pause() (*types.Transaction, error) {
	return _PancakePrediction.Contract.Pause(&_PancakePrediction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PancakePrediction *PancakePredictionTransactorSession) Pause() (*types.Transaction, error) {
	return _PancakePrediction.Contract.Pause(&_PancakePrediction.TransactOpts)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xb29a8140.
//
// Solidity: function recoverToken(address _token, uint256 _amount) returns()
func (_PancakePrediction *PancakePredictionTransactor) RecoverToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "recoverToken", _token, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xb29a8140.
//
// Solidity: function recoverToken(address _token, uint256 _amount) returns()
func (_PancakePrediction *PancakePredictionSession) RecoverToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.RecoverToken(&_PancakePrediction.TransactOpts, _token, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xb29a8140.
//
// Solidity: function recoverToken(address _token, uint256 _amount) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) RecoverToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.RecoverToken(&_PancakePrediction.TransactOpts, _token, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakePrediction *PancakePredictionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakePrediction *PancakePredictionSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakePrediction.Contract.RenounceOwnership(&_PancakePrediction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakePrediction *PancakePredictionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakePrediction.Contract.RenounceOwnership(&_PancakePrediction.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_PancakePrediction *PancakePredictionTransactor) SetAdmin(opts *bind.TransactOpts, _adminAddress common.Address) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "setAdmin", _adminAddress)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_PancakePrediction *PancakePredictionSession) SetAdmin(_adminAddress common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetAdmin(&_PancakePrediction.TransactOpts, _adminAddress)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) SetAdmin(_adminAddress common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetAdmin(&_PancakePrediction.TransactOpts, _adminAddress)
}

// SetBufferAndIntervalSeconds is a paid mutator transaction binding the contract method 0x890dc766.
//
// Solidity: function setBufferAndIntervalSeconds(uint256 _bufferSeconds, uint256 _intervalSeconds) returns()
func (_PancakePrediction *PancakePredictionTransactor) SetBufferAndIntervalSeconds(opts *bind.TransactOpts, _bufferSeconds *big.Int, _intervalSeconds *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "setBufferAndIntervalSeconds", _bufferSeconds, _intervalSeconds)
}

// SetBufferAndIntervalSeconds is a paid mutator transaction binding the contract method 0x890dc766.
//
// Solidity: function setBufferAndIntervalSeconds(uint256 _bufferSeconds, uint256 _intervalSeconds) returns()
func (_PancakePrediction *PancakePredictionSession) SetBufferAndIntervalSeconds(_bufferSeconds *big.Int, _intervalSeconds *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetBufferAndIntervalSeconds(&_PancakePrediction.TransactOpts, _bufferSeconds, _intervalSeconds)
}

// SetBufferAndIntervalSeconds is a paid mutator transaction binding the contract method 0x890dc766.
//
// Solidity: function setBufferAndIntervalSeconds(uint256 _bufferSeconds, uint256 _intervalSeconds) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) SetBufferAndIntervalSeconds(_bufferSeconds *big.Int, _intervalSeconds *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetBufferAndIntervalSeconds(&_PancakePrediction.TransactOpts, _bufferSeconds, _intervalSeconds)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_PancakePrediction *PancakePredictionTransactor) SetMinBetAmount(opts *bind.TransactOpts, _minBetAmount *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "setMinBetAmount", _minBetAmount)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_PancakePrediction *PancakePredictionSession) SetMinBetAmount(_minBetAmount *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetMinBetAmount(&_PancakePrediction.TransactOpts, _minBetAmount)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) SetMinBetAmount(_minBetAmount *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetMinBetAmount(&_PancakePrediction.TransactOpts, _minBetAmount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_PancakePrediction *PancakePredictionTransactor) SetOperator(opts *bind.TransactOpts, _operatorAddress common.Address) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "setOperator", _operatorAddress)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_PancakePrediction *PancakePredictionSession) SetOperator(_operatorAddress common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetOperator(&_PancakePrediction.TransactOpts, _operatorAddress)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) SetOperator(_operatorAddress common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetOperator(&_PancakePrediction.TransactOpts, _operatorAddress)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PancakePrediction *PancakePredictionTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PancakePrediction *PancakePredictionSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetOracle(&_PancakePrediction.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetOracle(&_PancakePrediction.TransactOpts, _oracle)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_PancakePrediction *PancakePredictionTransactor) SetOracleUpdateAllowance(opts *bind.TransactOpts, _oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "setOracleUpdateAllowance", _oracleUpdateAllowance)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_PancakePrediction *PancakePredictionSession) SetOracleUpdateAllowance(_oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetOracleUpdateAllowance(&_PancakePrediction.TransactOpts, _oracleUpdateAllowance)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) SetOracleUpdateAllowance(_oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetOracleUpdateAllowance(&_PancakePrediction.TransactOpts, _oracleUpdateAllowance)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_PancakePrediction *PancakePredictionTransactor) SetTreasuryFee(opts *bind.TransactOpts, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "setTreasuryFee", _treasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_PancakePrediction *PancakePredictionSession) SetTreasuryFee(_treasuryFee *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetTreasuryFee(&_PancakePrediction.TransactOpts, _treasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0x77e741c7.
//
// Solidity: function setTreasuryFee(uint256 _treasuryFee) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) SetTreasuryFee(_treasuryFee *big.Int) (*types.Transaction, error) {
	return _PancakePrediction.Contract.SetTreasuryFee(&_PancakePrediction.TransactOpts, _treasuryFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakePrediction *PancakePredictionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakePrediction *PancakePredictionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.TransferOwnership(&_PancakePrediction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakePrediction *PancakePredictionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakePrediction.Contract.TransferOwnership(&_PancakePrediction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PancakePrediction *PancakePredictionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePrediction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PancakePrediction *PancakePredictionSession) Unpause() (*types.Transaction, error) {
	return _PancakePrediction.Contract.Unpause(&_PancakePrediction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PancakePrediction *PancakePredictionTransactorSession) Unpause() (*types.Transaction, error) {
	return _PancakePrediction.Contract.Unpause(&_PancakePrediction.TransactOpts)
}

// PancakePredictionBetBearIterator is returned from FilterBetBear and is used to iterate over the raw logs and unpacked data for BetBear events raised by the PancakePrediction contract.
type PancakePredictionBetBearIterator struct {
	Event *PancakePredictionBetBear // Event containing the contract specifics and raw log

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
func (it *PancakePredictionBetBearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionBetBear)
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
		it.Event = new(PancakePredictionBetBear)
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
func (it *PancakePredictionBetBearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionBetBearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionBetBear represents a BetBear event raised by the PancakePrediction contract.
type PancakePredictionBetBear struct {
	Sender common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetBear is a free log retrieval operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) FilterBetBear(opts *bind.FilterOpts, sender []common.Address, epoch []*big.Int) (*PancakePredictionBetBearIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "BetBear", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionBetBearIterator{contract: _PancakePrediction.contract, event: "BetBear", logs: logs, sub: sub}, nil
}

// WatchBetBear is a free log subscription operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) WatchBetBear(opts *bind.WatchOpts, sink chan<- *PancakePredictionBetBear, sender []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "BetBear", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionBetBear)
				if err := _PancakePrediction.contract.UnpackLog(event, "BetBear", log); err != nil {
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

// ParseBetBear is a log parse operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) ParseBetBear(log types.Log) (*PancakePredictionBetBear, error) {
	event := new(PancakePredictionBetBear)
	if err := _PancakePrediction.contract.UnpackLog(event, "BetBear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionBetBullIterator is returned from FilterBetBull and is used to iterate over the raw logs and unpacked data for BetBull events raised by the PancakePrediction contract.
type PancakePredictionBetBullIterator struct {
	Event *PancakePredictionBetBull // Event containing the contract specifics and raw log

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
func (it *PancakePredictionBetBullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionBetBull)
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
		it.Event = new(PancakePredictionBetBull)
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
func (it *PancakePredictionBetBullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionBetBullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionBetBull represents a BetBull event raised by the PancakePrediction contract.
type PancakePredictionBetBull struct {
	Sender common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetBull is a free log retrieval operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) FilterBetBull(opts *bind.FilterOpts, sender []common.Address, epoch []*big.Int) (*PancakePredictionBetBullIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "BetBull", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionBetBullIterator{contract: _PancakePrediction.contract, event: "BetBull", logs: logs, sub: sub}, nil
}

// WatchBetBull is a free log subscription operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) WatchBetBull(opts *bind.WatchOpts, sink chan<- *PancakePredictionBetBull, sender []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "BetBull", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionBetBull)
				if err := _PancakePrediction.contract.UnpackLog(event, "BetBull", log); err != nil {
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

// ParseBetBull is a log parse operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) ParseBetBull(log types.Log) (*PancakePredictionBetBull, error) {
	event := new(PancakePredictionBetBull)
	if err := _PancakePrediction.contract.UnpackLog(event, "BetBull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the PancakePrediction contract.
type PancakePredictionClaimIterator struct {
	Event *PancakePredictionClaim // Event containing the contract specifics and raw log

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
func (it *PancakePredictionClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionClaim)
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
		it.Event = new(PancakePredictionClaim)
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
func (it *PancakePredictionClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionClaim represents a Claim event raised by the PancakePrediction contract.
type PancakePredictionClaim struct {
	Sender common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) FilterClaim(opts *bind.FilterOpts, sender []common.Address, epoch []*big.Int) (*PancakePredictionClaimIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "Claim", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionClaimIterator{contract: _PancakePrediction.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *PancakePredictionClaim, sender []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "Claim", senderRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionClaim)
				if err := _PancakePrediction.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed epoch, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) ParseClaim(log types.Log) (*PancakePredictionClaim, error) {
	event := new(PancakePredictionClaim)
	if err := _PancakePrediction.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionEndRoundIterator is returned from FilterEndRound and is used to iterate over the raw logs and unpacked data for EndRound events raised by the PancakePrediction contract.
type PancakePredictionEndRoundIterator struct {
	Event *PancakePredictionEndRound // Event containing the contract specifics and raw log

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
func (it *PancakePredictionEndRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionEndRound)
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
		it.Event = new(PancakePredictionEndRound)
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
func (it *PancakePredictionEndRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionEndRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionEndRound represents a EndRound event raised by the PancakePrediction contract.
type PancakePredictionEndRound struct {
	Epoch   *big.Int
	RoundId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEndRound is a free log retrieval operation binding the contract event 0xb6ff1fe915db84788cbbbc017f0d2bef9485fad9fd0bd8ce9340fde0d8410dd8.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_PancakePrediction *PancakePredictionFilterer) FilterEndRound(opts *bind.FilterOpts, epoch []*big.Int, roundId []*big.Int) (*PancakePredictionEndRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "EndRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionEndRoundIterator{contract: _PancakePrediction.contract, event: "EndRound", logs: logs, sub: sub}, nil
}

// WatchEndRound is a free log subscription operation binding the contract event 0xb6ff1fe915db84788cbbbc017f0d2bef9485fad9fd0bd8ce9340fde0d8410dd8.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_PancakePrediction *PancakePredictionFilterer) WatchEndRound(opts *bind.WatchOpts, sink chan<- *PancakePredictionEndRound, epoch []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "EndRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionEndRound)
				if err := _PancakePrediction.contract.UnpackLog(event, "EndRound", log); err != nil {
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

// ParseEndRound is a log parse operation binding the contract event 0xb6ff1fe915db84788cbbbc017f0d2bef9485fad9fd0bd8ce9340fde0d8410dd8.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_PancakePrediction *PancakePredictionFilterer) ParseEndRound(log types.Log) (*PancakePredictionEndRound, error) {
	event := new(PancakePredictionEndRound)
	if err := _PancakePrediction.contract.UnpackLog(event, "EndRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionLockRoundIterator is returned from FilterLockRound and is used to iterate over the raw logs and unpacked data for LockRound events raised by the PancakePrediction contract.
type PancakePredictionLockRoundIterator struct {
	Event *PancakePredictionLockRound // Event containing the contract specifics and raw log

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
func (it *PancakePredictionLockRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionLockRound)
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
		it.Event = new(PancakePredictionLockRound)
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
func (it *PancakePredictionLockRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionLockRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionLockRound represents a LockRound event raised by the PancakePrediction contract.
type PancakePredictionLockRound struct {
	Epoch   *big.Int
	RoundId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockRound is a free log retrieval operation binding the contract event 0x482e76a65b448a42deef26e99e58fb20c85e26f075defff8df6aa80459b39006.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_PancakePrediction *PancakePredictionFilterer) FilterLockRound(opts *bind.FilterOpts, epoch []*big.Int, roundId []*big.Int) (*PancakePredictionLockRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "LockRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionLockRoundIterator{contract: _PancakePrediction.contract, event: "LockRound", logs: logs, sub: sub}, nil
}

// WatchLockRound is a free log subscription operation binding the contract event 0x482e76a65b448a42deef26e99e58fb20c85e26f075defff8df6aa80459b39006.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_PancakePrediction *PancakePredictionFilterer) WatchLockRound(opts *bind.WatchOpts, sink chan<- *PancakePredictionLockRound, epoch []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "LockRound", epochRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionLockRound)
				if err := _PancakePrediction.contract.UnpackLog(event, "LockRound", log); err != nil {
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

// ParseLockRound is a log parse operation binding the contract event 0x482e76a65b448a42deef26e99e58fb20c85e26f075defff8df6aa80459b39006.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 indexed roundId, int256 price)
func (_PancakePrediction *PancakePredictionFilterer) ParseLockRound(log types.Log) (*PancakePredictionLockRound, error) {
	event := new(PancakePredictionLockRound)
	if err := _PancakePrediction.contract.UnpackLog(event, "LockRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionNewAdminAddressIterator is returned from FilterNewAdminAddress and is used to iterate over the raw logs and unpacked data for NewAdminAddress events raised by the PancakePrediction contract.
type PancakePredictionNewAdminAddressIterator struct {
	Event *PancakePredictionNewAdminAddress // Event containing the contract specifics and raw log

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
func (it *PancakePredictionNewAdminAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionNewAdminAddress)
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
		it.Event = new(PancakePredictionNewAdminAddress)
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
func (it *PancakePredictionNewAdminAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionNewAdminAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionNewAdminAddress represents a NewAdminAddress event raised by the PancakePrediction contract.
type PancakePredictionNewAdminAddress struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdminAddress is a free log retrieval operation binding the contract event 0x137b621413925496477d46e5055ac0d56178bdd724ba8bf843afceef18268ba3.
//
// Solidity: event NewAdminAddress(address admin)
func (_PancakePrediction *PancakePredictionFilterer) FilterNewAdminAddress(opts *bind.FilterOpts) (*PancakePredictionNewAdminAddressIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "NewAdminAddress")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionNewAdminAddressIterator{contract: _PancakePrediction.contract, event: "NewAdminAddress", logs: logs, sub: sub}, nil
}

// WatchNewAdminAddress is a free log subscription operation binding the contract event 0x137b621413925496477d46e5055ac0d56178bdd724ba8bf843afceef18268ba3.
//
// Solidity: event NewAdminAddress(address admin)
func (_PancakePrediction *PancakePredictionFilterer) WatchNewAdminAddress(opts *bind.WatchOpts, sink chan<- *PancakePredictionNewAdminAddress) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "NewAdminAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionNewAdminAddress)
				if err := _PancakePrediction.contract.UnpackLog(event, "NewAdminAddress", log); err != nil {
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

// ParseNewAdminAddress is a log parse operation binding the contract event 0x137b621413925496477d46e5055ac0d56178bdd724ba8bf843afceef18268ba3.
//
// Solidity: event NewAdminAddress(address admin)
func (_PancakePrediction *PancakePredictionFilterer) ParseNewAdminAddress(log types.Log) (*PancakePredictionNewAdminAddress, error) {
	event := new(PancakePredictionNewAdminAddress)
	if err := _PancakePrediction.contract.UnpackLog(event, "NewAdminAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionNewBufferAndIntervalSecondsIterator is returned from FilterNewBufferAndIntervalSeconds and is used to iterate over the raw logs and unpacked data for NewBufferAndIntervalSeconds events raised by the PancakePrediction contract.
type PancakePredictionNewBufferAndIntervalSecondsIterator struct {
	Event *PancakePredictionNewBufferAndIntervalSeconds // Event containing the contract specifics and raw log

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
func (it *PancakePredictionNewBufferAndIntervalSecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionNewBufferAndIntervalSeconds)
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
		it.Event = new(PancakePredictionNewBufferAndIntervalSeconds)
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
func (it *PancakePredictionNewBufferAndIntervalSecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionNewBufferAndIntervalSecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionNewBufferAndIntervalSeconds represents a NewBufferAndIntervalSeconds event raised by the PancakePrediction contract.
type PancakePredictionNewBufferAndIntervalSeconds struct {
	BufferSeconds   *big.Int
	IntervalSeconds *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewBufferAndIntervalSeconds is a free log retrieval operation binding the contract event 0xe60149e0431fec12df63dfab5fce2a9cefe9a4d3df5f41cb626f579ae1f2b91a.
//
// Solidity: event NewBufferAndIntervalSeconds(uint256 bufferSeconds, uint256 intervalSeconds)
func (_PancakePrediction *PancakePredictionFilterer) FilterNewBufferAndIntervalSeconds(opts *bind.FilterOpts) (*PancakePredictionNewBufferAndIntervalSecondsIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "NewBufferAndIntervalSeconds")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionNewBufferAndIntervalSecondsIterator{contract: _PancakePrediction.contract, event: "NewBufferAndIntervalSeconds", logs: logs, sub: sub}, nil
}

// WatchNewBufferAndIntervalSeconds is a free log subscription operation binding the contract event 0xe60149e0431fec12df63dfab5fce2a9cefe9a4d3df5f41cb626f579ae1f2b91a.
//
// Solidity: event NewBufferAndIntervalSeconds(uint256 bufferSeconds, uint256 intervalSeconds)
func (_PancakePrediction *PancakePredictionFilterer) WatchNewBufferAndIntervalSeconds(opts *bind.WatchOpts, sink chan<- *PancakePredictionNewBufferAndIntervalSeconds) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "NewBufferAndIntervalSeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionNewBufferAndIntervalSeconds)
				if err := _PancakePrediction.contract.UnpackLog(event, "NewBufferAndIntervalSeconds", log); err != nil {
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

// ParseNewBufferAndIntervalSeconds is a log parse operation binding the contract event 0xe60149e0431fec12df63dfab5fce2a9cefe9a4d3df5f41cb626f579ae1f2b91a.
//
// Solidity: event NewBufferAndIntervalSeconds(uint256 bufferSeconds, uint256 intervalSeconds)
func (_PancakePrediction *PancakePredictionFilterer) ParseNewBufferAndIntervalSeconds(log types.Log) (*PancakePredictionNewBufferAndIntervalSeconds, error) {
	event := new(PancakePredictionNewBufferAndIntervalSeconds)
	if err := _PancakePrediction.contract.UnpackLog(event, "NewBufferAndIntervalSeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionNewMinBetAmountIterator is returned from FilterNewMinBetAmount and is used to iterate over the raw logs and unpacked data for NewMinBetAmount events raised by the PancakePrediction contract.
type PancakePredictionNewMinBetAmountIterator struct {
	Event *PancakePredictionNewMinBetAmount // Event containing the contract specifics and raw log

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
func (it *PancakePredictionNewMinBetAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionNewMinBetAmount)
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
		it.Event = new(PancakePredictionNewMinBetAmount)
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
func (it *PancakePredictionNewMinBetAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionNewMinBetAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionNewMinBetAmount represents a NewMinBetAmount event raised by the PancakePrediction contract.
type PancakePredictionNewMinBetAmount struct {
	Epoch        *big.Int
	MinBetAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewMinBetAmount is a free log retrieval operation binding the contract event 0x90eb87c560a0213754ceb3a7fa3012f01acab0a35602c1e1995adf69dabc9d50.
//
// Solidity: event NewMinBetAmount(uint256 indexed epoch, uint256 minBetAmount)
func (_PancakePrediction *PancakePredictionFilterer) FilterNewMinBetAmount(opts *bind.FilterOpts, epoch []*big.Int) (*PancakePredictionNewMinBetAmountIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "NewMinBetAmount", epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionNewMinBetAmountIterator{contract: _PancakePrediction.contract, event: "NewMinBetAmount", logs: logs, sub: sub}, nil
}

// WatchNewMinBetAmount is a free log subscription operation binding the contract event 0x90eb87c560a0213754ceb3a7fa3012f01acab0a35602c1e1995adf69dabc9d50.
//
// Solidity: event NewMinBetAmount(uint256 indexed epoch, uint256 minBetAmount)
func (_PancakePrediction *PancakePredictionFilterer) WatchNewMinBetAmount(opts *bind.WatchOpts, sink chan<- *PancakePredictionNewMinBetAmount, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "NewMinBetAmount", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionNewMinBetAmount)
				if err := _PancakePrediction.contract.UnpackLog(event, "NewMinBetAmount", log); err != nil {
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

// ParseNewMinBetAmount is a log parse operation binding the contract event 0x90eb87c560a0213754ceb3a7fa3012f01acab0a35602c1e1995adf69dabc9d50.
//
// Solidity: event NewMinBetAmount(uint256 indexed epoch, uint256 minBetAmount)
func (_PancakePrediction *PancakePredictionFilterer) ParseNewMinBetAmount(log types.Log) (*PancakePredictionNewMinBetAmount, error) {
	event := new(PancakePredictionNewMinBetAmount)
	if err := _PancakePrediction.contract.UnpackLog(event, "NewMinBetAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionNewOperatorAddressIterator is returned from FilterNewOperatorAddress and is used to iterate over the raw logs and unpacked data for NewOperatorAddress events raised by the PancakePrediction contract.
type PancakePredictionNewOperatorAddressIterator struct {
	Event *PancakePredictionNewOperatorAddress // Event containing the contract specifics and raw log

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
func (it *PancakePredictionNewOperatorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionNewOperatorAddress)
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
		it.Event = new(PancakePredictionNewOperatorAddress)
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
func (it *PancakePredictionNewOperatorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionNewOperatorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionNewOperatorAddress represents a NewOperatorAddress event raised by the PancakePrediction contract.
type PancakePredictionNewOperatorAddress struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOperatorAddress is a free log retrieval operation binding the contract event 0xc47d127c07bdd56c5ccba00463ce3bd3c1bca71b4670eea6e5d0c02e4aa156e2.
//
// Solidity: event NewOperatorAddress(address operator)
func (_PancakePrediction *PancakePredictionFilterer) FilterNewOperatorAddress(opts *bind.FilterOpts) (*PancakePredictionNewOperatorAddressIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "NewOperatorAddress")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionNewOperatorAddressIterator{contract: _PancakePrediction.contract, event: "NewOperatorAddress", logs: logs, sub: sub}, nil
}

// WatchNewOperatorAddress is a free log subscription operation binding the contract event 0xc47d127c07bdd56c5ccba00463ce3bd3c1bca71b4670eea6e5d0c02e4aa156e2.
//
// Solidity: event NewOperatorAddress(address operator)
func (_PancakePrediction *PancakePredictionFilterer) WatchNewOperatorAddress(opts *bind.WatchOpts, sink chan<- *PancakePredictionNewOperatorAddress) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "NewOperatorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionNewOperatorAddress)
				if err := _PancakePrediction.contract.UnpackLog(event, "NewOperatorAddress", log); err != nil {
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

// ParseNewOperatorAddress is a log parse operation binding the contract event 0xc47d127c07bdd56c5ccba00463ce3bd3c1bca71b4670eea6e5d0c02e4aa156e2.
//
// Solidity: event NewOperatorAddress(address operator)
func (_PancakePrediction *PancakePredictionFilterer) ParseNewOperatorAddress(log types.Log) (*PancakePredictionNewOperatorAddress, error) {
	event := new(PancakePredictionNewOperatorAddress)
	if err := _PancakePrediction.contract.UnpackLog(event, "NewOperatorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionNewOracleIterator is returned from FilterNewOracle and is used to iterate over the raw logs and unpacked data for NewOracle events raised by the PancakePrediction contract.
type PancakePredictionNewOracleIterator struct {
	Event *PancakePredictionNewOracle // Event containing the contract specifics and raw log

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
func (it *PancakePredictionNewOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionNewOracle)
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
		it.Event = new(PancakePredictionNewOracle)
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
func (it *PancakePredictionNewOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionNewOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionNewOracle represents a NewOracle event raised by the PancakePrediction contract.
type PancakePredictionNewOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewOracle is a free log retrieval operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address oracle)
func (_PancakePrediction *PancakePredictionFilterer) FilterNewOracle(opts *bind.FilterOpts) (*PancakePredictionNewOracleIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "NewOracle")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionNewOracleIterator{contract: _PancakePrediction.contract, event: "NewOracle", logs: logs, sub: sub}, nil
}

// WatchNewOracle is a free log subscription operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address oracle)
func (_PancakePrediction *PancakePredictionFilterer) WatchNewOracle(opts *bind.WatchOpts, sink chan<- *PancakePredictionNewOracle) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "NewOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionNewOracle)
				if err := _PancakePrediction.contract.UnpackLog(event, "NewOracle", log); err != nil {
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

// ParseNewOracle is a log parse operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address oracle)
func (_PancakePrediction *PancakePredictionFilterer) ParseNewOracle(log types.Log) (*PancakePredictionNewOracle, error) {
	event := new(PancakePredictionNewOracle)
	if err := _PancakePrediction.contract.UnpackLog(event, "NewOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionNewOracleUpdateAllowanceIterator is returned from FilterNewOracleUpdateAllowance and is used to iterate over the raw logs and unpacked data for NewOracleUpdateAllowance events raised by the PancakePrediction contract.
type PancakePredictionNewOracleUpdateAllowanceIterator struct {
	Event *PancakePredictionNewOracleUpdateAllowance // Event containing the contract specifics and raw log

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
func (it *PancakePredictionNewOracleUpdateAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionNewOracleUpdateAllowance)
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
		it.Event = new(PancakePredictionNewOracleUpdateAllowance)
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
func (it *PancakePredictionNewOracleUpdateAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionNewOracleUpdateAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionNewOracleUpdateAllowance represents a NewOracleUpdateAllowance event raised by the PancakePrediction contract.
type PancakePredictionNewOracleUpdateAllowance struct {
	OracleUpdateAllowance *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewOracleUpdateAllowance is a free log retrieval operation binding the contract event 0x93ccaceac092ffb842c46b8718667a13a80e9058dcd0bd403d0b47215b30da07.
//
// Solidity: event NewOracleUpdateAllowance(uint256 oracleUpdateAllowance)
func (_PancakePrediction *PancakePredictionFilterer) FilterNewOracleUpdateAllowance(opts *bind.FilterOpts) (*PancakePredictionNewOracleUpdateAllowanceIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "NewOracleUpdateAllowance")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionNewOracleUpdateAllowanceIterator{contract: _PancakePrediction.contract, event: "NewOracleUpdateAllowance", logs: logs, sub: sub}, nil
}

// WatchNewOracleUpdateAllowance is a free log subscription operation binding the contract event 0x93ccaceac092ffb842c46b8718667a13a80e9058dcd0bd403d0b47215b30da07.
//
// Solidity: event NewOracleUpdateAllowance(uint256 oracleUpdateAllowance)
func (_PancakePrediction *PancakePredictionFilterer) WatchNewOracleUpdateAllowance(opts *bind.WatchOpts, sink chan<- *PancakePredictionNewOracleUpdateAllowance) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "NewOracleUpdateAllowance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionNewOracleUpdateAllowance)
				if err := _PancakePrediction.contract.UnpackLog(event, "NewOracleUpdateAllowance", log); err != nil {
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

// ParseNewOracleUpdateAllowance is a log parse operation binding the contract event 0x93ccaceac092ffb842c46b8718667a13a80e9058dcd0bd403d0b47215b30da07.
//
// Solidity: event NewOracleUpdateAllowance(uint256 oracleUpdateAllowance)
func (_PancakePrediction *PancakePredictionFilterer) ParseNewOracleUpdateAllowance(log types.Log) (*PancakePredictionNewOracleUpdateAllowance, error) {
	event := new(PancakePredictionNewOracleUpdateAllowance)
	if err := _PancakePrediction.contract.UnpackLog(event, "NewOracleUpdateAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionNewTreasuryFeeIterator is returned from FilterNewTreasuryFee and is used to iterate over the raw logs and unpacked data for NewTreasuryFee events raised by the PancakePrediction contract.
type PancakePredictionNewTreasuryFeeIterator struct {
	Event *PancakePredictionNewTreasuryFee // Event containing the contract specifics and raw log

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
func (it *PancakePredictionNewTreasuryFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionNewTreasuryFee)
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
		it.Event = new(PancakePredictionNewTreasuryFee)
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
func (it *PancakePredictionNewTreasuryFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionNewTreasuryFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionNewTreasuryFee represents a NewTreasuryFee event raised by the PancakePrediction contract.
type PancakePredictionNewTreasuryFee struct {
	Epoch       *big.Int
	TreasuryFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryFee is a free log retrieval operation binding the contract event 0xb1c4ee38d35556741133da7ff9b6f7ab0fa88d0406133126ff128f635490a857.
//
// Solidity: event NewTreasuryFee(uint256 indexed epoch, uint256 treasuryFee)
func (_PancakePrediction *PancakePredictionFilterer) FilterNewTreasuryFee(opts *bind.FilterOpts, epoch []*big.Int) (*PancakePredictionNewTreasuryFeeIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "NewTreasuryFee", epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionNewTreasuryFeeIterator{contract: _PancakePrediction.contract, event: "NewTreasuryFee", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryFee is a free log subscription operation binding the contract event 0xb1c4ee38d35556741133da7ff9b6f7ab0fa88d0406133126ff128f635490a857.
//
// Solidity: event NewTreasuryFee(uint256 indexed epoch, uint256 treasuryFee)
func (_PancakePrediction *PancakePredictionFilterer) WatchNewTreasuryFee(opts *bind.WatchOpts, sink chan<- *PancakePredictionNewTreasuryFee, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "NewTreasuryFee", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionNewTreasuryFee)
				if err := _PancakePrediction.contract.UnpackLog(event, "NewTreasuryFee", log); err != nil {
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

// ParseNewTreasuryFee is a log parse operation binding the contract event 0xb1c4ee38d35556741133da7ff9b6f7ab0fa88d0406133126ff128f635490a857.
//
// Solidity: event NewTreasuryFee(uint256 indexed epoch, uint256 treasuryFee)
func (_PancakePrediction *PancakePredictionFilterer) ParseNewTreasuryFee(log types.Log) (*PancakePredictionNewTreasuryFee, error) {
	event := new(PancakePredictionNewTreasuryFee)
	if err := _PancakePrediction.contract.UnpackLog(event, "NewTreasuryFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PancakePrediction contract.
type PancakePredictionOwnershipTransferredIterator struct {
	Event *PancakePredictionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakePredictionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionOwnershipTransferred)
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
		it.Event = new(PancakePredictionOwnershipTransferred)
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
func (it *PancakePredictionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionOwnershipTransferred represents a OwnershipTransferred event raised by the PancakePrediction contract.
type PancakePredictionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakePrediction *PancakePredictionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PancakePredictionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionOwnershipTransferredIterator{contract: _PancakePrediction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakePrediction *PancakePredictionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakePredictionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionOwnershipTransferred)
				if err := _PancakePrediction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakePrediction *PancakePredictionFilterer) ParseOwnershipTransferred(log types.Log) (*PancakePredictionOwnershipTransferred, error) {
	event := new(PancakePredictionOwnershipTransferred)
	if err := _PancakePrediction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PancakePrediction contract.
type PancakePredictionPauseIterator struct {
	Event *PancakePredictionPause // Event containing the contract specifics and raw log

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
func (it *PancakePredictionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionPause)
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
		it.Event = new(PancakePredictionPause)
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
func (it *PancakePredictionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionPause represents a Pause event raised by the PancakePrediction contract.
type PancakePredictionPause struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) FilterPause(opts *bind.FilterOpts, epoch []*big.Int) (*PancakePredictionPauseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "Pause", epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionPauseIterator{contract: _PancakePrediction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PancakePredictionPause, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "Pause", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionPause)
				if err := _PancakePrediction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) ParsePause(log types.Log) (*PancakePredictionPause, error) {
	event := new(PancakePredictionPause)
	if err := _PancakePrediction.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PancakePrediction contract.
type PancakePredictionPausedIterator struct {
	Event *PancakePredictionPaused // Event containing the contract specifics and raw log

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
func (it *PancakePredictionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionPaused)
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
		it.Event = new(PancakePredictionPaused)
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
func (it *PancakePredictionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionPaused represents a Paused event raised by the PancakePrediction contract.
type PancakePredictionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PancakePrediction *PancakePredictionFilterer) FilterPaused(opts *bind.FilterOpts) (*PancakePredictionPausedIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionPausedIterator{contract: _PancakePrediction.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PancakePrediction *PancakePredictionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PancakePredictionPaused) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionPaused)
				if err := _PancakePrediction.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PancakePrediction *PancakePredictionFilterer) ParsePaused(log types.Log) (*PancakePredictionPaused, error) {
	event := new(PancakePredictionPaused)
	if err := _PancakePrediction.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionRewardsCalculatedIterator is returned from FilterRewardsCalculated and is used to iterate over the raw logs and unpacked data for RewardsCalculated events raised by the PancakePrediction contract.
type PancakePredictionRewardsCalculatedIterator struct {
	Event *PancakePredictionRewardsCalculated // Event containing the contract specifics and raw log

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
func (it *PancakePredictionRewardsCalculatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionRewardsCalculated)
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
		it.Event = new(PancakePredictionRewardsCalculated)
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
func (it *PancakePredictionRewardsCalculatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionRewardsCalculatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionRewardsCalculated represents a RewardsCalculated event raised by the PancakePrediction contract.
type PancakePredictionRewardsCalculated struct {
	Epoch               *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	TreasuryAmount      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRewardsCalculated is a free log retrieval operation binding the contract event 0x6dfdfcb09c8804d0058826cd2539f1acfbe3cb887c9be03d928035bce0f1a58d.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount)
func (_PancakePrediction *PancakePredictionFilterer) FilterRewardsCalculated(opts *bind.FilterOpts, epoch []*big.Int) (*PancakePredictionRewardsCalculatedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "RewardsCalculated", epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionRewardsCalculatedIterator{contract: _PancakePrediction.contract, event: "RewardsCalculated", logs: logs, sub: sub}, nil
}

// WatchRewardsCalculated is a free log subscription operation binding the contract event 0x6dfdfcb09c8804d0058826cd2539f1acfbe3cb887c9be03d928035bce0f1a58d.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount)
func (_PancakePrediction *PancakePredictionFilterer) WatchRewardsCalculated(opts *bind.WatchOpts, sink chan<- *PancakePredictionRewardsCalculated, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "RewardsCalculated", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionRewardsCalculated)
				if err := _PancakePrediction.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
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

// ParseRewardsCalculated is a log parse operation binding the contract event 0x6dfdfcb09c8804d0058826cd2539f1acfbe3cb887c9be03d928035bce0f1a58d.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount)
func (_PancakePrediction *PancakePredictionFilterer) ParseRewardsCalculated(log types.Log) (*PancakePredictionRewardsCalculated, error) {
	event := new(PancakePredictionRewardsCalculated)
	if err := _PancakePrediction.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionStartRoundIterator is returned from FilterStartRound and is used to iterate over the raw logs and unpacked data for StartRound events raised by the PancakePrediction contract.
type PancakePredictionStartRoundIterator struct {
	Event *PancakePredictionStartRound // Event containing the contract specifics and raw log

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
func (it *PancakePredictionStartRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionStartRound)
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
		it.Event = new(PancakePredictionStartRound)
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
func (it *PancakePredictionStartRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionStartRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionStartRound represents a StartRound event raised by the PancakePrediction contract.
type PancakePredictionStartRound struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStartRound is a free log retrieval operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) FilterStartRound(opts *bind.FilterOpts, epoch []*big.Int) (*PancakePredictionStartRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionStartRoundIterator{contract: _PancakePrediction.contract, event: "StartRound", logs: logs, sub: sub}, nil
}

// WatchStartRound is a free log subscription operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) WatchStartRound(opts *bind.WatchOpts, sink chan<- *PancakePredictionStartRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionStartRound)
				if err := _PancakePrediction.contract.UnpackLog(event, "StartRound", log); err != nil {
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

// ParseStartRound is a log parse operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) ParseStartRound(log types.Log) (*PancakePredictionStartRound, error) {
	event := new(PancakePredictionStartRound)
	if err := _PancakePrediction.contract.UnpackLog(event, "StartRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionTokenRecoveryIterator is returned from FilterTokenRecovery and is used to iterate over the raw logs and unpacked data for TokenRecovery events raised by the PancakePrediction contract.
type PancakePredictionTokenRecoveryIterator struct {
	Event *PancakePredictionTokenRecovery // Event containing the contract specifics and raw log

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
func (it *PancakePredictionTokenRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionTokenRecovery)
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
		it.Event = new(PancakePredictionTokenRecovery)
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
func (it *PancakePredictionTokenRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionTokenRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionTokenRecovery represents a TokenRecovery event raised by the PancakePrediction contract.
type PancakePredictionTokenRecovery struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRecovery is a free log retrieval operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) FilterTokenRecovery(opts *bind.FilterOpts, token []common.Address) (*PancakePredictionTokenRecoveryIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionTokenRecoveryIterator{contract: _PancakePrediction.contract, event: "TokenRecovery", logs: logs, sub: sub}, nil
}

// WatchTokenRecovery is a free log subscription operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) WatchTokenRecovery(opts *bind.WatchOpts, sink chan<- *PancakePredictionTokenRecovery, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionTokenRecovery)
				if err := _PancakePrediction.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
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

// ParseTokenRecovery is a log parse operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) ParseTokenRecovery(log types.Log) (*PancakePredictionTokenRecovery, error) {
	event := new(PancakePredictionTokenRecovery)
	if err := _PancakePrediction.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionTreasuryClaimIterator is returned from FilterTreasuryClaim and is used to iterate over the raw logs and unpacked data for TreasuryClaim events raised by the PancakePrediction contract.
type PancakePredictionTreasuryClaimIterator struct {
	Event *PancakePredictionTreasuryClaim // Event containing the contract specifics and raw log

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
func (it *PancakePredictionTreasuryClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionTreasuryClaim)
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
		it.Event = new(PancakePredictionTreasuryClaim)
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
func (it *PancakePredictionTreasuryClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionTreasuryClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionTreasuryClaim represents a TreasuryClaim event raised by the PancakePrediction contract.
type PancakePredictionTreasuryClaim struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTreasuryClaim is a free log retrieval operation binding the contract event 0xb9197c6b8e21274bd1e2d9c956a88af5cfee510f630fab3f046300f88b422361.
//
// Solidity: event TreasuryClaim(uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) FilterTreasuryClaim(opts *bind.FilterOpts) (*PancakePredictionTreasuryClaimIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "TreasuryClaim")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionTreasuryClaimIterator{contract: _PancakePrediction.contract, event: "TreasuryClaim", logs: logs, sub: sub}, nil
}

// WatchTreasuryClaim is a free log subscription operation binding the contract event 0xb9197c6b8e21274bd1e2d9c956a88af5cfee510f630fab3f046300f88b422361.
//
// Solidity: event TreasuryClaim(uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) WatchTreasuryClaim(opts *bind.WatchOpts, sink chan<- *PancakePredictionTreasuryClaim) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "TreasuryClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionTreasuryClaim)
				if err := _PancakePrediction.contract.UnpackLog(event, "TreasuryClaim", log); err != nil {
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

// ParseTreasuryClaim is a log parse operation binding the contract event 0xb9197c6b8e21274bd1e2d9c956a88af5cfee510f630fab3f046300f88b422361.
//
// Solidity: event TreasuryClaim(uint256 amount)
func (_PancakePrediction *PancakePredictionFilterer) ParseTreasuryClaim(log types.Log) (*PancakePredictionTreasuryClaim, error) {
	event := new(PancakePredictionTreasuryClaim)
	if err := _PancakePrediction.contract.UnpackLog(event, "TreasuryClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PancakePrediction contract.
type PancakePredictionUnpauseIterator struct {
	Event *PancakePredictionUnpause // Event containing the contract specifics and raw log

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
func (it *PancakePredictionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionUnpause)
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
		it.Event = new(PancakePredictionUnpause)
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
func (it *PancakePredictionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionUnpause represents a Unpause event raised by the PancakePrediction contract.
type PancakePredictionUnpause struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) FilterUnpause(opts *bind.FilterOpts, epoch []*big.Int) (*PancakePredictionUnpauseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "Unpause", epochRule)
	if err != nil {
		return nil, err
	}
	return &PancakePredictionUnpauseIterator{contract: _PancakePrediction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PancakePredictionUnpause, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "Unpause", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionUnpause)
				if err := _PancakePrediction.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 indexed epoch)
func (_PancakePrediction *PancakePredictionFilterer) ParseUnpause(log types.Log) (*PancakePredictionUnpause, error) {
	event := new(PancakePredictionUnpause)
	if err := _PancakePrediction.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakePredictionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PancakePrediction contract.
type PancakePredictionUnpausedIterator struct {
	Event *PancakePredictionUnpaused // Event containing the contract specifics and raw log

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
func (it *PancakePredictionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakePredictionUnpaused)
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
		it.Event = new(PancakePredictionUnpaused)
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
func (it *PancakePredictionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakePredictionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakePredictionUnpaused represents a Unpaused event raised by the PancakePrediction contract.
type PancakePredictionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PancakePrediction *PancakePredictionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PancakePredictionUnpausedIterator, error) {

	logs, sub, err := _PancakePrediction.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PancakePredictionUnpausedIterator{contract: _PancakePrediction.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PancakePrediction *PancakePredictionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PancakePredictionUnpaused) (event.Subscription, error) {

	logs, sub, err := _PancakePrediction.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakePredictionUnpaused)
				if err := _PancakePrediction.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PancakePrediction *PancakePredictionFilterer) ParseUnpaused(log types.Log) (*PancakePredictionUnpaused, error) {
	event := new(PancakePredictionUnpaused)
	if err := _PancakePrediction.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
