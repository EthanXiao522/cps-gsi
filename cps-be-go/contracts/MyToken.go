// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// MyTokenMetaData contains all meta data concerning the MyToken contract.
var MyTokenMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
	ID:  "8165024181563831e8d57e825005ea4137",
	Bin: "0x0x60c060405260036080908152624d544b60e81b60a052600190610022908261016c565b506002805460ff1916601217905534801561003b575f5ffd5b50604051610bf3380380610bf383398101604081905261005a91610226565b5f610065828261016c565b506002546100779060ff16600a6103cf565b6100829060786103e1565b6003819055335f81815260046020908152604080832085905551938452919290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3506103f8565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806100fc57607f821691505b60208210810361011a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561016757805f5260205f20601f840160051c810160208510156101455750805b601f840160051c820191505b81811015610164575f8155600101610151565b50505b505050565b81516001600160401b03811115610185576101856100d4565b6101998161019384546100e8565b84610120565b6020601f8211600181146101cb575f83156101b45750848201515b5f19600385901b1c1916600184901b178455610164565b5f84815260208120601f198516915b828110156101fa57878501518255602094850194600190920191016101da565b508482101561021757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215610236575f5ffd5b81516001600160401b0381111561024b575f5ffd5b8201601f8101841361025b575f5ffd5b80516001600160401b03811115610274576102746100d4565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102a2576102a26100d4565b6040528181528282016020018610156102b9575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b634e487b7160e01b5f52601160045260245ffd5b6001815b600184111561032557808504811115610309576103096102d6565b600184161561031757908102905b60019390931c9280026102ee565b935093915050565b5f8261033b575060016103c9565b8161034757505f6103c9565b816001811461035d576002811461036757610383565b60019150506103c9565b60ff841115610378576103786102d6565b50506001821b6103c9565b5060208310610133831016604e8410600b84101617156103a6575081810a6103c9565b6103b25f1984846102ea565b805f19048211156103c5576103c56102d6565b0290505b92915050565b5f6103da838361032d565b9392505050565b80820281158282048414176103c9576103c96102d6565b6107ee806104055f395ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c8063313ce56711610063578063313ce567146100ff57806370a082311461011e57806395d89b411461013d578063a9059cbb14610145578063dd62ed3e14610158575f5ffd5b806306fdde0314610094578063095ea7b3146100b257806318160ddd146100d557806323b872dd146100ec575b5f5ffd5b61009c610182565b6040516100a991906105ba565b60405180910390f35b6100c56100c036600461060a565b61020d565b60405190151581526020016100a9565b6100de60035481565b6040519081526020016100a9565b6100c56100fa366004610632565b6102d8565b60025461010c9060ff1681565b60405160ff90911681526020016100a9565b6100de61012c36600461066c565b60046020525f908152604090205481565b61009c6104cc565b6100c561015336600461060a565b6104d9565b6100de61016636600461068c565b600560209081525f928352604080842090915290825290205481565b5f805461018e906106bd565b80601f01602080910402602001604051908101604052809291908181526020018280546101ba906106bd565b80156102055780601f106101dc57610100808354040283529160200191610205565b820191905f5260205f20905b8154815290600101906020018083116101e857829003601f168201915b505050505081565b5f6001600160a01b0383166102745760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084015b60405180910390fd5b335f8181526005602090815260408083206001600160a01b03881680855290835292819020869055518581529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a35060015b92915050565b5f6001600160a01b0383166102ff5760405162461bcd60e51b815260040161026b906106f5565b6001600160a01b0384165f90815260046020526040902054828110156103375760405162461bcd60e51b815260040161026b90610738565b6001600160a01b0385165f908152600560209081526040808320338452909152902054838110156103bb5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b606482015260840161026b565b6103c58483610792565b6001600160a01b038088165f9081526004602052604080822093909355908716815290812080548692906103fa9084906107a5565b9091555061040a90508482610792565b6001600160a01b038781165f81815260056020908152604080832033845282529182902094909455518781529188169290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36001600160a01b0386165f81815260056020908152604080832033808552908352928190205490519081529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350600195945050505050565b6001805461018e906106bd565b5f6001600160a01b0383166105005760405162461bcd60e51b815260040161026b906106f5565b335f908152600460205260409020548281101561052f5760405162461bcd60e51b815260040161026b90610738565b6105398382610792565b335f90815260046020526040808220929092556001600160a01b0386168152908120805485929061056b9084906107a5565b90915550506040518381526001600160a01b0385169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35060019392505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610605575f5ffd5b919050565b5f5f6040838503121561061b575f5ffd5b610624836105ef565b946020939093013593505050565b5f5f5f60608486031215610644575f5ffd5b61064d846105ef565b925061065b602085016105ef565b929592945050506040919091013590565b5f6020828403121561067c575f5ffd5b610685826105ef565b9392505050565b5f5f6040838503121561069d575f5ffd5b6106a6836105ef565b91506106b4602084016105ef565b90509250929050565b600181811c908216806106d157607f821691505b6020821081036106ef57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526026908201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604082015265616c616e636560d01b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b818103818111156102d2576102d261077e565b808201808211156102d2576102d261077e56fea2646970667358221220dde49aaba4e75bd04cc044c0c61b7b460de50c539645b7435f36a35de0b6792c64736f6c634300081c0033",
}

// MyToken is an auto generated Go binding around an Ethereum contract.
type MyToken struct {
	abi abi.ABI
}

// NewMyToken creates a new instance of MyToken.
func NewMyToken() *MyToken {
	parsed, err := MyTokenMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MyToken{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MyToken) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(string _name) returns()
func (myToken *MyToken) PackConstructor(_name string) []byte {
	enc, err := myToken.abi.Pack("", _name)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (myToken *MyToken) PackAllowance(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := myToken.abi.Pack("allowance", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (myToken *MyToken) TryPackAllowance(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return myToken.abi.Pack("allowance", arg0, arg1)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (myToken *MyToken) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := myToken.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (myToken *MyToken) PackApprove(spender common.Address, amount *big.Int) []byte {
	enc, err := myToken.abi.Pack("approve", spender, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (myToken *MyToken) TryPackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	return myToken.abi.Pack("approve", spender, amount)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (myToken *MyToken) UnpackApprove(data []byte) (bool, error) {
	out, err := myToken.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (myToken *MyToken) PackBalanceOf(arg0 common.Address) []byte {
	enc, err := myToken.abi.Pack("balanceOf", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (myToken *MyToken) TryPackBalanceOf(arg0 common.Address) ([]byte, error) {
	return myToken.abi.Pack("balanceOf", arg0)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (myToken *MyToken) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := myToken.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decimals() view returns(uint8)
func (myToken *MyToken) PackDecimals() []byte {
	enc, err := myToken.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decimals() view returns(uint8)
func (myToken *MyToken) TryPackDecimals() ([]byte, error) {
	return myToken.abi.Pack("decimals")
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (myToken *MyToken) UnpackDecimals(data []byte) (uint8, error) {
	out, err := myToken.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (myToken *MyToken) PackName() []byte {
	enc, err := myToken.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (myToken *MyToken) TryPackName() ([]byte, error) {
	return myToken.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (myToken *MyToken) UnpackName(data []byte) (string, error) {
	out, err := myToken.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function symbol() view returns(string)
func (myToken *MyToken) PackSymbol() []byte {
	enc, err := myToken.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function symbol() view returns(string)
func (myToken *MyToken) TryPackSymbol() ([]byte, error) {
	return myToken.abi.Pack("symbol")
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (myToken *MyToken) UnpackSymbol(data []byte) (string, error) {
	out, err := myToken.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (myToken *MyToken) PackTotalSupply() []byte {
	enc, err := myToken.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (myToken *MyToken) TryPackTotalSupply() ([]byte, error) {
	return myToken.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (myToken *MyToken) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := myToken.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (myToken *MyToken) PackTransfer(to common.Address, amount *big.Int) []byte {
	enc, err := myToken.abi.Pack("transfer", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (myToken *MyToken) TryPackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	return myToken.abi.Pack("transfer", to, amount)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (myToken *MyToken) UnpackTransfer(data []byte) (bool, error) {
	out, err := myToken.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (myToken *MyToken) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) []byte {
	enc, err := myToken.abi.Pack("transferFrom", from, to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (myToken *MyToken) TryPackTransferFrom(from common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	return myToken.abi.Pack("transferFrom", from, to, amount)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (myToken *MyToken) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := myToken.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// MyTokenApproval represents a Approval event raised by the MyToken contract.
type MyTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const MyTokenApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (MyTokenApproval) ContractEventName() string {
	return MyTokenApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (myToken *MyToken) UnpackApprovalEvent(log *types.Log) (*MyTokenApproval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != myToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyTokenApproval)
	if len(log.Data) > 0 {
		if err := myToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myToken.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// MyTokenTransfer represents a Transfer event raised by the MyToken contract.
type MyTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const MyTokenTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (MyTokenTransfer) ContractEventName() string {
	return MyTokenTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (myToken *MyToken) UnpackTransferEvent(log *types.Log) (*MyTokenTransfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != myToken.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyTokenTransfer)
	if len(log.Data) > 0 {
		if err := myToken.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myToken.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
