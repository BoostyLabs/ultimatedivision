// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardgenerator

import (
	"context"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/internal/logger"
)

// Config is the global configuration for cardgenerator.
type Config struct {
}

// Peer is the representation of a cardgenerator.
type Peer struct {
	Config Config
	Log    logger.Logger

	quantityOfCards int
}

// New is a constructor for cardgenerator.Peer.
func New(logger logger.Logger, config Config, quantity int) (peer *Peer, err error) {
	peer = &Peer{
		Log:             logger,
		Config:          config,
		quantityOfCards: quantity,
	}

	return peer, nil
}

// Run runs cardgenerator.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, _ := errgroup.WithContext(ctx)

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group
	return errlist.Err()
}
