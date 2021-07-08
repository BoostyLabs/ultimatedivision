// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"
)

var ErrAdmins = errs.Class("admins service error")

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

// encodePassword is method to encode password.
func (service *Service) encodePassword(admin Admin) error {
	hash, err := bcrypt.GenerateFromPassword(admin.PasswordHash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.PasswordHash = hash
	return nil
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
func (service *Service) Create(ctx context.Context, email string, password []byte) error {
	admin := Admin{
		ID: uuid.New(),
		Email: email,
		PasswordHash: password,
		CreatedAt: time.Now(),
	}

	err := service.encodePassword(admin)
	if err != nil {
		return ErrAdmins.Wrap(err)
	}

	return service.admins.Create(ctx, admin)
}

// Update updates admin from DB.
func (service *Service) Update(ctx context.Context, id uuid.UUID, newPassword []byte) error {
	admin, err := service.Get(ctx, id)
	if err != nil {
		return ErrAdmins.Wrap(err)
	}

	admin.PasswordHash = newPassword

	err = service.encodePassword(admin)
	if err != nil {
		return ErrAdmins.Wrap(err)
	}

	return service.admins.Update(ctx, admin)
}
