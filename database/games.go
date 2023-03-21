// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/gameplay/gameengine"
)

// ensures that gameDB implements game.DB.
var _ gameengine.DB = (*gameengineDB)(nil)

// ErrGames indicates that there was an error in the database.
var ErrGames = errs.Class("games repository error")

// gameengineDB provide access to games DB.
//
// architecture: Database
type gameengineDB struct {
	conn *sql.DB
}

// Create creates game in db.
func (gameengineDB *gameengineDB) Create(ctx context.Context, game gameengine.Game) error {
	tx, err := gameengineDB.conn.BeginTx(ctx, nil)
	if err != nil {
		return ErrGames.Wrap(err)
	}
	query := `INSERT INTO games(match_id, game_info)
              VALUES($1,$2)`

	gameInfo, err := json.Marshal(game.GameInfo)
	if err != nil {
		return ErrGames.Wrap(err)
	}

	_, err = gameengineDB.conn.ExecContext(ctx, query, game.MatchID, string(gameInfo))

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
func (gameengineDB *gameengineDB) List(ctx context.Context) ([]gameengine.Game, error) {
	query := `SELECT match_id, game_info
              FROM games`

	rows, err := gameengineDB.conn.QueryContext(ctx, query)
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
func (gameengineDB *gameengineDB) Get(ctx context.Context, matchID uuid.UUID) (gameengine.Game, error) {
	query := `SELECT match_id, game_info
              FROM games
              WHERE match_id = $1`

	row := gameengineDB.conn.QueryRowContext(ctx, query, matchID)

	var game gameengine.Game
	var gameData string

	err := row.Scan(&game.MatchID, &gameData)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return game, gameengine.ErrNoGames.Wrap(err)
		}

		return game, ErrGames.Wrap(err)
	}

	err = json.Unmarshal([]byte(gameData), &game.GameInfo)
	if err != nil {
		return game, ErrGames.Wrap(err)
	}

	return game, nil
}

// Update updates game info in the database by match id.
func (gameengineDB *gameengineDB) Update(ctx context.Context, game gameengine.Game) error {
	gameInfo, err := json.Marshal(game.GameInfo)
	if err != nil {
		return ErrGames.Wrap(err)
	}
	result, err := gameengineDB.conn.ExecContext(ctx, "UPDATE games SET game_info=$1 WHERE match_id=$2", string(gameInfo), game.MatchID)
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
func (gameengineDB *gameengineDB) Delete(ctx context.Context, matchID uuid.UUID) error {
	query := `DELETE FROM games
              WHERE match_id = $1`

	result, err := gameengineDB.conn.ExecContext(ctx, query, matchID)
	if err != nil {
		return ErrGames.Wrap(err)
	}
	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return gameengine.ErrNoGames.New("invalid query")
	}

	return ErrGames.Wrap(err)
}
