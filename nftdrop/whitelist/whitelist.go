// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"
	"regexp"

	"github.com/zeebo/errs"
)

// ErrNoWhitelist indicated that whitelist does not exist.
var ErrNoWhitelist = errs.Class("whitelist does not exist")

// RegularIsEthereumAddress indicated that expression is regular expression for ethereum address.
const RegularIsEthereumAddress = "^0x[0-9a-fA-F]{40}$"

// DB is exposing access to whitelist db.
//
// architecture: DB
type DB interface {
	// Create adds whitelist in the data base.
	Create(ctx context.Context, whitelist Whitelist) error
	// Get returns whitelist by address from the data base.
	Get(ctx context.Context, address string) (Whitelist, error)
	// List returns all whitelist from the data base.
	List(ctx context.Context) ([]Whitelist, error)
}

// Whitelist describes whitelist entity.
type Whitelist struct {
	Address  string `json:"address"`
	Password []byte `json:"password"`
}

// Request entity describes request values for create whitelist.
type Request struct {
	Address string `json:"address"`
	Signature string `json:"signature"`
}

// ValidateAddress checks if the address is valid.
func (w Request) ValidateAddress() bool {
	// TODO: rework.
	re := regexp.MustCompile(RegularIsEthereumAddress)
	return re.MatchString(w.Address)
}
