// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/users"
)

// ErrQueue indicated that there was an error in service.
var ErrQueue = errs.Class("queue service error")

// Service is handling queues related logic.
//
// architecture: Service
type Service struct {
	config Config
	queues DB
	users  *users.Service
	clubs  *clubs.Service
}

// NewService is a constructor for queues service.
func NewService(config Config, queues DB, users *users.Service, clubs *clubs.Service) *Service {
	return &Service{
		config: config,
		queues: queues,
		users:  users,
		clubs:  clubs,
	}
}

// Create adds client's queue in database.
func (service *Service) Create(ctx context.Context, client Client) error {
	if _, err := service.users.Get(ctx, client.UserID); err != nil {
		return ErrQueue.Wrap(err)
	}

	squad, err := service.clubs.GetSquad(ctx, client.SquadID)
	if err != nil {
		return ErrQueue.Wrap(err)
	}

	_, err = service.clubs.Get(ctx, squad.ClubID)
	if err != nil {
		return ErrQueue.Wrap(err)
	}

	// TODO: add division CardID to client.

	err = service.queues.Delete(client.UserID)
	if ErrNoClient.Has(err) || err == nil {
		service.queues.Create(client)
		return nil
	}

	return ErrQueue.Wrap(err)
}

// Get returns client from database.
func (service *Service) Get(userID uuid.UUID) (Client, error) {
	queue, err := service.queues.Get(userID)
	return queue, ErrQueue.Wrap(err)
}

// List returns clients from database.
func (service *Service) List() []Client {
	return service.queues.List()
}

// ListNotPlayingUsers returns clients who don't play game from database.
func (service *Service) ListNotPlayingUsers() []Client {
	return service.queues.ListNotPlayingUsers()
}

// UpdateIsPlaying updates is playing status of client in database.
func (service *Service) UpdateIsPlaying(userID uuid.UUID, isPlaying bool) error {
	return service.queues.UpdateIsPlaying(userID, isPlaying)
}

// Finish finishes client's queue in database.
func (service *Service) Finish(userID uuid.UUID) error {
	return service.queues.Delete(userID)
}
