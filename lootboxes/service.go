// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"
	"ultimatedivision/cards"

	"github.com/zeebo/errs"
)

// ErrLootBoxes indicates that there was an error in the service.
var ErrLootBoxes = errs.Class("lootboxes service error")

// Service is handling lootboxes related logic.
//
// architecture: Service
type Service struct {
	lootboxes DB
	config    Config
}

// NewService is a constructor for lootboxes service.
func NewService(lootboxes DB, config Config) *Service {
	return &Service{
		lootboxes: lootboxes,
		config:    config,
	}
}

// Create creates LootBox.
func (service *Service) Create(ctx context.Context, userLootBox UserLootBoxes) error {
	err := service.lootboxes.Create(ctx, userLootBox)

	return ErrLootBoxes.Wrap(err)
}

// Open opens lootbox by user.
func (service *Service) Open(ctx context.Context, userLootBox UserLootBoxes) ([]cards.Card, error) {
	probabilities := []int{service.config.Wood, service.config.Silver, service.config.Gold, service.config.Diamond}

	var lootBoxCards []cards.Card

	for i := 0; i < service.config.CardsNum; i++ {
		card, err := cards.Service.Create(cards.Service{}, ctx, userLootBox.UserID, probabilities)
		if err != nil {
			return lootBoxCards, ErrLootBoxes.Wrap(err)
		}

		lootBoxCards = append(lootBoxCards, card)
	}

	err := service.lootboxes.Delete(ctx, userLootBox)

	return lootBoxCards, ErrLootBoxes.Wrap(err)
}
