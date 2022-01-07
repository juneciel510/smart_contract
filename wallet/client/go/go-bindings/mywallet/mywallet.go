// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// MyWalletMetaData contains all meta data concerning the MyWallet contract.
var MyWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_weiAmount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_weiAmount\",\"type\":\"uint256\"}],\"name\":\"Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_weiAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d0e30db0": "deposit()",
		"12065fe0": "getBalance()",
		"8da5cb5b": "owner()",
		"a9059cbb": "transfer(address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610468806100326000396000f3fe60806040526004361061004a5760003560e01c806312065fe01461004f5780632e1a7d4d146100725780638da5cb5b14610094578063a9059cbb146100cc578063d0e30db0146100ec575b600080fd5b34801561005b57600080fd5b506001546040519081526020015b60405180910390f35b34801561007e57600080fd5b5061009261008d36600461039c565b6100f4565b005b3480156100a057600080fd5b506000546100b4906001600160a01b031681565b6040516001600160a01b039091168152602001610069565b3480156100d857600080fd5b506100926100e73660046103b5565b610215565b610092610351565b6000546001600160a01b0316331461014d5760405162461bcd60e51b815260206004820152601760248201527639b2b73232b91034b9903737ba103a34329037bbb732b960491b60448201526064015b60405180910390fd5b8060015410156101945760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b6044820152606401610144565b604051339082156108fc029083906000818181858888f193505050501580156101c1573d6000803e3d6000fd5b5080600160008282546101d49190610403565b909155505060408051338152602081018390527f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5910160405180910390a150565b6000546001600160a01b031633146102695760405162461bcd60e51b815260206004820152601760248201527639b2b73232b91034b9903737ba103a34329037bbb732b960491b6044820152606401610144565b8061027360015490565b10156102b65760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b6044820152606401610144565b6040516001600160a01b0383169082156108fc029083906000818181858888f193505050501580156102ec573d6000803e3d6000fd5b5080600160008282546102ff9190610403565b9091555050604080513381526001600160a01b03841660208201529081018290527fd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee9060600160405180910390a15050565b3460015461035f919061041a565b600155604080513381523460208201527f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4910160405180910390a1565b6000602082840312156103ae57600080fd5b5035919050565b600080604083850312156103c857600080fd5b82356001600160a01b03811681146103df57600080fd5b946020939093013593505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610415576104156103ed565b500390565b6000821982111561042d5761042d6103ed565b50019056fea264697066735822122051e136e831623f2981c61095b82036c1c50bdb1e07ee182dc481de87f10fb14964736f6c63430008090033",
}

// MyWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use MyWalletMetaData.ABI instead.
var MyWalletABI = MyWalletMetaData.ABI

// Deprecated: Use MyWalletMetaData.Sigs instead.
// MyWalletFuncSigs maps the 4-byte function signature to its string representation.
var MyWalletFuncSigs = MyWalletMetaData.Sigs

// MyWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyWalletMetaData.Bin instead.
var MyWalletBin = MyWalletMetaData.Bin

// DeployMyWallet deploys a new Ethereum contract, binding an instance of MyWallet to it.
func DeployMyWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyWallet, error) {
	parsed, err := MyWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyWallet{MyWalletCaller: MyWalletCaller{contract: contract}, MyWalletTransactor: MyWalletTransactor{contract: contract}, MyWalletFilterer: MyWalletFilterer{contract: contract}}, nil
}

// MyWallet is an auto generated Go binding around an Ethereum contract.
type MyWallet struct {
	MyWalletCaller     // Read-only binding to the contract
	MyWalletTransactor // Write-only binding to the contract
	MyWalletFilterer   // Log filterer for contract events
}

// MyWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyWalletSession struct {
	Contract     *MyWallet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyWalletCallerSession struct {
	Contract *MyWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MyWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyWalletTransactorSession struct {
	Contract     *MyWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MyWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyWalletRaw struct {
	Contract *MyWallet // Generic contract binding to access the raw methods on
}

// MyWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyWalletCallerRaw struct {
	Contract *MyWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MyWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyWalletTransactorRaw struct {
	Contract *MyWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyWallet creates a new instance of MyWallet, bound to a specific deployed contract.
func NewMyWallet(address common.Address, backend bind.ContractBackend) (*MyWallet, error) {
	contract, err := bindMyWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyWallet{MyWalletCaller: MyWalletCaller{contract: contract}, MyWalletTransactor: MyWalletTransactor{contract: contract}, MyWalletFilterer: MyWalletFilterer{contract: contract}}, nil
}

// NewMyWalletCaller creates a new read-only instance of MyWallet, bound to a specific deployed contract.
func NewMyWalletCaller(address common.Address, caller bind.ContractCaller) (*MyWalletCaller, error) {
	contract, err := bindMyWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyWalletCaller{contract: contract}, nil
}

// NewMyWalletTransactor creates a new write-only instance of MyWallet, bound to a specific deployed contract.
func NewMyWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MyWalletTransactor, error) {
	contract, err := bindMyWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyWalletTransactor{contract: contract}, nil
}

// NewMyWalletFilterer creates a new log filterer instance of MyWallet, bound to a specific deployed contract.
func NewMyWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MyWalletFilterer, error) {
	contract, err := bindMyWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyWalletFilterer{contract: contract}, nil
}

// bindMyWallet binds a generic wrapper to an already deployed contract.
func bindMyWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyWallet *MyWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyWallet.Contract.MyWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyWallet *MyWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyWallet.Contract.MyWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyWallet *MyWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyWallet.Contract.MyWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyWallet *MyWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyWallet *MyWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyWallet *MyWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyWallet.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_MyWallet *MyWalletCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyWallet.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_MyWallet *MyWalletSession) GetBalance() (*big.Int, error) {
	return _MyWallet.Contract.GetBalance(&_MyWallet.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_MyWallet *MyWalletCallerSession) GetBalance() (*big.Int, error) {
	return _MyWallet.Contract.GetBalance(&_MyWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyWallet *MyWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyWallet *MyWalletSession) Owner() (common.Address, error) {
	return _MyWallet.Contract.Owner(&_MyWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyWallet *MyWalletCallerSession) Owner() (common.Address, error) {
	return _MyWallet.Contract.Owner(&_MyWallet.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_MyWallet *MyWalletTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyWallet.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_MyWallet *MyWalletSession) Deposit() (*types.Transaction, error) {
	return _MyWallet.Contract.Deposit(&_MyWallet.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_MyWallet *MyWalletTransactorSession) Deposit() (*types.Transaction, error) {
	return _MyWallet.Contract.Deposit(&_MyWallet.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 weiAmount) returns()
func (_MyWallet *MyWalletTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, weiAmount *big.Int) (*types.Transaction, error) {
	return _MyWallet.contract.Transact(opts, "transfer", recipient, weiAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 weiAmount) returns()
func (_MyWallet *MyWalletSession) Transfer(recipient common.Address, weiAmount *big.Int) (*types.Transaction, error) {
	return _MyWallet.Contract.Transfer(&_MyWallet.TransactOpts, recipient, weiAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 weiAmount) returns()
func (_MyWallet *MyWalletTransactorSession) Transfer(recipient common.Address, weiAmount *big.Int) (*types.Transaction, error) {
	return _MyWallet.Contract.Transfer(&_MyWallet.TransactOpts, recipient, weiAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 weiAmount) returns()
func (_MyWallet *MyWalletTransactor) Withdraw(opts *bind.TransactOpts, weiAmount *big.Int) (*types.Transaction, error) {
	return _MyWallet.contract.Transact(opts, "withdraw", weiAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 weiAmount) returns()
func (_MyWallet *MyWalletSession) Withdraw(weiAmount *big.Int) (*types.Transaction, error) {
	return _MyWallet.Contract.Withdraw(&_MyWallet.TransactOpts, weiAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 weiAmount) returns()
func (_MyWallet *MyWalletTransactorSession) Withdraw(weiAmount *big.Int) (*types.Transaction, error) {
	return _MyWallet.Contract.Withdraw(&_MyWallet.TransactOpts, weiAmount)
}

// MyWalletDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the MyWallet contract.
type MyWalletDepositedIterator struct {
	Event *MyWalletDeposited // Event containing the contract specifics and raw log

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
func (it *MyWalletDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyWalletDeposited)
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
		it.Event = new(MyWalletDeposited)
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
func (it *MyWalletDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyWalletDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyWalletDeposited represents a Deposited event raised by the MyWallet contract.
type MyWalletDeposited struct {
	From      common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _from, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) FilterDeposited(opts *bind.FilterOpts) (*MyWalletDepositedIterator, error) {

	logs, sub, err := _MyWallet.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &MyWalletDepositedIterator{contract: _MyWallet.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _from, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *MyWalletDeposited) (event.Subscription, error) {

	logs, sub, err := _MyWallet.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyWalletDeposited)
				if err := _MyWallet.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _from, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) ParseDeposited(log types.Log) (*MyWalletDeposited, error) {
	event := new(MyWalletDeposited)
	if err := _MyWallet.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyWalletTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the MyWallet contract.
type MyWalletTransferredIterator struct {
	Event *MyWalletTransferred // Event containing the contract specifics and raw log

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
func (it *MyWalletTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyWalletTransferred)
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
		it.Event = new(MyWalletTransferred)
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
func (it *MyWalletTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyWalletTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyWalletTransferred represents a Transferred event raised by the MyWallet contract.
type MyWalletTransferred struct {
	From      common.Address
	To        common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _from, address _to, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) FilterTransferred(opts *bind.FilterOpts) (*MyWalletTransferredIterator, error) {

	logs, sub, err := _MyWallet.contract.FilterLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return &MyWalletTransferredIterator{contract: _MyWallet.contract, event: "Transferred", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _from, address _to, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *MyWalletTransferred) (event.Subscription, error) {

	logs, sub, err := _MyWallet.contract.WatchLogs(opts, "Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyWalletTransferred)
				if err := _MyWallet.contract.UnpackLog(event, "Transferred", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xd1ba4ac2e2a11b5101f6cb4d978f514a155b421e8ec396d2d9abaf0bb02917ee.
//
// Solidity: event Transferred(address _from, address _to, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) ParseTransferred(log types.Log) (*MyWalletTransferred, error) {
	event := new(MyWalletTransferred)
	if err := _MyWallet.contract.UnpackLog(event, "Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyWalletWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the MyWallet contract.
type MyWalletWithdrawnIterator struct {
	Event *MyWalletWithdrawn // Event containing the contract specifics and raw log

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
func (it *MyWalletWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyWalletWithdrawn)
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
		it.Event = new(MyWalletWithdrawn)
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
func (it *MyWalletWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyWalletWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyWalletWithdrawn represents a Withdrawn event raised by the MyWallet contract.
type MyWalletWithdrawn struct {
	To        common.Address
	WeiAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _to, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*MyWalletWithdrawnIterator, error) {

	logs, sub, err := _MyWallet.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &MyWalletWithdrawnIterator{contract: _MyWallet.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _to, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *MyWalletWithdrawn) (event.Subscription, error) {

	logs, sub, err := _MyWallet.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyWalletWithdrawn)
				if err := _MyWallet.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _to, uint256 _weiAmount)
func (_MyWallet *MyWalletFilterer) ParseWithdrawn(log types.Log) (*MyWalletWithdrawn, error) {
	event := new(MyWalletWithdrawn)
	if err := _MyWallet.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
