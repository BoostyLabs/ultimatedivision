// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/lootboxes"
)

// ErrLootBoxes indicates that there was an error in the database.
var ErrLootBoxes = errs.Class("lootboxes repository error")

// lootboxesDB provide access to lootboxes DB.
//
// architecture: Database
type lootboxesDB struct {
	conn *sql.DB
}

// Create creates opened lootbox in db.
func (lootboxesDB lootboxesDB) Create(ctx context.Context, lootBox lootboxes.UserLootBox) error {
	query := `INSERT INTO user_lootboxes(id,user_id, lootbox_id)
              VALUES($1,$2,$3)`

	_, err := lootboxesDB.conn.QueryContext(ctx, query, lootBox.ID, lootBox.UserID, lootBox.LootBoxID)

	return ErrLootBoxes.Wrap(err)
}

// CreateCards inserts cards getting from lootbox.
func (lootboxesDB lootboxesDB) CreateCards(ctx context.Context, loot lootboxes.UserLoot) error {
	query := `INSERT INTO lootboxes_cards(id,card_id)
              VALUES($1,$2)`

	_, err := lootboxesDB.conn.QueryContext(ctx, query, loot.ID, loot.CardID)

	return ErrLootBoxes.Wrap(err)
}

// Get returns all ids of opened lootboxes by userID.
func (lootboxesDB lootboxesDB) Get(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	query := `SELECT id
              FROM user_lootboxes
              WHERE user_id = $1`

	rows, err := lootboxesDB.conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, ErrLootBoxes.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []uuid.UUID
	for rows.Next() {
		var id uuid.UUID

		err = rows.Scan(&id)
		if err != nil {
			return nil, ErrLootBoxes.Wrap(err)
		}

		data = append(data, id)
	}

	return data, nil
}

// GetCards returns all ids of cards received from LootBox.
func (lootboxesDB lootboxesDB) GetCards(ctx context.Context, lootBoxID uuid.UUID) ([]uuid.UUID, error) {
	query := `SELECT card_id
              FROM lootboxes_cards
              WHERE id = $1`

	rows, err := lootboxesDB.conn.QueryContext(ctx, query, lootBoxID)
	if err != nil {
		return nil, ErrLootBoxes.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []uuid.UUID
	for rows.Next() {
		var cardID uuid.UUID

		err = rows.Scan(&cardID)
		if err != nil {
			return nil, ErrLootBoxes.Wrap(err)
		}

		data = append(data, cardID)
	}

	return data, nil
}
