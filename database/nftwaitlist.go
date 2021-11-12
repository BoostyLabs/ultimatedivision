// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts/nftwaitlist"
	"ultimatedivision/pkg/cryptoutils"
)

// ensures that nftwaitlistDB implements nftwaitlist.DB.
var _ nftwaitlist.DB = (*nftwaitlistDB)(nil)

// ErrNFTWaitlist indicates that there was an error in the database.
var ErrNFTWaitlist = errs.Class("ErrNFTWaitlist repository error")

// nftwaitlistDB provide access to nfts DB.
//
// architecture: Database
type nftwaitlistDB struct {
	conn *sql.DB
}

// Create creates nft for wait list in the database.
func (nftwaitlistDB *nftwaitlistDB) Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address) error {
	query := `INSERT INTO nfts_waitlist(card_id, wallet_address, password)
	          VALUES($1,$2,$3)`

	_, err := nftwaitlistDB.conn.ExecContext(ctx, query, cardID, wallet, "")
	return ErrNFTWaitlist.Wrap(err)
}

// List returns all nft for wait list from wait list from database.
func (nftwaitlistDB *nftwaitlistDB) List(ctx context.Context) ([]nftwaitlist.Item, error) {
	query := `SELECT *
	          FROM nfts_waitlist`

	var nftWaitList []nftwaitlist.Item

	rows, err := nftwaitlistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nftWaitList, ErrNFTWaitlist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		var nft nftwaitlist.Item
		err = rows.Scan(&nft.TokenID, &nft.CardID, &nft.Wallet, &nft.Password)
		if err != nil {
			return nftWaitList, ErrNFTWaitlist.Wrap(err)
		}

		nftWaitList = append(nftWaitList, nft)
	}
	if err = rows.Err(); err != nil {
		return nftWaitList, ErrNFTWaitlist.Wrap(err)
	}

	return nftWaitList, ErrNFTWaitlist.Wrap(err)
}

// ListWithoutPassword returns all nft for wait list without password from database.
func (nftwaitlistDB *nftwaitlistDB) ListWithoutPassword(ctx context.Context) ([]nftwaitlist.Item, error) {
	query :=
		`SELECT *
	     FROM nfts_waitlist
	     WHERE password = ''`

	rows, err := nftwaitlistDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrNFTWaitlist.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var nftWaitListWithoutPassword []nftwaitlist.Item
	for rows.Next() {
		var nft nftwaitlist.Item
		if err = rows.Scan(&nft.TokenID, &nft.CardID, &nft.Wallet, &nft.Password); err != nil {
			return nil, ErrNFTWaitlist.Wrap(err)
		}
		nftWaitListWithoutPassword = append(nftWaitListWithoutPassword, nft)
	}

	return nftWaitListWithoutPassword, ErrNFTWaitlist.Wrap(rows.Err())
}

// Get returns nft for wait list by card id.
func (nftwaitlistDB *nftwaitlistDB) Get(ctx context.Context, tokenID int) (nftwaitlist.Item, error) {
	query := `SELECT *
	          FROM nfts_waitlist
	          WHERE token_id = $1`

	var nftWaitList nftwaitlist.Item

	err := nftwaitlistDB.conn.QueryRowContext(ctx, query, tokenID).Scan(&nftWaitList.TokenID, &nftWaitList.CardID, &nftWaitList.Wallet, &nftWaitList.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return nftWaitList, nftwaitlist.ErrNoNFTWaitList.Wrap(err)
	}

	return nftWaitList, ErrNFTWaitlist.Wrap(err)
}

// GetLast returns id of last inserted nft for wait list.
func (nftwaitlistDB *nftwaitlistDB) GetLast(ctx context.Context) (int, error) {
	query := `SELECT token_id
	          FROM nfts_waitlist
	          ORDER BY token_id DESC
	          LIMIT 1`

	var lastToken int

	err := nftwaitlistDB.conn.QueryRowContext(ctx, query).Scan(&lastToken)
	if errors.Is(err, sql.ErrNoRows) {
		return lastToken, nftwaitlist.ErrNoNFTWaitList.Wrap(err)
	}

	return lastToken, ErrNFTWaitlist.Wrap(err)
}

// Delete deletes nft from wait list by id of token.
func (nftwaitlistDB *nftwaitlistDB) Delete(ctx context.Context, tokenIDs []int) error {
	query := `DELETE FROM nfts_waitlist
	          WHERE token_id = $1`

	preparedQuery, err := nftwaitlistDB.conn.PrepareContext(ctx, query)
	if err != nil {
		return ErrNFTWaitlist.Wrap(err)
	}
	defer func() {
		err = preparedQuery.Close()
	}()

	for _, tokenID := range tokenIDs {
		result, err := nftwaitlistDB.conn.ExecContext(ctx, query, tokenID)
		if err != nil {
			return ErrNFTWaitlist.Wrap(err)
		}

		rowNum, err := result.RowsAffected()
		if err != nil {
			return ErrNFTWaitlist.Wrap(err)
		}
		if rowNum == 0 {
			return nftwaitlist.ErrNoNFTWaitList.New("nft token does not exist")
		}
	}

	return nil
}
