// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
)

// ErrClubs indicates that there was an error in the database.
var ErrClubs = errs.Class("clubs repository error")

// ErrSquad indicates that there was an error in the database.
var ErrSquad = errs.Class("squad repository error")

// clubsDB provide access to club DB.
//
// architecture: Database
type clubsDB struct {
	conn *sql.DB
}

// Create creates empty club in the db.
func (clubsDB *clubsDB) Create(ctx context.Context, club clubs.Club) error {
	query := `INSERT INTO clubs(id, owner_id, club_name, created_at)
              VALUES($1,$2,$3,$4)`

	_, err := clubsDB.conn.ExecContext(ctx, query,
		club.ID, club.OwnerID, club.Name, club.CreatedAt)

	return ErrClubs.Wrap(err)
}

func (clubsDB *clubsDB) CreateSquad(ctx context.Context, squad clubs.Squad) error {
	query := `INSERT INTO squads(id, squad_name, club_id, tactic, formation)
              VALUES($1,$2,$3,$4,$5)`

	_, err := clubsDB.conn.ExecContext(ctx, query,
		squad.SquadID, squad.Name, squad.ClubID, squad.Tactic, squad.Formation)

	return ErrClubs.Wrap(err)
}

// Add inserts card to club.
func (clubsDB *clubsDB) AddSquadCard(ctx context.Context, squadCards clubs.SquadCard) error {
	query := `INSERT  INTO squad_cards(id, card_id, card_position, capitan_id)
              VALUES($1,$2,$3,$4)`

	_, err := clubsDB.conn.ExecContext(ctx, query,
		squadCards.SquadID, squadCards.CardID, squadCards.Position, squadCards.CaptainID)

	return ErrSquad.Wrap(err)
}

// DeleteSquadCard deletes card from squad.
func (clubsDB *clubsDB) DeleteSquadCard(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID) error {
	query := `DELETE FROM squad_cards
              WHERE id = $1 AND card_id = $2`

	_, err := clubsDB.conn.ExecContext(ctx, query, squadID, cardID)

	return ErrSquad.Wrap(err)
}

// GetByUserID returns club owned by the user.
func (clubsDB *clubsDB) GetByUserID(ctx context.Context, userID uuid.UUID) (clubs.Club, error) {
	query := `SELECT id, owner_id, club_name, created_at
			  FROM clubs
			  WHERE owner_id = $1`

	row := clubsDB.conn.QueryRowContext(ctx, query, userID)

	var club clubs.Club

	err := row.Scan(&club.ID, &club.OwnerID, &club.Name, &club.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return club, clubs.ErrNoClub.Wrap(err)
		}

		return club, clubs.ErrClubs.Wrap(err)
	}

	return club, nil
}

// GetSquad returns squad from database.
func (clubsDB *clubsDB) GetSquad(ctx context.Context, clubID uuid.UUID) (clubs.Squad, error) {
	query := `SELECT id, squad_name, club_id, tactic, formation 
			  FROM squads
			  WHERE club_id = $1`

	row := clubsDB.conn.QueryRowContext(ctx, query, clubID)

	var squad clubs.Squad

	err := row.Scan(&squad.SquadID, &squad.Name, &squad.ClubID, &squad.Tactic, &squad.Formation)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return squad, clubs.ErrNoSquad.Wrap(err)
		}

		return squad, ErrClubs.Wrap(err)
	}

	return squad, nil
}

// ListSquadCards returns all cards from squad.
func (clubsDB *clubsDB) ListSquadCards(ctx context.Context, squadID uuid.UUID) ([]clubs.SquadCard, error) {
	query := `SELECT id, card_id, card_position, capitan_id 
			  FROM squad_cards
			  WHERE id = $1`

	rows, err := clubsDB.conn.QueryContext(ctx, query, squadID)
	if err != nil {
		return nil, ErrSquad.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrSquad.Wrap(rows.Close()))
	}()

	var players []clubs.SquadCard

	for rows.Next() {
		var player clubs.SquadCard
		err = rows.Scan(&player.SquadID, &player.CardID, &player.Position, &player.CaptainID)
		if err != nil {
			return nil, clubs.ErrNoSquad.Wrap(err)
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrSquad.Wrap(err)
	}

	return players, nil
}

// UpdateTacticFormation updates tactic and formation for squad.
func (clubsDB *clubsDB) UpdateTacticFormation(ctx context.Context, squad clubs.Squad) error {
	query := `UPDATE squads
			  SET tactic = $1, formation = $2
  			  WHERE id = $3`

	_, err := clubsDB.conn.ExecContext(ctx, query, squad.Tactic, squad.Formation, squad.SquadID)

	return ErrSquad.Wrap(err)
}

// UpdateCaptain updates captain in the users team.
func (clubsDB *clubsDB) UpdateCaptain(ctx context.Context, captainID uuid.UUID, squadID uuid.UUID) error {
	query := `UPDATE squad_cards
			  SET capitan_id = $1
			  WHERE id = $2`

	_, err := clubsDB.conn.ExecContext(ctx, query, captainID, squadID)

	return ErrSquad.Wrap(err)
}

// UpdatePosition updates position of card in the squad.
func (clubsDB *clubsDB) UpdatePosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, newPosition clubs.Position) error {
	query := `UPDATE squad_cards
			  SET card_position = $1
			  WHERE card_id = $2 AND id = $3`

	_, err := clubsDB.conn.ExecContext(ctx, query, newPosition, cardID, squadID)

	return ErrSquad.Wrap(err)
}

// GetCaptainID returns id of captain of the users team.
func (clubsDB *clubsDB) GetCaptainID(ctx context.Context, squadID uuid.UUID) (uuid.UUID, error) {
	query := `SELECT capitan_id
			  FROM squad_cards
              WHERE id = $1`

	var id uuid.UUID

	row := clubsDB.conn.QueryRowContext(ctx, query, squadID)

	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.New(), clubs.ErrNoSquad.Wrap(err)
		}

		return uuid.New(), ErrSquad.Wrap(err)
	}

	return id, nil
}
