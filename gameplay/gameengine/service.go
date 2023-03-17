// Copyright (C) 2021 - 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"context"
	"sort"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/clubs"
)

// ErrGameEngine indicates that there was an error in the service.
var ErrGameEngine = errs.Class("game engine service error")

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
	clubs   *clubs.Service
	avatars *avatars.Service
	cards   *cards.Service
	config  Config
}

// NewService is a constructor for game engine service.
func NewService(clubs *clubs.Service, avatars *avatars.Service, cards *cards.Service, config Config) *Service {
	return &Service{
		clubs:   clubs,
		avatars: avatars,
		cards:   cards,
		config:  config,
	}
}

const (
	minPlace = 0
	maxPlace = 83
)

// GetCardMoves get all card possible moves.
func (service *Service) GetCardMoves(cardPlace int) ([]int, error) {
	top := []int{77, 70, 63, 56, 49, 42, 35, 28, 21, 14, 7, 0}
	bottom := []int{6, 13, 20, 27, 34, 41, 48, 55, 62, 69, 76, 83}
	exceptions := []int{1, 5, 78, 82}

	if cardPlace < minPlace || cardPlace > maxPlace {
		return []int{}, ErrGameEngine.New("player place can not be more 83 or les than 0, player place is %d", cardPlace)
	}
	var stepInWidth []int

	switch {
	case contains(top, cardPlace):
		stepInWidth = append(stepInWidth, cardPlace, cardPlace+1, cardPlace+2)

	case contains(bottom, cardPlace):
		stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace)

	case contains(exceptions, cardPlace):
		stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1)

	case cardPlace == 8:
		stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

	case cardPlace == 12:
		stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1)

	default:
		stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)
	}

	var moves []int

	for _, w := range stepInWidth {
		min := w - 14
		max := w + 14
		moves = append(moves, min, min+7, max-7, max, w)
	}

	sort.Ints(moves)
	moves = removeMin(moves, minPlace)
	moves = removeMax(moves, maxPlace)

	return moves, nil
}

func removeMin(l []int, min int) []int {
	for i, v := range l {
		if v >= min {
			return l[i:]
		}
	}
	return l
}
func removeMax(l []int, max int) []int {
	for i, v := range l {
		if v > max {
			return l[:i]
		}
	}
	return l
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GameInformation creates a player by user.
func (service *Service) GameInformation(ctx context.Context, players1SquadID, players2SquadID uuid.UUID) (MatchRepresentation, error) {
	var cardsWithPositionPlayer1 []CardWithPosition
	var cardsWithPositionPlayer2 []CardWithPosition
	var cardsAvailableAction []CardAvailableAction

	squadPlayer1, err := service.clubs.ListCards(ctx, players1SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	squadPlayer2, err := service.clubs.ListCards(ctx, players2SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	for _, squad := range squadPlayer1 {
		avatar, err := service.avatars.Get(ctx, squad.Card.ID)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardWithPositionPlayer := CardWithPosition{
			Card:          squad.Card,
			Avatar:        avatar,
			FieldPosition: service.squadPositionToFieldPositionLeftSide(squad.Position),
		}

		fieldPosition, err := service.GetCardMoves(cardWithPositionPlayer.FieldPosition)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardAvailableAction := CardAvailableAction{
			Action:        ActionMove,
			CardID:        squad.Card.ID,
			FieldPosition: fieldPosition,
		}

		cardsWithPositionPlayer1 = append(cardsWithPositionPlayer1, cardWithPositionPlayer)
		cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction)
	}

	for _, squad := range squadPlayer2 {
		avatar, err := service.avatars.Get(ctx, squad.Card.ID)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardWithPositionPlayer := CardWithPosition{
			Card:          squad.Card,
			Avatar:        avatar,
			FieldPosition: service.squadPositionToFieldPositionRightSide(squad.Position),
		}

		fieldPosition, err := service.GetCardMoves(cardWithPositionPlayer.FieldPosition)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardAvailableAction := CardAvailableAction{
			Action:        ActionMove,
			CardID:        squad.Card.ID,
			FieldPosition: fieldPosition,
		}

		cardsWithPositionPlayer2 = append(cardsWithPositionPlayer2, cardWithPositionPlayer)
		cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction)
	}

	return MatchRepresentation{
		User1CardsWithPosition: cardsWithPositionPlayer1,
		User2CardsWithPosition: cardsWithPositionPlayer2,
		BallPosition:           0,
		CardAvailableAction:    cardsAvailableAction,
	}, nil
}

func (service *Service) squadPositionToFieldPositionLeftSide(squadPosition clubs.Position) int {
	switch squadPosition {
	case clubs.GK:
		return service.config.LeftSide.Goalkeeper
	case clubs.LB:
		return service.config.LeftSide.LeftBack
	case clubs.RB:
		return service.config.LeftSide.RightBack
	case clubs.LM:
		return service.config.LeftSide.LeftMid
	case clubs.RM:
		return service.config.LeftSide.RightMid
	case clubs.LCD:
		return service.config.LeftSide.CenterBackLeft
	case clubs.RCD:
		return service.config.LeftSide.CenterBackRight
	case clubs.LCM:
		return service.config.LeftSide.CenterMidLeft
	case clubs.RCM:
		return service.config.LeftSide.CenterMidRight
	case clubs.LST:
		return service.config.LeftSide.ForwardLeft
	case clubs.RST:
		return service.config.LeftSide.ForwardRight
	}

	return 0
}

func (service *Service) squadPositionToFieldPositionRightSide(squadPosition clubs.Position) int {
	switch squadPosition {
	case clubs.GK:
		return service.config.RightSide.Goalkeeper
	case clubs.LB:
		return service.config.RightSide.LeftBack
	case clubs.RB:
		return service.config.RightSide.RightBack
	case clubs.LM:
		return service.config.RightSide.LeftMid
	case clubs.RM:
		return service.config.RightSide.RightMid
	case clubs.LCD:
		return service.config.RightSide.CenterBackLeft
	case clubs.RCD:
		return service.config.RightSide.CenterBackRight
	case clubs.LCM:
		return service.config.RightSide.CenterMidLeft
	case clubs.RCM:
		return service.config.RightSide.CenterMidRight
	case clubs.LST:
		return service.config.RightSide.ForwardLeft
	case clubs.RST:
		return service.config.RightSide.ForwardRight
	}
	return 0
}
