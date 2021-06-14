// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"time"

	"github.com/google/uuid"
)

//Admin represent table admins from DB
type Admin struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

//NewAdmin is constructor for Admin
func NewAdmin(email, password string) *Admin {
	return &Admin{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: password,
		CreatedAt:    time.Now(),
	}
}
