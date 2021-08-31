// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
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
func (lootboxesDB *lootboxesDB) Create(ctx context.Context, lootBox lootboxes.LootBox) error {
	tx, err := lootboxesDB.conn.BeginTx(ctx, nil)
	if err != nil {
		return ErrLootBoxes.Wrap(err)
	}
	query := `INSERT INTO lootboxes(lootbox_id, user_id, lootbox_name)
              VALUES($1,$2,$3)`

	_, err = lootboxesDB.conn.ExecContext(ctx, query, lootBox.UserID, lootBox.UserID, lootBox.Type)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return ErrLootBoxes.Wrap(err)
		}
		return ErrLootBoxes.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return ErrLootBoxes.Wrap(err)
	}

	return ErrLootBoxes.Wrap(err)
}

// Delete deletes opened lootbox by user in db.
func (lootboxesDB *lootboxesDB) Delete(ctx context.Context, lootBox lootboxes.LootBox) error {
	query := `DELETE FROM lootboxes
              WHERE user_id = $1 and lootbox_id = $2`

	_, err := lootboxesDB.conn.ExecContext(ctx, query, lootBox.UserID, lootBox.LootBoxID)

	return ErrLootBoxes.Wrap(err)
}

// GetByUserID returns all users loot box by id.
func (lootboxesDB *lootboxesDB) GetByUserID(ctx context.Context, id uuid.UUID) ([]lootboxes.LootBox, error) {
	query := `SELECT user_id, lootbox_id, lootbox_name
              FROM lootboxes
              WHERE user_id = $1`

	rows, err := lootboxesDB.conn.QueryContext(ctx, query, id)
	if err != nil{
		return nil, ErrLootBoxes.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var userLootBoxes []lootboxes.LootBox

	for rows.Next(){
		var userLootBox lootboxes.LootBox

		err = rows.Scan(&userLootBox.UserID, &userLootBox.LootBoxID, &userLootBox.Type)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, lootboxes.ErrNoLootBox.Wrap(err)
			}

			return nil, ErrLootBoxes.Wrap(err)
		}

		userLootBoxes = append(userLootBoxes, userLootBox)
	}

	return userLootBoxes, nil
}
