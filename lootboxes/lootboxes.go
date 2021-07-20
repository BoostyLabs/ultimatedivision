// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoLootBox indicates that lootbox does not exist.
var ErrNoLootBox = errs.Class("lootbox does not exist")

// DB is exposing access to lootboxes db.
//
// architecture: DB
type DB interface {
	// Create creates opened lootbox in db.
	Create(ctx context.Context, lootBox UserLootBox) error
	// CreateCards inserts cards getting from LootBox.
	CreateCards(ctx context.Context, loot UserLoot) error
	// Get returns all ids of opened lootboxes by userID.
	Get(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error)
	// GetCards returns all ids of cards received from LootBox.
	GetCards(ctx context.Context, lootBoxID uuid.UUID) ([]uuid.UUID, error)
}

// LootBox defines types of loot box.
type LootBox struct {
	ID       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Cost     int         `json:"cost"`
	CardsNum int         `json:"cardsNum"`
	Wood     Probability `json:"wood"`
	Silver   Probability `json:"silver"`
	Gold     Probability `json:"gold"`
	Diamond  Probability `json:"diamond"`
}

// Probability defines probability of getting different types of cards.
type Probability float32

// UserLootBox describes all opened LootBoxes by user.
type UserLootBox struct {
	ID        uuid.UUID `json:"id"` // unique id of lootbox.
	UserID    uuid.UUID `json:"userID"`
	LootBoxID uuid.UUID `json:"LootBoxID"` // type of lootbox.
}

// UserLoot describes all cards getting from LootBox.
type UserLoot struct {
	ID     uuid.UUID `json:"id"`
	CardID uuid.UUID `json:"cardID"`
}
