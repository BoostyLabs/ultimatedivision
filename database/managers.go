// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/managers"
)

// ensures that managersDB implements managers.DB.
var _ managers.DB = (*managersDB)(nil)

// ErrManagers indicates that there was an error in the database.
var ErrManagers = errs.Class("managers repository error")

// managersDB provide access to manager DB.
//
// architecture: Database
type managersDB struct {
	conn *sql.DB
}

// Create creates manager in the database.
func (managersDB *managersDB) Create(ctx context.Context, manager managers.Manager) error {
	query := `INSERT INTO managers(user_id, club_id, ended_at)
	          VALUES($1,$2,$3)`

	_, err := managersDB.conn.ExecContext(ctx, query, manager.UserID, manager.ClubID, manager.EndedAt)

	return ErrManagers.Wrap(err)
}

// List returns all managers from the database.
func (managersDB *managersDB) List(ctx context.Context) ([]managers.Manager, error) {
	query := `SELECT user_id, club_id, ended_at
	          FROM managers`

	rows, err := managersDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrManagers.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allManagers []managers.Manager
	for rows.Next() {
		var manager managers.Manager
		err = rows.Scan(&manager.UserID, &manager.ClubID, &manager.EndedAt)
		if err != nil {
			return nil, ErrManagers.Wrap(err)
		}

		allManagers = append(allManagers, manager)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrManagers.Wrap(err)
	}

	return allManagers, nil
}

// ListByUserID returns clubs managed by user from the database.
func (managersDB *managersDB) ListByUserID(ctx context.Context, userID uuid.UUID) ([]managers.Manager, error) {
	query := `SELECT user_id, club_id, ended_at
	          FROM managers
	          WHERE user_id = $1`

	rows, err := managersDB.conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, ErrManagers.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allManagers []managers.Manager
	for rows.Next() {
		var manager managers.Manager
		err = rows.Scan(&manager.UserID, &manager.ClubID, &manager.EndedAt)
		if err != nil {
			return nil, ErrManagers.Wrap(err)
		}

		allManagers = append(allManagers, manager)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrManagers.Wrap(err)
	}

	return allManagers, nil
}

// Delete deletes manager from the database.
func (managersDB *managersDB) Delete(ctx context.Context, userID, clubID uuid.UUID) error {
	query := `DELETE FROM managers
	          WHERE user_id = $1 and club_id = $2`

	result, err := managersDB.conn.ExecContext(ctx, query, userID, clubID)
	if err != nil {
		return ErrManagers.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if err != nil {
		return ErrManagers.Wrap(err)
	}
	if rowsNum == 0 {
		return managers.ErrNoManager.New("manager does not exist")
	}

	return ErrManagers.Wrap(err)
}
