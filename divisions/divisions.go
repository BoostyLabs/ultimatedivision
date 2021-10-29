// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package divisions

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoDivisions indicated that divisions does not exist.
var ErrNoDivisions = errs.Class("divisions does not exist")

// DB exposes access to divisions db.
//
// architecture: DB
type DB interface {
	// Create creates a division and writes to the database.
	Create(ctx context.Context, division Division) error
	// List returns all divisions from the data base.
	List(ctx context.Context) ([]Division, error)
	// Get returns division by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Division, error)
	// Delete deletes a division in the database.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Position defines the list of possible divisions positions.
type Position int

// Division describes divisions entity.
type Division struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	PercentOfPassing int       `json:"percent_of_passing"`
	CreatedAt        time.Time `json:"createdAt"`
}
