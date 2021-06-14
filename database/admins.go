// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/admin/admins"
)

//AdminRepository provide access to DB
type AdminRepository struct{
	conn *sql.DB
}

//List helps get all admins from db
func(a *AdminRepository) List(ctx context.Context) ([]admins.Admin,error){
	rows, err := a.conn.QueryContext(ctx, "SELECT id, email, password, creaed_at FROM admins")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var data []admins.Admin
	for rows.Next() {
		var admin admins.Admin
		err = rows.Scan(&admin.ID, &admin.Email, &admin.PasswordHash,&admin.CreatedAt)
		if err != nil {
			return nil, err
		}

		data = append(data, admin)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil

}

//Get helps get admin from db by id
func(a *AdminRepository) Get(id uuid.UUID) (admins.Admin, error){
	var admin admins.Admin

	row := a.conn.QueryRow("SELECT id, email, password, creaed_at FROM admins WHERE id=$1", id)
	err := row.Scan(&admin.ID, &admin.Email, &admin.PasswordHash, &admin.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return admin, admins.ErrNoAdmin.Wrap(err)
		}

		return admin, err
	}

	return admin, nil
}
