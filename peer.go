// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package ultimatedivision

import (
	"context"
	"errors"
	"net"
	"ultimatedivision/users/userauth"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"ultimatedivision/admin/adminauth"
	"ultimatedivision/admin/admins"
	"ultimatedivision/admin/adminserver"
	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/console/consoleserver"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	// Admins provides access to admins db.
	Admins() admins.DB
	// Users provides access to users db.
	Users() users.DB

	// Cards provides access to cards db.
	Cards() cards.DB

	// Clubs provides access to clubs db.
	Clubs() clubs.DB

	// Close closes underlying db connection.
	Close() error

	// CreateSchema create tables.
	CreateSchema(ctx context.Context) error
}

// Config is the global configuration for ultimatedivision.
type Config struct {
	Admins struct {
		Server adminserver.Config `json:"server"`
		Auth   struct {
			TokenAuthSecret string `json:"tokenAuthSecret"`
		} `json:"auth"`
	}

	Users struct {
		// Server userserver.Config `json:"server"`
		Auth struct {
			TokenAuthSecret string `json:"tokenAuthSecret"`
		} `json:"auth"`
	}

	Consoles struct {
		Server consoleserver.Config `json:"server"`
	}
}

// Peer is the representation of a ultimatedivision.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB

	// exposes admins relates logic.
	Admins struct {
		Service *admins.Service
		Auth    *adminauth.Service
	}

	// exposes users related logic.
	Users struct {
		Service *users.Service
		Auth    *userauth.Service
	}

	// exposes cards related logic.
	Cards struct {
		Service *cards.Service
	}

	// exposes clubs related logic
	Clubs struct {
		Service *clubs.Service
	}

	// Admin web server server with web UI.
	Admin struct {
		Listener net.Listener
		Endpoint *adminserver.Server
	}

	// Console web server server with web UI.
	Console struct {
		Listener net.Listener
		Endpoint *consoleserver.Server
	}
}

// New is a constructor for ultimatedivision.Peer.
func New(logger logger.Logger, config Config, db DB) (peer *Peer, err error) {
	peer = &Peer{
		Log:      logger,
		Database: db,
	}

	{ // users setup
		peer.Users.Service = users.NewService(
			peer.Database.Users(),
		)
		peer.Users.Auth = userauth.NewService(
			peer.Database.Users(),
			auth.TokenSigner{
				Secret: []byte(config.Users.Auth.TokenAuthSecret),
			},
		)
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

	{ // cards setup
		peer.Cards.Service = cards.NewService(
			peer.Database.Cards(),
		)
	}

	{ // clubs setup
		peer.Clubs.Service = clubs.NewService(
			peer.Database.Clubs(),
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
			peer.Users.Service,
			peer.Cards.Service,
		)
		if err != nil {
			return nil, err
		}
	}

	{ // console setup
		peer.Console.Listener, err = net.Listen("tcp", config.Consoles.Server.Address)
		if err != nil {
			return nil, err
		}

		peer.Console.Endpoint, err = consoleserver.NewServer(
			config.Consoles.Server,
			logger,
			peer.Console.Listener,
			peer.Cards.Service,
		)
		if err != nil {
			return nil, err
		}
	}

	return peer, nil
}

// Run runs ultimatedivision.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	// start ultimatedivision servers as a separate goroutines.
	group.Go(func() error {
		return ignoreCancel(peer.Admin.Endpoint.Run(ctx))
	})
	group.Go(func() error {
		return ignoreCancel(peer.Console.Endpoint.Run(ctx))
	})

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group

	errlist.Add(peer.Admin.Endpoint.Close())
	errlist.Add(peer.Console.Endpoint.Close())

	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
