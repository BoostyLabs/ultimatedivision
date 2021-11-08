// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts"
	"ultimatedivision/pkg/cryptoutils"
)

// ensures that nftsDB implements nfts.DB.
var _ nfts.DB = (*nftsDB)(nil)

// ErrNFTs indicates that there was an error in the database.
var ErrNFTs = errs.Class("NFTs repository error")

// nftsDB provide access to nfts DB.
//
// architecture: Database
type nftsDB struct {
	conn *sql.DB
}

// Create creates nft token in the database.
func (nftsDB *nftsDB) Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address, password cryptoutils.Signature) error {
	query := `INSERT INTO nfts_waitlist(card_id, wallet_address, password)
	          VALUES($1,$2,$3)`

	_, err := nftsDB.conn.ExecContext(ctx, query, cardID, wallet, password)
	return ErrNFTs.Wrap(err)
}

// List returns all nft token from wait list from database.
func (nftsDB *nftsDB) List(ctx context.Context) ([]nfts.NFTWaitList, error) {
	query := `SELECT *
	          FROM nfts_waitlist`

	var nftList []nfts.NFTWaitList

	rows, err := nftsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nftList, ErrNFTs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var nft nfts.NFTWaitList
		err = rows.Scan(&nft.TokenID, &nft.CardID, &nft.Wallet, &nft.Password)
		if err != nil {
			return nftList, ErrNFTs.Wrap(err)
		}

		nftList = append(nftList, nft)
	}
	if err = rows.Err(); err != nil {
		return nftList, ErrNFTs.Wrap(err)
	}

	return nftList, ErrNFTs.Wrap(err)
}

// ListWithoutPassword returns all nft tokens without password from database.
func (nftsDB *nftsDB) ListWithoutPassword(ctx context.Context) ([]nfts.NFTWaitList, error) {
	query :=
		`SELECT *
	     FROM nfts_waitlist
	     WHERE password = ''`

	rows, err := nftsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrNFTs.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var nftsWithoutPassword []nfts.NFTWaitList
	for rows.Next() {
		var nft nfts.NFTWaitList
		if err = rows.Scan(&nft.TokenID, &nft.CardID, &nft.Wallet, &nft.Password); err != nil {
			return nil, ErrNFTs.Wrap(err)
		}
		nftsWithoutPassword = append(nftsWithoutPassword, nft)
	}

	return nftsWithoutPassword, ErrNFTs.Wrap(rows.Err())
}

// Get returns nft token by card id.
func (nftsDB *nftsDB) Get(ctx context.Context, tokenID int) (nfts.NFTWaitList, error) {
	query := `SELECT *
	          FROM nfts_waitlist
	          WHERE token_id = $1`

	var nft nfts.NFTWaitList

	err := nftsDB.conn.QueryRowContext(ctx, query, tokenID).Scan(&nft.TokenID, &nft.CardID, &nft.Wallet, &nft.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return nft, nfts.ErrNoNFT.Wrap(err)
	}

	return nft, ErrNFTs.Wrap(err)
}

// GetLast returns id of last inserted token.
func (nftsDB *nftsDB) GetLast(ctx context.Context) (int, error) {
	query := `SELECT token_id
	          FROM nfts_waitlist
	          ORDER BY token_id DESC
	          LIMIT 1`

	var lastToken int

	err := nftsDB.conn.QueryRowContext(ctx, query).Scan(&lastToken)
	if errors.Is(err, sql.ErrNoRows) {
		return lastToken, nfts.ErrNoNFT.Wrap(err)
	}

	return lastToken, ErrNFTs.Wrap(err)
}

// Update updates signature by token id.
func (nftsDB *nftsDB) Update(ctx context.Context, tokenID int, signature cryptoutils.Signature) error {
	query := `UPDATE nfts_waitlist
	          SET password = $1
	          WHERE token_id = $2`

	result, err := nftsDB.conn.ExecContext(ctx, query, signature, tokenID)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if err != nil {
		return ErrNFTs.Wrap(err)
	}
	if rowNum == 0 {
		return nfts.ErrNoNFT.New("nft token does not exist")
	}
	return nil
}

// Delete deletes nfts from wait list by id of token.
func (nftsDB *nftsDB) Delete(ctx context.Context, tokenIDs []int) error {
	query := `DELETE FROM nfts_waitlist
	          WHERE token_id = $1`

	preparedQuery, err := nftsDB.conn.PrepareContext(ctx, query)
	if err != nil {
		return ErrNFTs.Wrap(err)
	}
	defer func() {
		err = preparedQuery.Close()
	}()

	for _, tokenID := range tokenIDs {
		result, err := nftsDB.conn.ExecContext(ctx, query, tokenID)
		if err != nil {
			return ErrNFTs.Wrap(err)
		}

		rowNum, err := result.RowsAffected()
		if err != nil {
			return ErrNFTs.Wrap(err)
		}
		if rowNum == 0 {
			return nfts.ErrNoNFT.New("nft token does not exist")
		}
	}

	return nil
}
