// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package divisions

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrDivisions indicates that there was an error in the service.
var ErrDivisions = errs.Class("divisions service error")

// Service is handling users related logic.
//
// architecture: Service
type Service struct {
	divisions DB
}

// NewService is a constructor for divisions service.
func NewService(divisions DB) *Service {
	return &Service{
		divisions: divisions,
	}
}

// Create creates a division.
func (service *Service) Create(ctx context.Context, name string) error {
	division := Division{
		ID:               uuid.UUID{},
		Name:             name,
		PercentOfPassing: 0,
		CreatedAt:        time.Time{},
	}

	return ErrDivisions.Wrap(service.divisions.Create(ctx, division))
}

// List returns all divisions from DB.
func (service *Service) List(ctx context.Context) ([]Division, error) {
	users, err := service.divisions.List(ctx)
	return users, ErrDivisions.Wrap(err)
}

// Get returns division from DB.
func (service *Service) Get(ctx context.Context, userID uuid.UUID) (Division, error) {
	user, err := service.divisions.Get(ctx, userID)
	return user, ErrDivisions.Wrap(err)
}

// Delete deletes a division.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrDivisions.Wrap(service.divisions.Delete(ctx, id))
}
