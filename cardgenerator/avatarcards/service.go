// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/pkg/fileutils"
	"ultimatedivision/pkg/nft"
)

// ErrCardWithLinkToAvatar indicated that there was an error in service.
var ErrCardWithLinkToAvatar = errs.Class("card with link to avatar service error")

// Service is handling cards with link to avatars related logic.
//
// architecture: Service
type Service struct {
	config  Config
	cards   *cards.Service
	avatars *avatars.Service
	nfts    *nfts.Service
}

// NewService is a constructor for card with link to avatar service.
func NewService(config Config, cards *cards.Service, avatars *avatars.Service, nfts *nfts.Service) *Service {
	return &Service{
		config:  config,
		cards:   cards,
		avatars: avatars,
		nfts:    nfts,
	}
}

// Generate generates cards with avatar link.
func (service *Service) Generate(ctx context.Context, nameFile int, playerName string) (nft.NFT, error) {
	percentageQualities := []int{
		service.config.PercentageQualities.Wood,
		service.config.PercentageQualities.Silver,
		service.config.PercentageQualities.Gold,
		service.config.PercentageQualities.Diamond,
	}

	card, err := service.cards.Generate(ctx, uuid.Nil, percentageQualities)
	if err != nil {
		return nft.NFT{}, ErrCardWithLinkToAvatar.Wrap(err)
	}
	card.PlayerName = playerName

	avatar, err := service.avatars.Generate(ctx, card, nameFile+1)
	if err != nil {
		return nft.NFT{}, ErrCardWithLinkToAvatar.Wrap(err)
	}

	nftCard, err := service.nfts.Generate(ctx, card, avatar.OriginalURL, fmt.Sprintf(service.config.NFTConfig.ExternalURL, nameFile+1))
	if err != nil {
		return nft.NFT{}, ErrCardWithLinkToAvatar.Wrap(err)
	}

	return nftCard, nil
}

// TestGenerate generates test version avatar cards.
func (service *Service) TestGenerate(ctx context.Context, count int) ([]avatars.Avatar, error) {
	var (
		err     error
		avatars []avatars.Avatar
	)

	id := uuid.New()
	percentageQualities := []int{
		service.config.PercentageQualities.Wood,
		service.config.PercentageQualities.Silver,
		service.config.PercentageQualities.Gold,
		service.config.PercentageQualities.Diamond,
	}

	allNames := make(map[string]struct{}, count)

	for i := 0; i < count; i++ {
		var avatarCard CardWithLinkToAvatar
		if avatarCard.Card, err = service.cards.Generate(ctx, id, percentageQualities); err != nil {
			return nil, ErrCardWithLinkToAvatar.Wrap(err)
		}

		for len(allNames) < count {
			if err = service.GenerateName(service.config.PathToNamesDataset, allNames); err != nil {
				return nil, ErrCardWithLinkToAvatar.Wrap(err)
			}
		}

		for name := range allNames {
			avatarCard.PlayerName = name
			delete(allNames, name)
			break
		}

		avatar, err := service.avatars.Generate(ctx, avatarCard.Card, i+1)
		if err != nil {
			return nil, ErrCardWithLinkToAvatar.Wrap(err)
		}

		avatars = append(avatars, avatar)
	}

	return avatars, nil
}

// GenerateName generates name of card.
func (service *Service) GenerateName(path string, names map[string]struct{}) error {
	file, err := os.Open(path)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, file.Close())
	}()

	rand.Seed(time.Now().UTC().UnixNano())

	totalCount, err := fileutils.CountLines(file)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}

	randomNum := rand.Intn(totalCount) + 1

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}

	name, err := fileutils.ReadLine(file, randomNum)
	if err != nil {
		return ErrCardWithLinkToAvatar.Wrap(err)
	}

	names[name] = struct{}{}

	return ErrCardWithLinkToAvatar.Wrap(err)
}
