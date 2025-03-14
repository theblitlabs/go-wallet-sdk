package paritysdk

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// StakeWalletContractABI is the input ABI used to generate the binding from.
const StakeWalletContractABI = `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_tokenAddress",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "string",
				"name": "deviceID",
				"type": "string"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "walletAddress",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "StakeDeposited",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "string",
				"name": "deviceID",
				"type": "string"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "walletAddress",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "StakeWithdrawn",
		"type": "event"
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

// GetStakeInfo retrieves stake information for a device ID
func (c *StakeWalletContractCaller) GetStakeInfo(opts *bind.CallOpts, deviceID string) (StakeInfo, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getStakeInfo", deviceID)
	if err != nil {
		return StakeInfo{}, err
	}
	return out[0].(StakeInfo), nil
}

// GetBalanceByDeviceID returns the stake balance for a device ID
func (c *StakeWalletContractCaller) GetBalanceByDeviceID(opts *bind.CallOpts, deviceID string) (*big.Int, error) {
	var out []interface{}
	err := c.contract.Call(opts, &out, "getBalanceByDeviceID", deviceID)
	if err != nil {
		return nil, err
	}
	return out[0].(*big.Int), nil
}

// Stake stakes tokens with device ID
func (t *StakeWalletContractTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, deviceID string, walletAddr common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "stake", amount, deviceID, walletAddr)
}

// TransferPayment transfers stake between devices
func (t *StakeWalletContractTransactor) TransferPayment(opts *bind.TransactOpts, creatorDeviceID string, solverDeviceID string, amount *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "transferPayment", creatorDeviceID, solverDeviceID, amount)
}

// WithdrawFunds withdraws staked tokens
func (t *StakeWalletContractTransactor) WithdrawFunds(opts *bind.TransactOpts, deviceID string, amount *big.Int) (*types.Transaction, error) {
	return t.contract.Transact(opts, "withdrawFunds", deviceID, amount)
}

// UpdateWalletAddress updates the wallet address for a device ID
func (t *StakeWalletContractTransactor) UpdateWalletAddress(opts *bind.TransactOpts, deviceID string, newWalletAddr common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "updateWalletAddress", deviceID, newWalletAddr)
}
