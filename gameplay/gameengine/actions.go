// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"fmt"
	"sort"
	"time"
	"ultimatedivision/cards"
	"ultimatedivision/pkg/rand"

	"github.com/zeebo/errs"

	"github.com/google/uuid"
)

// Action defines list of possible player action in the field.
type Action string

const (
	// ActionMove defines move action by player.
	ActionMove Action = "move"
	// ActionMoveWithBall defines move action by player with ball.
	ActionMoveWithBall Action = "moveWithBall"
	// ActionPass defines pass by player to another player.
	ActionPass Action = "pass"
	// ActionCrossPass defines passing the ball by throwing it into the air in the direction of a player on his team.
	ActionCrossPass Action = "crossPass"
	// ActionPassThrough defines pass in free zone on the move often between players of the other team.
	ActionPassThrough Action = "passTrough"
	// ActionDirectShot defines direct shot.
	ActionDirectShot Action = "directShot"
	// ActionCurlShot defines curl shot.
	ActionCurlShot Action = "curlShot"
	// ActionTakeawayShot defines powerful shot from the box.
	ActionTakeawayShot Action = "takeawayShot"
	// ActionTackle defines tackling the ball from an opponent.
	ActionTackle Action = "tackle"
	// ActionSlidingTackle defines tackle by sliding on the field.
	ActionSlidingTackle Action = "slidingTackle"
	// ActionDribbling defines action when player move with some feints ot tricks.
	ActionDribbling Action = "dribbling"
	// ActionFeints defines action when player show feints.
	ActionFeints Action = "feints"
)

// GameConfig contains config values related to game and characteristics of cards and there ratio.
type GameConfig struct {
	MoveAction struct {
		NumOfCellsForFirstRange  int `json:"numOfCellsForFirstRange"`
		NumOfCellsForSecondRange int `json:"numOfCellsForSecondRange"`
		NumOfCellsForThirdRange  int `json:"numOfCellsForThirdRange"`
		NumOfCellsForFourthRange int `json:"numOfCellsForFourthRange"`

		FirstRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"firstRange"`
		SecondRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"secondRange"`
		ThirdRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"thirdRange"`
		FourthRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"fourthRange"`
	} `json:"moveAction"`
}

// CardAvailableAction defines in which position card could be placed and which action it could do there.
type CardAvailableAction struct {
	CardID    uuid.UUID    `json:"cardId"`
	Action    Action       `json:"action"`
	Positions []Coordinate `json:"positions"`
}

// MatchRepresentation defines user1 and user2 cards with positions,
// ball position an the moment and available actions for user cards.
type MatchRepresentation struct {
	User1CardsWithPosition []CardWithCoordinate  `json:"user1CardsWithPosition"`
	User2CardsWithPosition []CardWithCoordinate  `json:"user2CardsWithPosition"`
	BallCoordinate         Coordinate            `json:"ballPosition"`
	Actions                []MakeAction          `json:"actions"`
	CardAvailableAction    []CardAvailableAction `json:"cardAvailableAction"`
}

// Result defines a list of possible action results.
type Result bool

const (
	// ResultSuccessful defines that result of action successful.
	ResultSuccessful Result = true
	// ResultUnsuccessful defines that result of action unsuccessful.
	ResultUnsuccessful Result = false
)

// MakeAction defines request for every action.
type MakeAction struct {
	CardsLayout       []CardWithCoordinate `json:"cardsLayout"`
	BallPosition      Coordinate           `json:"ballPosition"`
	PlayerID          uuid.UUID            `json:"playerId"`
	Action            Action               `json:"action"`
	StartCoordinate   Coordinate           `json:"startCoordinate"`
	EndCoordinate     Coordinate           `json:"endCoordinate"`
	ReceiverPlayerID  uuid.UUID            `json:"receiverPlayerId"`
	OpponentPlayerIDs []uuid.UUID          `json:"opponentPlayerIds"`
	ActionTime        time.Time            `json:"actionDate"`
	Result            Result               `json:"result"`
}

// IsValid checks is action request valid.
func (a MakeAction) IsValid() bool {
	switch {
	case a.Action == ActionMove || a.Action == ActionMoveWithBall:
		if a.ReceiverPlayerID != uuid.Nil {
			return false
		}
	default:
		if a.ReceiverPlayerID == uuid.Nil {
			return false
		}
	}

	if a.PlayerID == uuid.Nil {
		return false
	}

	if a.Action == ActionMove || a.Action == ActionMoveWithBall || a.Action == ActionPass ||
		a.Action == ActionCrossPass || a.Action == ActionPassThrough || a.Action == ActionDirectShot ||
		a.Action == ActionCurlShot || a.Action == ActionTakeawayShot || a.Action == ActionTackle ||
		a.Action == ActionSlidingTackle || a.Action == ActionDribbling || a.Action == ActionFeints {
		return true
	}

	return false
}

// ActionResult contains result for actions.
type ActionResult struct {
}

// HandleAction handles all match actions and generates result of actions.
func HandleAction(matchRepresentation MatchRepresentation, makeAction []MakeAction, gameConfig GameConfig, coordinatesConfig CoordinatesConfig) ([]MakeAction, error) {
	sort.Slice(makeAction, func(i, j int) bool {
		return makeAction[i].ActionTime.Before(makeAction[j].ActionTime)
	})

	var actions []MakeAction
	for _, action := range makeAction {
		if !action.IsValid() {
			return actions, errs.New("invalid action")
		}

		actionWithResult, err := GenerateActionResult(action, matchRepresentation, gameConfig, coordinatesConfig)
		if err != nil {
			return actions, err
		}

		actions = append(actions, actionWithResult)
	}

	return actions, nil
}

// GenerateActionResult generates result of action.
func GenerateActionResult(makeAction MakeAction, representation MatchRepresentation, gameConfig GameConfig, coordinatesConfig CoordinatesConfig) (MakeAction, error) {
	switch makeAction.Action {
	case ActionMove:
		availableActions, err := GenerateAvailableActions(representation.User1CardsWithPosition, representation.User2CardsWithPosition, representation.BallCoordinate, gameConfig, coordinatesConfig)
		if err != nil {
			return MakeAction{}, err
		}
		if makeAction.isActionInSlice(availableActions) {
			makeAction.Result = ResultSuccessful
		}
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

	return makeAction, nil
}

// isActionInSlice checks is available actions contains action.
func (a MakeAction) isActionInSlice(availableActions []CardAvailableAction) bool {
	for _, action := range availableActions {
		if a.Action == action.Action && a.PlayerID == action.CardID {
			for _, coordinate := range action.Positions {
				if coordinate.Compare(a.EndCoordinate) {
					return true
				}
			}
		}
	}
	return false
}

// GenerateAvailableActions generates available actions for card based card coordinate and ball position in the field.
func GenerateAvailableActions(alliesCardsWithCoordinates, opponentCardsWithCoordinates []CardWithCoordinate, ballCoordinate Coordinate, gameConfig GameConfig, coordinatesConfig CoordinatesConfig) ([]CardAvailableAction, error) {
	var availableActions []CardAvailableAction

	coordinatesOfOpponentCards := getCoordinates(opponentCardsWithCoordinates)

	for _, alliesCardWithCoordinates := range alliesCardsWithCoordinates {
		// action move
		numOfCells, err := getNumOfMaximumCells(gameConfig, alliesCardWithCoordinates.Card, ActionMove)
		if err != nil {
			return nil, err
		}
		fmt.Println("numOfCells", numOfCells)

		cells := generateCellsInRange(alliesCardWithCoordinates.Coordinate, numOfCells, coordinatesConfig.SizeOfFieldByOX, coordinatesConfig.SizeOfFieldByOY, coordinatesOfOpponentCards)
		fmt.Println(cells)
		moveAction := CardAvailableAction{
			CardID:    alliesCardWithCoordinates.Card.ID,
			Action:    ActionMove,
			Positions: cells,
		}

		availableActions = append(availableActions, moveAction)
	}

	return availableActions, nil
}

// GetCoordinates returns coordinates from slice with cards with coordinates.
func getCoordinates(cardsWithCoordinate []CardWithCoordinate) []Coordinate {
	var coordinates []Coordinate

	for _, cardWithCoordinate := range cardsWithCoordinate {
		coordinates = append(coordinates, Coordinate{
			X: cardWithCoordinate.Coordinate.X,
			Y: cardWithCoordinate.Coordinate.Y,
		})
	}

	return coordinates
}

// GetNumOfMaximumCells generates the maximum number of cells
// on which the action is distributed for a certain card.
func getNumOfMaximumCells(config GameConfig, card cards.Card, action Action) (int, error) {
	var numOfCells int
	switch {
	case action == ActionMove:
		runningSpeed, err := rand.RandomInRange(card.RunningSpeed)
		if err != nil {
			return 0, err
		}

		switch {
		case runningSpeed > config.MoveAction.FirstRange.Min && runningSpeed < config.MoveAction.FirstRange.Max:
			numOfCells = config.MoveAction.NumOfCellsForFirstRange
		case runningSpeed > config.MoveAction.SecondRange.Min && runningSpeed < config.MoveAction.SecondRange.Max:
			numOfCells = config.MoveAction.NumOfCellsForSecondRange
		case runningSpeed > config.MoveAction.ThirdRange.Min && runningSpeed < config.MoveAction.ThirdRange.Max:
			numOfCells = config.MoveAction.NumOfCellsForThirdRange
		case runningSpeed > config.MoveAction.FourthRange.Min && runningSpeed < config.MoveAction.ThirdRange.Max:
			numOfCells = config.MoveAction.NumOfCellsForFourthRange
		}
	case action == ActionMoveWithBall:
		// TODO: add others action here
	}

	return numOfCells, nil
}
