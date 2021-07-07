// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

var ErrAdminsService = errs.Class("admins service error")

// Service is handling admins related logic.
//
// architecture: Service
type Service struct {
	admins DB
}

// NewService is constructor for Service.
func NewService(admins DB) *Service {
	return &Service{
		admins: admins,
	}
}

// List returns all admins from DB.
func (service *Service) List(ctx context.Context) ([]Admin, error) {
	return service.admins.List(ctx)
}

// Get returns admin from DB.
func (service *Service) Get(ctx context.Context, id uuid.UUID) (Admin, error) {
	return service.admins.Get(ctx, id)
}

// Create insert admin to DB.
func (service *Service) Create(ctx context.Context, admin Admins) error {
	err := admin.EncodePassword()
	if err != nil {
		return ErrAdminsService.Wrap(err)
	}

	return service.admins.Create(ctx, admin.Admin)
}

// Update updates admin from DB.
func (service *Service) Update(ctx context.Context, admin Admins) error {
	err := admin.EncodePassword()
	if err != nil {
		return ErrAdminsService.Wrap(err)
	}

	return service.admins.Update(ctx, admin.Admin)
}

// Delete deletes admin from DB.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return service.admins.Delete(ctx, id)
}
