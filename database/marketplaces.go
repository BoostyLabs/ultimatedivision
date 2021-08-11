// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/marketplaces"
)

// ensures that marketplacesDB implements marketplaces.DB.
var _ marketplaces.DB = (*marketplacesDB)(nil)

// ErrMarketplace indicates that there was an error in the database.
var ErrMarketplace = errs.Class("marketplaces repository error")

// marketplacesDB provides access to marketplaces db.
//
// architecture: Database
type marketplacesDB struct {
	conn *sql.DB
}

// Create creates lot in the db.
func (marketplacesDB *marketplacesDB) Create(ctx context.Context, lot marketplaces.Lot) error {
	query :=
		`INSERT INTO 
			lots(id, item_id, type, user_id, shopper_id, status, start_price, max_price, current_price, start_time, end_time)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		`

	_, err := marketplacesDB.conn.ExecContext(ctx, query,
		lot.ID, lot.ItemID, lot.Type, lot.UserID, lot.ShopperID, lot.Status,
		lot.StartPrice, lot.MaxPrice, lot.CurrentPrice, lot.StartTime, lot.EndTime)

	return ErrMarketplace.Wrap(err)
}

// Get returns lot by id from the data base.
func (marketplacesDB *marketplacesDB) Get(ctx context.Context, id uuid.UUID) (marketplaces.Lot, error) {
	lot := marketplaces.Lot{}
	query :=
		`SELECT 
			id, item_id, type, status, start_price, max_price, current_price, start_time, end_time
        FROM 
            lots
        WHERE 
            id = $1
        `
	err := marketplacesDB.conn.QueryRowContext(ctx, query, id).Scan(
		&lot.ID, &lot.ItemID, &lot.Type, &lot.Status, &lot.StartPrice,
		&lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return lot, marketplaces.ErrNoLot.Wrap(err)
	case err != nil:
		return lot, ErrMarketplace.Wrap(err)
	default:
		return lot, nil
	}
}

// List returns active lots from the data base.
func (marketplacesDB *marketplacesDB) ListActive(ctx context.Context) ([]marketplaces.Lot, error) {
	query :=
		`SELECT 
			id, item_id, type, user_id, shopper_id, status, start_price, max_price, current_price, start_time, end_time 
        FROM 
            lots
		WHERE
			status = $1
        `

	rows, err := marketplacesDB.conn.QueryContext(ctx, query, marketplaces.StatusActive)
	if err != nil {
		return nil, ErrCard.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	lots := []marketplaces.Lot{}
	for rows.Next() {
		lot := marketplaces.Lot{}
		if err = rows.Scan(
			&lot.ID, &lot.ItemID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status,
			&lot.StartPrice, &lot.MaxPrice, &lot.CurrentPrice, &lot.StartTime, &lot.EndTime,
		); err != nil {
			return nil, marketplaces.ErrNoLot.Wrap(err)
		}

		lots = append(lots, lot)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	return lots, nil
}
