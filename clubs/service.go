// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import "github.com/zeebo/errs"

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// Service is handling users related logic.
//
// architecture: Service
type Service struct {
	clubs DB
}

// NewService is a constructor for clubs service.
func NewService(clubs DB) *Service {
	return &Service{
		clubs: clubs,
	}
}