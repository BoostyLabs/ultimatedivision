// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/signature"
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
	GetByAddress(ctx context.Context, address signature.Address) (Wallet, error)
	// List returns all whitelist from the database.
	List(ctx context.Context) ([]Wallet, error)
	// ListWithoutPassword returns whitelist without password from the database.
	ListWithoutPassword(ctx context.Context) ([]Wallet, error)
	// Update updates whitelist by address.
	Update(ctx context.Context, wallet Wallet) error
	// Delete deletes whitelist from the database.
	Delete(ctx context.Context, address signature.Address) error
}

// Wallet describes whitelist entity.
type Wallet struct {
	Address  signature.Address   `json:"address"`
	Password signature.Signature `json:"password"`
}

// CreateWallet entity describes request values for create whitelist.
type CreateWallet struct {
	Address    signature.Address    `json:"address"`
	PrivateKey signature.PrivateKey `json:"privateKey"`
}

// Config defines configuration for queue.
type Config struct {
	SmartContractAddress `json:"smartContractAddress"`
}

// SmartContractAddress entity describes smart contract addresses.
type SmartContractAddress struct {
	NFT     signature.Address `json:"nft"`
	NFTSale signature.Address `json:"nftSale"`
}

// Transaction entity describes password wallet and smart contract addresses.
type Transaction struct {
	Password             signature.Signature `json:"password"`
	SmartContractAddress `json:"smartContractAddress"`
}
