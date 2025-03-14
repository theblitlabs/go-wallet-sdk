package walletsdk

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// StakeWalletContractABI is the input ABI used to generate the binding from.
const StakeWalletContractABI = `[
    {
      "type": "constructor",
      "inputs": [
        {
          "name": "_tokenAddress",
          "type": "address",
          "internalType": "address"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "addFunds",
      "inputs": [
        {
          "name": "_amount",
          "type": "uint256",
          "internalType": "uint256"
        },
        {
          "name": "_deviceId",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "_walletAddress",
          "type": "address",
          "internalType": "address"
        }
      ],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "getBalance",
      "inputs": [
        {
          "name": "_deviceId",
          "type": "string",
          "internalType": "string"
        }
      ],
      "outputs": [
        {
          "name": "",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "stateMutability": "view"
    },
    {
      "type": "function",
      "name": "getWalletInfo",
      "inputs": [
        {
          "name": "_deviceId",
          "type": "string",
          "internalType": "string"
        }
      ],
      "outputs": [
        {
          "name": "balance",
          "type": "uint256",
          "internalType": "uint256"
        },
        {
          "name": "deviceId",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "walletAddress",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "exists",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "view"
    },
    {
      "type": "function",
      "name": "owner",
      "inputs": [],
      "outputs": [
        {
          "name": "",
          "type": "address",
          "internalType": "address"
        }
      ],
      "stateMutability": "view"
    },
    {
      "type": "function",
      "name": "recoverTokens",
      "inputs": [
        {
          "name": "_tokenAddress",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "_amount",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "renounceOwnership",
      "inputs": [],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "token",
      "inputs": [],
      "outputs": [
        {
          "name": "",
          "type": "address",
          "internalType": "contract IERC20"
        }
      ],
      "stateMutability": "view"
    },
    {
      "type": "function",
      "name": "transferOwnership",
      "inputs": [
        {
          "name": "newOwner",
          "type": "address",
          "internalType": "address"
        }
      ],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "transferPayment",
      "inputs": [
        {
          "name": "_creatorDeviceId",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "_solverDeviceId",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "_amount",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "updateWalletAddress",
      "inputs": [
        {
          "name": "_deviceId",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "_newWalletAddress",
          "type": "address",
          "internalType": "address"
        }
      ],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "wallets",
      "inputs": [
        {
          "name": "",
          "type": "string",
          "internalType": "string"
        }
      ],
      "outputs": [
        {
          "name": "balance",
          "type": "uint256",
          "internalType": "uint256"
        },
        {
          "name": "deviceId",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "walletAddress",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "exists",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "view"
    },
    {
      "type": "function",
      "name": "withdrawFunds",
      "inputs": [
        {
          "name": "_deviceId",
          "type": "string",
          "internalType": "string"
        },
        {
          "name": "_amount",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "event",
      "name": "FundsAdded",
      "inputs": [
        {
          "name": "deviceId",
          "type": "string",
          "indexed": true,
          "internalType": "string"
        },
        {
          "name": "from",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "amount",
          "type": "uint256",
          "indexed": false,
          "internalType": "uint256"
        }
      ],
      "anonymous": false
    },
    {
      "type": "event",
      "name": "FundsWithdrawn",
      "inputs": [
        {
          "name": "deviceId",
          "type": "string",
          "indexed": true,
          "internalType": "string"
        },
        {
          "name": "to",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "amount",
          "type": "uint256",
          "indexed": false,
          "internalType": "uint256"
        }
      ],
      "anonymous": false
    },
    {
      "type": "event",
      "name": "OwnershipTransferred",
      "inputs": [
        {
          "name": "previousOwner",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "newOwner",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        }
      ],
      "anonymous": false
    },
    {
      "type": "event",
      "name": "TaskPayment",
      "inputs": [
        {
          "name": "creatorDeviceId",
          "type": "string",
          "indexed": true,
          "internalType": "string"
        },
        {
          "name": "solverDeviceId",
          "type": "string",
          "indexed": true,
          "internalType": "string"
        },
        {
          "name": "amount",
          "type": "uint256",
          "indexed": false,
          "internalType": "uint256"
        }
      ],
      "anonymous": false
    },
    {
      "type": "event",
      "name": "TokenRecovered",
      "inputs": [
        {
          "name": "tokenAddress",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "amount",
          "type": "uint256",
          "indexed": false,
          "internalType": "uint256"
        }
      ],
      "anonymous": false
    },
    {
      "type": "error",
      "name": "OwnableInvalidOwner",
      "inputs": [
        {
          "name": "owner",
          "type": "address",
          "internalType": "address"
        }
      ]
    },
    {
      "type": "error",
      "name": "OwnableUnauthorizedAccount",
      "inputs": [
        {
          "name": "account",
          "type": "address",
          "internalType": "address"
        }
      ]
    }
  ]`

// StakeWalletContract represents a stake wallet contract
type StakeWalletContract struct {
	address                       common.Address
	StakeWalletContractCaller     // Read-only binding to the contract
	StakeWalletContractTransactor // Write-only binding to the contract
	StakeWalletContractFilterer   // Log filterer for contract events
}

// StakeWalletContractCaller contains read-only contract methods
type StakeWalletContractCaller struct {
	contract *bind.BoundContract
}

// StakeWalletContractTransactor contains write-only contract methods
type StakeWalletContractTransactor struct {
	contract *bind.BoundContract
}

// StakeWalletContractFilterer contains contract event filtering methods
type StakeWalletContractFilterer struct {
	contract *bind.BoundContract
}

// NewStakeWalletContract creates a new instance of StakeWalletContract
func NewStakeWalletContract(address common.Address, backend bind.ContractBackend) (*StakeWalletContract, error) {
	abi, err := abi.JSON(strings.NewReader(StakeWalletContractABI))
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(address, abi, backend, backend, backend)

	return &StakeWalletContract{
		address:                       address,
		StakeWalletContractCaller:     StakeWalletContractCaller{contract: contract},
		StakeWalletContractTransactor: StakeWalletContractTransactor{contract: contract},
		StakeWalletContractFilterer:   StakeWalletContractFilterer{contract: contract},
	}, nil
}

// WalletInfo represents wallet information from the contract
type WalletInfo struct {
	Balance       *big.Int
	DeviceID      string
	WalletAddress common.Address
	Exists        bool
}

// GetBalance returns the balance for a device ID
func (c *StakeWalletContractCaller) GetBalance(opts *bind.CallOpts, deviceID string) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getBalance", deviceID)
	if err != nil {
		return nil, err
	}
	return out[0].(*big.Int), nil
}

// GetWalletInfo retrieves wallet information for a device ID
func (c *StakeWalletContractCaller) GetWalletInfo(opts *bind.CallOpts, deviceID string) (WalletInfo, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getWalletInfo", deviceID)
	if err != nil {
		return WalletInfo{}, err
	}
	return WalletInfo{
		Balance:       out[0].(*big.Int),
		DeviceID:      out[1].(string),
		WalletAddress: out[2].(common.Address),
		Exists:        out[3].(bool),
	}, nil
}

// Owner returns the contract owner address
func (c *StakeWalletContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "owner")
	if err != nil {
		return common.Address{}, err
	}
	return out[0].(common.Address), nil
}

// Token returns the token contract address
func (c *StakeWalletContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "token")
	if err != nil {
		return common.Address{}, err
	}
	return out[0].(common.Address), nil
}

// Wallets returns wallet information for a device ID directly from storage
func (c *StakeWalletContractCaller) Wallets(opts *bind.CallOpts, deviceID string) (WalletInfo, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "wallets", deviceID)
	if err != nil {
		return WalletInfo{}, err
	}
	return WalletInfo{
		Balance:       out[0].(*big.Int),
		DeviceID:      out[1].(string),
		WalletAddress: out[2].(common.Address),
		Exists:        out[3].(bool),
	}, nil
}

// AddFunds adds funds to a device's wallet
func (t *StakeWalletContractTransactor) AddFunds(opts *bind.TransactOpts, amount *big.Int, deviceID string, walletAddress common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "addFunds", amount, deviceID, walletAddress)
}

// RecoverTokens recovers tokens from the contract
func (t *StakeWalletContractTransactor) RecoverTokens(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "recoverTokens", tokenAddress, amount)
}

// RenounceOwnership renounces ownership of the contract
func (t *StakeWalletContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership transfers ownership of the contract
func (t *StakeWalletContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferPayment transfers payment between devices
func (t *StakeWalletContractTransactor) TransferPayment(opts *bind.TransactOpts, creatorDeviceID string, solverDeviceID string, amount *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transferPayment", creatorDeviceID, solverDeviceID, amount)
}

// UpdateWalletAddress updates the wallet address for a device
func (t *StakeWalletContractTransactor) UpdateWalletAddress(opts *bind.TransactOpts, deviceID string, newWalletAddress common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "updateWalletAddress", deviceID, newWalletAddress)
}

// WithdrawFunds withdraws funds from a device's wallet
func (t *StakeWalletContractTransactor) WithdrawFunds(opts *bind.TransactOpts, deviceID string, amount *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "withdrawFunds", deviceID, amount)
}

// StakeWalletContractFundsAdded represents a FundsAdded event raised by the contract
type StakeWalletContractFundsAdded struct {
	DeviceID string
	From     common.Address
	Amount   *big.Int
	Raw      types.Log
}

// StakeWalletContractFundsWithdrawn represents a FundsWithdrawn event raised by the contract
type StakeWalletContractFundsWithdrawn struct {
	DeviceID string
	To       common.Address
	Amount   *big.Int
	Raw      types.Log
}

// StakeWalletContractOwnershipTransferred represents an OwnershipTransferred event raised by the contract
type StakeWalletContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

// StakeWalletContractTaskPayment represents a TaskPayment event raised by the contract
type StakeWalletContractTaskPayment struct {
	CreatorDeviceID string
	SolverDeviceID  string
	Amount          *big.Int
	Raw             types.Log
}

// StakeWalletContractTokenRecovered represents a TokenRecovered event raised by the contract
type StakeWalletContractTokenRecovered struct {
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log
}

// FilterFundsAdded retrieves FundsAdded events from the contract
func (f *StakeWalletContractFilterer) FilterFundsAdded(opts *bind.FilterOpts, deviceID []string, from []common.Address) (*StakeWalletContractFundsAddedIterator, error) {
	var deviceIDRule []interface{}
	for _, deviceIDItem := range deviceID {
		deviceIDRule = append(deviceIDRule, deviceIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logsChan, sub, err := f.contract.FilterLogs(opts, "FundsAdded", deviceIDRule, fromRule)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}

	return &StakeWalletContractFundsAddedIterator{contract: f.contract, event: "FundsAdded", logs: logs}, nil
}

// WatchFundsAdded subscribes to FundsAdded events
func (f *StakeWalletContractFilterer) WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *StakeWalletContractFundsAdded, deviceID []string, from []common.Address) (event.Subscription, error) {
	var deviceIDRule []interface{}
	for _, deviceIDItem := range deviceID {
		deviceIDRule = append(deviceIDRule, deviceIDItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FundsAdded", deviceIDRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(StakeWalletContractFundsAdded)
				if err := f.contract.UnpackLog(event, "FundsAdded", log); err != nil {
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

// FilterFundsWithdrawn retrieves FundsWithdrawn events from the contract
func (f *StakeWalletContractFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, deviceID []string, to []common.Address) (*StakeWalletContractFundsWithdrawnIterator, error) {
	var deviceIDRule []interface{}
	for _, deviceIDItem := range deviceID {
		deviceIDRule = append(deviceIDRule, deviceIDItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logsChan, sub, err := f.contract.FilterLogs(opts, "FundsWithdrawn", deviceIDRule, toRule)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}

	return &StakeWalletContractFundsWithdrawnIterator{contract: f.contract, event: "FundsWithdrawn", logs: logs}, nil
}

// WatchFundsWithdrawn subscribes to FundsWithdrawn events
func (f *StakeWalletContractFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeWalletContractFundsWithdrawn, deviceID []string, to []common.Address) (event.Subscription, error) {
	var deviceIDRule []interface{}
	for _, deviceIDItem := range deviceID {
		deviceIDRule = append(deviceIDRule, deviceIDItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "FundsWithdrawn", deviceIDRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(StakeWalletContractFundsWithdrawn)
				if err := f.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// FilterOwnershipTransferred retrieves OwnershipTransferred events from the contract
func (f *StakeWalletContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeWalletContractOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logsChan, sub, err := f.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}

	return &StakeWalletContractOwnershipTransferredIterator{contract: f.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchOwnershipTransferred subscribes to OwnershipTransferred events
func (f *StakeWalletContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeWalletContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(StakeWalletContractOwnershipTransferred)
				if err := f.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// FilterTaskPayment retrieves TaskPayment events from the contract
func (f *StakeWalletContractFilterer) FilterTaskPayment(opts *bind.FilterOpts, creatorDeviceID []string, solverDeviceID []string) (*StakeWalletContractTaskPaymentIterator, error) {
	var creatorDeviceIDRule []interface{}
	for _, creatorDeviceIDItem := range creatorDeviceID {
		creatorDeviceIDRule = append(creatorDeviceIDRule, creatorDeviceIDItem)
	}
	var solverDeviceIDRule []interface{}
	for _, solverDeviceIDItem := range solverDeviceID {
		solverDeviceIDRule = append(solverDeviceIDRule, solverDeviceIDItem)
	}

	logsChan, sub, err := f.contract.FilterLogs(opts, "TaskPayment", creatorDeviceIDRule, solverDeviceIDRule)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}

	return &StakeWalletContractTaskPaymentIterator{contract: f.contract, event: "TaskPayment", logs: logs}, nil
}

// WatchTaskPayment subscribes to TaskPayment events
func (f *StakeWalletContractFilterer) WatchTaskPayment(opts *bind.WatchOpts, sink chan<- *StakeWalletContractTaskPayment, creatorDeviceID []string, solverDeviceID []string) (event.Subscription, error) {
	var creatorDeviceIDRule []interface{}
	for _, creatorDeviceIDItem := range creatorDeviceID {
		creatorDeviceIDRule = append(creatorDeviceIDRule, creatorDeviceIDItem)
	}
	var solverDeviceIDRule []interface{}
	for _, solverDeviceIDItem := range solverDeviceID {
		solverDeviceIDRule = append(solverDeviceIDRule, solverDeviceIDItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "TaskPayment", creatorDeviceIDRule, solverDeviceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(StakeWalletContractTaskPayment)
				if err := f.contract.UnpackLog(event, "TaskPayment", log); err != nil {
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

// FilterTokenRecovered retrieves TokenRecovered events from the contract
func (f *StakeWalletContractFilterer) FilterTokenRecovered(opts *bind.FilterOpts, tokenAddress []common.Address) (*StakeWalletContractTokenRecoveredIterator, error) {
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logsChan, sub, err := f.contract.FilterLogs(opts, "TokenRecovered", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}

	return &StakeWalletContractTokenRecoveredIterator{contract: f.contract, event: "TokenRecovered", logs: logs}, nil
}

// WatchTokenRecovered subscribes to TokenRecovered events
func (f *StakeWalletContractFilterer) WatchTokenRecovered(opts *bind.WatchOpts, sink chan<- *StakeWalletContractTokenRecovered, tokenAddress []common.Address) (event.Subscription, error) {
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "TokenRecovered", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(StakeWalletContractTokenRecovered)
				if err := f.contract.UnpackLog(event, "TokenRecovered", log); err != nil {
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

// StakeWalletContractFundsAddedIterator is returned from FilterFundsAdded and is used to iterate over the raw logs and unpacked data
type StakeWalletContractFundsAddedIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// StakeWalletContractFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data
type StakeWalletContractFundsWithdrawnIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// StakeWalletContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data
type StakeWalletContractOwnershipTransferredIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// StakeWalletContractTaskPaymentIterator is returned from FilterTaskPayment and is used to iterate over the raw logs and unpacked data
type StakeWalletContractTaskPaymentIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// StakeWalletContractTokenRecoveredIterator is returned from FilterTokenRecovered and is used to iterate over the raw logs and unpacked data
type StakeWalletContractTokenRecoveredIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// Next advances the iterator to the subsequent event
func (it *StakeWalletContractFundsAddedIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		return false
	}
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering
func (it *StakeWalletContractFundsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process
func (it *StakeWalletContractFundsAddedIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *StakeWalletContractFundsAddedIterator) Event() (*StakeWalletContractFundsAdded, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(StakeWalletContractFundsAdded)
	if err := it.contract.UnpackLog(event, "FundsAdded", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}

// Next advances the iterator to the subsequent event
func (it *StakeWalletContractFundsWithdrawnIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		return false
	}
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering
func (it *StakeWalletContractFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process
func (it *StakeWalletContractFundsWithdrawnIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *StakeWalletContractFundsWithdrawnIterator) Event() (*StakeWalletContractFundsWithdrawn, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(StakeWalletContractFundsWithdrawn)
	if err := it.contract.UnpackLog(event, "FundsWithdrawn", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}

// Next advances the iterator to the subsequent event
func (it *StakeWalletContractOwnershipTransferredIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		return false
	}
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering
func (it *StakeWalletContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process
func (it *StakeWalletContractOwnershipTransferredIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *StakeWalletContractOwnershipTransferredIterator) Event() (*StakeWalletContractOwnershipTransferred, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(StakeWalletContractOwnershipTransferred)
	if err := it.contract.UnpackLog(event, "OwnershipTransferred", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}

// Next advances the iterator to the subsequent event
func (it *StakeWalletContractTaskPaymentIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		return false
	}
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering
func (it *StakeWalletContractTaskPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process
func (it *StakeWalletContractTaskPaymentIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *StakeWalletContractTaskPaymentIterator) Event() (*StakeWalletContractTaskPayment, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(StakeWalletContractTaskPayment)
	if err := it.contract.UnpackLog(event, "TaskPayment", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}

// Next advances the iterator to the subsequent event
func (it *StakeWalletContractTokenRecoveredIterator) Next() bool {
	if it.fail != nil {
		return false
	}
	if it.done {
		return false
	}
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering
func (it *StakeWalletContractTokenRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process
func (it *StakeWalletContractTokenRecoveredIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *StakeWalletContractTokenRecoveredIterator) Event() (*StakeWalletContractTokenRecovered, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(StakeWalletContractTokenRecovered)
	if err := it.contract.UnpackLog(event, "TokenRecovered", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}
