// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

//ErrNoAdmin represent new class of errors for nil rows in admin table
var ErrNoAdmin = errs.Class("admin does not exist")

type DB interface {
	List(ctx context.Context) ([]Admin, error)
	Get(id uuid.UUID) (Admin, error)
}

//Service struct gives access to database lvl
type Service struct {
	admins DB
	ctx   context.Context
}

//NewService is constructor for Service
func NewService(admins DB, ctx context.Context) *Service {
	return &Service{
		admins: admins,
		ctx:   ctx,
	}
}

//GetAll return all admins from db
func(s *Service) GetAll() ([]Admin,error){
	return s.admins.List(s.ctx)
}

//Get return admin by id
func(s *Service) Get(id uuid.UUID) (Admin,error){
	return s.admins.Get(id)
}



