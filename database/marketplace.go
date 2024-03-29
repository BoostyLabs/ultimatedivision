// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/zeebo/errs"

	"ultimatedivision/marketplace"
	"ultimatedivision/pkg/pagination"
)

// ensures that marketplaceDB implements marketplace.DB.
var _ marketplace.DB = (*marketplaceDB)(nil)

// ErrMarketplace indicates that there was an error in the database.
var ErrMarketplace = errs.Class("marketplace repository error")

// marketplaceDB provides access to marketplace db.
//
// architecture: Database
type marketplaceDB struct {
	conn *sql.DB
}

const (
	allFieldsOfLot = `card_id, type, user_id, shopper_id, status, start_price, max_price, current_price, start_time, end_time, period`
)

// CreateLot creates lot in the db.
func (marketplaceDB *marketplaceDB) CreateLot(ctx context.Context, lot marketplace.Lot) error {
	query :=
		`INSERT INTO 
			lots(` + allFieldsOfLot + ` )
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := marketplaceDB.conn.ExecContext(ctx, query,
		lot.CardID, lot.Type, lot.UserID, lot.ShopperID, lot.Status,
		lot.StartPrice.Bytes(), lot.MaxPrice.Bytes(), lot.CurrentPrice.Bytes(), lot.StartTime, lot.EndTime, lot.Period)

	return ErrMarketplace.Wrap(err)
}

// GetLotByID returns lot by id from the data base.
func (marketplaceDB *marketplaceDB) GetLotByID(ctx context.Context, id uuid.UUID) (marketplace.Lot, error) {
	var (
		startPrice   []byte
		maxPrice     []byte
		currentPrice []byte
		lot          marketplace.Lot
	)

	query :=
		`SELECT 
			lots.card_id, lots.type, lots.user_id, shopper_id, lots.status, start_price, max_price, current_price, start_time, end_time, period,
			cards.id, player_name, quality, height, weight, dominant_foot, is_tattoo, cards.status, cards.type,
			cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
			stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
			forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
			tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing
		FROM 
			lots
		LEFT JOIN 
			cards ON lots.card_id = cards.id
		WHERE 
			lots.card_id = $1`

	err := marketplaceDB.conn.QueryRowContext(ctx, query, id).Scan(
		&lot.CardID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status, &startPrice, &maxPrice, &currentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
		&lot.Card.ID, &lot.Card.PlayerName, &lot.Card.Quality, &lot.Card.Height, &lot.Card.Weight, &lot.Card.DominantFoot, &lot.Card.IsTattoo, &lot.Card.Status, &lot.Card.Type, &lot.Card.UserID, &lot.Card.Tactics, &lot.Card.Positioning,
		&lot.Card.Composure, &lot.Card.Aggression, &lot.Card.Vision, &lot.Card.Awareness, &lot.Card.Crosses, &lot.Card.Physique, &lot.Card.Acceleration, &lot.Card.RunningSpeed,
		&lot.Card.ReactionSpeed, &lot.Card.Agility, &lot.Card.Stamina, &lot.Card.Strength, &lot.Card.Jumping, &lot.Card.Balance, &lot.Card.Technique, &lot.Card.Dribbling,
		&lot.Card.BallControl, &lot.Card.WeakFoot, &lot.Card.SkillMoves, &lot.Card.Finesse, &lot.Card.Curve, &lot.Card.Volleys, &lot.Card.ShortPassing, &lot.Card.LongPassing,
		&lot.Card.ForwardPass, &lot.Card.Offence, &lot.Card.FinishingAbility, &lot.Card.ShotPower, &lot.Card.Accuracy, &lot.Card.Distance, &lot.Card.Penalty,
		&lot.Card.FreeKicks, &lot.Card.Corners, &lot.Card.HeadingAccuracy, &lot.Card.Defence, &lot.Card.OffsideTrap, &lot.Card.Sliding, &lot.Card.Tackles, &lot.Card.BallFocus,
		&lot.Card.Interceptions, &lot.Card.Vigilance, &lot.Card.Goalkeeping, &lot.Card.Reflexes, &lot.Card.Diving, &lot.Card.Handling, &lot.Card.Sweeping, &lot.Card.Throwing,
	)
	lot.StartPrice.SetBytes(startPrice)
	lot.MaxPrice.SetBytes(maxPrice)
	lot.CurrentPrice.SetBytes(currentPrice)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return lot, marketplace.ErrNoLot.Wrap(err)
	case err != nil:
		return lot, ErrMarketplace.Wrap(err)
	default:
		return lot, nil
	}
}

// GetLotEndTimeByID returns lot end time by id from data base.
func (marketplaceDB *marketplaceDB) GetLotEndTimeByID(ctx context.Context, id uuid.UUID) (time.Time, error) {
	var endTime time.Time

	query := `SELECT end_time FROM lots WHERE card_id = $1`

	err := marketplaceDB.conn.QueryRowContext(ctx, query, id).Scan(&endTime)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return endTime, marketplace.ErrNoLot.Wrap(err)
	case err != nil:
		return endTime, ErrMarketplace.Wrap(err)
	default:
		return endTime, nil
	}
}

// GetCurrentPriceByCardID returns current price by card id from the data base.
func (marketplaceDB *marketplaceDB) GetCurrentPriceByCardID(ctx context.Context, cardID uuid.UUID) (big.Int, error) {
	var (
		currentPrice    []byte
		currentPriceInt big.Int
	)

	query := `SELECT current_price FROM lots WHERE card_id = $1`

	err := marketplaceDB.conn.QueryRowContext(ctx, query, cardID).Scan(&currentPrice)
	currentPriceInt.SetBytes(currentPrice)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return currentPriceInt, marketplace.ErrNoLot.Wrap(err)
	case err != nil:
		return currentPriceInt, ErrMarketplace.Wrap(err)
	default:
		return currentPriceInt, nil
	}
}

// ListActiveLots returns active lots from the data base.
func (marketplaceDB *marketplaceDB) ListActiveLots(ctx context.Context, cursor pagination.Cursor) (marketplace.Page, error) {
	var (
		startPrice   []byte
		maxPrice     []byte
		currentPrice []byte
		lotsListPage marketplace.Page
	)

	offset := (cursor.Page - 1) * cursor.Limit
	query :=
		`SELECT 
			lots.card_id, lots.type, lots.user_id, shopper_id, lots.status, start_price, max_price, current_price, start_time, end_time, period,
			cards.id, player_name, quality, height, weight, dominant_foot, is_tattoo, cards.status, cards.type,
			cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
			stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
			forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
			tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing
		FROM 
			lots
		LEFT JOIN 
			cards ON lots.card_id = cards.id
		WHERE
			lots.status = $1
		LIMIT 
			$2
		OFFSET 
			$3`

	rows, err := marketplaceDB.conn.QueryContext(ctx, query, marketplace.StatusActive, cursor.Limit, offset)
	if err != nil {
		return lotsListPage, ErrMarketplace.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	lots := []marketplace.Lot{}
	for rows.Next() {
		lot := marketplace.Lot{}
		if err = rows.Scan(
			&lot.CardID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status, &startPrice, &maxPrice, &currentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
			&lot.Card.ID, &lot.Card.PlayerName, &lot.Card.Quality, &lot.Card.Height, &lot.Card.Weight,
			&lot.Card.DominantFoot, &lot.Card.IsTattoo, &lot.Card.Status, &lot.Card.Type, &lot.Card.UserID, &lot.Card.Tactics, &lot.Card.Positioning,
			&lot.Card.Composure, &lot.Card.Aggression, &lot.Card.Vision, &lot.Card.Awareness, &lot.Card.Crosses, &lot.Card.Physique, &lot.Card.Acceleration, &lot.Card.RunningSpeed,
			&lot.Card.ReactionSpeed, &lot.Card.Agility, &lot.Card.Stamina, &lot.Card.Strength, &lot.Card.Jumping, &lot.Card.Balance, &lot.Card.Technique, &lot.Card.Dribbling,
			&lot.Card.BallControl, &lot.Card.WeakFoot, &lot.Card.SkillMoves, &lot.Card.Finesse, &lot.Card.Curve, &lot.Card.Volleys, &lot.Card.ShortPassing, &lot.Card.LongPassing,
			&lot.Card.ForwardPass, &lot.Card.Offence, &lot.Card.FinishingAbility, &lot.Card.ShotPower, &lot.Card.Accuracy, &lot.Card.Distance, &lot.Card.Penalty,
			&lot.Card.FreeKicks, &lot.Card.Corners, &lot.Card.HeadingAccuracy, &lot.Card.Defence, &lot.Card.OffsideTrap, &lot.Card.Sliding, &lot.Card.Tackles, &lot.Card.BallFocus,
			&lot.Card.Interceptions, &lot.Card.Vigilance, &lot.Card.Goalkeeping, &lot.Card.Reflexes, &lot.Card.Diving, &lot.Card.Handling, &lot.Card.Sweeping, &lot.Card.Throwing,
		); err != nil {
			return lotsListPage, ErrMarketplace.Wrap(err)
		}
		lot.StartPrice.SetBytes(startPrice)
		lot.MaxPrice.SetBytes(maxPrice)
		lot.CurrentPrice.SetBytes(currentPrice)

		lots = append(lots, lot)
	}

	totalActiveCount, err := marketplaceDB.totalActiveCount(ctx)
	if err != nil {
		return lotsListPage, ErrMarketplace.Wrap(err)
	}

	lotsListPage, err = marketplaceDB.listPaginated(ctx, cursor, lots, totalActiveCount)
	return lotsListPage, ErrMarketplace.Wrap(err)
}

// ListActiveLotsByCardID returns active lots from the data base by card id.
func (marketplaceDB *marketplaceDB) ListActiveLotsByCardID(ctx context.Context, cardIds []uuid.UUID, cursor pagination.Cursor) (marketplace.Page, error) {
	var (
		startPrice   []byte
		maxPrice     []byte
		currentPrice []byte
		lotsListPage marketplace.Page
	)

	offset := (cursor.Page - 1) * cursor.Limit
	query :=
		`SELECT 
			lots.card_id, lots.type, lots.user_id, shopper_id, lots.status, start_price, max_price, current_price, start_time, end_time, period,
			cards.id, player_name, quality, height, weight, dominant_foot, is_tattoo, cards.status, cards.type,
			cards.user_id, tactics, positioning, composure, aggression, vision, awareness, crosses, physique, acceleration, running_speed, reaction_speed, agility,
			stamina, strength, jumping, balance, technique, dribbling, ball_control, weak_foot, skill_moves, finesse, curve, volleys, short_passing, long_passing,
			forward_pass, offense, finishing_ability, shot_power, accuracy, distance, penalty, free_kicks, corners, heading_accuracy, defence, offside_trap, sliding,
			tackles, ball_focus, interceptions, vigilance, goalkeeping, reflexes, diving, handling, sweeping, throwing
		FROM 
			lots
		LEFT JOIN 
			cards ON lots.card_id = cards.id
		WHERE
			lots.status = $1 AND lots.card_id = ANY($2)
		LIMIT 
			$3 
		OFFSET 
			$4`

	rows, err := marketplaceDB.conn.QueryContext(ctx, query, marketplace.StatusActive, pq.Array(cardIds), cursor.Limit, offset)
	if err != nil {
		return lotsListPage, ErrMarketplace.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var lots []marketplace.Lot
	for rows.Next() {
		lot := marketplace.Lot{}
		if err = rows.Scan(
			&lot.CardID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status, &startPrice, &maxPrice, &currentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
			&lot.Card.ID, &lot.Card.PlayerName, &lot.Card.Quality, &lot.Card.Height, &lot.Card.Weight, &lot.Card.DominantFoot, &lot.Card.IsTattoo, &lot.Card.Status, &lot.Card.Type, &lot.Card.UserID, &lot.Card.Tactics, &lot.Card.Positioning,
			&lot.Card.Composure, &lot.Card.Aggression, &lot.Card.Vision, &lot.Card.Awareness, &lot.Card.Crosses, &lot.Card.Physique, &lot.Card.Acceleration, &lot.Card.RunningSpeed,
			&lot.Card.ReactionSpeed, &lot.Card.Agility, &lot.Card.Stamina, &lot.Card.Strength, &lot.Card.Jumping, &lot.Card.Balance, &lot.Card.Technique, &lot.Card.Dribbling,
			&lot.Card.BallControl, &lot.Card.WeakFoot, &lot.Card.SkillMoves, &lot.Card.Finesse, &lot.Card.Curve, &lot.Card.Volleys, &lot.Card.ShortPassing, &lot.Card.LongPassing,
			&lot.Card.ForwardPass, &lot.Card.Offence, &lot.Card.FinishingAbility, &lot.Card.ShotPower, &lot.Card.Accuracy, &lot.Card.Distance, &lot.Card.Penalty,
			&lot.Card.FreeKicks, &lot.Card.Corners, &lot.Card.HeadingAccuracy, &lot.Card.Defence, &lot.Card.OffsideTrap, &lot.Card.Sliding, &lot.Card.Tackles, &lot.Card.BallFocus,
			&lot.Card.Interceptions, &lot.Card.Vigilance, &lot.Card.Goalkeeping, &lot.Card.Reflexes, &lot.Card.Diving, &lot.Card.Handling, &lot.Card.Sweeping, &lot.Card.Throwing,
		); err != nil {
			return lotsListPage, ErrMarketplace.Wrap(err)
		}
		lot.StartPrice.SetBytes(startPrice)
		lot.MaxPrice.SetBytes(maxPrice)
		lot.CurrentPrice.SetBytes(currentPrice)

		lots = append(lots, lot)
	}

	totalActiveCount, err := marketplaceDB.totalActiveCountWithFilters(ctx, cardIds)
	if err != nil {
		return lotsListPage, ErrCard.Wrap(err)
	}

	lotsListPage, err = marketplaceDB.listPaginated(ctx, cursor, lots, totalActiveCount)
	return lotsListPage, ErrMarketplace.Wrap(err)
}

// listPaginated returns paginated list of lots.
func (marketplaceDB *marketplaceDB) listPaginated(ctx context.Context, cursor pagination.Cursor, lotsList []marketplace.Lot, totalActiveCount int) (marketplace.Page, error) {
	var lotsListPage marketplace.Page
	offset := (cursor.Page - 1) * cursor.Limit
	pageCount := totalActiveCount / cursor.Limit
	if totalActiveCount%cursor.Limit != 0 {
		pageCount++
	}

	lotsListPage = marketplace.Page{
		Lots: lotsList,
		Page: pagination.Page{
			Offset:      offset,
			Limit:       cursor.Limit,
			CurrentPage: cursor.Page,
			PageCount:   pageCount,
			TotalCount:  totalActiveCount,
		},
	}

	return lotsListPage, nil
}

// totalActiveCount counts active lots in the table.
func (marketplaceDB *marketplaceDB) totalActiveCount(ctx context.Context) (int, error) {
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM lots WHERE lots.status = $1`)
	err := marketplaceDB.conn.QueryRowContext(ctx, query, marketplace.StatusActive).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, marketplace.ErrNoLot.Wrap(err)
	}
	return count, ErrMarketplace.Wrap(err)
}

// totalActiveCountWithFilters counts active lots with filtes in the table.
func (marketplaceDB *marketplaceDB) totalActiveCountWithFilters(ctx context.Context, itemIds []uuid.UUID) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM lots WHERE lots.status = $1 AND card_id = ANY($2)")
	err := marketplaceDB.conn.QueryRowContext(ctx, query, marketplace.StatusActive, pq.Array(itemIds)).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, marketplace.ErrNoLot.Wrap(err)
	}
	return count, ErrMarketplace.Wrap(err)
}

// ListExpiredLot returns lots where end time lower than or equal to time now UTC from the data base.
func (marketplaceDB *marketplaceDB) ListExpiredLot(ctx context.Context) ([]marketplace.Lot, error) {
	var (
		startPrice   []byte
		maxPrice     []byte
		currentPrice []byte
		lots         []marketplace.Lot
	)

	query :=
		`SELECT 
			` + allFieldsOfLot + ` 
		FROM 
			lots
		WHERE
			status = $1
		AND
			end_time <= $2
		`

	rows, err := marketplaceDB.conn.QueryContext(ctx, query, marketplace.StatusActive, time.Now().UTC())
	if err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	for rows.Next() {
		lot := marketplace.Lot{}
		if err = rows.Scan(
			&lot.CardID, &lot.Type, &lot.UserID, &lot.ShopperID, &lot.Status,
			&startPrice, &maxPrice, &currentPrice, &lot.StartTime, &lot.EndTime, &lot.Period,
		); err != nil {
			return nil, ErrMarketplace.Wrap(err)
		}
		lot.StartPrice.SetBytes(startPrice)
		lot.MaxPrice.SetBytes(maxPrice)
		lot.CurrentPrice.SetBytes(currentPrice)

		lots = append(lots, lot)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrMarketplace.Wrap(err)
	}

	return lots, nil
}

// UpdateShopperIDLot updates shopper id of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateShopperIDLot(ctx context.Context, id, shopperID uuid.UUID) error {
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET shopper_id = $1 WHERE card_id = $2", shopperID, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}

// UpdateStatusLot updates status of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateStatusLot(ctx context.Context, id uuid.UUID, status marketplace.Status) error {
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET status = $1 WHERE card_id = $2", status, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}

// UpdateCurrentPriceLot updates current price of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice big.Int) error {
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET current_price = $1 WHERE card_id = $2", currentPrice.Bytes(), id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}

// UpdateEndTimeLot updates end time of lot in the database.
func (marketplaceDB *marketplaceDB) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	result, err := marketplaceDB.conn.ExecContext(ctx, "UPDATE lots SET end_time = $1 WHERE card_id = $2", endTime, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowsNum, err := result.RowsAffected()
	if rowsNum == 0 {
		return marketplace.ErrNoLot.New("lot does not exist")
	}

	return ErrMarketplace.Wrap(err)
}

// Delete deletes lot in the database.
func (marketplaceDB *marketplaceDB) Delete(ctx context.Context, cardID uuid.UUID) error {
	query := "DELETE FROM lots WHERE card_id = $1"

	result, err := marketplaceDB.conn.ExecContext(ctx, query, cardID)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return marketplace.ErrNoLot.New("")
	}

	return ErrMarketplace.Wrap(err)
}
