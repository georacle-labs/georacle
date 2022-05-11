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
	Pubkey  []byte
	NetAddr [32]byte
}

// ProviderABI is the input ABI used to generate the binding from.
const ProviderABI = "[]"

// ProviderBin is the compiled bytecode used for deploying new contracts.
var ProviderBin = "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122091822ce63c1bdd66bae6e93c70cff8a111c36d9e72c28473ace13b3a9a52c57064736f6c634300080d0033"

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
const ProvidersABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"ProviderExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"name\":\"ProviderJoin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"lookup\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"internalType\":\"structProvider.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProvidersBin is the compiled bytecode used for deploying new contracts.
var ProvidersBin = "0x608060405234801561001057600080fd5b50611254806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063127bdf4a14610051578063d4b6b5da1461006d578063e9fad8ee1461009d578063edc922a9146100a7575b600080fd5b61006b60048036038101906100669190610c8e565b6100c5565b005b61008760048036038101906100829190610d48565b61013b565b6040516100949190610e49565b60405180910390f35b6100a5610205565b005b6100af61025e565b6040516100bc9190610f29565b60405180910390f35b6100dd33838360006102f4909392919063ffffffff16565b80826040516100ec9190610f87565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167f879fcd154725a29470ab6900655c940fab17573a678a4ddfaaf9262a3759014560405160405180910390a45050565b6101436109fe565b6101578260006104a590919063ffffffff16565b60405180604001604052908160008201805461017290610fcd565b80601f016020809104026020016040519081016040528092919081815260200182805461019e90610fcd565b80156101eb5780601f106101c0576101008083540402835291602001916101eb565b820191906000526020600020905b8154815290600101906020018083116101ce57829003601f168201915b505050505081526020016001820154815250509050919050565b6102193360006104f390919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167fafb4f5373b6024003a3099e27a07247ab27cdb54951710a3757eea1efaf8e4ec60405160405180910390a2565b606061026a600061073b565b8054806020026020016040519081016040528092919081815260200182805480156102ea57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116102a0575b5050505050905090565b60008460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600001805461034890610fcd565b90501461038a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103819061105b565b60405180910390fd5b6040518060400160405280856001018054905081526020016040518060400160405280858152602001848152508152508460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101600082015181600001908051906020019061042a929190610a1b565b5060208201518160010155505090505083600101839080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101905092915050565b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806040016040529081600082015481526020016001820160405180604001604052908160008201805461056b90610fcd565b80601f016020809104026020016040519081016040528092919081815260200182805461059790610fcd565b80156105e45780601f106105b9576101008083540402835291602001916105e4565b820191906000526020600020905b8154815290600101906020018083116105c757829003601f168201915b5050505050815260200160018201548152505081525050905060008160200151600001515114158015610621575082600101805490508160000151105b610660576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610657906110c7565b60405180910390fd5b6106828382600001516001866001018054905061067d9190611120565b610748565b8260010180548061069657610695611154565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590558260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160009055600182016000808201600061072a9190610aa1565b600182016000905550505050505050565b6000816001019050919050565b8260010180549050821080156107645750826001018054905081105b6107a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079a906111cf565b60405180910390fd5b60008360010183815481106107bb576107ba6111ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508360010182815481106107fe576107fd6111ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600101848154811061083f5761083e6111ef565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508084600101838154811061089e5761089d6111ef565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081846000016000866001018681548110610903576109026111ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055508284600001600086600101858154811061098c5761098b6111ef565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555050505050565b604051806040016040528060608152602001600080191681525090565b828054610a2790610fcd565b90600052602060002090601f016020900481019282610a495760008555610a90565b82601f10610a6257805160ff1916838001178555610a90565b82800160010185558215610a90579182015b82811115610a8f578251825591602001919060010190610a74565b5b509050610a9d9190610ae1565b5090565b508054610aad90610fcd565b6000825580601f10610abf5750610ade565b601f016020900490600052602060002090810190610add9190610ae1565b5b50565b5b80821115610afa576000816000905550600101610ae2565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b6582610b1c565b810181811067ffffffffffffffff82111715610b8457610b83610b2d565b5b80604052505050565b6000610b97610afe565b9050610ba38282610b5c565b919050565b600067ffffffffffffffff821115610bc357610bc2610b2d565b5b610bcc82610b1c565b9050602081019050919050565b82818337600083830152505050565b6000610bfb610bf684610ba8565b610b8d565b905082815260208101848484011115610c1757610c16610b17565b5b610c22848285610bd9565b509392505050565b600082601f830112610c3f57610c3e610b12565b5b8135610c4f848260208601610be8565b91505092915050565b6000819050919050565b610c6b81610c58565b8114610c7657600080fd5b50565b600081359050610c8881610c62565b92915050565b60008060408385031215610ca557610ca4610b08565b5b600083013567ffffffffffffffff811115610cc357610cc2610b0d565b5b610ccf85828601610c2a565b9250506020610ce085828601610c79565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d1582610cea565b9050919050565b610d2581610d0a565b8114610d3057600080fd5b50565b600081359050610d4281610d1c565b92915050565b600060208284031215610d5e57610d5d610b08565b5b6000610d6c84828501610d33565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610daf578082015181840152602081019050610d94565b83811115610dbe576000848401525b50505050565b6000610dcf82610d75565b610dd98185610d80565b9350610de9818560208601610d91565b610df281610b1c565b840191505092915050565b610e0681610c58565b82525050565b60006040830160008301518482036000860152610e298282610dc4565b9150506020830151610e3e6020860182610dfd565b508091505092915050565b60006020820190508181036000830152610e638184610e0c565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ea081610d0a565b82525050565b6000610eb28383610e97565b60208301905092915050565b6000602082019050919050565b6000610ed682610e6b565b610ee08185610e76565b9350610eeb83610e87565b8060005b83811015610f1c578151610f038882610ea6565b9750610f0e83610ebe565b925050600181019050610eef565b5085935050505092915050565b60006020820190508181036000830152610f438184610ecb565b905092915050565b600081905092915050565b6000610f6182610d75565b610f6b8185610f4b565b9350610f7b818560208601610d91565b80840191505092915050565b6000610f938284610f56565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610fe557607f821691505b602082108103610ff857610ff7610f9e565b5b50919050565b600082825260208201905092915050565b7f4475706c696361746520656c656d656e74000000000000000000000000000000600082015250565b6000611045601183610ffe565b91506110508261100f565b602082019050919050565b6000602082019050818103600083015261107481611038565b9050919050565b7f4e6f6e6578697374656e7420656c656d656e7400000000000000000000000000600082015250565b60006110b1601383610ffe565b91506110bc8261107b565b602082019050919050565b600060208201905081810360008301526110e0816110a4565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061112b826110e7565b9150611136836110e7565b925082821015611149576111486110f1565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f496e76616c696420696e64657800000000000000000000000000000000000000600082015250565b60006111b9600d83610ffe565b91506111c482611183565b602082019050919050565b600060208201905081810360008301526111e8816111ac565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212202359824ebb6fa946c9a2fb3e0c7517de59d7d38da1de816d893ab3a5e4f26aa364736f6c634300080d0033"

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
// Solidity: function lookup(address a) view returns((bytes,bytes32))
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
// Solidity: function lookup(address a) view returns((bytes,bytes32))
func (_Providers *ProvidersSession) Lookup(a common.Address) (ProviderNode, error) {
	return _Providers.Contract.Lookup(&_Providers.CallOpts, a)
}

// Lookup is a free data retrieval call binding the contract method 0xd4b6b5da.
//
// Solidity: function lookup(address a) view returns((bytes,bytes32))
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

// Join is a paid mutator transaction binding the contract method 0x127bdf4a.
//
// Solidity: function join(bytes pubkey, bytes32 netAddr) returns()
func (_Providers *ProvidersTransactor) Join(opts *bind.TransactOpts, pubkey []byte, netAddr [32]byte) (*types.Transaction, error) {
	return _Providers.contract.Transact(opts, "join", pubkey, netAddr)
}

// Join is a paid mutator transaction binding the contract method 0x127bdf4a.
//
// Solidity: function join(bytes pubkey, bytes32 netAddr) returns()
func (_Providers *ProvidersSession) Join(pubkey []byte, netAddr [32]byte) (*types.Transaction, error) {
	return _Providers.Contract.Join(&_Providers.TransactOpts, pubkey, netAddr)
}

// Join is a paid mutator transaction binding the contract method 0x127bdf4a.
//
// Solidity: function join(bytes pubkey, bytes32 netAddr) returns()
func (_Providers *ProvidersTransactorSession) Join(pubkey []byte, netAddr [32]byte) (*types.Transaction, error) {
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
	P   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProviderExit is a free log retrieval operation binding the contract event 0xafb4f5373b6024003a3099e27a07247ab27cdb54951710a3757eea1efaf8e4ec.
//
// Solidity: event ProviderExit(address indexed p)
func (_Providers *ProvidersFilterer) FilterProviderExit(opts *bind.FilterOpts, p []common.Address) (*ProvidersProviderExitIterator, error) {

	var pRule []interface{}
	for _, pItem := range p {
		pRule = append(pRule, pItem)
	}

	logs, sub, err := _Providers.contract.FilterLogs(opts, "ProviderExit", pRule)
	if err != nil {
		return nil, err
	}
	return &ProvidersProviderExitIterator{contract: _Providers.contract, event: "ProviderExit", logs: logs, sub: sub}, nil
}

// WatchProviderExit is a free log subscription operation binding the contract event 0xafb4f5373b6024003a3099e27a07247ab27cdb54951710a3757eea1efaf8e4ec.
//
// Solidity: event ProviderExit(address indexed p)
func (_Providers *ProvidersFilterer) WatchProviderExit(opts *bind.WatchOpts, sink chan<- *ProvidersProviderExit, p []common.Address) (event.Subscription, error) {

	var pRule []interface{}
	for _, pItem := range p {
		pRule = append(pRule, pItem)
	}

	logs, sub, err := _Providers.contract.WatchLogs(opts, "ProviderExit", pRule)
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

// ParseProviderExit is a log parse operation binding the contract event 0xafb4f5373b6024003a3099e27a07247ab27cdb54951710a3757eea1efaf8e4ec.
//
// Solidity: event ProviderExit(address indexed p)
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
	Pubkey  common.Hash
	NetAddr [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProviderJoin is a free log retrieval operation binding the contract event 0x879fcd154725a29470ab6900655c940fab17573a678a4ddfaaf9262a37590145.
//
// Solidity: event ProviderJoin(address indexed p, bytes indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) FilterProviderJoin(opts *bind.FilterOpts, p []common.Address, pubkey [][]byte, netAddr [][32]byte) (*ProvidersProviderJoinIterator, error) {

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

// WatchProviderJoin is a free log subscription operation binding the contract event 0x879fcd154725a29470ab6900655c940fab17573a678a4ddfaaf9262a37590145.
//
// Solidity: event ProviderJoin(address indexed p, bytes indexed pubkey, bytes32 indexed netAddr)
func (_Providers *ProvidersFilterer) WatchProviderJoin(opts *bind.WatchOpts, sink chan<- *ProvidersProviderJoin, p []common.Address, pubkey [][]byte, netAddr [][32]byte) (event.Subscription, error) {

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

// ParseProviderJoin is a log parse operation binding the contract event 0x879fcd154725a29470ab6900655c940fab17573a678a4ddfaaf9262a37590145.
//
// Solidity: event ProviderJoin(address indexed p, bytes indexed pubkey, bytes32 indexed netAddr)
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
var UnorderedSetBin = "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b81839c9a56dc5b4809a947dc55c0402b0117517f61229630601626b62b1c49264736f6c634300080d0033"

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
