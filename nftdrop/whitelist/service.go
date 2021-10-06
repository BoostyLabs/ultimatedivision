// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"

	"github.com/zeebo/errs"
)

// ErrWhitelist indicated that there was an error in service.
var ErrWhitelist = errs.Class("whitelist service error")

// Service is handling whitelist related logic.
//
// architecture: Service
type Service struct {
	whitelist DB
}

// NewService is a constructor for whitelist service.
func NewService(whitelist DB) *Service {
	return &Service{whitelist: whitelist}
}

// Create adds whitelist in the data base.
func (service Service) Create(ctx context.Context, whitelistRequest Request) error {
	whitelist := Whitelist{
		Address: whitelistRequest.Address,
		// TODO: generate password
		Password: []byte{},
	}
	return ErrWhitelist.Wrap(service.whitelist.Create(ctx, whitelist))
}

// Get returns whitelist by address from the data base.
func (service Service) Get(ctx context.Context, address string) (Whitelist, error) {
	whitelist, err := service.whitelist.Get(ctx, address)
	return whitelist, ErrWhitelist.Wrap(err)
}

// List returns all whitelist from the data base.
func (service Service) List(ctx context.Context) ([]Whitelist, error) {
	whitelistRecords, err := service.whitelist.List(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}
