// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/divisions"
)

// ErrDivisions indicates that there was an error in the database.
var ErrDivisions = errs.Class("divisions repository error")

// divisionsDB provides access to divisions db.
//
// architecture: Database
type divisionsDB struct {
	conn *sql.DB
}

// Create creates a division and writes to the database.
func (divisionsDB *divisionsDB) Create(ctx context.Context, division divisions.Division) error {
	query := `INSERT INTO divisions(id, name, percent_of_passing, created_at) 
	VALUES ($1, $2, $3, $4)`

	_, err := divisionsDB.conn.ExecContext(ctx, query, division.ID, division.Name, division.PercentOfPassing, division.CreatedAt)

	return ErrDivisions.Wrap(err)
}

// List returns all divisions from the data base.
func (divisionsDB *divisionsDB) List(ctx context.Context) ([]divisions.Division, error) {
	query := `SELECT id, name, percent_of_passing, created_at FROM divisions`

	rows, err := divisionsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrDivisions.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var dataDivisions []divisions.Division
	for rows.Next() {
		var division divisions.Division
		err := rows.Scan(&division.ID, &division.Name, &division.PercentOfPassing, &division.CreatedAt)
		if err != nil {
			return nil, ErrDivisions.Wrap(err)
		}

		dataDivisions = append(dataDivisions, division)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrDivisions.Wrap(err)
	}

	return dataDivisions, ErrDivisions.Wrap(err)
}

// Get returns division by id from the data base.
func (divisionsDB *divisionsDB) Get(ctx context.Context, id uuid.UUID) (divisions.Division, error) {
	query := `SELECT id, name, percent_of_passing, created_at FROM divisions WHERE id=$1`
	var division divisions.Division

	row := divisionsDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&division.ID, &division.Name, &division.PercentOfPassing, &division.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return division, divisions.ErrNoDivisions.Wrap(err)
		}

		return division, ErrDivisions.Wrap(err)
	}

	return division, ErrDivisions.Wrap(err)
}

// Delete deletes a division in the database.
func (divisionsDB *divisionsDB) Delete(ctx context.Context, id uuid.UUID) error {
	result, err := divisionsDB.conn.ExecContext(ctx, "DELETE FROM divisions WHERE id=$1", id)
	if err != nil {
		return ErrDivisions.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return divisions.ErrNoDivisions.New("division does not exist")
	}

	return ErrDivisions.Wrap(err)
}
