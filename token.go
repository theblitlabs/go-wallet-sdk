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

// ParityTokenABI is the input ABI used to generate the binding from.
const ParityTokenABI = `[
    {
      "type": "constructor",
      "inputs": [
        {
          "name": "initialSupply",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "allowance",
      "inputs": [
        {
          "name": "",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "",
          "type": "address",
          "internalType": "address"
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
      "name": "approve",
      "inputs": [
        {
          "name": "spender",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "value",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [
        {
          "name": "success",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "balanceOf",
      "inputs": [
        {
          "name": "",
          "type": "address",
          "internalType": "address"
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
      "name": "burn",
      "inputs": [
        {
          "name": "value",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [
        {
          "name": "success",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "decimals",
      "inputs": [],
      "outputs": [
        {
          "name": "",
          "type": "uint8",
          "internalType": "uint8"
        }
      ],
      "stateMutability": "view"
    },
    {
      "type": "function",
      "name": "mint",
      "inputs": [
        {
          "name": "to",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "value",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [
        {
          "name": "success",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "name",
      "inputs": [],
      "outputs": [
        {
          "name": "",
          "type": "string",
          "internalType": "string"
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
      "name": "renounceOwnership",
      "inputs": [],
      "outputs": [],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "symbol",
      "inputs": [],
      "outputs": [
        {
          "name": "",
          "type": "string",
          "internalType": "string"
        }
      ],
      "stateMutability": "view"
    },
    {
      "type": "function",
      "name": "totalSupply",
      "inputs": [],
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
      "name": "transfer",
      "inputs": [
        {
          "name": "to",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "value",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [
        {
          "name": "success",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "transferFrom",
      "inputs": [
        {
          "name": "from",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "to",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "value",
          "type": "uint256",
          "internalType": "uint256"
        }
      ],
      "outputs": [
        {
          "name": "success",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "nonpayable"
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
      "name": "transferWithData",
      "inputs": [
        {
          "name": "to",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "value",
          "type": "uint256",
          "internalType": "uint256"
        },
        {
          "name": "",
          "type": "bytes",
          "internalType": "bytes"
        }
      ],
      "outputs": [
        {
          "name": "success",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "function",
      "name": "transferWithDataAndCallback",
      "inputs": [
        {
          "name": "to",
          "type": "address",
          "internalType": "address"
        },
        {
          "name": "value",
          "type": "uint256",
          "internalType": "uint256"
        },
        {
          "name": "data",
          "type": "bytes",
          "internalType": "bytes"
        }
      ],
      "outputs": [
        {
          "name": "",
          "type": "bool",
          "internalType": "bool"
        }
      ],
      "stateMutability": "nonpayable"
    },
    {
      "type": "event",
      "name": "Approval",
      "inputs": [
        {
          "name": "owner",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "spender",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "value",
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
      "name": "Transfer",
      "inputs": [
        {
          "name": "from",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "to",
          "type": "address",
          "indexed": true,
          "internalType": "address"
        },
        {
          "name": "value",
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

// ParityToken represents a Parity token contract
type ParityToken struct {
	ParityTokenCaller     // Read-only binding to the contract
	ParityTokenTransactor // Write-only binding to the contract
	ParityTokenFilterer   // Log filterer for contract events
}

// ParityTokenCaller contains read-only contract methods
type ParityTokenCaller struct {
	contract *bind.BoundContract
}

// ParityTokenTransactor contains write-only contract methods
type ParityTokenTransactor struct {
	contract *bind.BoundContract
}

// ParityTokenFilterer contains contract event filtering methods
type ParityTokenFilterer struct {
	contract *bind.BoundContract
}

// NewParityToken creates a new instance of ParityToken
func NewParityToken(address common.Address, backend bind.ContractBackend) (*ParityToken, error) {
	abi, err := abi.JSON(strings.NewReader(ParityTokenABI))
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(address, abi, backend, backend, backend)

	return &ParityToken{
		ParityTokenCaller:     ParityTokenCaller{contract: contract},
		ParityTokenTransactor: ParityTokenTransactor{contract: contract},
		ParityTokenFilterer:   ParityTokenFilterer{contract: contract},
	}, nil
}

// Name returns the token name
func (c *ParityTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "name")
	if err != nil {
		return "", err
	}
	return out[0].(string), nil
}

// Symbol returns the token symbol
func (c *ParityTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "symbol")
	if err != nil {
		return "", err
	}
	return out[0].(string), nil
}

// Decimals returns the token decimals
func (c *ParityTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "decimals")
	if err != nil {
		return 0, err
	}
	return uint8(out[0].(uint8)), nil
}

// TotalSupply returns the total token supply
func (c *ParityTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "totalSupply")
	if err != nil {
		return nil, err
	}
	return out[0].(*big.Int), nil
}

// BalanceOf returns the token balance of an account
func (c *ParityTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "balanceOf", account)
	if err != nil {
		return nil, err
	}
	return out[0].(*big.Int), nil
}

// Allowance returns the token allowance for owner and spender
func (c *ParityTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "allowance", owner, spender)
	if err != nil {
		return nil, err
	}
	return out[0].(*big.Int), nil
}

// Transfer transfers tokens to an address
func (t *ParityTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transfer", to, value)
}

// Approve approves tokens for a spender
func (t *ParityTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "approve", spender, value)
}

// TransferFrom transfers tokens from one address to another
func (t *ParityTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transferFrom", from, to, value)
}

// Mint mints new tokens
func (t *ParityTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "mint", to, value)
}

// Burn burns tokens
func (t *ParityTokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "burn", value)
}

// TransferWithData transfers tokens with additional data
func (t *ParityTokenTransactor) TransferWithData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transferWithData", to, value, data)
}

// TransferWithDataAndCallback transfers tokens with data and callback
func (t *ParityTokenTransactor) TransferWithDataAndCallback(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transferWithDataAndCallback", to, value, data)
}

// Owner returns the owner address
func (c *ParityTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "owner")
	if err != nil {
		return common.Address{}, err
	}
	return out[0].(common.Address), nil
}

// RenounceOwnership renounces ownership of the contract
func (t *ParityTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return t.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership transfers ownership of the contract to a new address
func (t *ParityTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transferOwnership", newOwner)
}

// FilterTransfer is a free log retrieval operation binding the contract Transfer event
func (f *ParityTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ParityTokenTransferIterator, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logsChan, _, err := f.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}
	return &ParityTokenTransferIterator{contract: f.contract, event: "Transfer", logs: logs}, nil
}

// FilterApproval is a free log retrieval operation binding the contract Approval event
func (f *ParityTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ParityTokenApprovalIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logsChan, _, err := f.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}
	return &ParityTokenApprovalIterator{contract: f.contract, event: "Approval", logs: logs}, nil
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract OwnershipTransferred event
func (f *ParityTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParityTokenOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logsChan, _, err := f.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	logs := make([]types.Log, 0)
	for log := range logsChan {
		logs = append(logs, log)
	}
	return &ParityTokenOwnershipTransferredIterator{contract: f.contract, event: "OwnershipTransferred", logs: logs}, nil
}

// WatchTransfer is a free log subscription operation binding the contract Transfer event
func (f *ParityTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ParityTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ParityTokenTransfer)
				if err := f.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// WatchApproval is a free log subscription operation binding the contract Approval event
func (f *ParityTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ParityTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := f.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				event := new(ParityTokenApproval)
				if err := f.contract.UnpackLog(event, "Approval", log); err != nil {
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

// WatchOwnershipTransferred is a free log subscription operation binding the contract OwnershipTransferred event
func (f *ParityTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParityTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
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
				event := new(ParityTokenOwnershipTransferred)
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

// ParityTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ParityToken contract.
type ParityTokenTransferIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// ParityTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ParityToken contract.
type ParityTokenApprovalIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// ParityTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ParityToken contract.
type ParityTokenOwnershipTransferredIterator struct {
	event    string
	contract *bind.BoundContract
	logs     []types.Log
	done     bool
	fail     error
}

// ParityTokenTransfer represents a Transfer event raised by the ParityToken contract.
type ParityTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// ParityTokenApproval represents an Approval event raised by the ParityToken contract.
type ParityTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// ParityTokenOwnershipTransferred represents an OwnershipTransferred event raised by the ParityToken contract.
type ParityTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParityTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		return false
	}
	// Iterator still has logs, shift after the current value
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParityTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParityTokenTransferIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *ParityTokenTransferIterator) Event() (*ParityTokenTransfer, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(ParityTokenTransfer)
	if err := it.contract.UnpackLog(event, "Transfer", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParityTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		return false
	}
	// Iterator still has logs, shift after the current value
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParityTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParityTokenApprovalIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *ParityTokenApprovalIterator) Event() (*ParityTokenApproval, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(ParityTokenApproval)
	if err := it.contract.UnpackLog(event, "Approval", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParityTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		return false
	}
	// Iterator still has logs, shift after the current value
	if len(it.logs) == 0 {
		it.done = true
		return false
	}
	it.logs = it.logs[1:]
	return true
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParityTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParityTokenOwnershipTransferredIterator) Close() error {
	it.done = true
	return nil
}

// Event returns the parsed event data for the current log
func (it *ParityTokenOwnershipTransferredIterator) Event() (*ParityTokenOwnershipTransferred, error) {
	if len(it.logs) == 0 {
		return nil, fmt.Errorf("no more events")
	}
	event := new(ParityTokenOwnershipTransferred)
	if err := it.contract.UnpackLog(event, "OwnershipTransferred", it.logs[0]); err != nil {
		return nil, err
	}
	event.Raw = it.logs[0]
	return event, nil
}
