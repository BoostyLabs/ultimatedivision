// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/pagination"
)

// ErrNoMatch indicated that match does not exist.
var ErrNoMatch = errs.Class("match does not exist")

// DB is exposing access to matches db.
//
// architecture: DB
type DB interface {
	// Create creates match in the database.
	Create(ctx context.Context, match Match) error
	// Get returns match from the database.
	Get(ctx context.Context, id uuid.UUID) (Match, error)
	// GetGoals counts goals for user's squad in the match.
	GetGoals(ctx context.Context, matchID uuid.UUID, userID uuid.UUID) (int, error)
	// ListMatches returns page of matches from the database.
	ListMatches(ctx context.Context, cursor pagination.Cursor) (Page, error)
	// Delete deletes match from the database.
	Delete(ctx context.Context, id uuid.UUID) error
	// AddGoal adds new goal in the match.
	AddGoal(ctx context.Context, matchGoal MatchGoals) error
	// ListMatchGoals returns all goals from the match from the database.
	ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]MatchGoals, error)
}

// Config defines configuration for marketplace.
type Config struct {
	pagination.Cursor `json:"cursor"`
}

// Match describes match entity.
type Match struct {
	ID      uuid.UUID `json:"Id"`
	User1ID uuid.UUID `json:"user1Id"`
	User2ID uuid.UUID `json:"user2Id"`
}

// MatchGoals defines goals scored by clubs.
type MatchGoals struct {
	ID      uuid.UUID `json:"id"`
	MatchID uuid.UUID `json:"matchId"`
	UserID  uuid.UUID `json:"userId"`
	CardID  uuid.UUID `json:"cardId"`
	Minute  int       `json:"minute"`
}

// Page holds match page entity which is used to show listed page of matches.
type Page struct {
	Matches []Match         `json:"matches"`
	Page    pagination.Page `json:"page"`
}
