// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm

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

// ProviderNode is an auto generated low-level Go binding around an user-defined struct.
type ProviderNode struct {
	Pubkey  [32]byte
	NetAddr [32]byte
}

// ProviderABI is the input ABI used to generate the binding from.
const ProviderABI = "[]"

// ProviderBin is the compiled bytecode used for deploying new contracts.
var ProviderBin = "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b9d838faa9b8c7a3ed907c8088e5b6a9c6828252202c03724a7edaee7854a35f64736f6c634300080e0033"

// DeployProvider deploys a new Ethereum contract, binding an instance of Provider to it.
func DeployProvider(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Provider, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProviderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Provider{ProviderCaller: ProviderCaller{contract: contract}, ProviderTransactor: ProviderTransactor{contract: contract}, ProviderFilterer: ProviderFilterer{contract: contract}}, nil
}

// Provider is an auto generated Go binding around an Ethereum contract.
type Provider struct {
	ProviderCaller     // Read-only binding to the contract
	ProviderTransactor // Write-only binding to the contract
	ProviderFilterer   // Log filterer for contract events
}

// ProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderSession struct {
	Contract     *Provider         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderCallerSession struct {
	Contract *ProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderTransactorSession struct {
	Contract     *ProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderRaw struct {
	Contract *Provider // Generic contract binding to access the raw methods on
}

// ProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderCallerRaw struct {
	Contract *ProviderCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderTransactorRaw struct {
	Contract *ProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProvider creates a new instance of Provider, bound to a specific deployed contract.
func NewProvider(address common.Address, backend bind.ContractBackend) (*Provider, error) {
	contract, err := bindProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Provider{ProviderCaller: ProviderCaller{contract: contract}, ProviderTransactor: ProviderTransactor{contract: contract}, ProviderFilterer: ProviderFilterer{contract: contract}}, nil
}

// NewProviderCaller creates a new read-only instance of Provider, bound to a specific deployed contract.
func NewProviderCaller(address common.Address, caller bind.ContractCaller) (*ProviderCaller, error) {
	contract, err := bindProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderCaller{contract: contract}, nil
}

// NewProviderTransactor creates a new write-only instance of Provider, bound to a specific deployed contract.
func NewProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderTransactor, error) {
	contract, err := bindProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderTransactor{contract: contract}, nil
}

// NewProviderFilterer creates a new log filterer instance of Provider, bound to a specific deployed contract.
func NewProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderFilterer, error) {
	contract, err := bindProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderFilterer{contract: contract}, nil
}

// bindProvider binds a generic wrapper to an already deployed contract.
func bindProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provider *ProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provider.Contract.ProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provider *ProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.Contract.ProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provider *ProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provider.Contract.ProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provider *ProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provider *ProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provider *ProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provider.Contract.contract.Transact(opts, method, params...)
}

// ProvidersABI is the input ABI used to generate the binding from.
const ProvidersABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"name\":\"ProviderExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"name\":\"ProviderJoin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"lookup\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pubkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"internalType\":\"structProvider.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProvidersBin is the compiled bytecode used for deploying new contracts.
var ProvidersBin = "0x608060405234801561001057600080fd5b50610e40806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80635b419a6514610051578063d4b6b5da1461006d578063e9fad8ee1461009d578063edc922a9146100a7575b600080fd5b61006b6004803603810190610066919061097a565b6100c5565b005b61008760048036038101906100829190610a18565b61016e565b6040516100949190610a83565b60405180910390f35b6100a56101b0565b005b6100af610218565b6040516100bc9190610b5c565b60405180910390f35b8160008160001c0361010c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010390610bdb565b60405180910390fd5b61012433848460006102ae909392919063ffffffff16565b81833373ffffffffffffffffffffffffffffffffffffffff167f15b88bb7d80477950c43dc4f6fd378c0dd62f99cffcc5f99f30dbd117c4f57ff60405160405180910390a4505050565b61017661091f565b61018a82600061044390919063ffffffff16565b604051806040016040529081600082015481526020016001820154815250509050919050565b60006101c633600061049190919063ffffffff16565b9050806020015181600001513373ffffffffffffffffffffffffffffffffffffffff167fc72f425864663e90861a0510ca302816b37bb1f9ab5203223088f8a4a25fb55860405160405180910390a450565b6060610224600061065c565b8054806020026020016040519081016040528092919081815260200182805480156102a457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161025a575b5050505050905090565b60008460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016000015460001c1461033b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033290610c47565b60405180910390fd5b6040518060400160405280856001018054905081526020016040518060400160405280858152602001848152508152508460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001016000820151816000015560208201518160010155505090505083600101839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101905092915050565b61049961091f565b60008360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806040016040529081600082015481526020016001820160405180604001604052908160008201548152602001600182015481525050815250509050600081602001516000015160001c14158015610541575083600101805490508160000151105b610580576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057790610cb3565b60405180910390fd5b6105a28482600001516001876001018054905061059d9190610d0c565b610669565b836001018054806105b6576105b5610d40565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055806020015191508360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160009055600182016000808201600090556001820160009055505050505092915050565b6000816001019050919050565b8260010180549050821080156106855750826001018054905081105b6106c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bb90610dbb565b60405180910390fd5b60008360010183815481106106dc576106db610ddb565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083600101828154811061071f5761071e610ddb565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168460010184815481106107605761075f610ddb565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808460010183815481106107bf576107be610ddb565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508184600001600086600101868154811061082457610823610ddb565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550828460000160008660010185815481106108ad576108ac610ddb565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555050505050565b604051806040016040528060008019168152602001600080191681525090565b600080fd5b6000819050919050565b61095781610944565b811461096257600080fd5b50565b6000813590506109748161094e565b92915050565b600080604083850312156109915761099061093f565b5b600061099f85828601610965565b92505060206109b085828601610965565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109e5826109ba565b9050919050565b6109f5816109da565b8114610a0057600080fd5b50565b600081359050610a12816109ec565b92915050565b600060208284031215610a2e57610a2d61093f565b5b6000610a3c84828501610a03565b91505092915050565b610a4e81610944565b82525050565b604082016000820151610a6a6000850182610a45565b506020820151610a7d6020850182610a45565b50505050565b6000604082019050610a986000830184610a54565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ad3816109da565b82525050565b6000610ae58383610aca565b60208301905092915050565b6000602082019050919050565b6000610b0982610a9e565b610b138185610aa9565b9350610b1e83610aba565b8060005b83811015610b4f578151610b368882610ad9565b9750610b4183610af1565b925050600181019050610b22565b5085935050505092915050565b60006020820190508181036000830152610b768184610afe565b905092915050565b600082825260208201905092915050565b7f496e76616c6964207075626c6963206b65790000000000000000000000000000600082015250565b6000610bc5601283610b7e565b9150610bd082610b8f565b602082019050919050565b60006020820190508181036000830152610bf481610bb8565b9050919050565b7f4475706c696361746520656c656d656e74000000000000000000000000000000600082015250565b6000610c31601183610b7e565b9150610c3c82610bfb565b602082019050919050565b60006020820190508181036000830152610c6081610c24565b9050919050565b7f4e6f6e6578697374656e7420656c656d656e7400000000000000000000000000600082015250565b6000610c9d601383610b7e565b9150610ca882610c67565b602082019050919050565b60006020820190508181036000830152610ccc81610c90565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d1782610cd3565b9150610d2283610cd3565b925082821015610d3557610d34610cdd565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f496e76616c696420696e64657800000000000000000000000000000000000000600082015250565b6000610da5600d83610b7e565b9150610db082610d6f565b602082019050919050565b60006020820190508181036000830152610dd481610d98565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220a8333226837d875415a70c7a1835420739cc0f25e8f4c1ba8f788835b2e2dbbc64736f6c634300080e0033"

// DeployProviders deploys a new Ethereum contract, binding an instance of Providers to it.
func DeployProviders(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Providers, error) {
	parsed, err := abi.JSON(strings.NewReader(ProvidersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProvidersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Providers{ProvidersCaller: ProvidersCaller{contract: contract}, ProvidersTransactor: ProvidersTransactor{contract: contract}, ProvidersFilterer: ProvidersFilterer{contract: contract}}, nil
}

// Providers is an auto generated Go binding around an Ethereum contract.
type Providers struct {
	ProvidersCaller     // Read-only binding to the contract
	ProvidersTransactor // Write-only binding to the contract
	ProvidersFilterer   // Log filterer for contract events
}

// ProvidersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProvidersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvidersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProvidersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvidersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProvidersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvidersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProvidersSession struct {
	Contract     *Providers        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProvidersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProvidersCallerSession struct {
	Contract *ProvidersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ProvidersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProvidersTransactorSession struct {
	Contract     *ProvidersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProvidersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProvidersRaw struct {
	Contract *Providers // Generic contract binding to access the raw methods on
}

// ProvidersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProvidersCallerRaw struct {
	Contract *ProvidersCaller // Generic read-only contract binding to access the raw methods on
}

// ProvidersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProvidersTransactorRaw struct {
	Contract *ProvidersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProviders creates a new instance of Providers, bound to a specific deployed contract.
func NewProviders(address common.Address, backend bind.ContractBackend) (*Providers, error) {
	contract, err := bindProviders(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Providers{ProvidersCaller: ProvidersCaller{contract: contract}, ProvidersTransactor: ProvidersTransactor{contract: contract}, ProvidersFilterer: ProvidersFilterer{contract: contract}}, nil
}

// NewProvidersCaller creates a new read-only instance of Providers, bound to a specific deployed contract.
func NewProvidersCaller(address common.Address, caller bind.ContractCaller) (*ProvidersCaller, error) {
	contract, err := bindProviders(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProvidersCaller{contract: contract}, nil
}

// NewProvidersTransactor creates a new write-only instance of Providers, bound to a specific deployed contract.
func NewProvidersTransactor(address common.Address, transactor bind.ContractTransactor) (*ProvidersTransactor, error) {
	contract, err := bindProviders(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProvidersTransactor{contract: contract}, nil
}

// NewProvidersFilterer creates a new log filterer instance of Providers, bound to a specific deployed contract.
func NewProvidersFilterer(address common.Address, filterer bind.ContractFilterer) (*ProvidersFilterer, error) {
	contract, err := bindProviders(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProvidersFilterer{contract: contract}, nil
}

// bindProviders binds a generic wrapper to an already deployed contract.
func bindProviders(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProvidersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Providers *ProvidersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Providers.Contract.ProvidersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Providers *ProvidersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Providers.Contract.ProvidersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Providers *ProvidersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Providers.Contract.ProvidersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Providers *ProvidersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Providers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Providers *ProvidersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Providers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Providers *ProvidersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Providers.Contract.contract.Transact(opts, method, params...)
}

// GetProviders is a free data retrieval call binding the contract method 0xedc922a9.
//
// Solidity: function getProviders() view returns(address[])
func (_Providers *ProvidersCaller) GetProviders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Providers.contract.Call(opts, &out, "getProviders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetProviders is a free data retrieval call binding the contract method 0xedc922a9.
//
// Solidity: function getProviders() view returns(address[])
func (_Providers *ProvidersSession) GetProviders() ([]common.Address, error) {
	return _Providers.Contract.GetProviders(&_Providers.CallOpts)
}

// GetProviders is a free data retrieval call binding the contract method 0xedc922a9.
//
// Solidity: function getProviders() view returns(address[])
func (_Providers *ProvidersCallerSession) GetProviders() ([]common.Address, error) {
	return _Providers.Contract.GetProviders(&_Providers.CallOpts)
}

// Lookup is a free data retrieval call binding the contract method 0xd4b6b5da.
//
// Solidity: function lookup(address a) view returns((bytes32,bytes32))
func (_Providers *ProvidersCaller) Lookup(opts *bind.CallOpts, a common.Address) (ProviderNode, error) {
	var out []interface{}
	err := _Providers.contract.Call(opts, &out, "lookup", a)

	if err != nil {
		return *new(ProviderNode), err
	}

	out0 := *abi.ConvertType(out[0], new(ProviderNode)).(*ProviderNode)

	return out0, err

}

// Lookup is a free data retrieval call binding the contract method 0xd4b6b5da.
//
// Solidity: function lookup(address a) view returns((bytes32,bytes32))
func (_Providers *ProvidersSession) Lookup(a common.Address) (ProviderNode, error) {
	return _Providers.Contract.Lookup(&_Providers.CallOpts, a)
}

// Lookup is a free data retrieval call binding the contract method 0xd4b6b5da.
//
// Solidity: function lookup(address a) view returns((bytes32,bytes32))
func (_Providers *ProvidersCallerSession) Lookup(a common.Address) (ProviderNode, error) {
	return _Providers.Contract.Lookup(&_Providers.CallOpts, a)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Providers *ProvidersTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Providers.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Providers *ProvidersSession) Exit() (*types.Transaction, error) {
	return _Providers.Contract.Exit(&_Providers.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Providers *ProvidersTransactorSession) Exit() (*types.Transaction, error) {
	return _Providers.Contract.Exit(&_Providers.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0x5b419a65.
//
// Solidity: function join(bytes32 pubkey, bytes32 netAddr) returns()
func (_Providers *ProvidersTransactor) Join(opts *bind.TransactOpts, pubkey [32]byte, netAddr [32]byte) (*types.Transaction, error) {
	return _Providers.contract.Transact(opts, "join", pubkey, netAddr)
}

// Join is a paid mutator transaction binding the contract method 0x5b419a65.
//
// Solidity: function join(bytes32 pubkey, bytes32 netAddr) returns()
func (_Providers *ProvidersSession) Join(pubkey [32]byte, netAddr [32]byte) (*types.Transaction, error) {
	return _Providers.Contract.Join(&_Providers.TransactOpts, pubkey, netAddr)
}

// Join is a paid mutator transaction binding the contract method 0x5b419a65.
//
// Solidity: function join(bytes32 pubkey, bytes32 netAddr) returns()
func (_Providers *ProvidersTransactorSession) Join(pubkey [32]byte, netAddr [32]byte) (*types.Transaction, error) {
	return _Providers.Contract.Join(&_Providers.TransactOpts, pubkey, netAddr)
}

// ProvidersProviderExitIterator is returned from FilterProviderExit and is used to iterate over the raw logs and unpacked data for ProviderExit events raised by the Providers contract.
type ProvidersProviderExitIterator struct {
	Event *ProvidersProviderExit // Event containing the contract specifics and raw log

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
func (it *ProvidersProviderExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvidersProviderExit)
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
		it.Event = new(ProvidersProviderExit)
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
func (it *ProvidersProviderExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvidersProviderExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvidersProviderExit represents a ProviderExit event raised by the Providers contract.
type ProvidersProviderExit struct {
	P       common.Address
	Pubkey  [32]byte
	NetAddr [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProviderExit is a free log retrieval operation binding the contract event 0xc72f425864663e90861a0510ca302816b37bb1f9ab5203223088f8a4a25fb558.
//
// Solidity: event ProviderExit(address indexed p, bytes32 indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) FilterProviderExit(opts *bind.FilterOpts, p []common.Address, pubkey [][32]byte, netAddr [][32]byte) (*ProvidersProviderExitIterator, error) {

	var pRule []interface{}
	for _, pItem := range p {
		pRule = append(pRule, pItem)
	}
	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var netAddrRule []interface{}
	for _, netAddrItem := range netAddr {
		netAddrRule = append(netAddrRule, netAddrItem)
	}

	logs, sub, err := _Providers.contract.FilterLogs(opts, "ProviderExit", pRule, pubkeyRule, netAddrRule)
	if err != nil {
		return nil, err
	}
	return &ProvidersProviderExitIterator{contract: _Providers.contract, event: "ProviderExit", logs: logs, sub: sub}, nil
}

// WatchProviderExit is a free log subscription operation binding the contract event 0xc72f425864663e90861a0510ca302816b37bb1f9ab5203223088f8a4a25fb558.
//
// Solidity: event ProviderExit(address indexed p, bytes32 indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) WatchProviderExit(opts *bind.WatchOpts, sink chan<- *ProvidersProviderExit, p []common.Address, pubkey [][32]byte, netAddr [][32]byte) (event.Subscription, error) {

	var pRule []interface{}
	for _, pItem := range p {
		pRule = append(pRule, pItem)
	}
	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var netAddrRule []interface{}
	for _, netAddrItem := range netAddr {
		netAddrRule = append(netAddrRule, netAddrItem)
	}

	logs, sub, err := _Providers.contract.WatchLogs(opts, "ProviderExit", pRule, pubkeyRule, netAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvidersProviderExit)
				if err := _Providers.contract.UnpackLog(event, "ProviderExit", log); err != nil {
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

// ParseProviderExit is a log parse operation binding the contract event 0xc72f425864663e90861a0510ca302816b37bb1f9ab5203223088f8a4a25fb558.
//
// Solidity: event ProviderExit(address indexed p, bytes32 indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) ParseProviderExit(log types.Log) (*ProvidersProviderExit, error) {
	event := new(ProvidersProviderExit)
	if err := _Providers.contract.UnpackLog(event, "ProviderExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvidersProviderJoinIterator is returned from FilterProviderJoin and is used to iterate over the raw logs and unpacked data for ProviderJoin events raised by the Providers contract.
type ProvidersProviderJoinIterator struct {
	Event *ProvidersProviderJoin // Event containing the contract specifics and raw log

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
func (it *ProvidersProviderJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvidersProviderJoin)
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
		it.Event = new(ProvidersProviderJoin)
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
func (it *ProvidersProviderJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvidersProviderJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvidersProviderJoin represents a ProviderJoin event raised by the Providers contract.
type ProvidersProviderJoin struct {
	P       common.Address
	Pubkey  [32]byte
	NetAddr [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProviderJoin is a free log retrieval operation binding the contract event 0x15b88bb7d80477950c43dc4f6fd378c0dd62f99cffcc5f99f30dbd117c4f57ff.
//
// Solidity: event ProviderJoin(address indexed p, bytes32 indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) FilterProviderJoin(opts *bind.FilterOpts, p []common.Address, pubkey [][32]byte, netAddr [][32]byte) (*ProvidersProviderJoinIterator, error) {

	var pRule []interface{}
	for _, pItem := range p {
		pRule = append(pRule, pItem)
	}
	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var netAddrRule []interface{}
	for _, netAddrItem := range netAddr {
		netAddrRule = append(netAddrRule, netAddrItem)
	}

	logs, sub, err := _Providers.contract.FilterLogs(opts, "ProviderJoin", pRule, pubkeyRule, netAddrRule)
	if err != nil {
		return nil, err
	}
	return &ProvidersProviderJoinIterator{contract: _Providers.contract, event: "ProviderJoin", logs: logs, sub: sub}, nil
}

// WatchProviderJoin is a free log subscription operation binding the contract event 0x15b88bb7d80477950c43dc4f6fd378c0dd62f99cffcc5f99f30dbd117c4f57ff.
//
// Solidity: event ProviderJoin(address indexed p, bytes32 indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) WatchProviderJoin(opts *bind.WatchOpts, sink chan<- *ProvidersProviderJoin, p []common.Address, pubkey [][32]byte, netAddr [][32]byte) (event.Subscription, error) {

	var pRule []interface{}
	for _, pItem := range p {
		pRule = append(pRule, pItem)
	}
	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var netAddrRule []interface{}
	for _, netAddrItem := range netAddr {
		netAddrRule = append(netAddrRule, netAddrItem)
	}

	logs, sub, err := _Providers.contract.WatchLogs(opts, "ProviderJoin", pRule, pubkeyRule, netAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvidersProviderJoin)
				if err := _Providers.contract.UnpackLog(event, "ProviderJoin", log); err != nil {
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

// ParseProviderJoin is a log parse operation binding the contract event 0x15b88bb7d80477950c43dc4f6fd378c0dd62f99cffcc5f99f30dbd117c4f57ff.
//
// Solidity: event ProviderJoin(address indexed p, bytes32 indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) ParseProviderJoin(log types.Log) (*ProvidersProviderJoin, error) {
	event := new(ProvidersProviderJoin)
	if err := _Providers.contract.UnpackLog(event, "ProviderJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnorderedSetABI is the input ABI used to generate the binding from.
const UnorderedSetABI = "[]"

// UnorderedSetBin is the compiled bytecode used for deploying new contracts.
var UnorderedSetBin = "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c8d422dae4808943c827a278a50d5310221a6778b6e431d4e39742d36e235b0764736f6c634300080e0033"

// DeployUnorderedSet deploys a new Ethereum contract, binding an instance of UnorderedSet to it.
func DeployUnorderedSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnorderedSet, error) {
	parsed, err := abi.JSON(strings.NewReader(UnorderedSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnorderedSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnorderedSet{UnorderedSetCaller: UnorderedSetCaller{contract: contract}, UnorderedSetTransactor: UnorderedSetTransactor{contract: contract}, UnorderedSetFilterer: UnorderedSetFilterer{contract: contract}}, nil
}

// UnorderedSet is an auto generated Go binding around an Ethereum contract.
type UnorderedSet struct {
	UnorderedSetCaller     // Read-only binding to the contract
	UnorderedSetTransactor // Write-only binding to the contract
	UnorderedSetFilterer   // Log filterer for contract events
}

// UnorderedSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnorderedSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnorderedSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnorderedSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnorderedSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnorderedSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnorderedSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnorderedSetSession struct {
	Contract     *UnorderedSet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnorderedSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnorderedSetCallerSession struct {
	Contract *UnorderedSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UnorderedSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnorderedSetTransactorSession struct {
	Contract     *UnorderedSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UnorderedSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnorderedSetRaw struct {
	Contract *UnorderedSet // Generic contract binding to access the raw methods on
}

// UnorderedSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnorderedSetCallerRaw struct {
	Contract *UnorderedSetCaller // Generic read-only contract binding to access the raw methods on
}

// UnorderedSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnorderedSetTransactorRaw struct {
	Contract *UnorderedSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnorderedSet creates a new instance of UnorderedSet, bound to a specific deployed contract.
func NewUnorderedSet(address common.Address, backend bind.ContractBackend) (*UnorderedSet, error) {
	contract, err := bindUnorderedSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnorderedSet{UnorderedSetCaller: UnorderedSetCaller{contract: contract}, UnorderedSetTransactor: UnorderedSetTransactor{contract: contract}, UnorderedSetFilterer: UnorderedSetFilterer{contract: contract}}, nil
}

// NewUnorderedSetCaller creates a new read-only instance of UnorderedSet, bound to a specific deployed contract.
func NewUnorderedSetCaller(address common.Address, caller bind.ContractCaller) (*UnorderedSetCaller, error) {
	contract, err := bindUnorderedSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnorderedSetCaller{contract: contract}, nil
}

// NewUnorderedSetTransactor creates a new write-only instance of UnorderedSet, bound to a specific deployed contract.
func NewUnorderedSetTransactor(address common.Address, transactor bind.ContractTransactor) (*UnorderedSetTransactor, error) {
	contract, err := bindUnorderedSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnorderedSetTransactor{contract: contract}, nil
}

// NewUnorderedSetFilterer creates a new log filterer instance of UnorderedSet, bound to a specific deployed contract.
func NewUnorderedSetFilterer(address common.Address, filterer bind.ContractFilterer) (*UnorderedSetFilterer, error) {
	contract, err := bindUnorderedSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnorderedSetFilterer{contract: contract}, nil
}

// bindUnorderedSet binds a generic wrapper to an already deployed contract.
func bindUnorderedSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnorderedSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnorderedSet *UnorderedSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnorderedSet.Contract.UnorderedSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnorderedSet *UnorderedSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnorderedSet.Contract.UnorderedSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnorderedSet *UnorderedSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnorderedSet.Contract.UnorderedSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnorderedSet *UnorderedSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnorderedSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnorderedSet *UnorderedSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnorderedSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnorderedSet *UnorderedSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnorderedSet.Contract.contract.Transact(opts, method, params...)
}
