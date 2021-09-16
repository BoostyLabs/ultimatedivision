// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"fmt"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/sync"
	"ultimatedivision/users"
)

var (
	// ChoreError represents place chore error type.
	ChoreError = errs.Class("expiration place chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log     logger.Logger
	service *Service
	Loop    *sync.Cycle
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, queue DB, users *users.Service) *Chore {
	return &Chore{
		log: log,
		service: NewService(
			config,
			queue,
			users,
		),
		Loop: sync.NewCycle(config.PlaceRenewalInterval),
	}
}

// Run starts the chore for re-check the expiration time of the place.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		places, err := chore.service.List(ctx)

		if len(places) > 2 {
			for k := range places {
				if (places[k] != Place{} && places[k+1] != Place{}) {
					fmt.Println("lalal 11111")
				}
				fmt.Println("lalal 22222")
			}
		}
		fmt.Println("lalal 33333")
		return ChoreError.Wrap(err)
	})
}
