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

// Repo structure for connect to postgres data base
type repo struct {
	conn *sql.DB
}

// get all users from the data base
func (r *repo) List(ctx context.Context) ([]users.User, error) {
	rows, err := r.conn.QueryContext(ctx, "SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
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

//get user by id from the data base
func (r *repo) Get(id uuid.UUID) (users.User, error) {
	var user users.User

	row := r.conn.QueryRow("SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, err
	}

	return user, nil
}

//get user by email from the data base
func (r *repo) GetByEmail(email string) (users.User, error) {
	var user users.User

	row := r.conn.QueryRow("SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, users.ErrNoUser.Wrap(err)
		}

		return user, err
	}

	return user, nil
}
