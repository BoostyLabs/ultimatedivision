// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/users"
)

// usersDB provides access to users db.
//
// architecture: Database
type usersDB struct {
	conn *sql.DB
}

// List returns all users from the data base.
func (usersDB *usersDB) List(ctx context.Context) ([]users.User, error) {
	rows, err := usersDB.conn.QueryContext(ctx, "SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		data = append(data, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// Get returns user by id from the data base.
func (usersDB *usersDB) Get(ctx context.Context, id uuid.UUID) (users.User, error) {
	var user users.User

	row, err := usersDB.conn.QueryContext(ctx, "SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}

	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, err
	}

	return user, nil
}

// GetByEmail returns user by email from the data base.
func (usersDB *usersDB) GetByEmail(ctx context.Context, email string) (users.User, error) {
	var user users.User

	row, err := usersDB.conn.QueryContext(ctx, "SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users WHERE email=$1", email)
	if err != nil {
		return user, err
	}

	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, err
	}

	return user, nil
}

// Create, creates a user and writes to the database.
func (usersDB *usersDB) Create(ctx context.Context, user users.User) (string, error) {
	var email string
	row, err := usersDB.conn.QueryContext(ctx,
		"INSERT INTO  users(email, password, nick_name, first_name, last_name, last_login, status, creaed_at) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING email", user.Email, user.PasswordHash,
		user.NickName, user.FirstName, user.LastName, user.LastLogin, user.Status, user.CreatedAt)
	if err != nil {
		return email, err
	}

	err = row.Scan(&email)
	if err != nil {
		return email, err
	}

	return email, nil
}
