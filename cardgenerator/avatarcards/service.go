// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"context"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/pkg/filereading"
)

// ErrAvatarCard indicated that there was an error in service.
var ErrAvatarCard = errs.Class("avatar card service error")

// Service is handling avatars related logic.
//
// architecture: Service
type Service struct {
	config  Config
	cards   *cards.Service
	avatars *avatars.Service
}

// NewService is a constructor for avatar card service.
func NewService(config Config, cards *cards.Service, avatars *avatars.Service) *Service {
	return &Service{
		config:  config,
		cards:   cards,
		avatars: avatars,
	}
}

// Generate generates avatar cards.
func (service *Service) Generate(ctx context.Context, count int) ([]AvatarCards, error) {
	var (
		err         error
		avatarCards []AvatarCards
	)

	id := uuid.New()
	percentageQualities := []int{
		service.config.PercentageQualities.Wood,
		service.config.PercentageQualities.Silver,
		service.config.PercentageQualities.Gold,
		service.config.PercentageQualities.Diamond,
	}

	allNames := make([]string, 0, count)

	for i := 0; i < count; i++ {
		var avatarCard AvatarCards
		var avatar avatars.Avatar
		if avatarCard.Card, err = service.cards.Generate(ctx, id, percentageQualities); err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		if avatar, err = service.avatars.Generate(ctx, avatarCard.Card.ID, avatarCard.Card.IsTattoo, strconv.Itoa(i)); err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		avatarCard.OriginalURL = avatar.OriginalURL

		avatarCard.PlayerName, err = generateName(service.config.PathToNamesDataset, allNames)
		if err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		allNames = append(allNames, avatarCard.PlayerName)

		avatarCards = append(avatarCards, avatarCard)
	}

	return avatarCards, nil
}

// generateName generates name of card.
func generateName(path string, names []string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", ErrAvatarCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, file.Close())
	}()

	rand.Seed(time.Now().UTC().UnixNano())

	totalCount, err := filereading.CountLines(file)
	if err != nil {
		return "", ErrAvatarCard.Wrap(err)
	}

	randomNum := rand.Intn(totalCount) + 1

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", ErrAvatarCard.Wrap(err)
	}

	name, err := filereading.ReadCertainLine(file, randomNum)
	if err != nil {
		return "", ErrAvatarCard.Wrap(err)
	}

	return name, ErrAvatarCard.Wrap(err)
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

	for i := 0; i < count; i++ {
		var avatarCard AvatarCards
		if avatarCard.Card, err = service.cards.Generate(ctx, id, percentageQualities); err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		avatar, err := service.avatars.Generate(ctx, avatarCard.Card.ID, avatarCard.Card.IsTattoo, avatarCard.Card.ID.String())
		if err != nil {
			return nil, ErrAvatarCard.Wrap(err)
		}

		avatars = append(avatars, avatar)
	}

	return avatars, nil
}
