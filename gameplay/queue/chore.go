// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queue

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/gameplay/matches"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/pkg/sync"
	"ultimatedivision/seasons"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

var (
	// ChoreError represents place chore error type.
	ChoreError = errs.Class("expiration place chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	config           Config
	log              logger.Logger
	service          *Service
	Loop             *sync.Cycle
	matches          *matches.Service
	seasons          *seasons.Service
	clubs            *clubs.Service
	currencywaitlist *currencywaitlist.Service
	users            *users.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, log logger.Logger, service *Service, matches *matches.Service, seasons *seasons.Service, clubs *clubs.Service, currencywaitlist *currencywaitlist.Service, users *users.Service) *Chore {
	return &Chore{
		config:           config,
		log:              log,
		service:          service,
		Loop:             sync.NewCycle(config.PlaceRenewalInterval),
		matches:          matches,
		seasons:          seasons,
		clubs:            clubs,
		currencywaitlist: currencywaitlist,
		users:            users,
	}
}

// Run starts the chore for re-check the expiration time of the token.
func (chore *Chore) Run(ctx context.Context) (err error) {
	firstRequestChan := make(chan Request)
	secondRequestChan := make(chan Request)

	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		clients := chore.service.ListNotPlayingUsers()

		if len(clients) >= 2 {
			for k := range clients {
				isEvenNumber := (k%2 != 1)
				if isEvenNumber {
					continue
				}

				go func(clients []Client, k int) {
					firstClient := clients[k-1]
					secondClient := clients[k]

					if err = chore.service.UpdateIsPlaying(firstClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
					}
					if err = chore.service.UpdateIsPlaying(secondClient.UserID, true); err != nil {
						chore.log.Error("could not update is play", ChoreError.Wrap(err))
					}

					if err := firstClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
					}
					if err := secondClient.WriteJSON(http.StatusOK, "you confirm play?"); err != nil {
						chore.log.Error("could not write json", ChoreError.Wrap(err))
					}

					go func() {
						request, err := firstClient.ReadJSON()
						if err != nil {
							chore.log.Error("could not read json", ChoreError.Wrap(err))
						}
						firstRequestChan <- request
					}()

					go func() {
						request, err := secondClient.ReadJSON()
						if err != nil {
							chore.log.Error("could not read json", ChoreError.Wrap(err))
						}
						secondRequestChan <- request
					}()

					var firstRequest, secondRequest Request
					for {
						select {
						case firstRequest = <-firstRequestChan:
							if (firstRequest != Request{}) {
								if firstRequest.Action != ActionConfirm && firstRequest.Action != ActionReject {
									if err := firstClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
										chore.log.Error("could not write json", ChoreError.Wrap(err))
									}

									if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									return
								}
							}
						case secondRequest = <-secondRequestChan:
							if (secondRequest != Request{}) {
								if secondRequest.Action != ActionConfirm && secondRequest.Action != ActionReject {
									if err := secondClient.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
										chore.log.Error("could not write json", ChoreError.Wrap(err))
									}

									if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
										chore.log.Error("could not update is play", ChoreError.Wrap(err))
									}
									return
								}
							}
						}

						if (firstRequest == Request{} && secondRequest == Request{}) {
							continue
						}

						if firstRequest.Action == ActionReject || secondRequest.Action == ActionReject {
							if err := firstClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
							}
							if err := secondClient.WriteJSON(http.StatusOK, "you are still in search!"); err != nil {
								chore.log.Error("could not write json", ChoreError.Wrap(err))
							}

							if err = chore.service.Finish(firstClient.UserID); err != nil {
								chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
							}
							if err = chore.service.Finish(secondClient.UserID); err != nil {
								chore.log.Error("could not delete client from queue", ChoreError.Wrap(err))
							}
							return
						}

						if (firstRequest == Request{} || secondRequest == Request{}) {
							continue
						}

						if err = chore.Play(ctx, firstClient, secondClient); err != nil {
							if err = chore.service.UpdateIsPlaying(firstClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							if err = chore.service.UpdateIsPlaying(secondClient.UserID, false); err != nil {
								chore.log.Error("could not update is play", ChoreError.Wrap(err))
							}
							chore.log.Error("could not play game", ChoreError.Wrap(err))
						}
						return
					}
				}(clients, k)
			}
		}
		return ChoreError.Wrap(err)
	})
}

// Play method contains all the logic for playing matches.
func (chore *Chore) Play(ctx context.Context, firstClient, secondClient Client) error {
	squadCardsFirstClient, err := chore.service.clubs.ListSquadCards(ctx, firstClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}
	if len(squadCardsFirstClient) != clubs.SquadSize {
		if err := firstClient.WriteJSON(http.StatusBadRequest, "squad is not full"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	firstClientCardsWithPositions, err := chore.matches.ConvertPositionsForGameplay(ctx, squadCardsFirstClient)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	ballPosition := chore.matches.GenerateBallPosition()

	var matchResponse matches.GetMatchResponse

	firstClientPossibleActions := chore.matches.GenerateActionsForCards(ctx, firstClientCardsWithPositions, ballPosition)

	firstPlayer := matches.GetMatchPlayerResponse{
		UserID:          firstClient.UserID,
		SquadCards:      firstClientCardsWithPositions,
		PossibleActions: firstClientPossibleActions,
	}

	squadCardsSecondClient, err := chore.service.clubs.ListSquadCards(ctx, secondClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}
	if len(squadCardsSecondClient) != clubs.SquadSize {
		if err := secondClient.WriteJSON(http.StatusBadRequest, "squad is not full"); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	secondClientCardsWithPositions, err := chore.matches.ConvertPositionsForGameplay(ctx, squadCardsSecondClient)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	// reflect position for second client.
	secondClientCardsWithPositions = chore.matches.ReflectPositions(ctx, secondClientCardsWithPositions)

	secondClientPossibleActions := chore.matches.GenerateActionsForCards(ctx, secondClientCardsWithPositions, ballPosition)

	secondPlayer := matches.GetMatchPlayerResponse{
		UserID:          secondClient.UserID,
		SquadCards:      secondClientCardsWithPositions,
		PossibleActions: secondClientPossibleActions,
	}

	matchResponse.UserSquads = []matches.GetMatchPlayerResponse{firstPlayer, secondPlayer}
	matchResponse.BallPosition = ballPosition

	// TODO: add slice with positions that will be in the center with ball.

	if err := firstClient.WriteJSON(http.StatusOK, matchResponse); err != nil {
		return ChoreError.Wrap(err)
	}
	if err := secondClient.WriteJSON(http.StatusOK, matchResponse); err != nil {
		return ChoreError.Wrap(err)
	}

	firstClientSquad, err := chore.clubs.GetSquad(ctx, firstClient.SquadID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	firstClientClub, err := chore.clubs.Get(ctx, firstClientSquad.ClubID)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	season, err := chore.seasons.GetSeasonByDivisionID(ctx, firstClientClub.DivisionID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "could not get season id"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "could not get season id"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	matchesID, err := chore.matches.Create(ctx, firstClient.SquadID, secondClient.SquadID, firstClient.UserID, secondClient.UserID, season.ID)
	if err != nil {
		if err := firstClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
		if err := secondClient.WriteJSON(http.StatusInternalServerError, "match error"); err != nil {
			return ChoreError.Wrap(err)
		}
		return ChoreError.Wrap(err)
	}

	for i := 1; i < chore.config.NumberOfRounds; i++ {
		if i == chore.config.NumberOfRounds/2+1 {
			ballPosition = chore.matches.GenerateBallPosition()

			firstPlayer = matches.GetMatchPlayerResponse{
				UserID:     firstClient.UserID,
				SquadCards: chore.matches.ReflectPositions(ctx, firstPlayer.SquadCards),
				Movements:  matchResponse.Movements,
			}

			firstPlayer.PossibleActions = chore.matches.GenerateActionsForCards(ctx, firstPlayer.SquadCards, ballPosition)

			secondPlayer = matches.GetMatchPlayerResponse{
				UserID:          secondClient.UserID,
				SquadCards:      chore.matches.ReflectPositions(ctx, secondPlayer.SquadCards),
				PossibleActions: secondClientPossibleActions,
				Movements:       matchResponse.Movements,
			}

			secondPlayer.PossibleActions = chore.matches.GenerateActionsForCards(ctx, secondPlayer.SquadCards, ballPosition)

			matchResponse.UserSquads = []matches.GetMatchPlayerResponse{firstPlayer, secondPlayer}
			matchResponse.BallPosition = ballPosition

			if err := firstClient.WriteJSON(http.StatusOK, matchResponse); err != nil {
				return ChoreError.Wrap(err)
			}
			if err := secondClient.WriteJSON(http.StatusOK, matchResponse); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		ticker := time.NewTicker(chore.config.MatchActionRenewalInterval)
		done := make(chan bool)
		go func() {
			time.Sleep(chore.config.RoundDuration)
			done <- true
		}()

		var firstPlayerActions []matches.ActionRequest
		var secondPlayerActions []matches.ActionRequest
	Loop:
		for {
			select {
			case <-ticker.C:
				if len(firstPlayerActions) == 0 {
					firstPlayerActions, err = firstClient.ReadActionJSON()
					if err != nil {
						chore.log.Error("could not read json", ChoreError.Wrap(err))
					}
				}

				if len(secondPlayerActions) == 0 {
					secondPlayerActions, err = secondClient.ReadActionJSON()
					if err != nil {
						chore.log.Error("could not read json", ChoreError.Wrap(err))
					}
				}

				if len(firstPlayerActions) != 0 && len(secondPlayerActions) != 0 {
					break Loop
				}
			case <-done:
				firstPlayerActions, err = firstClient.ReadActionJSON()
				if err != nil {
					chore.log.Error("could not read json", ChoreError.Wrap(err))
				}

				if len(firstPlayerActions) == 0 {
					if err := firstClient.WriteJSON(http.StatusBadRequest, "invalid numbers of actions"); err != nil {
						return ChoreError.Wrap(err)
					}
				}

				secondPlayerActions, err = secondClient.ReadActionJSON()
				if err != nil {
					chore.log.Error("could not read json", ChoreError.Wrap(err))
				}
				if len(secondPlayerActions) == 0 {
					if err := secondClient.WriteJSON(http.StatusBadRequest, "invalid numbers of actions"); err != nil {
						return ChoreError.Wrap(err)
					}
				}
				break Loop
			}
		}
		if len(firstPlayerActions)+len(secondPlayerActions) > 0 {
			var actions []matches.ActionRequest
			_ = copy(actions, firstPlayerActions)
			actions = append(actions, secondPlayerActions...)

			matchResponse, err = chore.matches.HandleActions(ctx, actions, matchResponse)
			if err != nil {
				if err = firstClient.WriteJSON(http.StatusBadRequest, "could not handle actions"); err != nil {
					return ChoreError.Wrap(err)
				}
				if err = secondClient.WriteJSON(http.StatusBadRequest, "could not handle actions"); err != nil {
					return ChoreError.Wrap(err)
				}
			}

			for _, userSquad := range matchResponse.UserSquads {
				userSquad.PossibleActions = chore.matches.GenerateActionsForCards(ctx, firstClientCardsWithPositions, ballPosition)
			}

			if i != chore.config.NumberOfRounds/2 {
				if err = firstClient.WriteJSON(http.StatusOK, matchResponse); err != nil {
					return ChoreError.Wrap(err)
				}
				if err = secondClient.WriteJSON(http.StatusOK, matchResponse); err != nil {
					return ChoreError.Wrap(err)
				}
			}
		}
	}

	gameResult, err := chore.matches.GetGameResult(ctx, matchesID)
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
	value.SetString(chore.config.WinValue, 10)

	switch {
	case firstClientResult.MatchResults[0].QuantityGoals > secondClientResult.MatchResults[0].QuantityGoals:
		winResult := WinResult{
			Client:     firstClient,
			GameResult: firstClientResult,
			Value:      value,
		}

		go chore.FinishWithWinResult(ctx, winResult)
		go chore.Finish(secondClient, secondClientResult)
	case firstClientResult.MatchResults[0].QuantityGoals < secondClientResult.MatchResults[0].QuantityGoals:
		winResult := WinResult{
			Client:     secondClient,
			GameResult: secondClientResult,
			Value:      value,
		}

		go chore.FinishWithWinResult(ctx, winResult)
		go chore.Finish(firstClient, firstClientResult)
	default:
		var value = new(big.Int)
		value.SetString(chore.config.DrawValue, 10)

		winResult := WinResult{
			Client:     firstClient,
			GameResult: firstClientResult,
			Value:      value,
		}
		go chore.FinishWithWinResult(ctx, winResult)

		winResult = WinResult{
			Client:     secondClient,
			GameResult: secondClientResult,
			Value:      value,
		}
		go chore.FinishWithWinResult(ctx, winResult)
	}

	return nil
}

// FinishWithWinResult sends win result and finishes the connection.
func (chore *Chore) FinishWithWinResult(ctx context.Context, winResult WinResult) {
	user, err := chore.users.Get(ctx, winResult.Client.UserID)
	if err != nil {
		chore.log.Error("could not get user", ChoreError.Wrap(err))
		return
	}

	winResult.GameResult.Question = "do you allow us to take your address?"
	winResult.GameResult.Transaction.Value = cryptoutils.WeiToEthereum(winResult.Value).String()
	winResult.GameResult.Transaction.UDTContract.Address = chore.config.UDTContract.Address
	if err := winResult.Client.WriteJSON(http.StatusOK, winResult.GameResult); err != nil {
		chore.log.Error("could not write json", ChoreError.Wrap(err))
		return
	}

	request, err := winResult.Client.ReadJSON()
	if err != nil {
		chore.log.Error("could not read json", ChoreError.Wrap(err))
		return
	}

	if request.Action != ActionForbidAddress && request.Action != ActionAllowAddress {
		if err := winResult.Client.WriteJSON(http.StatusBadRequest, "wrong action"); err != nil {
			chore.log.Error("could not write json", ChoreError.Wrap(err))
			return
		}
	}

	if request.Action == ActionAllowAddress {
		if !request.WalletAddress.IsValidAddress() {
			if err := winResult.Client.WriteJSON(http.StatusBadRequest, "invalid address of user's wallet"); err != nil {
				chore.log.Error("could not write json", ChoreError.Wrap(err))
				return
			}
		}

		if err = chore.users.UpdateWalletAddress(ctx, request.WalletAddress, winResult.Client.UserID); err != nil {
			if !users.ErrWalletAddressAlreadyInUse.Has(err) {
				chore.log.Error("could not update user's wallet address", ChoreError.Wrap(err))
				return
			}
		}

		if winResult.GameResult.Transaction, err = chore.currencywaitlist.Create(ctx, user.ID, *winResult.Value, request.Nonce); err != nil {
			chore.log.Error("could not create item of currencywaitlist", ChoreError.Wrap(err))
			return
		}
	}
	chore.Finish(winResult.Client, winResult.GameResult)
}

// Finish sends result and finishes the connection.
func (chore *Chore) Finish(client Client, gameResult matches.GameResult) {
	var err error

	if err = client.WriteJSON(http.StatusOK, gameResult); err != nil {
		chore.log.Error("could not write json", ChoreError.Wrap(err))
		return
	}

	if err = chore.service.Finish(client.UserID); err != nil {
		chore.log.Error("could not finish match", ChoreError.Wrap(err))
		return
	}
	defer func() {
		if err = client.Connection.Close(); err != nil {
			chore.log.Error("could not close websocket", ChoreError.Wrap(err))
		}
	}()
}

// Close closes the chore for re-check the expiration time of the token.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
