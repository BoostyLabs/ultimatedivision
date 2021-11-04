// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts"
)

// ensures that lootboxDB implements lootbox.DB.
var _ nfts.DB = (*nftsDB)(nil)

// ErrNFTs indicates that there was an error in the database.
var ErrNFTs = errs.Class("NFTs repository error")

// nftsDB provide access to nfts DB.
//
// architecture: Database
type nftsDB struct {
	conn *sql.DB
}

func Create(ctx context.Context, )


