// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/remotefilestorage/storj"
)

// ErrNoItem indicates that item for wait list does not exist.
var ErrNoItem = errs.Class("item for wait list does not exist")

// DB is exposing access to waitlist db.
//
// architecture: DB
type DB interface {
	// Create creates nft for wait list in the database.
	Create(ctx context.Context, cardID uuid.UUID, wallet evmsignature.Address) error
	// Get returns nft for wait list by token id.
	GetByTokenID(ctx context.Context, tokenID int64) (Item, error)
	// GetByCardID returns nft for wait list by card id.
	GetByCardID(ctx context.Context, cardID uuid.UUID) (Item, error)
	// GetLastTokenID returns id of last inserted token.
	GetLastTokenID(ctx context.Context) (int64, error)
	// List returns all nft tokens from wait list from database.
	List(ctx context.Context) ([]Item, error)
	// ListWithoutPassword returns nfts for wait list without password from database.
	ListWithoutPassword(ctx context.Context) ([]Item, error)
	// Delete deletes nft from wait list by id of token.
	Delete(ctx context.Context, tokenIDs []int64) error
	// Update updates signature to nft token.
	Update(ctx context.Context, tokenID int64, password evmsignature.Signature) error
}

// Item entity describes item fot wait list nfts.
type Item struct {
	TokenID  int64                  `json:"tokenId"`
	CardID   uuid.UUID              `json:"cardId"`
	Wallet   evmsignature.Address   `json:"wallet"`
	Password evmsignature.Signature `json:"password"`
}

// CreateNFT describes body of request for creating nft token.
type CreateNFT struct {
	CardID        uuid.UUID            `json:"cardId"`
	WalletAddress evmsignature.Address `json:"walletAddress"`
	UserID        uuid.UUID            `json:"userId"`
}

// Transaction entity describes password wallet, smart contracts address and token id.
type Transaction struct {
	Password evmsignature.Signature `json:"password"`
	Contract evmsignature.Contract  `json:"contract"`
	TokenID  int64                  `json:"tokenId"`
}

// Config defines values needed by check mint nft in blockchain.
type Config struct {
	WaitListRenewalInterval time.Duration `json:"waitListRenewalInterval"`
	WaitListCheckSignature  time.Duration `json:"waitListCheckSignature"`
	NFTContract             struct {
		Address      evmsignature.Address `json:"address"`
		AddressEvent evmsignature.Hex     `json:"addressEvent"`
	} `json:"nftContract"`
	evmsignature.Contract `json:"contract"`
	AddressNodeServer     string       `json:"addressNodeServer"`
	Storj                 storj.Config `json:"storj"`
	Bucket                string       `json:"bucket"`
	URLToAvatar           string       `json:"urlToAvatar"`
}
