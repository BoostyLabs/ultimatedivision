// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftwaitlist

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrNoNFTWaitList indicates that nft for wait list does not exist.
var ErrNoNFTWaitList = errs.Class("nft for wait list does not exist")

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create creates nft for wait list in the database.
	Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address) error
	// Get returns nft for wait list by card id.
	Get(ctx context.Context, tokenID int) (NFTWaitList, error)
	// GetLast returns id of last inserted token.
	GetLast(ctx context.Context) (int, error)
	// List returns all nft tokens from wait list from database.
	List(ctx context.Context) ([]NFTWaitList, error)
	// ListWithoutPassword returns nfts for wait list without password from database.
	ListWithoutPassword(ctx context.Context) ([]NFTWaitList, error)
	// Delete deletes nft from wait list by id of token.
	Delete(ctx context.Context, tokenIDs []int) error
}

// NFTWaitList describes wait list of nft tokens entity.
type NFTWaitList struct {
	TokenID  int                   `json:"tokenId"`
	CardID   uuid.UUID             `json:"cardId"`
	Wallet   cryptoutils.Address   `json:"wallet"`
	Password cryptoutils.Signature `json:"password"`
}

// CreateNFT describes body of request for creating nft token.
type CreateNFT struct {
	CardID        uuid.UUID           `json:"cardId"`
	WalletAddress cryptoutils.Address `json:"walletAddress"`
	UserID        uuid.UUID           `json:"userId"`
}
