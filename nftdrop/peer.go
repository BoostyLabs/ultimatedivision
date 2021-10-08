// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftdrop

import (
	"context"
	"errors"
	"net"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/admin/adminauth"
	"ultimatedivision/nftdrop/admin/admins"
	"ultimatedivision/nftdrop/admin/adminserver"
	"ultimatedivision/nftdrop/server"
	"ultimatedivision/nftdrop/whitelist"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// Whitelist provides access to whitelist db.
	Whitelist() whitelist.DB

	// Admins provides access to admins db.
	Admins() admins.DB

	// Close closes underlying db connection.
	Close() error

	// CreateSchema create tables.
	CreateSchema(ctx context.Context) error
}

// Config is the global configuration for nftdrop.
type Config struct {
	Landing struct {
		Server server.Config `json:"server"`
	} `json:"landing"`

	Admins struct {
		Server adminserver.Config `json:"server"`
		Auth   struct {
			TokenAuthSecret string `json:"tokenAuthSecret"`
		} `json:"auth"`
	} `json:"admins"`
}

// Peer is the representation of a nftdrop.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB

	// exposes whitelist related logic.
	Whitelist struct {
		Service *whitelist.Service
	}

	// exposes admins relates logic.
	Admins struct {
		Service *admins.Service
		Auth    *adminauth.Service
	}

	// Landing web server with web UI.
	Landing struct {
		Listener net.Listener
		Endpoint *server.Server
	}

	// Admin web server server with web UI.
	Admin struct {
		Listener net.Listener
		Endpoint *adminserver.Server
	}
}

// New is a constructor for nftdrop.Peer.
func New(logger logger.Logger, config Config, db DB) (peer *Peer, err error) {
	peer = &Peer{
		Log:      logger,
		Database: db,
	}

	{ // whitelist setup
		peer.Whitelist.Service = whitelist.NewService(peer.Database.Whitelist())
	}

	{ // admins setup
		peer.Admins.Service = admins.NewService(
			peer.Database.Admins(),
		)
		peer.Admins.Auth = adminauth.NewService(
			peer.Database.Admins(),
			auth.TokenSigner{
				Secret: []byte(config.Admins.Auth.TokenAuthSecret),
			},
		)
	}

	{ // admin setup
		peer.Admin.Listener, err = net.Listen("tcp", config.Admins.Server.Address)
		if err != nil {
			return nil, err
		}

		peer.Admin.Endpoint, err = adminserver.NewServer(
			config.Admins.Server,
			logger,
			peer.Admin.Listener,
			peer.Admins.Auth,
			peer.Admins.Service,
			peer.Whitelist.Service,
		)
		if err != nil {
			return nil, err
		}
	}

	{ // landing setup
		peer.Landing.Listener, err = net.Listen("tcp", config.Landing.Server.Address)
		if err != nil {
			return nil, err
		}

		peer.Landing.Endpoint = server.NewServer(
			config.Landing.Server,
			logger,
			peer.Landing.Listener,
			peer.Whitelist.Service,
		)
	}

	return peer, nil
}

// Run runs nftdrop.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	// start nftdrop servers as a separate goroutines.
	group.Go(func() error {
		return ignoreCancel(peer.Landing.Endpoint.Run(ctx))
	})
	group.Go(func() error {
		return ignoreCancel(peer.Admin.Endpoint.Run(ctx))
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
