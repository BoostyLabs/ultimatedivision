// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/nftdrop/emails"
)

// ErrEmails indicates that there was an error in the database.
var ErrEmails = errs.Class("emails repository error")

// emailsDB provides access to emails db.
//
// architecture: Database
type emailsDB struct {
	conn *sql.DB
}

// List returns all emails from the data base.
func (emailsDB *emailsDB) List(ctx context.Context) ([]emails.Email, error) {
	rows, err := emailsDB.conn.QueryContext(ctx, "SELECT id, email, created_at FROM emails")
	if err != nil {
		return nil, ErrEmails.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrEmails.Wrap(rows.Close()))
	}()

	var data []emails.Email
	for rows.Next() {
		var email emails.Email
		err := rows.Scan(&email.ID, &email.Name, &email.CreatedAt)
		if err != nil {
			return nil, emails.ErrNoEmail.Wrap(err)
		}

		data = append(data, email)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrEmails.Wrap(err)
	}

	return data, nil
}

// Get returns email by id from the data base.
func (emailsDB *emailsDB) Get(ctx context.Context, id uuid.UUID) (emails.Email, error) {
	var email emails.Email

	row := emailsDB.conn.QueryRowContext(ctx, "SELECT id, email, created_at FROM emails WHERE id=$1", id)

	err := row.Scan(&email.ID, &email.Name, &email.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return email, emails.ErrNoEmail.Wrap(err)
		}

		return email, ErrEmails.Wrap(err)
	}

	return email, nil
}

// GetByEmail returns email by email from the data base.
func (emailsDB *emailsDB) GetByEmail(ctx context.Context, name string) (emails.Email, error) {
	var email emails.Email
	emailNormalized := normalizeEmail(name)

	row := emailsDB.conn.QueryRowContext(ctx, "SELECT id, email, created_at FROM emails WHERE email_normalized=$1", emailNormalized)

	err := row.Scan(&email.ID, &email.Name, &email.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return email, emails.ErrNoEmail.Wrap(err)
		}

		return email, ErrEmails.Wrap(err)
	}

	return email, nil
}

// Create creates a email and writes to the database.
func (emailsDB *emailsDB) Create(ctx context.Context, email emails.Email) error {
	emailNormalized := normalizeEmail(email.Name)
	query := `INSERT INTO emails(
                  id, 
                  email, 
                  email_normalized, 
                  created_at) 
                  VALUES ($1, $2, $3, $4)`

	_, err := emailsDB.conn.QueryContext(ctx, query, email.ID, email.Name, emailNormalized, email.CreatedAt)
	if err != nil {
		return ErrEmails.Wrap(err)
	}

	return nil
}

// Delete deletes a email in the database.
func (emailsDB *emailsDB) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := emailsDB.conn.QueryContext(ctx, "DELETE FROM emails WHERE id=$1", id)
	if err != nil {
		return ErrEmails.Wrap(err)
	}

	return nil
}

// Update updates a email in the database.
func (emailsDB *emailsDB) Update(ctx context.Context, email string, id uuid.UUID) error {
	_, err := emailsDB.conn.QueryContext(ctx, "UPDATE emails SET email=$1 WHERE id=$2", email, id)
	if err != nil {
		return ErrEmails.Wrap(err)
	}

	return nil
}
