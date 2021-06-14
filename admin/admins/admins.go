// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"time"

	"github.com/google/uuid"
)

// Admin describes user entity.
type Admin struct {
	ID           uuid.UUID
	Email        string
	PasswordHash []byte
	CreatedAt    time.Time
}

// NewAdmin is constructor for Admin.
func NewAdmin(email string,  password []byte) *Admin {
	return &Admin{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: password,
		CreatedAt:    time.Now(),
	}
}
