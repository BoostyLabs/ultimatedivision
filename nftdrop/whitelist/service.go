// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import "github.com/zeebo/errs"

// ErrWhitelist indicated that there was an error in service.
var ErrWhitelist = errs.Class("whitelist service error")

// Service is handling whitelist related logic.
//
// architecture: Service
type Service struct {
	whitelist  DB
}

// NewService is a constructor for cards service.
func NewService(whitelist DB) *Service {
	return &Service{
		whitelist:  whitelist,
	}
}
