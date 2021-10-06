// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/zeebo/errs"
)

// ErrNoAddress indicated that address does not exist.
var ErrNoAddress = errs.Class("address does not exist")

// DB is exposing access to whitelist db.
//
// architecture: DB
type DB interface {
	// Add adds wallet to the database.
	Add(ctx context.Context, whitelist Whitelist) error
	// Get returns password of address.
	Get(ctx context.Context, address string) ([]byte, error)
}

type Whitelist struct {
	Address  string `json:"address"`
	Password []byte `json:"password"`
}
