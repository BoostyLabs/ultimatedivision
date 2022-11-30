// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencywaitlist

import (
	"context"
	"math/big"
	"time"
	"ultimatedivision/users"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"
)

// ErrNoItem indicates that item of currency wait list does not exist.
var ErrNoItem = errs.Class("item of currency wait list does not exist")

// DB is exposing access to currencywaitlist db.
//
// architecture: DB.
type DB interface {
	// Create creates item of currency waitlist in the database.
	Create(ctx context.Context, item Item) error
	// GetByWalletAddressAndNonce returns item of currency wait list by wallet address and nonce.
	GetByWalletAddressAndNonce(ctx context.Context, walletAddress common.Address, nonce int64) (Item, error)
	// List returns items of currency waitlist from database.
	List(ctx context.Context) ([]Item, error)
	// GetNonce returns number of nonce. from database.
	GetNonce(ctx context.Context) (int64, error)
	// ListWithoutSignature returns items of currency waitlist without signature from database.
	ListWithoutSignature(ctx context.Context) ([]Item, error)
	// Update updates item by wallet address and nonce in the database.
	Update(ctx context.Context, item Item) error
	// UpdateSignature updates signature of item by wallet address and nonce in the database.
	UpdateSignature(ctx context.Context, signature evmsignature.Signature, walletAddress common.Address, nonce int64) error
	// Delete deletes item of currency waitlist by wallet address and nonce in the database.
	Delete(ctx context.Context, walletAddress common.Address, nonce int64) error
}

// Item entity describes item of currency wait list.
type Item struct {
	WalletAddress common.Address         `json:"walletAddress"`
	WalletType    users.WalletType       `json:"walleType"`
	Value         big.Int                `json:"value"`
	Nonce         int64                  `json:"nonce"`
	Signature     evmsignature.Signature `json:"signature"`
}

// Transaction entity describes values for creating transaction to contract.
type Transaction struct {
	Signature   evmsignature.Signature `json:"signature"`
	UDTContract evmsignature.Contract  `json:"udtContract"`
	Value       string                 `json:"value"`
}

// Config defines values needed by mint udt tokens in blockchain.
type Config struct {
	IntervalSignatureCheck time.Duration         `json:"intervalSignatureCheck"`
	UDTContract            evmsignature.Contract `json:"udtContract"`
}
