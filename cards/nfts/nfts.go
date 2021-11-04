// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
)

// ErrNoNFT indicates that nft token does not exist.
var ErrNoNFT = errs.Class("nft does not exist")

// Storage is exposing access to nfts storage.
//
// architecture: Storage
type Storage interface {
	// Save saves nft in the storage.
	Save(ctx context.Context, nft NFT) error
}

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create creates nft token in the database.
	Create(ctx context.Context, cardID uuid.UUID, wallet cryptoutils.Address) error
	// Get returns nft token by card id.
	Get(ctx context.Context, tokenID int) (NFTWaitList, error)
	// GetLast returns id of last inserted token.
	GetLast(ctx context.Context) (int, error)
	// List returns all nft token from wait list from database.
	List(ctx context.Context) ([]NFTWaitList, error)
	// Delete deletes nft from wait list by id of token.
	Delete(ctx context.Context, tokenIDs []int) error
}

// NFTWaitList describes list of nft tokens entity.
type NFTWaitList struct {
	TokenID int                 `json:"tokenId"`
	CardID  uuid.UUID           `json:"cardId"`
	Wallet  cryptoutils.Address `json:"wallet"`
}

// NFT entity describes nft token format erc-721.
type NFT struct {
	Attributes  []Attribute `json:"attributes"`
	Description string      `json:"description"`
	ExternalURL string      `json:"external_url"`
	Image       string      `json:"image"`
	Name        string      `json:"name"`
}

// Attribute entity describes attributes for nft token.
type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
	MaxValue  interface{} `json:"max_value"`
}

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100
