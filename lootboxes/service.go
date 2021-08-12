// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// ErrLootBoxes indicates that there was an error in the service.
var ErrLootBoxes = errs.Class("lootboxes service error")

// Service is handling lootboxes related logic.
//
// architecture: Service
type Service struct {
	config    Config
	lootboxes DB
	cards     *cards.Service
}

// NewService is a constructor for lootboxes service.
func NewService(config Config, lootboxes DB, cards *cards.Service) *Service {
	return &Service{
		config:    config,
		lootboxes: lootboxes,
		cards:     cards,
	}
}

// Create creates LootBox.
func (service *Service) Create(ctx context.Context, userLootBox LootBox) error {
	err := service.lootboxes.Create(ctx, userLootBox)

	return ErrLootBoxes.Wrap(err)
}

// Open opens lootbox by user.
func (service *Service) Open(ctx context.Context, userLootBox LootBox) ([]cards.Card, error) {
	cardsNum := 0
	probabilities := make([]int, 0, 5)

	if userLootBox.Name == RegularBox {
		cardsNum = service.config.ConfigRegularBox.CardsNum
		probabilities = []int{service.config.ConfigRegularBox.Wood, service.config.ConfigRegularBox.Silver, service.config.ConfigRegularBox.Gold, service.config.ConfigRegularBox.Diamond}
	} else if userLootBox.Name == UDReleaseCelebrationBox {
		cardsNum = service.config.ConfigUDReleaseCelebrationBox.CardsNum
		probabilities = []int{service.config.ConfigUDReleaseCelebrationBox.Wood, service.config.ConfigUDReleaseCelebrationBox.Silver, service.config.ConfigUDReleaseCelebrationBox.Gold, service.config.ConfigUDReleaseCelebrationBox.Diamond}
	}

	var lootBoxCards []cards.Card

	for i := 0; i < cardsNum; i++ {
		card, err := service.cards.Create(ctx, userLootBox.UserID, probabilities)
		if err != nil {
			return lootBoxCards, ErrLootBoxes.Wrap(err)
		}

		lootBoxCards = append(lootBoxCards, card)
	}

	err := service.lootboxes.Delete(ctx, userLootBox)

	return lootBoxCards, ErrLootBoxes.Wrap(err)
}
