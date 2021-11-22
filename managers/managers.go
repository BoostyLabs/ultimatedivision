// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package managers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoManager indicated that manager does not exist.
var ErrNoManager = errs.Class("manager does not exist")

// DB is exposing access to managers db.
//
// architecture: DB
type DB interface {
	// Create creates manager in the database.
	Create(ctx context.Context, managers Manager) error
	// List returns all managers from the database.
	List(ctx context.Context) ([]Manager, error)
	// ListByUserID returns clubs which user manages by id from the database.
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]Manager, error)
	// Delete deletes manager from the database.
	Delete(ctx context.Context, userID, clubID uuid.UUID) error
}

// Manager defines managers entity.
type Manager struct {
	UserID  uuid.UUID `json:"userId"`
	ClubID  uuid.UUID `json:"clubId"`
	EndedAt time.Time `json:"endedAt"`
}

// Config defines configuration for managers chore.
type Config struct {
	RenewalTime time.Duration `json:"renewalTime"`
}
