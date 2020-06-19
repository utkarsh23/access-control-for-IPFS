// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AccessControlIPFSABI is the input ABI used to generate the binding from.
const AccessControlIPFSABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"chunk\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"grantees\",\"type\":\"address[]\"}],\"name\":\"addBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"chunks\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"grantees\",\"type\":\"address[]\"}],\"name\":\"addBlockMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"chunk\",\"type\":\"bytes32\"}],\"name\":\"checkAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasAccess\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"chunks\",\"type\":\"bytes32[]\"}],\"name\":\"checkAccessMultiple\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_checkAccessMultiple\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"chunk\",\"type\":\"bytes32\"}],\"name\":\"deleteBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"chunks\",\"type\":\"bytes32[]\"}],\"name\":\"deleteBlockMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"chunk\",\"type\":\"bytes32\"}],\"name\":\"grantAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"chunks\",\"type\":\"bytes32[]\"}],\"name\":\"grantAccessMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"chunk\",\"type\":\"bytes32\"}],\"name\":\"removeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"chunks\",\"type\":\"bytes32[]\"}],\"name\":\"removeAccessMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccessControlIPFSFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlIPFSFuncSigs = map[string]string{
	"1bd8e589": "addBlock(address,bytes32,address[])",
	"5294fddc": "addBlockMultiple(address,bytes32[],address[])",
	"7461f24a": "checkAccess(address,bytes32)",
	"4f897b0b": "checkAccessMultiple(address,bytes32[])",
	"deeb2613": "deleteBlock(address,bytes32)",
	"28db8a01": "deleteBlockMultiple(address,bytes32[])",
	"d9896430": "grantAccess(address,address,bytes32)",
	"6a71b26c": "grantAccessMultiple(address,address,bytes32[])",
	"750c6e5c": "removeAccess(address,address,bytes32)",
	"78595290": "removeAccessMultiple(address,address,bytes32[])",
}

// AccessControlIPFSBin is the compiled bytecode used for deploying new contracts.
var AccessControlIPFSBin = "0x608060405234801561001057600080fd5b50610ed2806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637461f24a116100665780637461f24a146104fa578063750c6e5c1461053a5780637859529014610570578063d98964301461062a578063deeb2613146106605761009e565b80631bd8e589146100a357806328db8a011461015b5780634f897b0b1461020c5780635294fddc1461030d5780636a71b26c14610440575b600080fd5b610159600480360360608110156100b957600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b8111156100e857600080fd5b8201836020820111156100fa57600080fd5b803590602001918460208302840111600160201b8311171561011b57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061068c945050505050565b005b6101596004803603604081101561017157600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561019b57600080fd5b8201836020820111156101ad57600080fd5b803590602001918460208302840111600160201b831117156101ce57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610775945050505050565b6102bd6004803603604081101561022257600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561024c57600080fd5b82018360208201111561025e57600080fd5b803590602001918460208302840111600160201b8311171561027f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506107f7945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102f95781810151838201526020016102e1565b505050509050019250505060405180910390f35b6101596004803603606081101561032357600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561034d57600080fd5b82018360208201111561035f57600080fd5b803590602001918460208302840111600160201b8311171561038057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156103cf57600080fd5b8201836020820111156103e157600080fd5b803590602001918460208302840111600160201b8311171561040257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610917945050505050565b6101596004803603606081101561045657600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561048957600080fd5b82018360208201111561049b57600080fd5b803590602001918460208302840111600160201b831117156104bc57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610a54945050505050565b6105266004803603604081101561051057600080fd5b506001600160a01b038135169060200135610b2d565b604080519115158252519081900360200190f35b6101596004803603606081101561055057600080fd5b506001600160a01b03813581169160208101359091169060400135610ba7565b6101596004803603606081101561058657600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b8111156105b957600080fd5b8201836020820111156105cb57600080fd5b803590602001918460208302840111600160201b831117156105ec57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610c11945050505050565b6101596004803603606081101561064057600080fd5b506001600160a01b03813581169160208101359091169060400135610cba565b6101596004803603604081101561067657600080fd5b506001600160a01b038135169060200135610d4d565b8160001a60f81b6001600160f81b0319166106a657610770565b600080546001600160a01b038086166001600160a01b03199283161780845585845260208490526040842080549093169116178155600280548392916106f0918184019190610d95565b50600091508190505b835181101561076c5783818151811061070e57fe5b6020908102919091018101516001600160a01b03811660008181526001878101855260408220805460ff19168217905560028801805480830182559083529490912090930180546001600160a01b03191690911790559250016106f9565b5050505b505050565b6000805b82518110156107f15782818151811061078e57fe5b60200260200101519150816000602081106107a557fe5b1a60f81b6001600160f81b0319166107bc576107e9565b600082815260208190526040812080546001600160a01b0319168155906107e66002830182610de5565b50505b600101610779565b50505050565b6060815167ffffffffffffffff8111801561081157600080fd5b5060405190808252806020026020018201604052801561083b578160200160208202803683370190505b5090506000805b835181101561090f5783818151811061085757fe5b60200260200101519150600083828151811061086f57fe5b911515602092830291909101909101528160001a60f81b6001600160f81b03191661089957610907565b600082815260208190526040902080546001600160a01b03878116911614806108e157506001600160a01b0386166000908152600180830160205260409091205460ff161515145b156109055760018483815181106108f457fe5b911515602092830291909101909101525b505b600101610842565b505092915050565b6000806060815b8551811015610a4b5785818151811061093357fe5b602002602001015193508360006020811061094a57fe5b1a60f81b6001600160f81b03191661096157610a43565b6040805180820182526001600160a01b0389811682526020808301868152600089815280835294909420835181546001600160a01b031916931692909217825592518051929391926109b99260028501920190610e06565b5060009150505b8551811015610a41578581815181106109d557fe5b602090810291909101810151600087815280835260408082206001600160a01b03841680845260018281018752928420805460ff1916841790558386526002909101805480840182559084529490922090930180546001600160a01b03191690911790559450016109c0565b505b60010161091e565b50505050505050565b6000805b8251811015610b2657828181518110610a6d57fe5b6020026020010151915081600060208110610a8457fe5b1a60f81b6001600160f81b031916610a9b57610b1e565b600082815260208190526040902081158015610ac4575080546001600160a01b03878116911614155b15610ad157505050610770565b6001600160a01b03851660008181526001838101602090815260408320805460ff191683179055600290940180549182018155825292902090910180546001600160a01b03191690911790555b600101610a58565b5050505050565b600081811a60f81b6001600160f81b031916610b4857610ba1565b600082815260208190526040902080546001600160a01b0385811691161480610b9057506001600160a01b0384166000908152600180830160205260409091205460ff161515145b15610b9f575060019050610ba1565b505b92915050565b8060001a60f81b6001600160f81b031916610bc157610770565b600081815260208190526040902080546001600160a01b03858116911614610be95750610770565b6001600160a01b0383166000908152600190910160205260409020805460ff19169055505050565b6000805b8251811015610b2657828181518110610c2a57fe5b6020026020010151915081600060208110610c4157fe5b1a60f81b6001600160f81b031916610c5857610cb2565b600082815260208190526040902081158015610c81575080546001600160a01b03878116911614155b15610c8e57505050610770565b6001600160a01b0385166000908152600190910160205260409020805460ff191690555b600101610c15565b8060001a60f81b6001600160f81b031916610cd457610770565b600081815260208190526040902080546001600160a01b03858116911614610cfc5750610770565b6001600160a01b03831660008181526001838101602090815260408320805460ff191683179055600290940180549182018155825292902090910180546001600160a01b0319169091179055505050565b8060001a60f81b6001600160f81b031916610d6757610d91565b600081815260208190526040812080546001600160a01b0319168155906107f16002830182610de5565b5050565b828054828255906000526020600020908101928215610dd55760005260206000209182015b82811115610dd5578254825591600101919060010190610dba565b50610de1929150610e5b565b5090565b5080546000825590600052602060002090810190610e039190610e82565b50565b828054828255906000526020600020908101928215610dd5579160200282015b82811115610dd557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610e26565b610e7f91905b80821115610de15780546001600160a01b0319168155600101610e61565b90565b610e7f91905b80821115610de15760008155600101610e8856fea26469706673582212201fd347c73008a71155da4491368225efc958590637588d82d4ddc7212dd4efb164736f6c63430006080033"

// DeployAccessControlIPFS deploys a new Ethereum contract, binding an instance of AccessControlIPFS to it.
func DeployAccessControlIPFS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControlIPFS, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlIPFSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlIPFSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlIPFS{AccessControlIPFSCaller: AccessControlIPFSCaller{contract: contract}, AccessControlIPFSTransactor: AccessControlIPFSTransactor{contract: contract}, AccessControlIPFSFilterer: AccessControlIPFSFilterer{contract: contract}}, nil
}

// AccessControlIPFS is an auto generated Go binding around an Ethereum contract.
type AccessControlIPFS struct {
	AccessControlIPFSCaller     // Read-only binding to the contract
	AccessControlIPFSTransactor // Write-only binding to the contract
	AccessControlIPFSFilterer   // Log filterer for contract events
}

// AccessControlIPFSCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlIPFSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlIPFSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlIPFSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlIPFSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlIPFSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlIPFSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlIPFSSession struct {
	Contract     *AccessControlIPFS // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccessControlIPFSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlIPFSCallerSession struct {
	Contract *AccessControlIPFSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AccessControlIPFSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlIPFSTransactorSession struct {
	Contract     *AccessControlIPFSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AccessControlIPFSRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlIPFSRaw struct {
	Contract *AccessControlIPFS // Generic contract binding to access the raw methods on
}

// AccessControlIPFSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlIPFSCallerRaw struct {
	Contract *AccessControlIPFSCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlIPFSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlIPFSTransactorRaw struct {
	Contract *AccessControlIPFSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlIPFS creates a new instance of AccessControlIPFS, bound to a specific deployed contract.
func NewAccessControlIPFS(address common.Address, backend bind.ContractBackend) (*AccessControlIPFS, error) {
	contract, err := bindAccessControlIPFS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlIPFS{AccessControlIPFSCaller: AccessControlIPFSCaller{contract: contract}, AccessControlIPFSTransactor: AccessControlIPFSTransactor{contract: contract}, AccessControlIPFSFilterer: AccessControlIPFSFilterer{contract: contract}}, nil
}

// NewAccessControlIPFSCaller creates a new read-only instance of AccessControlIPFS, bound to a specific deployed contract.
func NewAccessControlIPFSCaller(address common.Address, caller bind.ContractCaller) (*AccessControlIPFSCaller, error) {
	contract, err := bindAccessControlIPFS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlIPFSCaller{contract: contract}, nil
}

// NewAccessControlIPFSTransactor creates a new write-only instance of AccessControlIPFS, bound to a specific deployed contract.
func NewAccessControlIPFSTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlIPFSTransactor, error) {
	contract, err := bindAccessControlIPFS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlIPFSTransactor{contract: contract}, nil
}

// NewAccessControlIPFSFilterer creates a new log filterer instance of AccessControlIPFS, bound to a specific deployed contract.
func NewAccessControlIPFSFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlIPFSFilterer, error) {
	contract, err := bindAccessControlIPFS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlIPFSFilterer{contract: contract}, nil
}

// bindAccessControlIPFS binds a generic wrapper to an already deployed contract.
func bindAccessControlIPFS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlIPFSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlIPFS *AccessControlIPFSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControlIPFS.Contract.AccessControlIPFSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlIPFS *AccessControlIPFSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.AccessControlIPFSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlIPFS *AccessControlIPFSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.AccessControlIPFSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlIPFS *AccessControlIPFSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControlIPFS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlIPFS *AccessControlIPFSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlIPFS *AccessControlIPFSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.contract.Transact(opts, method, params...)
}

// CheckAccess is a free data retrieval call binding the contract method 0x7461f24a.
//
// Solidity: function checkAccess(address grantee, bytes32 chunk) view returns(bool hasAccess)
func (_AccessControlIPFS *AccessControlIPFSCaller) CheckAccess(opts *bind.CallOpts, grantee common.Address, chunk [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccessControlIPFS.contract.Call(opts, out, "checkAccess", grantee, chunk)
	return *ret0, err
}

// CheckAccess is a free data retrieval call binding the contract method 0x7461f24a.
//
// Solidity: function checkAccess(address grantee, bytes32 chunk) view returns(bool hasAccess)
func (_AccessControlIPFS *AccessControlIPFSSession) CheckAccess(grantee common.Address, chunk [32]byte) (bool, error) {
	return _AccessControlIPFS.Contract.CheckAccess(&_AccessControlIPFS.CallOpts, grantee, chunk)
}

// CheckAccess is a free data retrieval call binding the contract method 0x7461f24a.
//
// Solidity: function checkAccess(address grantee, bytes32 chunk) view returns(bool hasAccess)
func (_AccessControlIPFS *AccessControlIPFSCallerSession) CheckAccess(grantee common.Address, chunk [32]byte) (bool, error) {
	return _AccessControlIPFS.Contract.CheckAccess(&_AccessControlIPFS.CallOpts, grantee, chunk)
}

// CheckAccessMultiple is a free data retrieval call binding the contract method 0x4f897b0b.
//
// Solidity: function checkAccessMultiple(address grantee, bytes32[] chunks) view returns(bool[] _checkAccessMultiple)
func (_AccessControlIPFS *AccessControlIPFSCaller) CheckAccessMultiple(opts *bind.CallOpts, grantee common.Address, chunks [][32]byte) ([]bool, error) {
	var (
		ret0 = new([]bool)
	)
	out := ret0
	err := _AccessControlIPFS.contract.Call(opts, out, "checkAccessMultiple", grantee, chunks)
	return *ret0, err
}

// CheckAccessMultiple is a free data retrieval call binding the contract method 0x4f897b0b.
//
// Solidity: function checkAccessMultiple(address grantee, bytes32[] chunks) view returns(bool[] _checkAccessMultiple)
func (_AccessControlIPFS *AccessControlIPFSSession) CheckAccessMultiple(grantee common.Address, chunks [][32]byte) ([]bool, error) {
	return _AccessControlIPFS.Contract.CheckAccessMultiple(&_AccessControlIPFS.CallOpts, grantee, chunks)
}

// CheckAccessMultiple is a free data retrieval call binding the contract method 0x4f897b0b.
//
// Solidity: function checkAccessMultiple(address grantee, bytes32[] chunks) view returns(bool[] _checkAccessMultiple)
func (_AccessControlIPFS *AccessControlIPFSCallerSession) CheckAccessMultiple(grantee common.Address, chunks [][32]byte) ([]bool, error) {
	return _AccessControlIPFS.Contract.CheckAccessMultiple(&_AccessControlIPFS.CallOpts, grantee, chunks)
}

// AddBlock is a paid mutator transaction binding the contract method 0x1bd8e589.
//
// Solidity: function addBlock(address sender, bytes32 chunk, address[] grantees) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) AddBlock(opts *bind.TransactOpts, sender common.Address, chunk [32]byte, grantees []common.Address) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "addBlock", sender, chunk, grantees)
}

// AddBlock is a paid mutator transaction binding the contract method 0x1bd8e589.
//
// Solidity: function addBlock(address sender, bytes32 chunk, address[] grantees) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) AddBlock(sender common.Address, chunk [32]byte, grantees []common.Address) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.AddBlock(&_AccessControlIPFS.TransactOpts, sender, chunk, grantees)
}

// AddBlock is a paid mutator transaction binding the contract method 0x1bd8e589.
//
// Solidity: function addBlock(address sender, bytes32 chunk, address[] grantees) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) AddBlock(sender common.Address, chunk [32]byte, grantees []common.Address) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.AddBlock(&_AccessControlIPFS.TransactOpts, sender, chunk, grantees)
}

// AddBlockMultiple is a paid mutator transaction binding the contract method 0x5294fddc.
//
// Solidity: function addBlockMultiple(address sender, bytes32[] chunks, address[] grantees) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) AddBlockMultiple(opts *bind.TransactOpts, sender common.Address, chunks [][32]byte, grantees []common.Address) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "addBlockMultiple", sender, chunks, grantees)
}

// AddBlockMultiple is a paid mutator transaction binding the contract method 0x5294fddc.
//
// Solidity: function addBlockMultiple(address sender, bytes32[] chunks, address[] grantees) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) AddBlockMultiple(sender common.Address, chunks [][32]byte, grantees []common.Address) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.AddBlockMultiple(&_AccessControlIPFS.TransactOpts, sender, chunks, grantees)
}

// AddBlockMultiple is a paid mutator transaction binding the contract method 0x5294fddc.
//
// Solidity: function addBlockMultiple(address sender, bytes32[] chunks, address[] grantees) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) AddBlockMultiple(sender common.Address, chunks [][32]byte, grantees []common.Address) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.AddBlockMultiple(&_AccessControlIPFS.TransactOpts, sender, chunks, grantees)
}

// DeleteBlock is a paid mutator transaction binding the contract method 0xdeeb2613.
//
// Solidity: function deleteBlock(address sender, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) DeleteBlock(opts *bind.TransactOpts, sender common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "deleteBlock", sender, chunk)
}

// DeleteBlock is a paid mutator transaction binding the contract method 0xdeeb2613.
//
// Solidity: function deleteBlock(address sender, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) DeleteBlock(sender common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.DeleteBlock(&_AccessControlIPFS.TransactOpts, sender, chunk)
}

// DeleteBlock is a paid mutator transaction binding the contract method 0xdeeb2613.
//
// Solidity: function deleteBlock(address sender, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) DeleteBlock(sender common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.DeleteBlock(&_AccessControlIPFS.TransactOpts, sender, chunk)
}

// DeleteBlockMultiple is a paid mutator transaction binding the contract method 0x28db8a01.
//
// Solidity: function deleteBlockMultiple(address sender, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) DeleteBlockMultiple(opts *bind.TransactOpts, sender common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "deleteBlockMultiple", sender, chunks)
}

// DeleteBlockMultiple is a paid mutator transaction binding the contract method 0x28db8a01.
//
// Solidity: function deleteBlockMultiple(address sender, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) DeleteBlockMultiple(sender common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.DeleteBlockMultiple(&_AccessControlIPFS.TransactOpts, sender, chunks)
}

// DeleteBlockMultiple is a paid mutator transaction binding the contract method 0x28db8a01.
//
// Solidity: function deleteBlockMultiple(address sender, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) DeleteBlockMultiple(sender common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.DeleteBlockMultiple(&_AccessControlIPFS.TransactOpts, sender, chunks)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xd9896430.
//
// Solidity: function grantAccess(address sender, address grantee, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) GrantAccess(opts *bind.TransactOpts, sender common.Address, grantee common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "grantAccess", sender, grantee, chunk)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xd9896430.
//
// Solidity: function grantAccess(address sender, address grantee, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) GrantAccess(sender common.Address, grantee common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.GrantAccess(&_AccessControlIPFS.TransactOpts, sender, grantee, chunk)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xd9896430.
//
// Solidity: function grantAccess(address sender, address grantee, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) GrantAccess(sender common.Address, grantee common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.GrantAccess(&_AccessControlIPFS.TransactOpts, sender, grantee, chunk)
}

// GrantAccessMultiple is a paid mutator transaction binding the contract method 0x6a71b26c.
//
// Solidity: function grantAccessMultiple(address sender, address grantee, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) GrantAccessMultiple(opts *bind.TransactOpts, sender common.Address, grantee common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "grantAccessMultiple", sender, grantee, chunks)
}

// GrantAccessMultiple is a paid mutator transaction binding the contract method 0x6a71b26c.
//
// Solidity: function grantAccessMultiple(address sender, address grantee, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) GrantAccessMultiple(sender common.Address, grantee common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.GrantAccessMultiple(&_AccessControlIPFS.TransactOpts, sender, grantee, chunks)
}

// GrantAccessMultiple is a paid mutator transaction binding the contract method 0x6a71b26c.
//
// Solidity: function grantAccessMultiple(address sender, address grantee, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) GrantAccessMultiple(sender common.Address, grantee common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.GrantAccessMultiple(&_AccessControlIPFS.TransactOpts, sender, grantee, chunks)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x750c6e5c.
//
// Solidity: function removeAccess(address sender, address grantee, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) RemoveAccess(opts *bind.TransactOpts, sender common.Address, grantee common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "removeAccess", sender, grantee, chunk)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x750c6e5c.
//
// Solidity: function removeAccess(address sender, address grantee, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) RemoveAccess(sender common.Address, grantee common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.RemoveAccess(&_AccessControlIPFS.TransactOpts, sender, grantee, chunk)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x750c6e5c.
//
// Solidity: function removeAccess(address sender, address grantee, bytes32 chunk) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) RemoveAccess(sender common.Address, grantee common.Address, chunk [32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.RemoveAccess(&_AccessControlIPFS.TransactOpts, sender, grantee, chunk)
}

// RemoveAccessMultiple is a paid mutator transaction binding the contract method 0x78595290.
//
// Solidity: function removeAccessMultiple(address sender, address grantee, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactor) RemoveAccessMultiple(opts *bind.TransactOpts, sender common.Address, grantee common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.contract.Transact(opts, "removeAccessMultiple", sender, grantee, chunks)
}

// RemoveAccessMultiple is a paid mutator transaction binding the contract method 0x78595290.
//
// Solidity: function removeAccessMultiple(address sender, address grantee, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSSession) RemoveAccessMultiple(sender common.Address, grantee common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.RemoveAccessMultiple(&_AccessControlIPFS.TransactOpts, sender, grantee, chunks)
}

// RemoveAccessMultiple is a paid mutator transaction binding the contract method 0x78595290.
//
// Solidity: function removeAccessMultiple(address sender, address grantee, bytes32[] chunks) returns()
func (_AccessControlIPFS *AccessControlIPFSTransactorSession) RemoveAccessMultiple(sender common.Address, grantee common.Address, chunks [][32]byte) (*types.Transaction, error) {
	return _AccessControlIPFS.Contract.RemoveAccessMultiple(&_AccessControlIPFS.TransactOpts, sender, grantee, chunks)
}
