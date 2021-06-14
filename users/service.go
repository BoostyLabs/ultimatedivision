// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"github.com/google/uuid"
)

//Service struct gives access to DB interface methods
type Service struct {
	users DB
	ctx   context.Context
}

//returns service with db interface
func NewService(users DB, ctx context.Context) *Service {
	return &Service{
		users: users,
		ctx:   ctx,
	}
}

// GetUser return user from DB
func (service *Service) GetUser(userID uuid.UUID) (User, error) {
	return service.users.Get(userID)
}

// GetUser return user by email from DB
func (service *Service) GetUserByEmail(email string) (User, error) {
	return service.users.GetByEmail(email)
}

// GetUser return all users from DB
func (service *Service) GetListOfUsers() ([]User, error) {
	return service.users.List(service.ctx)
}
