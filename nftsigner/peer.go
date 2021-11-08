// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"time"

	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// NFTs provides access to nfts db.
	NFTs() nfts.DB
}

// Config is the global configuration for nftsigner.
type Config struct {
	RenewalInterval time.Duration `json:"renewalInterval"`
}

// Peer is the representation of a nftsigner.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB

	Chore *Chore
}

// New is a constructor for nftsigner.Peer.
func New(logger logger.Logger, config Config, database DB) (peer *Peer, err error) {
	peer = &Peer{
		Log:    logger,
		Config: config,
	}

	{
		peer.Chore = NewChore(logger, config, database.NFTs())
	}

	return peer, nil
}

// Run runs ultimatedivision.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return ignoreCancel(peer.Chore.Run(ctx))
	})
	return group.Wait()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
