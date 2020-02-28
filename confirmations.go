// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gost

import (
	"math/big"
	"strings"

	"github.com/gochain/gochain/v3"
	"github.com/gochain/gochain/v3/accounts/abi"
	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/core/types"
	"github.com/gochain/gochain/v3/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = gochain.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressSetABI is the input ABI used to generate the binding from.
const AddressSetABI = "[]"

// AddressSetFuncSigs maps the 4-byte function signature to its string representation.
var AddressSetFuncSigs = map[string]string{
	"e207783c": "majority(AddressSet.Set storage)",
}

// AddressSetBin is the compiled bytecode used for deploying new contracts.
const AddressSetBin = "0x60a4610024600b82828239805160001a607314601757fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063e207783c146038575b600080fd5b605260048036036020811015604c57600080fd5b50356064565b60408051918252519081900360200190f35b54600290046001019056fea265627a7a7231582081dfa07501e6d34e5edde93510f6808a70476bbdb4c818ccf5dbdfdd219b202664736f6c63430005100032"

// DeployAddressSet deploys a new GoChain contract, binding an instance of AddressSet to it.
func DeployAddressSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressSet, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	addressSetBin := AddressSetBin

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(addressSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressSet{AddressSetCaller: AddressSetCaller{contract: contract}, AddressSetTransactor: AddressSetTransactor{contract: contract}, AddressSetFilterer: AddressSetFilterer{contract: contract}}, nil
}

// AddressSet is an auto generated Go binding around an GoChain contract.
type AddressSet struct {
	AddressSetCaller     // Read-only binding to the contract
	AddressSetTransactor // Write-only binding to the contract
	AddressSetFilterer   // Log filterer for contract events
}

// AddressSetCaller is an auto generated read-only Go binding around an GoChain contract.
type AddressSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSetTransactor is an auto generated write-only Go binding around an GoChain contract.
type AddressSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSetFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type AddressSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSetSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type AddressSetSession struct {
	Contract     *AddressSet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressSetCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type AddressSetCallerSession struct {
	Contract *AddressSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AddressSetTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type AddressSetTransactorSession struct {
	Contract     *AddressSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddressSetRaw is an auto generated low-level Go binding around an GoChain contract.
type AddressSetRaw struct {
	Contract *AddressSet // Generic contract binding to access the raw methods on
}

// AddressSetCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type AddressSetCallerRaw struct {
	Contract *AddressSetCaller // Generic read-only contract binding to access the raw methods on
}

// AddressSetTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type AddressSetTransactorRaw struct {
	Contract *AddressSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressSet creates a new instance of AddressSet, bound to a specific deployed contract.
func NewAddressSet(address common.Address, backend bind.ContractBackend) (*AddressSet, error) {
	contract, err := bindAddressSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressSet{AddressSetCaller: AddressSetCaller{contract: contract}, AddressSetTransactor: AddressSetTransactor{contract: contract}, AddressSetFilterer: AddressSetFilterer{contract: contract}}, nil
}

// NewAddressSetCaller creates a new read-only instance of AddressSet, bound to a specific deployed contract.
func NewAddressSetCaller(address common.Address, caller bind.ContractCaller) (*AddressSetCaller, error) {
	contract, err := bindAddressSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressSetCaller{contract: contract}, nil
}

// NewAddressSetTransactor creates a new write-only instance of AddressSet, bound to a specific deployed contract.
func NewAddressSetTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressSetTransactor, error) {
	contract, err := bindAddressSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressSetTransactor{contract: contract}, nil
}

// NewAddressSetFilterer creates a new log filterer instance of AddressSet, bound to a specific deployed contract.
func NewAddressSetFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressSetFilterer, error) {
	contract, err := bindAddressSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressSetFilterer{contract: contract}, nil
}

// bindAddressSet binds a generic wrapper to an already deployed contract.
func bindAddressSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressSet *AddressSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressSet.Contract.AddressSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressSet *AddressSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressSet.Contract.AddressSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressSet *AddressSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressSet.Contract.AddressSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressSet *AddressSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressSet *AddressSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressSet *AddressSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressSet.Contract.contract.Transact(opts, method, params...)
}

// AuthABI is the input ABI used to generate the binding from.
const AuthABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AuthFuncSigs maps the 4-byte function signature to its string representation.
var AuthFuncSigs = map[string]string{
	"7849ed5a": "count(address,bool,bool)",
	"3ffefe4e": "getSigner(uint256)",
	"d07bff0c": "getVoter(uint256)",
	"7df73e27": "isSigner(address)",
	"a7771ee3": "isVoter(address)",
	"48c5000d": "setVote(address,bool,bool)",
	"41f684f3": "signersLength()",
	"6c6c39fb": "votersLength()",
	"d8bff5a5": "votes(address)",
}

// Auth is an auto generated Go binding around an GoChain contract.
type Auth struct {
	AuthCaller     // Read-only binding to the contract
	AuthTransactor // Write-only binding to the contract
	AuthFilterer   // Log filterer for contract events
}

// AuthCaller is an auto generated read-only Go binding around an GoChain contract.
type AuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthTransactor is an auto generated write-only Go binding around an GoChain contract.
type AuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type AuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type AuthSession struct {
	Contract     *Auth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type AuthCallerSession struct {
	Contract *AuthCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AuthTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type AuthTransactorSession struct {
	Contract     *AuthTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthRaw is an auto generated low-level Go binding around an GoChain contract.
type AuthRaw struct {
	Contract *Auth // Generic contract binding to access the raw methods on
}

// AuthCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type AuthCallerRaw struct {
	Contract *AuthCaller // Generic read-only contract binding to access the raw methods on
}

// AuthTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type AuthTransactorRaw struct {
	Contract *AuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuth creates a new instance of Auth, bound to a specific deployed contract.
func NewAuth(address common.Address, backend bind.ContractBackend) (*Auth, error) {
	contract, err := bindAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// NewAuthCaller creates a new read-only instance of Auth, bound to a specific deployed contract.
func NewAuthCaller(address common.Address, caller bind.ContractCaller) (*AuthCaller, error) {
	contract, err := bindAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthCaller{contract: contract}, nil
}

// NewAuthTransactor creates a new write-only instance of Auth, bound to a specific deployed contract.
func NewAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthTransactor, error) {
	contract, err := bindAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthTransactor{contract: contract}, nil
}

// NewAuthFilterer creates a new log filterer instance of Auth, bound to a specific deployed contract.
func NewAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthFilterer, error) {
	contract, err := bindAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthFilterer{contract: contract}, nil
}

// bindAuth binds a generic wrapper to an already deployed contract.
func bindAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.AuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Auth *AuthCaller) Count(opts *bind.CallOpts, arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "count", arg0, arg1, arg2)
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Auth *AuthSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Auth.Contract.Count(&_Auth.CallOpts, arg0, arg1, arg2)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Auth *AuthCallerSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Auth.Contract.Count(&_Auth.CallOpts, arg0, arg1, arg2)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Auth *AuthCaller) GetSigner(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "getSigner", i)
	return *ret0, err
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Auth *AuthSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Auth.Contract.GetSigner(&_Auth.CallOpts, i)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Auth *AuthCallerSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Auth.Contract.GetSigner(&_Auth.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Auth *AuthCaller) GetVoter(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "getVoter", i)
	return *ret0, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Auth *AuthSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Auth.Contract.GetVoter(&_Auth.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Auth *AuthCallerSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Auth.Contract.GetVoter(&_Auth.CallOpts, i)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Auth *AuthCaller) IsSigner(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "isSigner", voter)
	return *ret0, err
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Auth *AuthSession) IsSigner(voter common.Address) (bool, error) {
	return _Auth.Contract.IsSigner(&_Auth.CallOpts, voter)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Auth *AuthCallerSession) IsSigner(voter common.Address) (bool, error) {
	return _Auth.Contract.IsSigner(&_Auth.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Auth *AuthCaller) IsVoter(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "isVoter", voter)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Auth *AuthSession) IsVoter(voter common.Address) (bool, error) {
	return _Auth.Contract.IsVoter(&_Auth.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Auth *AuthCallerSession) IsVoter(voter common.Address) (bool, error) {
	return _Auth.Contract.IsVoter(&_Auth.CallOpts, voter)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Auth *AuthCaller) SignersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "signersLength")
	return *ret0, err
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Auth *AuthSession) SignersLength() (*big.Int, error) {
	return _Auth.Contract.SignersLength(&_Auth.CallOpts)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Auth *AuthCallerSession) SignersLength() (*big.Int, error) {
	return _Auth.Contract.SignersLength(&_Auth.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Auth *AuthCaller) VotersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Auth.contract.Call(opts, out, "votersLength")
	return *ret0, err
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Auth *AuthSession) VotersLength() (*big.Int, error) {
	return _Auth.Contract.VotersLength(&_Auth.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Auth *AuthCallerSession) VotersLength() (*big.Int, error) {
	return _Auth.Contract.VotersLength(&_Auth.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Auth *AuthCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	ret := new(struct {
		Addr  common.Address
		Voter bool
		Add   bool
	})
	out := ret
	err := _Auth.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Auth *AuthSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Auth.Contract.Votes(&_Auth.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Auth *AuthCallerSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Auth.Contract.Votes(&_Auth.CallOpts, arg0)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Auth *AuthTransactor) SetVote(opts *bind.TransactOpts, addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Auth.contract.Transact(opts, "setVote", addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Auth *AuthSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Auth.Contract.SetVote(&_Auth.TransactOpts, addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Auth *AuthTransactorSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Auth.Contract.SetVote(&_Auth.TransactOpts, addr, voter, add)
}

// ConfirmationsABI is the input ABI used to generate the binding from.
const ConfirmationsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"Confirmed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_confirmedGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"confirmGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"request\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"shouldConfirm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumIConfirmations.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalConfirmGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"voter\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ConfirmationsFuncSigs maps the 4-byte function signature to its string representation.
var ConfirmationsFuncSigs = map[string]string{
	"2ea6ece8": "_confirmedGas()",
	"8f7fd0d6": "confirm(uint256,uint256,bytes32,bool)",
	"42715aec": "confirmGas()",
	"7849ed5a": "count(address,bool,bool)",
	"3ffefe4e": "getSigner(uint256)",
	"d07bff0c": "getVoter(uint256)",
	"7df73e27": "isSigner(address)",
	"a7771ee3": "isVoter(address)",
	"03aca792": "pendingList(uint256)",
	"26b60630": "pendingListLength()",
	"32030803": "request(uint256,uint256,bytes32)",
	"48c5000d": "setVote(address,bool,bool)",
	"859f32c2": "shouldConfirm(uint256,uint256,bytes32)",
	"41f684f3": "signersLength()",
	"fad3ffd6": "status(uint256,uint256,bytes32)",
	"5c5f6b39": "totalConfirmGas()",
	"6c6c39fb": "votersLength()",
	"d8bff5a5": "votes(address)",
}

// ConfirmationsBin is the compiled bytecode used for deploying new contracts.
const ConfirmationsBin = "0x60806040523480156200001157600080fd5b5060405162001aeb38038062001aeb833981810160405260208110156200003757600080fd5b5051806200005360008262000073602090811b6200119f17901c565b6200006b6002826200007360201b6200119f1760201c565b505062000114565b6001600160a01b038116600090815260018301602052604090205415620000d2576040805162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481a5b881cd95d60921b604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b6119c780620001246000396000f3fe6080604052600436106101095760003560e01c80636c6c39fb116100955780638f7fd0d6116100645780638f7fd0d614610375578063a7771ee3146103b3578063d07bff0c146103e6578063d8bff5a514610410578063fad3ffd61461046d57610109565b80636c6c39fb146102855780637849ed5a1461029a5780637df73e27146102dd578063859f32c21461032457610109565b80633ffefe4e116100dc5780633ffefe4e146101bd57806341f684f31461020357806342715aec1461021857806348c5000d1461022d5780635c5f6b391461027057610109565b806303aca7921461010e57806326b60630146101565780632ea6ece81461017d5780633203080314610192575b600080fd5b34801561011a57600080fd5b506101386004803603602081101561013157600080fd5b50356104c7565b60408051938452602084019290925282820152519081900360600190f35b34801561016257600080fd5b5061016b6104f7565b60408051918252519081900360200190f35b34801561018957600080fd5b5061016b6104fe565b6101bb600480360360608110156101a857600080fd5b5080359060208101359060400135610505565b005b3480156101c957600080fd5b506101e7600480360360208110156101e057600080fd5b503561074a565b604080516001600160a01b039092168252519081900360200190f35b34801561020f57600080fd5b5061016b610777565b34801561022457600080fd5b5061016b61077d565b34801561023957600080fd5b506101bb6004803603606081101561025057600080fd5b506001600160a01b03813516906020810135151590604001351515610784565b34801561027c57600080fd5b5061016b610a65565b34801561029157600080fd5b5061016b610a6c565b3480156102a657600080fd5b5061016b600480360360608110156102bd57600080fd5b506001600160a01b03813516906020810135151590604001351515610a72565b3480156102e957600080fd5b506103106004803603602081101561030057600080fd5b50356001600160a01b0316610a95565b604080519115158252519081900360200190f35b34801561033057600080fd5b5061035a6004803603606081101561034757600080fd5b5080359060208101359060400135610ab2565b60408051921515835260208301919091528051918290030190f35b34801561038157600080fd5b506101bb6004803603608081101561039857600080fd5b50803590602081013590604081013590606001351515610bc5565b3480156103bf57600080fd5b50610310600480360360208110156103d657600080fd5b50356001600160a01b0316610dd2565b3480156103f257600080fd5b506101e76004803603602081101561040957600080fd5b5035610def565b34801561041c57600080fd5b506104436004803603602081101561043357600080fd5b50356001600160a01b0316610e00565b604080516001600160a01b0390941684529115156020840152151582820152519081900360600190f35b34801561047957600080fd5b506104a36004803603606081101561049057600080fd5b5080359060208101359060400135610e31565b604051808260038111156104b357fe5b60ff16815260200191505060405180910390f35b600881815481106104d457fe5b600091825260209091206003909102018054600182015460029092015490925083565b6008545b90565b6202710081565b6000838152600660209081526040808320858452825280832084845290915281205460ff16600381111561053557fe5b1461057d576040805162461bcd60e51b8152602060048201526013602482015272537461747573206d757374206265204e6f6e6560681b604482015290519081900360640190fd5b600083815260066020908152604080832085845282528083208484529091529020805460ff191660011790553a622625a002348111156105ee5760405162461bcd60e51b815260040180806020018281038252602781526020018061196c6027913960400191505060405180910390fd5b803411156106275760405133903483900380156108fc02916000818181858888f19350505050158015610625573d6000803e3d6000fd5b505b6000848152600760209081526040808320868452825280832085845282528083203a6004909101558051606081018252878152808301878152818301878152600880546001810180835591885293517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee360039095029485015591517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee4840155517ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee59092019190915587845260098352818420878552835281842086855283529281902083905580518581529051869288927f2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03929081900390910190a35050505050565b60006002600001828154811061075c57fe5b6000918252602090912001546001600160a01b031692915050565b60025490565b62015f9081565b336000908152600160205260409020546107da576040805162461bcd60e51b815260206004820152601260248201527129b2b73232b9103737ba1030903b37ba32b960711b604482015290519081900360640190fd5b33600090815260046020526040902080546001600160a01b0316156108ae5780546001600160a01b0385811691161480156108245750805460ff600160a01b909104161515831515145b801561083f5750805460ff600160a81b909104161515821515145b1561084a5750610a60565b80546001600160a01b038116600090815260056020908152604080832060ff600160a01b8604811615158552908352818420600160a81b90950416151583529281528282208054600019019055338252600490522080546001600160b01b03191690555b6001600160a01b0384166108c25750610a60565b82156108ff5781156108e7576108d784610dd2565b156108e25750610a60565b6108fa565b6108f084610dd2565b6108fa5750610a60565b61092c565b81156109195761090e84610a95565b156108fa5750610a60565b61092284610a95565b61092c5750610a60565b604080516060810182526001600160a01b03808716808352861515602080850182815288151586880181815233600090815260048086528a822099518a54955193511515600160a81b0260ff60a81b19941515600160a01b0260ff60a01b1992909b166001600160a01b03199097169690961716989098179190911692909217909655928352600581528583209183529081528482209382529283528381208054600101908190558451633881de0f60e21b815292830191909152925173__$5c8db3f5ca5c76e13c1ce5ed01517fc035$__9263e207783c9260248082019391829003018186803b158015610a2057600080fd5b505af4158015610a34573d6000803e3d6000fd5b505050506040513d6020811015610a4a57600080fd5b50518110610a5d57610a5d858585610e57565b50505b505050565b622625a081565b60005490565b600560209081526000938452604080852082529284528284209052825290205481565b6001600160a01b0316600090815260036020526040902054151590565b336000908152600360205260408120548190610b0b576040805162461bcd60e51b815260206004820152601360248201527229b2b73232b9103737ba10309039b4b3b732b960691b604482015290519081900360640190fd5b60016000868152600660209081526040808320888452825280832087845290915290205460ff166003811115610b3d57fe5b14610b4d57506000905080610bbd565b600085815260076020908152604080832087845282528083208684528252808320338452600181019092529091205415610b8e575060009150819050610bbd565b33600090815260038201602052604090205415610bb2575060009150819050610bbd565b600401546001925090505b935093915050565b33600090815260036020526040902054610c1c576040805162461bcd60e51b815260206004820152601360248201527229b2b73232b9103737ba10309039b4b3b732b960691b604482015290519081900360640190fd5b60016000858152600660209081526040808320878452825280832086845290915290205460ff166003811115610c4e57fe5b14610c5857610dcc565b60008481526007602090815260408083208684528252808320858452909152902060048101543a14610cbb5760405162461bcd60e51b815260040180806020018281038252602281526020018061194a6022913960400191505060405180910390fd5b808215610ceb5733600090815260038301602052604090205415610ce657610ce68260020133611086565b610d12565b50336000908152600182016020526040902054600282019015610d1257610d128233611086565b336000908152600182016020526040902054610dc957610d32813361119f565b73__$5c8db3f5ca5c76e13c1ce5ed01517fc035$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610d8257600080fd5b505af4158015610d96573d6000803e3d6000fd5b505050506040513d6020811015610dac57600080fd5b505181541015610dbd575050610dcc565b610dc98686868661123f565b50505b50505050565b6001600160a01b0316600090815260016020526040902054151590565b600080600001828154811061075c57fe5b6004602052600090815260409020546001600160a01b0381169060ff600160a01b8204811691600160a81b90041683565b600660209081526000938452604080852082529284528284209052825290205460ff1681565b60005b600054811015610f14576000806000018281548110610e7557fe5b60009182526020808320909101546001600160a01b0390811680845260049092526040909220805491935091908116908716148015610ec35750805460ff600160a01b909104161515851515145b8015610ede5750805460ff600160a81b909104161515841515145b15610f0a576001600160a01b038216600090815260046020526040902080546001600160b01b03191690555b5050600101610e5a565b506001600160a01b038316600090815260056020908152604080832085158015855290835281842085151585529092528220919091558290610f535750805b15610fad576001600160a01b038316600090815260036020526040902054610f8057610f8060028461119f565b6001600160a01b038316600090815260016020526040902054610fa857610fa860008461119f565b610a60565b818015610fb8575080155b15610fe4576001600160a01b03831660009081526001602052604090205415610fa857610fa88361151d565b81158015610fef5750805b1561101c576001600160a01b038316600090815260036020526040902054610fa857610fa860028461119f565b81158015611028575080155b15610a60576001600160a01b03831660009081526001602052604090205415611054576110548361151d565b6001600160a01b03831660009081526003602052604090205415610a605761107d600284611086565b610a608361170e565b6001600160a01b0381166000908152600183016020526040902054806110e0576040805162461bcd60e51b815260206004820152600a602482015269139bdd081a5b881cd95d60b21b604482015290519081900360640190fd5b6001600160a01b0382166000908152600184016020526040812055825483600019820182811061110c57fe5b60009182526020909120015484546001600160a01b03909116908590600019850190811061113657fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555083600001600182038154811061117657fe5b600091825260209091200180546001600160a01b03191690558354610a5d856000198301611896565b6001600160a01b0381166000908152600183016020526040902054156111fd576040805162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481a5b881cd95d60921b604482015290519081900360640190fd5b815460018181018085556000858152602080822090940180546001600160a01b039096166001600160a01b031990961686179055938452930190526040902055565b6000811561124f57506003611253565b5060025b600085815260066020908152604080832087845282528083208684529091529020805482919060ff1916600183600381111561128b57fe5b021790555060008581526007602090815260408083208784528252808320868452909152902080836112bd5750600281015b805460048301546223b4a00260008282816112d457fe5b04905060005b838110156113405760008560000182815481106112f357fe5b60009182526020822001546040516001600160a01b039091169250829185156108fc02918691818181858888f19350505050158015611336573d6000803e3d6000fd5b50506001016112da565b506004850154604051848302840391620271000290339082840180156108fc02916000818181858888f19350505050158015611380573d6000803e3d6000fd5b5060008c81526007602090815260408083208e845282528083208d845290915281209081816113af82826118ba565b50506002820160006113c182826118ba565b50506000600492909201829055508c81526009602090815260408083208e845282528083208d84529091528120805491905560085460018111156114875760006008600183038154811061141157fe5b90600052602060002090600302019050806008600185038154811061143257fe5b6000918252602080832084546003909302019182556001808501548184015560029485015492850192909255845483526009815260408084209286015484529181528183209490930154825292909152208290555b6008600182038154811061149757fe5b600091825260208220600390910201818155600181018290556002015560088054906114c79060001983016118d8565b508c8e7fdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c8e8e60405180838152602001821515151581526020019250505060405180910390a35050505050505050505050505050565b611528600082611086565b6001600160a01b0380821660009081526004602052604090208054909116156115b35780546001600160a01b03808216600090815260056020908152604080832060ff600160a01b8704811615158552908352818420600160a81b90960416151583529381528382208054600019019055918516815260049091522080546001600160b01b03191690555b600073__$5c8db3f5ca5c76e13c1ce5ed01517fc035$__63e207783c60006040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561160557600080fd5b505af4158015611619573d6000803e3d6000fd5b505050506040513d602081101561162f57600080fd5b5051905060005b600054811015610dcc576004600080600001838154811061165357fe5b60009182526020808320909101546001600160a01b0390811684529083019390935260409091019020805490945016156117035782546001600160a01b038116600090815260056020908152604080832060ff600160a01b8604811615158552908352818420600160a81b9095041615158352929052205482116117035782546116fb906001600160a01b0381169060ff600160a01b8204811691600160a81b900416610e57565b50505061170b565b600101611636565b50565b600073__$5c8db3f5ca5c76e13c1ce5ed01517fc035$__63e207783c60026040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561176057600080fd5b505af4158015611774573d6000803e3d6000fd5b505050506040513d602081101561178a57600080fd5b5051905060005b600854811015610a60576000600882815481106117aa57fe5b6000918252602080832060039092029091018054835260078252604080842060018084015486529084528185206002840154865284528185206001600160a01b038a1686529081019093529092205491925090156118115761180c8186611086565b61183e565b6001600160a01b03851660009081526003820160205260409020541561183e5761183e8160020186611086565b805484116118645761185f826000015483600101548460020154600161123f565b61188f565b600281015484116118885761185f826000015483600101548460020154600061123f565b6001909201915b5050611791565b815481835581811115610a6057600083815260209020610a60918101908301611904565b508054600082559060005260206000209081019061170b9190611904565b815481835581811115610a6057600302816003028360005260206000209182019101610a609190611922565b6104fb91905b8082111561191e576000815560010161190a565b5090565b6104fb91905b8082111561191e57600080825560018201819055600282015560030161192856fe47617320707269636520646f65736e2774206d61746368207265717565737465642e54782076616c756520646f65736e277420636f76657220636f6e6669726d6174696f6e20666565a265627a7a72315820597b51da58ea8b45a374e03aa284afe8e983d750357139d4dc5d50cfa7f6961064736f6c63430005100032"

// DeployConfirmations deploys a new GoChain contract, binding an instance of Confirmations to it.
func DeployConfirmations(auth *bind.TransactOpts, backend bind.ContractBackend, voter common.Address) (common.Address, *types.Transaction, *Confirmations, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	confirmationsBin := ConfirmationsBin

	if auth.Nonce == nil {
		if err := bind.EnsureNonce(auth, backend); err != nil {
			return common.Address{}, nil, nil, err
		}
	}

	addressSetAddr, _, _, err := DeployAddressSet(auth, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	confirmationsBin = strings.Replace(confirmationsBin, "__$5c8db3f5ca5c76e13c1ce5ed01517fc035$__", addressSetAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(confirmationsBin), backend, voter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Confirmations{ConfirmationsCaller: ConfirmationsCaller{contract: contract}, ConfirmationsTransactor: ConfirmationsTransactor{contract: contract}, ConfirmationsFilterer: ConfirmationsFilterer{contract: contract}}, nil
}

// Confirmations is an auto generated Go binding around an GoChain contract.
type Confirmations struct {
	ConfirmationsCaller     // Read-only binding to the contract
	ConfirmationsTransactor // Write-only binding to the contract
	ConfirmationsFilterer   // Log filterer for contract events
}

// ConfirmationsCaller is an auto generated read-only Go binding around an GoChain contract.
type ConfirmationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsTransactor is an auto generated write-only Go binding around an GoChain contract.
type ConfirmationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type ConfirmationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmationsSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type ConfirmationsSession struct {
	Contract     *Confirmations    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfirmationsCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type ConfirmationsCallerSession struct {
	Contract *ConfirmationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConfirmationsTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type ConfirmationsTransactorSession struct {
	Contract     *ConfirmationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConfirmationsRaw is an auto generated low-level Go binding around an GoChain contract.
type ConfirmationsRaw struct {
	Contract *Confirmations // Generic contract binding to access the raw methods on
}

// ConfirmationsCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type ConfirmationsCallerRaw struct {
	Contract *ConfirmationsCaller // Generic read-only contract binding to access the raw methods on
}

// ConfirmationsTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type ConfirmationsTransactorRaw struct {
	Contract *ConfirmationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfirmations creates a new instance of Confirmations, bound to a specific deployed contract.
func NewConfirmations(address common.Address, backend bind.ContractBackend) (*Confirmations, error) {
	contract, err := bindConfirmations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Confirmations{ConfirmationsCaller: ConfirmationsCaller{contract: contract}, ConfirmationsTransactor: ConfirmationsTransactor{contract: contract}, ConfirmationsFilterer: ConfirmationsFilterer{contract: contract}}, nil
}

// NewConfirmationsCaller creates a new read-only instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsCaller(address common.Address, caller bind.ContractCaller) (*ConfirmationsCaller, error) {
	contract, err := bindConfirmations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsCaller{contract: contract}, nil
}

// NewConfirmationsTransactor creates a new write-only instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfirmationsTransactor, error) {
	contract, err := bindConfirmations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsTransactor{contract: contract}, nil
}

// NewConfirmationsFilterer creates a new log filterer instance of Confirmations, bound to a specific deployed contract.
func NewConfirmationsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfirmationsFilterer, error) {
	contract, err := bindConfirmations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsFilterer{contract: contract}, nil
}

// bindConfirmations binds a generic wrapper to an already deployed contract.
func bindConfirmations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirmations *ConfirmationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Confirmations.Contract.ConfirmationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirmations *ConfirmationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirmations.Contract.ConfirmationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirmations *ConfirmationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirmations.Contract.ConfirmationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirmations *ConfirmationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Confirmations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirmations *ConfirmationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirmations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirmations *ConfirmationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirmations.Contract.contract.Transact(opts, method, params...)
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) ConfirmedGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "_confirmedGas")
	return *ret0, err
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) ConfirmedGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmedGas(&_Confirmations.CallOpts)
}

// ConfirmedGas is a free data retrieval call binding the contract method 0x2ea6ece8.
//
// Solidity: function _confirmedGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) ConfirmedGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmedGas(&_Confirmations.CallOpts)
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) ConfirmGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "confirmGas")
	return *ret0, err
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) ConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmGas(&_Confirmations.CallOpts)
}

// ConfirmGas is a free data retrieval call binding the contract method 0x42715aec.
//
// Solidity: function confirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) ConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.ConfirmGas(&_Confirmations.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) Count(opts *bind.CallOpts, arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "count", arg0, arg1, arg2)
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Confirmations.Contract.Count(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// Count is a free data retrieval call binding the contract method 0x7849ed5a.
//
// Solidity: function count(address , bool , bool ) constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) Count(arg0 common.Address, arg1 bool, arg2 bool) (*big.Int, error) {
	return _Confirmations.Contract.Count(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCaller) GetSigner(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "getSigner", i)
	return *ret0, err
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetSigner(&_Confirmations.CallOpts, i)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCallerSession) GetSigner(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetSigner(&_Confirmations.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCaller) GetVoter(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "getVoter", i)
	return *ret0, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetVoter(&_Confirmations.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 i) constant returns(address)
func (_Confirmations *ConfirmationsCallerSession) GetVoter(i *big.Int) (common.Address, error) {
	return _Confirmations.Contract.GetVoter(&_Confirmations.CallOpts, i)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCaller) IsSigner(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "isSigner", voter)
	return *ret0, err
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsSession) IsSigner(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsSigner(&_Confirmations.CallOpts, voter)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCallerSession) IsSigner(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsSigner(&_Confirmations.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCaller) IsVoter(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "isVoter", voter)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsSession) IsVoter(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsVoter(&_Confirmations.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address voter) constant returns(bool)
func (_Confirmations *ConfirmationsCallerSession) IsVoter(voter common.Address) (bool, error) {
	return _Confirmations.Contract.IsVoter(&_Confirmations.CallOpts, voter)
}

// PendingList is a free data retrieval call binding the contract method 0x03aca792.
//
// Solidity: function pendingList(uint256 ) constant returns(uint256 blockNum, uint256 logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsCaller) PendingList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
}, error) {
	ret := new(struct {
		BlockNum  *big.Int
		LogIndex  *big.Int
		EventHash [32]byte
	})
	out := ret
	err := _Confirmations.contract.Call(opts, out, "pendingList", arg0)
	return *ret, err
}

// PendingList is a free data retrieval call binding the contract method 0x03aca792.
//
// Solidity: function pendingList(uint256 ) constant returns(uint256 blockNum, uint256 logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsSession) PendingList(arg0 *big.Int) (struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
}, error) {
	return _Confirmations.Contract.PendingList(&_Confirmations.CallOpts, arg0)
}

// PendingList is a free data retrieval call binding the contract method 0x03aca792.
//
// Solidity: function pendingList(uint256 ) constant returns(uint256 blockNum, uint256 logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsCallerSession) PendingList(arg0 *big.Int) (struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
}, error) {
	return _Confirmations.Contract.PendingList(&_Confirmations.CallOpts, arg0)
}

// PendingListLength is a free data retrieval call binding the contract method 0x26b60630.
//
// Solidity: function pendingListLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) PendingListLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "pendingListLength")
	return *ret0, err
}

// PendingListLength is a free data retrieval call binding the contract method 0x26b60630.
//
// Solidity: function pendingListLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) PendingListLength() (*big.Int, error) {
	return _Confirmations.Contract.PendingListLength(&_Confirmations.CallOpts)
}

// PendingListLength is a free data retrieval call binding the contract method 0x26b60630.
//
// Solidity: function pendingListLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) PendingListLength() (*big.Int, error) {
	return _Confirmations.Contract.PendingListLength(&_Confirmations.CallOpts)
}

// ShouldConfirm is a free data retrieval call binding the contract method 0x859f32c2.
//
// Solidity: function shouldConfirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(bool, uint256)
func (_Confirmations *ConfirmationsCaller) ShouldConfirm(opts *bind.CallOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Confirmations.contract.Call(opts, out, "shouldConfirm", blockNum, logIndex, eventHash)
	return *ret0, *ret1, err
}

// ShouldConfirm is a free data retrieval call binding the contract method 0x859f32c2.
//
// Solidity: function shouldConfirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(bool, uint256)
func (_Confirmations *ConfirmationsSession) ShouldConfirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (bool, *big.Int, error) {
	return _Confirmations.Contract.ShouldConfirm(&_Confirmations.CallOpts, blockNum, logIndex, eventHash)
}

// ShouldConfirm is a free data retrieval call binding the contract method 0x859f32c2.
//
// Solidity: function shouldConfirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(bool, uint256)
func (_Confirmations *ConfirmationsCallerSession) ShouldConfirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (bool, *big.Int, error) {
	return _Confirmations.Contract.ShouldConfirm(&_Confirmations.CallOpts, blockNum, logIndex, eventHash)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) SignersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "signersLength")
	return *ret0, err
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) SignersLength() (*big.Int, error) {
	return _Confirmations.Contract.SignersLength(&_Confirmations.CallOpts)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) SignersLength() (*big.Int, error) {
	return _Confirmations.Contract.SignersLength(&_Confirmations.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsCaller) Status(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "status", arg0, arg1, arg2)
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsSession) Status(arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	return _Confirmations.Contract.Status(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 , uint256 , bytes32 ) constant returns(uint8)
func (_Confirmations *ConfirmationsCallerSession) Status(arg0 *big.Int, arg1 *big.Int, arg2 [32]byte) (uint8, error) {
	return _Confirmations.Contract.Status(&_Confirmations.CallOpts, arg0, arg1, arg2)
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) TotalConfirmGas(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "totalConfirmGas")
	return *ret0, err
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) TotalConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.TotalConfirmGas(&_Confirmations.CallOpts)
}

// TotalConfirmGas is a free data retrieval call binding the contract method 0x5c5f6b39.
//
// Solidity: function totalConfirmGas() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) TotalConfirmGas() (*big.Int, error) {
	return _Confirmations.Contract.TotalConfirmGas(&_Confirmations.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCaller) VotersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Confirmations.contract.Call(opts, out, "votersLength")
	return *ret0, err
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsSession) VotersLength() (*big.Int, error) {
	return _Confirmations.Contract.VotersLength(&_Confirmations.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_Confirmations *ConfirmationsCallerSession) VotersLength() (*big.Int, error) {
	return _Confirmations.Contract.VotersLength(&_Confirmations.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsCaller) Votes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	ret := new(struct {
		Addr  common.Address
		Voter bool
		Add   bool
	})
	out := ret
	err := _Confirmations.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Confirmations.Contract.Votes(&_Confirmations.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) constant returns(address addr, bool voter, bool add)
func (_Confirmations *ConfirmationsCallerSession) Votes(arg0 common.Address) (struct {
	Addr  common.Address
	Voter bool
	Add   bool
}, error) {
	return _Confirmations.Contract.Votes(&_Confirmations.CallOpts, arg0)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsTransactor) Confirm(opts *bind.TransactOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "confirm", blockNum, logIndex, eventHash, valid)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsSession) Confirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.Contract.Confirm(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash, valid)
}

// Confirm is a paid mutator transaction binding the contract method 0x8f7fd0d6.
//
// Solidity: function confirm(uint256 blockNum, uint256 logIndex, bytes32 eventHash, bool valid) returns()
func (_Confirmations *ConfirmationsTransactorSession) Confirm(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte, valid bool) (*types.Transaction, error) {
	return _Confirmations.Contract.Confirm(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash, valid)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsTransactor) Request(opts *bind.TransactOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "request", blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.Contract.Request(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_Confirmations *ConfirmationsTransactorSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _Confirmations.Contract.Request(&_Confirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsTransactor) SetVote(opts *bind.TransactOpts, addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.contract.Transact(opts, "setVote", addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.Contract.SetVote(&_Confirmations.TransactOpts, addr, voter, add)
}

// SetVote is a paid mutator transaction binding the contract method 0x48c5000d.
//
// Solidity: function setVote(address addr, bool voter, bool add) returns()
func (_Confirmations *ConfirmationsTransactorSession) SetVote(addr common.Address, voter bool, add bool) (*types.Transaction, error) {
	return _Confirmations.Contract.SetVote(&_Confirmations.TransactOpts, addr, voter, add)
}

// ConfirmationsConfirmationRequestedIterator is returned from FilterConfirmationRequested and is used to iterate over the raw logs and unpacked data for ConfirmationRequested events raised by the Confirmations contract.
type ConfirmationsConfirmationRequestedIterator struct {
	Event *ConfirmationsConfirmationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfirmationsConfirmationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmationsConfirmationRequested)
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
		it.Event = new(ConfirmationsConfirmationRequested)
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
func (it *ConfirmationsConfirmationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfirmationsConfirmationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfirmationsConfirmationRequested represents a ConfirmationRequested event raised by the Confirmations contract.
type ConfirmationsConfirmationRequested struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmationRequested is a free log retrieval operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) FilterConfirmationRequested(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*ConfirmationsConfirmationRequestedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.FilterLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsConfirmationRequestedIterator{contract: _Confirmations.contract, event: "ConfirmationRequested", logs: logs, sub: sub}, nil
}

// WatchConfirmationRequested is a free log subscription operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) WatchConfirmationRequested(opts *bind.WatchOpts, sink chan<- *ConfirmationsConfirmationRequested, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.WatchLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfirmationsConfirmationRequested)
				if err := _Confirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
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

// ParseConfirmationRequested is a log parse operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_Confirmations *ConfirmationsFilterer) ParseConfirmationRequested(log types.Log) (*ConfirmationsConfirmationRequested, error) {
	event := new(ConfirmationsConfirmationRequested)
	if err := _Confirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConfirmationsConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the Confirmations contract.
type ConfirmationsConfirmedIterator struct {
	Event *ConfirmationsConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfirmationsConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmationsConfirmed)
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
		it.Event = new(ConfirmationsConfirmed)
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
func (it *ConfirmationsConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfirmationsConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfirmationsConfirmed represents a Confirmed event raised by the Confirmations contract.
type ConfirmationsConfirmed struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Valid     bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) FilterConfirmed(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*ConfirmationsConfirmedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.FilterLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmationsConfirmedIterator{contract: _Confirmations.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *ConfirmationsConfirmed, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _Confirmations.contract.WatchLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfirmationsConfirmed)
				if err := _Confirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
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

// ParseConfirmed is a log parse operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_Confirmations *ConfirmationsFilterer) ParseConfirmed(log types.Log) (*ConfirmationsConfirmed, error) {
	event := new(ConfirmationsConfirmed)
	if err := _Confirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IAuthABI is the input ABI used to generate the binding from.
const IAuthABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"signersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IAuthFuncSigs maps the 4-byte function signature to its string representation.
var IAuthFuncSigs = map[string]string{
	"3ffefe4e": "getSigner(uint256)",
	"d07bff0c": "getVoter(uint256)",
	"7df73e27": "isSigner(address)",
	"a7771ee3": "isVoter(address)",
	"41f684f3": "signersLength()",
	"6c6c39fb": "votersLength()",
}

// IAuth is an auto generated Go binding around an GoChain contract.
type IAuth struct {
	IAuthCaller     // Read-only binding to the contract
	IAuthTransactor // Write-only binding to the contract
	IAuthFilterer   // Log filterer for contract events
}

// IAuthCaller is an auto generated read-only Go binding around an GoChain contract.
type IAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthTransactor is an auto generated write-only Go binding around an GoChain contract.
type IAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type IAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type IAuthSession struct {
	Contract     *IAuth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type IAuthCallerSession struct {
	Contract *IAuthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type IAuthTransactorSession struct {
	Contract     *IAuthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthRaw is an auto generated low-level Go binding around an GoChain contract.
type IAuthRaw struct {
	Contract *IAuth // Generic contract binding to access the raw methods on
}

// IAuthCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type IAuthCallerRaw struct {
	Contract *IAuthCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type IAuthTransactorRaw struct {
	Contract *IAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuth creates a new instance of IAuth, bound to a specific deployed contract.
func NewIAuth(address common.Address, backend bind.ContractBackend) (*IAuth, error) {
	contract, err := bindIAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuth{IAuthCaller: IAuthCaller{contract: contract}, IAuthTransactor: IAuthTransactor{contract: contract}, IAuthFilterer: IAuthFilterer{contract: contract}}, nil
}

// NewIAuthCaller creates a new read-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthCaller(address common.Address, caller bind.ContractCaller) (*IAuthCaller, error) {
	contract, err := bindIAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthCaller{contract: contract}, nil
}

// NewIAuthTransactor creates a new write-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthTransactor, error) {
	contract, err := bindIAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthTransactor{contract: contract}, nil
}

// NewIAuthFilterer creates a new log filterer instance of IAuth, bound to a specific deployed contract.
func NewIAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthFilterer, error) {
	contract, err := bindIAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthFilterer{contract: contract}, nil
}

// bindIAuth binds a generic wrapper to an already deployed contract.
func bindIAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.IAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transact(opts, method, params...)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 ) constant returns(address)
func (_IAuth *IAuthCaller) GetSigner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IAuth.contract.Call(opts, out, "getSigner", arg0)
	return *ret0, err
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 ) constant returns(address)
func (_IAuth *IAuthSession) GetSigner(arg0 *big.Int) (common.Address, error) {
	return _IAuth.Contract.GetSigner(&_IAuth.CallOpts, arg0)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 ) constant returns(address)
func (_IAuth *IAuthCallerSession) GetSigner(arg0 *big.Int) (common.Address, error) {
	return _IAuth.Contract.GetSigner(&_IAuth.CallOpts, arg0)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 ) constant returns(address)
func (_IAuth *IAuthCaller) GetVoter(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IAuth.contract.Call(opts, out, "getVoter", arg0)
	return *ret0, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 ) constant returns(address)
func (_IAuth *IAuthSession) GetVoter(arg0 *big.Int) (common.Address, error) {
	return _IAuth.Contract.GetVoter(&_IAuth.CallOpts, arg0)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(uint256 ) constant returns(address)
func (_IAuth *IAuthCallerSession) GetVoter(arg0 *big.Int) (common.Address, error) {
	return _IAuth.Contract.GetVoter(&_IAuth.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) constant returns(bool)
func (_IAuth *IAuthCaller) IsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IAuth.contract.Call(opts, out, "isSigner", arg0)
	return *ret0, err
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) constant returns(bool)
func (_IAuth *IAuthSession) IsSigner(arg0 common.Address) (bool, error) {
	return _IAuth.Contract.IsSigner(&_IAuth.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) constant returns(bool)
func (_IAuth *IAuthCallerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _IAuth.Contract.IsSigner(&_IAuth.CallOpts, arg0)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) constant returns(bool)
func (_IAuth *IAuthCaller) IsVoter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IAuth.contract.Call(opts, out, "isVoter", arg0)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) constant returns(bool)
func (_IAuth *IAuthSession) IsVoter(arg0 common.Address) (bool, error) {
	return _IAuth.Contract.IsVoter(&_IAuth.CallOpts, arg0)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) constant returns(bool)
func (_IAuth *IAuthCallerSession) IsVoter(arg0 common.Address) (bool, error) {
	return _IAuth.Contract.IsVoter(&_IAuth.CallOpts, arg0)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_IAuth *IAuthCaller) SignersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IAuth.contract.Call(opts, out, "signersLength")
	return *ret0, err
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_IAuth *IAuthSession) SignersLength() (*big.Int, error) {
	return _IAuth.Contract.SignersLength(&_IAuth.CallOpts)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() constant returns(uint256)
func (_IAuth *IAuthCallerSession) SignersLength() (*big.Int, error) {
	return _IAuth.Contract.SignersLength(&_IAuth.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_IAuth *IAuthCaller) VotersLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IAuth.contract.Call(opts, out, "votersLength")
	return *ret0, err
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_IAuth *IAuthSession) VotersLength() (*big.Int, error) {
	return _IAuth.Contract.VotersLength(&_IAuth.CallOpts)
}

// VotersLength is a free data retrieval call binding the contract method 0x6c6c39fb.
//
// Solidity: function votersLength() constant returns(uint256)
func (_IAuth *IAuthCallerSession) VotersLength() (*big.Int, error) {
	return _IAuth.Contract.VotersLength(&_IAuth.CallOpts)
}

// IConfirmationsABI is the input ABI used to generate the binding from.
const IConfirmationsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"Confirmed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"request\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"eventHash\",\"type\":\"bytes32\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumIConfirmations.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IConfirmationsFuncSigs maps the 4-byte function signature to its string representation.
var IConfirmationsFuncSigs = map[string]string{
	"32030803": "request(uint256,uint256,bytes32)",
	"fad3ffd6": "status(uint256,uint256,bytes32)",
}

// IConfirmations is an auto generated Go binding around an GoChain contract.
type IConfirmations struct {
	IConfirmationsCaller     // Read-only binding to the contract
	IConfirmationsTransactor // Write-only binding to the contract
	IConfirmationsFilterer   // Log filterer for contract events
}

// IConfirmationsCaller is an auto generated read-only Go binding around an GoChain contract.
type IConfirmationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfirmationsTransactor is an auto generated write-only Go binding around an GoChain contract.
type IConfirmationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfirmationsFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type IConfirmationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfirmationsSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type IConfirmationsSession struct {
	Contract     *IConfirmations   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IConfirmationsCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type IConfirmationsCallerSession struct {
	Contract *IConfirmationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IConfirmationsTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type IConfirmationsTransactorSession struct {
	Contract     *IConfirmationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IConfirmationsRaw is an auto generated low-level Go binding around an GoChain contract.
type IConfirmationsRaw struct {
	Contract *IConfirmations // Generic contract binding to access the raw methods on
}

// IConfirmationsCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type IConfirmationsCallerRaw struct {
	Contract *IConfirmationsCaller // Generic read-only contract binding to access the raw methods on
}

// IConfirmationsTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type IConfirmationsTransactorRaw struct {
	Contract *IConfirmationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIConfirmations creates a new instance of IConfirmations, bound to a specific deployed contract.
func NewIConfirmations(address common.Address, backend bind.ContractBackend) (*IConfirmations, error) {
	contract, err := bindIConfirmations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IConfirmations{IConfirmationsCaller: IConfirmationsCaller{contract: contract}, IConfirmationsTransactor: IConfirmationsTransactor{contract: contract}, IConfirmationsFilterer: IConfirmationsFilterer{contract: contract}}, nil
}

// NewIConfirmationsCaller creates a new read-only instance of IConfirmations, bound to a specific deployed contract.
func NewIConfirmationsCaller(address common.Address, caller bind.ContractCaller) (*IConfirmationsCaller, error) {
	contract, err := bindIConfirmations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IConfirmationsCaller{contract: contract}, nil
}

// NewIConfirmationsTransactor creates a new write-only instance of IConfirmations, bound to a specific deployed contract.
func NewIConfirmationsTransactor(address common.Address, transactor bind.ContractTransactor) (*IConfirmationsTransactor, error) {
	contract, err := bindIConfirmations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IConfirmationsTransactor{contract: contract}, nil
}

// NewIConfirmationsFilterer creates a new log filterer instance of IConfirmations, bound to a specific deployed contract.
func NewIConfirmationsFilterer(address common.Address, filterer bind.ContractFilterer) (*IConfirmationsFilterer, error) {
	contract, err := bindIConfirmations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IConfirmationsFilterer{contract: contract}, nil
}

// bindIConfirmations binds a generic wrapper to an already deployed contract.
func bindIConfirmations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IConfirmationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfirmations *IConfirmationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IConfirmations.Contract.IConfirmationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfirmations *IConfirmationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfirmations.Contract.IConfirmationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfirmations *IConfirmationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfirmations.Contract.IConfirmationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfirmations *IConfirmationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IConfirmations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfirmations *IConfirmationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfirmations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfirmations *IConfirmationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfirmations.Contract.contract.Transact(opts, method, params...)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(uint8)
func (_IConfirmations *IConfirmationsCaller) Status(opts *bind.CallOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IConfirmations.contract.Call(opts, out, "status", blockNum, logIndex, eventHash)
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(uint8)
func (_IConfirmations *IConfirmationsSession) Status(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (uint8, error) {
	return _IConfirmations.Contract.Status(&_IConfirmations.CallOpts, blockNum, logIndex, eventHash)
}

// Status is a free data retrieval call binding the contract method 0xfad3ffd6.
//
// Solidity: function status(uint256 blockNum, uint256 logIndex, bytes32 eventHash) constant returns(uint8)
func (_IConfirmations *IConfirmationsCallerSession) Status(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (uint8, error) {
	return _IConfirmations.Contract.Status(&_IConfirmations.CallOpts, blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_IConfirmations *IConfirmationsTransactor) Request(opts *bind.TransactOpts, blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _IConfirmations.contract.Transact(opts, "request", blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_IConfirmations *IConfirmationsSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _IConfirmations.Contract.Request(&_IConfirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// Request is a paid mutator transaction binding the contract method 0x32030803.
//
// Solidity: function request(uint256 blockNum, uint256 logIndex, bytes32 eventHash) returns()
func (_IConfirmations *IConfirmationsTransactorSession) Request(blockNum *big.Int, logIndex *big.Int, eventHash [32]byte) (*types.Transaction, error) {
	return _IConfirmations.Contract.Request(&_IConfirmations.TransactOpts, blockNum, logIndex, eventHash)
}

// IConfirmationsConfirmationRequestedIterator is returned from FilterConfirmationRequested and is used to iterate over the raw logs and unpacked data for ConfirmationRequested events raised by the IConfirmations contract.
type IConfirmationsConfirmationRequestedIterator struct {
	Event *IConfirmationsConfirmationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IConfirmationsConfirmationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfirmationsConfirmationRequested)
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
		it.Event = new(IConfirmationsConfirmationRequested)
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
func (it *IConfirmationsConfirmationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfirmationsConfirmationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfirmationsConfirmationRequested represents a ConfirmationRequested event raised by the IConfirmations contract.
type IConfirmationsConfirmationRequested struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmationRequested is a free log retrieval operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_IConfirmations *IConfirmationsFilterer) FilterConfirmationRequested(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*IConfirmationsConfirmationRequestedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _IConfirmations.contract.FilterLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &IConfirmationsConfirmationRequestedIterator{contract: _IConfirmations.contract, event: "ConfirmationRequested", logs: logs, sub: sub}, nil
}

// WatchConfirmationRequested is a free log subscription operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_IConfirmations *IConfirmationsFilterer) WatchConfirmationRequested(opts *bind.WatchOpts, sink chan<- *IConfirmationsConfirmationRequested, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _IConfirmations.contract.WatchLogs(opts, "ConfirmationRequested", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfirmationsConfirmationRequested)
				if err := _IConfirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
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

// ParseConfirmationRequested is a log parse operation binding the contract event 0x2382b1fef9adb7c2a096ef31e2e402e08d3b5f9eb5fc18ea9a56835e44627b03.
//
// Solidity: event ConfirmationRequested(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash)
func (_IConfirmations *IConfirmationsFilterer) ParseConfirmationRequested(log types.Log) (*IConfirmationsConfirmationRequested, error) {
	event := new(IConfirmationsConfirmationRequested)
	if err := _IConfirmations.contract.UnpackLog(event, "ConfirmationRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IConfirmationsConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the IConfirmations contract.
type IConfirmationsConfirmedIterator struct {
	Event *IConfirmationsConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IConfirmationsConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IConfirmationsConfirmed)
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
		it.Event = new(IConfirmationsConfirmed)
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
func (it *IConfirmationsConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IConfirmationsConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IConfirmationsConfirmed represents a Confirmed event raised by the IConfirmations contract.
type IConfirmationsConfirmed struct {
	BlockNum  *big.Int
	LogIndex  *big.Int
	EventHash [32]byte
	Valid     bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_IConfirmations *IConfirmationsFilterer) FilterConfirmed(opts *bind.FilterOpts, blockNum []*big.Int, logIndex []*big.Int) (*IConfirmationsConfirmedIterator, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _IConfirmations.contract.FilterLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return &IConfirmationsConfirmedIterator{contract: _IConfirmations.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_IConfirmations *IConfirmationsFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *IConfirmationsConfirmed, blockNum []*big.Int, logIndex []*big.Int) (event.Subscription, error) {

	var blockNumRule []interface{}
	for _, blockNumItem := range blockNum {
		blockNumRule = append(blockNumRule, blockNumItem)
	}
	var logIndexRule []interface{}
	for _, logIndexItem := range logIndex {
		logIndexRule = append(logIndexRule, logIndexItem)
	}

	logs, sub, err := _IConfirmations.contract.WatchLogs(opts, "Confirmed", blockNumRule, logIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IConfirmationsConfirmed)
				if err := _IConfirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
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

// ParseConfirmed is a log parse operation binding the contract event 0xdf1051063b9bab79d715a3919f387eb9ee4291be1e9241dffb4500694141f25c.
//
// Solidity: event Confirmed(uint256 indexed blockNum, uint256 indexed logIndex, bytes32 eventHash, bool valid)
func (_IConfirmations *IConfirmationsFilterer) ParseConfirmed(log types.Log) (*IConfirmationsConfirmed, error) {
	event := new(IConfirmationsConfirmed)
	if err := _IConfirmations.contract.UnpackLog(event, "Confirmed", log); err != nil {
		return nil, err
	}
	return event, nil
}
