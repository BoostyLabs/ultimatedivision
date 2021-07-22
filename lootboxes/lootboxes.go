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
	Create(ctx context.Context, lootBox OpenedLootBoxes) error
}

// LootBox defines types of loot box.
type LootBox struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Probability defines probability of getting different types of cards.
type Probability float32

// OpenedLootBoxes describes all opened LootBoxes by user.
type OpenedLootBoxes struct {
	UserID    uuid.UUID `json:"userID"`
	LootBoxID uuid.UUID `json:"LootBoxID"`
}
