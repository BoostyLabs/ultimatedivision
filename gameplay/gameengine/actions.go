// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/pkg/rand"
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

// MakeAction defines fields that describes football action.
type MakeAction struct {
	CardsLayout       []CardWithCoordinate `json:"cardsLayout"`
	BallPosition      Coordinate           `json:"ballPosition"`
	PlayerID          uuid.UUID            `json:"playerId"`
	Action            Action               `json:"action"`
	StartCoordinate   Coordinate           `json:"startCoordinate"`
	EndCoordinate     Coordinate           `json:"endCoordinate"`
	ReceiverPlayerID  uuid.UUID            `json:"receiverPlayerId"`
	OpponentPlayerIDs []uuid.UUID          `json:"opponentPlayerIds"`
	ActionTime        time.Time            `json:"actionTime"`
	Result            Result               `json:"result"`
}

// IsValid checks is action valid.
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

// PrepareMatchRepresentation prepares data to send for both match participants which contains
// cards with coordinates, ball position and available actions for all cards.
func PrepareMatchRepresentation(coordinateConfig CoordinatesConfig, gameConfig GameConfig, firstSquad, secondSquad []clubs.SquadCard) (MatchRepresentation, error) {
	firstClientCardsWithCoordinates := ConvertPositionToCoordinate(coordinateConfig, firstSquad)
	ballCoordinate := GetCenterOfField(coordinateConfig.SizeOfFieldByOX, coordinateConfig.SizeOfFieldByOY)
	secondClientCardsWithCoordinates := ConvertPositionToCoordinate(coordinateConfig, secondSquad)
	secondClientCardsWithCoordinates = ReflectCoordinates(secondClientCardsWithCoordinates, coordinateConfig.SizeOfFieldByOX, coordinateConfig.SizeOfFieldByOY)
	// generates available action for each squad.
	var cardsAvailableActions []CardAvailableAction
	firstSquadCardAvailableActions, err := GenerateAvailableActions(firstClientCardsWithCoordinates, secondClientCardsWithCoordinates, ballCoordinate, gameConfig, coordinateConfig)
	if err != nil {
		return MatchRepresentation{}, err
	}
	secondSquadCardAvailableActions, err := GenerateAvailableActions(secondClientCardsWithCoordinates, firstClientCardsWithCoordinates, ballCoordinate,
		gameConfig, coordinateConfig)
	if err != nil {
		return MatchRepresentation{}, err
	}

	cardsAvailableActions = append(cardsAvailableActions, firstSquadCardAvailableActions...)
	cardsAvailableActions = append(cardsAvailableActions, secondSquadCardAvailableActions...)

	matchRepresentation := MatchRepresentation{
		User1CardsWithPosition: firstClientCardsWithCoordinates,
		User2CardsWithPosition: secondClientCardsWithCoordinates,
		BallCoordinate:         ballCoordinate,
		CardAvailableAction:    cardsAvailableActions,
	}

	return matchRepresentation, nil
}

// HandleAction handles all match actions and generates result of them.
func HandleAction(matchRepresentation MatchRepresentation, makeAction []MakeAction, gameConfig GameConfig, coordinatesConfig CoordinatesConfig) ([]MakeAction, error) {
	sort.Slice(makeAction, func(i, j int) bool {
		return makeAction[i].ActionTime.Before(makeAction[j].ActionTime)
	})

	var actions []MakeAction
	for _, action := range makeAction {
		if !action.IsValid() {
			return actions, errs.New("invalid action")
		}

		actionWithResult, err := generateActionResult(action, matchRepresentation, gameConfig, coordinatesConfig)
		if err != nil {
			return actions, err
		}

		actions = append(actions, actionWithResult)
	}

	return actions, nil
}

// GenerateActionResult generates result of action.
func generateActionResult(action MakeAction, representation MatchRepresentation, gameConfig GameConfig, coordinatesConfig CoordinatesConfig) (MakeAction, error) {
	availableActions, err := GenerateAvailableActions(representation.User1CardsWithPosition, representation.User2CardsWithPosition, representation.BallCoordinate, gameConfig, coordinatesConfig)
	if err != nil {
		return action, err
	}

	if action.isActionInSlice(availableActions) {
		action.Result = ResultSuccessful
	}

	return action, nil
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
		// action move.
		numOfCells, err := getNumOfMaximumCells(gameConfig, alliesCardWithCoordinates.Card, ActionMove)
		if err != nil {
			return nil, err
		}

		cells := generateCellsInRange(alliesCardWithCoordinates.Coordinate, numOfCells, coordinatesConfig.SizeOfFieldByOX, coordinatesConfig.SizeOfFieldByOY, coordinatesOfOpponentCards)
		moveAction := CardAvailableAction{
			CardID:    alliesCardWithCoordinates.Card.ID,
			Action:    ActionMove,
			Positions: cells,
		}

		availableActions = append(availableActions, moveAction)
	}

	return availableActions, nil
}

// getCoordinates returns coordinates of opponent cards.
func getCoordinates(cardsWithCoordinate []CardWithCoordinate) []Coordinate {
	var coordinates []Coordinate

	for i := 0; i < len(cardsWithCoordinate); i++ {
		coordinates = append(coordinates, Coordinate{
			X: cardsWithCoordinate[i].Coordinate.X,
			Y: cardsWithCoordinate[i].Coordinate.Y,
		})
	}

	return coordinates
}

// getNumOfMaximumCells generates the maximum number of cells
// on which the action could be distributed for a certain card.
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
		// TODO: add others action here.
	}

	return numOfCells, nil
}
