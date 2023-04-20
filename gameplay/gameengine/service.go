// Copyright (C) 2021 - 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"context"
	"encoding/json"
	"math/rand"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/clubs"
	"ultimatedivision/gameplay/matches"
)

// ErrGameEngine indicates that there was an error in the service.
var ErrGameEngine = errs.Class("game engine service error")

var (
	topLine    = []int{77, 70, 63, 56, 49, 42, 35, 28, 21, 14, 7, 0}
	bottomLine = []int{6, 13, 20, 27, 34, 41, 48, 55, 62, 69, 76, 83}
)

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
	games   DB
	clubs   *clubs.Service
	avatars *avatars.Service
	cards   *cards.Service
	matches *matches.Service
	config  Config
}

// NewService is a constructor for game engine service.
func NewService(games DB, clubs *clubs.Service, avatars *avatars.Service, cards *cards.Service, matches *matches.Service, config Config) *Service {
	return &Service{
		games:   games,
		clubs:   clubs,
		avatars: avatars,
		cards:   cards,
		matches: matches,
		config:  config,
	}
}

const (
	minPlace = 0
	maxPlace = 83
)

// GetCardMoves get all card possible moves.
func (service *Service) GetCardMoves(cardPlace int, isThreeSteps bool) ([]int, error) {
	bottom1 := []int{82, 75, 68, 61, 54, 47, 40, 33, 26, 19, 12, 5}
	bottom2 := []int{81, 74, 67, 60, 53, 46, 39, 32, 25, 18, 11, 4}

	top1 := []int{71, 64, 57, 50, 43, 36, 29, 22, 15}
	top2 := []int{72, 65, 58, 51, 44, 37, 30, 23, 16, 9, 2, 79}

	exceptions := []int{1, 5, 78, 82, 8}

	if cardPlace < minPlace || cardPlace > maxPlace {
		return []int{}, ErrGameEngine.New("player place can not be more 83 or les than 0, player place is %d", cardPlace)
	}

	var stepInWidth []int
	var moves []int

	if isThreeSteps == true {

		switch {
		case contains(topLine, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace, cardPlace+1, cardPlace+2, cardPlace+3)

		case contains(top1, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		case contains(top2, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2, cardPlace+3)

		case contains(bottomLine, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace)

		case contains(bottom1, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1)

		case contains(bottom2, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		case contains(exceptions, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2, cardPlace+3)

		case cardPlace == 12:
			stepInWidth = append(stepInWidth, cardPlace-3, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1)

		}

		for _, w := range stepInWidth {
			min := w - 14
			max := w + 14
			min21 := w - 21
			max21 := w + 21
			moves = append(moves, min, min+7, max-7, max, w, min21, max21)
		}

	} else {
		switch {
		case contains(topLine, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace, cardPlace+1, cardPlace+2)

		case contains(bottomLine, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace)

		case contains(exceptions, cardPlace):
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1)

		case cardPlace == 8:
			stepInWidth = append(stepInWidth, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		case cardPlace == 12:
			stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1)

		default:
		}
		stepInWidth = append(stepInWidth, cardPlace-2, cardPlace-1, cardPlace, cardPlace+1, cardPlace+2)

		for _, w := range stepInWidth {
			min := w - 14
			max := w + 14
			moves = append(moves, min, min+7, max-7, max, w)
		}
	}

	sort.Ints(moves)
	moves = removeMin(moves, minPlace)
	moves = removeMax(moves, maxPlace)

	return moves, nil
}

// IsInterrupted check if pass is interrupted.
func (service *Service) IsInterrupted(shortPassing, interceptions int) bool {
	if shortPassing > interceptions {
		return true
	}
	return false
}

// GetCardPasses get all field cells possible to pass.
func (service *Service) GetCardPasses(teamPositions, availablePassCells []int) []int {
	var availablePasses []int
	for _, teammatePosition := range teamPositions {
		if contains(availablePassCells, teammatePosition) {
			availablePasses = append(availablePasses, teammatePosition)
		}
	}
	return availablePasses
}

// GivePass get info about pass and return final ball cell.
func (service *Service) GivePass(passWay []int, passReceiverStats CardWithPosition, passGiverCard, passReceiverCard cards.Card, opponentsWithPosition []CardWithPosition) int {
	var passPosition int
	for _, opponent := range opponentsWithPosition {
		if contains(passWay, opponent.FieldPosition) {
			if opponent.Card.Interceptions > passGiverCard.ShortPassing {
				return opponent.FieldPosition
			}
		}
	}
	if passReceiverCard.BallControl > 10 {
		return passReceiverStats.FieldPosition
	} else {
		return ballBounce(passReceiverStats.FieldPosition)
	}

	return passPosition
}

// ballBounce calculates the position of the ball bounce.
func ballBounce(position int) int {
	var bounceBall []int

	switch {
	case contains(topLine, position):
		bounceBall = append(bounceBall, position, position+1)

	case contains(bottomLine, position):
		bounceBall = append(bounceBall, position-1, position)

	case position == 8:
		bounceBall = append(bounceBall, position-1, position, position+1)

	case position == 12:
		bounceBall = append(bounceBall, position-1, position)
	default:
		bounceBall = append(bounceBall, position-1, position, position+1)
	}

	var ballCells []int

	for _, i2 := range bounceBall {
		ballCells = append(ballCells, i2+7, i2-7, i2)
	}
	sort.Ints(ballCells)
	ballCells = removeMin(ballCells, minPlace)
	ballCells = removeMax(ballCells, maxPlace)

	// Seed the random number generator with the current time.
	rand.Seed(time.Now().UnixNano())

	// Generate a random index within the range of the array.
	randomIndex := rand.Intn(len(ballCells))

	// Retrieve the element at the random index.
	randomElement := ballCells[randomIndex]

	return randomElement
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

func removeIntersections(moves, playerPositions []int) (movesWithoutIntersections []int) {
	for _, v := range moves {
		if !contains(playerPositions, v) {
			movesWithoutIntersections = append(movesWithoutIntersections, v)
		}
	}
	return movesWithoutIntersections
}

// Move update card moves and get possible moves cells.
func (service *Service) Move(ctx context.Context, matchID uuid.UUID, card CardIDWithPosition, isCardFast bool) (CardAvailableAction, error) {
	var cardAvailableAction CardAvailableAction

	gameInfoJSON, err := service.games.Get(ctx, matchID)
	if err != nil {
		return CardAvailableAction{}, ErrGameEngine.Wrap(err)
	}

	var game Game
	game.MatchID = matchID

	err = json.Unmarshal([]byte(gameInfoJSON), &game.GameInfo)
	if err != nil {
		return CardAvailableAction{}, ErrGameEngine.Wrap(err)
	}

	var moves []int
	var allPositionsInUse []int

	allPositionsInUse = append(allPositionsInUse, card.Position)

	for _, cardWithPosition := range game.GameInfo {
		allPositionsInUse = append(allPositionsInUse, cardWithPosition.Position)
	}

	// check whether the position we want to go to is occupied.
	if contains(allPositionsInUse, card.Position) {
		return CardAvailableAction{}, ErrGameEngine.New("Can not move to position, already in use")
	}

	// check, Update and get all possible moves.
	for i, cardData := range game.GameInfo {
		if cardData.CardID == card.CardID {
			game.GameInfo[i].Position = card.Position

			newGameInfoJSON, err := json.Marshal(game.GameInfo)
			if err != nil {
				return CardAvailableAction{}, ErrGameEngine.Wrap(err)
			}

			err = service.games.Update(ctx, matchID, string(newGameInfoJSON))
			if err != nil {
				return CardAvailableAction{}, ErrGameEngine.Wrap(err)
			}

			moves, err = service.GetCardMoves(card.Position, isCardFast)
			if err != nil {
				return CardAvailableAction{}, ErrGameEngine.Wrap(err)
			}
		}
	}

	// remove already occupied positions.
	moves = removeIntersections(moves, allPositionsInUse)

	cardAvailableAction = CardAvailableAction{
		Action:        ActionMove,
		CardID:        card.CardID,
		FieldPosition: moves,
	}

	return cardAvailableAction, nil
}

// GameInformation creates a player by user.
func (service *Service) GameInformation(ctx context.Context, player1SquadID, player2SquadID uuid.UUID) (MatchRepresentation, error) {
	var cardsWithPositionPlayer1 []CardWithPosition
	var cardsWithPositionPlayer2 []CardWithPosition
	var cardsWithPositionLeftPlayer []CardWithPosition
	var cardsWithPositionRightPlayer []CardWithPosition
	var cardsAvailableAction []CardAvailableAction
	var ballPosition int

	squadCardsPlayer1, err := service.clubs.ListCards(ctx, player1SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	squadCardsPlayer2, err := service.clubs.ListCards(ctx, player2SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	squadPlayer1, err := service.clubs.GetSquad(ctx, player1SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	squadPlayer2, err := service.clubs.GetSquad(ctx, player2SquadID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	clubPlayer1, err := service.clubs.Get(ctx, squadPlayer1.ClubID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	clubPlayer2, err := service.clubs.Get(ctx, squadPlayer2.ClubID)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	var matchInfo []CardIDWithPosition
	var leftSidePositions []int

	for _, sqCard := range squadCardsPlayer1 {
		fieldPosition := service.squadPositionToFieldPositionLeftSide(sqCard.Position)

		if ballPosition < fieldPosition {
			ballPosition = fieldPosition
		}

		leftSidePositions = append(leftSidePositions, fieldPosition)

		cardWithPositionLeftPlayer := CardWithPosition{
			Card:          sqCard.Card,
			FieldPosition: fieldPosition,
		}

		cardsWithPositionLeftPlayer = append(cardsWithPositionLeftPlayer, cardWithPositionLeftPlayer)
	}

	var rightSidePositions []int

	for _, sqCard := range squadCardsPlayer2 {
		fieldPosition := service.squadPositionToFieldPositionRightSide(sqCard.Position)

		rightSidePositions = append(rightSidePositions, fieldPosition)

		cardWithPositionRightPlayer := CardWithPosition{
			Card:          sqCard.Card,
			FieldPosition: fieldPosition,
		}

		cardsWithPositionRightPlayer = append(cardsWithPositionRightPlayer, cardWithPositionRightPlayer)
	}

	for _, sqCard := range squadCardsPlayer1 {
		avatar, err := service.avatars.Get(ctx, sqCard.Card.ID)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardWithPositionPlayer := CardWithPosition{
			Card:          sqCard.Card,
			Avatar:        avatar,
			FieldPosition: service.squadPositionToFieldPositionLeftSide(sqCard.Position),
		}

		cardInfo := CardIDWithPosition{
			CardID:   sqCard.Card.ID,
			Position: cardWithPositionPlayer.FieldPosition,
		}

		matchInfo = append(matchInfo, cardInfo)

		isCardFast := false
		if sqCard.Card.RunningSpeed > 80 && cardInfo.Position == ballPosition || sqCard.Card.RunningSpeed > 70 {
			isCardFast = true
		}

		fieldPosition, err := service.GetCardMoves(cardWithPositionPlayer.FieldPosition, isCardFast)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardAvailableAction := CardAvailableAction{
			Action:        ActionMove,
			CardID:        sqCard.Card.ID,
			FieldPosition: fieldPosition,
		}

		if cardInfo.Position == ballPosition {
			passOptions := service.GetCardPasses(leftSidePositions, fieldPosition)

			cardAvailablePasses := CardAvailableAction{
				Action:        ActionPass,
				CardID:        sqCard.Card.ID,
				FieldPosition: passOptions,
			}
			cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction, cardAvailablePasses)
		} else {
			cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction)
		}

		cardsWithPositionPlayer1 = append(cardsWithPositionPlayer1, cardWithPositionPlayer)
	}

	for _, sqCard := range squadCardsPlayer2 {
		avatar, err := service.avatars.Get(ctx, sqCard.Card.ID)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardWithPositionPlayer := CardWithPosition{
			Card:          sqCard.Card,
			Avatar:        avatar,
			FieldPosition: service.squadPositionToFieldPositionRightSide(sqCard.Position),
		}

		isCardFast := false
		if sqCard.Card.RunningSpeed > 70 {
			isCardFast = true
		}
		cardWithPositionPlayer.FieldPosition = 9
		fieldPosition, err := service.GetCardMoves(cardWithPositionPlayer.FieldPosition, isCardFast)
		if err != nil {
			return MatchRepresentation{}, ErrGameEngine.Wrap(err)
		}

		cardAvailableAction := CardAvailableAction{
			Action:        ActionMove,
			CardID:        sqCard.Card.ID,
			FieldPosition: fieldPosition,
		}

		cardInfo := CardIDWithPosition{
			CardID:   sqCard.Card.ID,
			Position: cardWithPositionPlayer.FieldPosition,
		}

		matchInfo = append(matchInfo, cardInfo)

		if cardInfo.Position == ballPosition {
			passOptions := service.GetCardPasses(rightSidePositions, fieldPosition)

			cardAvailablePasses := CardAvailableAction{
				Action:        ActionPass,
				CardID:        sqCard.Card.ID,
				FieldPosition: passOptions,
			}
			cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction, cardAvailablePasses)
		} else {
			cardsAvailableAction = append(cardsAvailableAction, cardAvailableAction)
		}

		cardsWithPositionPlayer2 = append(cardsWithPositionPlayer2, cardWithPositionPlayer)
	}

	matchID, err := service.matches.CreateMatchID(ctx, player1SquadID, player2SquadID, clubPlayer1.OwnerID, clubPlayer2.OwnerID, 1)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	cardIDsWithPositionWithBallPosition := CardIDsWithPositionWithBallPosition{
		CardIDsWithPosition: matchInfo,
		BallPosition:        ballPosition,
	}

	gameInfo, err := json.Marshal(cardIDsWithPositionWithBallPosition)
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	err = service.games.Create(ctx, matchID, string(gameInfo))
	if err != nil {
		return MatchRepresentation{}, ErrGameEngine.Wrap(err)
	}

	return MatchRepresentation{
		MatchID:                matchID,
		User1CardsWithPosition: cardsWithPositionPlayer1,
		User2CardsWithPosition: cardsWithPositionPlayer2,
		BallPosition:           ballPosition,
		CardAvailableAction:    cardsAvailableAction,
		User1ClubInformation:   clubPlayer1,
		User2ClubInformation:   clubPlayer2,
		User1SquadInformation:  squadPlayer1,
		User2SquadInformation:  squadPlayer2,
		Rounds:                 service.config.Rounds,
	}, nil
}

// GameLogicByAction returns game logic by action.
func (service *Service) GameLogicByAction(ctx context.Context, matchID uuid.UUID, cardIDWithPosition CardIDWithPosition, action Action) (CardAvailableAction, error) {
	hasBall, err := service.ifCardHasBall(ctx, matchID, cardIDWithPosition.CardID)
	if err != nil {
		return CardAvailableAction{}, ErrGameEngine.Wrap(err)
	}

	card, err := service.cards.Get(ctx, cardIDWithPosition.CardID)
	if err != nil {
		return CardAvailableAction{}, ErrGameEngine.Wrap(err)
	}

	switch action {
	case ActionMove:
		isCardFast := false
		if hasBall && card.RunningSpeed > 80 || !hasBall && card.RunningSpeed > 70 {
			isCardFast = true
		}
		return service.Move(ctx, matchID, cardIDWithPosition, isCardFast)
	case ActionPass:

	}

	return CardAvailableAction{}, nil
}

func (service *Service) ifCardHasBall(ctx context.Context, matchID uuid.UUID, cardID uuid.UUID) (bool, error) {
	var cardIDsWithPositionWithBallPosition CardIDsWithPositionWithBallPosition

	gameInfo, err := service.games.Get(ctx, matchID)
	if err != nil {
		return false, ErrGameEngine.Wrap(err)
	}

	err = json.Unmarshal([]byte(gameInfo), &cardIDsWithPositionWithBallPosition)
	if err != nil {
		return false, ErrGameEngine.Wrap(err)
	}

	for _, position := range cardIDsWithPositionWithBallPosition.CardIDsWithPosition {
		if cardID == position.CardID {
			if position.Position == cardIDsWithPositionWithBallPosition.BallPosition {
				return true, nil
			}
		}
	}
	return false, nil
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
