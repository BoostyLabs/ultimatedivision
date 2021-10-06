// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftdrop

import (
	"context"
	"errors"
	"net"
	landingserver "nftdrop/landing"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"nftdrop/internal/logger"
	"nftdrop/whitelist"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// Whitelist provides access to whitelist db.
	Whitelist() whitelist.DB

	// Close closes underlying db connection.
	Close() error

	// CreateSchema create tables.
	CreateSchema(ctx context.Context) error
}

// Config is the global configuration for nftdrop.
type Config struct {
	Landing struct {
		Server landingserver.Config `json:"server"`
	} `json:"landing"`
}

// Peer is the representation of a nftdrop.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB

	// Landing web server with web UI.
	Landing struct {
		Listener net.Listener
		Endpoint *landingserver.Server
	}

	// exposes whitelist relates logic.
	Whitelist struct {
		Service *whitelist.Service
	}
}

// New is a constructor for nftdrop.Peer.
func New(logger logger.Logger, config Config, db DB) (peer *Peer, err error) {
	peer = &Peer{
		Log:      logger,
		Database: db,
	}

	{ // landing setup
		peer.Landing.Listener, err = net.Listen("tcp", config.Landing.Server.Address)
		if err != nil {
			return nil, err
		}

		peer.Landing.Endpoint = landingserver.NewServer(
			config.Landing.Server,
			logger,
			peer.Landing.Listener,
		)
	}

	peer.Whitelist.Service = whitelist.NewService(
		peer.Database.Whitelist(),
	)

	return peer, nil
}

// Run runs nftdrop.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	// start nftdrop servers as a separate goroutines.
	group.Go(func() error {
		return ignoreCancel(peer.Landing.Endpoint.Run(ctx))
	})

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group

	errlist.Add(peer.Landing.Endpoint.Close())

	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
