// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"

	"github.com/google/uuid"
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

// Config defines configuration for LootBox.
type Config struct {
	Cost     int         `json:"cost"`
	CardsNum int         `json:"cardsNum"`
	Wood     Probability `json:"wood"`
	Silver   Probability `json:"silver"`
	Gold     Probability `json:"gold"`
	Diamond  Probability `json:"diamond"`
}

// NewService is a constructor for lootboxes service.
func NewService(lootboxes DB, config Config) *Service {
	return &Service{
		lootboxes: lootboxes,
		config:    config,
	}
}

// Create creates opened LootBox.
func (service *Service) Create(ctx context.Context, userID uuid.UUID, lootBoxID uuid.UUID) error {
	openedLootBox := OpenedLootBoxes{
		UserID:    userID,
		LootBoxID: lootBoxID,
	}

	err := service.lootboxes.Create(ctx, openedLootBox)

	// TODO: call create cards method and return slice of generated cards.
	// TODO: check if user has enough money for lootbox.

	return ErrLootBoxes.Wrap(err)
}
