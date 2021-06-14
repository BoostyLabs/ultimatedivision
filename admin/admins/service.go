// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoAdmin indicates that user does not exist.
var ErrNoAdmin = errs.Class("admin does not exist")

// DB exposes access to admin db.
type DB interface {
	List(ctx context.Context) ([]Admin, error)
	Get(ctx context.Context,id uuid.UUID) (Admin, error)
}

//Service struct gives access to database lvl
type Service struct {
	admins DB
}

// NewService is constructor for Service
func NewService(admins DB) *Service {
	return &Service{
		admins: admins,
	}
}

// GetAll return all admins from db
func(service *Service) GetAll(ctx context.Context) ([]Admin,error){
	return service.admins.List(ctx)
}

// Get return admin by id
func(service *Service) Get(ctx context.Context,id uuid.UUID) (Admin,error){
	return service.admins.Get(ctx,id)
}



