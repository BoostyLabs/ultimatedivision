// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package managers

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/sync"
)

var (
	// ChoreError represents contract chore error type.
	ChoreError = errs.Class("expiration contract chore error")
)

// Chore requests access managers service, re-requests it after manager's contract expiration time.
//
// architecture: Chore
type Chore struct {
	service *Service
	Loop    *sync.Cycle
}

// NewChore instantiates Chore.
func NewChore(config Config, service *Service) *Chore {
	return &Chore{
		service: service,
		Loop:    sync.NewCycle(config.RenewalTime),
	}
}

// Run starts the chore for re-check the expiration time of the manager.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		allManagers, err := chore.service.List(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, manager := range allManagers {
			if manager.EndedAt.Before(time.Now().UTC()) {
				err = chore.service.Delete(ctx, manager.UserID, manager.ClubID)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			}
		}

		return ChoreError.Wrap(err)
	})
}
