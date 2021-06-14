// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package ultimatedivision

import (
	"context"

	"ultimatedivision/admin/admins"
	"ultimatedivision/internal/logger"
)

// DB provides access to all databases and database related functionality.
//
// architecture: Master Database.
type DB interface {
	//Admins provided access to admins db.
	Admins() admins.DB

	// Close closes underlying db connection.
	Close() error

	// CreateSchema create tables.
	CreateSchema(ctx context.Context) (err error)
}

// Config is the global configuration for ultimatedivision.
type Config struct {
}

// Peer is the representation of a ultimatedivision.
type Peer struct {
	Config   Config
	Log      logger.Logger
	Database DB
	Admins struct {
		Service *admins.Service
	}
}

func New(logger logger.Logger, config Config,ctx context.Context, db DB) (*Peer, error) {
	peer := &Peer{
		Log:      logger,
		Database: db,
	}

	{ // admins setup
		peer.Admins.Service = admins.NewService(
			peer.Database.Admins(),
			ctx,
		)
	}

	return peer, nil

}