// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts"
)

// ensures that nftsDB implements nfts.DB.
var _ nfts.DB = (*nftsDB)(nil)

// ErrNFTs indicates that there was an error in the database.
var ErrNFTs = errs.Class("ErrNFTs repository error")

// nftsDB provide access to nfts DB.
//
// architecture: Database
type nftsDB struct {
	conn *sql.DB
}

// Create creates nft in the database.
func (nftsDB *nftsDB) Create(ctx context.Context, nft nfts.NFT) error {
	query := `INSERT INTO nfts(card_id, token_id, chain, wallet_address)
	          VALUES($1,$2,$3,$4)`

	_, err := nftsDB.conn.ExecContext(ctx, query, nft.CardID, nft.TokenID, nft.Chain, nft.WalletAddress)
	return ErrNFTs.Wrap(err)
}

// List returns nfts from database.
func (nftsDB *nftsDB) List(ctx context.Context) ([]nfts.NFT, error) {
	var nftList []nfts.NFT
	query := `SELECT * FROM nfts`

	rows, err := nftsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nftList, ErrNFTs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var nft nfts.NFT

		if err = rows.Scan(&nft.CardID, &nft.TokenID, &nft.Chain, &nft.WalletAddress); err != nil {
			return nftList, ErrNFTs.Wrap(err)
		}
		nftList = append(nftList, nft)
	}

	return nftList, ErrNFTs.Wrap(rows.Err())
}