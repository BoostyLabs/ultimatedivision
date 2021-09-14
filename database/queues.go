// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/pagination"
	"ultimatedivision/queues"
)

// ensures that queuesDB implements queues.DB.
var _ queues.DB = (*queuesDB)(nil)

// ErrQueue indicates that there was an error in the database.
var ErrQueue = errs.Class("queues repository error")

// queuesDB provides access to queues database.
//
// architecture: Database
type queuesDB struct {
	conn *sql.DB
}

// Create adds queue in the database.
func (queuesDB *queuesDB) Create(ctx context.Context, queue queues.Queue) error {
	query :=
		`INSERT INTO
			queues(user_id, status) 
		VALUES 
			($1, $2)`

	_, err := queuesDB.conn.ExecContext(ctx, query, queue.UserID, queue.Status)
	return ErrQueue.Wrap(err)
}

// Get returns queue by id of user from the database.
func (queuesDB *queuesDB) Get(ctx context.Context, id uuid.UUID) (queues.Queue, error) {
	queue := queues.Queue{}
	query :=
		`SELECT 
			user_id, status
		FROM 
			queues
		WHERE 
			user_id = $1`

	err := queuesDB.conn.QueryRowContext(ctx, query, id).Scan(&queue.UserID, &queue.Status)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return queue, queues.ErrNoQueue.Wrap(err)
	case err != nil:
		return queue, ErrQueue.Wrap(err)
	default:
		return queue, nil
	}
}

// ListPaginated returns queues in page from the database.
func (queuesDB *queuesDB) ListPaginated(ctx context.Context, cursor pagination.Cursor) (queues.Page, error) {
	var queuesListPage queues.Page
	offset := (cursor.Page - 1) * cursor.Limit
	query :=
		`SELECT 
			user_id, status 
		FROM 
			queues 
		LIMIT 
			$1
		OFFSET 
			$2`

	rows, err := queuesDB.conn.QueryContext(ctx, query, cursor.Limit, offset)
	if err != nil {
		return queuesListPage, ErrQueue.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	data := []queues.Queue{}
	for rows.Next() {
		queue := queues.Queue{}
		if err = rows.Scan(&queue.UserID, &queue.Status); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return queuesListPage, queues.ErrNoQueue.Wrap(err)
			}
			return queuesListPage, ErrQueue.Wrap(err)
		}
		data = append(data, queue)
	}
	if err = rows.Err(); err != nil {
		return queuesListPage, ErrQueue.Wrap(err)
	}

	queuesListPage, err = queuesDB.listPaginated(ctx, cursor, data)
	return queuesListPage, ErrQueue.Wrap(err)
}

// listPaginated returns paginated list of queues.
func (queuesDB *queuesDB) listPaginated(ctx context.Context, cursor pagination.Cursor, queuesList []queues.Queue) (queues.Page, error) {
	var queuesListPage queues.Page
	offset := (cursor.Page - 1) * cursor.Limit

	totalCount, err := queuesDB.totalCount(ctx)
	if err != nil {
		return queuesListPage, ErrQueue.Wrap(err)
	}

	pageCount := totalCount / cursor.Limit
	if totalCount%cursor.Limit != 0 {
		pageCount++
	}

	queuesListPage = queues.Page{
		Queues: queuesList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalCount,
		},
	}
	return queuesListPage, nil
}

// totalCount counts all the queues in the table.
func (queuesDB *queuesDB) totalCount(ctx context.Context) (int, error) {
	var count int
	query :=
		`SELECT 
			COUNT(*) 
		FROM 
			queues`

	err := queuesDB.conn.QueryRowContext(ctx, query).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, queues.ErrNoQueue.Wrap(err)
	}
	return count, ErrQueue.Wrap(err)
}

// UpdateStatus updates status of queue in the database.
func (queuesDB *queuesDB) UpdateStatus(ctx context.Context, id uuid.UUID, status queues.Status) error {
	query :=
		`UPDATE
			queues 
		SET 
			status = $1 
		WHERE 
			user_id = $2`

	_, err := queuesDB.conn.ExecContext(ctx, query, status, id)
	return ErrQueue.Wrap(err)
}

// Delete deletes record queue in the database.
func (queuesDB *queuesDB) Delete(ctx context.Context, id uuid.UUID) error {
	query :=
		`DELETE FROM
			queues
		WHERE 
			user_id = $1`

	_, err := queuesDB.conn.ExecContext(ctx, query, id)
	return ErrQueue.Wrap(err)
}
