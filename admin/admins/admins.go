// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"time"

	"github.com/google/uuid"
)

// Admin describes admin entity.
type Admin struct {
	ID           uuid.UUID
	Email        string
	PasswordHash []byte
	CreatedAt    time.Time
}
