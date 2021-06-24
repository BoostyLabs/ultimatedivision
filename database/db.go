// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // using postgres driver
	"github.com/zeebo/errs"

	"ultimatedivision"
	"ultimatedivision/admin/admins"
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
            id               BYTEA PRIMARY KEY        NOT NULL,
            email            VARCHAR                  NOT NULL,
            email_normalized VARCHAR                  NOT NULL,
            password_hash    BYTEA                    NOT NULL,
            nick_name        VARCHAR                  NOT NULL,
            first_name       VARCHAR                  NOT NULL,
            last_name        VARCHAR                  NOT NULL,
            last_login       TIMESTAMP WITH TIME ZONE NOT NULL,
            status           INTEGER                  NOT NULL,
            created_at       TIMESTAMP WITH TIME ZONE NOT NULL
		);
		CREATE TABLE IF NOT EXISTS picture_types (
			id    INTEGER  PRIMARY KEY  NOT NULL,
			name  VARCHAR               NOT NULL
		);
		CREATE TABLE IF NOT EXISTS skin_colors (
			id    INTEGER  PRIMARY KEY  NOT NULL,
			name  VARCHAR               NOT NULL
		);
		CREATE TABLE IF NOT EXISTS hair_styles (
			id    INTEGER  PRIMARY KEY  NOT NULL,
			name  VARCHAR               NOT NULL
		);
		CREATE TABLE IF NOT EXISTS hair_colors (
			id    INTEGER  PRIMARY KEY  NOT NULL,
			name  VARCHAR               NOT NULL
		);
		CREATE TABLE IF NOT EXISTS accessories (
			id    INTEGER  PRIMARY KEY  NOT NULL,
			name  VARCHAR               NOT NULL
		);
		CREATE TABLE IF NOT EXISTS images (
			img   BYTEA    PRIMARY KEY  NOT NULL,
			name  VARCHAR               NOT NULL
		);
		CREATE TABLE IF NOT EXISTS cards (
			id                BYTEA    PRIMARY KEY                   NOT NULL,
			player_name       VARCHAR                                NOT NULL,
			quality           VARCHAR                                NOT NULL,
			picture_type      INTEGER  REFERENCES picture_types(id)  NOT NULL,
			height            DECIMAL                                NOT NULL,
			weight            DECIMAL                                NOT NULL,
			skin_color        INTEGER  REFERENCES skin_colors(id)    NOT NULL,
			hair_style        INTEGER  REFERENCES hair_styles(id)    NOT NULL,
			hair_color        INTEGER  REFERENCES hair_colors(id)    NOT NULL,
			dominant_foot     VARCHAR                                NOT NULL,
			user_id           BYTEA    REFERENCES users(id)          NOT NULL,
			image_id          BYTEA    REFERENCES images(img)        NOT NULL,
			positioning       INTEGER                                NOT NULL,
			composure         INTEGER                                NOT NULL,
			aggression        INTEGER                                NOT NULL,
			vision            INTEGER                                NOT NULL,
			awareness         INTEGER                                NOT NULL,
			crosses           INTEGER                                NOT NULL,
			acceleration      INTEGER                                NOT NULL,
			running_speed     INTEGER                                NOT NULL,
			reaction_speed    INTEGER                                NOT NULL,
			agility           INTEGER                                NOT NULL,
			stamina           INTEGER                                NOT NULL,
			strength          INTEGER                                NOT NULL,
			jumping           INTEGER                                NOT NULL,
			balance           INTEGER                                NOT NULL,
			dribbling         INTEGER                                NOT NULL,
			ball_control      INTEGER                                NOT NULL,
			weak_foot         INTEGER                                NOT NULL,
			skill_moves       INTEGER                                NOT NULL,
			finesse           INTEGER                                NOT NULL,
			curve             INTEGER                                NOT NULL,
			volleys           INTEGER                                NOT NULL,
			short_passing     INTEGER                                NOT NULL,
			long_passing      INTEGER                                NOT NULL,
			forward_pass      INTEGER                                NOT NULL,
			finishing_ability INTEGER                                NOT NULL,
			shot_power        INTEGER                                NOT NULL,
			accuracy          INTEGER                                NOT NULL,
			distance          INTEGER                                NOT NULL,
			penalty           INTEGER                                NOT NULL,
			free_kicks        INTEGER                                NOT NULL,
			corners           INTEGER                                NOT NULL,
			heading_accuracy  INTEGER                                NOT NULL,
			offside_trap      INTEGER                                NOT NULL,
			sliding           INTEGER                                NOT NULL,
			tackles           INTEGER                                NOT NULL,
			ball_focus        INTEGER                                NOT NULL,
			interceptions     INTEGER                                NOT NULL,
			vigilance         INTEGER                                NOT NULL,
			reflexes          INTEGER                                NOT NULL,
			diving            INTEGER                                NOT NULL,
			handling          INTEGER                                NOT NULL,
			sweeping          INTEGER                                NOT NULL,
			throwing          INTEGER                                NOT NULL
		);
		CREATE TABLE IF NOT EXISTS cards_accessories (
			card_id       BYTEA    REFERENCES cards(id)         NOT NULL,
			accessory_id  BYTEA    REFERENCES accessories(img)  NOT NULL,
			PRIMARY KEY(accessory_id, card_id)
		);
        CREATE TABLE IF NOT EXISTS admins (
            id            BYTEA     PRIMARY KEY    NOT NULL,
            email         VARCHAR                  NOT NULL,
            password_hash BYTEA                    NOT NULL,
            created_at    TIMESTAMP WITH TIME ZONE NOT NULL
        );`

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

// adminRepository provided access to accounts db.
func (db *database) Admins() admins.DB {
	return &adminsDB{conn: db.conn}
}

// usersDB provided access to accounts db.
func (db *database) Users() users.DB {
	return &usersDB{conn: db.conn}
}

// cardsDB provided access to accounts db.
func (db *database) Cards() cards.DB {
	return &cardsDB{conn: db.conn}
}
