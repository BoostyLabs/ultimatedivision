// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/gameplay/gameengine"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
	"ultimatedivision/seasons"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

// ErrQueue indicated that there was an error in service.
var ErrQueue = errs.Class("queue service error")

// Service is handling queues related logic.
//
// architecture: Service
type Service struct {
	config Config
	log    logger.Logger

	queues DB

	users            *users.Service
	clubs            *clubs.Service
	matches          *matches.Service
	seasons          *seasons.Service
	currencywaitlist *currencywaitlist.Service
}

// NewService is a constructor for queues service.
func NewService(config Config, log logger.Logger, queues DB, users *users.Service, clubs *clubs.Service,
	matches *matches.Service, currencywaitlist *currencywaitlist.Service, seasons *seasons.Service) *Service {
	return &Service{
		config:           config,
		log:              log,
		queues:           queues,
		users:            users,
		clubs:            clubs,
		matches:          matches,
		seasons:          seasons,
		currencywaitlist: currencywaitlist,
	}
}

// Create adds client to queue.
func (service *Service) Create(ctx context.Context, client Client) error {
	if _, err := service.users.Get(ctx, client.UserID); err != nil {
		return ErrQueue.Wrap(err)
	}

	squad, err := service.clubs.GetSquad(ctx, client.SquadID)
	if err != nil {
		return ErrQueue.Wrap(err)
	}

	_, err = service.clubs.Get(ctx, squad.ClubID)
	if err != nil {
		return ErrQueue.Wrap(err)
	}

	// TODO: add division ID to client.

	err = service.queues.Delete(client.UserID)
	if ErrNoClient.Has(err) || err == nil {
		service.queues.Create(client)
		return nil
	}

	return ErrQueue.Wrap(err)
}

// Get returns client from database.
func (service *Service) Get(userID uuid.UUID) (Client, error) {
	queue, err := service.queues.Get(userID)
	return queue, ErrQueue.Wrap(err)
}

// List returns clients from database.
func (service *Service) List() []Client {
	return service.queues.List()
}

// ListNotPlayingUsers returns clients who don't play game from database.
func (service *Service) ListNotPlayingUsers() []Client {
	return service.queues.ListNotPlayingUsers()
}

// UpdateIsPlaying updates is playing status of client in database.
func (service *Service) UpdateIsPlaying(userID uuid.UUID, isPlaying bool) error {
	return service.queues.UpdateIsPlaying(userID, isPlaying)
}

// Finish finishes client's queue in database.
func (service *Service) Finish(userID uuid.UUID) error {
	return service.queues.Delete(userID)
}

// Play method contains all the logic for playing matches.
func (service *Service) Play(ctx context.Context, firstClient, secondClient Client) error {
	squadCardsFirstClient, err := service.clubs.ListSquadCards(ctx, firstClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}
	if len(squadCardsFirstClient) != clubs.SquadSize {
		if err := firstClient.WriteJSON(http.StatusBadRequest, "squad is not full"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	squadCardsSecondClient, err := service.clubs.ListSquadCards(ctx, secondClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}
	if len(squadCardsSecondClient) != clubs.SquadSize {
		if err := secondClient.WriteJSON(http.StatusBadRequest, "squad is not full"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	firstClientSquad, err := service.clubs.GetSquad(ctx, firstClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	firstClientClub, err := service.clubs.Get(ctx, firstClientSquad.ClubID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	season, err := service.seasons.GetSeasonByDivisionID(ctx, firstClientClub.DivisionID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not season id"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	matchesID, err := service.matches.Create(ctx, firstClient.SquadID, secondClient.SquadID, firstClient.UserID, secondClient.UserID, season.ID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	matchRepresentation, err := gameengine.PrepareMatchRepresentation(service.config.CoordinateConfig, service.config.GameConfig, squadCardsFirstClient, squadCardsSecondClient)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, matchRepresentation); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, matchRepresentation); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	if err := firstClient.WriteJSON(http.StatusOK, matchRepresentation); err != nil {
		return ChoreError.Wrap(err)
	}
	if err := secondClient.WriteJSON(http.StatusOK, matchRepresentation); err != nil {
		return ChoreError.Wrap(err)
	}

	for i := 1; i < service.config.NumberOfRounds; i++ {
		if i == service.config.NumberOfRounds/2+1 {
			matchRepresentation.User1CardsWithPosition = gameengine.ReflectCoordinates(matchRepresentation.User1CardsWithPosition, service.config.CoordinateConfig.SizeOfFieldByOX, service.config.CoordinateConfig.SizeOfFieldByOY)
			matchRepresentation.User2CardsWithPosition = gameengine.ReflectCoordinates(matchRepresentation.User2CardsWithPosition, service.config.CoordinateConfig.SizeOfFieldByOX, service.config.CoordinateConfig.SizeOfFieldByOY)
			matchRepresentation.BallCoordinate = gameengine.GetCenterOfField(service.config.CoordinateConfig.SizeOfFieldByOX, service.config.CoordinateConfig.SizeOfFieldByOY)
			matchRepresentation.CardAvailableAction, err = gameengine.GenerateAvailableActions(matchRepresentation.User1CardsWithPosition, matchRepresentation.User2CardsWithPosition, matchRepresentation.BallCoordinate, service.config.GameConfig, service.config.CoordinateConfig)
			if err != nil {
				if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				return ChoreError.Wrap(err)
			}
			firstSquadCardAvailableActions, err := gameengine.GenerateAvailableActions(matchRepresentation.User1CardsWithPosition, matchRepresentation.User2CardsWithPosition,
				matchRepresentation.BallCoordinate, service.config.GameConfig, service.config.CoordinateConfig)
			if err != nil {
				if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				return ChoreError.Wrap(err)
			}
			secondSquadCardAvailableActions, err := gameengine.GenerateAvailableActions(matchRepresentation.User2CardsWithPosition, matchRepresentation.User1CardsWithPosition,
				matchRepresentation.BallCoordinate, service.config.GameConfig, service.config.CoordinateConfig)
			if err != nil {
				if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				return ChoreError.Wrap(err)
			}

			var cardsAvailableActions []gameengine.CardAvailableAction

			cardsAvailableActions = append(cardsAvailableActions, firstSquadCardAvailableActions...)
			cardsAvailableActions = append(cardsAvailableActions, secondSquadCardAvailableActions...)

			matchRepresentation.CardAvailableAction = cardsAvailableActions
		}

		ticker := time.NewTicker(service.config.MatchActionRenewalInterval)
		done := make(chan bool)
		go func() {
			time.Sleep(service.config.RoundDuration)
			done <- true
		}()

		var firstPlayerActions []gameengine.MakeAction
		var secondPlayerActions []gameengine.MakeAction

	loop:
		for {
			select {
			case <-ticker.C:
				firstPlayerActions, err = firstClient.ReadActionJSON()
				if err != nil {
					return ErrQueue.Wrap(err)
				}

				secondPlayerActions, err = secondClient.ReadActionJSON()
				if err != nil {
					return ErrQueue.Wrap(err)
				}

				if len(firstPlayerActions) != 0 && len(secondPlayerActions) != 0 {
					continue
				}
			case <-done:
				firstPlayerActions, err = firstClient.ReadActionJSON()
				if err != nil {
					return ErrQueue.Wrap(err)
				}

				if len(firstPlayerActions) == 0 {
					if err := firstClient.WriteJSON(http.StatusBadRequest, "invalid numbers of actions"); err != nil {
						return ChoreError.Wrap(err)
					}
				}

				secondPlayerActions, err = secondClient.ReadActionJSON()
				if err != nil {
					return ErrQueue.Wrap(err)
				}
				if len(secondPlayerActions) == 0 {
					if err := secondClient.WriteJSON(http.StatusBadRequest, "invalid numbers of actions"); err != nil {
						return ChoreError.Wrap(err)
					}
				}
				break loop
			}
		}
		if len(firstPlayerActions)+len(secondPlayerActions) > 0 {
			var actions []gameengine.MakeAction
			_ = copy(actions, firstPlayerActions)
			actions = append(actions, secondPlayerActions...)

			matchRepresentation.Actions, err = gameengine.HandleAction(matchRepresentation, actions, service.config.GameConfig, service.config.CoordinateConfig)
			if err != nil {
				if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not handle actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not handle actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				return ChoreError.Wrap(err)
			}

			firstSquadCardAvailableActions, err := gameengine.GenerateAvailableActions(matchRepresentation.User1CardsWithPosition, matchRepresentation.User2CardsWithPosition,
				matchRepresentation.BallCoordinate, service.config.GameConfig, service.config.CoordinateConfig)
			if err != nil {
				if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				return ChoreError.Wrap(err)
			}
			secondSquadCardAvailableActions, err := gameengine.GenerateAvailableActions(matchRepresentation.User2CardsWithPosition, matchRepresentation.User1CardsWithPosition,
				matchRepresentation.BallCoordinate, service.config.GameConfig, service.config.CoordinateConfig)
			if err != nil {
				if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not generate possible actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				return ChoreError.Wrap(err)
			}

			var cardsAvailableActions []gameengine.CardAvailableAction

			cardsAvailableActions = append(cardsAvailableActions, firstSquadCardAvailableActions...)
			cardsAvailableActions = append(cardsAvailableActions, secondSquadCardAvailableActions...)

			matchRepresentation.CardAvailableAction = cardsAvailableActions

			if i != service.config.NumberOfRounds/2 {
				if err = firstClient.WriteJSON(http.StatusOK, matchRepresentation); err != nil {
					return ChoreError.Wrap(err)
				}
				if err = secondClient.WriteJSON(http.StatusOK, matchRepresentation); err != nil {
					return ChoreError.Wrap(err)
				}
			}
		}
	}

	gameResult, err := service.matches.GetGameResult(ctx, matchesID)
	if err != nil {
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not get result of match"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	var firstClientResult matches.GameResult
	var secondClientResult matches.GameResult

	firstClientResult.MatchResults = make([]matches.MatchResult, len(gameResult.MatchResults))
	_ = copy(firstClientResult.MatchResults, gameResult.MatchResults)
	secondClientResult.MatchResults = make([]matches.MatchResult, len(gameResult.MatchResults))
	_ = copy(secondClientResult.MatchResults, gameResult.MatchResults)

	switch {
	case firstClient.UserID == gameResult.MatchResults[0].UserID:
		secondClientResult.MatchResults = matches.Swap(gameResult.MatchResults)
	case secondClient.UserID == gameResult.MatchResults[0].UserID:
		firstClientResult.MatchResults = matches.Swap(gameResult.MatchResults)
	}

	var value = new(big.Int)
	value.SetString(service.config.WinValue, 10)

	switch {
	case firstClientResult.MatchResults[0].QuantityGoals > secondClientResult.MatchResults[0].QuantityGoals:
		winResult := WinResult{
			Client:     firstClient,
			GameResult: firstClientResult,
			Value:      value,
		}

		go func() {
			err = service.FinishWithWinResult(ctx, winResult)
			if err != nil {
				service.log.Error(fmt.Sprintf("could not finish with result for user %v", winResult.Client.UserID), ErrQueue.Wrap(err))
			}
		}()
		go func() {
			err = service.FinishGame(secondClient, secondClientResult)
			if err != nil {
				service.log.Error(fmt.Sprintf("could not finish game for user %v", secondClient.UserID), ErrQueue.Wrap(err))
			}
		}()
	case firstClientResult.MatchResults[0].QuantityGoals < secondClientResult.MatchResults[0].QuantityGoals:
		winResult := WinResult{
			Client:     secondClient,
			GameResult: secondClientResult,
			Value:      value,
		}

		go func() {
			err = service.FinishWithWinResult(ctx, winResult)
			if err != nil {
				service.log.Error(fmt.Sprintf("could not finish with result for user %v", winResult.Client.UserID), ErrQueue.Wrap(err))
			}
		}()
		go func() {
			err = service.FinishGame(firstClient, firstClientResult)
			if err != nil {
				service.log.Error(fmt.Sprintf("could not finish game for user %v", firstClient.UserID), ErrQueue.Wrap(err))
			}
		}()
	default:
		var value = new(big.Int)
		value.SetString(service.config.DrawValue, 10)

		winResult := WinResult{
			Client:     firstClient,
			GameResult: firstClientResult,
			Value:      value,
		}
		go func() {
			err = service.FinishWithWinResult(ctx, winResult)
			if err != nil {
				service.log.Error(fmt.Sprintf("could not finish with result for user %v", winResult.Client.UserID), ErrQueue.Wrap(err))
			}
		}()

		winResult = WinResult{
			Client:     secondClient,
			GameResult: secondClientResult,
			Value:      value,
		}
		go func() {
			err = service.FinishWithWinResult(ctx, winResult)
			if err != nil {
				service.log.Error(fmt.Sprintf("could not finish with result for user %v", winResult.Client.UserID), ErrQueue.Wrap(err))
			}
		}()
	}

	return nil
}

// FinishWithWinResult sends win result and finishes the connection.
func (service *Service) FinishWithWinResult(ctx context.Context, winResult WinResult) error {
	user, err := service.users.Get(ctx, winResult.Client.UserID)
	if err != nil {
		return ErrQueue.Wrap(err)
	}

	winResult.GameResult.Question = "do you allow us to take your address?"
	winResult.GameResult.Transaction.Value = evmsignature.WeiToEthereum(winResult.Value).String()
	winResult.GameResult.Transaction.UDTContract.Address = service.config.UDTContract.Address
	if err := winResult.Client.WriteJSON(http.StatusOK, winResult.GameResult); err != nil {
		return ErrQueue.Wrap(err)
	}

	request, err := winResult.Client.ReadJSON()
	if err != nil {
		return ErrQueue.Wrap(err)
	}

	if request.Action != ActionForbidAddress && request.Action != ActionAllowAddress {
		if err := winResult.Client.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
			return ErrQueue.Wrap(err)
		}
	}

	if request.Action == ActionAllowAddress {
		if !request.WalletAddress.IsValidAddress() {
			if err := winResult.Client.WriteJSON(http.StatusBadRequest, "invalid address of user's wallet"); err != nil {
				return ErrQueue.Wrap(err)
			}
		}

		if err = service.users.UpdateWalletAddress(ctx, request.WalletAddress, winResult.Client.UserID); err != nil {
			if !users.ErrWalletAddressAlreadyInUse.Has(err) {
				return ErrQueue.Wrap(err)
			}
		}

		if winResult.GameResult.Transaction, err = service.currencywaitlist.Create(ctx, user.ID, *winResult.Value, request.Nonce); err != nil {
			return ErrQueue.Wrap(err)
		}
	}

	return service.FinishGame(winResult.Client, winResult.GameResult)
}

// FinishGame sends result and finishes the connection.
func (service *Service) FinishGame(client Client, gameResult matches.GameResult) error {
	var err error

	if err = client.WriteJSON(http.StatusOK, gameResult); err != nil {
		return ErrQueue.Wrap(err)
	}

	if err = service.Finish(client.UserID); err != nil {
		return ErrQueue.Wrap(err)
	}

	defer func() {
		if err = client.Connection.Close(); err != nil {
			service.log.Error("could not close websocket", ChoreError.Wrap(err))
			return
		}
	}()

	return nil
}
