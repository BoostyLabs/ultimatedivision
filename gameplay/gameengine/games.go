// Copyright (C) 2021 - 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoGames indicates that game does not exist.
var ErrNoGames = errs.Class("game does not exist")

// DB is exposing access to games db.
//
// architecture: DB.
type DB interface {
	// Create creates game in db.
	Create(ctx context.Context, gameInfo Game) error
	// List returns all games.
	List(ctx context.Context) ([]Game, error)
	// Get returns game by match id.
	Get(ctx context.Context, matchID uuid.UUID) (Game, error)
	// Update updates game info in the database by match id.
	Update(ctx context.Context, gameInfo Game) error
	// Delete deletes game by match id from db.
	Delete(ctx context.Context, gameID uuid.UUID) error
}

// Game defines game data.
type Game struct {
	MatchID  uuid.UUID
	GameInfo []Card
}

// Card defines Card with moves for game.
type Card struct {
	CardID   uuid.UUID `json:"cardID"`
	Position int       `json:"position"`
}
