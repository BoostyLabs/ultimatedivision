// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"
)

// ErrNoWhitelist indicated that whitelist does not exist.
var ErrNoWhitelist = errs.Class("whitelist does not exist")

// DB is exposing access to whitelist db.
//
// architecture: DB
type DB interface {
	// Create adds whitelist in the database.
	Create(ctx context.Context, wallet Wallet) error
	// GetByAddress returns whitelist by address from the database.
	GetByAddress(ctx context.Context, address Address) (Wallet, error)
	// List returns all whitelist from the database.
	List(ctx context.Context) ([]Wallet, error)
	// ListWithoutPassword returns whitelist without password from the database.
	ListWithoutPassword(ctx context.Context) ([]Wallet, error)
	// Update updates whitelist by address.
	Update(ctx context.Context, wallet Wallet) error
	// Delete deletes whitelist from the database.
	Delete(ctx context.Context, address Address) error
}

// Wallet describes whitelist entity.
type Wallet struct {
	Address  Address   `json:"address"`
	Password Signature `json:"password"`
}

// Address defines address type.
type Address string

// Signature defines signature type.
type Signature string

// PrivateKey defines private key type.
type PrivateKey string

// LengthPrivateKey defines length private key.
const LengthPrivateKey int = 64

// PrivateKeyV defines v of private key type.
type PrivateKeyV int

const (
	// PrivateKeyVZero indicates that the v of private key is 0.
	PrivateKeyVZero PrivateKeyV = 0
	// PrivateKeyVOne indicates that the v of private key is 1.
	PrivateKeyVOne PrivateKeyV = 1
	// PrivateKeyVTwentySeven indicates that the v of private key is 27.
	PrivateKeyVTwentySeven PrivateKeyV = 27
	// PrivateKeyVTwentyEight indicates that the v of private key is 28.
	PrivateKeyVTwentyEight PrivateKeyV = 28
)

// IsValidAddress checks if the address is valid.
func (address Address) IsValidAddress() bool {
	return common.IsHexAddress(string(address))
}

// IsValidPrivateKey validates whether each byte is valid hexadecimal private key.
func (privateKey PrivateKey) IsValidPrivateKey() bool {
	if len(string(privateKey)) != LengthPrivateKey {
		return false
	}
	for _, c := range []byte(string(privateKey)) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// CreateWallet entity describes request values for create whitelist.
type CreateWallet struct {
	Address    Address    `json:"address"`
	PrivateKey PrivateKey `json:"privateKey"`
}

// Config defines configuration for queue.
type Config struct {
	SmartContractAddress `json:"smartContractAddress"`
}

// SmartContractAddress entity describes smart contract addresses.
type SmartContractAddress struct {
	NFT     Address `json:"NFT"`
	NFTSale Address `json:"NFTSale"`
}

// Transaction entity describes password wallet and smart contract addresses.
type Transaction struct {
	Password             Signature `json:"password"`
	SmartContractAddress `json:"smartContractAddress"`
}

// EthereumSignedMessageHash defines message for sinbature.
const EthereumSignedMessageHash string = "19457468657265756d205369676e6564204d6573736167653a0a3332"
