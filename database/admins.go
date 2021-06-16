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

// adminRepository provide access to admin DB.
//
// Architecture: Database
type adminRepository struct{
	conn *sql.DB
}

// List returns all admins from db.
func(adminRepository *adminRepository) List(ctx context.Context) ([]admins.Admin,error){
	rows, err := adminRepository.conn.QueryContext(ctx, "SELECT id, email, password_hash, creaed_at FROM admins")
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

// Get returns admin from db by id.
func(adminRepository *adminRepository) Get(ctx context.Context,id uuid.UUID) (admins.Admin, error){
	var admin admins.Admin

	row,err := adminRepository.conn.QueryContext(ctx,"SELECT id, email, password_hash, creaed_at FROM admins WHERE id=$1", id)
	if err != nil {
		return admins.Admin{}, err
	}

	err = row.Scan(&admin.ID, &admin.Email, &admin.PasswordHash, &admin.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return admin, admins.ErrNoAdmin.Wrap(err)
		}

		return admin, err
	}
	return admin, nil
}

// Create inserts admin to DB.
func(adminRepository *adminRepository) Create(ctx context.Context,admin admins.Admin) error{
	_,err := adminRepository.conn.QueryContext(ctx,
		`INSERT INTO admins(id,email,password_hash,created_at)
		VALUES($1,$2,$3,$4)`,admin.ID,admin.Email,admin.PasswordHash,admin.CreatedAt)
	return err
}
