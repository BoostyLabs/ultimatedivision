// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/pkg/pagination"
)

// ErrMatches indicates that there was an error in the service.
var ErrMatches = errs.Class("matches service error")

// Service is handling matches related logic.
//
// architecture: Service
type Service struct {
	matches DB
	config  Config
	clubs   *clubs.Service
	cards   *cards.Service
}

// NewService is a constructor for matches service.
func NewService(matches DB, config Config, clubs *clubs.Service, cards *cards.Service) *Service {
	return &Service{
		matches: matches,
		config:  config,
		clubs:   clubs,
		cards:   cards,
	}
}

// MinNumberOfMatches defines minimal number of matches to participate in weekly competition.
const MinNumberOfMatches = 3

// maxNumberOfMatches defines maximal number of matches in weekly competition.
const maxNumberOfMatches = 30

// Play initiates match between users, calls methods to generate result.
func (service *Service) Play(ctx context.Context, match Match, squadCards1 []clubs.SquadCard, squadCards2 []clubs.SquadCard) error {
	return nil
}

// Create creates new match.
func (service *Service) Create(ctx context.Context, squad1ID uuid.UUID, squad2ID uuid.UUID, user1ID, user2ID uuid.UUID, seasonID int) (uuid.UUID, error) {
	_, err := service.clubs.ListSquadCardIDs(ctx, squad1ID)
	if err != nil {
		return uuid.Nil, ErrMatches.Wrap(err)
	}

	_, err = service.clubs.ListSquadCardIDs(ctx, squad2ID)
	if err != nil {
		return uuid.Nil, ErrMatches.Wrap(err)
	}

	newMatch := Match{
		ID:       uuid.New(),
		User1ID:  user1ID,
		Squad1ID: squad1ID,
		User2ID:  user2ID,
		Squad2ID: squad2ID,
		SeasonID: seasonID,
	}

	if err = service.matches.Create(ctx, newMatch); err != nil {
		return uuid.Nil, ErrMatches.Wrap(err)
	}

	return newMatch.ID, ErrMatches.Wrap(err)
}

// ConvertPositionsForGameplay converts positions for cards from 0-10 view to coordinates.
func (service *Service) ConvertPositionsForGameplay(ctx context.Context, squadCards []clubs.GetSquadCard) ([]SquadCardWithPosition, error) {
	var PositionToCoordinates = map[clubs.Position]PositionInTheField{
		clubs.GK:   service.config.Positions.GK.PositionInTheField,
		clubs.LB:   service.config.Positions.LB.PositionInTheField,
		clubs.LCD:  service.config.Positions.LCB.PositionInTheField,
		clubs.CCD:  service.config.Positions.CCB.PositionInTheField,
		clubs.RCD:  service.config.Positions.RCB.PositionInTheField,
		clubs.RB:   service.config.Positions.RB.PositionInTheField,
		clubs.LCDM: service.config.Positions.LCDM.PositionInTheField,
		clubs.RCDM: service.config.Positions.RCDM.PositionInTheField,
		clubs.CCDM: service.config.Positions.CCDM.PositionInTheField,
		clubs.CCM:  service.config.Positions.CCM.PositionInTheField,
		clubs.RCM:  service.config.Positions.RCM.PositionInTheField,
		clubs.LCM:  service.config.Positions.LCM.PositionInTheField,
		clubs.LM:   service.config.Positions.LM.PositionInTheField,
		clubs.RM:   service.config.Positions.RM.PositionInTheField,
		clubs.CCAM: service.config.Positions.CCAM.PositionInTheField,
		clubs.RCAM: service.config.Positions.RCAM.PositionInTheField,
		clubs.LCAM: service.config.Positions.LCAM.PositionInTheField,
		clubs.LWB:  service.config.Positions.LWB.PositionInTheField,
		clubs.RWB:  service.config.Positions.RWB.PositionInTheField,
		clubs.LW:   service.config.Positions.LW.PositionInTheField,
		clubs.RW:   service.config.Positions.RW.PositionInTheField,
		clubs.RST:  service.config.Positions.RST.PositionInTheField,
		clubs.CST:  service.config.Positions.CST.PositionInTheField,
		clubs.LST:  service.config.Positions.LST.PositionInTheField,
	}

	var cardsWithPositions []SquadCardWithPosition
	var err error

	for _, card := range squadCards {
		var cardWithPosition SquadCardWithPosition

		cardWithPosition.Card, err = service.cards.Get(ctx, card.Card.ID)
		if err != nil {
			fmt.Println("could not convert positions", cardWithPosition, "\n",  err)
			return cardsWithPositions, ErrMatches.Wrap(err)
		}

		cardWithPosition.Position = PositionToCoordinates[card.Position]

		cardsWithPositions = append(cardsWithPositions, cardWithPosition)
	}

	return cardsWithPositions, nil
}

// ReflectPositions reflects card positions to another part of field.
func (service *Service) ReflectPositions(ctx context.Context, cardsWithPositions []SquadCardWithPosition) []SquadCardWithPosition {
	var newCardsWithPositions []SquadCardWithPosition

	for _, cardWithPositions := range cardsWithPositions {
		var newCardWithPositions SquadCardWithPosition

		newCardWithPositions.Card = cardWithPositions.Card
		newCardWithPositions.Position.X += service.config.SizeOfFieldByOX - cardWithPositions.Position.X
		fmt.Println(service.config.SizeOfFieldByOX - cardWithPositions.Position.X)

		newCardsWithPositions = append(newCardsWithPositions, newCardWithPositions)
	}

	return newCardsWithPositions
}

// GenerateBallPosition returns middle field position.
func (service *Service) GenerateBallPosition() PositionInTheField {
	return PositionInTheField{
		X : service.config.SizeOfFieldByOX / 2,
		Y : service.config.SizeOfFieldByOY / 2,
	}
}

// GenerateActionsForCards generate possible actions for each card from squad.
func (service *Service) GenerateActionsForCards(ctx context.Context, cardsWithPositions []SquadCardWithPosition, ballPosition PositionInTheField) {
	var possibleActions []CardPossibleAction

	for _, card := range cardsWithPositions {
		if card.Position.Compare(ballPosition) {
			// actions with ball.
		}
		// move action.
		var minCoordinateOX int
		var minCoordinateOY int
		var maxCoordinateOX int
		var maxCoordinateOY int

		var numOfCells int

		switch {
		// TODO: put to config.
		case card.Card.RunningSpeed > 0 && card.Card.RunningSpeed < 49 :
			numOfCells = 2
		case card.Card.RunningSpeed > 50 && card.Card.RunningSpeed < 69 :
			numOfCells = 3
		case card.Card.RunningSpeed > 70 && card.Card.RunningSpeed < 89 :
			numOfCells = 4
		case card.Card.RunningSpeed > 90 && card.Card.RunningSpeed < 100 :
			numOfCells = 5
		}

		minCoordinateOX = card.Position.X - numOfCells
		minCoordinateOY = card.Position.Y - numOfCells
		maxCoordinateOX = card.Position.X + numOfCells
		maxCoordinateOY = card.Position.X + numOfCells

		for i := minCoordinateOX; i <= maxCoordinateOX; i++ {
			for j := minCoordinateOY; j <= maxCoordinateOY; j++ {
				if i == service.config.SizeOfFieldByOX || i == service.config.SizeOfFieldByOY || i < 0 || j < 0 {
					continue
				}

				possibleActions = append(possibleActions, CardPossibleAction{
					CardID: card.Card.ID,
					Action: ActionMove,
					Positions: 	PositionInTheField {
						X : i,
						Y : j,
					},
				})
			}
		}
	}
}

// GenerateActionResult generates result of action.
func (service *Service) GenerateActionResult(ctx context.Context, cards []SquadCardWithPosition, action Action) {
	switch action {
	case ActionMove:
		// formula
	case ActionMoveWithBall:
		// formula
	case ActionPass:
		// formula
	case ActionCrossPass:
		// formula
	case ActionPassThrough:
		// formula
	case ActionDirectShot:
		// formula
	case ActionCurlShot:
		// formula
	case ActionTakeawayShot:
		// formula
	case ActionTackle:
		// formula
	case ActionSlidingTackle:
		// formula
	case ActionDribbling:
		// formula
	case ActionFeints:
		// formula
	}
}

// Get returns match by id.
func (service *Service) Get(ctx context.Context, matchID uuid.UUID) (Match, error) {
	match, err := service.matches.Get(ctx, matchID)

	return match, ErrMatches.Wrap(err)
}

// List returns page of matches.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	allMatches, err := service.matches.ListMatches(ctx, cursor)

	return allMatches, ErrMatches.Wrap(err)
}

// Delete deletes match.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrMatches.Wrap(service.matches.Delete(ctx, id))
}

// ListMatchGoals returns all goals scored in the match.
func (service *Service) ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]MatchGoals, error) {
	matchGoals, err := service.matches.ListMatchGoals(ctx, matchID)

	return matchGoals, ErrMatches.Wrap(err)
}

// GetGameResult returns goals of each user in the match.
func (service *Service) GetGameResult(ctx context.Context, matchID uuid.UUID) (GameResult, error) {
	matchResults, err := service.matches.GetMatchResult(ctx, matchID)
	if err != nil {
		return GameResult{}, ErrMatches.Wrap(err)
	}

	matchGoals, err := service.ListMatchGoals(ctx, matchID)
	if err != nil {
		return GameResult{}, ErrMatches.Wrap(err)
	}

	gameResult := GameResult{
		MatchResults: matchResults,
	}

	if len(gameResult.MatchResults) == 2 {
		for k, result := range gameResult.MatchResults {
			for _, goal := range matchGoals {
				if goal.UserID == result.UserID {
					card, err := service.cards.Get(ctx, goal.CardID)
					if err != nil {
						return gameResult, ErrMatches.Wrap(err)
					}

					gameResult.MatchResults[k].Goalscorers = append(gameResult.MatchResults[k].Goalscorers, Goalscorer{
						Card:   card,
						Minute: goal.Minute,
					})
				}
			}
		}

		return gameResult, nil
	}

	match, err := service.matches.Get(ctx, matchID)
	if err != nil {
		return gameResult, ErrMatches.Wrap(err)
	}

	var newGameResult GameResult
	newGameResult.MatchResults = append(newGameResult.MatchResults, MatchResult{UserID: match.User1ID})
	newGameResult.MatchResults = append(newGameResult.MatchResults, MatchResult{UserID: match.User2ID})

	for k, result := range newGameResult.MatchResults {
		for _, res := range matchResults {
			if result.UserID == res.UserID {
				newGameResult.MatchResults[k].QuantityGoals = res.QuantityGoals
			}
		}

		for _, goal := range matchGoals {
			if goal.UserID == result.UserID {
				card, err := service.cards.Get(ctx, goal.CardID)
				if err != nil {
					return newGameResult, ErrMatches.Wrap(err)
				}

				newGameResult.MatchResults[k].Goalscorers = append(newGameResult.MatchResults[k].Goalscorers, Goalscorer{
					Card:   card,
					Minute: goal.Minute,
				})
			}
		}
	}

	return newGameResult, nil
}

// ListSquadMatches returns all club matches in season.
func (service *Service) ListSquadMatches(ctx context.Context, seasonID int) ([]Match, error) {
	allMatches, err := service.matches.ListSquadMatches(ctx, seasonID)
	return allMatches, ErrMatches.Wrap(err)
}

// RankMatch evaluates how many points each user receive per match.
func (service *Service) RankMatch(ctx context.Context, match Match, matchGoals []MatchGoals) error {
	var (
		user1Goals int
		user2Goals int
	)

	for _, goal := range matchGoals {
		if goal.UserID == match.User1ID {
			user1Goals++
			continue
		}
		user2Goals++
	}

	switch {
	case user1Goals > user2Goals:
		match.User1Points = service.config.NumberOfPointsForWin
		match.User2Points = service.config.NumberOfPointsForLosing
	case user1Goals < user2Goals:
		match.User1Points = service.config.NumberOfPointsForLosing
		match.User2Points = service.config.NumberOfPointsForWin
	case user1Goals == user2Goals:
		match.User1Points = service.config.NumberOfPointsForDraw
		match.User2Points = service.config.NumberOfPointsForDraw
	}

	return ErrMatches.Wrap(service.matches.UpdateMatch(ctx, match))
}

// GetStatistic returns statistic of club in season.
func (service *Service) GetStatistic(ctx context.Context, club clubs.Club, seasonID int) (Statistic, error) {
	var statistic Statistic

	allMatches, err := service.ListSquadMatches(ctx, seasonID)
	if err != nil {
		return statistic, ErrMatches.Wrap(err)
	}

	if len(allMatches) < MinNumberOfMatches {
		return statistic, nil
	}

	if len(allMatches) > maxNumberOfMatches {
		allMatches = allMatches[:maxNumberOfMatches]
	}

	var (
		goalScored    int
		goalsConceded int
	)

	for _, match := range allMatches {
		statistic.MatchPlayed++

		if match.User1ID == club.OwnerID {
			switch {
			case match.User1Points == service.config.NumberOfPointsForWin:
				statistic.Wins++
			case match.User1Points == service.config.NumberOfPointsForDraw:
				statistic.Draws++
			case match.User1Points == service.config.NumberOfPointsForLosing:
				statistic.Losses++
			}
		} else if match.User2ID == club.OwnerID {
			switch {
			case match.User2Points == service.config.NumberOfPointsForWin:
				statistic.Wins++
			case match.User2Points == service.config.NumberOfPointsForDraw:
				statistic.Draws++
			case match.User2Points == service.config.NumberOfPointsForLosing:
				statistic.Losses++
			}
		} else {
			return statistic, err
		}
		matchGoals, err := service.ListMatchGoals(ctx, match.ID)
		if err != nil {
			return statistic, ErrMatches.Wrap(err)
		}
		for _, goal := range matchGoals {
			if goal.UserID == club.OwnerID {
				goalScored++
				continue
			}
			goalsConceded++
		}
	}

	statistic.Points = service.config.NumberOfPointsForWin*statistic.Wins + service.config.NumberOfPointsForDraw*statistic.Draws +
		+service.config.NumberOfPointsForLosing*statistic.Losses

	statistic.GoalDifference = goalScored - goalsConceded
	statistic.Club = club

	return statistic, nil
}
