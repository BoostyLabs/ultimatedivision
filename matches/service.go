// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/pagination"
)

// ErrMatches indicates that there was an error in the service.
var ErrMatches = errs.Class("matches service error")

// Service is handling matches related logic.
//
// architecture: Service
type Service struct {
	matches DB
	config  Config
}

// NewService is a constructor for matches service.
func NewService(clubs DB, config Config) *Service {
	return &Service{
		matches: clubs,
		config:  config,
	}
}

// Create creates new match.
func (service *Service) Create(ctx context.Context, user1ID, user2ID uuid.UUID) error {
	newMatch := Match{
		ID:      uuid.New(),
		User1ID: user1ID,
		User2ID: user2ID,
		Score:   "",
	}

	return ErrMatches.Wrap(service.matches.Create(ctx, newMatch))
}

// Get returns match by id.
func (service *Service) Get(ctx context.Context, matchID uuid.UUID) (Match, error) {
	match, err := service.matches.Get(ctx, matchID)

	return match, ErrMatches.Wrap(err)
}

// Update updates score in the match.
func (service *Service) Update(ctx context.Context, matchID uuid.UUID, score string) error {
	return ErrMatches.Wrap(service.matches.Update(ctx, matchID, score))
}

// List returns all matches.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	allMatches, err := service.matches.ListMatches(ctx, cursor)

	return allMatches, ErrMatches.Wrap(err)
}

// Delete deletes match.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrMatches.Wrap(service.matches.Delete(ctx, id))
}

// AddGoal adds goal in the match.
func (service *Service) AddGoal(ctx context.Context, matchGoal MatchGoals) error {
	return ErrMatches.Wrap(service.matches.AddGoal(ctx, matchGoal))
}

// ListMatchGoals returns all goals scored in the match.
func (service *Service) ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]MatchGoals, error) {
	matchGoals, err := service.matches.ListMatchGoals(ctx, matchID)

	return matchGoals, ErrMatches.Wrap(err)
}
