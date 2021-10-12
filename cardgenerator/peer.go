// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardgenerator

import (
	"context"
	"errors"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/internal/logger"
)

// Config is the global configuration for nftdrop.
type Config struct {
}

// Peer is the representation of a nftdrop.
type Peer struct {
	Config   Config
	Log      logger.Logger
}

// New is a constructor for cardgenerator.Peer.
func New(logger logger.Logger, config Config) (peer *Peer, err error) {
	peer = &Peer{
		Log:      logger,
		Config: config,
	}

	return peer, nil
}

// Run runs cardgenerator.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group
	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
