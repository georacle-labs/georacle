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

// ProvidersProvider is an auto generated low-level Go binding around an user-defined struct.
type ProvidersProvider struct {
	Pubkey  []byte
	NetAddr [32]byte
}

// ProvidersABI is the input ABI used to generate the binding from.
const ProvidersABI = "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structProviders.Provider\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"ProviderExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structProviders.Provider\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"ProviderJoin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Oracles\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"netAddr\",\"type\":\"bytes32\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ProvidersBin is the compiled bytecode used for deploying new contracts.
var ProvidersBin = "0x608060405234801561001057600080fd5b50610901806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063127bdf4a14610046578063cbebd1e114610062578063e9fad8ee14610093575b600080fd5b610060600480360381019061005b9190610605565b61009d565b005b61007c600480360381019061007791906106bf565b61015b565b60405161008a929190610783565b60405180910390f35b61009b610207565b005b60006040518060400160405280848152602001838152509050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019080519060200190610111929190610392565b50602082015181600101559050507f95153673f75302857771bc6cf6ad1ca5f17342f2dd9f16411147be957a67074a8160405161014e9190610849565b60405180910390a1505050565b600060205280600052604060002060009150905080600001805461017e9061089a565b80601f01602080910402602001604051908101604052809291908181526020018280546101aa9061089a565b80156101f75780601f106101cc576101008083540402835291602001916101f7565b820191906000526020600020905b8154815290600101906020018083116101da57829003601f168201915b5050505050908060010154905082565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820180546102629061089a565b80601f016020809104026020016040519081016040528092919081815260200182805461028e9061089a565b80156102db5780601f106102b0576101008083540402835291602001916102db565b820191906000526020600020905b8154815290600101906020018083116102be57829003601f168201915b5050505050815260200160018201548152505090506000801b81602001511461038f576000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600061034d9190610418565b600182016000905550507f363758eb8b472af706762400ca9bce45c190aa4eeb38196ea7d083373b117f00816040516103869190610849565b60405180910390a15b50565b82805461039e9061089a565b90600052602060002090601f0160209004810192826103c05760008555610407565b82601f106103d957805160ff1916838001178555610407565b82800160010185558215610407579182015b828111156104065782518255916020019190600101906103eb565b5b5090506104149190610458565b5090565b5080546104249061089a565b6000825580601f106104365750610455565b601f0160209004906000526020600020908101906104549190610458565b5b50565b5b80821115610471576000816000905550600101610459565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104dc82610493565b810181811067ffffffffffffffff821117156104fb576104fa6104a4565b5b80604052505050565b600061050e610475565b905061051a82826104d3565b919050565b600067ffffffffffffffff82111561053a576105396104a4565b5b61054382610493565b9050602081019050919050565b82818337600083830152505050565b600061057261056d8461051f565b610504565b90508281526020810184848401111561058e5761058d61048e565b5b610599848285610550565b509392505050565b600082601f8301126105b6576105b5610489565b5b81356105c684826020860161055f565b91505092915050565b6000819050919050565b6105e2816105cf565b81146105ed57600080fd5b50565b6000813590506105ff816105d9565b92915050565b6000806040838503121561061c5761061b61047f565b5b600083013567ffffffffffffffff81111561063a57610639610484565b5b610646858286016105a1565b9250506020610657858286016105f0565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061068c82610661565b9050919050565b61069c81610681565b81146106a757600080fd5b50565b6000813590506106b981610693565b92915050565b6000602082840312156106d5576106d461047f565b5b60006106e3848285016106aa565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561072657808201518184015260208101905061070b565b83811115610735576000848401525b50505050565b6000610746826106ec565b61075081856106f7565b9350610760818560208601610708565b61076981610493565b840191505092915050565b61077d816105cf565b82525050565b6000604082019050818103600083015261079d818561073b565b90506107ac6020830184610774565b9392505050565b600082825260208201905092915050565b60006107cf826106ec565b6107d981856107b3565b93506107e9818560208601610708565b6107f281610493565b840191505092915050565b610806816105cf565b82525050565b6000604083016000830151848203600086015261082982826107c4565b915050602083015161083e60208601826107fd565b508091505092915050565b60006020820190508181036000830152610863818461080c565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806108b257607f821691505b6020821081036108c5576108c461086b565b5b5091905056fea26469706673582212208c8428dcbe4099b495608cc9ece6db9208d0ed1df2cbc4acaff26253c0bc861264736f6c634300080d0033"

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

// Oracles is a free data retrieval call binding the contract method 0xcbebd1e1.
//
// Solidity: function Oracles(address ) view returns(bytes pubkey, bytes32 netAddr)
func (_Providers *ProvidersCaller) Oracles(opts *bind.CallOpts, arg0 common.Address) (struct {
	Pubkey  []byte
	NetAddr [32]byte
}, error) {
	var out []interface{}
	err := _Providers.contract.Call(opts, &out, "Oracles", arg0)

	outstruct := new(struct {
		Pubkey  []byte
		NetAddr [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.NetAddr = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Oracles is a free data retrieval call binding the contract method 0xcbebd1e1.
//
// Solidity: function Oracles(address ) view returns(bytes pubkey, bytes32 netAddr)
func (_Providers *ProvidersSession) Oracles(arg0 common.Address) (struct {
	Pubkey  []byte
	NetAddr [32]byte
}, error) {
	return _Providers.Contract.Oracles(&_Providers.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xcbebd1e1.
//
// Solidity: function Oracles(address ) view returns(bytes pubkey, bytes32 netAddr)
func (_Providers *ProvidersCallerSession) Oracles(arg0 common.Address) (struct {
	Pubkey  []byte
	NetAddr [32]byte
}, error) {
	return _Providers.Contract.Oracles(&_Providers.CallOpts, arg0)
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
	P   ProvidersProvider
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProviderExit is a free log retrieval operation binding the contract event 0x363758eb8b472af706762400ca9bce45c190aa4eeb38196ea7d083373b117f00.
//
// Solidity: event ProviderExit((bytes,bytes32) p)
func (_Providers *ProvidersFilterer) FilterProviderExit(opts *bind.FilterOpts) (*ProvidersProviderExitIterator, error) {

	logs, sub, err := _Providers.contract.FilterLogs(opts, "ProviderExit")
	if err != nil {
		return nil, err
	}
	return &ProvidersProviderExitIterator{contract: _Providers.contract, event: "ProviderExit", logs: logs, sub: sub}, nil
}

// WatchProviderExit is a free log subscription operation binding the contract event 0x363758eb8b472af706762400ca9bce45c190aa4eeb38196ea7d083373b117f00.
//
// Solidity: event ProviderExit((bytes,bytes32) p)
func (_Providers *ProvidersFilterer) WatchProviderExit(opts *bind.WatchOpts, sink chan<- *ProvidersProviderExit) (event.Subscription, error) {

	logs, sub, err := _Providers.contract.WatchLogs(opts, "ProviderExit")
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

// ParseProviderExit is a log parse operation binding the contract event 0x363758eb8b472af706762400ca9bce45c190aa4eeb38196ea7d083373b117f00.
//
// Solidity: event ProviderExit((bytes,bytes32) p)
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
	P   ProvidersProvider
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProviderJoin is a free log retrieval operation binding the contract event 0x95153673f75302857771bc6cf6ad1ca5f17342f2dd9f16411147be957a67074a.
//
// Solidity: event ProviderJoin((bytes,bytes32) p)
func (_Providers *ProvidersFilterer) FilterProviderJoin(opts *bind.FilterOpts) (*ProvidersProviderJoinIterator, error) {

	logs, sub, err := _Providers.contract.FilterLogs(opts, "ProviderJoin")
	if err != nil {
		return nil, err
	}
	return &ProvidersProviderJoinIterator{contract: _Providers.contract, event: "ProviderJoin", logs: logs, sub: sub}, nil
}

// WatchProviderJoin is a free log subscription operation binding the contract event 0x95153673f75302857771bc6cf6ad1ca5f17342f2dd9f16411147be957a67074a.
//
// Solidity: event ProviderJoin((bytes,bytes32) p)
func (_Providers *ProvidersFilterer) WatchProviderJoin(opts *bind.WatchOpts, sink chan<- *ProvidersProviderJoin) (event.Subscription, error) {

	logs, sub, err := _Providers.contract.WatchLogs(opts, "ProviderJoin")
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

// ParseProviderJoin is a log parse operation binding the contract event 0x95153673f75302857771bc6cf6ad1ca5f17342f2dd9f16411147be957a67074a.
//
// Solidity: event ProviderJoin((bytes,bytes32) p)
func (_Providers *ProvidersFilterer) ParseProviderJoin(log types.Log) (*ProvidersProviderJoin, error) {
	event := new(ProvidersProviderJoin)
	if err := _Providers.contract.UnpackLog(event, "ProviderJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
