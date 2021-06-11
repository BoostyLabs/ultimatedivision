// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"database/sql"
	"ultimatedivision/users"
)

type UsersRepository struct {
	db *database
}

func (r *UsersRepository) GetAll() ([]users.User, error) {
	rows, err := r.db.conn.Query("SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		data = append(data, user)
	}

	return data, nil
}

func (r *UsersRepository) GetById(id string) (*users.User, error) {
	var user users.User

	row := r.db.conn.QueryRow("SELECT id, email, password, nick_name, first_name, last_name, last_login, status, creaed_at FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.NickName, &user.FirstName, &user.LastName, &user.LastLogin, &user.Status, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &user, nil
}
