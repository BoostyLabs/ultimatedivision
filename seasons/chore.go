// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/sync"
)

var (
	// ChoreError represents season chore error type.
	ChoreError = errs.Class("expiration season chore error")
)

// Chore requests access token for counties api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log     logger.Logger
	service *Service
	Loop    *sync.Cycle
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, season DB) *Chore {
	return &Chore{
		log: log,
		service: NewService(
			season,
			config,
		),
		Loop: sync.NewCycle(config.SeasonsRenewalInterval),
	}
}

// Run starts the chore for re-check the expiration time of the season.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		err := chore.service.Create(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		return ChoreError.Wrap(err)
	})
}
