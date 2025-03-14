package walletsdk

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ParityTokenABI is the input ABI used to generate the binding from.
const ParityTokenABI = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "initialSupply",
				"type": "uint256"
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
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
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
