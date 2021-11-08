// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"

	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/cryptoutils"
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
	RenewalInterval      time.Duration          `json:"renewalInterval"`
	PrivateKey           cryptoutils.PrivateKey `json:"privateKey"`
	SmartContractAddress cryptoutils.Address    `json:"smartContractAddress"`
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
		Log:      logger,
		Config:   config,
		Database: database,
	}

	{ // chore setup
		peer.Chore = NewChore(logger, config, peer.Database.NFTs())
	}

	return peer, nil
}

// Run runs nftsigner.Peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return ignoreCancel(peer.Chore.Run(ctx, peer.Config.SmartContractAddress, peer.Config.PrivateKey))
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
