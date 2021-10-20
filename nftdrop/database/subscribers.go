// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/mail"
	"ultimatedivision/nftdrop/subscribers"
)

// ErrSubscribers indicates that there was an error in the database.
var ErrSubscribers = errs.Class("subscribers repository error")

// subscribersDB provides access to subscribers db.
//
// architecture: Database
type subscribersDB struct {
	conn *sql.DB
}

// List returns all subscribers from the data base.
func (subscribersDB *subscribersDB) List(ctx context.Context) ([]subscribers.Subscriber, error) {
	rows, err := subscribersDB.conn.QueryContext(ctx, "SELECT email, created_at FROM subscribers")
	if err != nil {
		return nil, ErrSubscribers.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrSubscribers.Wrap(rows.Close()))
	}()

	var Subscribers []subscribers.Subscriber
	for rows.Next() {
		var subscriber subscribers.Subscriber
		err := rows.Scan(&subscriber.Email, &subscriber.CreatedAt)
		if err != nil {
			return nil, subscribers.ErrNoSubscriber.Wrap(err)
		}

		Subscribers = append(Subscribers, subscriber)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrSubscribers.Wrap(err)
	}

	return Subscribers, nil
}

// GetByEmail returns subscriber by email from the data base.
func (subscribersDB *subscribersDB) GetByEmail(ctx context.Context, email string) (subscribers.Subscriber, error) {
	var subscriber subscribers.Subscriber
	emailNormalized := mail.NormalizeEmail(email)

	row := subscribersDB.conn.QueryRowContext(ctx, "SELECT email, created_at FROM subscribers WHERE email_normalized=$1", emailNormalized)

	err := row.Scan(&subscriber.Email, &subscriber.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return subscriber, subscribers.ErrNoSubscriber.Wrap(err)
		}

		return subscriber, ErrSubscribers.Wrap(err)
	}

	return subscriber, nil
}

// Create creates a subscriber and writes to the database.
func (subscribersDB *subscribersDB) Create(ctx context.Context, subscriber subscribers.Subscriber) error {
	emailNormalized := mail.NormalizeEmail(subscriber.Email)
	query := `INSERT INTO subscribers(
                  email, 
                  email_normalized, 
                  created_at) 
                  VALUES ($1, $2, $3)`

	_, err := subscribersDB.conn.ExecContext(ctx, query, subscriber.Email, emailNormalized, subscriber.CreatedAt)

	return ErrSubscribers.Wrap(err)
}

// Delete deletes a subscriber in the database.
func (subscribersDB *subscribersDB) Delete(ctx context.Context, email string) error {
	_, err := subscribersDB.conn.QueryContext(ctx, "DELETE FROM subscribers WHERE email=$1", email)

	return ErrSubscribers.Wrap(err)
}
