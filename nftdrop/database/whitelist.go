// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/zeebo/errs"

	"nftdrop/whitelist"
)

// ensures that whitelistDB implements whitelist.DB.
var _ whitelist.DB = (*whitelistDB)(nil)


// ErrWhitelist indicates that there was an error in the database.
var ErrWhitelist = errs.Class("whitelist repository error")

// whitelistDB provide access to admin DB.
//
// architecture: Database
type whitelistDB struct {
	conn *sql.DB
}

// Add adds wallet to the database.
func (whitelistDB *whitelistDB) Add(ctx context.Context, whitelist whitelist.Whitelist) error {
	return nil
}

// Get returns password of address.
func (whitelistDB *whitelistDB) Get(ctx context.Context, address string) ([]byte, error) {
	return nil, nil
}

