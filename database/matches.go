// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/matches"
)

// ensures that matchesDB implements matches.DB.
var _ matches.DB = (*matchesDB)(nil)

// ErrMatches indicates that there was an error in the database.
var ErrMatches = errs.Class("matches repository error")

// matchesDB provide access to matches DB.
//
// architecture: Database
type matchesDB struct {
	conn *sql.DB
}

// Create inserts match in the database.
func (matchesDB *matchesDB) Create(ctx context.Context, match matches.Match) error {
	query := `INSERT INTO matches(id, user1_id, user2_id, score)
              VALUES($1,$2,$3,$4)`

	_, err := matchesDB.conn.ExecContext(ctx, query, match.ID, match.User1ID, match.User2ID, match.Score)

	return ErrMatches.Wrap(err)
}

// Get returns match from the database.
func (matchesDB *matchesDB) Get(ctx context.Context, id uuid.UUID) (matches.Match, error) {
	query := `SELECT id, user1_id, user2_id, score
              FROM matches
              WHERE id = $1`

	var match matches.Match

	row := matchesDB.conn.QueryRowContext(ctx, query, id)

	err := row.Scan(&match.ID, &match.User1ID, &match.User2ID, &match.Score)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return match, matches.ErrNoMatch.Wrap(err)
		}

		return match, ErrMatches.Wrap(err)
	}

	return match, ErrMatches.Wrap(err)
}

// Update updates score in the match in the database.
func (matchesDB *matchesDB) Update(ctx context.Context, matchID uuid.UUID, score string) error {
	query := `UPDATE matches
              SET score = $1
              WHERE id = $2`

	_, err := matchesDB.conn.ExecContext(ctx, query, score, matchID)

	return ErrMatches.Wrap(err)
}

// ListMatches returns all matches from the database.
func (matchesDB *matchesDB) ListMatches(ctx context.Context) ([]matches.Match, error) {
	query := `SELECT id, user1_id, user2_id, score
             FROM matches`

	rows, err := matchesDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrMatches.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allMatches []matches.Match

	for rows.Next() {
		var match matches.Match
		err = rows.Scan(&match.ID, &match.User1ID, &match.User2ID, &match.Score)
		if err != nil {
			return nil, ErrMatches.Wrap(err)
		}

		allMatches = append(allMatches, match)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMatches.Wrap(err)
	}

	return allMatches, ErrMatches.Wrap(err)
}

// Delete deletes match from the database.
func (matchesDB *matchesDB) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM matches
              WHERE id = $1`

	_, err := matchesDB.conn.ExecContext(ctx, query, id)

	return ErrMatches.Wrap(err)
}

// AddGoal adds new goal in the match.
func (matchesDB *matchesDB) AddGoal(ctx context.Context, matchGoal matches.MatchGoals) error {
	query := `INSERT INTO match_goals(id, match_id, user_id, card_id, minute)
              VALUES($1,$2,$3,$4,$5)`

	_, err := matchesDB.conn.ExecContext(ctx, query,
		matchGoal.ID, matchGoal.MatchID, matchGoal.UserID, matchGoal.CardID, matchGoal.Minute)

	return ErrMatches.Wrap(err)
}

// ListMatchGoals returns all goals from the match from the database.
func (matchesDB *matchesDB) ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]matches.MatchGoals, error) {
	query := `SELECT id, match_id, user_id, card_id, minute
              FROM match_goals
              WHERE match_id = $1`

	rows, err := matchesDB.conn.QueryContext(ctx, query, matchID)
	if err != nil {
		return nil, ErrMatches.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var goals []matches.MatchGoals

	for rows.Next() {
		var goal matches.MatchGoals
		err = rows.Scan(&goal.ID, &goal.MatchID, &goal.UserID, &goal.CardID, &goal.Minute)
		if err != nil {
			return nil, ErrMatches.Wrap(err)
		}

		goals = append(goals, goal)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMatches.Wrap(err)
	}

	return goals, ErrMatches.Wrap(err)
}
