// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package managers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrManagers indicates that there was an error in the service.
var ErrManagers = errs.Class("managers service error")

// Service is handling managers related logic.
//
// architecture: Service
type Service struct {
	managers DB
}

// NewService is constructor for Service.
func NewService(managers DB) *Service {
	return &Service{
		managers: managers,
	}
}

// Create creates manager.
func (service *Service) Create(ctx context.Context, endedAt time.Time, userID, clubID uuid.UUID) error {
	newManager := Manager{
		UserID:  userID,
		ClubID:  clubID,
		EndedAt: endedAt.UTC(),
	}

	return ErrManagers.Wrap(service.managers.Create(ctx, newManager))
}

// List returns all managers.
func (service *Service) List(ctx context.Context) ([]Manager, error) {
	allManagers, err := service.managers.List(ctx)
	return allManagers, ErrManagers.Wrap(err)
}

// ListByUserID returns clubs which user manages by id from the database.
func (service *Service) ListByUserID(ctx context.Context, userID uuid.UUID) ([]Manager, error) {
	allManagers, err := service.managers.ListByUserID(ctx, userID)
	return allManagers, ErrManagers.Wrap(err)
}

// Delete deletes manager.
func (service *Service) Delete(ctx context.Context, userID, clubID uuid.UUID) error {
	return ErrManagers.Wrap(service.managers.Delete(ctx, userID, clubID))
}
