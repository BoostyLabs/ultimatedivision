// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/sync"
	"ultimatedivision/queue/queuehub"
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
	Hub     *queuehub.Hub
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, queue DB, users *users.Service, hub *queuehub.Hub) *Chore {
	return &Chore{
		log: log,
		service: NewService(
			config,
			queue,
			users,
		),
		Loop: sync.NewCycle(config.PlaceRenewalInterval),
		Hub:  hub,
	}
}

// Run starts the chore for re-check the expiration time of the place.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		places, err := chore.service.List(ctx)

		if len(places) >= 2 {
			for k := range places {
				if k%2 != 0 || (places[k] == Place{} && places[k+1] == Place{}) {
					continue
				}

				firstUserID := places[k].UserID
				secondUserID := places[k+1].UserID

				firstUser, err := chore.service.users.Get(ctx, firstUserID)
				if err != nil {
					return ChoreError.Wrap(err)
				}
				secondUser, err := chore.service.users.Get(ctx, secondUserID)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				// TODO: maybe send squad
				err = chore.Hub.SendInvite(firstUserID, secondUser)
				if err != nil {
					return ChoreError.Wrap(err)
				}
				err = chore.Hub.SendInvite(secondUserID, firstUser)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				firstIsInvite, err := chore.Hub.ReadPlay(firstUserID)
				if err != nil {
					return ChoreError.Wrap(err)
				}
				secondIsInvite, err := chore.Hub.ReadPlay(secondUserID)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				firstClient := queuehub.Client{
					UserID: firstUserID,
					Conn:   chore.Hub.Clients[firstUserID],
				}
				secondClient := queuehub.Client{
					UserID: secondUserID,
					Conn:   chore.Hub.Clients[secondUserID],
				}

				if !firstIsInvite || !secondIsInvite {
					message := queuehub.NewMessage(http.StatusOK, "you are still in search!")
					if err = chore.Hub.SendMessage(firstClient, *message); err != nil {
						return ChoreError.Wrap(err)
					}

					if err = chore.Hub.SendMessage(secondClient, *message); err != nil {
						return ChoreError.Wrap(err)
					}
					continue
				}

				if err = chore.Hub.RemoveClient(firstClient); err != nil {
					return ChoreError.Wrap(err)
				}
				if err = chore.Hub.RemoveClient(secondClient); err != nil {
					return ChoreError.Wrap(err)
				}

				if err = chore.service.UpdateStatus(ctx, firstUserID, StatusPlays); err != nil {
					return ChoreError.Wrap(err)
				}
				if err = chore.service.UpdateStatus(ctx, secondUserID, StatusPlays); err != nil {
					return ChoreError.Wrap(err)
				}

				// TODO: add to match and send result

				if err = chore.service.Finish(ctx, firstUserID); err != nil {
					return ChoreError.Wrap(err)
				}
				if err = chore.service.Finish(ctx, secondUserID); err != nil {
					return ChoreError.Wrap(err)
				}
			}
		}
		return ChoreError.Wrap(err)
	})
}
