// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardgenerator

import (
	"ultimatedivision/cardgenerator/avatarcards"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/internal/logger"
)

// Config is the global configuration for cardgenerator.
type Config struct {
	AvatarCards struct {
		avatarcards.Config
	} `json:"avatarCards"`
}

// Peer is the representation of a cardgenerator.
type Peer struct {
	Config Config
	Log    logger.Logger

	// exposes cards related logic.
	Cards struct {
		Service *cards.Service
	}

	// exposes avatars related logic.
	Avatars struct {
		Service *avatars.Service
	}

	// exposes avatar cards related logic.
	AvatarCards struct {
		Service *avatarcards.Service
	}
}

// New is a constructor for cardgenerator.Peer.
func New(logger logger.Logger, config Config) (peer *Peer, err error) {
	peer = &Peer{
		Log:    logger,
		Config: config,
	}

	{ // Avatars setup
		peer.Avatars.Service = avatars.NewService(
			nil,
			config.AvatarCards.AvatarConfig,
		)
	}

	{ // cards setup
		peer.Cards.Service = cards.NewService(
			nil,
			config.AvatarCards.CardConfig,
			peer.Avatars.Service,
		)
	}

	{ // avatar cards setup
		peer.AvatarCards.Service = avatarcards.NewService(
			config.AvatarCards.Config,
			peer.Cards.Service,
			peer.Avatars.Service,
		)
	}

	return peer, nil
}
