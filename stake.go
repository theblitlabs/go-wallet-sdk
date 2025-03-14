package walletsdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// StakeInfo represents stake information for a device
type StakeInfo struct {
	Amount        *big.Int
	DeviceID      string
	WalletAddress common.Address
	Exists        bool
}

// StakeWallet manages staking operations
type StakeWallet struct {
	client    *Client
	contract  *StakeWalletContract
	tokenAddr common.Address
}

// NewStakeWallet creates a new stake wallet instance
func NewStakeWallet(client *Client, contractAddr, tokenAddr common.Address) (*StakeWallet, error) {
	contract, err := NewStakeWalletContract(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &StakeWallet{
		client:    client,
		contract:  contract,
		tokenAddr: tokenAddr,
	}, nil
}

// GetStakeInfo retrieves stake information for a device ID
func (s *StakeWallet) GetStakeInfo(deviceID string) (StakeInfo, error) {
	return s.contract.GetStakeInfo(&bind.CallOpts{}, deviceID)
}

// Stake tokens with device ID
func (s *StakeWallet) Stake(amount *big.Int, deviceID string) (*types.Transaction, error) {
	// First approve the contract to spend tokens
	tokenClient, err := NewParityToken(s.tokenAddr, s.client)
	if err != nil {
		return nil, err
	}

	opts, err := s.client.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	// Approve the contract to spend tokens
	tx, err := tokenClient.Approve(opts, s.contract.address, amount)
	if err != nil {
		return nil, err
	}

	// Wait for approval to be mined
	_, err = bind.WaitMined(context.Background(), s.client, tx)
	if err != nil {
		return nil, err
	}

	// Now stake the tokens
	return s.contract.Stake(opts, amount, deviceID, s.client.Address())
}

// TransferPayment transfers stake between devices
func (s *StakeWallet) TransferPayment(creatorDeviceID, solverDeviceID string, amount *big.Int) (*types.Transaction, error) {
	opts, err := s.client.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	return s.contract.TransferPayment(opts, creatorDeviceID, solverDeviceID, amount)
}

// GetBalance returns the stake balance for a device ID
func (s *StakeWallet) GetBalance(deviceID string) (*big.Int, error) {
	return s.contract.GetBalanceByDeviceID(&bind.CallOpts{}, deviceID)
}

// WithdrawStake withdraws staked tokens
func (s *StakeWallet) WithdrawStake(deviceID string, amount *big.Int) (*types.Transaction, error) {
	opts, err := s.client.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	return s.contract.WithdrawFunds(opts, deviceID, amount)
}

// UpdateWalletAddress updates the wallet address for a device ID
func (s *StakeWallet) UpdateWalletAddress(deviceID string, newWalletAddr common.Address) (*types.Transaction, error) {
	opts, err := s.client.GetTransactOpts()
	if err != nil {
		return nil, err
	}

	return s.contract.UpdateWalletAddress(opts, deviceID, newWalletAddr)
}
