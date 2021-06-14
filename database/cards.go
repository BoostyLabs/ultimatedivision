// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
)

// cardsDB provides access to cards db.
//
// architecture: Database
type cardsDB struct {
	conn *sql.DB
}

const (
	allFields = `id, playerName, quality, pictureType, height, weight, skinColor, hairStyle, hairColor, accessories, dominantFoot,
				positioning, composure, aggression, vision, awareness, crosses,
				acceleration, runningSpeed, reactionSpeed, agility, stamina, strength, jumping, balance,
				dribbling, ballControl, weakFoot, skillMoves, finesse, curve, volleys, shortPassing, longPassing, forwardPass,
				finishingAbility, shotPower, accuracy, distance, penalty, freeKicks, corners, headingAccuracy,
				offsideTrap, sliding, tackles, ballFocus, interceptions, vigilance,
				reflexes, diving, handling, sweeping, throwing
				`
)

// Create add card in the data base.
func (cardsDB *cardsDB) Create(ctx context.Context, card *cards.Card) error {
	query :=
		`
		INSERT INTO 
			cards(` + allFields + `) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27,
				$28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54)
		`
	_, err := cardsDB.conn.ExecContext(ctx, query,
		card.Id, card.PlayerName, card.Quality, card.PictureType, card.Height, card.Weight, card.SkinColor, card.HairStyle, card.HairColor, card.Accessories, card.DominantFoot,
		card.Positioning, card.Composure, card.Aggression, card.Vision, card.Awareness, card.Crosses,
		card.Acceleration, card.RunningSpeed, card.ReactionSpeed, card.Agility, card.Stamina, card.Strength, card.Jumping, card.Balance,
		card.Dribbling, card.BallControl, card.WeakFoot, card.SkillMoves, card.Finesse, card.Curve, card.Volleys, card.ShortPassing, card.LongPassing, card.ForwardPass,
		card.FinishingAbility, card.ShotPower, card.Accuracy, card.Distance, card.Penalty, card.FreeKicks, card.Corners, card.HeadingAccuracy,
		card.OffsideTrap, card.Sliding, card.Tackles, card.BallFocus, card.Interceptions, card.Vigilance,
		card.Reflexes, card.Diving, card.Handling, card.Sweeping, card.Throwing,
	)

	if err != nil {
		return err
	}

	return nil
}

// Get returns card by id from the data base.
func (cardsDB *cardsDB) Get(ctx context.Context, id uuid.UUID) ([]*cards.Card, error) {
	query := "SELECT " + allFields + " FROM cards WHERE id=$1"
	rows, err := cardsDB.conn.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	data := []*cards.Card{}
	for rows.Next() {
		card := &cards.Card{}
		if err = rows.Scan(
			&card.Id, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle, &card.HairColor, &card.Accessories, &card.DominantFoot,
			&card.Positioning, &card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses,
			&card.Acceleration, &card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance,
			&card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing, &card.ForwardPass,
			&card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy,
			&card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions, &card.Vigilance,
			&card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
		); err != nil {
			return nil, err
		}
		data = append(data, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// List returns all cards from the data base.
func (cardsDB *cardsDB) List(ctx context.Context) ([]*cards.Card, error) {
	query := "SELECT " + allFields + " FROM cards"
	rows, err := cardsDB.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	data := []*cards.Card{}
	for rows.Next() {
		card := &cards.Card{}
		if err = rows.Scan(
			&card.Id, &card.PlayerName, &card.Quality, &card.PictureType, &card.Height, &card.Weight, &card.SkinColor, &card.HairStyle, &card.HairColor, &card.Accessories, &card.DominantFoot,
			&card.Positioning, &card.Composure, &card.Aggression, &card.Vision, &card.Awareness, &card.Crosses,
			&card.Acceleration, &card.RunningSpeed, &card.ReactionSpeed, &card.Agility, &card.Stamina, &card.Strength, &card.Jumping, &card.Balance,
			&card.Dribbling, &card.BallControl, &card.WeakFoot, &card.SkillMoves, &card.Finesse, &card.Curve, &card.Volleys, &card.ShortPassing, &card.LongPassing, &card.ForwardPass,
			&card.FinishingAbility, &card.ShotPower, &card.Accuracy, &card.Distance, &card.Penalty, &card.FreeKicks, &card.Corners, &card.HeadingAccuracy,
			&card.OffsideTrap, &card.Sliding, &card.Tackles, &card.BallFocus, &card.Interceptions, &card.Vigilance,
			&card.Reflexes, &card.Diving, &card.Handling, &card.Sweeping, &card.Throwing,
		); err != nil {
			return nil, err
		}
		data = append(data, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
