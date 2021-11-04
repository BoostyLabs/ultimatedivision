// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nfts

import (
	"context"

	"github.com/google/uuid"
)

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create creates nft token in the database.
	Create(ctx context.Context, tokenID int, cardID uuid.UUID) error
	// Get returns nft token by card id.
	Get(ctx context.Context, tokenID int, cardID uuid.UUID)
	// GetLast returns id of last inserted token.
	GetLast(ctx context.Context) int
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
