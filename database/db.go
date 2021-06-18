// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision"
	"ultimatedivision/cards"
	"ultimatedivision/users"
)

// ensures that database implements ultimatedivision.DB.
var _ ultimatedivision.DB = (*database)(nil)

var (
	// Error is the default ultimatedivision error class.
	Error = errs.Class("ultimatedivision db error")
)

// database combines access to different database tables with a record
// of the db driver, db implementation, and db source URL.
//
// architecture: Master Database
type database struct {
	conn *sql.DB
}

// New returns ultimatedivision.DB postgresql implementation.
func New(databaseURL string) (ultimatedivision.DB, error) {
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	return &database{conn: conn}, nil
}

// CreateSchema create schema for all tables and databases.
func (db *database) CreateSchema(ctx context.Context) (err error) {
	createTableQuery :=
		`CREATE TABLE IF NOT EXISTS users (
            id         BYTEA     PRIMARY KEY 	NOT NULL,
            email      VARCHAR                  NOT NULL,
            password   BYTEA                    NOT NULL,
            nick_name  VARCHAR                  NOT NULL,
            first_name VARCHAR                  NOT NULL,
            last_name  VARCHAR                  NOT NULL,
            last_login TIMESTAMP WITH TIME ZONE NOT NULL,
            status     INTEGER                  NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL
		);
		CREATE TABLE IF NOT EXISTS cards (
            id                BYTEA    PRIMARY KEY           NOT NULL,
            player_name       VARCHAR                        NOT NULL,
            quality           VARCHAR                        NOT NULL,
            picture_type      VARCHAR                        NOT NULL,
            height            DECIMAL                        NOT NULL,
            weight            DECIMAL                        NOT NULL,
            skin_color        INTEGER                        NOT NULL,
			hair_style        INTEGER                        NOT NULL,
			hair_color        INTEGER                        NOT NULL,
			accessories       INTEGER[]                      NOT NULL,
			dominant_foot     VARCHAR                        NOT NULL,
			user_id           BYTEA  REFERENCES users(id)    NOT NULL,
			positioning       INTEGER                        NOT NULL,
			composure         INTEGER                        NOT NULL,
			aggression        INTEGER                        NOT NULL,
			vision            INTEGER                        NOT NULL,
			awareness         INTEGER                        NOT NULL,
			crosses           INTEGER                        NOT NULL,
			acceleration      INTEGER                        NOT NULL,
			running_speed     INTEGER                        NOT NULL,
			reaction_speed    INTEGER                        NOT NULL,
			agility           INTEGER                        NOT NULL,
			stamina           INTEGER                        NOT NULL,
			strength          INTEGER                        NOT NULL,
			jumping           INTEGER                        NOT NULL,
			balance           INTEGER                        NOT NULL,
			dribbling         INTEGER                        NOT NULL,
			ball_control      INTEGER                        NOT NULL,
			weak_foot         INTEGER                        NOT NULL,
			skill_moves       INTEGER                        NOT NULL,
			finesse           INTEGER                        NOT NULL,
			curve             INTEGER                        NOT NULL,
			volleys           INTEGER                        NOT NULL,
			short_passing     INTEGER                        NOT NULL,
			long_passing      INTEGER                        NOT NULL,
			forward_pass      INTEGER                        NOT NULL,
			finishing_ability INTEGER                        NOT NULL,
			shot_power        INTEGER                        NOT NULL,
			accuracy          INTEGER                        NOT NULL,
			distance          INTEGER                        NOT NULL,
			penalty           INTEGER                        NOT NULL,
			free_kicks        INTEGER                        NOT NULL,
			corners           INTEGER                        NOT NULL,
			heading_accuracy  INTEGER                        NOT NULL,
			offside_trap      INTEGER                        NOT NULL,
			sliding           INTEGER                        NOT NULL,
			tackles           INTEGER                        NOT NULL,
			ball_focus        INTEGER                        NOT NULL,
			interceptions     INTEGER                        NOT NULL,
			vigilance         INTEGER                        NOT NULL,
			reflexes          INTEGER                        NOT NULL,
			diving            INTEGER                        NOT NULL,
			handling          INTEGER                        NOT NULL,
			sweeping          INTEGER                        NOT NULL,
			throwing          INTEGER                        NOT NULL
		);
		`

	_, err = db.conn.ExecContext(ctx, createTableQuery)
	if err != nil {
		return Error.Wrap(err)
	}

	return nil
}

// Close closes underlying db connection.
func (db *database) Close() error {
	return Error.Wrap(db.conn.Close())
}

// usersDB provided access to accounts db.
func (db *database) Users() users.DB {
	return &usersDB{conn: db.conn}
}

// cardsDB provided access to accounts db.
func (db *database) Cards() cards.DB {
	return &cardsDB{conn: db.conn}
}
