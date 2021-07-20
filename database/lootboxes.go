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
	query := `INSERT INTO user_lootbox(id,user_id, lootbox_id)
     VALUES($1,$2,$3)`

	_, err := lootboxesDB.conn.QueryContext(ctx,query, lootBox.ID, lootBox.UserID, lootBox.LootBoxID)

	return ErrLootBoxes.Wrap(err)
}

// CreateLoot inserts cards getting from lootbox.
func (lootboxesDB lootboxesDB) CreateLoot(ctx context.Context, loot lootboxes.UserLoot) error {
	query := `INSERT INTO user_loot(id,card_id)
     VALUES($1,$2)`

	_, err := lootboxesDB.conn.QueryContext(ctx, query, loot.ID, loot.CardID)

	return ErrLootBoxes.Wrap(err)
}

// Get returns all id of opened loot boxes.
func (lootboxesDB lootboxesDB) Get(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	query := `Select id
      FROM user_lootbox
     WHERE user_id = $1`

	rows, err := lootboxesDB.conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, ErrLootBoxes.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []lootboxes.UserLootBox

	for rows.Next() {
		var lootBox lootboxes.UserLootBox
		err = rows.Scan(&lootBox.ID)
		if err != nil {
			return nil, ErrLootBoxes.Wrap(err)
		}

		data = append(data, lootBox)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrLootBoxes.Wrap(err)
	}

	var LootsID []uuid.UUID

	for _, value := range data {
		LootsID = append(LootsID, value.ID)
	}

	return LootsID, nil
}

// GetLoot returns all cards received from loot box.
func (lootboxesDB lootboxesDB) GetLoot(ctx context.Context, lootBoxID uuid.UUID) ([]uuid.UUID, error) {
	query := `Select id,card_id
      FROM user_loot
      WHERE id = $1`

	rows,err := lootboxesDB.conn.QueryContext(ctx, query, lootBoxID)
	if err != nil {
		return nil, ErrLootBoxes.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []lootboxes.UserLoot

	for rows.Next(){
		var userLoot lootboxes.UserLoot

		err = rows.Scan(&userLoot.ID, &userLoot.CardID)
		if err != nil {
			return nil , ErrLootBoxes.Wrap(err)
		}

		data = append(data, userLoot)
	}

	var CardsID []uuid.UUID

	for _, value := range data {
		CardsID = append(CardsID, value.CardID)
	}

	return CardsID, nil
}