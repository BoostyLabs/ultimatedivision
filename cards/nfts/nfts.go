// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"

	"github.com/google/uuid"

	"ultimatedivision/pkg/cryptoutils"
)

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create creates nft token in the database.
	Create(ctx context.Context, nft NFT) error
	// List returns all nft token from database.
	List(ctx context.Context) ([]NFT, error)
}

// NFT entity describes values released nft token.
type NFT struct {
	CardID        uuid.UUID           `json:"cardId"`
	TokenID       int                 `json:"tokenId"`
	Chain         cryptoutils.Сhain   `json:"chain"`
	WalletAddress cryptoutils.Address `json:"walletAddress"`
}

// MaxValueGameParameter indicates that max value game parameter is 100.
const MaxValueGameParameter = 100

// Config defines values needed by create nft.
type Config struct {
	Description string `json:"description"`
	ExternalURL string `json:"externalUrl"`
}
