// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package games

import (
	"context"

	"github.com/zeebo/errs"
)

// ErrGames indicates that there was an error in the service.
var ErrGames = errs.Class("game service error")

// Service is handling matches related logic.
//
// architecture: Service
type Service struct {
	games  DB
	config Config
}

// NewService is a constructor for games service.
func NewService(matches DB, config Config) *Service {
	return &Service{
		games:  matches,
		config: config,
	}
}

// Create adds card in DB.
func (service *Service) Create(ctx context.Context, gameInfo Game) error {
	return ErrGames.Wrap(service.games.Create(ctx, gameInfo))
}

// List returns all games from DB.
func (service *Service) List(ctx context.Context) ([]Game, error) {
	games, err := service.games.List(ctx)
	return games, ErrGames.Wrap(err)
}

// Update updates game info in the database by match id.
func (service *Service) Update(ctx context.Context, gameInfo Game) error {
	return ErrGames.Wrap(service.games.Update(ctx, gameInfo))
}
