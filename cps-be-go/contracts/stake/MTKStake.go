// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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

// MTKStakeStake is an auto generated low-level Go binding around an user-defined struct.
type MTKStakeStake struct {
	StakeId        *big.Int
	Amount         *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	RewardRateYear *big.Int
	IsActive       bool
	Period         uint8
}

// MTKStakeMetaData contains all meta data concerning the MTKStake contract.
var MTKStakeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_mtkToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"apy\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumMTKStake.StakingPeriod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"durations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumMTKStake.StakingPeriod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserActiveStakes\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMTKStake.Stake[]\",\"components\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardRateYear\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumMTKStake.StakingPeriod\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumMTKStake.StakingPeriod\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeIdOwner\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userStakes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardRateYear\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumMTKStake.StakingPeriod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"stakeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumMTKStake.StakingPeriod\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stakeId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"principal\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MTKStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use MTKStakeMetaData.ABI instead.
var MTKStakeABI = MTKStakeMetaData.ABI

// MTKStake is an auto generated Go binding around an Ethereum contract.
type MTKStake struct {
	MTKStakeCaller     // Read-only binding to the contract
	MTKStakeTransactor // Write-only binding to the contract
	MTKStakeFilterer   // Log filterer for contract events
}

// MTKStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MTKStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTKStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MTKStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTKStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MTKStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTKStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MTKStakeSession struct {
	Contract     *MTKStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MTKStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MTKStakeCallerSession struct {
	Contract *MTKStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MTKStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MTKStakeTransactorSession struct {
	Contract     *MTKStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MTKStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MTKStakeRaw struct {
	Contract *MTKStake // Generic contract binding to access the raw methods on
}

// MTKStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MTKStakeCallerRaw struct {
	Contract *MTKStakeCaller // Generic read-only contract binding to access the raw methods on
}

// MTKStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MTKStakeTransactorRaw struct {
	Contract *MTKStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMTKStake creates a new instance of MTKStake, bound to a specific deployed contract.
func NewMTKStake(address common.Address, backend bind.ContractBackend) (*MTKStake, error) {
	contract, err := bindMTKStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MTKStake{MTKStakeCaller: MTKStakeCaller{contract: contract}, MTKStakeTransactor: MTKStakeTransactor{contract: contract}, MTKStakeFilterer: MTKStakeFilterer{contract: contract}}, nil
}

// NewMTKStakeCaller creates a new read-only instance of MTKStake, bound to a specific deployed contract.
func NewMTKStakeCaller(address common.Address, caller bind.ContractCaller) (*MTKStakeCaller, error) {
	contract, err := bindMTKStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MTKStakeCaller{contract: contract}, nil
}

// NewMTKStakeTransactor creates a new write-only instance of MTKStake, bound to a specific deployed contract.
func NewMTKStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*MTKStakeTransactor, error) {
	contract, err := bindMTKStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MTKStakeTransactor{contract: contract}, nil
}

// NewMTKStakeFilterer creates a new log filterer instance of MTKStake, bound to a specific deployed contract.
func NewMTKStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*MTKStakeFilterer, error) {
	contract, err := bindMTKStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MTKStakeFilterer{contract: contract}, nil
}

// bindMTKStake binds a generic wrapper to an already deployed contract.
func bindMTKStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MTKStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MTKStake *MTKStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MTKStake.Contract.MTKStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MTKStake *MTKStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTKStake.Contract.MTKStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MTKStake *MTKStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MTKStake.Contract.MTKStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MTKStake *MTKStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MTKStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MTKStake *MTKStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTKStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MTKStake *MTKStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MTKStake.Contract.contract.Transact(opts, method, params...)
}

// Apy is a free data retrieval call binding the contract method 0x1f1accb2.
//
// Solidity: function apy(uint8 ) view returns(uint256)
func (_MTKStake *MTKStakeCaller) Apy(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _MTKStake.contract.Call(opts, &out, "apy", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apy is a free data retrieval call binding the contract method 0x1f1accb2.
//
// Solidity: function apy(uint8 ) view returns(uint256)
func (_MTKStake *MTKStakeSession) Apy(arg0 uint8) (*big.Int, error) {
	return _MTKStake.Contract.Apy(&_MTKStake.CallOpts, arg0)
}

// Apy is a free data retrieval call binding the contract method 0x1f1accb2.
//
// Solidity: function apy(uint8 ) view returns(uint256)
func (_MTKStake *MTKStakeCallerSession) Apy(arg0 uint8) (*big.Int, error) {
	return _MTKStake.Contract.Apy(&_MTKStake.CallOpts, arg0)
}

// Durations is a free data retrieval call binding the contract method 0x0ae355d3.
//
// Solidity: function durations(uint8 ) view returns(uint256)
func (_MTKStake *MTKStakeCaller) Durations(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _MTKStake.contract.Call(opts, &out, "durations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Durations is a free data retrieval call binding the contract method 0x0ae355d3.
//
// Solidity: function durations(uint8 ) view returns(uint256)
func (_MTKStake *MTKStakeSession) Durations(arg0 uint8) (*big.Int, error) {
	return _MTKStake.Contract.Durations(&_MTKStake.CallOpts, arg0)
}

// Durations is a free data retrieval call binding the contract method 0x0ae355d3.
//
// Solidity: function durations(uint8 ) view returns(uint256)
func (_MTKStake *MTKStakeCallerSession) Durations(arg0 uint8) (*big.Int, error) {
	return _MTKStake.Contract.Durations(&_MTKStake.CallOpts, arg0)
}

// GetUserActiveStakes is a free data retrieval call binding the contract method 0xa262ab35.
//
// Solidity: function getUserActiveStakes(address user) view returns((uint256,uint256,uint256,uint256,uint256,bool,uint8)[])
func (_MTKStake *MTKStakeCaller) GetUserActiveStakes(opts *bind.CallOpts, user common.Address) ([]MTKStakeStake, error) {
	var out []interface{}
	err := _MTKStake.contract.Call(opts, &out, "getUserActiveStakes", user)

	if err != nil {
		return *new([]MTKStakeStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]MTKStakeStake)).(*[]MTKStakeStake)

	return out0, err

}

// GetUserActiveStakes is a free data retrieval call binding the contract method 0xa262ab35.
//
// Solidity: function getUserActiveStakes(address user) view returns((uint256,uint256,uint256,uint256,uint256,bool,uint8)[])
func (_MTKStake *MTKStakeSession) GetUserActiveStakes(user common.Address) ([]MTKStakeStake, error) {
	return _MTKStake.Contract.GetUserActiveStakes(&_MTKStake.CallOpts, user)
}

// GetUserActiveStakes is a free data retrieval call binding the contract method 0xa262ab35.
//
// Solidity: function getUserActiveStakes(address user) view returns((uint256,uint256,uint256,uint256,uint256,bool,uint8)[])
func (_MTKStake *MTKStakeCallerSession) GetUserActiveStakes(user common.Address) ([]MTKStakeStake, error) {
	return _MTKStake.Contract.GetUserActiveStakes(&_MTKStake.CallOpts, user)
}

// StakeIdOwner is a free data retrieval call binding the contract method 0xeb992847.
//
// Solidity: function stakeIdOwner(uint256 ) view returns(address)
func (_MTKStake *MTKStakeCaller) StakeIdOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MTKStake.contract.Call(opts, &out, "stakeIdOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeIdOwner is a free data retrieval call binding the contract method 0xeb992847.
//
// Solidity: function stakeIdOwner(uint256 ) view returns(address)
func (_MTKStake *MTKStakeSession) StakeIdOwner(arg0 *big.Int) (common.Address, error) {
	return _MTKStake.Contract.StakeIdOwner(&_MTKStake.CallOpts, arg0)
}

// StakeIdOwner is a free data retrieval call binding the contract method 0xeb992847.
//
// Solidity: function stakeIdOwner(uint256 ) view returns(address)
func (_MTKStake *MTKStakeCallerSession) StakeIdOwner(arg0 *big.Int) (common.Address, error) {
	return _MTKStake.Contract.StakeIdOwner(&_MTKStake.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_MTKStake *MTKStakeCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MTKStake.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_MTKStake *MTKStakeSession) StakingToken() (common.Address, error) {
	return _MTKStake.Contract.StakingToken(&_MTKStake.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_MTKStake *MTKStakeCallerSession) StakingToken() (common.Address, error) {
	return _MTKStake.Contract.StakingToken(&_MTKStake.CallOpts)
}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256 stakeId, uint256 amount, uint256 startTime, uint256 endTime, uint256 rewardRateYear, bool isActive, uint8 period)
func (_MTKStake *MTKStakeCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeId        *big.Int
	Amount         *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	RewardRateYear *big.Int
	IsActive       bool
	Period         uint8
}, error) {
	var out []interface{}
	err := _MTKStake.contract.Call(opts, &out, "userStakes", arg0, arg1)

	outstruct := new(struct {
		StakeId        *big.Int
		Amount         *big.Int
		StartTime      *big.Int
		EndTime        *big.Int
		RewardRateYear *big.Int
		IsActive       bool
		Period         uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RewardRateYear = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Period = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256 stakeId, uint256 amount, uint256 startTime, uint256 endTime, uint256 rewardRateYear, bool isActive, uint8 period)
func (_MTKStake *MTKStakeSession) UserStakes(arg0 common.Address, arg1 *big.Int) (struct {
	StakeId        *big.Int
	Amount         *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	RewardRateYear *big.Int
	IsActive       bool
	Period         uint8
}, error) {
	return _MTKStake.Contract.UserStakes(&_MTKStake.CallOpts, arg0, arg1)
}

// UserStakes is a free data retrieval call binding the contract method 0xb5d5b5fa.
//
// Solidity: function userStakes(address , uint256 ) view returns(uint256 stakeId, uint256 amount, uint256 startTime, uint256 endTime, uint256 rewardRateYear, bool isActive, uint8 period)
func (_MTKStake *MTKStakeCallerSession) UserStakes(arg0 common.Address, arg1 *big.Int) (struct {
	StakeId        *big.Int
	Amount         *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	RewardRateYear *big.Int
	IsActive       bool
	Period         uint8
}, error) {
	return _MTKStake.Contract.UserStakes(&_MTKStake.CallOpts, arg0, arg1)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period) returns()
func (_MTKStake *MTKStakeTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, period uint8) (*types.Transaction, error) {
	return _MTKStake.contract.Transact(opts, "stake", amount, period)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period) returns()
func (_MTKStake *MTKStakeSession) Stake(amount *big.Int, period uint8) (*types.Transaction, error) {
	return _MTKStake.Contract.Stake(&_MTKStake.TransactOpts, amount, period)
}

// Stake is a paid mutator transaction binding the contract method 0x10087fb1.
//
// Solidity: function stake(uint256 amount, uint8 period) returns()
func (_MTKStake *MTKStakeTransactorSession) Stake(amount *big.Int, period uint8) (*types.Transaction, error) {
	return _MTKStake.Contract.Stake(&_MTKStake.TransactOpts, amount, period)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId) returns()
func (_MTKStake *MTKStakeTransactor) Withdraw(opts *bind.TransactOpts, stakeId *big.Int) (*types.Transaction, error) {
	return _MTKStake.contract.Transact(opts, "withdraw", stakeId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId) returns()
func (_MTKStake *MTKStakeSession) Withdraw(stakeId *big.Int) (*types.Transaction, error) {
	return _MTKStake.Contract.Withdraw(&_MTKStake.TransactOpts, stakeId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 stakeId) returns()
func (_MTKStake *MTKStakeTransactorSession) Withdraw(stakeId *big.Int) (*types.Transaction, error) {
	return _MTKStake.Contract.Withdraw(&_MTKStake.TransactOpts, stakeId)
}

// MTKStakeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the MTKStake contract.
type MTKStakeStakedIterator struct {
	Event *MTKStakeStaked // Event containing the contract specifics and raw log

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
func (it *MTKStakeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTKStakeStaked)
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
		it.Event = new(MTKStakeStaked)
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
func (it *MTKStakeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTKStakeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTKStakeStaked represents a Staked event raised by the MTKStake contract.
type MTKStakeStaked struct {
	User      common.Address
	StakeId   *big.Int
	Amount    *big.Int
	Period    uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xcc10169be2ad544347561e230939849af48d1714c052d7fe247d12f3decb4896.
//
// Solidity: event Staked(address indexed user, uint256 stakeId, uint256 amount, uint8 period, uint256 timestamp)
func (_MTKStake *MTKStakeFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*MTKStakeStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MTKStake.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &MTKStakeStakedIterator{contract: _MTKStake.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xcc10169be2ad544347561e230939849af48d1714c052d7fe247d12f3decb4896.
//
// Solidity: event Staked(address indexed user, uint256 stakeId, uint256 amount, uint8 period, uint256 timestamp)
func (_MTKStake *MTKStakeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *MTKStakeStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MTKStake.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTKStakeStaked)
				if err := _MTKStake.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xcc10169be2ad544347561e230939849af48d1714c052d7fe247d12f3decb4896.
//
// Solidity: event Staked(address indexed user, uint256 stakeId, uint256 amount, uint8 period, uint256 timestamp)
func (_MTKStake *MTKStakeFilterer) ParseStaked(log types.Log) (*MTKStakeStaked, error) {
	event := new(MTKStakeStaked)
	if err := _MTKStake.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTKStakeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MTKStake contract.
type MTKStakeWithdrawIterator struct {
	Event *MTKStakeWithdraw // Event containing the contract specifics and raw log

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
func (it *MTKStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTKStakeWithdraw)
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
		it.Event = new(MTKStakeWithdraw)
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
func (it *MTKStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTKStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTKStakeWithdraw represents a Withdraw event raised by the MTKStake contract.
type MTKStakeWithdraw struct {
	User        common.Address
	StakeId     *big.Int
	Principal   *big.Int
	Reward      *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe08737ac48a1dab4b1a46c7dc9398bd5bfc6d7ad6fabb7cd8caa254de14def35.
//
// Solidity: event Withdraw(address indexed user, uint256 stakeId, uint256 principal, uint256 reward, uint256 totalAmount)
func (_MTKStake *MTKStakeFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*MTKStakeWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MTKStake.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &MTKStakeWithdrawIterator{contract: _MTKStake.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe08737ac48a1dab4b1a46c7dc9398bd5bfc6d7ad6fabb7cd8caa254de14def35.
//
// Solidity: event Withdraw(address indexed user, uint256 stakeId, uint256 principal, uint256 reward, uint256 totalAmount)
func (_MTKStake *MTKStakeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MTKStakeWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MTKStake.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTKStakeWithdraw)
				if err := _MTKStake.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe08737ac48a1dab4b1a46c7dc9398bd5bfc6d7ad6fabb7cd8caa254de14def35.
//
// Solidity: event Withdraw(address indexed user, uint256 stakeId, uint256 principal, uint256 reward, uint256 totalAmount)
func (_MTKStake *MTKStakeFilterer) ParseWithdraw(log types.Log) (*MTKStakeWithdraw, error) {
	event := new(MTKStakeWithdraw)
	if err := _MTKStake.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
