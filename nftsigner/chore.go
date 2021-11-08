// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/sync"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log  logger.Logger
	nfts *nfts.Service
	loop *sync.Cycle
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, db DB) *Chore {
	return &Chore{
		log:  log,
		loop: sync.NewCycle(config.RenewalInterval),
		nfts: nfts.NewService(nfts.Config{}, nil, nil, nil, db.NFTs(),
		),
	}
}

// Run starts the chore for signing unsigned nft token from ultimatedivision.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.loop.Run(ctx, func(ctx context.Context) error {
		return ChoreError.Wrap(err)
	})
}

// Stop stops the chore for signing unsigned nft token from ultimatedivision.
func (chore *Chore) Stop(ctx context.Context) {
	chore.loop.Stop()
}
