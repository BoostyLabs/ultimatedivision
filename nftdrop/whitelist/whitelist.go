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
	Create(ctx context.Context, whitelist Whitelist) error
	// GetByAddress returns whitelist by address from the database.
	GetByAddress(ctx context.Context, address Address) (Whitelist, error)
	// List returns all whitelist from the database.
	List(ctx context.Context) ([]Whitelist, error)
	// ListWithoutPassword returns whitelist without password from the database.
	ListWithoutPassword(ctx context.Context) ([]Whitelist, error)
	// Update updates whitelist by address.
	Update(ctx context.Context, whitelist Whitelist) error
	// Delete deletes whitelist from the database.
	Delete(ctx context.Context, address Address) error
	// Update updates a whitelists password in the data base.
	Update(ctx context.Context, whitelist Whitelist) error
	// ListWithoutPassword returns all whitelist address from the data base.
	ListWithoutPassword(ctx context.Context) ([]Whitelist, error)
}

// Whitelist describes whitelist entity.
type Whitelist struct {
	Address  Address `json:"address"`
	Password []byte  `json:"password"`
}

// Address defines address of user's wallet.
type Address string

// ValidateAddress checks if the address is valid.
func (address Address) ValidateAddress() bool {
	return common.IsHexAddress(string(address))
}

// Request entity describes request values for create whitelist.
type Request struct {
	Address Address `json:"address"`
	Key     string  `json:"key"`
}

// Config defines configuration for queue.
type Config struct {
	SmartContract struct {
		Address string  `json:"address"`
		Price   float64 `json:"price"`
	} `json:"smartContract"`
}
