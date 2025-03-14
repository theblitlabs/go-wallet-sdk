// Package walletsdk provides a unified SDK for Parity token and staking operations
package walletsdk

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents a unified Parity SDK client
type Client struct {
	*ethclient.Client
	chainID     *big.Int
	auth        *bind.TransactOpts
	privateKey  *ecdsa.PrivateKey
	address     common.Address
	token       *ParityToken
	stakeWallet *StakeWallet
}

// ClientConfig represents the configuration for creating a new client
type ClientConfig struct {
	RPCURL       string
	ChainID      int64
	TokenAddress common.Address
	StakeAddress common.Address
	PrivateKey   string
}

// NewClient creates a new Parity SDK client
func NewClient(config ClientConfig) (*Client, error) {
	ethClient, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	token, err := NewParityToken(config.TokenAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create token contract: %w", err)
	}

	client := &Client{
		Client:  ethClient,
		chainID: big.NewInt(config.ChainID),
		token:   token,
	}

	if config.PrivateKey != "" {
		if err := client.SetPrivateKey(config.PrivateKey); err != nil {
			return nil, err
		}
	}

	if config.StakeAddress != (common.Address{}) {
		stakeWallet, err := NewStakeWallet(client, config.StakeAddress, config.TokenAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to create stake wallet: %w", err)
		}
		client.stakeWallet = stakeWallet
	}

	return client, nil
}

// SetPrivateKey sets the private key for the client
func (c *Client) SetPrivateKey(privateKey string) error {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, c.chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	c.privateKey = key
	c.auth = auth
	c.address = crypto.PubkeyToAddress(key.PublicKey)

	return nil
}

// Address returns the wallet address
func (c *Client) Address() common.Address {
	return c.address
}

// GetTransactOpts returns the transaction options for contract interactions
func (c *Client) GetTransactOpts() (*bind.TransactOpts, error) {
	if c.auth == nil {
		return nil, fmt.Errorf("wallet not authenticated")
	}
	return c.auth, nil
}

// GetBalance returns the token balance for an address
func (c *Client) GetBalance(address common.Address) (*big.Int, error) {
	return c.token.BalanceOf(&bind.CallOpts{}, address)
}

// GetTokenInfo returns token information
func (c *Client) GetTokenInfo() (name string, symbol string, decimals uint8, err error) {
	opts := &bind.CallOpts{}

	name, err = c.token.Name(opts)
	if err != nil {
		return "", "", 0, err
	}

	symbol, err = c.token.Symbol(opts)
	if err != nil {
		return "", "", 0, err
	}

	decimals, err = c.token.Decimals(opts)
	if err != nil {
		return "", "", 0, err
	}

	return name, symbol, decimals, nil
}

// GetAllowance returns the token allowance for owner and spender
func (c *Client) GetAllowance(owner, spender common.Address) (*big.Int, error) {
	return c.token.Allowance(&bind.CallOpts{}, owner, spender)
}

// GetTotalSupply returns the total token supply
func (c *Client) GetTotalSupply() (*big.Int, error) {
	return c.token.TotalSupply(&bind.CallOpts{})
}

// Transfer transfers tokens to an address
func (c *Client) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	opts, err := c.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	return c.token.Transfer(opts, to, amount)
}

// Approve approves tokens for a spender
func (c *Client) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	opts, err := c.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	return c.token.Approve(opts, spender, amount)
}

// TransferFrom transfers tokens from one address to another
func (c *Client) TransferFrom(from, to common.Address, amount *big.Int) (*types.Transaction, error) {
	opts, err := c.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	return c.token.TransferFrom(opts, from, to, amount)
}

// Mint mints new tokens
func (c *Client) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	opts, err := c.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	return c.token.Mint(opts, to, amount)
}

// Burn burns tokens
func (c *Client) Burn(amount *big.Int) (*types.Transaction, error) {
	opts, err := c.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	return c.token.Burn(opts, amount)
}

// TransferWithData transfers tokens with additional data
func (c *Client) TransferWithData(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	opts, err := c.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	return c.token.TransferWithData(opts, to, amount, data)
}

// TransferWithDataAndCallback transfers tokens with data and callback
func (c *Client) TransferWithDataAndCallback(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	opts, err := c.GetTransactOpts()
	if err != nil {
		return nil, err
	}
	return c.token.TransferWithDataAndCallback(opts, to, amount, data)
}

// GetStakeInfo retrieves stake information for a device ID
func (c *Client) GetStakeInfo(deviceID string) (StakeInfo, error) {
	if c.stakeWallet == nil {
		return StakeInfo{}, fmt.Errorf("stake wallet not initialized")
	}
	return c.stakeWallet.GetStakeInfo(deviceID)
}

// AddFunds adds funds to a device's wallet
func (c *Client) AddFunds(amount *big.Int, deviceID string) (*types.Transaction, error) {
	if c.stakeWallet == nil {
		return nil, fmt.Errorf("stake wallet not initialized")
	}
	return c.stakeWallet.Stake(amount, deviceID)
}

// TransferPayment transfers stake between devices
func (c *Client) TransferPayment(creatorDeviceID, solverDeviceID string, amount *big.Int) (*types.Transaction, error) {
	if c.stakeWallet == nil {
		return nil, fmt.Errorf("stake wallet not initialized")
	}
	return c.stakeWallet.TransferPayment(creatorDeviceID, solverDeviceID, amount)
}

// GetStakeBalance returns the stake balance for a device ID
func (c *Client) GetStakeBalance(deviceID string) (*big.Int, error) {
	if c.stakeWallet == nil {
		return nil, fmt.Errorf("stake wallet not initialized")
	}
	return c.stakeWallet.GetBalance(deviceID)
}

// WithdrawFunds withdraws staked tokens
func (c *Client) WithdrawFunds(deviceID string, amount *big.Int) (*types.Transaction, error) {
	if c.stakeWallet == nil {
		return nil, fmt.Errorf("stake wallet not initialized")
	}
	return c.stakeWallet.WithdrawStake(deviceID, amount)
}

// UpdateWalletAddress updates the wallet address for a device ID
func (c *Client) UpdateWalletAddress(deviceID string, newWalletAddr common.Address) (*types.Transaction, error) {
	if c.stakeWallet == nil {
		return nil, fmt.Errorf("stake wallet not initialized")
	}
	return c.stakeWallet.UpdateWalletAddress(deviceID, newWalletAddr)
}
