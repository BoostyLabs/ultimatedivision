// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package emails

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoEmail indicated that user does not exist.
var ErrNoEmail = errs.Class("email does not exist")

// DB exposes access to emails db.
//
// architecture: DB
type DB interface {
	// Create creates a email and writes to the database.
	Create(ctx context.Context, email Email) error
	// List returns all emails from the data base.
	List(ctx context.Context) ([]Email, error)
	// Get returns email by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Email, error)
	// Update updates a status in the database.
	Update(ctx context.Context, email string, id uuid.UUID) error
	// Delete deletes a email in the database.
	Delete(ctx context.Context, id uuid.UUID) error
	// GetByEmail returns email by email from the data base.
	GetByEmail(ctx context.Context, name string) (Email, error)
}

// Email describes email entity.
type Email struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// CreateEmailFields for crete email.
type CreateEmailFields struct {
	Email string `json:"email"`
}
