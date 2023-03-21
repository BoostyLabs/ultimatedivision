// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"ultimatedivision/gameplay/gameengine"
)

// ensures that gameDB implements game.DB.
var _ gameengine.DB = (*gamesDB)(nil)

// ErrGames indicates that there was an error in the database.
var ErrGames = errs.Class("games repository error")

// gamesDB provide access to games DB.
//
// architecture: Database
type gamesDB struct {
	conn *sql.DB
}

// Create creates game in db.
func (gamesDB *gamesDB) Create(ctx context.Context, game gameengine.Game) error {
	tx, err := gamesDB.conn.BeginTx(ctx, nil)
	if err != nil {
		return ErrGames.Wrap(err)
	}
	query := `INSERT INTO games(match_id, game_info)
              VALUES($1,$2)`

	_, err = gamesDB.conn.ExecContext(ctx, query, game.MatchID, game.GameInfo)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return ErrGames.Wrap(err)
		}
		return ErrGames.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return ErrGames.Wrap(err)
	}

	return ErrGames.Wrap(err)
}

// List returns all games.
func (gamesDB *gamesDB) List(ctx context.Context) ([]gameengine.Game, error) {
	query := `SELECT match_id, game_info
              FROM games`

	rows, err := gamesDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrGames.Wrap(err)
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var allGames []gameengine.Game

	for rows.Next() {
		var game gameengine.Game

		err = rows.Scan(&game.MatchID, &game.GameInfo)
		if err != nil {
			return nil, ErrGames.Wrap(err)
		}

		allGames = append(allGames, game)
	}

	return allGames, nil
}

// Get returns game by match id.
func (gamesDB *gamesDB) Get(ctx context.Context, matchID uuid.UUID) (gameengine.Game, error) {
	query := `SELECT match_id, game_info
              FROM games
              WHERE match_id = $1`

	row := gamesDB.conn.QueryRowContext(ctx, query, matchID)
	fmt.Println("row ---> ", row)
	var game gameengine.Game

	err := row.Scan(&game.MatchID, &game.GameInfo)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return game, gameengine.ErrNoGames.Wrap(err)
		}

		return game, ErrGames.Wrap(err)
	}

	return game, nil
}

// Update updates game info in the database by match id.
func (gamesDB *gamesDB) Update(ctx context.Context, gameInfo gameengine.Game) error {
	result, err := gamesDB.conn.ExecContext(ctx, "UPDATE games SET game_info=$1 WHERE match_id=$2", gameInfo.MatchID, gameInfo.GameInfo)
	if err != nil {
		return ErrGames.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return gameengine.ErrNoGames.New("")
	}

	return ErrGames.Wrap(err)
}

// Delete deletes game by match id from db.
func (gamesDB *gamesDB) Delete(ctx context.Context, matchID uuid.UUID) error {
	query := `DELETE FROM games
              WHERE match_id = $1`

	result, err := gamesDB.conn.ExecContext(ctx, query, matchID)
	if err != nil {
		return ErrGames.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return gameengine.ErrNoGames.New("invalid query")
	}

	return ErrGames.Wrap(err)
}
