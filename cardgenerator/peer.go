// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cardgenerator

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"ultimatedivision/cardgenerator/avatarcards"
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
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

	quantityOfCard int

	// exposes cards related logic.
	Cards struct {
		Service *cards.Service
	}

	// exposes avatars related logic.
	Avatars struct {
		Service *avatars.Service
	}

	// exposes nfts related logic.
	NFTs struct {
		Service *nfts.Service
	}

	// exposes avatar cards related logic.
	AvatarCards struct {
		Service *avatarcards.Service
	}
}

// New is a constructor for cardgenerator.Peer.
func New(logger logger.Logger, config Config, quantityOfCard int) (peer *Peer, err error) {
	peer = &Peer{
		Log:            logger,
		Config:         config,
		quantityOfCard: quantityOfCard,
	}

	{ // cards setup
		peer.Cards.Service = cards.NewService(
			nil,
			config.AvatarCards.CardConfig,
		)
	}

	{ // avatars setup
		peer.Avatars.Service = avatars.NewService(
			nil,
			nil,
			config.AvatarCards.AvatarConfig,
		)
	}

	{ // nfts setup
		peer.NFTs.Service = nfts.NewService(
			config.AvatarCards.NFTConfig,
			nil,
			peer.Cards.Service,
			peer.Avatars.Service,
			nil,
		)
	}

	{ // avatar cards setup
		peer.AvatarCards.Service = avatarcards.NewService(
			config.AvatarCards.Config,
			peer.Cards.Service,
			peer.Avatars.Service,
			peer.NFTs.Service,
		)
	}

	return peer, nil
}

// Generate initiates generation of avatar cards.
func (peer *Peer) Generate(ctx context.Context) error {
	for i := 0; i < peer.quantityOfCard; i++ {

		allNames := make(map[string]struct{}, peer.quantityOfCard)
		for len(allNames) <= peer.quantityOfCard {
			if err := peer.AvatarCards.Service.GenerateName(peer.Config.AvatarCards.PathToNamesDataset, allNames); err != nil {
				return err
			}
		}

		var playerName string
		for name := range allNames {
			playerName = name
			delete(allNames, name)
			break
		}

		nft, err := peer.AvatarCards.Service.Generate(ctx, i, playerName)
		if err != nil {
			return err
		}

		file, err := json.MarshalIndent(nft, "", " ")
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile(filepath.Join(peer.Config.AvatarCards.PathToOutputJSONFile, strconv.Itoa(i+1)+".json"), file, 0644); err != nil {
			return err
		}
	}

	return nil
}

// TestGenerate initiates generation test version of avatar cards.
func (peer *Peer) TestGenerate(ctx context.Context) error {
	avatars, err := peer.AvatarCards.Service.TestGenerate(ctx, peer.quantityOfCard)
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(avatars, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(peer.Config.AvatarCards.PathToOutputJSONFile, "test.json"), file, 0644)
}
