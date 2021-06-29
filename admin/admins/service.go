// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package admins

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// salt is made up of random bits.
const salt string = "^AF*(SF(A*^->%s<-s)#)("

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
func (service *Service) Create(ctx context.Context, admin Admin) error {
	return service.admins.Create(ctx, admin)
}

// GeneratePasswordHash generates a hash for the password.
func (service *Service) GeneratePasswordHash(password string) ([]byte, error) {
	if val, err := bcrypt.GenerateFromPassword(addSalt(password), bcrypt.DefaultCost); err != nil {
		return nil, fmt.Errorf("can`t to create password hash, err:%s", err)
	} else {
		return val, nil
	}
}

// addSalt add salt string to the password.
func addSalt(pwd string) []byte {
	data := []byte(fmt.Sprintf(salt, pwd))
	bt := md5.Sum(data)
	bytes := make([]byte, 16)
	copy(bytes, bt[0:])
	return []byte(hex.EncodeToString(bytes))
}
