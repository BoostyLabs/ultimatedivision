// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package emails

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrEmails indicates that there was an error in the service.
var ErrEmails = errs.Class("emails service error")

// Service is handling emails related logic.
//
// architecture: Service
type Service struct {
	emails DB
}

// NewService is a constructor for emails service.
func NewService(emails DB) *Service {
	return &Service{
		emails: emails,
	}
}

// Get returns email from DB.
func (service *Service) Get(ctx context.Context, emailID uuid.UUID) (Email, error) {
	email, err := service.emails.Get(ctx, emailID)
	return email, ErrEmails.Wrap(err)
}

// GetByEmail returns email by email from DB.
func (service *Service) GetByEmail(ctx context.Context, name string) (Email, error) {
	email, err := service.emails.GetByEmail(ctx, name)
	return email, ErrEmails.Wrap(err)
}

// List returns all emails from DB.
func (service *Service) List(ctx context.Context) ([]Email, error) {
	emails, err := service.emails.List(ctx)
	return emails, ErrEmails.Wrap(err)
}

// Create creates a email.
func (service *Service) Create(ctx context.Context, name string) error {
	email := Email{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	return ErrEmails.Wrap(service.emails.Create(ctx, email))
}

// Delete deletes a email.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrEmails.Wrap(service.emails.Delete(ctx, id))
}

// Update updates a emails name.
func (service *Service) Update(ctx context.Context, name string, id uuid.UUID) error {
	return ErrEmails.Wrap(service.emails.Update(ctx, name, id))
}
