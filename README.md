# Parity SDK for Go

## Features

- Complete token management functionality
  - Transfer tokens
  - Check balances
  - Approve spending
  - Mint and burn tokens
  - Transfer with data and callbacks
- Full staking capabilities
  - Stake tokens with device IDs
  - Transfer payments between devices
  - Withdraw stakes
  - Update wallet addresses
  - Check stake information
- Unified client interface
- Built-in Ethereum integration
- Type-safe contract bindings

## Installation

```bash
go get github.com/theblitlabs/go-wallet-sdk
```

## Quick Start

```go
package main

import (
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/theblitlabs/go-wallet-sdk"
)

func main() {
    // Create a new client
    config := walletsdk.ClientConfig{
        RPCURL:       "https://your-ethereum-node.com",
        ChainID:      1, // Mainnet
        TokenAddress: common.HexToAddress("0xtoken_address"),
        StakeAddress: common.HexToAddress("0xstake_address"),
        PrivateKey:   "your_private_key",
    }

    client, err := walletsdk.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    // Get token information
    name, symbol, decimals, err := client.GetTokenInfo()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Token: %s (%s) with %d decimals", name, symbol, decimals)

    // Check balance
    balance, err := client.GetBalance(client.Address())
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Balance: %s", balance.String())

    // Stake tokens
    amount := big.NewInt(1000000000000000000) // 1 token
    deviceID := "device123"
    tx, err := client.Stake(amount, deviceID)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Stake transaction hash: %s", tx.Hash().Hex())

    // Get stake information
    stakeInfo, err := client.GetStakeInfo(deviceID)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Stake amount: %s", stakeInfo.Amount.String())
}
```

## Token Operations

### Transfer Tokens

```go
to := common.HexToAddress("0xrecipient")
amount := big.NewInt(1000000000000000000) // 1 token
tx, err := client.Transfer(to, amount)
```

### Check Balance

```go
address := common.HexToAddress("0xaddress")
balance, err := client.GetBalance(address)
```

### Approve Spending

```go
spender := common.HexToAddress("0xspender")
amount := big.NewInt(1000000000000000000)
tx, err := client.Approve(spender, amount)
```

## Staking Operations

### Stake Tokens

```go
amount := big.NewInt(1000000000000000000)
deviceID := "device123"
tx, err := client.Stake(amount, deviceID)
```

### Transfer Payment

```go
creatorDeviceID := "device123"
solverDeviceID := "device456"
amount := big.NewInt(500000000000000000)
tx, err := client.TransferPayment(creatorDeviceID, solverDeviceID, amount)
```

### Withdraw Stake

```go
deviceID := "device123"
amount := big.NewInt(1000000000000000000)
tx, err := client.WithdrawStake(deviceID, amount)
```

## Error Handling

The SDK uses standard Go error handling patterns. All operations that can fail return an error as the last return value. Always check these errors in production code.

```go
tx, err := client.Transfer(to, amount)
if err != nil {
    // Handle error appropriately
    log.Printf("Transfer failed: %v", err)
    return
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
